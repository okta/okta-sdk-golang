# IdpPolicyRuleActionProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | IdP types of &#x60;OKTA&#x60;, &#x60;AgentlessDSSO&#x60;, and &#x60;IWA&#x60; don&#39;t require an ID. | [optional] 
**Name** | Pointer to **string** | Provider &#x60;name&#x60; in Okta. Optional. Supported in &#x60;IDENTITY ENGINE&#x60;. | [optional] 
**Type** | Pointer to **string** | The IdP object&#39;s &#x60;type&#x60; property identifies the social or enterprise IdP used for authentication. Each IdP uses a specific protocol, therefore the &#x60;protocol&#x60; object must correspond with the IdP &#x60;type&#x60;. If the protocol is OAuth 2.0-based, the &#x60;protocol&#x60; object&#39;s &#x60;scopes&#x60; property must also correspond with the scopes supported by the IdP &#x60;type&#x60;. For policy actions supported by each IdP type, see [IdP type policy actions](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path&#x3D;policy&amp;t&#x3D;request).  | Type               | Description                                                                                                                                           | Corresponding protocol | Corresponding protocol scopes                                         | | ------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------- | --------------------------------------------------------------------  | | &#x60;AMAZON&#x60;           | [Amazon](https://developer.amazon.com/settings/console/registration?return_to&#x3D;/)&amp;nbsp;as the IdP                                        | OpenID Connect         | &#x60;profile&#x60;, &#x60;profile:user_id&#x60;                                          | | &#x60;APPLE&#x60;            | [Apple](https://developer.apple.com/sign-in-with-apple/)&amp;nbsp;as the IdP                                                                | OpenID Connect         | &#x60;names&#x60;, &#x60;email&#x60;, &#x60;openid&#x60;                                            | | &#x60;DISCORD&#x60;          | [Discord](https://discord.com/login)&amp;nbsp;as the IdP                                                                                    | OAuth 2.0              | &#x60;identify&#x60;, &#x60;email&#x60;                                                   | | &#x60;FACEBOOK&#x60;         | [Facebook](https://developers.facebook.com)&amp;nbsp;as the IdP                                                                             | OAuth 2.0              | &#x60;public_profile&#x60;, &#x60;email&#x60;                                             | | &#x60;GITHUB&#x60;           | [GitHub](https://github.com/join)&amp;nbsp;as the IdP                                                                                       | OAuth 2.0              | &#x60;user&#x60;                                                                | | &#x60;GITLAB&#x60;           | [GitLab](https://gitlab.com/users/sign_in)&amp;nbsp;as the IdP                                                                              | OpenID Connect         | &#x60;openid&#x60;, &#x60;read_user&#x60;, &#x60;profile&#x60;, &#x60;email&#x60;                             | | &#x60;GOOGLE&#x60;           | [Google](https://accounts.google.com/signup)&amp;nbsp;as the IdP                                                                            | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                          | | &#x60;IDV_PERSONA&#x60;      | [Persona](https://app.withpersona.com/dashboard/login)&amp;nbsp;as the IDV IdP                                                              | ID verification        |                                                                       | | &#x60;IDV_CLEAR&#x60;        | [CLEAR Verified](https://www.clearme.com/)&amp;nbsp;as the IDV IdP                                                                          | ID verification        | &#x60;openid&#x60;, &#x60;profile&#x60;, &#x60;identity_assurance&#x60;                             | | &#x60;IDV_INCODE&#x60;       | [Incode](https://incode.com/)&amp;nbsp;as the IDV IdP                                                                                       | ID verification        | &#x60;openid&#x60;, &#x60;profile&#x60;, &#x60;identity_assurance&#x60;                             | | &#x60;LINKEDIN&#x60;         | [LinkedIn](https://developer.linkedin.com/)&amp;nbsp;as the IdP                                                                             | OAuth 2.0              | &#x60;r_emailaddress&#x60;, &#x60;r_liteprofile&#x60;                                     | | &#x60;LOGINGOV&#x60;         | [Login.gov](https://developers.login.gov/)&amp;nbsp;as the IdP                                                                              | OpenID Connect         | &#x60;email&#x60;, &#x60;profile&#x60;, &#x60;profile:name&#x60;                                    | | &#x60;LOGINGOV_SANDBOX&#x60; | [Login.gov&#39;s identity sandbox](https://developers.login.gov/testing/)&amp;nbsp;as the IdP                                                   | OpenID Connect         | &#x60;email&#x60;, &#x60;profile&#x60;, &#x60;profile:name&#x60;                                    | | &#x60;MICROSOFT&#x60;        | [Microsoft Enterprise SSO](https://azure.microsoft.com/)&amp;nbsp;as the IdP                                                                | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;, &#x60;https://graph.microsoft.com/User.Read&#x60; | | &#x60;OIDC&#x60;             | IdP that supports [OpenID Connect](https://openid.net/specs/openid-connect-core-1_0.html)                                               | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                          | | &#x60;PAYPAL&#x60;           | [Paypal](https://www.paypal.com/signin)&amp;nbsp;as the IdP                                                                                 | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                          | | &#x60;PAYPAL_SANDBOX&#x60;   | [Paypal Sandbox](https://developer.paypal.com/tools/sandbox/)&amp;nbsp;as the IdP                                                           | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                          | | &#x60;SALESFORCE&#x60;       | [SalesForce](https://login.salesforce.com/)&amp;nbsp;as the IdP                                                                             | OAuth 2.0              | &#x60;id&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                              | | &#x60;SAML2&#x60;            | Enterprise IdP that supports the [SAML 2.0 Web Browser SSO Profile](https://docs.oasis-open.org/security/saml/v2.0/saml-profiles-2.0-os.pdf)| SAML 2.0  |                                                                                | | &#x60;SPOTIFY&#x60;          | [Spotify](https://developer.spotify.com/)&amp;nbsp;as the IdP                                                                               | OpenID Connect         | &#x60;user-read-email&#x60;, &#x60;user-read-private&#x60;                                | | &#x60;X509&#x60;             | [Smart Card IdP](https://tools.ietf.org/html/rfc5280)                                                                                   | Mutual TLS             |                                                                       | | &#x60;XERO&#x60;             | [Xero](https://www.xero.com/us/signup/api/)&amp;nbsp;as the IdP                                                                             | OpenID Connect         | &#x60;openid&#x60;, &#x60;profile&#x60;, &#x60;email&#x60;                                          | | &#x60;YAHOO&#x60;            | [Yahoo](https://login.yahoo.com/)&amp;nbsp;as the IdP                                                                                       | OpenID Connect         | &#x60;openid&#x60;, &#x60;profile&#x60;, &#x60;email&#x60;                                          | | &#x60;YAHOOJP&#x60;          | [Yahoo Japan](https://login.yahoo.co.jp/config/login)&amp;nbsp;as the IdP                                                                   | OpenID Connect         | &#x60;openid&#x60;, &#x60;profile&#x60;, &#x60;email&#x60;                                          | | &#x60;OKTA_INTEGRATION&#x60;             | IdP that supports the [OpenID Connect](https://openid.net/specs/openid-connect-core-1_0.html) Org2Org IdP                                               | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                          | | [optional] 

## Methods

### NewIdpPolicyRuleActionProvider

`func NewIdpPolicyRuleActionProvider() *IdpPolicyRuleActionProvider`

NewIdpPolicyRuleActionProvider instantiates a new IdpPolicyRuleActionProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpPolicyRuleActionProviderWithDefaults

`func NewIdpPolicyRuleActionProviderWithDefaults() *IdpPolicyRuleActionProvider`

NewIdpPolicyRuleActionProviderWithDefaults instantiates a new IdpPolicyRuleActionProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdpPolicyRuleActionProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdpPolicyRuleActionProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdpPolicyRuleActionProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdpPolicyRuleActionProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdpPolicyRuleActionProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpPolicyRuleActionProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpPolicyRuleActionProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdpPolicyRuleActionProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IdpPolicyRuleActionProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdpPolicyRuleActionProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdpPolicyRuleActionProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdpPolicyRuleActionProvider) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


