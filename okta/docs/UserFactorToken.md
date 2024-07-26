# UserFactorToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorType** | Pointer to **interface{}** |  | [optional] 
**Profile** | Pointer to [**UserFactorTokenProfile**](UserFactorTokenProfile.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Verify** | Pointer to [**UserFactorTokenAllOfVerify**](UserFactorTokenAllOfVerify.md) |  | [optional] 

## Methods

### NewUserFactorToken

`func NewUserFactorToken() *UserFactorToken`

NewUserFactorToken instantiates a new UserFactorToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorTokenWithDefaults

`func NewUserFactorTokenWithDefaults() *UserFactorToken`

NewUserFactorTokenWithDefaults instantiates a new UserFactorToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorType

`func (o *UserFactorToken) GetFactorType() interface{}`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *UserFactorToken) GetFactorTypeOk() (*interface{}, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *UserFactorToken) SetFactorType(v interface{})`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *UserFactorToken) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *UserFactorToken) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *UserFactorToken) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetProfile

`func (o *UserFactorToken) GetProfile() UserFactorTokenProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorToken) GetProfileOk() (*UserFactorTokenProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorToken) SetProfile(v UserFactorTokenProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorToken) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *UserFactorToken) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserFactorToken) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserFactorToken) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserFactorToken) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetVerify

`func (o *UserFactorToken) GetVerify() UserFactorTokenAllOfVerify`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *UserFactorToken) GetVerifyOk() (*UserFactorTokenAllOfVerify, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *UserFactorToken) SetVerify(v UserFactorTokenAllOfVerify)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *UserFactorToken) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


