# LogUserBehavior

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The unique identifier of the user behavior detection model | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The name of the user behavior detection model [configured by admins](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Behavior/) | [optional] [readonly] 
**Result** | Pointer to **NullableString** | The result of the user behavior analysis | [optional] [readonly] 

## Methods

### NewLogUserBehavior

`func NewLogUserBehavior() *LogUserBehavior`

NewLogUserBehavior instantiates a new LogUserBehavior object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogUserBehaviorWithDefaults

`func NewLogUserBehaviorWithDefaults() *LogUserBehavior`

NewLogUserBehaviorWithDefaults instantiates a new LogUserBehavior object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogUserBehavior) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogUserBehavior) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogUserBehavior) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogUserBehavior) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LogUserBehavior) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LogUserBehavior) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *LogUserBehavior) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogUserBehavior) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogUserBehavior) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogUserBehavior) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LogUserBehavior) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LogUserBehavior) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResult

`func (o *LogUserBehavior) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *LogUserBehavior) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *LogUserBehavior) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *LogUserBehavior) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *LogUserBehavior) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *LogUserBehavior) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


