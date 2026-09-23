# PublicAppInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to [**ApplicationAccessibility**](ApplicationAccessibility.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the application object was created | [optional] [readonly] 
**ExpressConfiguration** | Pointer to [**ApplicationExpressConfiguration**](ApplicationExpressConfiguration.md) |  | [optional] 
**Features** | Pointer to **[]string** | Enabled app features &gt; **Note:** See [Application Features](/openapi/okta-management/management/tags/applicationfeatures/) for app provisioning features.  | [optional] [readonly] 
**Id** | Pointer to **string** | Unique ID for the app instance | [optional] [readonly] 
**Label** | **string** | User-defined display name for app | 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the application object was last updated | [optional] [readonly] 
**Licensing** | Pointer to [**ApplicationLicensing**](ApplicationLicensing.md) |  | [optional] 
**Orn** | Pointer to **string** | The Okta resource name (ORN) for the current app instance | [optional] [readonly] 
**Profile** | Pointer to **map[string]interface{}** | Contains any valid JSON schema for specifying properties that can be referenced from a request (only available to OAuth 2.0 client apps). For example, add an app manager contact email address or define an allowlist of groups that you can then reference using the Okta Expression Language &#x60;getFilteredGroups&#x60; function.  &gt; **Notes:** &gt; * &#x60;profile&#x60; isn&#39;t encrypted, so don&#39;t store sensitive data in it. &gt; * &#x60;profile&#x60; doesn&#39;t limit the level of nesting in the JSON schema you created, but there is a practical size limit. Okta recommends a JSON schema size of 1 MB or less for best performance. | [optional] 
**SignOnMode** | **string** | Authentication mode for the app  | signOnMode | Description | | ---------- | ----------- | | AUTO_LOGIN | Secure Web Authentication (SWA) | | BASIC_AUTH | HTTP Basic Authentication with Okta Browser Plugin | | BOOKMARK | Just a bookmark (no-authentication) | | BROWSER_PLUGIN | Secure Web Authentication (SWA) with Okta Browser Plugin | | OPENID_CONNECT | Federated Authentication with OpenID Connect (OIDC) | | SAML_1_1 | Federated Authentication with SAML 1.1 WebSSO (not supported for custom apps) | | SAML_2_0 | Federated Authentication with SAML 2.0 WebSSO | | SECURE_PASSWORD_STORE | Secure Web Authentication (SWA) with POST (plugin not required) | | WS_FEDERATION | Federated Authentication with WS-Federation Passive Requestor Profile |  Select the &#x60;signOnMode&#x60; for your custom app:  | 
**Status** | Pointer to **string** | App instance status | [optional] [readonly] 
**Types** | Pointer to **[]string** | App type categorization tags used by the Apps list page facets. Values align with the [&#x60;/apps/types&#x60;](/openapi/okta-management/management/tag/Application/#tag/Application/operation/listAppTypeUsage) taxonomy (for example, &#x60;SAML&#x60;, &#x60;OIDC&#x60;, &#x60;PROVISIONING&#x60;, &#x60;UNIVERSAL_LOGOUT&#x60;). Multi-valued because an app can belong to several buckets.  | [optional] [readonly] 
**UniversalLogout** | Pointer to [**ApplicationUniversalLogout**](ApplicationUniversalLogout.md) |  | [optional] 
**Visibility** | Pointer to [**ApplicationVisibility**](ApplicationVisibility.md) |  | [optional] 
**Embedded** | Pointer to [**ApplicationEmbedded**](ApplicationEmbedded.md) |  | [optional] 
**Links** | Pointer to [**ApplicationLinks**](ApplicationLinks.md) |  | [optional] 

## Methods

### NewPublicAppInstance

`func NewPublicAppInstance(label string, signOnMode string, ) *PublicAppInstance`

NewPublicAppInstance instantiates a new PublicAppInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAppInstanceWithDefaults

`func NewPublicAppInstanceWithDefaults() *PublicAppInstance`

NewPublicAppInstanceWithDefaults instantiates a new PublicAppInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *PublicAppInstance) GetAccessibility() ApplicationAccessibility`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *PublicAppInstance) GetAccessibilityOk() (*ApplicationAccessibility, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *PublicAppInstance) SetAccessibility(v ApplicationAccessibility)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *PublicAppInstance) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetCreated

`func (o *PublicAppInstance) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PublicAppInstance) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PublicAppInstance) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PublicAppInstance) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpressConfiguration

