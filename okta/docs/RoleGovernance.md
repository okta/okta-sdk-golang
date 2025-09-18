# RoleGovernance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grants** | Pointer to [**[]RoleGovernanceSource**](RoleGovernanceSource.md) |  | [optional] 
**Links** | Pointer to [**LinksGovernanceSources**](LinksGovernanceSources.md) |  | [optional] 

## Methods

### NewRoleGovernance

`func NewRoleGovernance() *RoleGovernance`

NewRoleGovernance instantiates a new RoleGovernance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleGovernanceWithDefaults

`func NewRoleGovernanceWithDefaults() *RoleGovernance`

NewRoleGovernanceWithDefaults instantiates a new RoleGovernance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrants

`func (o *RoleGovernance) GetGrants() []RoleGovernanceSource`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *RoleGovernance) GetGrantsOk() (*[]RoleGovernanceSource, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *RoleGovernance) SetGrants(v []RoleGovernanceSource)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *RoleGovernance) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetLinks

`func (o *RoleGovernance) GetLinks() LinksGovernanceSources`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoleGovernance) GetLinksOk() (*LinksGovernanceSources, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoleGovernance) SetLinks(v LinksGovernanceSources)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoleGovernance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


