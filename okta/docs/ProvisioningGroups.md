# ProvisioningGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ProvisioningGroupsAction**](ProvisioningGroupsAction.md) |  | [optional] 
**Assignments** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to **[]string** |  | [optional] 
**SourceAttributeName** | Pointer to **string** |  | [optional] 

## Methods

### NewProvisioningGroups

`func NewProvisioningGroups() *ProvisioningGroups`

NewProvisioningGroups instantiates a new ProvisioningGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningGroupsWithDefaults

`func NewProvisioningGroupsWithDefaults() *ProvisioningGroups`

NewProvisioningGroupsWithDefaults instantiates a new ProvisioningGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ProvisioningGroups) GetAction() ProvisioningGroupsAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProvisioningGroups) GetActionOk() (*ProvisioningGroupsAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProvisioningGroups) SetAction(v ProvisioningGroupsAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ProvisioningGroups) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAssignments

`func (o *ProvisioningGroups) GetAssignments() []string`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *ProvisioningGroups) GetAssignmentsOk() (*[]string, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *ProvisioningGroups) SetAssignments(v []string)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *ProvisioningGroups) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetFilter

`func (o *ProvisioningGroups) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ProvisioningGroups) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ProvisioningGroups) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ProvisioningGroups) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSourceAttributeName

`func (o *ProvisioningGroups) GetSourceAttributeName() string`

GetSourceAttributeName returns the SourceAttributeName field if non-nil, zero value otherwise.

### GetSourceAttributeNameOk

`func (o *ProvisioningGroups) GetSourceAttributeNameOk() (*string, bool)`

GetSourceAttributeNameOk returns a tuple with the SourceAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttributeName

`func (o *ProvisioningGroups) SetSourceAttributeName(v string)`

SetSourceAttributeName sets SourceAttributeName field to given value.

### HasSourceAttributeName

`func (o *ProvisioningGroups) HasSourceAttributeName() bool`

HasSourceAttributeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


