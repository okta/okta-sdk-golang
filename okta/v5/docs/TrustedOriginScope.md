# TrustedOriginScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedOktaApps** | Pointer to **[]string** | The allowed Okta apps for the Trusted Origin scope | [optional] 
**Type** | Pointer to **string** | The scope type. Supported values: When you use &#x60;IFRAME_EMBED&#x60; as the scope type, leave the allowedOktaApps property empty to allow iFrame embedding of only Okta sign-in pages. Include &#x60;OKTA_ENDUSER&#x60; as a value for the allowedOktaApps property to allow iFrame embedding of both Okta sign-in pages and the Okta End-User Dashboard.  | [optional] 

## Methods

### NewTrustedOriginScope

`func NewTrustedOriginScope() *TrustedOriginScope`

NewTrustedOriginScope instantiates a new TrustedOriginScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedOriginScopeWithDefaults

`func NewTrustedOriginScopeWithDefaults() *TrustedOriginScope`

NewTrustedOriginScopeWithDefaults instantiates a new TrustedOriginScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedOktaApps

`func (o *TrustedOriginScope) GetAllowedOktaApps() []string`

GetAllowedOktaApps returns the AllowedOktaApps field if non-nil, zero value otherwise.

### GetAllowedOktaAppsOk

`func (o *TrustedOriginScope) GetAllowedOktaAppsOk() (*[]string, bool)`

GetAllowedOktaAppsOk returns a tuple with the AllowedOktaApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOktaApps

`func (o *TrustedOriginScope) SetAllowedOktaApps(v []string)`

SetAllowedOktaApps sets AllowedOktaApps field to given value.

### HasAllowedOktaApps

`func (o *TrustedOriginScope) HasAllowedOktaApps() bool`

HasAllowedOktaApps returns a boolean if a field has been set.

### GetType

`func (o *TrustedOriginScope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrustedOriginScope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrustedOriginScope) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TrustedOriginScope) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


