# SamlAttributeStatementExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the attribute in your app. The attribute name must be unique across all user and group attribute statements. | [optional] 
**Namespace** | Pointer to **string** | The name format of the attribute. Supported values: | [optional] 
**Type** | Pointer to **string** | The type of attribute statements object | [optional] 
**Values** | Pointer to **[]string** | The attribute values (supports [Okta Expression Language](https://developer.okta.com/docs/reference/okta-expression-language/)) | [optional] 

## Methods

### NewSamlAttributeStatementExpression

`func NewSamlAttributeStatementExpression() *SamlAttributeStatementExpression`

NewSamlAttributeStatementExpression instantiates a new SamlAttributeStatementExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAttributeStatementExpressionWithDefaults

`func NewSamlAttributeStatementExpressionWithDefaults() *SamlAttributeStatementExpression`

NewSamlAttributeStatementExpressionWithDefaults instantiates a new SamlAttributeStatementExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SamlAttributeStatementExpression) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlAttributeStatementExpression) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlAttributeStatementExpression) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SamlAttributeStatementExpression) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SamlAttributeStatementExpression) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SamlAttributeStatementExpression) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SamlAttributeStatementExpression) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SamlAttributeStatementExpression) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetType

`func (o *SamlAttributeStatementExpression) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SamlAttributeStatementExpression) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SamlAttributeStatementExpression) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SamlAttributeStatementExpression) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValues

`func (o *SamlAttributeStatementExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SamlAttributeStatementExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SamlAttributeStatementExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *SamlAttributeStatementExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


