package conjungo

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Sirupsen/logrus"
)

type funcSelector struct {
	typeFuncs   map[reflect.Type]MergeFunc
	kindFuncs   map[reflect.Kind]MergeFunc
	defaultFunc MergeFunc
}

func newFuncSelector() *funcSelector {
	return &funcSelector{
		typeFuncs: map[reflect.Type]MergeFunc{},
		kindFuncs: map[reflect.Kind]MergeFunc{
			reflect.Map:    mergeMap,
			reflect.Slice:  mergeSlice,
			reflect.Struct: mergeStruct,
		},
		defaultFunc: defaultMergeFunc,
	}
}

func (f *funcSelector) SetTypeMergeFunc(t reflect.Type, mf MergeFunc) {
	if nil == f.typeFuncs {
		f.typeFuncs = map[reflect.Type]MergeFunc{}
	}
	f.typeFuncs[t] = mf
}

func (f *funcSelector) SetKindMergeFunc(k reflect.Kind, mf MergeFunc) {
	if nil == f.kindFuncs {
		f.kindFuncs = map[reflect.Kind]MergeFunc{}
	}
	f.kindFuncs[k] = mf
}

func (f *funcSelector) SetDefaultMergeFunc(mf MergeFunc) {
	f.defaultFunc = mf
}

// Get func must always return a function.
// First looks for a merge func defined for its type. Type is the most specific way to categorize something,
// for example, struct type foo of package bar or map[string]string. Next it looks for a merge func defined for its
// kind, for example, struct or map. At this point, if nothing matches, it will fall back to the default merge definition.
func (f *funcSelector) GetFunc(v reflect.Value) MergeFunc {
	// prioritize a specific 'type' definition
	ti := v.Type()

	if fx, ok := f.typeFuncs[ti]; ok {
		return fx
	}

	// then look for a more general 'kind'.
	if fx, ok := f.kindFuncs[ti.Kind()]; ok {
		return fx
	}

	if f.defaultFunc != nil {
		return f.defaultFunc
	}

	return defaultMergeFunc
}

// A function which defines how two items of the same type are merged together.
// Options are also passed in and it is the responsibility of the merge function to handle
// any variations in behavior that should occur. The value returned from the function will be
// written to directly to the target map, as long as there is no error.
type MergeFunc func(reflect.Value, reflect.Value, *Options) (reflect.Value, error)

// The most basic merge function to be used as default behavior. In overwrite mode, it returns the source. Otherwise,
// it returns the target.
func defaultMergeFunc(t, s reflect.Value, o *Options) (reflect.Value, error) {
	if o.Overwrite {
		return s, nil
	}

	return t, nil
}

func mergeMap(t, s reflect.Value, o *Options) (v reflect.Value, err error) {
	if t.Kind() != reflect.Map || s.Kind() != reflect.Map {
		return reflect.Value{}, fmt.Errorf("got non-map type (tagret: %v; source: %v)", t.Kind(), s.Kind())
	}

	keys := s.MapKeys()

	defer func() {
		if r := recover(); r != nil {
			vr := reflect.ValueOf(r)
			if vr.Kind() == reflect.String {
				//TODO: make this easier to debug
				err = errors.New("failed to merge map: " + r.(string))
			}
		}
	}()

	for _, k := range keys {
		logrus.Debugf("MERGE T<>S '%s' :: %v <> %v", k, t.MapIndex(k), s.MapIndex(k))
		val, err := merge(t.MapIndex(k), s.MapIndex(k), o)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("key '%s': %v", k, err)
		}
		t.SetMapIndex(k, val)
	}

	v = t
	return
}

// Merges two slices of the same type by appending source to target.
func mergeSlice(t, s reflect.Value, o *Options) (reflect.Value, error) {
	if t.Type() != s.Type() {
		return reflect.Value{}, fmt.Errorf("slices must have same type: T: %v S: %v", t.Type(), s.Type())
	}

	return reflect.AppendSlice(t, s), nil
}

// This func is designed to be called by merge().
// It should not be used on its own because it will panic.
func mergeStruct(t, s reflect.Value, o *Options) (reflect.Value, error) {
	// accept pointer values, but dereference them
	valT := reflect.Indirect(t)
	valS := reflect.Indirect(s)
	kindT := valT.Kind()
	kindS := valS.Kind()

	newT := reflect.New(valT.Type()).Elem() //a new instance of the struct type that can be set

	okT := kindT == reflect.Struct
	okS := kindS == reflect.Struct
	if !okT || !okS {
		return reflect.Value{}, fmt.Errorf("got non-struct kind (tagret: %v; source: %v)", kindT, kindS)
	}

	for i := 0; i < valS.NumField(); i++ {
		fieldT := newT.Field(i)
		logrus.Debug("merging struct field %s", fieldT)

		// field is addressable because it's created above. So this means it is unexported.
		if !fieldT.CanSet() {
			if o.ErrorOnUnexported {
				return reflect.Value{}, fmt.Errorf("struct of type %v has unexported field: %s",
					t.Type().Name(), newT.Type().Field(i).Name)
			}

			// revert to using the default func instead to treat the struct as single entity
			return defaultMergeFunc(t, s, o)
		}

		//fieldT should always be valid because it's created above
		merged, err := merge(valT.Field(i), valS.Field(i), o)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("failed to merge field `%s.%s`: %v",
				newT.Type().Name(), newT.Type().Field(i).Name, err)
		}

		if fieldT.Kind() != reflect.Interface && fieldT.Type() != merged.Type() {
			return reflect.Value{}, fmt.Errorf("types dont match %v <> %v", fieldT.Type(), merged.Type())
		}

		fieldT.Set(merged)
	}

	return newT, nil
}
