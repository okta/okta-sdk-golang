# GroupSchemaBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Properties** | Pointer to [**GroupSchemaBaseProperties**](GroupSchemaBaseProperties.md) |  | [optional] 
**Required** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupSchemaBase

`func NewGroupSchemaBase() *GroupSchemaBase`

NewGroupSchemaBase instantiates a new GroupSchemaBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSchemaBaseWithDefaults

`func NewGroupSchemaBaseWithDefaults() *GroupSchemaBase`

NewGroupSchemaBaseWithDefaults instantiates a new GroupSchemaBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupSchemaBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupSchemaBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupSchemaBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupSchemaBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProperties

`func (o *GroupSchemaBase) GetProperties() GroupSchemaBaseProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GroupSchemaBase) GetPropertiesOk() (*GroupSchemaBaseProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GroupSchemaBase) SetProperties(v GroupSchemaBaseProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GroupSchemaBase) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRequired

`func (o *GroupSchemaBase) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GroupSchemaBase) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GroupSchemaBase) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GroupSchemaBase) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetType

`func (o *GroupSchemaBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupSchemaBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupSchemaBase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupSchemaBase) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


