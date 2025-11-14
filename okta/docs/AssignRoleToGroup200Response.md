# AssignRoleToGroup200Response

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

### NewAssignRoleToGroup200Response

`func NewAssignRoleToGroup200Response() *AssignRoleToGroup200Response`

NewAssignRoleToGroup200Response instantiates a new AssignRoleToGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignRoleToGroup200ResponseWithDefaults

`func NewAssignRoleToGroup200ResponseWithDefaults() *AssignRoleToGroup200Response`

NewAssignRoleToGroup200ResponseWithDefaults instantiates a new AssignRoleToGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentType

`func (o *AssignRoleToGroup200Response) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *AssignRoleToGroup200Response) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *AssignRoleToGroup200Response) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *AssignRoleToGroup200Response) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetCreated

`func (o *AssignRoleToGroup200Response) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AssignRoleToGroup200Response) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AssignRoleToGroup200Response) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AssignRoleToGroup200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AssignRoleToGroup200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssignRoleToGroup200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssignRoleToGroup200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssignRoleToGroup200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *AssignRoleToGroup200Response) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AssignRoleToGroup200Response) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AssignRoleToGroup200Response) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AssignRoleToGroup200Response) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AssignRoleToGroup200Response) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AssignRoleToGroup200Response) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AssignRoleToGroup200Response) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AssignRoleToGroup200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *AssignRoleToGroup200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssignRoleToGroup200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssignRoleToGroup200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssignRoleToGroup200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *AssignRoleToGroup200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignRoleToGroup200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignRoleToGroup200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssignRoleToGroup200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *AssignRoleToGroup200Response) GetEmbedded() StandardRoleEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AssignRoleToGroup200Response) GetEmbeddedOk() (*StandardRoleEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AssignRoleToGroup200Response) SetEmbedded(v StandardRoleEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AssignRoleToGroup200Response) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *AssignRoleToGroup200Response) GetLinks() LinksCustomRoleResponse`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AssignRoleToGroup200Response) GetLinksOk() (*LinksCustomRoleResponse, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AssignRoleToGroup200Response) SetLinks(v LinksCustomRoleResponse)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AssignRoleToGroup200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResourceSet

`func (o *AssignRoleToGroup200Response) GetResourceSet() string`

GetResourceSet returns the ResourceSet field if non-nil, zero value otherwise.

### GetResourceSetOk

`func (o *AssignRoleToGroup200Response) GetResourceSetOk() (*string, bool)`

GetResourceSetOk returns a tuple with the ResourceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSet

`func (o *AssignRoleToGroup200Response) SetResourceSet(v string)`

SetResourceSet sets ResourceSet field to given value.

### HasResourceSet

`func (o *AssignRoleToGroup200Response) HasResourceSet() bool`

HasResourceSet returns a boolean if a field has been set.

### GetRole

`func (o *AssignRoleToGroup200Response) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AssignRoleToGroup200Response) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AssignRoleToGroup200Response) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AssignRoleToGroup200Response) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


