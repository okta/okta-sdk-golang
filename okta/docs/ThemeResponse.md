# ThemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundImage** | Pointer to **string** |  | [optional] [readonly] 
**EmailTemplateTouchPointVariant** | Pointer to [**EmailTemplateTouchPointVariant**](EmailTemplateTouchPointVariant.md) |  | [optional] 
**EndUserDashboardTouchPointVariant** | Pointer to [**EndUserDashboardTouchPointVariant**](EndUserDashboardTouchPointVariant.md) |  | [optional] 
**ErrorPageTouchPointVariant** | Pointer to [**ErrorPageTouchPointVariant**](ErrorPageTouchPointVariant.md) |  | [optional] 
**Favicon** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LoadingPageTouchPointVariant** | Pointer to [**LoadingPageTouchPointVariant**](LoadingPageTouchPointVariant.md) |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] [readonly] 
**PrimaryColorContrastHex** | Pointer to **string** |  | [optional] 
**PrimaryColorHex** | Pointer to **string** |  | [optional] 
**SecondaryColorContrastHex** | Pointer to **string** |  | [optional] 
**SecondaryColorHex** | Pointer to **string** |  | [optional] 
**SignInPageTouchPointVariant** | Pointer to [**SignInPageTouchPointVariant**](SignInPageTouchPointVariant.md) |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewThemeResponse

`func NewThemeResponse() *ThemeResponse`

NewThemeResponse instantiates a new ThemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThemeResponseWithDefaults

`func NewThemeResponseWithDefaults() *ThemeResponse`

NewThemeResponseWithDefaults instantiates a new ThemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundImage

`func (o *ThemeResponse) GetBackgroundImage() string`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *ThemeResponse) GetBackgroundImageOk() (*string, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *ThemeResponse) SetBackgroundImage(v string)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *ThemeResponse) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### GetEmailTemplateTouchPointVariant

`func (o *ThemeResponse) GetEmailTemplateTouchPointVariant() EmailTemplateTouchPointVariant`

GetEmailTemplateTouchPointVariant returns the EmailTemplateTouchPointVariant field if non-nil, zero value otherwise.

### GetEmailTemplateTouchPointVariantOk

`func (o *ThemeResponse) GetEmailTemplateTouchPointVariantOk() (*EmailTemplateTouchPointVariant, bool)`

GetEmailTemplateTouchPointVariantOk returns a tuple with the EmailTemplateTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateTouchPointVariant

`func (o *ThemeResponse) SetEmailTemplateTouchPointVariant(v EmailTemplateTouchPointVariant)`

SetEmailTemplateTouchPointVariant sets EmailTemplateTouchPointVariant field to given value.

### HasEmailTemplateTouchPointVariant

`func (o *ThemeResponse) HasEmailTemplateTouchPointVariant() bool`

HasEmailTemplateTouchPointVariant returns a boolean if a field has been set.

### GetEndUserDashboardTouchPointVariant

`func (o *ThemeResponse) GetEndUserDashboardTouchPointVariant() EndUserDashboardTouchPointVariant`

GetEndUserDashboardTouchPointVariant returns the EndUserDashboardTouchPointVariant field if non-nil, zero value otherwise.

### GetEndUserDashboardTouchPointVariantOk

`func (o *ThemeResponse) GetEndUserDashboardTouchPointVariantOk() (*EndUserDashboardTouchPointVariant, bool)`

GetEndUserDashboardTouchPointVariantOk returns a tuple with the EndUserDashboardTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserDashboardTouchPointVariant

`func (o *ThemeResponse) SetEndUserDashboardTouchPointVariant(v EndUserDashboardTouchPointVariant)`

SetEndUserDashboardTouchPointVariant sets EndUserDashboardTouchPointVariant field to given value.

### HasEndUserDashboardTouchPointVariant

`func (o *ThemeResponse) HasEndUserDashboardTouchPointVariant() bool`

HasEndUserDashboardTouchPointVariant returns a boolean if a field has been set.

### GetErrorPageTouchPointVariant

`func (o *ThemeResponse) GetErrorPageTouchPointVariant() ErrorPageTouchPointVariant`

GetErrorPageTouchPointVariant returns the ErrorPageTouchPointVariant field if non-nil, zero value otherwise.

### GetErrorPageTouchPointVariantOk

`func (o *ThemeResponse) GetErrorPageTouchPointVariantOk() (*ErrorPageTouchPointVariant, bool)`

GetErrorPageTouchPointVariantOk returns a tuple with the ErrorPageTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPageTouchPointVariant

`func (o *ThemeResponse) SetErrorPageTouchPointVariant(v ErrorPageTouchPointVariant)`

SetErrorPageTouchPointVariant sets ErrorPageTouchPointVariant field to given value.

### HasErrorPageTouchPointVariant

`func (o *ThemeResponse) HasErrorPageTouchPointVariant() bool`

HasErrorPageTouchPointVariant returns a boolean if a field has been set.

### GetFavicon

