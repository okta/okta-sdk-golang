# OpenIdConnectApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**OAuthApplicationCredentials**](OAuthApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [default to "oidc_client"]
**Settings** | Pointer to [**OpenIdConnectApplicationSettings**](OpenIdConnectApplicationSettings.md) |  | [optional] 

## Methods

### NewOpenIdConnectApplication

`func NewOpenIdConnectApplication() *OpenIdConnectApplication`

NewOpenIdConnectApplication instantiates a new OpenIdConnectApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectApplicationWithDefaults

`func NewOpenIdConnectApplicationWithDefaults() *OpenIdConnectApplication`

NewOpenIdConnectApplicationWithDefaults instantiates a new OpenIdConnectApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *OpenIdConnectApplication) GetCredentials() OAuthApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OpenIdConnectApplication) GetCredentialsOk() (*OAuthApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OpenIdConnectApplication) SetCredentials(v OAuthApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OpenIdConnectApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *OpenIdConnectApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenIdConnectApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenIdConnectApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenIdConnectApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *OpenIdConnectApplication) GetSettings() OpenIdConnectApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *OpenIdConnectApplication) GetSettingsOk() (*OpenIdConnectApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *OpenIdConnectApplication) SetSettings(v OpenIdConnectApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *OpenIdConnectApplication) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


