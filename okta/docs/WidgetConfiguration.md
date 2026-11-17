# WidgetConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandColors** | Pointer to [**WidgetConfigurationBrandColors**](WidgetConfigurationBrandColors.md) |  | [optional] 
**BrandName** | Pointer to **string** | Brand name displayed on the widget | [optional] 
**Features** | Pointer to [**WidgetConfigurationFeatures**](WidgetConfigurationFeatures.md) |  | [optional] 
**HelpLinks** | Pointer to [**WidgetConfigurationHelpLinks**](WidgetConfigurationHelpLinks.md) |  | [optional] 
**HelpSupportNumber** | Pointer to **string** | Support phone number displayed on the widget | [optional] 
**I18n** | Pointer to **map[string]map[string]string** | Internationalization overrides. Keys are locale codes (e.g. \&quot;en\&quot;, \&quot;ja\&quot;, \&quot;fr\&quot;), values are objects mapping i18n keys to string overrides. Override values can&#39;t contain &#x60;&lt;&#x60; or &#x60;&gt;&#x60; characters. | [optional] 
**Language** | Pointer to **string** | Language code for the widget | [optional] 
**Logo** | Pointer to **string** | URL for the logo displayed on the Sign-In Widget | [optional] 
**LogoText** | Pointer to **string** | Alt text for the logo | [optional] 
**SignOutLink** | Pointer to **string** | Custom sign-out URL | [optional] 
**Theme** | Pointer to [**WidgetConfigurationTheme**](WidgetConfigurationTheme.md) |  | [optional] 

## Methods

### NewWidgetConfiguration

`func NewWidgetConfiguration() *WidgetConfiguration`

NewWidgetConfiguration instantiates a new WidgetConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetConfigurationWithDefaults

`func NewWidgetConfigurationWithDefaults() *WidgetConfiguration`

NewWidgetConfigurationWithDefaults instantiates a new WidgetConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandColors

`func (o *WidgetConfiguration) GetBrandColors() WidgetConfigurationBrandColors`

GetBrandColors returns the BrandColors field if non-nil, zero value otherwise.

### GetBrandColorsOk

`func (o *WidgetConfiguration) GetBrandColorsOk() (*WidgetConfigurationBrandColors, bool)`

GetBrandColorsOk returns a tuple with the BrandColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandColors

`func (o *WidgetConfiguration) SetBrandColors(v WidgetConfigurationBrandColors)`

SetBrandColors sets BrandColors field to given value.

### HasBrandColors

`func (o *WidgetConfiguration) HasBrandColors() bool`

HasBrandColors returns a boolean if a field has been set.

### GetBrandName

`func (o *WidgetConfiguration) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *WidgetConfiguration) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *WidgetConfiguration) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *WidgetConfiguration) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### GetFeatures

`func (o *WidgetConfiguration) GetFeatures() WidgetConfigurationFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *WidgetConfiguration) GetFeaturesOk() (*WidgetConfigurationFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *WidgetConfiguration) SetFeatures(v WidgetConfigurationFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *WidgetConfiguration) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHelpLinks

`func (o *WidgetConfiguration) GetHelpLinks() WidgetConfigurationHelpLinks`

GetHelpLinks returns the HelpLinks field if non-nil, zero value otherwise.

### GetHelpLinksOk

`func (o *WidgetConfiguration) GetHelpLinksOk() (*WidgetConfigurationHelpLinks, bool)`

GetHelpLinksOk returns a tuple with the HelpLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpLinks

`func (o *WidgetConfiguration) SetHelpLinks(v WidgetConfigurationHelpLinks)`

SetHelpLinks sets HelpLinks field to given value.

### HasHelpLinks

`func (o *WidgetConfiguration) HasHelpLinks() bool`

HasHelpLinks returns a boolean if a field has been set.

### GetHelpSupportNumber

`func (o *WidgetConfiguration) GetHelpSupportNumber() string`

GetHelpSupportNumber returns the HelpSupportNumber field if non-nil, zero value otherwise.

### GetHelpSupportNumberOk

`func (o *WidgetConfiguration) GetHelpSupportNumberOk() (*string, bool)`

GetHelpSupportNumberOk returns a tuple with the HelpSupportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpSupportNumber

`func (o *WidgetConfiguration) SetHelpSupportNumber(v string)`

SetHelpSupportNumber sets HelpSupportNumber field to given value.

### HasHelpSupportNumber

`func (o *WidgetConfiguration) HasHelpSupportNumber() bool`

HasHelpSupportNumber returns a boolean if a field has been set.

### GetI18n

`func (o *WidgetConfiguration) GetI18n() map[string]map[string]string`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *WidgetConfiguration) GetI18nOk() (*map[string]map[string]string, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *WidgetConfiguration) SetI18n(v map[string]map[string]string)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *WidgetConfiguration) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetLanguage

`func (o *WidgetConfiguration) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WidgetConfiguration) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WidgetConfiguration) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *WidgetConfiguration) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLogo

`func (o *WidgetConfiguration) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *WidgetConfiguration) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *WidgetConfiguration) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *WidgetConfiguration) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoText

`func (o *WidgetConfiguration) GetLogoText() string`

GetLogoText returns the LogoText field if non-nil, zero value otherwise.

### GetLogoTextOk

`func (o *WidgetConfiguration) GetLogoTextOk() (*string, bool)`

GetLogoTextOk returns a tuple with the LogoText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoText

`func (o *WidgetConfiguration) SetLogoText(v string)`

SetLogoText sets LogoText field to given value.

### HasLogoText

`func (o *WidgetConfiguration) HasLogoText() bool`

HasLogoText returns a boolean if a field has been set.

### GetSignOutLink

`func (o *WidgetConfiguration) GetSignOutLink() string`

GetSignOutLink returns the SignOutLink field if non-nil, zero value otherwise.

### GetSignOutLinkOk

`func (o *WidgetConfiguration) GetSignOutLinkOk() (*string, bool)`

GetSignOutLinkOk returns a tuple with the SignOutLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOutLink

`func (o *WidgetConfiguration) SetSignOutLink(v string)`

SetSignOutLink sets SignOutLink field to given value.

### HasSignOutLink

`func (o *WidgetConfiguration) HasSignOutLink() bool`

HasSignOutLink returns a boolean if a field has been set.

### GetTheme

`func (o *WidgetConfiguration) GetTheme() WidgetConfigurationTheme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *WidgetConfiguration) GetThemeOk() (*WidgetConfigurationTheme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *WidgetConfiguration) SetTheme(v WidgetConfigurationTheme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *WidgetConfiguration) HasTheme() bool`

HasTheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


