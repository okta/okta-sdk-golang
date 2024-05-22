# TestInfoOidcTestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idp** | Pointer to **bool** | Read only.&lt;br&gt;Indicates if your integration supports IdP-initiated sign-in flows. If [&#x60;sso.oidc.initiateLoginUri&#x60;](/openapi/okta-management/management/tag/YourOinIntegrations/#tag/YourOinIntegrations/operation/createSubmission!path&#x3D;sso/oidc/initiateLoginUri&amp;t&#x3D;request) is specified, this property is set to &#x60;true&#x60;. If [&#x60;sso.oidc.initiateLoginUri&#x60;](/openapi/okta-management/management/tag/YourOinIntegrations/#tag/YourOinIntegrations/operation/createSubmission!path&#x3D;sso/oidc/initiateLoginUri&amp;t&#x3D;request) isn&#39;t set for the integration submission, this property is set to &#x60;false&#x60; | [optional] [readonly] 
**Sp** | Pointer to **bool** | Read only.&lt;br&gt;Indicates if your integration supports SP-initiated sign-in flows and is always set to &#x60;true&#x60; for OIDC SSO | [optional] [readonly] 
**Jit** | Pointer to **bool** | Indicates if your integration supports Just-In-Time (JIT) provisioning | [optional] 
**SpInitiateUrl** | **string** | URL for SP-initiated sign-in flows (required if &#x60;sp &#x3D; true&#x60;) | 

## Methods

### NewTestInfoOidcTestConfiguration

`func NewTestInfoOidcTestConfiguration(spInitiateUrl string, ) *TestInfoOidcTestConfiguration`

NewTestInfoOidcTestConfiguration instantiates a new TestInfoOidcTestConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestInfoOidcTestConfigurationWithDefaults

`func NewTestInfoOidcTestConfigurationWithDefaults() *TestInfoOidcTestConfiguration`

NewTestInfoOidcTestConfigurationWithDefaults instantiates a new TestInfoOidcTestConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdp

`func (o *TestInfoOidcTestConfiguration) GetIdp() bool`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *TestInfoOidcTestConfiguration) GetIdpOk() (*bool, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *TestInfoOidcTestConfiguration) SetIdp(v bool)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *TestInfoOidcTestConfiguration) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetSp

`func (o *TestInfoOidcTestConfiguration) GetSp() bool`

GetSp returns the Sp field if non-nil, zero value otherwise.

### GetSpOk

`func (o *TestInfoOidcTestConfiguration) GetSpOk() (*bool, bool)`

GetSpOk returns a tuple with the Sp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSp

`func (o *TestInfoOidcTestConfiguration) SetSp(v bool)`

SetSp sets Sp field to given value.

### HasSp

`func (o *TestInfoOidcTestConfiguration) HasSp() bool`

HasSp returns a boolean if a field has been set.

### GetJit

`func (o *TestInfoOidcTestConfiguration) GetJit() bool`

GetJit returns the Jit field if non-nil, zero value otherwise.

### GetJitOk

`func (o *TestInfoOidcTestConfiguration) GetJitOk() (*bool, bool)`

GetJitOk returns a tuple with the Jit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJit

`func (o *TestInfoOidcTestConfiguration) SetJit(v bool)`

SetJit sets Jit field to given value.

### HasJit

`func (o *TestInfoOidcTestConfiguration) HasJit() bool`

HasJit returns a boolean if a field has been set.

### GetSpInitiateUrl

`func (o *TestInfoOidcTestConfiguration) GetSpInitiateUrl() string`

GetSpInitiateUrl returns the SpInitiateUrl field if non-nil, zero value otherwise.

### GetSpInitiateUrlOk

`func (o *TestInfoOidcTestConfiguration) GetSpInitiateUrlOk() (*string, bool)`

GetSpInitiateUrlOk returns a tuple with the SpInitiateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpInitiateUrl

`func (o *TestInfoOidcTestConfiguration) SetSpInitiateUrl(v string)`

SetSpInitiateUrl sets SpInitiateUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


