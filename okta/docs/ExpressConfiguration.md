# ExpressConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationClientId** | **string** | The client ID of the SaaS application that end users sign in to | 
**Capabilities** | Pointer to **[]string** | The Express Configuration capabilities to enable on the submission. If omitted, all capabilities are automatically configured based on the OIN integration&#39;s supported protocols. | [optional] 
**InitiateLoginUriTemplate** | Pointer to **string** | The URL template that Okta uses to launch the app from the end-user dashboard. Supports template variables &#x60;{organization_name}&#x60;, &#x60;{organization_id}&#x60;, and &#x60;{connection_name}&#x60;. | [optional] 
**LoginDomain** | **string** | The Auth0 admin login domain that Okta redirects to as part of the consent flow in a web browser. Use the custom domain name if one is configured in Auth0. | 
**OinClientId** | **string** | The client ID of the additional client application that Auth0 creates for the OIN administrator consent flow between Okta and Auth0 | 
**TenantDomain** | **string** | The Auth0 tenant domain (for example, &#x60;example.auth0.com&#x60;) | 

## Methods

### NewExpressConfiguration

`func NewExpressConfiguration(applicationClientId string, loginDomain string, oinClientId string, tenantDomain string, ) *ExpressConfiguration`

NewExpressConfiguration instantiates a new ExpressConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpressConfigurationWithDefaults

`func NewExpressConfigurationWithDefaults() *ExpressConfiguration`

NewExpressConfigurationWithDefaults instantiates a new ExpressConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationClientId

`func (o *ExpressConfiguration) GetApplicationClientId() string`

GetApplicationClientId returns the ApplicationClientId field if non-nil, zero value otherwise.

### GetApplicationClientIdOk

`func (o *ExpressConfiguration) GetApplicationClientIdOk() (*string, bool)`

GetApplicationClientIdOk returns a tuple with the ApplicationClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationClientId

`func (o *ExpressConfiguration) SetApplicationClientId(v string)`

SetApplicationClientId sets ApplicationClientId field to given value.


### GetCapabilities

`func (o *ExpressConfiguration) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ExpressConfiguration) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ExpressConfiguration) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ExpressConfiguration) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetInitiateLoginUriTemplate

`func (o *ExpressConfiguration) GetInitiateLoginUriTemplate() string`

GetInitiateLoginUriTemplate returns the InitiateLoginUriTemplate field if non-nil, zero value otherwise.

### GetInitiateLoginUriTemplateOk

`func (o *ExpressConfiguration) GetInitiateLoginUriTemplateOk() (*string, bool)`

GetInitiateLoginUriTemplateOk returns a tuple with the InitiateLoginUriTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateLoginUriTemplate

`func (o *ExpressConfiguration) SetInitiateLoginUriTemplate(v string)`

SetInitiateLoginUriTemplate sets InitiateLoginUriTemplate field to given value.

### HasInitiateLoginUriTemplate

`func (o *ExpressConfiguration) HasInitiateLoginUriTemplate() bool`

HasInitiateLoginUriTemplate returns a boolean if a field has been set.

### GetLoginDomain

`func (o *ExpressConfiguration) GetLoginDomain() string`

GetLoginDomain returns the LoginDomain field if non-nil, zero value otherwise.

### GetLoginDomainOk

`func (o *ExpressConfiguration) GetLoginDomainOk() (*string, bool)`

GetLoginDomainOk returns a tuple with the LoginDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDomain

`func (o *ExpressConfiguration) SetLoginDomain(v string)`

SetLoginDomain sets LoginDomain field to given value.


### GetOinClientId

`func (o *ExpressConfiguration) GetOinClientId() string`

GetOinClientId returns the OinClientId field if non-nil, zero value otherwise.

### GetOinClientIdOk

`func (o *ExpressConfiguration) GetOinClientIdOk() (*string, bool)`

GetOinClientIdOk returns a tuple with the OinClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOinClientId

`func (o *ExpressConfiguration) SetOinClientId(v string)`

SetOinClientId sets OinClientId field to given value.


### GetTenantDomain

`func (o *ExpressConfiguration) GetTenantDomain() string`

GetTenantDomain returns the TenantDomain field if non-nil, zero value otherwise.

### GetTenantDomainOk

`func (o *ExpressConfiguration) GetTenantDomainOk() (*string, bool)`

GetTenantDomainOk returns a tuple with the TenantDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantDomain

`func (o *ExpressConfiguration) SetTenantDomain(v string)`

SetTenantDomain sets TenantDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


