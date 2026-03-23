# LogIpServiceCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAnonymous** | Pointer to **NullableBool** | Indicates whether the service is an anonymizer | [optional] [readonly] 
**Operator** | Pointer to **NullableString** | The name of the associated operator | [optional] [readonly] 
**Type** | Pointer to **NullableString** | The type of service provided from this IP address | [optional] [readonly] 

## Methods

### NewLogIpServiceCategory

`func NewLogIpServiceCategory() *LogIpServiceCategory`

NewLogIpServiceCategory instantiates a new LogIpServiceCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogIpServiceCategoryWithDefaults

`func NewLogIpServiceCategoryWithDefaults() *LogIpServiceCategory`

NewLogIpServiceCategoryWithDefaults instantiates a new LogIpServiceCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAnonymous

`func (o *LogIpServiceCategory) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *LogIpServiceCategory) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *LogIpServiceCategory) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.

### HasIsAnonymous

`func (o *LogIpServiceCategory) HasIsAnonymous() bool`

HasIsAnonymous returns a boolean if a field has been set.

### SetIsAnonymousNil

`func (o *LogIpServiceCategory) SetIsAnonymousNil(b bool)`

 SetIsAnonymousNil sets the value for IsAnonymous to be an explicit nil

### UnsetIsAnonymous
`func (o *LogIpServiceCategory) UnsetIsAnonymous()`

UnsetIsAnonymous ensures that no value is present for IsAnonymous, not even an explicit nil
### GetOperator

`func (o *LogIpServiceCategory) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LogIpServiceCategory) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LogIpServiceCategory) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *LogIpServiceCategory) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *LogIpServiceCategory) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *LogIpServiceCategory) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetType

`func (o *LogIpServiceCategory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogIpServiceCategory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogIpServiceCategory) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LogIpServiceCategory) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *LogIpServiceCategory) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LogIpServiceCategory) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


