# AutoLoginApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**SchemeApplicationCredentials**](SchemeApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**AutoLoginApplicationSettings**](AutoLoginApplicationSettings.md) |  | [optional] 

## Methods

### NewAutoLoginApplication

`func NewAutoLoginApplication() *AutoLoginApplication`

NewAutoLoginApplication instantiates a new AutoLoginApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoLoginApplicationWithDefaults

`func NewAutoLoginApplicationWithDefaults() *AutoLoginApplication`

NewAutoLoginApplicationWithDefaults instantiates a new AutoLoginApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *AutoLoginApplication) GetCredentials() SchemeApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AutoLoginApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AutoLoginApplication) SetCredentials(v SchemeApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AutoLoginApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *AutoLoginApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoLoginApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoLoginApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoLoginApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *AutoLoginApplication) GetSettings() AutoLoginApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AutoLoginApplication) GetSettingsOk() (*AutoLoginApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AutoLoginApplication) SetSettings(v AutoLoginApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AutoLoginApplication) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


