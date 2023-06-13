# OpenIdConnectApplicationSettingsClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationType** | Pointer to [**OpenIdConnectApplicationType**](OpenIdConnectApplicationType.md) |  | [optional] 
**ClientUri** | Pointer to **string** |  | [optional] 
**ConsentMethod** | Pointer to [**OpenIdConnectApplicationConsentMethod**](OpenIdConnectApplicationConsentMethod.md) |  | [optional] 
**GrantTypes** | Pointer to [**[]OAuthGrantType**](OAuthGrantType.md) |  | [optional] 
**IdpInitiatedLogin** | Pointer to [**OpenIdConnectApplicationIdpInitiatedLogin**](OpenIdConnectApplicationIdpInitiatedLogin.md) |  | [optional] 
**InitiateLoginUri** | Pointer to **string** |  | [optional] 
**IssuerMode** | Pointer to [**OpenIdConnectApplicationIssuerMode**](OpenIdConnectApplicationIssuerMode.md) |  | [optional] 
**Jwks** | Pointer to [**OpenIdConnectApplicationSettingsClientKeys**](OpenIdConnectApplicationSettingsClientKeys.md) |  | [optional] 
**LogoUri** | Pointer to **string** |  | [optional] 
**PolicyUri** | Pointer to **string** |  | [optional] 
**PostLogoutRedirectUris** | Pointer to **[]string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**RefreshToken** | Pointer to [**OpenIdConnectApplicationSettingsRefreshToken**](OpenIdConnectApplicationSettingsRefreshToken.md) |  | [optional] 
**ResponseTypes** | Pointer to [**[]OAuthResponseType**](OAuthResponseType.md) |  | [optional] 
**TosUri** | Pointer to **string** |  | [optional] 
**WildcardRedirect** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenIdConnectApplicationSettingsClient

`func NewOpenIdConnectApplicationSettingsClient() *OpenIdConnectApplicationSettingsClient`

NewOpenIdConnectApplicationSettingsClient instantiates a new OpenIdConnectApplicationSettingsClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectApplicationSettingsClientWithDefaults

`func NewOpenIdConnectApplicationSettingsClientWithDefaults() *OpenIdConnectApplicationSettingsClient`

NewOpenIdConnectApplicationSettingsClientWithDefaults instantiates a new OpenIdConnectApplicationSettingsClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationType

`func (o *OpenIdConnectApplicationSettingsClient) GetApplicationType() OpenIdConnectApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *OpenIdConnectApplicationSettingsClient) GetApplicationTypeOk() (*OpenIdConnectApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *OpenIdConnectApplicationSettingsClient) SetApplicationType(v OpenIdConnectApplicationType)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *OpenIdConnectApplicationSettingsClient) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetClientUri

`func (o *OpenIdConnectApplicationSettingsClient) GetClientUri() string`

GetClientUri returns the ClientUri field if non-nil, zero value otherwise.

### GetClientUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetClientUriOk() (*string, bool)`

GetClientUriOk returns a tuple with the ClientUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUri

`func (o *OpenIdConnectApplicationSettingsClient) SetClientUri(v string)`

SetClientUri sets ClientUri field to given value.

### HasClientUri

`func (o *OpenIdConnectApplicationSettingsClient) HasClientUri() bool`

HasClientUri returns a boolean if a field has been set.

### GetConsentMethod

`func (o *OpenIdConnectApplicationSettingsClient) GetConsentMethod() OpenIdConnectApplicationConsentMethod`

GetConsentMethod returns the ConsentMethod field if non-nil, zero value otherwise.

### GetConsentMethodOk

`func (o *OpenIdConnectApplicationSettingsClient) GetConsentMethodOk() (*OpenIdConnectApplicationConsentMethod, bool)`

GetConsentMethodOk returns a tuple with the ConsentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentMethod

`func (o *OpenIdConnectApplicationSettingsClient) SetConsentMethod(v OpenIdConnectApplicationConsentMethod)`

SetConsentMethod sets ConsentMethod field to given value.

### HasConsentMethod

`func (o *OpenIdConnectApplicationSettingsClient) HasConsentMethod() bool`

HasConsentMethod returns a boolean if a field has been set.

### GetGrantTypes

`func (o *OpenIdConnectApplicationSettingsClient) GetGrantTypes() []OAuthGrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *OpenIdConnectApplicationSettingsClient) GetGrantTypesOk() (*[]OAuthGrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *OpenIdConnectApplicationSettingsClient) SetGrantTypes(v []OAuthGrantType)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *OpenIdConnectApplicationSettingsClient) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetIdpInitiatedLogin

`func (o *OpenIdConnectApplicationSettingsClient) GetIdpInitiatedLogin() OpenIdConnectApplicationIdpInitiatedLogin`

GetIdpInitiatedLogin returns the IdpInitiatedLogin field if non-nil, zero value otherwise.

### GetIdpInitiatedLoginOk

`func (o *OpenIdConnectApplicationSettingsClient) GetIdpInitiatedLoginOk() (*OpenIdConnectApplicationIdpInitiatedLogin, bool)`

GetIdpInitiatedLoginOk returns a tuple with the IdpInitiatedLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpInitiatedLogin

`func (o *OpenIdConnectApplicationSettingsClient) SetIdpInitiatedLogin(v OpenIdConnectApplicationIdpInitiatedLogin)`

SetIdpInitiatedLogin sets IdpInitiatedLogin field to given value.

### HasIdpInitiatedLogin

`func (o *OpenIdConnectApplicationSettingsClient) HasIdpInitiatedLogin() bool`

HasIdpInitiatedLogin returns a boolean if a field has been set.

### GetInitiateLoginUri

`func (o *OpenIdConnectApplicationSettingsClient) GetInitiateLoginUri() string`

GetInitiateLoginUri returns the InitiateLoginUri field if non-nil, zero value otherwise.

### GetInitiateLoginUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetInitiateLoginUriOk() (*string, bool)`

GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateLoginUri

`func (o *OpenIdConnectApplicationSettingsClient) SetInitiateLoginUri(v string)`

SetInitiateLoginUri sets InitiateLoginUri field to given value.

### HasInitiateLoginUri

`func (o *OpenIdConnectApplicationSettingsClient) HasInitiateLoginUri() bool`

HasInitiateLoginUri returns a boolean if a field has been set.

### GetIssuerMode

`func (o *OpenIdConnectApplicationSettingsClient) GetIssuerMode() OpenIdConnectApplicationIssuerMode`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *OpenIdConnectApplicationSettingsClient) GetIssuerModeOk() (*OpenIdConnectApplicationIssuerMode, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *OpenIdConnectApplicationSettingsClient) SetIssuerMode(v OpenIdConnectApplicationIssuerMode)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *OpenIdConnectApplicationSettingsClient) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetJwks

`func (o *OpenIdConnectApplicationSettingsClient) GetJwks() OpenIdConnectApplicationSettingsClientKeys`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *OpenIdConnectApplicationSettingsClient) GetJwksOk() (*OpenIdConnectApplicationSettingsClientKeys, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *OpenIdConnectApplicationSettingsClient) SetJwks(v OpenIdConnectApplicationSettingsClientKeys)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *OpenIdConnectApplicationSettingsClient) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetLogoUri

`func (o *OpenIdConnectApplicationSettingsClient) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *OpenIdConnectApplicationSettingsClient) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *OpenIdConnectApplicationSettingsClient) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.

### GetPolicyUri

`func (o *OpenIdConnectApplicationSettingsClient) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *OpenIdConnectApplicationSettingsClient) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *OpenIdConnectApplicationSettingsClient) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### GetPostLogoutRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) GetPostLogoutRedirectUris() []string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *OpenIdConnectApplicationSettingsClient) GetPostLogoutRedirectUrisOk() (*[]string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) SetPostLogoutRedirectUris(v []string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OpenIdConnectApplicationSettingsClient) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRefreshToken

`func (o *OpenIdConnectApplicationSettingsClient) GetRefreshToken() OpenIdConnectApplicationSettingsRefreshToken`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OpenIdConnectApplicationSettingsClient) GetRefreshTokenOk() (*OpenIdConnectApplicationSettingsRefreshToken, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OpenIdConnectApplicationSettingsClient) SetRefreshToken(v OpenIdConnectApplicationSettingsRefreshToken)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OpenIdConnectApplicationSettingsClient) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetResponseTypes

`func (o *OpenIdConnectApplicationSettingsClient) GetResponseTypes() []OAuthResponseType`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *OpenIdConnectApplicationSettingsClient) GetResponseTypesOk() (*[]OAuthResponseType, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *OpenIdConnectApplicationSettingsClient) SetResponseTypes(v []OAuthResponseType)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *OpenIdConnectApplicationSettingsClient) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetTosUri

`func (o *OpenIdConnectApplicationSettingsClient) GetTosUri() string`

GetTosUri returns the TosUri field if non-nil, zero value otherwise.

### GetTosUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetTosUriOk() (*string, bool)`

GetTosUriOk returns a tuple with the TosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUri

`func (o *OpenIdConnectApplicationSettingsClient) SetTosUri(v string)`

SetTosUri sets TosUri field to given value.

### HasTosUri

`func (o *OpenIdConnectApplicationSettingsClient) HasTosUri() bool`

HasTosUri returns a boolean if a field has been set.

### GetWildcardRedirect

`func (o *OpenIdConnectApplicationSettingsClient) GetWildcardRedirect() string`

GetWildcardRedirect returns the WildcardRedirect field if non-nil, zero value otherwise.

### GetWildcardRedirectOk

`func (o *OpenIdConnectApplicationSettingsClient) GetWildcardRedirectOk() (*string, bool)`

GetWildcardRedirectOk returns a tuple with the WildcardRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardRedirect

`func (o *OpenIdConnectApplicationSettingsClient) SetWildcardRedirect(v string)`

SetWildcardRedirect sets WildcardRedirect field to given value.

### HasWildcardRedirect

`func (o *OpenIdConnectApplicationSettingsClient) HasWildcardRedirect() bool`

HasWildcardRedirect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


