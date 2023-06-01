# CustomHotpUserFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorProfileId** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to [**CustomHotpUserFactorProfile**](CustomHotpUserFactorProfile.md) |  | [optional] 

## Methods

### NewCustomHotpUserFactor

`func NewCustomHotpUserFactor() *CustomHotpUserFactor`

NewCustomHotpUserFactor instantiates a new CustomHotpUserFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomHotpUserFactorWithDefaults

`func NewCustomHotpUserFactorWithDefaults() *CustomHotpUserFactor`

NewCustomHotpUserFactorWithDefaults instantiates a new CustomHotpUserFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorProfileId

`func (o *CustomHotpUserFactor) GetFactorProfileId() string`

GetFactorProfileId returns the FactorProfileId field if non-nil, zero value otherwise.

### GetFactorProfileIdOk

`func (o *CustomHotpUserFactor) GetFactorProfileIdOk() (*string, bool)`

GetFactorProfileIdOk returns a tuple with the FactorProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorProfileId

`func (o *CustomHotpUserFactor) SetFactorProfileId(v string)`

SetFactorProfileId sets FactorProfileId field to given value.

### HasFactorProfileId

`func (o *CustomHotpUserFactor) HasFactorProfileId() bool`

HasFactorProfileId returns a boolean if a field has been set.

### GetProfile

`func (o *CustomHotpUserFactor) GetProfile() CustomHotpUserFactorProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CustomHotpUserFactor) GetProfileOk() (*CustomHotpUserFactorProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CustomHotpUserFactor) SetProfile(v CustomHotpUserFactorProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CustomHotpUserFactor) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


