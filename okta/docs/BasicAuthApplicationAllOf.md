# BasicAuthApplicationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**SchemeApplicationCredentials**](SchemeApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [default to "template_basic_auth"]
**Settings** | Pointer to [**BasicApplicationSettings**](BasicApplicationSettings.md) |  | [optional] 

## Methods

### NewBasicAuthApplicationAllOf

`func NewBasicAuthApplicationAllOf() *BasicAuthApplicationAllOf`

NewBasicAuthApplicationAllOf instantiates a new BasicAuthApplicationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicAuthApplicationAllOfWithDefaults

`func NewBasicAuthApplicationAllOfWithDefaults() *BasicAuthApplicationAllOf`

NewBasicAuthApplicationAllOfWithDefaults instantiates a new BasicAuthApplicationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *BasicAuthApplicationAllOf) GetCredentials() SchemeApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BasicAuthApplicationAllOf) GetCredentialsOk() (*SchemeApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BasicAuthApplicationAllOf) SetCredentials(v SchemeApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BasicAuthApplicationAllOf) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *BasicAuthApplicationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicAuthApplicationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicAuthApplicationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicAuthApplicationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *BasicAuthApplicationAllOf) GetSettings() BasicApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *BasicAuthApplicationAllOf) GetSettingsOk() (*BasicApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *BasicAuthApplicationAllOf) SetSettings(v BasicApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *BasicAuthApplicationAllOf) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


