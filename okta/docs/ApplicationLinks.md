# ApplicationLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPolicy** | Pointer to [**AccessPolicyLink**](AccessPolicyLink.md) |  | [optional] 
**Activate** | Pointer to [**HrefObjectActivateLink**](HrefObjectActivateLink.md) |  | [optional] 
**AppLinks** | Pointer to [**[]HrefObject**](HrefObject.md) | List of app link resources | [optional] 
**Deactivate** | Pointer to [**HrefObjectDeactivateLink**](HrefObjectDeactivateLink.md) |  | [optional] 
**Groups** | Pointer to [**GroupsLink**](GroupsLink.md) |  | [optional] 
**Help** | Pointer to [**HelpLink**](HelpLink.md) |  | [optional] 
**Logo** | Pointer to [**[]HrefObject**](HrefObject.md) | List of app logo resources | [optional] 
**Metadata** | Pointer to [**MetadataLink**](MetadataLink.md) |  | [optional] 
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**Users** | Pointer to [**UsersLink**](UsersLink.md) |  | [optional] 

## Methods

### NewApplicationLinks

`func NewApplicationLinks() *ApplicationLinks`

NewApplicationLinks instantiates a new ApplicationLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLinksWithDefaults

`func NewApplicationLinksWithDefaults() *ApplicationLinks`

NewApplicationLinksWithDefaults instantiates a new ApplicationLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPolicy

`func (o *ApplicationLinks) GetAccessPolicy() AccessPolicyLink`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *ApplicationLinks) GetAccessPolicyOk() (*AccessPolicyLink, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *ApplicationLinks) SetAccessPolicy(v AccessPolicyLink)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *ApplicationLinks) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.

### GetActivate

`func (o *ApplicationLinks) GetActivate() HrefObjectActivateLink`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *ApplicationLinks) GetActivateOk() (*HrefObjectActivateLink, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *ApplicationLinks) SetActivate(v HrefObjectActivateLink)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *ApplicationLinks) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetAppLinks

`func (o *ApplicationLinks) GetAppLinks() []HrefObject`

GetAppLinks returns the AppLinks field if non-nil, zero value otherwise.

### GetAppLinksOk

`func (o *ApplicationLinks) GetAppLinksOk() (*[]HrefObject, bool)`

GetAppLinksOk returns a tuple with the AppLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLinks

`func (o *ApplicationLinks) SetAppLinks(v []HrefObject)`

SetAppLinks sets AppLinks field to given value.

### HasAppLinks

`func (o *ApplicationLinks) HasAppLinks() bool`

HasAppLinks returns a boolean if a field has been set.

### GetDeactivate

`func (o *ApplicationLinks) GetDeactivate() HrefObjectDeactivateLink`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *ApplicationLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *ApplicationLinks) SetDeactivate(v HrefObjectDeactivateLink)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *ApplicationLinks) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetGroups

`func (o *ApplicationLinks) GetGroups() GroupsLink`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ApplicationLinks) GetGroupsOk() (*GroupsLink, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ApplicationLinks) SetGroups(v GroupsLink)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ApplicationLinks) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHelp

`func (o *ApplicationLinks) GetHelp() HelpLink`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *ApplicationLinks) GetHelpOk() (*HelpLink, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *ApplicationLinks) SetHelp(v HelpLink)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *ApplicationLinks) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetLogo

`func (o *ApplicationLinks) GetLogo() []HrefObject`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ApplicationLinks) GetLogoOk() (*[]HrefObject, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ApplicationLinks) SetLogo(v []HrefObject)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ApplicationLinks) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationLinks) GetMetadata() MetadataLink`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationLinks) GetMetadataOk() (*MetadataLink, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationLinks) SetMetadata(v MetadataLink)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationLinks) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSelf

`func (o *ApplicationLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ApplicationLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ApplicationLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ApplicationLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetUsers

`func (o *ApplicationLinks) GetUsers() UsersLink`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ApplicationLinks) GetUsersOk() (*UsersLink, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ApplicationLinks) SetUsers(v UsersLink)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ApplicationLinks) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


