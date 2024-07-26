# PasswordSettingObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Change** | Pointer to **string** | Determines whether a change in a user&#39;s password also updates the user&#39;s password in the application | [optional] [default to "KEEP_EXISTING"]
**Seed** | Pointer to **string** | Determines whether the generated password is the user&#39;s Okta password or a randomly generated password | [optional] [default to "RANDOM"]
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewPasswordSettingObject

`func NewPasswordSettingObject() *PasswordSettingObject`

NewPasswordSettingObject instantiates a new PasswordSettingObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordSettingObjectWithDefaults

`func NewPasswordSettingObjectWithDefaults() *PasswordSettingObject`

NewPasswordSettingObjectWithDefaults instantiates a new PasswordSettingObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChange

`func (o *PasswordSettingObject) GetChange() string`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *PasswordSettingObject) GetChangeOk() (*string, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *PasswordSettingObject) SetChange(v string)`

SetChange sets Change field to given value.

### HasChange

`func (o *PasswordSettingObject) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetSeed

`func (o *PasswordSettingObject) GetSeed() string`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *PasswordSettingObject) GetSeedOk() (*string, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *PasswordSettingObject) SetSeed(v string)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *PasswordSettingObject) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### GetStatus

`func (o *PasswordSettingObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PasswordSettingObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PasswordSettingObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PasswordSettingObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


