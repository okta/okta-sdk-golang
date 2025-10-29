# OrgCreationAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**OrgCreationAdminCredentials**](OrgCreationAdminCredentials.md) |  | [optional] 
**Profile** | [**OrgCreationAdminProfile**](OrgCreationAdminProfile.md) |  | 

## Methods

### NewOrgCreationAdmin

`func NewOrgCreationAdmin(profile OrgCreationAdminProfile, ) *OrgCreationAdmin`

NewOrgCreationAdmin instantiates a new OrgCreationAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgCreationAdminWithDefaults

`func NewOrgCreationAdminWithDefaults() *OrgCreationAdmin`

NewOrgCreationAdminWithDefaults instantiates a new OrgCreationAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *OrgCreationAdmin) GetCredentials() OrgCreationAdminCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OrgCreationAdmin) GetCredentialsOk() (*OrgCreationAdminCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OrgCreationAdmin) SetCredentials(v OrgCreationAdminCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OrgCreationAdmin) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetProfile

`func (o *OrgCreationAdmin) GetProfile() OrgCreationAdminProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OrgCreationAdmin) GetProfileOk() (*OrgCreationAdminProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OrgCreationAdmin) SetProfile(v OrgCreationAdminProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


