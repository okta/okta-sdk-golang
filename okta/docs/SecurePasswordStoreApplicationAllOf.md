# SecurePasswordStoreApplicationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**SchemeApplicationCredentials**](SchemeApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [default to "template_sps"]
**Settings** | Pointer to [**SecurePasswordStoreApplicationSettings**](SecurePasswordStoreApplicationSettings.md) |  | [optional] 

## Methods

### NewSecurePasswordStoreApplicationAllOf

`func NewSecurePasswordStoreApplicationAllOf() *SecurePasswordStoreApplicationAllOf`

NewSecurePasswordStoreApplicationAllOf instantiates a new SecurePasswordStoreApplicationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurePasswordStoreApplicationAllOfWithDefaults

`func NewSecurePasswordStoreApplicationAllOfWithDefaults() *SecurePasswordStoreApplicationAllOf`

NewSecurePasswordStoreApplicationAllOfWithDefaults instantiates a new SecurePasswordStoreApplicationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *SecurePasswordStoreApplicationAllOf) GetCredentials() SchemeApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SecurePasswordStoreApplicationAllOf) GetCredentialsOk() (*SchemeApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SecurePasswordStoreApplicationAllOf) SetCredentials(v SchemeApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SecurePasswordStoreApplicationAllOf) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *SecurePasswordStoreApplicationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurePasswordStoreApplicationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurePasswordStoreApplicationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurePasswordStoreApplicationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *SecurePasswordStoreApplicationAllOf) GetSettings() SecurePasswordStoreApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SecurePasswordStoreApplicationAllOf) GetSettingsOk() (*SecurePasswordStoreApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SecurePasswordStoreApplicationAllOf) SetSettings(v SecurePasswordStoreApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SecurePasswordStoreApplicationAllOf) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


