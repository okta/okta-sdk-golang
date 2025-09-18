# CustomRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentType** | Pointer to **string** | Role assignment type | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**Id** | Pointer to **string** | Binding object ID | [optional] [readonly] 
**Label** | Pointer to **string** | Label for the custom role assignment | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**ResourceSet** | Pointer to **string** | Resource set ID | [optional] [readonly] 
**Role** | Pointer to **string** | Custom role ID | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the custom role assignment | [optional] 
**Type** | Pointer to **string** | CUSTOM for a custom role | [optional] 
**Links** | Pointer to [**LinksCustomRoleResponse**](LinksCustomRoleResponse.md) |  | [optional] 

## Methods

### NewCustomRole

`func NewCustomRole() *CustomRole`

NewCustomRole instantiates a new CustomRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRoleWithDefaults

`func NewCustomRoleWithDefaults() *CustomRole`

NewCustomRoleWithDefaults instantiates a new CustomRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentType

`func (o *CustomRole) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *CustomRole) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *CustomRole) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *CustomRole) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetCreated

`func (o *CustomRole) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomRole) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomRole) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CustomRole) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *CustomRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *CustomRole) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomRole) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomRole) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomRole) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CustomRole) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CustomRole) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CustomRole) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CustomRole) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetResourceSet

`func (o *CustomRole) GetResourceSet() string`

GetResourceSet returns the ResourceSet field if non-nil, zero value otherwise.

### GetResourceSetOk

`func (o *CustomRole) GetResourceSetOk() (*string, bool)`

GetResourceSetOk returns a tuple with the ResourceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSet

`func (o *CustomRole) SetResourceSet(v string)`

SetResourceSet sets ResourceSet field to given value.

### HasResourceSet

`func (o *CustomRole) HasResourceSet() bool`

HasResourceSet returns a boolean if a field has been set.

### GetRole

`func (o *CustomRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CustomRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CustomRole) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CustomRole) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *CustomRole) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomRole) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomRole) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomRole) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *CustomRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomRole) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomRole) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *CustomRole) GetLinks() LinksCustomRoleResponse`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomRole) GetLinksOk() (*LinksCustomRoleResponse, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomRole) SetLinks(v LinksCustomRoleResponse)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomRole) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


