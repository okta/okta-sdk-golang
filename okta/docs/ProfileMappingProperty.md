# ProfileMappingProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | Combination or single source properties that are mapped to the target property | [optional] 
**PushStatus** | Pointer to **string** | Indicates whether to update target properties for user create and update or just for user create.  Having a pushStatus of &#x60;PUSH&#x60; causes properties in the target to be updated on create and update. Having a pushStatus of &#x60;DONT_PUSH&#x60; causes properties in the target to be updated only on create. | [optional] 

## Methods

### NewProfileMappingProperty

`func NewProfileMappingProperty() *ProfileMappingProperty`

NewProfileMappingProperty instantiates a new ProfileMappingProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileMappingPropertyWithDefaults

`func NewProfileMappingPropertyWithDefaults() *ProfileMappingProperty`

NewProfileMappingPropertyWithDefaults instantiates a new ProfileMappingProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *ProfileMappingProperty) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ProfileMappingProperty) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ProfileMappingProperty) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ProfileMappingProperty) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetPushStatus

`func (o *ProfileMappingProperty) GetPushStatus() string`

GetPushStatus returns the PushStatus field if non-nil, zero value otherwise.

### GetPushStatusOk

`func (o *ProfileMappingProperty) GetPushStatusOk() (*string, bool)`

GetPushStatusOk returns a tuple with the PushStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushStatus

`func (o *ProfileMappingProperty) SetPushStatus(v string)`

SetPushStatus sets PushStatus field to given value.

### HasPushStatus

`func (o *ProfileMappingProperty) HasPushStatus() bool`

HasPushStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


