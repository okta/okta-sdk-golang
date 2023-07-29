# Theme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundImage** | Pointer to **string** |  | [optional] [readonly] 
**EmailTemplateTouchPointVariant** | Pointer to [**EmailTemplateTouchPointVariant**](EmailTemplateTouchPointVariant.md) |  | [optional] 
**EndUserDashboardTouchPointVariant** | Pointer to [**EndUserDashboardTouchPointVariant**](EndUserDashboardTouchPointVariant.md) |  | [optional] 
**ErrorPageTouchPointVariant** | Pointer to [**ErrorPageTouchPointVariant**](ErrorPageTouchPointVariant.md) |  | [optional] 
**LoadingPageTouchPointVariant** | Pointer to [**LoadingPageTouchPointVariant**](LoadingPageTouchPointVariant.md) |  | [optional] 
**PrimaryColorContrastHex** | Pointer to **string** |  | [optional] 
**PrimaryColorHex** | Pointer to **string** |  | [optional] 
**SecondaryColorContrastHex** | Pointer to **string** |  | [optional] 
**SecondaryColorHex** | Pointer to **string** |  | [optional] 
**SignInPageTouchPointVariant** | Pointer to [**SignInPageTouchPointVariant**](SignInPageTouchPointVariant.md) |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewTheme

`func NewTheme() *Theme`

NewTheme instantiates a new Theme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThemeWithDefaults

`func NewThemeWithDefaults() *Theme`

NewThemeWithDefaults instantiates a new Theme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundImage

`func (o *Theme) GetBackgroundImage() string`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *Theme) GetBackgroundImageOk() (*string, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *Theme) SetBackgroundImage(v string)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *Theme) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### GetEmailTemplateTouchPointVariant

`func (o *Theme) GetEmailTemplateTouchPointVariant() EmailTemplateTouchPointVariant`

GetEmailTemplateTouchPointVariant returns the EmailTemplateTouchPointVariant field if non-nil, zero value otherwise.

### GetEmailTemplateTouchPointVariantOk

`func (o *Theme) GetEmailTemplateTouchPointVariantOk() (*EmailTemplateTouchPointVariant, bool)`

GetEmailTemplateTouchPointVariantOk returns a tuple with the EmailTemplateTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateTouchPointVariant

`func (o *Theme) SetEmailTemplateTouchPointVariant(v EmailTemplateTouchPointVariant)`

SetEmailTemplateTouchPointVariant sets EmailTemplateTouchPointVariant field to given value.

### HasEmailTemplateTouchPointVariant

`func (o *Theme) HasEmailTemplateTouchPointVariant() bool`

HasEmailTemplateTouchPointVariant returns a boolean if a field has been set.

### GetEndUserDashboardTouchPointVariant

`func (o *Theme) GetEndUserDashboardTouchPointVariant() EndUserDashboardTouchPointVariant`

GetEndUserDashboardTouchPointVariant returns the EndUserDashboardTouchPointVariant field if non-nil, zero value otherwise.

### GetEndUserDashboardTouchPointVariantOk

`func (o *Theme) GetEndUserDashboardTouchPointVariantOk() (*EndUserDashboardTouchPointVariant, bool)`

GetEndUserDashboardTouchPointVariantOk returns a tuple with the EndUserDashboardTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserDashboardTouchPointVariant

`func (o *Theme) SetEndUserDashboardTouchPointVariant(v EndUserDashboardTouchPointVariant)`

SetEndUserDashboardTouchPointVariant sets EndUserDashboardTouchPointVariant field to given value.

### HasEndUserDashboardTouchPointVariant

`func (o *Theme) HasEndUserDashboardTouchPointVariant() bool`

HasEndUserDashboardTouchPointVariant returns a boolean if a field has been set.

### GetErrorPageTouchPointVariant

`func (o *Theme) GetErrorPageTouchPointVariant() ErrorPageTouchPointVariant`

GetErrorPageTouchPointVariant returns the ErrorPageTouchPointVariant field if non-nil, zero value otherwise.

### GetErrorPageTouchPointVariantOk

`func (o *Theme) GetErrorPageTouchPointVariantOk() (*ErrorPageTouchPointVariant, bool)`

GetErrorPageTouchPointVariantOk returns a tuple with the ErrorPageTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPageTouchPointVariant

`func (o *Theme) SetErrorPageTouchPointVariant(v ErrorPageTouchPointVariant)`

SetErrorPageTouchPointVariant sets ErrorPageTouchPointVariant field to given value.

### HasErrorPageTouchPointVariant

`func (o *Theme) HasErrorPageTouchPointVariant() bool`

HasErrorPageTouchPointVariant returns a boolean if a field has been set.

### GetLoadingPageTouchPointVariant

`func (o *Theme) GetLoadingPageTouchPointVariant() LoadingPageTouchPointVariant`

