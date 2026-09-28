# ProximityProviderAuthenticatorCharacteristics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HardwareProtected** | Pointer to **bool** | Indicates if this provider adds the hardware-protected authenticator characteristic. The hardware-protected characteristic means that authenticators require a physical device to authenticate. | [optional] [readonly] 

## Methods

### NewProximityProviderAuthenticatorCharacteristics

`func NewProximityProviderAuthenticatorCharacteristics() *ProximityProviderAuthenticatorCharacteristics`

NewProximityProviderAuthenticatorCharacteristics instantiates a new ProximityProviderAuthenticatorCharacteristics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProximityProviderAuthenticatorCharacteristicsWithDefaults

`func NewProximityProviderAuthenticatorCharacteristicsWithDefaults() *ProximityProviderAuthenticatorCharacteristics`

NewProximityProviderAuthenticatorCharacteristicsWithDefaults instantiates a new ProximityProviderAuthenticatorCharacteristics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHardwareProtected

`func (o *ProximityProviderAuthenticatorCharacteristics) GetHardwareProtected() bool`

GetHardwareProtected returns the HardwareProtected field if non-nil, zero value otherwise.

### GetHardwareProtectedOk

`func (o *ProximityProviderAuthenticatorCharacteristics) GetHardwareProtectedOk() (*bool, bool)`

GetHardwareProtectedOk returns a tuple with the HardwareProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareProtected

`func (o *ProximityProviderAuthenticatorCharacteristics) SetHardwareProtected(v bool)`

SetHardwareProtected sets HardwareProtected field to given value.

### HasHardwareProtected

`func (o *ProximityProviderAuthenticatorCharacteristics) HasHardwareProtected() bool`

HasHardwareProtected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


