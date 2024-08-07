# SecurePasswordStoreApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**SchemeApplicationCredentials**](SchemeApplicationCredentials.md) |  | [optional] 
**Name** | **string** | &#x60;template_sps&#x60; is the key name for a SWA app instance that uses HTTP POST and doesn&#39;t require a browser plugin | 
**Settings** | [**SecurePasswordStoreApplicationSettings**](SecurePasswordStoreApplicationSettings.md) |  | 

## Methods

### NewSecurePasswordStoreApplication

`func NewSecurePasswordStoreApplication(name string, settings SecurePasswordStoreApplicationSettings, ) *SecurePasswordStoreApplication`

NewSecurePasswordStoreApplication instantiates a new SecurePasswordStoreApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurePasswordStoreApplicationWithDefaults

`func NewSecurePasswordStoreApplicationWithDefaults() *SecurePasswordStoreApplication`

NewSecurePasswordStoreApplicationWithDefaults instantiates a new SecurePasswordStoreApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *SecurePasswordStoreApplication) GetCredentials() SchemeApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SecurePasswordStoreApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SecurePasswordStoreApplication) SetCredentials(v SchemeApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SecurePasswordStoreApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *SecurePasswordStoreApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurePasswordStoreApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurePasswordStoreApplication) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *SecurePasswordStoreApplication) GetSettings() SecurePasswordStoreApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SecurePasswordStoreApplication) GetSettingsOk() (*SecurePasswordStoreApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SecurePasswordStoreApplication) SetSettings(v SecurePasswordStoreApplicationSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


