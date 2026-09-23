# OSAccountDisplayName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sensitive** | Pointer to **bool** | Indicates whether the associated value is Personal Identifiable Information (PII) and requires masking | [optional] [default to false]
**Value** | Pointer to **string** | Display name of the OS account | [optional] 

## Methods

### NewOSAccountDisplayName

`func NewOSAccountDisplayName() *OSAccountDisplayName`

NewOSAccountDisplayName instantiates a new OSAccountDisplayName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSAccountDisplayNameWithDefaults

`func NewOSAccountDisplayNameWithDefaults() *OSAccountDisplayName`

NewOSAccountDisplayNameWithDefaults instantiates a new OSAccountDisplayName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSensitive

`func (o *OSAccountDisplayName) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *OSAccountDisplayName) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *OSAccountDisplayName) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *OSAccountDisplayName) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetValue

`func (o *OSAccountDisplayName) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OSAccountDisplayName) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OSAccountDisplayName) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OSAccountDisplayName) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


