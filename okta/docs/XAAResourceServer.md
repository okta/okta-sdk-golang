# XAAResourceServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudTenantRequired** | Pointer to **bool** | Indicates whether the ID-JAG audience for this resource server must carry a tenant claim. When &#x60;true&#x60;, a tenant value is collected from the customer when the app is configured and is included in the ID-JAG JWT audience. When absent, this property defaults to &#x60;false&#x60;, which is the case for single-tenant resource servers.  | [optional] [readonly] 
**CimdSupported** | Pointer to **bool** | Indicates whether this resource server supports Client ID Metadata Documents (CIMD)  | [optional] [readonly] 
**Issuer** | Pointer to **string** | The issuer URL that identifies the authorization server for the app (resource server).  The &#x60;issuer&#x60; property from the OIN catalog integration can be an expression that supports the [Okta Expression Language&#39;s app properties](https://developer.okta.com/docs/reference/okta-expression-language/#application-properties), and contains the app properties that represent the customer tenant. For example, when the &#x60;issuer&#x60; on the OIN catalog integration is &#x60;https://{app.subdomain}.example.com/&#x60;, the app instance&#39;s &#x60;issuer&#x60; property is resolved to &#x60;https://mydomain.example.com&#x60; if the &#x60;subdomain&#x60; app property variable is set to &#x60;mydomain&#x60;. The fully resolved value is returned.  | [optional] [readonly] 
**ProtectedResources** | Pointer to **[]string** | The resource URLs protected by this resource server | [optional] [readonly] 
**Scopes** | Pointer to **[]string** | The OAuth 2.0 scopes that this resource server accepts | [optional] [readonly] 

## Methods

### NewXAAResourceServer

`func NewXAAResourceServer() *XAAResourceServer`

NewXAAResourceServer instantiates a new XAAResourceServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXAAResourceServerWithDefaults

`func NewXAAResourceServerWithDefaults() *XAAResourceServer`

NewXAAResourceServerWithDefaults instantiates a new XAAResourceServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudTenantRequired

`func (o *XAAResourceServer) GetAudTenantRequired() bool`

GetAudTenantRequired returns the AudTenantRequired field if non-nil, zero value otherwise.

### GetAudTenantRequiredOk

`func (o *XAAResourceServer) GetAudTenantRequiredOk() (*bool, bool)`

GetAudTenantRequiredOk returns a tuple with the AudTenantRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudTenantRequired

`func (o *XAAResourceServer) SetAudTenantRequired(v bool)`

SetAudTenantRequired sets AudTenantRequired field to given value.

### HasAudTenantRequired

`func (o *XAAResourceServer) HasAudTenantRequired() bool`

HasAudTenantRequired returns a boolean if a field has been set.

### GetCimdSupported

`func (o *XAAResourceServer) GetCimdSupported() bool`

GetCimdSupported returns the CimdSupported field if non-nil, zero value otherwise.

### GetCimdSupportedOk

`func (o *XAAResourceServer) GetCimdSupportedOk() (*bool, bool)`

GetCimdSupportedOk returns a tuple with the CimdSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCimdSupported

`func (o *XAAResourceServer) SetCimdSupported(v bool)`

SetCimdSupported sets CimdSupported field to given value.

### HasCimdSupported

`func (o *XAAResourceServer) HasCimdSupported() bool`

HasCimdSupported returns a boolean if a field has been set.

### GetIssuer

`func (o *XAAResourceServer) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *XAAResourceServer) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *XAAResourceServer) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *XAAResourceServer) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetProtectedResources

`func (o *XAAResourceServer) GetProtectedResources() []string`

GetProtectedResources returns the ProtectedResources field if non-nil, zero value otherwise.

### GetProtectedResourcesOk

`func (o *XAAResourceServer) GetProtectedResourcesOk() (*[]string, bool)`

GetProtectedResourcesOk returns a tuple with the ProtectedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedResources

`func (o *XAAResourceServer) SetProtectedResources(v []string)`

SetProtectedResources sets ProtectedResources field to given value.

### HasProtectedResources

`func (o *XAAResourceServer) HasProtectedResources() bool`

HasProtectedResources returns a boolean if a field has been set.

### GetScopes

`func (o *XAAResourceServer) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *XAAResourceServer) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *XAAResourceServer) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *XAAResourceServer) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


