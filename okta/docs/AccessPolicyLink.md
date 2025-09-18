# AccessPolicyLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hints** | Pointer to [**HrefHints**](HrefHints.md) |  | [optional] 
**Href** | **string** | Link URI | 
**Name** | Pointer to **string** | Link name | [optional] [readonly] 
**Templated** | Pointer to **bool** | Indicates whether the link object&#39;s &#x60;href&#x60; property is a URI template. | [optional] [readonly] 
**Type** | Pointer to **string** | The media type of the link. If omitted, it is implicitly &#x60;application/json&#x60;. | [optional] [readonly] 

## Methods

### NewAccessPolicyLink

`func NewAccessPolicyLink(href string, ) *AccessPolicyLink`

NewAccessPolicyLink instantiates a new AccessPolicyLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyLinkWithDefaults

`func NewAccessPolicyLinkWithDefaults() *AccessPolicyLink`

NewAccessPolicyLinkWithDefaults instantiates a new AccessPolicyLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHints

`func (o *AccessPolicyLink) GetHints() HrefHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *AccessPolicyLink) GetHintsOk() (*HrefHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *AccessPolicyLink) SetHints(v HrefHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *AccessPolicyLink) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetHref

`func (o *AccessPolicyLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AccessPolicyLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AccessPolicyLink) SetHref(v string)`

SetHref sets Href field to given value.


### GetName

`func (o *AccessPolicyLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessPolicyLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessPolicyLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessPolicyLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplated

`func (o *AccessPolicyLink) GetTemplated() bool`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *AccessPolicyLink) GetTemplatedOk() (*bool, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *AccessPolicyLink) SetTemplated(v bool)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *AccessPolicyLink) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.

### GetType

`func (o *AccessPolicyLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessPolicyLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessPolicyLink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccessPolicyLink) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


