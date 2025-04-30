# GroupSchemaCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Properties** | Pointer to [**map[string]GroupSchemaAttribute**](GroupSchemaAttribute.md) |  | [optional] 
**Required** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupSchemaCustom

`func NewGroupSchemaCustom() *GroupSchemaCustom`

NewGroupSchemaCustom instantiates a new GroupSchemaCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSchemaCustomWithDefaults

`func NewGroupSchemaCustomWithDefaults() *GroupSchemaCustom`

NewGroupSchemaCustomWithDefaults instantiates a new GroupSchemaCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupSchemaCustom) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupSchemaCustom) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupSchemaCustom) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupSchemaCustom) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProperties

`func (o *GroupSchemaCustom) GetProperties() map[string]GroupSchemaAttribute`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GroupSchemaCustom) GetPropertiesOk() (*map[string]GroupSchemaAttribute, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GroupSchemaCustom) SetProperties(v map[string]GroupSchemaAttribute)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GroupSchemaCustom) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRequired

`func (o *GroupSchemaCustom) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GroupSchemaCustom) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GroupSchemaCustom) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GroupSchemaCustom) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetType

`func (o *GroupSchemaCustom) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupSchemaCustom) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupSchemaCustom) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupSchemaCustom) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


