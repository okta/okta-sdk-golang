# PasswordPolicyPasswordSettingsComplexity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dictionary** | Pointer to [**PasswordDictionary**](PasswordDictionary.md) |  | [optional] 
**ExcludeAttributes** | Pointer to **[]string** |  | [optional] 
**ExcludeUsername** | Pointer to **bool** |  | [optional] [default to true]
**MinLength** | Pointer to **int32** |  | [optional] 
**MinLowerCase** | Pointer to **int32** |  | [optional] 
**MinNumber** | Pointer to **int32** |  | [optional] 
**MinSymbol** | Pointer to **int32** |  | [optional] 
**MinUpperCase** | Pointer to **int32** |  | [optional] 

## Methods

### NewPasswordPolicyPasswordSettingsComplexity

`func NewPasswordPolicyPasswordSettingsComplexity() *PasswordPolicyPasswordSettingsComplexity`

NewPasswordPolicyPasswordSettingsComplexity instantiates a new PasswordPolicyPasswordSettingsComplexity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyPasswordSettingsComplexityWithDefaults

`func NewPasswordPolicyPasswordSettingsComplexityWithDefaults() *PasswordPolicyPasswordSettingsComplexity`

NewPasswordPolicyPasswordSettingsComplexityWithDefaults instantiates a new PasswordPolicyPasswordSettingsComplexity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDictionary

`func (o *PasswordPolicyPasswordSettingsComplexity) GetDictionary() PasswordDictionary`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *PasswordPolicyPasswordSettingsComplexity) GetDictionaryOk() (*PasswordDictionary, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *PasswordPolicyPasswordSettingsComplexity) SetDictionary(v PasswordDictionary)`

SetDictionary sets Dictionary field to given value.

### HasDictionary

`func (o *PasswordPolicyPasswordSettingsComplexity) HasDictionary() bool`

HasDictionary returns a boolean if a field has been set.

### GetExcludeAttributes

`func (o *PasswordPolicyPasswordSettingsComplexity) GetExcludeAttributes() []string`

GetExcludeAttributes returns the ExcludeAttributes field if non-nil, zero value otherwise.

### GetExcludeAttributesOk

`func (o *PasswordPolicyPasswordSettingsComplexity) GetExcludeAttributesOk() (*[]string, bool)`

GetExcludeAttributesOk returns a tuple with the ExcludeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttributes

`func (o *PasswordPolicyPasswordSettingsComplexity) SetExcludeAttributes(v []string)`

SetExcludeAttributes sets ExcludeAttributes field to given value.

### HasExcludeAttributes

`func (o *PasswordPolicyPasswordSettingsComplexity) HasExcludeAttributes() bool`

HasExcludeAttributes returns a boolean if a field has been set.

### GetExcludeUsername

`func (o *PasswordPolicyPasswordSettingsComplexity) GetExcludeUsername() bool`

GetExcludeUsername returns the ExcludeUsername field if non-nil, zero value otherwise.

### GetExcludeUsernameOk

`func (o *PasswordPolicyPasswordSettingsComplexity) GetExcludeUsernameOk() (*bool, bool)`

GetExcludeUsernameOk returns a tuple with the ExcludeUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsername

`func (o *PasswordPolicyPasswordSettingsComplexity) SetExcludeUsername(v bool)`

SetExcludeUsername sets ExcludeUsername field to given value.

### HasExcludeUsername

`func (o *PasswordPolicyPasswordSettingsComplexity) HasExcludeUsername() bool`

HasExcludeUsername returns a boolean if a field has been set.

### GetMinLength

`func (o *PasswordPolicyPasswordSettingsComplexity) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PasswordPolicyPasswordSettingsComplexity) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PasswordPolicyPasswordSettingsComplexity) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *PasswordPolicyPasswordSettingsComplexity) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMinLowerCase

`func (o *PasswordPolicyPasswordSettingsComplexity) GetMinLowerCase() int32`

GetMinLowerCase returns the MinLowerCase field if non-nil, zero value otherwise.

### GetMinLowerCaseOk

`func (o *PasswordPolicyPasswordSettingsComplexity) GetMinLowerCaseOk() (*int32, bool)`

GetMinLowerCaseOk returns a tuple with the MinLowerCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLowerCase

`func (o *PasswordPolicyPasswordSettingsComplexity) SetMinLowerCase(v int32)`

SetMinLowerCase sets MinLowerCase field to given value.

### HasMinLowerCase

`func (o *PasswordPolicyPasswordSettingsComplexity) HasMinLowerCase() bool`

HasMinLowerCase returns a boolean if a field has been set.

### GetMinNumber

`func (o *PasswordPolicyPasswordSettingsComplexity) GetMinNumber() int32`

GetMinNumber returns the MinNumber field if non-nil, zero value otherwise.

### GetMinNumberOk

`func (o *PasswordPolicyPasswordSettingsComplexity) GetMinNumberOk() (*int32, bool)`

GetMinNumberOk returns a tuple with the MinNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumber

`func (o *PasswordPolicyPasswordSettingsComplexity) SetMinNumber(v int32)`

SetMinNumber sets MinNumber field to given value.

### HasMinNumber

`func (o *PasswordPolicyPasswordSettingsComplexity) HasMinNumber() bool`

HasMinNumber returns a boolean if a field has been set.

### GetMinSymbol

`func (o *PasswordPolicyPasswordSettingsComplexity) GetMinSymbol() int32`

GetMinSymbol returns the MinSymbol field if non-nil, zero value otherwise.

### GetMinSymbolOk

`func (o *PasswordPolicyPasswordSettingsComplexity) GetMinSymbolOk() (*int32, bool)`

GetMinSymbolOk returns a tuple with the MinSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSymbol

`func (o *PasswordPolicyPasswordSettingsComplexity) SetMinSymbol(v int32)`

SetMinSymbol sets MinSymbol field to given value.

### HasMinSymbol

`func (o *PasswordPolicyPasswordSettingsComplexity) HasMinSymbol() bool`

HasMinSymbol returns a boolean if a field has been set.

### GetMinUpperCase

`func (o *PasswordPolicyPasswordSettingsComplexity) GetMinUpperCase() int32`

GetMinUpperCase returns the MinUpperCase field if non-nil, zero value otherwise.

### GetMinUpperCaseOk

`func (o *PasswordPolicyPasswordSettingsComplexity) GetMinUpperCaseOk() (*int32, bool)`

GetMinUpperCaseOk returns a tuple with the MinUpperCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpperCase

`func (o *PasswordPolicyPasswordSettingsComplexity) SetMinUpperCase(v int32)`

SetMinUpperCase sets MinUpperCase field to given value.

### HasMinUpperCase

`func (o *PasswordPolicyPasswordSettingsComplexity) HasMinUpperCase() bool`

HasMinUpperCase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


