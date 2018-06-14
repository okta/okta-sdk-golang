package conjungo

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type Options struct {
	// Overwrite a target value with source value even if it already exists
	Overwrite bool

	// Unexported fields on a struct can not be set. When a struct contains an unexported
	// field, the default behavior is to treat the entire struct as a single entity and
	// replace according to Overwrite settings. If this is enabled, an error will be thrown instead.
	//
	// Note: this is used by the default mergeStruct function, and may not apply if that is
	// overwritten with a custom function. Custom struct merge functions should consider
	// using this value as well.
	ErrorOnUnexported bool

	// A set of default and customizable functions that define how values are merged
	MergeFuncs *funcSelector

	// To be used by merge functions to pass values down into recursive calls freely
	Context context.Context
}

func NewOptions() *Options {
	return &Options{
		Overwrite:  true,
		MergeFuncs: newFuncSelector(),
	}
}

// public wrapper
func Merge(target, source interface{}, opt *Options) error {
	vT := reflect.ValueOf(target)
	vS := reflect.ValueOf(source)

	if vT.Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
	}

	if !reflect.Indirect(vT).IsValid() {
		return errors.New("target can not be zero value")
	}

	// use defaults if none are provided
	if opt == nil {
		opt = NewOptions()
	}

	//make a copy here so if there is an error mid way, the target stays in tact
	cp := vT.Elem()

	//TODO reflect.Indirect(vS)?
	merged, err := merge(cp, vS, opt)
	if err != nil {
		return err
	}

	if !isSettable(vT.Elem(), merged) {
		return fmt.Errorf("Merge failed: expected merged result to be %v but got %v",
			vT.Elem().Type(), merged.Type())
	}

	vT.Elem().Set(merged)
	return nil
}

func isSettable(t, s reflect.Value) bool {
	if t.Kind() != reflect.Interface && t.Type() != s.Type() {
		return false
	}

	return true
}

func merge(valT, valS reflect.Value, opt *Options) (reflect.Value, error) {
	// if source is nil, skip
	if !valS.IsValid() ||
		valS.Kind() == reflect.Ptr && valS.IsNil() {
		return valT, nil
	}

	// if target is nil write to it
	if !valT.IsValid() ||
		valT.Kind() == reflect.Ptr && valT.IsNil() {
		return valS, nil
	}

	// get to the real type
	if valT.Kind() == reflect.Interface || valS.Kind() == reflect.Interface {
		valT = reflect.ValueOf(valT.Interface())
		valS = reflect.ValueOf(valS.Interface())
	}

	// if types do not match, bail
	if valT.Type() != valS.Type() {
		return reflect.Value{}, fmt.Errorf("Types do not match: %v, %v", valT.Type(), valS.Type())
	}

	// look for a merge function
	f := opt.MergeFuncs.GetFunc(valT)
	val, err := f(valT, valS, opt)
	if err != nil {
		return reflect.Value{}, err
	}

	return val, nil
}
