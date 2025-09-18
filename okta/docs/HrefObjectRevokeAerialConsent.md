# HrefObjectRevokeAerialConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hints** | Pointer to [**HrefHints**](HrefHints.md) |  | [optional] 
**Href** | **string** | Link URI | 
**Name** | Pointer to **string** | Link name | [optional] [readonly] 
**Templated** | Pointer to **bool** | Indicates whether the link object&#39;s &#x60;href&#x60; property is a URI template. | [optional] [readonly] 
**Type** | Pointer to **string** | The media type of the link. If omitted, it is implicitly &#x60;application/json&#x60;. | [optional] [readonly] 

## Methods

### NewHrefObjectRevokeAerialConsent

`func NewHrefObjectRevokeAerialConsent(href string, ) *HrefObjectRevokeAerialConsent`

NewHrefObjectRevokeAerialConsent instantiates a new HrefObjectRevokeAerialConsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHrefObjectRevokeAerialConsentWithDefaults

`func NewHrefObjectRevokeAerialConsentWithDefaults() *HrefObjectRevokeAerialConsent`

NewHrefObjectRevokeAerialConsentWithDefaults instantiates a new HrefObjectRevokeAerialConsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHints

`func (o *HrefObjectRevokeAerialConsent) GetHints() HrefHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *HrefObjectRevokeAerialConsent) GetHintsOk() (*HrefHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *HrefObjectRevokeAerialConsent) SetHints(v HrefHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *HrefObjectRevokeAerialConsent) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetHref

`func (o *HrefObjectRevokeAerialConsent) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *HrefObjectRevokeAerialConsent) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *HrefObjectRevokeAerialConsent) SetHref(v string)`

SetHref sets Href field to given value.


### GetName

`func (o *HrefObjectRevokeAerialConsent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HrefObjectRevokeAerialConsent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HrefObjectRevokeAerialConsent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HrefObjectRevokeAerialConsent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplated

`func (o *HrefObjectRevokeAerialConsent) GetTemplated() bool`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *HrefObjectRevokeAerialConsent) GetTemplatedOk() (*bool, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *HrefObjectRevokeAerialConsent) SetTemplated(v bool)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *HrefObjectRevokeAerialConsent) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.

### GetType

`func (o *HrefObjectRevokeAerialConsent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HrefObjectRevokeAerialConsent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HrefObjectRevokeAerialConsent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HrefObjectRevokeAerialConsent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


