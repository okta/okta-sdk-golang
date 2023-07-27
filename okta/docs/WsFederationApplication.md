# WsFederationApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ApplicationCredentials**](ApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [default to "template_wsfed"]
**Settings** | Pointer to [**WsFederationApplicationSettings**](WsFederationApplicationSettings.md) |  | [optional] 

## Methods

### NewWsFederationApplication

`func NewWsFederationApplication() *WsFederationApplication`

NewWsFederationApplication instantiates a new WsFederationApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsFederationApplicationWithDefaults

`func NewWsFederationApplicationWithDefaults() *WsFederationApplication`

NewWsFederationApplicationWithDefaults instantiates a new WsFederationApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *WsFederationApplication) GetCredentials() ApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *WsFederationApplication) GetCredentialsOk() (*ApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *WsFederationApplication) SetCredentials(v ApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *WsFederationApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *WsFederationApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WsFederationApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WsFederationApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WsFederationApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *WsFederationApplication) GetSettings() WsFederationApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WsFederationApplication) GetSettingsOk() (*WsFederationApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WsFederationApplication) SetSettings(v WsFederationApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *WsFederationApplication) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


