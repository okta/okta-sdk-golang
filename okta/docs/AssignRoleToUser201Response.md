# AssignRoleToUser201Response

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

### NewAssignRoleToUser201Response

`func NewAssignRoleToUser201Response() *AssignRoleToUser201Response`

NewAssignRoleToUser201Response instantiates a new AssignRoleToUser201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignRoleToUser201ResponseWithDefaults

`func NewAssignRoleToUser201ResponseWithDefaults() *AssignRoleToUser201Response`

NewAssignRoleToUser201ResponseWithDefaults instantiates a new AssignRoleToUser201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentType

`func (o *AssignRoleToUser201Response) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *AssignRoleToUser201Response) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *AssignRoleToUser201Response) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *AssignRoleToUser201Response) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetCreated

`func (o *AssignRoleToUser201Response) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AssignRoleToUser201Response) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AssignRoleToUser201Response) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AssignRoleToUser201Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AssignRoleToUser201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssignRoleToUser201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssignRoleToUser201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssignRoleToUser201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *AssignRoleToUser201Response) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AssignRoleToUser201Response) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AssignRoleToUser201Response) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AssignRoleToUser201Response) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AssignRoleToUser201Response) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AssignRoleToUser201Response) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AssignRoleToUser201Response) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AssignRoleToUser201Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *AssignRoleToUser201Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssignRoleToUser201Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssignRoleToUser201Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssignRoleToUser201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *AssignRoleToUser201Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignRoleToUser201Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignRoleToUser201Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssignRoleToUser201Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *AssignRoleToUser201Response) GetEmbedded() StandardRoleEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AssignRoleToUser201Response) GetEmbeddedOk() (*StandardRoleEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AssignRoleToUser201Response) SetEmbedded(v StandardRoleEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AssignRoleToUser201Response) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *AssignRoleToUser201Response) GetLinks() LinksCustomRoleResponse`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AssignRoleToUser201Response) GetLinksOk() (*LinksCustomRoleResponse, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AssignRoleToUser201Response) SetLinks(v LinksCustomRoleResponse)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AssignRoleToUser201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResourceSet

`func (o *AssignRoleToUser201Response) GetResourceSet() string`

GetResourceSet returns the ResourceSet field if non-nil, zero value otherwise.

### GetResourceSetOk

`func (o *AssignRoleToUser201Response) GetResourceSetOk() (*string, bool)`

GetResourceSetOk returns a tuple with the ResourceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSet

`func (o *AssignRoleToUser201Response) SetResourceSet(v string)`

SetResourceSet sets ResourceSet field to given value.

### HasResourceSet

`func (o *AssignRoleToUser201Response) HasResourceSet() bool`

HasResourceSet returns a boolean if a field has been set.

### GetRole

`func (o *AssignRoleToUser201Response) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AssignRoleToUser201Response) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AssignRoleToUser201Response) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AssignRoleToUser201Response) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


