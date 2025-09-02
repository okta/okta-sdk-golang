# SamlAttributeStatementGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** | The operation to filter groups based on &#x60;filterValue&#x60; | [optional] 
**FilterValue** | Pointer to **string** | Filter the groups based on a specific value. | [optional] 
**Name** | Pointer to **string** | The name of the group attribute in your app. The attribute name must be unique across all user and group attribute statements. | [optional] 
**Namespace** | Pointer to **string** | The name format of the group attribute. Supported values: | [optional] 
**Type** | Pointer to **string** | The type of attribute statements object | [optional] 

## Methods

### NewSamlAttributeStatementGroup

`func NewSamlAttributeStatementGroup() *SamlAttributeStatementGroup`

NewSamlAttributeStatementGroup instantiates a new SamlAttributeStatementGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAttributeStatementGroupWithDefaults

`func NewSamlAttributeStatementGroupWithDefaults() *SamlAttributeStatementGroup`

NewSamlAttributeStatementGroupWithDefaults instantiates a new SamlAttributeStatementGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *SamlAttributeStatementGroup) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *SamlAttributeStatementGroup) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *SamlAttributeStatementGroup) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *SamlAttributeStatementGroup) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetFilterValue

`func (o *SamlAttributeStatementGroup) GetFilterValue() string`

GetFilterValue returns the FilterValue field if non-nil, zero value otherwise.

### GetFilterValueOk

`func (o *SamlAttributeStatementGroup) GetFilterValueOk() (*string, bool)`

GetFilterValueOk returns a tuple with the FilterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValue

`func (o *SamlAttributeStatementGroup) SetFilterValue(v string)`

SetFilterValue sets FilterValue field to given value.

### HasFilterValue

`func (o *SamlAttributeStatementGroup) HasFilterValue() bool`

HasFilterValue returns a boolean if a field has been set.

### GetName

`func (o *SamlAttributeStatementGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlAttributeStatementGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlAttributeStatementGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SamlAttributeStatementGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SamlAttributeStatementGroup) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SamlAttributeStatementGroup) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SamlAttributeStatementGroup) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SamlAttributeStatementGroup) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetType

`func (o *SamlAttributeStatementGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SamlAttributeStatementGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SamlAttributeStatementGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SamlAttributeStatementGroup) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


