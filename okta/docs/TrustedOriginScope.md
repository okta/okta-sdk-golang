# TrustedOriginScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedOktaApps** | Pointer to [**[]IframeEmbedScopeAllowedApps**](IframeEmbedScopeAllowedApps.md) |  | [optional] 
**Type** | Pointer to [**TrustedOriginScopeType**](TrustedOriginScopeType.md) |  | [optional] 

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

`func (o *TrustedOriginScope) GetAllowedOktaApps() []IframeEmbedScopeAllowedApps`

GetAllowedOktaApps returns the AllowedOktaApps field if non-nil, zero value otherwise.

### GetAllowedOktaAppsOk

`func (o *TrustedOriginScope) GetAllowedOktaAppsOk() (*[]IframeEmbedScopeAllowedApps, bool)`

GetAllowedOktaAppsOk returns a tuple with the AllowedOktaApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOktaApps

`func (o *TrustedOriginScope) SetAllowedOktaApps(v []IframeEmbedScopeAllowedApps)`

SetAllowedOktaApps sets AllowedOktaApps field to given value.

### HasAllowedOktaApps

`func (o *TrustedOriginScope) HasAllowedOktaApps() bool`

HasAllowedOktaApps returns a boolean if a field has been set.

### GetType

`func (o *TrustedOriginScope) GetType() TrustedOriginScopeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrustedOriginScope) GetTypeOk() (*TrustedOriginScopeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrustedOriginScope) SetType(v TrustedOriginScopeType)`

SetType sets Type field to given value.

### HasType

`func (o *TrustedOriginScope) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


