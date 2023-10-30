# SsprStepUpRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Methods** | Pointer to **[]string** | Authenticator methods required for secondary authentication step of password recovery. Specify this value only when &#x60;required&#x60; is true and &#x60;security_question&#x60; is permitted for the secondary authentication. | [optional] 
**Required** | Pointer to **bool** |  | [optional] 

## Methods

### NewSsprStepUpRequirement

`func NewSsprStepUpRequirement() *SsprStepUpRequirement`

NewSsprStepUpRequirement instantiates a new SsprStepUpRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsprStepUpRequirementWithDefaults

`func NewSsprStepUpRequirementWithDefaults() *SsprStepUpRequirement`

NewSsprStepUpRequirementWithDefaults instantiates a new SsprStepUpRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethods

`func (o *SsprStepUpRequirement) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *SsprStepUpRequirement) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *SsprStepUpRequirement) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *SsprStepUpRequirement) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetRequired

`func (o *SsprStepUpRequirement) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SsprStepUpRequirement) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SsprStepUpRequirement) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SsprStepUpRequirement) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


