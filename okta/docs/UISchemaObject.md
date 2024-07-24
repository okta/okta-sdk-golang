# UISchemaObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ButtonLabel** | Pointer to **string** | Specifies the button label for the &#x60;Submit&#x60; button at the bottom of the enrollment form. | [optional] [default to "Submit"]
**Elements** | Pointer to [**UIElement**](UIElement.md) |  | [optional] 
**Label** | Pointer to **string** | Specifies the label at the top of the enrollment form under the logo. | [optional] [default to "Sign in"]
**Type** | Pointer to **string** | Specifies the type of layout | [optional] 

## Methods

### NewUISchemaObject

`func NewUISchemaObject() *UISchemaObject`

NewUISchemaObject instantiates a new UISchemaObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUISchemaObjectWithDefaults

`func NewUISchemaObjectWithDefaults() *UISchemaObject`

NewUISchemaObjectWithDefaults instantiates a new UISchemaObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtonLabel

`func (o *UISchemaObject) GetButtonLabel() string`

GetButtonLabel returns the ButtonLabel field if non-nil, zero value otherwise.

### GetButtonLabelOk

`func (o *UISchemaObject) GetButtonLabelOk() (*string, bool)`

GetButtonLabelOk returns a tuple with the ButtonLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonLabel

`func (o *UISchemaObject) SetButtonLabel(v string)`

SetButtonLabel sets ButtonLabel field to given value.

### HasButtonLabel

`func (o *UISchemaObject) HasButtonLabel() bool`

HasButtonLabel returns a boolean if a field has been set.

### GetElements

`func (o *UISchemaObject) GetElements() UIElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *UISchemaObject) GetElementsOk() (*UIElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *UISchemaObject) SetElements(v UIElement)`

SetElements sets Elements field to given value.

### HasElements

`func (o *UISchemaObject) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetLabel

`func (o *UISchemaObject) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UISchemaObject) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UISchemaObject) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UISchemaObject) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *UISchemaObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UISchemaObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UISchemaObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UISchemaObject) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