`func (o *PublicAppInstance) GetExpressConfiguration() ApplicationExpressConfiguration`

GetExpressConfiguration returns the ExpressConfiguration field if non-nil, zero value otherwise.

### GetExpressConfigurationOk

`func (o *PublicAppInstance) GetExpressConfigurationOk() (*ApplicationExpressConfiguration, bool)`

GetExpressConfigurationOk returns a tuple with the ExpressConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressConfiguration

`func (o *PublicAppInstance) SetExpressConfiguration(v ApplicationExpressConfiguration)`

SetExpressConfiguration sets ExpressConfiguration field to given value.

### HasExpressConfiguration

`func (o *PublicAppInstance) HasExpressConfiguration() bool`

HasExpressConfiguration returns a boolean if a field has been set.

### GetFeatures

`func (o *PublicAppInstance) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PublicAppInstance) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PublicAppInstance) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PublicAppInstance) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *PublicAppInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicAppInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicAppInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicAppInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PublicAppInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PublicAppInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PublicAppInstance) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLastUpdated

`func (o *PublicAppInstance) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PublicAppInstance) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PublicAppInstance) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PublicAppInstance) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLicensing

`func (o *PublicAppInstance) GetLicensing() ApplicationLicensing`

GetLicensing returns the Licensing field if non-nil, zero value otherwise.

### GetLicensingOk

`func (o *PublicAppInstance) GetLicensingOk() (*ApplicationLicensing, bool)`

GetLicensingOk returns a tuple with the Licensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensing

`func (o *PublicAppInstance) SetLicensing(v ApplicationLicensing)`

SetLicensing sets Licensing field to given value.

### HasLicensing

`func (o *PublicAppInstance) HasLicensing() bool`

HasLicensing returns a boolean if a field has been set.

### GetOrn

`func (o *PublicAppInstance) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *PublicAppInstance) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *PublicAppInstance) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *PublicAppInstance) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetProfile

`func (o *PublicAppInstance) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PublicAppInstance) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PublicAppInstance) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PublicAppInstance) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSignOnMode

`func (o *PublicAppInstance) GetSignOnMode() string`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *PublicAppInstance) GetSignOnModeOk() (*string, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *PublicAppInstance) SetSignOnMode(v string)`

SetSignOnMode sets SignOnMode field to given value.


### GetStatus

`func (o *PublicAppInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicAppInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicAppInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicAppInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTypes

`func (o *PublicAppInstance) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *PublicAppInstance) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *PublicAppInstance) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *PublicAppInstance) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUniversalLogout

`func (o *PublicAppInstance) GetUniversalLogout() ApplicationUniversalLogout`

GetUniversalLogout returns the UniversalLogout field if non-nil, zero value otherwise.

### GetUniversalLogoutOk

`func (o *PublicAppInstance) GetUniversalLogoutOk() (*ApplicationUniversalLogout, bool)`

GetUniversalLogoutOk returns a tuple with the UniversalLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalLogout

`func (o *PublicAppInstance) SetUniversalLogout(v ApplicationUniversalLogout)`

SetUniversalLogout sets UniversalLogout field to given value.

### HasUniversalLogout

`func (o *PublicAppInstance) HasUniversalLogout() bool`

HasUniversalLogout returns a boolean if a field has been set.

### GetVisibility

`func (o *PublicAppInstance) GetVisibility() ApplicationVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PublicAppInstance) GetVisibilityOk() (*ApplicationVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PublicAppInstance) SetVisibility(v ApplicationVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PublicAppInstance) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEmbedded

`func (o *PublicAppInstance) GetEmbedded() ApplicationEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *PublicAppInstance) GetEmbeddedOk() (*ApplicationEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *PublicAppInstance) SetEmbedded(v ApplicationEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *PublicAppInstance) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *PublicAppInstance) GetLinks() ApplicationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PublicAppInstance) GetLinksOk() (*ApplicationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PublicAppInstance) SetLinks(v ApplicationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PublicAppInstance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


