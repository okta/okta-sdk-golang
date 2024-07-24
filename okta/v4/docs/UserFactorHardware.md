# UserFactorHardware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorType** | Pointer to **interface{}** |  | [optional] 
**Profile** | Pointer to [**UserFactorHardwareProfile**](UserFactorHardwareProfile.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Verify** | Pointer to [**UserFactorHardwareAllOfVerify**](UserFactorHardwareAllOfVerify.md) |  | [optional] 

## Methods

### NewUserFactorHardware

`func NewUserFactorHardware() *UserFactorHardware`

NewUserFactorHardware instantiates a new UserFactorHardware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorHardwareWithDefaults

`func NewUserFactorHardwareWithDefaults() *UserFactorHardware`

NewUserFactorHardwareWithDefaults instantiates a new UserFactorHardware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorType

`func (o *UserFactorHardware) GetFactorType() interface{}`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *UserFactorHardware) GetFactorTypeOk() (*interface{}, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *UserFactorHardware) SetFactorType(v interface{})`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *UserFactorHardware) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *UserFactorHardware) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *UserFactorHardware) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetProfile

`func (o *UserFactorHardware) GetProfile() UserFactorHardwareProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorHardware) GetProfileOk() (*UserFactorHardwareProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorHardware) SetProfile(v UserFactorHardwareProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorHardware) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *UserFactorHardware) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserFactorHardware) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserFactorHardware) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserFactorHardware) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetVerify

`func (o *UserFactorHardware) GetVerify() UserFactorHardwareAllOfVerify`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *UserFactorHardware) GetVerifyOk() (*UserFactorHardwareAllOfVerify, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *UserFactorHardware) SetVerify(v UserFactorHardwareAllOfVerify)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *UserFactorHardware) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


