# OINAutoLoginApplicationSettingsSignOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignOnMode** | Pointer to **interface{}** |  | [optional] [default to AUTO_LOGIN]
**LoginUrl** | **string** | Primary URL of the sign-in page for this app | 
**RedirectUrl** | Pointer to **string** | Secondary URL of the sign-in page for this app | [optional] 

## Methods

### NewOINAutoLoginApplicationSettingsSignOn

`func NewOINAutoLoginApplicationSettingsSignOn(loginUrl string, ) *OINAutoLoginApplicationSettingsSignOn`

NewOINAutoLoginApplicationSettingsSignOn instantiates a new OINAutoLoginApplicationSettingsSignOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOINAutoLoginApplicationSettingsSignOnWithDefaults

`func NewOINAutoLoginApplicationSettingsSignOnWithDefaults() *OINAutoLoginApplicationSettingsSignOn`

NewOINAutoLoginApplicationSettingsSignOnWithDefaults instantiates a new OINAutoLoginApplicationSettingsSignOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignOnMode

`func (o *OINAutoLoginApplicationSettingsSignOn) GetSignOnMode() interface{}`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *OINAutoLoginApplicationSettingsSignOn) GetSignOnModeOk() (*interface{}, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *OINAutoLoginApplicationSettingsSignOn) SetSignOnMode(v interface{})`

SetSignOnMode sets SignOnMode field to given value.

### HasSignOnMode

`func (o *OINAutoLoginApplicationSettingsSignOn) HasSignOnMode() bool`

HasSignOnMode returns a boolean if a field has been set.

### SetSignOnModeNil

`func (o *OINAutoLoginApplicationSettingsSignOn) SetSignOnModeNil(b bool)`

 SetSignOnModeNil sets the value for SignOnMode to be an explicit nil

### UnsetSignOnMode
`func (o *OINAutoLoginApplicationSettingsSignOn) UnsetSignOnMode()`

UnsetSignOnMode ensures that no value is present for SignOnMode, not even an explicit nil
### GetLoginUrl

`func (o *OINAutoLoginApplicationSettingsSignOn) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *OINAutoLoginApplicationSettingsSignOn) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *OINAutoLoginApplicationSettingsSignOn) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.


### GetRedirectUrl

`func (o *OINAutoLoginApplicationSettingsSignOn) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *OINAutoLoginApplicationSettingsSignOn) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *OINAutoLoginApplicationSettingsSignOn) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *OINAutoLoginApplicationSettingsSignOn) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


