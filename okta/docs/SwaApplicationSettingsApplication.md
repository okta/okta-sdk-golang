# SwaApplicationSettingsApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ButtonField** | **string** | CSS selector for the **Sign-In** button in the sign-in form (for SWA apps with the &#x60;template_swa&#x60; app name definition) | 
**ButtonSelector** | Pointer to **string** | CSS selector for the **Sign-In**  button in the sign-in form (for three-field SWA apps with the &#x60;template_swa3field&#x60; app name definition) | [optional] 
**ExtraFieldSelector** | Pointer to **string** | Enter the CSS selector for the extra field (for three-field SWA apps with the &#x60;template_swa3field&#x60; app name definition). | [optional] 
**ExtraFieldValue** | Pointer to **string** | Enter the value for the extra field in the form (for three-field SWA apps with the &#x60;template_swa3field&#x60; app name definition). | [optional] 
**LoginUrlRegex** | Pointer to **string** | A regular expression that further restricts targetURL to the specified regular expression | [optional] 
**PasswordField** | **string** | CSS selector for the **Password** field in the sign-in form (for SWA apps with the &#x60;template_swa&#x60; app name definition) | 
**PasswordSelector** | Pointer to **string** | CSS selector for the **Password** field in the sign-in form (for three-field SWA apps with the &#x60;template_swa3field&#x60; app name definition) | [optional] 
**TargetURL** | Pointer to **string** | The URL of the sign-in page for this app (for three-field SWA apps with the &#x60;template_swa3field&#x60; app name definition) | [optional] 
**Url** | **string** | The URL of the sign-in page for this app (for SWA apps with the &#x60;template_swa&#x60; app name definition) | 
**UsernameField** | **string** | CSS selector for the **Username** field in the sign-in form (for SWA apps with the &#x60;template_swa&#x60; app name definition) | 
**UserNameSelector** | Pointer to **string** | CSS selector for the **Username** field in the sign-in form (for three-field SWA apps with the &#x60;template_swa3field&#x60; app name definition) | [optional] 

## Methods

### NewSwaApplicationSettingsApplication

`func NewSwaApplicationSettingsApplication(buttonField string, passwordField string, url string, usernameField string, ) *SwaApplicationSettingsApplication`

NewSwaApplicationSettingsApplication instantiates a new SwaApplicationSettingsApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwaApplicationSettingsApplicationWithDefaults

`func NewSwaApplicationSettingsApplicationWithDefaults() *SwaApplicationSettingsApplication`

NewSwaApplicationSettingsApplicationWithDefaults instantiates a new SwaApplicationSettingsApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtonField

`func (o *SwaApplicationSettingsApplication) GetButtonField() string`

GetButtonField returns the ButtonField field if non-nil, zero value otherwise.

### GetButtonFieldOk

`func (o *SwaApplicationSettingsApplication) GetButtonFieldOk() (*string, bool)`

GetButtonFieldOk returns a tuple with the ButtonField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonField

`func (o *SwaApplicationSettingsApplication) SetButtonField(v string)`

SetButtonField sets ButtonField field to given value.


### GetButtonSelector

`func (o *SwaApplicationSettingsApplication) GetButtonSelector() string`

GetButtonSelector returns the ButtonSelector field if non-nil, zero value otherwise.

### GetButtonSelectorOk

`func (o *SwaApplicationSettingsApplication) GetButtonSelectorOk() (*string, bool)`

GetButtonSelectorOk returns a tuple with the ButtonSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonSelector

`func (o *SwaApplicationSettingsApplication) SetButtonSelector(v string)`

SetButtonSelector sets ButtonSelector field to given value.

### HasButtonSelector

`func (o *SwaApplicationSettingsApplication) HasButtonSelector() bool`

HasButtonSelector returns a boolean if a field has been set.

### GetExtraFieldSelector

`func (o *SwaApplicationSettingsApplication) GetExtraFieldSelector() string`

GetExtraFieldSelector returns the ExtraFieldSelector field if non-nil, zero value otherwise.

### GetExtraFieldSelectorOk

`func (o *SwaApplicationSettingsApplication) GetExtraFieldSelectorOk() (*string, bool)`

GetExtraFieldSelectorOk returns a tuple with the ExtraFieldSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFieldSelector

`func (o *SwaApplicationSettingsApplication) SetExtraFieldSelector(v string)`

SetExtraFieldSelector sets ExtraFieldSelector field to given value.

### HasExtraFieldSelector

`func (o *SwaApplicationSettingsApplication) HasExtraFieldSelector() bool`

HasExtraFieldSelector returns a boolean if a field has been set.

### GetExtraFieldValue

`func (o *SwaApplicationSettingsApplication) GetExtraFieldValue() string`

GetExtraFieldValue returns the ExtraFieldValue field if non-nil, zero value otherwise.

### GetExtraFieldValueOk

`func (o *SwaApplicationSettingsApplication) GetExtraFieldValueOk() (*string, bool)`

GetExtraFieldValueOk returns a tuple with the ExtraFieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFieldValue

`func (o *SwaApplicationSettingsApplication) SetExtraFieldValue(v string)`

SetExtraFieldValue sets ExtraFieldValue field to given value.

### HasExtraFieldValue

`func (o *SwaApplicationSettingsApplication) HasExtraFieldValue() bool`

HasExtraFieldValue returns a boolean if a field has been set.

### GetLoginUrlRegex

`func (o *SwaApplicationSettingsApplication) GetLoginUrlRegex() string`

GetLoginUrlRegex returns the LoginUrlRegex field if non-nil, zero value otherwise.

### GetLoginUrlRegexOk

`func (o *SwaApplicationSettingsApplication) GetLoginUrlRegexOk() (*string, bool)`

GetLoginUrlRegexOk returns a tuple with the LoginUrlRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrlRegex

`func (o *SwaApplicationSettingsApplication) SetLoginUrlRegex(v string)`

SetLoginUrlRegex sets LoginUrlRegex field to given value.

### HasLoginUrlRegex

`func (o *SwaApplicationSettingsApplication) HasLoginUrlRegex() bool`

HasLoginUrlRegex returns a boolean if a field has been set.

### GetPasswordField

`func (o *SwaApplicationSettingsApplication) GetPasswordField() string`

GetPasswordField returns the PasswordField field if non-nil, zero value otherwise.

### GetPasswordFieldOk

`func (o *SwaApplicationSettingsApplication) GetPasswordFieldOk() (*string, bool)`

GetPasswordFieldOk returns a tuple with the PasswordField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordField

`func (o *SwaApplicationSettingsApplication) SetPasswordField(v string)`

SetPasswordField sets PasswordField field to given value.


### GetPasswordSelector

`func (o *SwaApplicationSettingsApplication) GetPasswordSelector() string`

GetPasswordSelector returns the PasswordSelector field if non-nil, zero value otherwise.

### GetPasswordSelectorOk

`func (o *SwaApplicationSettingsApplication) GetPasswordSelectorOk() (*string, bool)`

GetPasswordSelectorOk returns a tuple with the PasswordSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSelector

`func (o *SwaApplicationSettingsApplication) SetPasswordSelector(v string)`

SetPasswordSelector sets PasswordSelector field to given value.

### HasPasswordSelector

`func (o *SwaApplicationSettingsApplication) HasPasswordSelector() bool`

HasPasswordSelector returns a boolean if a field has been set.

### GetTargetURL

`func (o *SwaApplicationSettingsApplication) GetTargetURL() string`

GetTargetURL returns the TargetURL field if non-nil, zero value otherwise.

### GetTargetURLOk

`func (o *SwaApplicationSettingsApplication) GetTargetURLOk() (*string, bool)`

GetTargetURLOk returns a tuple with the TargetURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetURL

`func (o *SwaApplicationSettingsApplication) SetTargetURL(v string)`

SetTargetURL sets TargetURL field to given value.

### HasTargetURL

`func (o *SwaApplicationSettingsApplication) HasTargetURL() bool`

HasTargetURL returns a boolean if a field has been set.

### GetUrl

`func (o *SwaApplicationSettingsApplication) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SwaApplicationSettingsApplication) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SwaApplicationSettingsApplication) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsernameField

`func (o *SwaApplicationSettingsApplication) GetUsernameField() string`

GetUsernameField returns the UsernameField field if non-nil, zero value otherwise.

### GetUsernameFieldOk

`func (o *SwaApplicationSettingsApplication) GetUsernameFieldOk() (*string, bool)`

GetUsernameFieldOk returns a tuple with the UsernameField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameField

`func (o *SwaApplicationSettingsApplication) SetUsernameField(v string)`

SetUsernameField sets UsernameField field to given value.


### GetUserNameSelector

`func (o *SwaApplicationSettingsApplication) GetUserNameSelector() string`

GetUserNameSelector returns the UserNameSelector field if non-nil, zero value otherwise.

### GetUserNameSelectorOk

`func (o *SwaApplicationSettingsApplication) GetUserNameSelectorOk() (*string, bool)`

GetUserNameSelectorOk returns a tuple with the UserNameSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameSelector

`func (o *SwaApplicationSettingsApplication) SetUserNameSelector(v string)`

SetUserNameSelector sets UserNameSelector field to given value.

### HasUserNameSelector

`func (o *SwaApplicationSettingsApplication) HasUserNameSelector() bool`

HasUserNameSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


