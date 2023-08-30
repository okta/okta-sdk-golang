# SelfServicePasswordResetActionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of rule action | [optional] [readonly] 
**Requirement** | Pointer to [**SsprRequirement**](SsprRequirement.md) |  | [optional] 

## Methods

### NewSelfServicePasswordResetActionAllOf

`func NewSelfServicePasswordResetActionAllOf() *SelfServicePasswordResetActionAllOf`

NewSelfServicePasswordResetActionAllOf instantiates a new SelfServicePasswordResetActionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServicePasswordResetActionAllOfWithDefaults

`func NewSelfServicePasswordResetActionAllOfWithDefaults() *SelfServicePasswordResetActionAllOf`

NewSelfServicePasswordResetActionAllOfWithDefaults instantiates a new SelfServicePasswordResetActionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SelfServicePasswordResetActionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelfServicePasswordResetActionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelfServicePasswordResetActionAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SelfServicePasswordResetActionAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequirement

`func (o *SelfServicePasswordResetActionAllOf) GetRequirement() SsprRequirement`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *SelfServicePasswordResetActionAllOf) GetRequirementOk() (*SsprRequirement, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *SelfServicePasswordResetActionAllOf) SetRequirement(v SsprRequirement)`

SetRequirement sets Requirement field to given value.

### HasRequirement

`func (o *SelfServicePasswordResetActionAllOf) HasRequirement() bool`

HasRequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


