# ProvisioningGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Provisioning action for the IdP user&#39;s group memberships  | Enum     | Description                                                                                                                                                       | Existing OKTA_GROUP Memberships                                                                       | Existing APP_GROUP Memberships | Existing BUILT_IN Memberships | | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ------------------------------ | ----------------------------- | | &#x60;APPEND&#x60; | Adds a user to any group defined by the IdP as a value of the &#x60;sourceAttributeName&#x60; array that matches the name of the allow listed group defined in the &#x60;filter&#x60; | Unchanged                                                                                             | Unchanged                      | Unchanged                     | | &#x60;ASSIGN&#x60; | Assigns a user to groups defined in the &#x60;assignments&#x60; array                                                                                                       | Unchanged                                                                                             | Unchanged                      | Unchanged                     | | &#x60;NONE&#x60;   | Skips processing of group memberships                                                                                                                             | Unchanged                                                                                             | Unchanged                      | Unchanged                     | | &#x60;SYNC&#x60;   | Group memberships are sourced by the IdP as a value of the &#x60;sourceAttributeName&#x60; array that matches the name of the group defined in the &#x60;filter&#x60;                 | Removed if not defined by the IdP in &#x60;sourceAttributeName&#x60; and matching name of the group in &#x60;filter&#x60; | Unchanged                      | Unchanged                     |  &gt; **Note:** Group provisioning action is processed independently from profile sourcing. You can sync group memberships through SAML with profile sourcing disabled. | [optional] 
**Assignments** | Pointer to **[]string** | List of &#x60;OKTA_GROUP&#x60; group identifiers to add an IdP user as a member with the &#x60;ASSIGN&#x60; action | [optional] 
**Filter** | Pointer to **[]string** | Allowlist of &#x60;OKTA_GROUP&#x60; group identifiers for the &#x60;APPEND&#x60; or &#x60;SYNC&#x60; provisioning action | [optional] 
**SourceAttributeName** | Pointer to **string** | IdP user profile attribute name (case-insensitive) for an array value that contains group memberships | [optional] 

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

`func (o *ProvisioningGroups) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProvisioningGroups) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProvisioningGroups) SetAction(v string)`

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


