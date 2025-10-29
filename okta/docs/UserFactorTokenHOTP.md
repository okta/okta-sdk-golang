# UserFactorTokenHOTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorProfileId** | Pointer to **string** | ID of an existing Custom TOTP factor profile. To create this, see [Custom TOTP factor](https://help.okta.com/okta_help.htm?id&#x3D;ext-mfa-totp). | [optional] 
**FactorType** | Pointer to **interface{}** |  | [optional] 
**Profile** | Pointer to [**UserFactorTokenHOTPProfile**](UserFactorTokenHOTPProfile.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewUserFactorTokenHOTP

`func NewUserFactorTokenHOTP() *UserFactorTokenHOTP`

NewUserFactorTokenHOTP instantiates a new UserFactorTokenHOTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorTokenHOTPWithDefaults

`func NewUserFactorTokenHOTPWithDefaults() *UserFactorTokenHOTP`

NewUserFactorTokenHOTPWithDefaults instantiates a new UserFactorTokenHOTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorProfileId

`func (o *UserFactorTokenHOTP) GetFactorProfileId() string`

GetFactorProfileId returns the FactorProfileId field if non-nil, zero value otherwise.

### GetFactorProfileIdOk

`func (o *UserFactorTokenHOTP) GetFactorProfileIdOk() (*string, bool)`

GetFactorProfileIdOk returns a tuple with the FactorProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorProfileId

`func (o *UserFactorTokenHOTP) SetFactorProfileId(v string)`

SetFactorProfileId sets FactorProfileId field to given value.

### HasFactorProfileId

`func (o *UserFactorTokenHOTP) HasFactorProfileId() bool`

HasFactorProfileId returns a boolean if a field has been set.

### GetFactorType

`func (o *UserFactorTokenHOTP) GetFactorType() interface{}`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *UserFactorTokenHOTP) GetFactorTypeOk() (*interface{}, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *UserFactorTokenHOTP) SetFactorType(v interface{})`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *UserFactorTokenHOTP) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *UserFactorTokenHOTP) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *UserFactorTokenHOTP) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetProfile

`func (o *UserFactorTokenHOTP) GetProfile() UserFactorTokenHOTPProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorTokenHOTP) GetProfileOk() (*UserFactorTokenHOTPProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorTokenHOTP) SetProfile(v UserFactorTokenHOTPProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorTokenHOTP) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *UserFactorTokenHOTP) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserFactorTokenHOTP) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserFactorTokenHOTP) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserFactorTokenHOTP) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


