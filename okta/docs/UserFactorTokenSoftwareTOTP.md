# UserFactorTokenSoftwareTOTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorType** | Pointer to **interface{}** |  | [optional] 
**Profile** | Pointer to [**UserFactorTokenProfile**](UserFactorTokenProfile.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewUserFactorTokenSoftwareTOTP

`func NewUserFactorTokenSoftwareTOTP() *UserFactorTokenSoftwareTOTP`

NewUserFactorTokenSoftwareTOTP instantiates a new UserFactorTokenSoftwareTOTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorTokenSoftwareTOTPWithDefaults

`func NewUserFactorTokenSoftwareTOTPWithDefaults() *UserFactorTokenSoftwareTOTP`

NewUserFactorTokenSoftwareTOTPWithDefaults instantiates a new UserFactorTokenSoftwareTOTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorType

`func (o *UserFactorTokenSoftwareTOTP) GetFactorType() interface{}`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *UserFactorTokenSoftwareTOTP) GetFactorTypeOk() (*interface{}, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *UserFactorTokenSoftwareTOTP) SetFactorType(v interface{})`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *UserFactorTokenSoftwareTOTP) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *UserFactorTokenSoftwareTOTP) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *UserFactorTokenSoftwareTOTP) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetProfile

`func (o *UserFactorTokenSoftwareTOTP) GetProfile() UserFactorTokenProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorTokenSoftwareTOTP) GetProfileOk() (*UserFactorTokenProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorTokenSoftwareTOTP) SetProfile(v UserFactorTokenProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorTokenSoftwareTOTP) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *UserFactorTokenSoftwareTOTP) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserFactorTokenSoftwareTOTP) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserFactorTokenSoftwareTOTP) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserFactorTokenSoftwareTOTP) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


