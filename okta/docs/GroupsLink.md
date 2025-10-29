# GroupsLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hints** | Pointer to [**HrefHints**](HrefHints.md) |  | [optional] 
**Href** | **string** | Link URI | 
**Name** | Pointer to **string** | Link name | [optional] [readonly] 
**Templated** | Pointer to **bool** | Indicates whether the link object&#39;s &#x60;href&#x60; property is a URI template. | [optional] [readonly] 
**Type** | Pointer to **string** | The media type of the link. If omitted, it is implicitly &#x60;application/json&#x60;. | [optional] [readonly] 

## Methods

### NewGroupsLink

`func NewGroupsLink(href string, ) *GroupsLink`

NewGroupsLink instantiates a new GroupsLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsLinkWithDefaults

`func NewGroupsLinkWithDefaults() *GroupsLink`

NewGroupsLinkWithDefaults instantiates a new GroupsLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHints

`func (o *GroupsLink) GetHints() HrefHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *GroupsLink) GetHintsOk() (*HrefHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *GroupsLink) SetHints(v HrefHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *GroupsLink) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetHref

`func (o *GroupsLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GroupsLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GroupsLink) SetHref(v string)`

SetHref sets Href field to given value.


### GetName

`func (o *GroupsLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupsLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupsLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupsLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplated

`func (o *GroupsLink) GetTemplated() bool`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *GroupsLink) GetTemplatedOk() (*bool, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *GroupsLink) SetTemplated(v bool)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *GroupsLink) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.

### GetType

`func (o *GroupsLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupsLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupsLink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupsLink) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


