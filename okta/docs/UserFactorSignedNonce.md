# UserFactorSignedNonce

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorType** | Pointer to **interface{}** |  | [optional] 
**Profile** | Pointer to [**UserFactorSignedNonceProfile**](UserFactorSignedNonceProfile.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewUserFactorSignedNonce

`func NewUserFactorSignedNonce() *UserFactorSignedNonce`

NewUserFactorSignedNonce instantiates a new UserFactorSignedNonce object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorSignedNonceWithDefaults

`func NewUserFactorSignedNonceWithDefaults() *UserFactorSignedNonce`

NewUserFactorSignedNonceWithDefaults instantiates a new UserFactorSignedNonce object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorType

`func (o *UserFactorSignedNonce) GetFactorType() interface{}`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *UserFactorSignedNonce) GetFactorTypeOk() (*interface{}, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *UserFactorSignedNonce) SetFactorType(v interface{})`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *UserFactorSignedNonce) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *UserFactorSignedNonce) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *UserFactorSignedNonce) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetProfile

`func (o *UserFactorSignedNonce) GetProfile() UserFactorSignedNonceProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorSignedNonce) GetProfileOk() (*UserFactorSignedNonceProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorSignedNonce) SetProfile(v UserFactorSignedNonceProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorSignedNonce) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *UserFactorSignedNonce) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserFactorSignedNonce) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserFactorSignedNonce) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserFactorSignedNonce) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


