# UserIdentificationSecurityMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fastpass** | Pointer to [**UserIdentificationFastpassSettings**](UserIdentificationFastpassSettings.md) |  | [optional] 
**Passkey** | Pointer to [**UserIdentificationPasskeySettings**](UserIdentificationPasskeySettings.md) |  | [optional] 

## Methods

### NewUserIdentificationSecurityMethods

`func NewUserIdentificationSecurityMethods() *UserIdentificationSecurityMethods`

NewUserIdentificationSecurityMethods instantiates a new UserIdentificationSecurityMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentificationSecurityMethodsWithDefaults

`func NewUserIdentificationSecurityMethodsWithDefaults() *UserIdentificationSecurityMethods`

NewUserIdentificationSecurityMethodsWithDefaults instantiates a new UserIdentificationSecurityMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFastpass

`func (o *UserIdentificationSecurityMethods) GetFastpass() UserIdentificationFastpassSettings`

GetFastpass returns the Fastpass field if non-nil, zero value otherwise.

### GetFastpassOk

`func (o *UserIdentificationSecurityMethods) GetFastpassOk() (*UserIdentificationFastpassSettings, bool)`

GetFastpassOk returns a tuple with the Fastpass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastpass

`func (o *UserIdentificationSecurityMethods) SetFastpass(v UserIdentificationFastpassSettings)`

SetFastpass sets Fastpass field to given value.

### HasFastpass

`func (o *UserIdentificationSecurityMethods) HasFastpass() bool`

HasFastpass returns a boolean if a field has been set.

### GetPasskey

`func (o *UserIdentificationSecurityMethods) GetPasskey() UserIdentificationPasskeySettings`

GetPasskey returns the Passkey field if non-nil, zero value otherwise.

### GetPasskeyOk

`func (o *UserIdentificationSecurityMethods) GetPasskeyOk() (*UserIdentificationPasskeySettings, bool)`

GetPasskeyOk returns a tuple with the Passkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasskey

`func (o *UserIdentificationSecurityMethods) SetPasskey(v UserIdentificationPasskeySettings)`

SetPasskey sets Passkey field to given value.

### HasPasskey

`func (o *UserIdentificationSecurityMethods) HasPasskey() bool`

HasPasskey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


