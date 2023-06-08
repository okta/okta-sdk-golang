# GroupSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] [readonly] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**Definitions** | Pointer to [**GroupSchemaDefinitions**](GroupSchemaDefinitions.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Properties** | Pointer to [**UserSchemaProperties**](UserSchemaProperties.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewGroupSchema

`func NewGroupSchema() *GroupSchema`

NewGroupSchema instantiates a new GroupSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSchemaWithDefaults

`func NewGroupSchemaWithDefaults() *GroupSchema`

NewGroupSchemaWithDefaults instantiates a new GroupSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *GroupSchema) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *GroupSchema) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *GroupSchema) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *GroupSchema) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetCreated

`func (o *GroupSchema) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupSchema) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupSchema) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GroupSchema) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDefinitions

`func (o *GroupSchema) GetDefinitions() GroupSchemaDefinitions`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *GroupSchema) GetDefinitionsOk() (*GroupSchemaDefinitions, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *GroupSchema) SetDefinitions(v GroupSchemaDefinitions)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *GroupSchema) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetDescription

`func (o *GroupSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *GroupSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GroupSchema) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GroupSchema) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GroupSchema) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GroupSchema) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *GroupSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *GroupSchema) GetProperties() UserSchemaProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GroupSchema) GetPropertiesOk() (*UserSchemaProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GroupSchema) SetProperties(v UserSchemaProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GroupSchema) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTitle

`func (o *GroupSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GroupSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GroupSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GroupSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *GroupSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *GroupSchema) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GroupSchema) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GroupSchema) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GroupSchema) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


