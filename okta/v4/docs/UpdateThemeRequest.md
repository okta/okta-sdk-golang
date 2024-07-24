# UpdateThemeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTemplateTouchPointVariant** | **string** | Variant for email templates. You can publish a theme for email templates with different combinations of assets. Variants are preset combinations of those assets.  | [default to "OKTA_DEFAULT"]
**EndUserDashboardTouchPointVariant** | **string** | Variant for the Okta End-User Dashboard. You can publish a theme for end-user dashboard with different combinations of assets. Variants are preset combinations of those assets.  | [default to "OKTA_DEFAULT"]
**ErrorPageTouchPointVariant** | **string** | Variant for the error page. You can publish a theme for error page with different combinations of assets. Variants are preset combinations of those assets.  | [default to "OKTA_DEFAULT"]
**LoadingPageTouchPointVariant** | Pointer to **string** | Variant for the Okta loading page. You can publish a theme for Okta loading page with different combinations of assets. Variants are preset combinations of those assets.  | [optional] [default to "OKTA_DEFAULT"]
**PrimaryColorContrastHex** | Pointer to **string** | Primary color contrast hex code | [optional] 
**PrimaryColorHex** | **string** | Primary color hex code | 
**SecondaryColorContrastHex** | Pointer to **string** | Secondary color contrast hex code | [optional] 
**SecondaryColorHex** | **string** | Secondary color hex code | 
**SignInPageTouchPointVariant** | **string** | Variant for the Okta sign-in page. You can publish a theme for sign-in page with different combinations of assets. Variants are preset combinations of those assets. &gt; **Note:**  For a non-&#x60;OKTA_DEFAULT&#x60; variant, &#x60;primaryColorHex&#x60; is used for button background color and &#x60;primaryColorContrastHex&#x60; is used to optimize the opacity for button text.  | 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewUpdateThemeRequest

`func NewUpdateThemeRequest(emailTemplateTouchPointVariant string, endUserDashboardTouchPointVariant string, errorPageTouchPointVariant string, primaryColorHex string, secondaryColorHex string, signInPageTouchPointVariant string, ) *UpdateThemeRequest`

NewUpdateThemeRequest instantiates a new UpdateThemeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateThemeRequestWithDefaults

`func NewUpdateThemeRequestWithDefaults() *UpdateThemeRequest`

NewUpdateThemeRequestWithDefaults instantiates a new UpdateThemeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailTemplateTouchPointVariant

`func (o *UpdateThemeRequest) GetEmailTemplateTouchPointVariant() string`

GetEmailTemplateTouchPointVariant returns the EmailTemplateTouchPointVariant field if non-nil, zero value otherwise.

### GetEmailTemplateTouchPointVariantOk

`func (o *UpdateThemeRequest) GetEmailTemplateTouchPointVariantOk() (*string, bool)`

GetEmailTemplateTouchPointVariantOk returns a tuple with the EmailTemplateTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateTouchPointVariant

`func (o *UpdateThemeRequest) SetEmailTemplateTouchPointVariant(v string)`

SetEmailTemplateTouchPointVariant sets EmailTemplateTouchPointVariant field to given value.


### GetEndUserDashboardTouchPointVariant

`func (o *UpdateThemeRequest) GetEndUserDashboardTouchPointVariant() string`

GetEndUserDashboardTouchPointVariant returns the EndUserDashboardTouchPointVariant field if non-nil, zero value otherwise.

### GetEndUserDashboardTouchPointVariantOk

`func (o *UpdateThemeRequest) GetEndUserDashboardTouchPointVariantOk() (*string, bool)`

GetEndUserDashboardTouchPointVariantOk returns a tuple with the EndUserDashboardTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserDashboardTouchPointVariant

`func (o *UpdateThemeRequest) SetEndUserDashboardTouchPointVariant(v string)`

SetEndUserDashboardTouchPointVariant sets EndUserDashboardTouchPointVariant field to given value.


### GetErrorPageTouchPointVariant

`func (o *UpdateThemeRequest) GetErrorPageTouchPointVariant() string`

GetErrorPageTouchPointVariant returns the ErrorPageTouchPointVariant field if non-nil, zero value otherwise.

### GetErrorPageTouchPointVariantOk

`func (o *UpdateThemeRequest) GetErrorPageTouchPointVariantOk() (*string, bool)`

GetErrorPageTouchPointVariantOk returns a tuple with the ErrorPageTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPageTouchPointVariant

`func (o *UpdateThemeRequest) SetErrorPageTouchPointVariant(v string)`

SetErrorPageTouchPointVariant sets ErrorPageTouchPointVariant field to given value.


### GetLoadingPageTouchPointVariant