GetLoadingPageTouchPointVariant returns the LoadingPageTouchPointVariant field if non-nil, zero value otherwise.

### GetLoadingPageTouchPointVariantOk

`func (o *Theme) GetLoadingPageTouchPointVariantOk() (*LoadingPageTouchPointVariant, bool)`

GetLoadingPageTouchPointVariantOk returns a tuple with the LoadingPageTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingPageTouchPointVariant

`func (o *Theme) SetLoadingPageTouchPointVariant(v LoadingPageTouchPointVariant)`

SetLoadingPageTouchPointVariant sets LoadingPageTouchPointVariant field to given value.

### HasLoadingPageTouchPointVariant

`func (o *Theme) HasLoadingPageTouchPointVariant() bool`

HasLoadingPageTouchPointVariant returns a boolean if a field has been set.

### GetPrimaryColorContrastHex

`func (o *Theme) GetPrimaryColorContrastHex() string`

GetPrimaryColorContrastHex returns the PrimaryColorContrastHex field if non-nil, zero value otherwise.

### GetPrimaryColorContrastHexOk

`func (o *Theme) GetPrimaryColorContrastHexOk() (*string, bool)`

GetPrimaryColorContrastHexOk returns a tuple with the PrimaryColorContrastHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColorContrastHex

`func (o *Theme) SetPrimaryColorContrastHex(v string)`

SetPrimaryColorContrastHex sets PrimaryColorContrastHex field to given value.

### HasPrimaryColorContrastHex

`func (o *Theme) HasPrimaryColorContrastHex() bool`

HasPrimaryColorContrastHex returns a boolean if a field has been set.

### GetPrimaryColorHex

`func (o *Theme) GetPrimaryColorHex() string`

GetPrimaryColorHex returns the PrimaryColorHex field if non-nil, zero value otherwise.

### GetPrimaryColorHexOk

`func (o *Theme) GetPrimaryColorHexOk() (*string, bool)`

GetPrimaryColorHexOk returns a tuple with the PrimaryColorHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColorHex

`func (o *Theme) SetPrimaryColorHex(v string)`

SetPrimaryColorHex sets PrimaryColorHex field to given value.

### HasPrimaryColorHex

`func (o *Theme) HasPrimaryColorHex() bool`

HasPrimaryColorHex returns a boolean if a field has been set.

### GetSecondaryColorContrastHex

`func (o *Theme) GetSecondaryColorContrastHex() string`

GetSecondaryColorContrastHex returns the SecondaryColorContrastHex field if non-nil, zero value otherwise.

### GetSecondaryColorContrastHexOk

`func (o *Theme) GetSecondaryColorContrastHexOk() (*string, bool)`

GetSecondaryColorContrastHexOk returns a tuple with the SecondaryColorContrastHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryColorContrastHex

`func (o *Theme) SetSecondaryColorContrastHex(v string)`

SetSecondaryColorContrastHex sets SecondaryColorContrastHex field to given value.

### HasSecondaryColorContrastHex

`func (o *Theme) HasSecondaryColorContrastHex() bool`

HasSecondaryColorContrastHex returns a boolean if a field has been set.

### GetSecondaryColorHex

`func (o *Theme) GetSecondaryColorHex() string`

GetSecondaryColorHex returns the SecondaryColorHex field if non-nil, zero value otherwise.

### GetSecondaryColorHexOk

`func (o *Theme) GetSecondaryColorHexOk() (*string, bool)`

GetSecondaryColorHexOk returns a tuple with the SecondaryColorHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryColorHex

`func (o *Theme) SetSecondaryColorHex(v string)`

SetSecondaryColorHex sets SecondaryColorHex field to given value.

### HasSecondaryColorHex

`func (o *Theme) HasSecondaryColorHex() bool`

HasSecondaryColorHex returns a boolean if a field has been set.

### GetSignInPageTouchPointVariant

`func (o *Theme) GetSignInPageTouchPointVariant() SignInPageTouchPointVariant`

GetSignInPageTouchPointVariant returns the SignInPageTouchPointVariant field if non-nil, zero value otherwise.

### GetSignInPageTouchPointVariantOk

`func (o *Theme) GetSignInPageTouchPointVariantOk() (*SignInPageTouchPointVariant, bool)`

GetSignInPageTouchPointVariantOk returns a tuple with the SignInPageTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInPageTouchPointVariant

`func (o *Theme) SetSignInPageTouchPointVariant(v SignInPageTouchPointVariant)`

SetSignInPageTouchPointVariant sets SignInPageTouchPointVariant field to given value.

### HasSignInPageTouchPointVariant

`func (o *Theme) HasSignInPageTouchPointVariant() bool`

HasSignInPageTouchPointVariant returns a boolean if a field has been set.

### GetLinks

`func (o *Theme) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Theme) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Theme) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Theme) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


