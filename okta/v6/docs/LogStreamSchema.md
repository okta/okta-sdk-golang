# LogStreamSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | JSON schema version identifier | [optional] [readonly] 
**ErrorMessage** | Pointer to **map[string]interface{}** | A collection of error messages for individual properties in the schema. Okta implements a subset of [ajv-errors](https://github.com/ajv-validator/ajv-errors). | [optional] 
**Id** | Pointer to **string** | URI of log stream schema | [optional] [readonly] 
**OneOf** | Pointer to [**[]UserSchemaAttributeEnum**](UserSchemaAttributeEnum.md) | Non-empty array of valid JSON schemas.  Okta only supports &#x60;oneOf&#x60; for specifying display names for an &#x60;enum&#x60;. Each schema has the following format:  &#x60;&#x60;&#x60; {   \&quot;const\&quot;: \&quot;enumValue\&quot;,   \&quot;title\&quot;: \&quot;display name\&quot; } &#x60;&#x60;&#x60; | [optional] 
**Pattern** | Pointer to **string** | For &#x60;string&#x60; log stream schema property type, specifies the regular expression used to validate the property | [optional] 
**Properties** | Pointer to **map[string]interface{}** | log stream schema properties object | [optional] 
**Required** | Pointer to **[]string** | Required properties for this log stream schema object | [optional] 
**Title** | Pointer to **string** | Name of the log streaming integration | [optional] 
**Type** | Pointer to **string** | Type of log stream schema property | [optional] [readonly] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

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

### GetOneOf

`func (o *LogStreamSchema) GetOneOf() []UserSchemaAttributeEnum`

GetOneOf returns the OneOf field if non-nil, zero value otherwise.

### GetOneOfOk

`func (o *LogStreamSchema) GetOneOfOk() (*[]UserSchemaAttributeEnum, bool)`

GetOneOfOk returns a tuple with the OneOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneOf

`func (o *LogStreamSchema) SetOneOf(v []UserSchemaAttributeEnum)`

SetOneOf sets OneOf field to given value.

### HasOneOf

`func (o *LogStreamSchema) HasOneOf() bool`

HasOneOf returns a boolean if a field has been set.

### SetOneOfNil

`func (o *LogStreamSchema) SetOneOfNil(b bool)`

 SetOneOfNil sets the value for OneOf to be an explicit nil

### UnsetOneOf
`func (o *LogStreamSchema) UnsetOneOf()`

UnsetOneOf ensures that no value is present for OneOf, not even an explicit nil
### GetPattern

`func (o *LogStreamSchema) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *LogStreamSchema) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *LogStreamSchema) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *LogStreamSchema) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

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

`func (o *LogStreamSchema) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LogStreamSchema) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LogStreamSchema) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LogStreamSchema) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