`func (o *UpdateThemeRequest) GetLoadingPageTouchPointVariant() string`

GetLoadingPageTouchPointVariant returns the LoadingPageTouchPointVariant field if non-nil, zero value otherwise.

### GetLoadingPageTouchPointVariantOk

`func (o *UpdateThemeRequest) GetLoadingPageTouchPointVariantOk() (*string, bool)`

GetLoadingPageTouchPointVariantOk returns a tuple with the LoadingPageTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingPageTouchPointVariant

`func (o *UpdateThemeRequest) SetLoadingPageTouchPointVariant(v string)`

SetLoadingPageTouchPointVariant sets LoadingPageTouchPointVariant field to given value.

### HasLoadingPageTouchPointVariant

`func (o *UpdateThemeRequest) HasLoadingPageTouchPointVariant() bool`

HasLoadingPageTouchPointVariant returns a boolean if a field has been set.

### GetPrimaryColorContrastHex

`func (o *UpdateThemeRequest) GetPrimaryColorContrastHex() string`

GetPrimaryColorContrastHex returns the PrimaryColorContrastHex field if non-nil, zero value otherwise.

### GetPrimaryColorContrastHexOk

`func (o *UpdateThemeRequest) GetPrimaryColorContrastHexOk() (*string, bool)`

GetPrimaryColorContrastHexOk returns a tuple with the PrimaryColorContrastHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColorContrastHex

`func (o *UpdateThemeRequest) SetPrimaryColorContrastHex(v string)`

SetPrimaryColorContrastHex sets PrimaryColorContrastHex field to given value.

### HasPrimaryColorContrastHex

`func (o *UpdateThemeRequest) HasPrimaryColorContrastHex() bool`

HasPrimaryColorContrastHex returns a boolean if a field has been set.

### GetPrimaryColorHex

`func (o *UpdateThemeRequest) GetPrimaryColorHex() string`

GetPrimaryColorHex returns the PrimaryColorHex field if non-nil, zero value otherwise.

### GetPrimaryColorHexOk

`func (o *UpdateThemeRequest) GetPrimaryColorHexOk() (*string, bool)`

GetPrimaryColorHexOk returns a tuple with the PrimaryColorHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColorHex

`func (o *UpdateThemeRequest) SetPrimaryColorHex(v string)`

SetPrimaryColorHex sets PrimaryColorHex field to given value.


### GetSecondaryColorContrastHex

`func (o *UpdateThemeRequest) GetSecondaryColorContrastHex() string`

GetSecondaryColorContrastHex returns the SecondaryColorContrastHex field if non-nil, zero value otherwise.

### GetSecondaryColorContrastHexOk

`func (o *UpdateThemeRequest) GetSecondaryColorContrastHexOk() (*string, bool)`

GetSecondaryColorContrastHexOk returns a tuple with the SecondaryColorContrastHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryColorContrastHex

`func (o *UpdateThemeRequest) SetSecondaryColorContrastHex(v string)`

SetSecondaryColorContrastHex sets SecondaryColorContrastHex field to given value.

### HasSecondaryColorContrastHex

`func (o *UpdateThemeRequest) HasSecondaryColorContrastHex() bool`

HasSecondaryColorContrastHex returns a boolean if a field has been set.

### GetSecondaryColorHex

`func (o *UpdateThemeRequest) GetSecondaryColorHex() string`

GetSecondaryColorHex returns the SecondaryColorHex field if non-nil, zero value otherwise.

### GetSecondaryColorHexOk

`func (o *UpdateThemeRequest) GetSecondaryColorHexOk() (*string, bool)`

GetSecondaryColorHexOk returns a tuple with the SecondaryColorHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryColorHex

`func (o *UpdateThemeRequest) SetSecondaryColorHex(v string)`

SetSecondaryColorHex sets SecondaryColorHex field to given value.


### GetSignInPageTouchPointVariant

`func (o *UpdateThemeRequest) GetSignInPageTouchPointVariant() string`

GetSignInPageTouchPointVariant returns the SignInPageTouchPointVariant field if non-nil, zero value otherwise.

### GetSignInPageTouchPointVariantOk

`func (o *UpdateThemeRequest) GetSignInPageTouchPointVariantOk() (*string, bool)`

GetSignInPageTouchPointVariantOk returns a tuple with the SignInPageTouchPointVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInPageTouchPointVariant

`func (o *UpdateThemeRequest) SetSignInPageTouchPointVariant(v string)`

SetSignInPageTouchPointVariant sets SignInPageTouchPointVariant field to given value.


### GetLinks

`func (o *UpdateThemeRequest) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateThemeRequest) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateThemeRequest) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateThemeRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


