# Parameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The update action to take | [optional] 
**Attribute** | Pointer to **string** | The attribute that tracks group memberships in AD. This should be &#x60;member&#x60; for AD. | [optional] 
**Values** | Pointer to **[]string** | List of user IDs whose group memberships to update | [optional] 

## Methods

### NewParameters

`func NewParameters() *Parameters`

NewParameters instantiates a new Parameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametersWithDefaults

`func NewParametersWithDefaults() *Parameters`

NewParametersWithDefaults instantiates a new Parameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Parameters) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Parameters) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Parameters) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Parameters) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAttribute

`func (o *Parameters) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Parameters) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Parameters) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *Parameters) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValues

`func (o *Parameters) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Parameters) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Parameters) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *Parameters) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


