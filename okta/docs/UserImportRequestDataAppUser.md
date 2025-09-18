# UserImportRequestDataAppUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **map[string]string** | Provides the name-value pairs of the attributes contained in the app user profile of the user who is being imported. You can change the values of attributes in the user&#39;s app profile by means of the &#x60;commands&#x60; object you return. If you change attributes in the app profile, they then flow through to the Okta user profile, based on matching and mapping rules. | [optional] 

## Methods

### NewUserImportRequestDataAppUser

`func NewUserImportRequestDataAppUser() *UserImportRequestDataAppUser`

NewUserImportRequestDataAppUser instantiates a new UserImportRequestDataAppUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportRequestDataAppUserWithDefaults

`func NewUserImportRequestDataAppUserWithDefaults() *UserImportRequestDataAppUser`

NewUserImportRequestDataAppUserWithDefaults instantiates a new UserImportRequestDataAppUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *UserImportRequestDataAppUser) GetProfile() map[string]string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserImportRequestDataAppUser) GetProfileOk() (*map[string]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserImportRequestDataAppUser) SetProfile(v map[string]string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserImportRequestDataAppUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


