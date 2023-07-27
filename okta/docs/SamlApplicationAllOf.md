# SamlApplicationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ApplicationCredentials**](ApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**SamlApplicationSettings**](SamlApplicationSettings.md) |  | [optional] 

## Methods

### NewSamlApplicationAllOf

`func NewSamlApplicationAllOf() *SamlApplicationAllOf`

NewSamlApplicationAllOf instantiates a new SamlApplicationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlApplicationAllOfWithDefaults

`func NewSamlApplicationAllOfWithDefaults() *SamlApplicationAllOf`

NewSamlApplicationAllOfWithDefaults instantiates a new SamlApplicationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *SamlApplicationAllOf) GetCredentials() ApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SamlApplicationAllOf) GetCredentialsOk() (*ApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SamlApplicationAllOf) SetCredentials(v ApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SamlApplicationAllOf) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *SamlApplicationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlApplicationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlApplicationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SamlApplicationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *SamlApplicationAllOf) GetSettings() SamlApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SamlApplicationAllOf) GetSettingsOk() (*SamlApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SamlApplicationAllOf) SetSettings(v SamlApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SamlApplicationAllOf) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


