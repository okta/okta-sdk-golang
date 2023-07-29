# UserSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] [readonly] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**Definitions** | Pointer to [**UserSchemaDefinitions**](UserSchemaDefinitions.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Properties** | Pointer to [**UserSchemaProperties**](UserSchemaProperties.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewUserSchema

`func NewUserSchema() *UserSchema`

NewUserSchema instantiates a new UserSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSchemaWithDefaults

`func NewUserSchemaWithDefaults() *UserSchema`

NewUserSchemaWithDefaults instantiates a new UserSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *UserSchema) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *UserSchema) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *UserSchema) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *UserSchema) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetCreated

`func (o *UserSchema) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserSchema) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserSchema) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserSchema) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDefinitions

`func (o *UserSchema) GetDefinitions() UserSchemaDefinitions`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *UserSchema) GetDefinitionsOk() (*UserSchemaDefinitions, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *UserSchema) SetDefinitions(v UserSchemaDefinitions)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *UserSchema) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetId

`func (o *UserSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UserSchema) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UserSchema) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UserSchema) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UserSchema) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *UserSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *UserSchema) GetProperties() UserSchemaProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserSchema) GetPropertiesOk() (*UserSchemaProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserSchema) SetProperties(v UserSchemaProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UserSchema) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTitle

`func (o *UserSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *UserSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *UserSchema) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSchema) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSchema) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserSchema) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


