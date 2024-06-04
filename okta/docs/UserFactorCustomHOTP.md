# UserFactorCustomHOTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorProfileId** | Pointer to **string** | ID of an existing Custom TOTP Factor profile. To create this, see [Custom TOTP Factor](https://help.okta.com/okta_help.htm?id&#x3D;ext-mfa-totp). | [optional] 
**Profile** | Pointer to [**UserFactorCustomHOTPProfile**](UserFactorCustomHOTPProfile.md) |  | [optional] 

## Methods

### NewUserFactorCustomHOTP

`func NewUserFactorCustomHOTP() *UserFactorCustomHOTP`

NewUserFactorCustomHOTP instantiates a new UserFactorCustomHOTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorCustomHOTPWithDefaults

`func NewUserFactorCustomHOTPWithDefaults() *UserFactorCustomHOTP`

NewUserFactorCustomHOTPWithDefaults instantiates a new UserFactorCustomHOTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorProfileId

`func (o *UserFactorCustomHOTP) GetFactorProfileId() string`

GetFactorProfileId returns the FactorProfileId field if non-nil, zero value otherwise.

### GetFactorProfileIdOk

`func (o *UserFactorCustomHOTP) GetFactorProfileIdOk() (*string, bool)`

GetFactorProfileIdOk returns a tuple with the FactorProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorProfileId

`func (o *UserFactorCustomHOTP) SetFactorProfileId(v string)`

SetFactorProfileId sets FactorProfileId field to given value.

### HasFactorProfileId

`func (o *UserFactorCustomHOTP) HasFactorProfileId() bool`

HasFactorProfileId returns a boolean if a field has been set.

### GetProfile

`func (o *UserFactorCustomHOTP) GetProfile() UserFactorCustomHOTPProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorCustomHOTP) GetProfileOk() (*UserFactorCustomHOTPProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorCustomHOTP) SetProfile(v UserFactorCustomHOTPProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorCustomHOTP) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


