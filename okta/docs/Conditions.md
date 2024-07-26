# Conditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to [**Expression**](Expression.md) |  | [optional] 
**ProfileSourceId** | Pointer to **string** |  | [optional] 

## Methods

### NewConditions

`func NewConditions() *Conditions`

NewConditions instantiates a new Conditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionsWithDefaults

`func NewConditionsWithDefaults() *Conditions`

NewConditionsWithDefaults instantiates a new Conditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *Conditions) GetExpression() Expression`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Conditions) GetExpressionOk() (*Expression, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Conditions) SetExpression(v Expression)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *Conditions) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetProfileSourceId

`func (o *Conditions) GetProfileSourceId() string`

GetProfileSourceId returns the ProfileSourceId field if non-nil, zero value otherwise.

### GetProfileSourceIdOk

`func (o *Conditions) GetProfileSourceIdOk() (*string, bool)`

GetProfileSourceIdOk returns a tuple with the ProfileSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileSourceId

`func (o *Conditions) SetProfileSourceId(v string)`

SetProfileSourceId sets ProfileSourceId field to given value.

### HasProfileSourceId

`func (o *Conditions) HasProfileSourceId() bool`

HasProfileSourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


