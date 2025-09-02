# SimulateResultPoliciesItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]SimulateResultConditions**](SimulateResultConditions.md) | List of all conditions involved for this policy evaluation | [optional] 
**Id** | Pointer to **string** | ID of the specified policy type | [optional] 
**Name** | Pointer to **string** | Policy name | [optional] 
**Rules** | Pointer to [**[]SimulateResultRules**](SimulateResultRules.md) |  | [optional] 
**Status** | Pointer to **string** | The result of this entity evaluation | [optional] 

## Methods

### NewSimulateResultPoliciesItems

`func NewSimulateResultPoliciesItems() *SimulateResultPoliciesItems`

NewSimulateResultPoliciesItems instantiates a new SimulateResultPoliciesItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateResultPoliciesItemsWithDefaults

`func NewSimulateResultPoliciesItemsWithDefaults() *SimulateResultPoliciesItems`

NewSimulateResultPoliciesItemsWithDefaults instantiates a new SimulateResultPoliciesItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *SimulateResultPoliciesItems) GetConditions() []SimulateResultConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SimulateResultPoliciesItems) GetConditionsOk() (*[]SimulateResultConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SimulateResultPoliciesItems) SetConditions(v []SimulateResultConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *SimulateResultPoliciesItems) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetId

`func (o *SimulateResultPoliciesItems) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimulateResultPoliciesItems) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimulateResultPoliciesItems) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimulateResultPoliciesItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SimulateResultPoliciesItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimulateResultPoliciesItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimulateResultPoliciesItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimulateResultPoliciesItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *SimulateResultPoliciesItems) GetRules() []SimulateResultRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SimulateResultPoliciesItems) GetRulesOk() (*[]SimulateResultRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SimulateResultPoliciesItems) SetRules(v []SimulateResultRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SimulateResultPoliciesItems) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetStatus

`func (o *SimulateResultPoliciesItems) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimulateResultPoliciesItems) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimulateResultPoliciesItems) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimulateResultPoliciesItems) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


