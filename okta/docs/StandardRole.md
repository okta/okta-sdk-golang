# StandardRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentType** | Pointer to **string** | Role assignment type | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**Id** | Pointer to **string** | Role assignment ID | [optional] [readonly] 
**Label** | Pointer to **string** | Label for the role assignment | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the role assignment | [optional] 
**Type** | Pointer to **string** | Standard role type | [optional] 
**Embedded** | Pointer to [**StandardRoleEmbedded**](StandardRoleEmbedded.md) |  | [optional] 
**Links** | Pointer to [**LinksAssignee**](LinksAssignee.md) |  | [optional] 

## Methods

### NewStandardRole

`func NewStandardRole() *StandardRole`

NewStandardRole instantiates a new StandardRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardRoleWithDefaults

`func NewStandardRoleWithDefaults() *StandardRole`

NewStandardRoleWithDefaults instantiates a new StandardRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentType

`func (o *StandardRole) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *StandardRole) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *StandardRole) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *StandardRole) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetCreated

`func (o *StandardRole) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StandardRole) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StandardRole) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StandardRole) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *StandardRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandardRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandardRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StandardRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *StandardRole) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StandardRole) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StandardRole) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *StandardRole) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastUpdated

`func (o *StandardRole) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *StandardRole) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *StandardRole) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *StandardRole) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *StandardRole) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandardRole) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandardRole) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StandardRole) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *StandardRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StandardRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StandardRole) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StandardRole) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *StandardRole) GetEmbedded() StandardRoleEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *StandardRole) GetEmbeddedOk() (*StandardRoleEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *StandardRole) SetEmbedded(v StandardRoleEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *StandardRole) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *StandardRole) GetLinks() LinksAssignee`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StandardRole) GetLinksOk() (*LinksAssignee, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StandardRole) SetLinks(v LinksAssignee)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StandardRole) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


