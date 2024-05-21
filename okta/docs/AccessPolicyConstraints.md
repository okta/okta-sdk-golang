# AccessPolicyConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Knowledge** | Pointer to [**KnowledgeConstraint**](KnowledgeConstraint.md) |  | [optional] 
**Possession** | Pointer to [**PossessionConstraint**](PossessionConstraint.md) |  | [optional] 

## Methods

### NewAccessPolicyConstraints

`func NewAccessPolicyConstraints() *AccessPolicyConstraints`

NewAccessPolicyConstraints instantiates a new AccessPolicyConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyConstraintsWithDefaults

`func NewAccessPolicyConstraintsWithDefaults() *AccessPolicyConstraints`

NewAccessPolicyConstraintsWithDefaults instantiates a new AccessPolicyConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnowledge

`func (o *AccessPolicyConstraints) GetKnowledge() KnowledgeConstraint`

GetKnowledge returns the Knowledge field if non-nil, zero value otherwise.

### GetKnowledgeOk

`func (o *AccessPolicyConstraints) GetKnowledgeOk() (*KnowledgeConstraint, bool)`

GetKnowledgeOk returns a tuple with the Knowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledge

`func (o *AccessPolicyConstraints) SetKnowledge(v KnowledgeConstraint)`

SetKnowledge sets Knowledge field to given value.

### HasKnowledge

`func (o *AccessPolicyConstraints) HasKnowledge() bool`

HasKnowledge returns a boolean if a field has been set.

### GetPossession

`func (o *AccessPolicyConstraints) GetPossession() PossessionConstraint`

GetPossession returns the Possession field if non-nil, zero value otherwise.

### GetPossessionOk

`func (o *AccessPolicyConstraints) GetPossessionOk() (*PossessionConstraint, bool)`

GetPossessionOk returns a tuple with the Possession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossession

`func (o *AccessPolicyConstraints) SetPossession(v PossessionConstraint)`

SetPossession sets Possession field to given value.

### HasPossession

`func (o *AccessPolicyConstraints) HasPossession() bool`

HasPossession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


