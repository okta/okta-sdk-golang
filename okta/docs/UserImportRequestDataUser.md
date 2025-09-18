# UserImportRequestDataUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **map[string]string** | The &#x60;data.user.profile&#x60; contains the name-value pairs of the attributes in the user profile. If the user has been matched to an existing Okta user, a &#x60;data.user.id&#x60; object is included, containing the unique identifier of the Okta user profile.  You can change the values of the attributes by means of the &#x60;commands&#x60; object you return. | [optional] 

## Methods

### NewUserImportRequestDataUser

`func NewUserImportRequestDataUser() *UserImportRequestDataUser`

NewUserImportRequestDataUser instantiates a new UserImportRequestDataUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportRequestDataUserWithDefaults

`func NewUserImportRequestDataUserWithDefaults() *UserImportRequestDataUser`

NewUserImportRequestDataUserWithDefaults instantiates a new UserImportRequestDataUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *UserImportRequestDataUser) GetProfile() map[string]string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserImportRequestDataUser) GetProfileOk() (*map[string]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserImportRequestDataUser) SetProfile(v map[string]string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserImportRequestDataUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


