# LogStreamSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] [readonly] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**ErrorMessage** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Required** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewLogStreamSchema

`func NewLogStreamSchema() *LogStreamSchema`

NewLogStreamSchema instantiates a new LogStreamSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamSchemaWithDefaults

`func NewLogStreamSchemaWithDefaults() *LogStreamSchema`

NewLogStreamSchemaWithDefaults instantiates a new LogStreamSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *LogStreamSchema) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *LogStreamSchema) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *LogStreamSchema) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *LogStreamSchema) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetCreated

`func (o *LogStreamSchema) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LogStreamSchema) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LogStreamSchema) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LogStreamSchema) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetErrorMessage

`func (o *LogStreamSchema) GetErrorMessage() map[string]interface{}`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *LogStreamSchema) GetErrorMessageOk() (*map[string]interface{}, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *LogStreamSchema) SetErrorMessage(v map[string]interface{})`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *LogStreamSchema) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetId

`func (o *LogStreamSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogStreamSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogStreamSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogStreamSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *LogStreamSchema) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *LogStreamSchema) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *LogStreamSchema) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *LogStreamSchema) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *LogStreamSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogStreamSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogStreamSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogStreamSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *LogStreamSchema) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *LogStreamSchema) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *LogStreamSchema) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *LogStreamSchema) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRequired

`func (o *LogStreamSchema) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *LogStreamSchema) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *LogStreamSchema) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *LogStreamSchema) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetTitle

`func (o *LogStreamSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LogStreamSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LogStreamSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LogStreamSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *LogStreamSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogStreamSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogStreamSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LogStreamSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *LogStreamSchema) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LogStreamSchema) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LogStreamSchema) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LogStreamSchema) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


