# AAGUIDAuthenticatorCharacteristics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FipsCompliant** | Pointer to **bool** | Indicates whether the authenticator meets Federal Information Processing Standards (FIPS) compliance requirements | [optional] 
**HardwareProtected** | Pointer to **bool** | Indicates whether the authenticator stores the private key on a hardware component | [optional] 
**PlatformAttached** | Pointer to **bool** | Indicates whether the custom AAGUID is built into the authenticator (&#x60;true&#x60;) or if it&#39;s a separate, external authenticator | [optional] 

## Methods

### NewAAGUIDAuthenticatorCharacteristics

`func NewAAGUIDAuthenticatorCharacteristics() *AAGUIDAuthenticatorCharacteristics`

NewAAGUIDAuthenticatorCharacteristics instantiates a new AAGUIDAuthenticatorCharacteristics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAAGUIDAuthenticatorCharacteristicsWithDefaults

`func NewAAGUIDAuthenticatorCharacteristicsWithDefaults() *AAGUIDAuthenticatorCharacteristics`

NewAAGUIDAuthenticatorCharacteristicsWithDefaults instantiates a new AAGUIDAuthenticatorCharacteristics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFipsCompliant

`func (o *AAGUIDAuthenticatorCharacteristics) GetFipsCompliant() bool`

GetFipsCompliant returns the FipsCompliant field if non-nil, zero value otherwise.

### GetFipsCompliantOk

`func (o *AAGUIDAuthenticatorCharacteristics) GetFipsCompliantOk() (*bool, bool)`

GetFipsCompliantOk returns a tuple with the FipsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsCompliant

`func (o *AAGUIDAuthenticatorCharacteristics) SetFipsCompliant(v bool)`

SetFipsCompliant sets FipsCompliant field to given value.

### HasFipsCompliant

`func (o *AAGUIDAuthenticatorCharacteristics) HasFipsCompliant() bool`

HasFipsCompliant returns a boolean if a field has been set.

### GetHardwareProtected

`func (o *AAGUIDAuthenticatorCharacteristics) GetHardwareProtected() bool`

GetHardwareProtected returns the HardwareProtected field if non-nil, zero value otherwise.

### GetHardwareProtectedOk

`func (o *AAGUIDAuthenticatorCharacteristics) GetHardwareProtectedOk() (*bool, bool)`

GetHardwareProtectedOk returns a tuple with the HardwareProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareProtected

`func (o *AAGUIDAuthenticatorCharacteristics) SetHardwareProtected(v bool)`

SetHardwareProtected sets HardwareProtected field to given value.

### HasHardwareProtected

`func (o *AAGUIDAuthenticatorCharacteristics) HasHardwareProtected() bool`

HasHardwareProtected returns a boolean if a field has been set.

### GetPlatformAttached

`func (o *AAGUIDAuthenticatorCharacteristics) GetPlatformAttached() bool`

GetPlatformAttached returns the PlatformAttached field if non-nil, zero value otherwise.

### GetPlatformAttachedOk

`func (o *AAGUIDAuthenticatorCharacteristics) GetPlatformAttachedOk() (*bool, bool)`

GetPlatformAttachedOk returns a tuple with the PlatformAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAttached

`func (o *AAGUIDAuthenticatorCharacteristics) SetPlatformAttached(v bool)`

SetPlatformAttached sets PlatformAttached field to given value.

### HasPlatformAttached

`func (o *AAGUIDAuthenticatorCharacteristics) HasPlatformAttached() bool`

HasPlatformAttached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


