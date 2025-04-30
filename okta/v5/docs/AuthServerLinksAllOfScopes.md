# AuthServerLinksAllOfScopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hints** | Pointer to [**HrefHints**](HrefHints.md) |  | [optional] 
**Href** | **string** | Link URI | 
**Name** | Pointer to **string** | Link name | [optional] 
**Templated** | Pointer to **bool** | Indicates whether the Link Object&#39;s &#x60;href&#x60; property is a URI template. | [optional] 
**Type** | Pointer to **string** | The media type of the link. If omitted, it is implicitly &#x60;application/json&#x60;. | [optional] 

## Methods

### NewAuthServerLinksAllOfScopes

`func NewAuthServerLinksAllOfScopes(href string, ) *AuthServerLinksAllOfScopes`

NewAuthServerLinksAllOfScopes instantiates a new AuthServerLinksAllOfScopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthServerLinksAllOfScopesWithDefaults

`func NewAuthServerLinksAllOfScopesWithDefaults() *AuthServerLinksAllOfScopes`

NewAuthServerLinksAllOfScopesWithDefaults instantiates a new AuthServerLinksAllOfScopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHints

`func (o *AuthServerLinksAllOfScopes) GetHints() HrefHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *AuthServerLinksAllOfScopes) GetHintsOk() (*HrefHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *AuthServerLinksAllOfScopes) SetHints(v HrefHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *AuthServerLinksAllOfScopes) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetHref

`func (o *AuthServerLinksAllOfScopes) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AuthServerLinksAllOfScopes) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AuthServerLinksAllOfScopes) SetHref(v string)`

SetHref sets Href field to given value.


### GetName

`func (o *AuthServerLinksAllOfScopes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthServerLinksAllOfScopes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthServerLinksAllOfScopes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthServerLinksAllOfScopes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplated

`func (o *AuthServerLinksAllOfScopes) GetTemplated() bool`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *AuthServerLinksAllOfScopes) GetTemplatedOk() (*bool, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *AuthServerLinksAllOfScopes) SetTemplated(v bool)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *AuthServerLinksAllOfScopes) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.

### GetType

`func (o *AuthServerLinksAllOfScopes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthServerLinksAllOfScopes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthServerLinksAllOfScopes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthServerLinksAllOfScopes) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


