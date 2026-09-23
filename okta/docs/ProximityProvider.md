# ProximityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorCharacteristics** | Pointer to [**ProximityProviderAuthenticatorCharacteristics**](ProximityProviderAuthenticatorCharacteristics.md) |  | [optional] 
**Id** | Pointer to **string** | Unique key for the proximity provider | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the proximity provider | [optional] [readonly] 

## Methods

### NewProximityProvider

`func NewProximityProvider() *ProximityProvider`

NewProximityProvider instantiates a new ProximityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProximityProviderWithDefaults

`func NewProximityProviderWithDefaults() *ProximityProvider`

NewProximityProviderWithDefaults instantiates a new ProximityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorCharacteristics

`func (o *ProximityProvider) GetAuthenticatorCharacteristics() ProximityProviderAuthenticatorCharacteristics`

GetAuthenticatorCharacteristics returns the AuthenticatorCharacteristics field if non-nil, zero value otherwise.

### GetAuthenticatorCharacteristicsOk

`func (o *ProximityProvider) GetAuthenticatorCharacteristicsOk() (*ProximityProviderAuthenticatorCharacteristics, bool)`

GetAuthenticatorCharacteristicsOk returns a tuple with the AuthenticatorCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorCharacteristics

`func (o *ProximityProvider) SetAuthenticatorCharacteristics(v ProximityProviderAuthenticatorCharacteristics)`

SetAuthenticatorCharacteristics sets AuthenticatorCharacteristics field to given value.

### HasAuthenticatorCharacteristics

`func (o *ProximityProvider) HasAuthenticatorCharacteristics() bool`

HasAuthenticatorCharacteristics returns a boolean if a field has been set.

### GetId

`func (o *ProximityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProximityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProximityProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProximityProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProximityProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProximityProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProximityProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProximityProvider) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


