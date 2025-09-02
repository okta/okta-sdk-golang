# HrefObjectResourceSetResourcesLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hints** | Pointer to [**HrefHints**](HrefHints.md) |  | [optional] 
**Href** | **string** | Link URI | 
**Name** | Pointer to **string** | Link name | [optional] [readonly] 
**Templated** | Pointer to **bool** | Indicates whether the link object&#39;s &#x60;href&#x60; property is a URI template. | [optional] [readonly] 
**Type** | Pointer to **string** | The media type of the link. If omitted, it is implicitly &#x60;application/json&#x60;. | [optional] [readonly] 

## Methods

### NewHrefObjectResourceSetResourcesLink

`func NewHrefObjectResourceSetResourcesLink(href string, ) *HrefObjectResourceSetResourcesLink`

NewHrefObjectResourceSetResourcesLink instantiates a new HrefObjectResourceSetResourcesLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHrefObjectResourceSetResourcesLinkWithDefaults

`func NewHrefObjectResourceSetResourcesLinkWithDefaults() *HrefObjectResourceSetResourcesLink`

NewHrefObjectResourceSetResourcesLinkWithDefaults instantiates a new HrefObjectResourceSetResourcesLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHints

`func (o *HrefObjectResourceSetResourcesLink) GetHints() HrefHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *HrefObjectResourceSetResourcesLink) GetHintsOk() (*HrefHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *HrefObjectResourceSetResourcesLink) SetHints(v HrefHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *HrefObjectResourceSetResourcesLink) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetHref

`func (o *HrefObjectResourceSetResourcesLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *HrefObjectResourceSetResourcesLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *HrefObjectResourceSetResourcesLink) SetHref(v string)`

SetHref sets Href field to given value.


### GetName

`func (o *HrefObjectResourceSetResourcesLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HrefObjectResourceSetResourcesLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HrefObjectResourceSetResourcesLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HrefObjectResourceSetResourcesLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplated

`func (o *HrefObjectResourceSetResourcesLink) GetTemplated() bool`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *HrefObjectResourceSetResourcesLink) GetTemplatedOk() (*bool, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *HrefObjectResourceSetResourcesLink) SetTemplated(v bool)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *HrefObjectResourceSetResourcesLink) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.

### GetType

`func (o *HrefObjectResourceSetResourcesLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HrefObjectResourceSetResourcesLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HrefObjectResourceSetResourcesLink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HrefObjectResourceSetResourcesLink) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


