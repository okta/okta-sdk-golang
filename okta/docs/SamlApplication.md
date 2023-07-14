# SamlApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ApplicationCredentials**](ApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**SamlApplicationSettings**](SamlApplicationSettings.md) |  | [optional] 

## Methods

### NewSamlApplication

`func NewSamlApplication() *SamlApplication`

NewSamlApplication instantiates a new SamlApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlApplicationWithDefaults

`func NewSamlApplicationWithDefaults() *SamlApplication`

NewSamlApplicationWithDefaults instantiates a new SamlApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *SamlApplication) GetCredentials() ApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SamlApplication) GetCredentialsOk() (*ApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SamlApplication) SetCredentials(v ApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SamlApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *SamlApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SamlApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *SamlApplication) GetSettings() SamlApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SamlApplication) GetSettingsOk() (*SamlApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SamlApplication) SetSettings(v SamlApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SamlApplication) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


