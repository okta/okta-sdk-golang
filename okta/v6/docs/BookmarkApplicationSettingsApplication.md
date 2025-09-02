# BookmarkApplicationSettingsApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestIntegration** | Pointer to **bool** | Would you like Okta to add an integration for this app? | [optional] [default to false]
**Url** | **string** | The URL of the launch page for this app | 

## Methods

### NewBookmarkApplicationSettingsApplication

`func NewBookmarkApplicationSettingsApplication(url string, ) *BookmarkApplicationSettingsApplication`

NewBookmarkApplicationSettingsApplication instantiates a new BookmarkApplicationSettingsApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkApplicationSettingsApplicationWithDefaults

`func NewBookmarkApplicationSettingsApplicationWithDefaults() *BookmarkApplicationSettingsApplication`

NewBookmarkApplicationSettingsApplicationWithDefaults instantiates a new BookmarkApplicationSettingsApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestIntegration

`func (o *BookmarkApplicationSettingsApplication) GetRequestIntegration() bool`

GetRequestIntegration returns the RequestIntegration field if non-nil, zero value otherwise.

### GetRequestIntegrationOk

`func (o *BookmarkApplicationSettingsApplication) GetRequestIntegrationOk() (*bool, bool)`

GetRequestIntegrationOk returns a tuple with the RequestIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIntegration

`func (o *BookmarkApplicationSettingsApplication) SetRequestIntegration(v bool)`

SetRequestIntegration sets RequestIntegration field to given value.

### HasRequestIntegration

`func (o *BookmarkApplicationSettingsApplication) HasRequestIntegration() bool`

HasRequestIntegration returns a boolean if a field has been set.

### GetUrl

`func (o *BookmarkApplicationSettingsApplication) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BookmarkApplicationSettingsApplication) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BookmarkApplicationSettingsApplication) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


