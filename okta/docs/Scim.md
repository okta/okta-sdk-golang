# Scim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | **string** | The authentication mode for requests to your SCIM server  | authMode | Description | | -------- | ----------- | | &#x60;header&#x60; | Uses authorization header with a customer-provided token value in the following format: &#x60;Authorization: {API token}&#x60; | | &#x60;bearer&#x60; | Uses authorization header with a customer-provided bearer token in the following format: &#x60;Authorization: Bearer {API token}&#x60; | | {authModeId} | The ID of the auth mode object that contains OAuth 2.0 credentials. &lt;br&gt; **Note:** Use the &#x60;/integrations/api/v1/internal/authModes&#x60; endpoint to create the auth mode object. | | 
**BaseUri** | **string** | The base URL that Okta uses to send outbound calls to your SCIM server. Only the HTTPS protocol is supported. You can use the app-level variables defined in the &#x60;config&#x60; array for the base URL. For example, if you have a &#x60;subdomain&#x60; variable defined in the &#x60;config&#x60; array and the URL to retrieve SCIM users for your integration is &#x60;https://${subdomain}.example.com/scim/v2/Users&#x60;, then specify the following base URL: &#x60;&#39;https://&#39; + app.subdomain + &#39;.example.com/scim/v2&#39;&#x60;. | 
**EntitlementTypes** | Pointer to [**[]EntitlementTypesInner**](EntitlementTypesInner.md) | List of supported entitlement types | [optional] 
**ScimServerConfig** | [**ScimScimServerConfig**](ScimScimServerConfig.md) |  | 
**SetupInstructionsUri** | **string** | The URL to your customer-facing instructions for configuring your SCIM integration. See [Customer configuration document guidelines](https://developer.okta.com/docs/guides/submit-app-prereq/main/#customer-configuration-document-guidelines). | 

## Methods

### NewScim

`func NewScim(authMode string, baseUri string, scimServerConfig ScimScimServerConfig, setupInstructionsUri string, ) *Scim`

NewScim instantiates a new Scim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimWithDefaults

`func NewScimWithDefaults() *Scim`

NewScimWithDefaults instantiates a new Scim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *Scim) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *Scim) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *Scim) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.


### GetBaseUri

`func (o *Scim) GetBaseUri() string`

GetBaseUri returns the BaseUri field if non-nil, zero value otherwise.

### GetBaseUriOk

`func (o *Scim) GetBaseUriOk() (*string, bool)`

GetBaseUriOk returns a tuple with the BaseUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUri

`func (o *Scim) SetBaseUri(v string)`

SetBaseUri sets BaseUri field to given value.


### GetEntitlementTypes

`func (o *Scim) GetEntitlementTypes() []EntitlementTypesInner`

GetEntitlementTypes returns the EntitlementTypes field if non-nil, zero value otherwise.

### GetEntitlementTypesOk

`func (o *Scim) GetEntitlementTypesOk() (*[]EntitlementTypesInner, bool)`

GetEntitlementTypesOk returns a tuple with the EntitlementTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementTypes

`func (o *Scim) SetEntitlementTypes(v []EntitlementTypesInner)`

SetEntitlementTypes sets EntitlementTypes field to given value.

### HasEntitlementTypes

`func (o *Scim) HasEntitlementTypes() bool`

HasEntitlementTypes returns a boolean if a field has been set.

### GetScimServerConfig

`func (o *Scim) GetScimServerConfig() ScimScimServerConfig`

GetScimServerConfig returns the ScimServerConfig field if non-nil, zero value otherwise.

### GetScimServerConfigOk

`func (o *Scim) GetScimServerConfigOk() (*ScimScimServerConfig, bool)`

GetScimServerConfigOk returns a tuple with the ScimServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimServerConfig

`func (o *Scim) SetScimServerConfig(v ScimScimServerConfig)`

SetScimServerConfig sets ScimServerConfig field to given value.


### GetSetupInstructionsUri

`func (o *Scim) GetSetupInstructionsUri() string`

GetSetupInstructionsUri returns the SetupInstructionsUri field if non-nil, zero value otherwise.

### GetSetupInstructionsUriOk

`func (o *Scim) GetSetupInstructionsUriOk() (*string, bool)`

GetSetupInstructionsUriOk returns a tuple with the SetupInstructionsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupInstructionsUri

`func (o *Scim) SetSetupInstructionsUri(v string)`

SetSetupInstructionsUri sets SetupInstructionsUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


