# UsersLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hints** | Pointer to [**HrefHints**](HrefHints.md) |  | [optional] 
**Href** | **string** | Link URI | 
**Name** | Pointer to **string** | Link name | [optional] [readonly] 
**Templated** | Pointer to **bool** | Indicates whether the link object&#39;s &#x60;href&#x60; property is a URI template. | [optional] [readonly] 
**Type** | Pointer to **string** | The media type of the link. If omitted, it is implicitly &#x60;application/json&#x60;. | [optional] [readonly] 

## Methods

### NewUsersLink

`func NewUsersLink(href string, ) *UsersLink`

NewUsersLink instantiates a new UsersLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersLinkWithDefaults

`func NewUsersLinkWithDefaults() *UsersLink`

NewUsersLinkWithDefaults instantiates a new UsersLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHints

`func (o *UsersLink) GetHints() HrefHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *UsersLink) GetHintsOk() (*HrefHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *UsersLink) SetHints(v HrefHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *UsersLink) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetHref

`func (o *UsersLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UsersLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UsersLink) SetHref(v string)`

SetHref sets Href field to given value.


### GetName

`func (o *UsersLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsersLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsersLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsersLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplated

`func (o *UsersLink) GetTemplated() bool`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *UsersLink) GetTemplatedOk() (*bool, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *UsersLink) SetTemplated(v bool)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *UsersLink) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.

### GetType

`func (o *UsersLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UsersLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UsersLink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UsersLink) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


