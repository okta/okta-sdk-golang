# SimulateResultRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]SimulateResultConditions**](SimulateResultConditions.md) | List of all conditions involved for this rule evaluation | [optional] 
**Id** | Pointer to **string** | The unique ID number of the policy rule | [optional] 
**Name** | Pointer to **string** | The name of the policy rule | [optional] 
**Status** | Pointer to **string** | The result of this entity evaluation | [optional] 

## Methods

### NewSimulateResultRules

`func NewSimulateResultRules() *SimulateResultRules`

NewSimulateResultRules instantiates a new SimulateResultRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateResultRulesWithDefaults

`func NewSimulateResultRulesWithDefaults() *SimulateResultRules`

NewSimulateResultRulesWithDefaults instantiates a new SimulateResultRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *SimulateResultRules) GetConditions() []SimulateResultConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SimulateResultRules) GetConditionsOk() (*[]SimulateResultConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SimulateResultRules) SetConditions(v []SimulateResultConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *SimulateResultRules) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetId

`func (o *SimulateResultRules) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimulateResultRules) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimulateResultRules) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimulateResultRules) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SimulateResultRules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimulateResultRules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimulateResultRules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimulateResultRules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *SimulateResultRules) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimulateResultRules) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimulateResultRules) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimulateResultRules) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


