# ProfileMappingProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** |  | [optional] 
**PushStatus** | Pointer to [**ProfileMappingPropertyPushStatus**](ProfileMappingPropertyPushStatus.md) |  | [optional] 

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

`func (o *ProfileMappingProperty) GetPushStatus() ProfileMappingPropertyPushStatus`

GetPushStatus returns the PushStatus field if non-nil, zero value otherwise.

### GetPushStatusOk

`func (o *ProfileMappingProperty) GetPushStatusOk() (*ProfileMappingPropertyPushStatus, bool)`

GetPushStatusOk returns a tuple with the PushStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushStatus

`func (o *ProfileMappingProperty) SetPushStatus(v ProfileMappingPropertyPushStatus)`

SetPushStatus sets PushStatus field to given value.

### HasPushStatus

`func (o *ProfileMappingProperty) HasPushStatus() bool`

HasPushStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


