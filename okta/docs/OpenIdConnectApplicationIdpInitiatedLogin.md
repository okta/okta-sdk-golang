# OpenIdConnectApplicationIdpInitiatedLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultScope** | Pointer to **[]string** | The scopes to use for the request when &#x60;mode&#x60; is &#x60;OKTA&#x60; | [optional] 
**Mode** | **string** | The mode to use for the IdP-initiated sign-in flow. For &#x60;OKTA&#x60; or &#x60;SPEC&#x60; modes, the client must have an &#x60;initiate_login_uri&#x60; registered. &gt; **Note:** For web and SPA apps, if the mode is &#x60;SPEC&#x60; or &#x60;OKTA&#x60;, you must set &#x60;grant_types&#x60; to &#x60;authorization_code&#x60;, &#x60;implicit&#x60;, or &#x60;interaction_code&#x60;.  | 

## Methods

### NewOpenIdConnectApplicationIdpInitiatedLogin

`func NewOpenIdConnectApplicationIdpInitiatedLogin(mode string, ) *OpenIdConnectApplicationIdpInitiatedLogin`

NewOpenIdConnectApplicationIdpInitiatedLogin instantiates a new OpenIdConnectApplicationIdpInitiatedLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectApplicationIdpInitiatedLoginWithDefaults

`func NewOpenIdConnectApplicationIdpInitiatedLoginWithDefaults() *OpenIdConnectApplicationIdpInitiatedLogin`

NewOpenIdConnectApplicationIdpInitiatedLoginWithDefaults instantiates a new OpenIdConnectApplicationIdpInitiatedLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultScope

`func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetDefaultScope() []string`

GetDefaultScope returns the DefaultScope field if non-nil, zero value otherwise.

### GetDefaultScopeOk

`func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetDefaultScopeOk() (*[]string, bool)`

GetDefaultScopeOk returns a tuple with the DefaultScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultScope

`func (o *OpenIdConnectApplicationIdpInitiatedLogin) SetDefaultScope(v []string)`

SetDefaultScope sets DefaultScope field to given value.

### HasDefaultScope

`func (o *OpenIdConnectApplicationIdpInitiatedLogin) HasDefaultScope() bool`

HasDefaultScope returns a boolean if a field has been set.

### GetMode

`func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OpenIdConnectApplicationIdpInitiatedLogin) SetMode(v string)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


