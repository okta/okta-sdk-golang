# UserSchemaPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**map[string]UserSchemaAttribute**](UserSchemaAttribute.md) |  | [optional] 
**Required** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewUserSchemaPublic

`func NewUserSchemaPublic() *UserSchemaPublic`

NewUserSchemaPublic instantiates a new UserSchemaPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSchemaPublicWithDefaults

`func NewUserSchemaPublicWithDefaults() *UserSchemaPublic`

NewUserSchemaPublicWithDefaults instantiates a new UserSchemaPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSchemaPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSchemaPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSchemaPublic) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserSchemaPublic) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProperties

`func (o *UserSchemaPublic) GetProperties() map[string]UserSchemaAttribute`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserSchemaPublic) GetPropertiesOk() (*map[string]UserSchemaAttribute, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserSchemaPublic) SetProperties(v map[string]UserSchemaAttribute)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UserSchemaPublic) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRequired

`func (o *UserSchemaPublic) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UserSchemaPublic) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UserSchemaPublic) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UserSchemaPublic) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetType

`func (o *UserSchemaPublic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSchemaPublic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSchemaPublic) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSchemaPublic) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


