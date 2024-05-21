# AuthorizationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | Pointer to **[]string** | The recipients that the tokens are intended for. This becomes the &#x60;aud&#x60; claim in an access token. Okta currently supports only one audience. | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Credentials** | Pointer to [**AuthorizationServerCredentials**](AuthorizationServerCredentials.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the custom authorization server | [optional] 
**Id** | Pointer to **string** | The ID of the custom authorization server | [optional] [readonly] 
**Issuer** | Pointer to **string** | The complete URL for the custom authorization server. This becomes the &#x60;iss&#x60; claim in an access token. | [optional] 
**IssuerMode** | Pointer to **string** | Indicates which value is specified in the issuer of the tokens that a custom authorization server returns: the Okta org domain URL or a custom domain URL.  &#x60;issuerMode&#x60; is visible if you have a custom URL domain configured or the Dynamic Issuer Mode feature enabled. If you have a custom URL domain configured, you can set a custom domain URL in a custom authorization server, and this property is returned in the appropriate responses.  When set to &#x60;ORG_URL&#x60;, then in responses, &#x60;issuer&#x60; is the Okta org domain URL: &#x60;https://${yourOktaDomain}&#x60;.  When set to &#x60;CUSTOM_URL&#x60;, then in responses, &#x60;issuer&#x60; is the custom domain URL configured in the administration user interface.  When set to &#x60;DYNAMIC&#x60;, then in responses, &#x60;issuer&#x60; is the custom domain URL if the OAuth 2.0 request was sent to the custom domain, or is the Okta org&#39;s domain URL if the OAuth 2.0 request was sent to the original Okta org domain.  After you configure a custom URL domain, all new custom authorization servers use &#x60;CUSTOM_URL&#x60; by default. If the Dynamic Issuer Mode feature is enabled, then all new custom authorization servers use &#x60;DYNAMIC&#x60; by default. All existing custom authorization servers continue to use the original value until they&#39;re changed using the Admin Console or the API. This way, existing integrations with the client and resource server continue to work after the feature is enabled. | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the custom authorization server | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**AuthServerLinks**](AuthServerLinks.md) |  | [optional] 

## Methods

### NewAuthorizationServer

`func NewAuthorizationServer() *AuthorizationServer`

NewAuthorizationServer instantiates a new AuthorizationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerWithDefaults

`func NewAuthorizationServerWithDefaults() *AuthorizationServer`

NewAuthorizationServerWithDefaults instantiates a new AuthorizationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *AuthorizationServer) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *AuthorizationServer) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *AuthorizationServer) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *AuthorizationServer) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### GetCreated

`func (o *AuthorizationServer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthorizationServer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthorizationServer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthorizationServer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentials

`func (o *AuthorizationServer) GetCredentials() AuthorizationServerCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AuthorizationServer) GetCredentialsOk() (*AuthorizationServerCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AuthorizationServer) SetCredentials(v AuthorizationServerCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AuthorizationServer) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizationServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizationServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizationServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizationServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AuthorizationServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizationServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizationServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizationServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuer

`func (o *AuthorizationServer) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AuthorizationServer) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AuthorizationServer) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *AuthorizationServer) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetIssuerMode

`func (o *AuthorizationServer) GetIssuerMode() string`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *AuthorizationServer) GetIssuerModeOk() (*string, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *AuthorizationServer) SetIssuerMode(v string)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *AuthorizationServer) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AuthorizationServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AuthorizationServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AuthorizationServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AuthorizationServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *AuthorizationServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizationServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizationServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizationServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorizationServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorizationServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorizationServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorizationServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *AuthorizationServer) GetLinks() AuthServerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizationServer) GetLinksOk() (*AuthServerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizationServer) SetLinks(v AuthServerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizationServer) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


