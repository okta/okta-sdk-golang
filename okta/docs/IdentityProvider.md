# IdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique key for the IdP | [optional] [readonly] 
**IssuerMode** | Pointer to **string** | Indicates whether Okta uses the original Okta org domain URL or a custom domain URL in the request to the social IdP | [optional] [default to "DYNAMIC"]
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**Name** | Pointer to **string** | Unique name for the IdP | [optional] 
**Policy** | Pointer to [**IdentityProviderPolicy**](IdentityProviderPolicy.md) |  | [optional] 
**Properties** | Pointer to [**NullableIdentityProviderProperties**](IdentityProviderProperties.md) |  | [optional] 
**Protocol** | Pointer to [**IdentityProviderProtocol**](IdentityProviderProtocol.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The IdP object&#39;s &#x60;type&#x60; property identifies the social or enterprise IdP used for authentication. Each IdP uses a specific protocol, therefore the &#x60;protocol&#x60; object must correspond with the IdP &#x60;type&#x60;. If the protocol is OAuth 2.0-based, the &#x60;protocol&#x60; object&#39;s &#x60;scopes&#x60; property must also correspond with the scopes supported by the IdP &#x60;type&#x60;. For policy actions supported by each IdP type, see [IdP type policy actions](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path&#x3D;policy&amp;t&#x3D;request).  | Type               | Description                                                                                                                                           | Corresponding protocol | Corresponding protocol scopes                                         | | ------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------- | --------------------------------------------------------------------  | | &#x60;AMAZON&#x60;           | [Amazon](https://developer.amazon.com/settings/console/registration?return_to&#x3D;/)&amp;nbsp;as the IdP                                        | OpenID Connect         | &#x60;profile&#x60;, &#x60;profile:user_id&#x60;                                          | | &#x60;APPLE&#x60;            | [Apple](https://developer.apple.com/sign-in-with-apple/)&amp;nbsp;as the IdP                                                                | OpenID Connect         | &#x60;names&#x60;, &#x60;email&#x60;, &#x60;openid&#x60;                                            | | &#x60;DISCORD&#x60;          | [Discord](https://discord.com/login)&amp;nbsp;as the IdP                                                                                    | OAuth 2.0              | &#x60;identify&#x60;, &#x60;email&#x60;                                                   | | &#x60;FACEBOOK&#x60;         | [Facebook](https://developers.facebook.com)&amp;nbsp;as the IdP                                                                             | OAuth 2.0              | &#x60;public_profile&#x60;, &#x60;email&#x60;                                             | | &#x60;GITHUB&#x60;           | [GitHub](https://github.com/join)&amp;nbsp;as the IdP                                                                                       | OAuth 2.0              | &#x60;user&#x60;                                                                | | &#x60;GITLAB&#x60;           | [GitLab](https://gitlab.com/users/sign_in)&amp;nbsp;as the IdP                                                                              | OpenID Connect         | &#x60;openid&#x60;, &#x60;read_user&#x60;, &#x60;profile&#x60;, &#x60;email&#x60;                             | | &#x60;GOOGLE&#x60;           | [Google](https://accounts.google.com/signup)&amp;nbsp;as the IdP                                                                            | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                          | | &#x60;IDV_PERSONA&#x60;      | [Persona](https://app.withpersona.com/dashboard/login)&amp;nbsp;as the IDV IdP                                                              | ID verification        |                                                                       | | &#x60;IDV_CLEAR&#x60;        | [CLEAR Verified](https://www.clearme.com/)&amp;nbsp;as the IDV IdP                                                                          | ID verification        | &#x60;openid&#x60;, &#x60;profile&#x60;, &#x60;identity_assurance&#x60;                             | | &#x60;IDV_INCODE&#x60;       | [Incode](https://incode.com/)&amp;nbsp;as the IDV IdP                                                                                       | ID verification        | &#x60;openid&#x60;, &#x60;profile&#x60;, &#x60;identity_assurance&#x60;                             | | &#x60;LINKEDIN&#x60;         | [LinkedIn](https://developer.linkedin.com/)&amp;nbsp;as the IdP                                                                             | OAuth 2.0              | &#x60;r_emailaddress&#x60;, &#x60;r_liteprofile&#x60;                                     | | &#x60;LOGINGOV&#x60;         | [Login.gov](https://developers.login.gov/)&amp;nbsp;as the IdP                                                                              | OpenID Connect         | &#x60;email&#x60;, &#x60;profile&#x60;, &#x60;profile:name&#x60;                                    | | &#x60;LOGINGOV_SANDBOX&#x60; | [Login.gov&#39;s identity sandbox](https://developers.login.gov/testing/)&amp;nbsp;as the IdP                                                   | OpenID Connect         | &#x60;email&#x60;, &#x60;profile&#x60;, &#x60;profile:name&#x60;                                    | | &#x60;MICROSOFT&#x60;        | [Microsoft Enterprise SSO](https://azure.microsoft.com/)&amp;nbsp;as the IdP                                                                | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;, &#x60;https://graph.microsoft.com/User.Read&#x60; | | &#x60;OIDC&#x60;             | IdP that supports [OpenID Connect](https://openid.net/specs/openid-connect-core-1_0.html)                                               | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                          | | &#x60;PAYPAL&#x60;           | [Paypal](https://www.paypal.com/signin)&amp;nbsp;as the IdP                                                                                 | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                          | | &#x60;PAYPAL_SANDBOX&#x60;   | [Paypal Sandbox](https://developer.paypal.com/tools/sandbox/)&amp;nbsp;as the IdP                                                           | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                          | | &#x60;SALESFORCE&#x60;       | [SalesForce](https://login.salesforce.com/)&amp;nbsp;as the IdP                                                                             | OAuth 2.0              | &#x60;id&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                              | | &#x60;SAML2&#x60;            | Enterprise IdP that supports the [SAML 2.0 Web Browser SSO Profile](https://docs.oasis-open.org/security/saml/v2.0/saml-profiles-2.0-os.pdf)| SAML 2.0  |                                                                                | | &#x60;SPOTIFY&#x60;          | [Spotify](https://developer.spotify.com/)&amp;nbsp;as the IdP                                                                               | OpenID Connect         | &#x60;user-read-email&#x60;, &#x60;user-read-private&#x60;                                | | &#x60;X509&#x60;             | [Smart Card IdP](https://tools.ietf.org/html/rfc5280)                                                                                   | Mutual TLS             |                                                                       | | &#x60;XERO&#x60;             | [Xero](https://www.xero.com/us/signup/api/)&amp;nbsp;as the IdP                                                                             | OpenID Connect         | &#x60;openid&#x60;, &#x60;profile&#x60;, &#x60;email&#x60;                                          | | &#x60;YAHOO&#x60;            | [Yahoo](https://login.yahoo.com/)&amp;nbsp;as the IdP                                                                                       | OpenID Connect         | &#x60;openid&#x60;, &#x60;profile&#x60;, &#x60;email&#x60;                                          | | &#x60;YAHOOJP&#x60;          | [Yahoo Japan](https://login.yahoo.co.jp/config/login)&amp;nbsp;as the IdP                                                                   | OpenID Connect         | &#x60;openid&#x60;, &#x60;profile&#x60;, &#x60;email&#x60;                                          | | &#x60;OKTA_INTEGRATION&#x60;             | IdP that supports the [OpenID Connect](https://openid.net/specs/openid-connect-core-1_0.html) Org2Org IdP                                               | OpenID Connect         | &#x60;openid&#x60;, &#x60;email&#x60;, &#x60;profile&#x60;                                          | | [optional] 
**Links** | Pointer to [**IdentityProviderLinks**](IdentityProviderLinks.md) |  | [optional] 

## Methods

### NewIdentityProvider

`func NewIdentityProvider() *IdentityProvider`

NewIdentityProvider instantiates a new IdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderWithDefaults

`func NewIdentityProviderWithDefaults() *IdentityProvider`

NewIdentityProviderWithDefaults instantiates a new IdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IdentityProvider) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IdentityProvider) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IdentityProvider) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IdentityProvider) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *IdentityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuerMode

