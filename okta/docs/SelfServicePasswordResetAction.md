# SelfServicePasswordResetAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type of rule action | [optional] [readonly] 
**Requirement** | Pointer to [**SsprRequirement**](SsprRequirement.md) |  | [optional] 

## Methods

### NewSelfServicePasswordResetAction

`func NewSelfServicePasswordResetAction() *SelfServicePasswordResetAction`

NewSelfServicePasswordResetAction instantiates a new SelfServicePasswordResetAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServicePasswordResetActionWithDefaults

`func NewSelfServicePasswordResetActionWithDefaults() *SelfServicePasswordResetAction`

NewSelfServicePasswordResetActionWithDefaults instantiates a new SelfServicePasswordResetAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *SelfServicePasswordResetAction) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SelfServicePasswordResetAction) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SelfServicePasswordResetAction) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SelfServicePasswordResetAction) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetType

`func (o *SelfServicePasswordResetAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelfServicePasswordResetAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelfServicePasswordResetAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SelfServicePasswordResetAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequirement

`func (o *SelfServicePasswordResetAction) GetRequirement() SsprRequirement`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *SelfServicePasswordResetAction) GetRequirementOk() (*SsprRequirement, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *SelfServicePasswordResetAction) SetRequirement(v SsprRequirement)`

SetRequirement sets Requirement field to given value.

### HasRequirement

`func (o *SelfServicePasswordResetAction) HasRequirement() bool`

HasRequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


