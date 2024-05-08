# UIElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label name for the UI element | [optional] 
**Options** | Pointer to [**UIElementOptions**](UIElementOptions.md) |  | [optional] 
**Scope** | Pointer to **string** | Specifies the property bound to the input field. It must follow the format &#x60;#/properties/PROPERTY_NAME&#x60; where &#x60;PROPERTY_NAME&#x60; is a variable name for an attribute in &#x60;profile editor&#x60;. | [optional] 
**Type** | Pointer to **string** | Specifies the relationship between this input element and &#x60;scope&#x60;. The &#x60;Control&#x60; value specifies that this input controls the value represented by &#x60;scope&#x60;. | [optional] 

## Methods

### NewUIElement

`func NewUIElement() *UIElement`

NewUIElement instantiates a new UIElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUIElementWithDefaults

`func NewUIElementWithDefaults() *UIElement`

NewUIElementWithDefaults instantiates a new UIElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UIElement) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UIElement) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UIElement) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UIElement) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOptions

`func (o *UIElement) GetOptions() UIElementOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UIElement) GetOptionsOk() (*UIElementOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UIElement) SetOptions(v UIElementOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UIElement) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetScope

`func (o *UIElement) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UIElement) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UIElement) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UIElement) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetType

`func (o *UIElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UIElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UIElement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UIElement) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


