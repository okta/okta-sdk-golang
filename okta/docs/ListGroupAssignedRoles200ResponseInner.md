# ListGroupAssignedRoles200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentType** | Pointer to **string** | Role assignment type | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**Id** | Pointer to **string** | Binding object ID | [optional] [readonly] 
**Label** | Pointer to **string** | Label for the custom role assignment | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the custom role assignment | [optional] 
**Type** | Pointer to **string** | Standard role type | [optional] 
**Embedded** | Pointer to [**StandardRoleEmbedded**](StandardRoleEmbedded.md) |  | [optional] 
**Links** | Pointer to [**LinksCustomRoleResponse**](LinksCustomRoleResponse.md) |  | [optional] 
**ResourceSet** | Pointer to **string** | Resource set ID | [optional] [readonly] 
**Role** | Pointer to **string** | Custom role ID | [optional] [readonly] 

## Methods

### NewListGroupAssignedRoles200ResponseInner

`func NewListGroupAssignedRoles200ResponseInner() *ListGroupAssignedRoles200ResponseInner`

NewListGroupAssignedRoles200ResponseInner instantiates a new ListGroupAssignedRoles200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGroupAssignedRoles200ResponseInnerWithDefaults

`func NewListGroupAssignedRoles200ResponseInnerWithDefaults() *ListGroupAssignedRoles200ResponseInner`

NewListGroupAssignedRoles200ResponseInnerWithDefaults instantiates a new ListGroupAssignedRoles200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentType

`func (o *ListGroupAssignedRoles200ResponseInner) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *ListGroupAssignedRoles200ResponseInner) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *ListGroupAssignedRoles200ResponseInner) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetCreated

`func (o *ListGroupAssignedRoles200ResponseInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListGroupAssignedRoles200ResponseInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListGroupAssignedRoles200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ListGroupAssignedRoles200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListGroupAssignedRoles200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListGroupAssignedRoles200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ListGroupAssignedRoles200ResponseInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListGroupAssignedRoles200ResponseInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ListGroupAssignedRoles200ResponseInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListGroupAssignedRoles200ResponseInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListGroupAssignedRoles200ResponseInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListGroupAssignedRoles200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *ListGroupAssignedRoles200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListGroupAssignedRoles200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListGroupAssignedRoles200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ListGroupAssignedRoles200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListGroupAssignedRoles200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListGroupAssignedRoles200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *ListGroupAssignedRoles200ResponseInner) GetEmbedded() StandardRoleEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetEmbeddedOk() (*StandardRoleEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListGroupAssignedRoles200ResponseInner) SetEmbedded(v StandardRoleEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ListGroupAssignedRoles200ResponseInner) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *ListGroupAssignedRoles200ResponseInner) GetLinks() LinksCustomRoleResponse`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetLinksOk() (*LinksCustomRoleResponse, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListGroupAssignedRoles200ResponseInner) SetLinks(v LinksCustomRoleResponse)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListGroupAssignedRoles200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResourceSet

`func (o *ListGroupAssignedRoles200ResponseInner) GetResourceSet() string`

GetResourceSet returns the ResourceSet field if non-nil, zero value otherwise.

### GetResourceSetOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetResourceSetOk() (*string, bool)`

GetResourceSetOk returns a tuple with the ResourceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSet

`func (o *ListGroupAssignedRoles200ResponseInner) SetResourceSet(v string)`

SetResourceSet sets ResourceSet field to given value.

### HasResourceSet

`func (o *ListGroupAssignedRoles200ResponseInner) HasResourceSet() bool`

HasResourceSet returns a boolean if a field has been set.

### GetRole

`func (o *ListGroupAssignedRoles200ResponseInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ListGroupAssignedRoles200ResponseInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ListGroupAssignedRoles200ResponseInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ListGroupAssignedRoles200ResponseInner) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


