# WebAuthnRpId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to [**WebAuthnRpIdDomain**](WebAuthnRpIdDomain.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the RP ID is active and is used for WebAuthn operations. It can only be set to &#x60;true&#x60; once the &#x60;validationStatus&#x60; of the &#x60;domain&#x60; object is &#x60;VERIFIED&#x60;. &#x60;enabled&#x60; can only be &#x60;true&#x60; for this same &#x60;domain&#x60;. Its value must be &#x60;false&#x60; to be able to configure the &#x60;domain&#x60;. | [optional] [default to false]

## Methods

### NewWebAuthnRpId

`func NewWebAuthnRpId() *WebAuthnRpId`

NewWebAuthnRpId instantiates a new WebAuthnRpId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAuthnRpIdWithDefaults

`func NewWebAuthnRpIdWithDefaults() *WebAuthnRpId`

NewWebAuthnRpIdWithDefaults instantiates a new WebAuthnRpId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *WebAuthnRpId) GetDomain() WebAuthnRpIdDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WebAuthnRpId) GetDomainOk() (*WebAuthnRpIdDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WebAuthnRpId) SetDomain(v WebAuthnRpIdDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WebAuthnRpId) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEnabled

`func (o *WebAuthnRpId) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebAuthnRpId) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebAuthnRpId) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebAuthnRpId) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