`func (o *IdentityProvider) GetIssuerMode() string`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *IdentityProvider) GetIssuerModeOk() (*string, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *IdentityProvider) SetIssuerMode(v string)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *IdentityProvider) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetLastUpdated

`func (o *IdentityProvider) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IdentityProvider) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IdentityProvider) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *IdentityProvider) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *IdentityProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *IdentityProvider) GetPolicy() IdentityProviderPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *IdentityProvider) GetPolicyOk() (*IdentityProviderPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *IdentityProvider) SetPolicy(v IdentityProviderPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *IdentityProvider) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProperties

`func (o *IdentityProvider) GetProperties() IdentityProviderProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IdentityProvider) GetPropertiesOk() (*IdentityProviderProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IdentityProvider) SetProperties(v IdentityProviderProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IdentityProvider) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *IdentityProvider) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *IdentityProvider) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetProtocol

`func (o *IdentityProvider) GetProtocol() IdentityProviderProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IdentityProvider) GetProtocolOk() (*IdentityProviderProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IdentityProvider) SetProtocol(v IdentityProviderProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IdentityProvider) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityProvider) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityProvider) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityProvider) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityProvider) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *IdentityProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityProvider) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *IdentityProvider) GetLinks() IdentityProviderLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdentityProvider) GetLinksOk() (*IdentityProviderLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdentityProvider) SetLinks(v IdentityProviderLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IdentityProvider) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


