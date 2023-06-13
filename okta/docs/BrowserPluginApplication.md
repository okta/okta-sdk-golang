# BrowserPluginApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**SchemeApplicationCredentials**](SchemeApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**SwaApplicationSettings**](SwaApplicationSettings.md) |  | [optional] 

## Methods

### NewBrowserPluginApplication

`func NewBrowserPluginApplication() *BrowserPluginApplication`

NewBrowserPluginApplication instantiates a new BrowserPluginApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserPluginApplicationWithDefaults

`func NewBrowserPluginApplicationWithDefaults() *BrowserPluginApplication`

NewBrowserPluginApplicationWithDefaults instantiates a new BrowserPluginApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *BrowserPluginApplication) GetCredentials() SchemeApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BrowserPluginApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BrowserPluginApplication) SetCredentials(v SchemeApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BrowserPluginApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *BrowserPluginApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrowserPluginApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrowserPluginApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BrowserPluginApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *BrowserPluginApplication) GetSettings() SwaApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *BrowserPluginApplication) GetSettingsOk() (*SwaApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *BrowserPluginApplication) SetSettings(v SwaApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *BrowserPluginApplication) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