`func (o *ThemeResponse) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *ThemeResponse) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *ThemeResponse) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *ThemeResponse) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetId

`func (o *ThemeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThemeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThemeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThemeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadingPageTouchPointVariant

`func (o *ThemeResponse) GetLoadingPageTouchPointVariant() LoadingPageTouchPointVariant`

GetLoadingPageTouchPointVariant returns the LoadingPageTouchPointVariant field if non-nil, zero value otherwise.

### GetLoadingPageTouchPointVariantOk

`func (o *ThemeResponse) GetLoadingPageTouchPointVariantOk() (*LoadingPageTouchPointVariant, bool)`

GetLoadingPageTouchPointVariantOk returns a tuple with the LoadingPageTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingPageTouchPointVariant

`func (o *ThemeResponse) SetLoadingPageTouchPointVariant(v LoadingPageTouchPointVariant)`

SetLoadingPageTouchPointVariant sets LoadingPageTouchPointVariant field to given value.

### HasLoadingPageTouchPointVariant

`func (o *ThemeResponse) HasLoadingPageTouchPointVariant() bool`

HasLoadingPageTouchPointVariant returns a boolean if a field has been set.

### GetLogo

`func (o *ThemeResponse) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ThemeResponse) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ThemeResponse) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ThemeResponse) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPrimaryColorContrastHex

`func (o *ThemeResponse) GetPrimaryColorContrastHex() string`

GetPrimaryColorContrastHex returns the PrimaryColorContrastHex field if non-nil, zero value otherwise.

### GetPrimaryColorContrastHexOk

`func (o *ThemeResponse) GetPrimaryColorContrastHexOk() (*string, bool)`

GetPrimaryColorContrastHexOk returns a tuple with the PrimaryColorContrastHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColorContrastHex

`func (o *ThemeResponse) SetPrimaryColorContrastHex(v string)`

SetPrimaryColorContrastHex sets PrimaryColorContrastHex field to given value.

### HasPrimaryColorContrastHex

`func (o *ThemeResponse) HasPrimaryColorContrastHex() bool`

HasPrimaryColorContrastHex returns a boolean if a field has been set.

### GetPrimaryColorHex

`func (o *ThemeResponse) GetPrimaryColorHex() string`

GetPrimaryColorHex returns the PrimaryColorHex field if non-nil, zero value otherwise.

### GetPrimaryColorHexOk

`func (o *ThemeResponse) GetPrimaryColorHexOk() (*string, bool)`

GetPrimaryColorHexOk returns a tuple with the PrimaryColorHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColorHex

`func (o *ThemeResponse) SetPrimaryColorHex(v string)`

SetPrimaryColorHex sets PrimaryColorHex field to given value.

### HasPrimaryColorHex

`func (o *ThemeResponse) HasPrimaryColorHex() bool`

HasPrimaryColorHex returns a boolean if a field has been set.

### GetSecondaryColorContrastHex

`func (o *ThemeResponse) GetSecondaryColorContrastHex() string`

GetSecondaryColorContrastHex returns the SecondaryColorContrastHex field if non-nil, zero value otherwise.

### GetSecondaryColorContrastHexOk

`func (o *ThemeResponse) GetSecondaryColorContrastHexOk() (*string, bool)`

GetSecondaryColorContrastHexOk returns a tuple with the SecondaryColorContrastHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryColorContrastHex

`func (o *ThemeResponse) SetSecondaryColorContrastHex(v string)`

SetSecondaryColorContrastHex sets SecondaryColorContrastHex field to given value.

### HasSecondaryColorContrastHex

`func (o *ThemeResponse) HasSecondaryColorContrastHex() bool`

HasSecondaryColorContrastHex returns a boolean if a field has been set.

### GetSecondaryColorHex

`func (o *ThemeResponse) GetSecondaryColorHex() string`

GetSecondaryColorHex returns the SecondaryColorHex field if non-nil, zero value otherwise.

### GetSecondaryColorHexOk

`func (o *ThemeResponse) GetSecondaryColorHexOk() (*string, bool)`

GetSecondaryColorHexOk returns a tuple with the SecondaryColorHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryColorHex

`func (o *ThemeResponse) SetSecondaryColorHex(v string)`

SetSecondaryColorHex sets SecondaryColorHex field to given value.

### HasSecondaryColorHex

`func (o *ThemeResponse) HasSecondaryColorHex() bool`

HasSecondaryColorHex returns a boolean if a field has been set.

### GetSignInPageTouchPointVariant

`func (o *ThemeResponse) GetSignInPageTouchPointVariant() SignInPageTouchPointVariant`

GetSignInPageTouchPointVariant returns the SignInPageTouchPointVariant field if non-nil, zero value otherwise.

### GetSignInPageTouchPointVariantOk

`func (o *ThemeResponse) GetSignInPageTouchPointVariantOk() (*SignInPageTouchPointVariant, bool)`

GetSignInPageTouchPointVariantOk returns a tuple with the SignInPageTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInPageTouchPointVariant

`func (o *ThemeResponse) SetSignInPageTouchPointVariant(v SignInPageTouchPointVariant)`

SetSignInPageTouchPointVariant sets SignInPageTouchPointVariant field to given value.

### HasSignInPageTouchPointVariant

`func (o *ThemeResponse) HasSignInPageTouchPointVariant() bool`

HasSignInPageTouchPointVariant returns a boolean if a field has been set.

### GetLinks

`func (o *ThemeResponse) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ThemeResponse) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ThemeResponse) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ThemeResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


