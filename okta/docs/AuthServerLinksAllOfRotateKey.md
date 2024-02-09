# AuthServerLinksAllOfRotateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hints** | Pointer to [**HrefObjectHints**](HrefObjectHints.md) |  | [optional] 
**Href** | **string** | Link URI | 
**Name** | Pointer to **string** | Link name | [optional] 
**Type** | Pointer to **string** | The media type of the link. If omitted, it is implicitly &#x60;application/json&#x60;. | [optional] 

## Methods

### NewAuthServerLinksAllOfRotateKey

`func NewAuthServerLinksAllOfRotateKey(href string, ) *AuthServerLinksAllOfRotateKey`

NewAuthServerLinksAllOfRotateKey instantiates a new AuthServerLinksAllOfRotateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthServerLinksAllOfRotateKeyWithDefaults

`func NewAuthServerLinksAllOfRotateKeyWithDefaults() *AuthServerLinksAllOfRotateKey`

NewAuthServerLinksAllOfRotateKeyWithDefaults instantiates a new AuthServerLinksAllOfRotateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHints

`func (o *AuthServerLinksAllOfRotateKey) GetHints() HrefObjectHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *AuthServerLinksAllOfRotateKey) GetHintsOk() (*HrefObjectHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *AuthServerLinksAllOfRotateKey) SetHints(v HrefObjectHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *AuthServerLinksAllOfRotateKey) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetHref

`func (o *AuthServerLinksAllOfRotateKey) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AuthServerLinksAllOfRotateKey) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AuthServerLinksAllOfRotateKey) SetHref(v string)`

SetHref sets Href field to given value.


### GetName

`func (o *AuthServerLinksAllOfRotateKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthServerLinksAllOfRotateKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthServerLinksAllOfRotateKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthServerLinksAllOfRotateKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AuthServerLinksAllOfRotateKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthServerLinksAllOfRotateKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthServerLinksAllOfRotateKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthServerLinksAllOfRotateKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


