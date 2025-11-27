# SignInPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageContent** | Pointer to **string** | The HTML for the page | [optional] 
**ContentSecurityPolicySetting** | Pointer to [**ContentSecurityPolicySetting**](ContentSecurityPolicySetting.md) |  | [optional] 
**WidgetCustomizations** | Pointer to [**SignInPageAllOfWidgetCustomizations**](SignInPageAllOfWidgetCustomizations.md) |  | [optional] 
**WidgetVersion** | Pointer to **string** | The version specified as a [Semantic Version](https://semver.org/). This value can be a wildcard (&#x60;*&#x60;), a major version range (for example, &#x60;^2&#x60;), a major-only version (for example, &#x60;7&#x60;), or a specific &#x60;Major.Minor&#x60; version (for example, &#x60;5.15&#x60;). | [optional] 

## Methods

### NewSignInPage

`func NewSignInPage() *SignInPage`

NewSignInPage instantiates a new SignInPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignInPageWithDefaults

`func NewSignInPageWithDefaults() *SignInPage`

NewSignInPageWithDefaults instantiates a new SignInPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageContent

`func (o *SignInPage) GetPageContent() string`

GetPageContent returns the PageContent field if non-nil, zero value otherwise.

### GetPageContentOk

`func (o *SignInPage) GetPageContentOk() (*string, bool)`

GetPageContentOk returns a tuple with the PageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageContent

`func (o *SignInPage) SetPageContent(v string)`

SetPageContent sets PageContent field to given value.

### HasPageContent

`func (o *SignInPage) HasPageContent() bool`

HasPageContent returns a boolean if a field has been set.

### GetContentSecurityPolicySetting

`func (o *SignInPage) GetContentSecurityPolicySetting() ContentSecurityPolicySetting`

GetContentSecurityPolicySetting returns the ContentSecurityPolicySetting field if non-nil, zero value otherwise.

### GetContentSecurityPolicySettingOk

`func (o *SignInPage) GetContentSecurityPolicySettingOk() (*ContentSecurityPolicySetting, bool)`

GetContentSecurityPolicySettingOk returns a tuple with the ContentSecurityPolicySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSecurityPolicySetting

`func (o *SignInPage) SetContentSecurityPolicySetting(v ContentSecurityPolicySetting)`

SetContentSecurityPolicySetting sets ContentSecurityPolicySetting field to given value.

### HasContentSecurityPolicySetting

`func (o *SignInPage) HasContentSecurityPolicySetting() bool`

HasContentSecurityPolicySetting returns a boolean if a field has been set.

### GetWidgetCustomizations

`func (o *SignInPage) GetWidgetCustomizations() SignInPageAllOfWidgetCustomizations`

GetWidgetCustomizations returns the WidgetCustomizations field if non-nil, zero value otherwise.

### GetWidgetCustomizationsOk

`func (o *SignInPage) GetWidgetCustomizationsOk() (*SignInPageAllOfWidgetCustomizations, bool)`

GetWidgetCustomizationsOk returns a tuple with the WidgetCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetCustomizations

`func (o *SignInPage) SetWidgetCustomizations(v SignInPageAllOfWidgetCustomizations)`

SetWidgetCustomizations sets WidgetCustomizations field to given value.

### HasWidgetCustomizations

`func (o *SignInPage) HasWidgetCustomizations() bool`

HasWidgetCustomizations returns a boolean if a field has been set.

### GetWidgetVersion

`func (o *SignInPage) GetWidgetVersion() string`

GetWidgetVersion returns the WidgetVersion field if non-nil, zero value otherwise.

### GetWidgetVersionOk

`func (o *SignInPage) GetWidgetVersionOk() (*string, bool)`

GetWidgetVersionOk returns a tuple with the WidgetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetVersion

`func (o *SignInPage) SetWidgetVersion(v string)`

SetWidgetVersion sets WidgetVersion field to given value.

### HasWidgetVersion

`func (o *SignInPage) HasWidgetVersion() bool`

HasWidgetVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


