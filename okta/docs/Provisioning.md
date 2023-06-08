# Provisioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ProvisioningAction**](ProvisioningAction.md) |  | [optional] 
**Conditions** | Pointer to [**ProvisioningConditions**](ProvisioningConditions.md) |  | [optional] 
**Groups** | Pointer to [**ProvisioningGroups**](ProvisioningGroups.md) |  | [optional] 
**ProfileMaster** | Pointer to **bool** |  | [optional] 

## Methods

### NewProvisioning

`func NewProvisioning() *Provisioning`

NewProvisioning instantiates a new Provisioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningWithDefaults

`func NewProvisioningWithDefaults() *Provisioning`

NewProvisioningWithDefaults instantiates a new Provisioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Provisioning) GetAction() ProvisioningAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Provisioning) GetActionOk() (*ProvisioningAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Provisioning) SetAction(v ProvisioningAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *Provisioning) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetConditions

`func (o *Provisioning) GetConditions() ProvisioningConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Provisioning) GetConditionsOk() (*ProvisioningConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Provisioning) SetConditions(v ProvisioningConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Provisioning) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetGroups

`func (o *Provisioning) GetGroups() ProvisioningGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Provisioning) GetGroupsOk() (*ProvisioningGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Provisioning) SetGroups(v ProvisioningGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Provisioning) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetProfileMaster

`func (o *Provisioning) GetProfileMaster() bool`

GetProfileMaster returns the ProfileMaster field if non-nil, zero value otherwise.

### GetProfileMasterOk

`func (o *Provisioning) GetProfileMasterOk() (*bool, bool)`

GetProfileMasterOk returns a tuple with the ProfileMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMaster

`func (o *Provisioning) SetProfileMaster(v bool)`

SetProfileMaster sets ProfileMaster field to given value.

### HasProfileMaster

`func (o *Provisioning) HasProfileMaster() bool`

HasProfileMaster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


