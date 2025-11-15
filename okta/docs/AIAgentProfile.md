# AIAgentProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the AI agent | [optional] 
**Name** | **string** | Unique name of the AI agent | 

## Methods

### NewAIAgentProfile

`func NewAIAgentProfile(name string, ) *AIAgentProfile`

NewAIAgentProfile instantiates a new AIAgentProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAgentProfileWithDefaults

`func NewAIAgentProfileWithDefaults() *AIAgentProfile`

NewAIAgentProfileWithDefaults instantiates a new AIAgentProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AIAgentProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIAgentProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIAgentProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIAgentProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AIAgentProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIAgentProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIAgentProfile) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


