# AppConfigActiveDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinguishedName** | **string** | The distinguished name of the group in Active Directory | 
**GroupScope** | **string** | The scope of the group in Active Directory | 
**GroupType** | **string** | The type of the group in Active Directory | 
**SamAccountName** | **string** | The SAM account name of the group in Active Directory | 

## Methods

### NewAppConfigActiveDirectory

`func NewAppConfigActiveDirectory(distinguishedName string, groupScope string, groupType string, samAccountName string, ) *AppConfigActiveDirectory`

NewAppConfigActiveDirectory instantiates a new AppConfigActiveDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConfigActiveDirectoryWithDefaults

`func NewAppConfigActiveDirectoryWithDefaults() *AppConfigActiveDirectory`

NewAppConfigActiveDirectoryWithDefaults instantiates a new AppConfigActiveDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinguishedName

`func (o *AppConfigActiveDirectory) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *AppConfigActiveDirectory) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *AppConfigActiveDirectory) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.


### GetGroupScope

`func (o *AppConfigActiveDirectory) GetGroupScope() string`

GetGroupScope returns the GroupScope field if non-nil, zero value otherwise.

### GetGroupScopeOk

`func (o *AppConfigActiveDirectory) GetGroupScopeOk() (*string, bool)`

GetGroupScopeOk returns a tuple with the GroupScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupScope

`func (o *AppConfigActiveDirectory) SetGroupScope(v string)`

SetGroupScope sets GroupScope field to given value.


### GetGroupType

`func (o *AppConfigActiveDirectory) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *AppConfigActiveDirectory) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *AppConfigActiveDirectory) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.


### GetSamAccountName

`func (o *AppConfigActiveDirectory) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *AppConfigActiveDirectory) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *AppConfigActiveDirectory) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


