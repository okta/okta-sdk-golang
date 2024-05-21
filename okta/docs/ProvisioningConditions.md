# ProvisioningConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deprovisioned** | Pointer to [**ProvisioningDeprovisionedCondition**](ProvisioningDeprovisionedCondition.md) |  | [optional] 
**Suspended** | Pointer to [**ProvisioningSuspendedCondition**](ProvisioningSuspendedCondition.md) |  | [optional] 

## Methods

### NewProvisioningConditions

`func NewProvisioningConditions() *ProvisioningConditions`

NewProvisioningConditions instantiates a new ProvisioningConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConditionsWithDefaults

`func NewProvisioningConditionsWithDefaults() *ProvisioningConditions`

NewProvisioningConditionsWithDefaults instantiates a new ProvisioningConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeprovisioned

`func (o *ProvisioningConditions) GetDeprovisioned() ProvisioningDeprovisionedCondition`

GetDeprovisioned returns the Deprovisioned field if non-nil, zero value otherwise.

### GetDeprovisionedOk

`func (o *ProvisioningConditions) GetDeprovisionedOk() (*ProvisioningDeprovisionedCondition, bool)`

GetDeprovisionedOk returns a tuple with the Deprovisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprovisioned

`func (o *ProvisioningConditions) SetDeprovisioned(v ProvisioningDeprovisionedCondition)`

SetDeprovisioned sets Deprovisioned field to given value.

### HasDeprovisioned

`func (o *ProvisioningConditions) HasDeprovisioned() bool`

HasDeprovisioned returns a boolean if a field has been set.

### GetSuspended

`func (o *ProvisioningConditions) GetSuspended() ProvisioningSuspendedCondition`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ProvisioningConditions) GetSuspendedOk() (*ProvisioningSuspendedCondition, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ProvisioningConditions) SetSuspended(v ProvisioningSuspendedCondition)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ProvisioningConditions) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


