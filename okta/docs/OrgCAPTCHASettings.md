# OrgCAPTCHASettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptchaId** | Pointer to **string** | The unique key of the associated CAPTCHA instance | [optional] 
**EnabledPages** | Pointer to **[]string** | An array of pages that have CAPTCHA enabled | [optional] 
**Links** | Pointer to [**OrgCAPTCHASettingsLinks**](OrgCAPTCHASettingsLinks.md) |  | [optional] 

## Methods

### NewOrgCAPTCHASettings

`func NewOrgCAPTCHASettings() *OrgCAPTCHASettings`

NewOrgCAPTCHASettings instantiates a new OrgCAPTCHASettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgCAPTCHASettingsWithDefaults

`func NewOrgCAPTCHASettingsWithDefaults() *OrgCAPTCHASettings`

NewOrgCAPTCHASettingsWithDefaults instantiates a new OrgCAPTCHASettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptchaId

`func (o *OrgCAPTCHASettings) GetCaptchaId() string`

GetCaptchaId returns the CaptchaId field if non-nil, zero value otherwise.

### GetCaptchaIdOk

`func (o *OrgCAPTCHASettings) GetCaptchaIdOk() (*string, bool)`

GetCaptchaIdOk returns a tuple with the CaptchaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaId

`func (o *OrgCAPTCHASettings) SetCaptchaId(v string)`

SetCaptchaId sets CaptchaId field to given value.

### HasCaptchaId

`func (o *OrgCAPTCHASettings) HasCaptchaId() bool`

HasCaptchaId returns a boolean if a field has been set.

### GetEnabledPages

`func (o *OrgCAPTCHASettings) GetEnabledPages() []string`

GetEnabledPages returns the EnabledPages field if non-nil, zero value otherwise.

### GetEnabledPagesOk

`func (o *OrgCAPTCHASettings) GetEnabledPagesOk() (*[]string, bool)`

GetEnabledPagesOk returns a tuple with the EnabledPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledPages

`func (o *OrgCAPTCHASettings) SetEnabledPages(v []string)`

SetEnabledPages sets EnabledPages field to given value.

### HasEnabledPages

`func (o *OrgCAPTCHASettings) HasEnabledPages() bool`

HasEnabledPages returns a boolean if a field has been set.

### GetLinks

`func (o *OrgCAPTCHASettings) GetLinks() OrgCAPTCHASettingsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrgCAPTCHASettings) GetLinksOk() (*OrgCAPTCHASettingsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrgCAPTCHASettings) SetLinks(v OrgCAPTCHASettingsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrgCAPTCHASettings) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


