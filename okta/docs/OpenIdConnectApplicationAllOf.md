# OpenIdConnectApplicationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**OAuthApplicationCredentials**](OAuthApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [default to "oidc_client"]
**Settings** | Pointer to [**OpenIdConnectApplicationSettings**](OpenIdConnectApplicationSettings.md) |  | [optional] 

## Methods

### NewOpenIdConnectApplicationAllOf

`func NewOpenIdConnectApplicationAllOf() *OpenIdConnectApplicationAllOf`

NewOpenIdConnectApplicationAllOf instantiates a new OpenIdConnectApplicationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectApplicationAllOfWithDefaults

`func NewOpenIdConnectApplicationAllOfWithDefaults() *OpenIdConnectApplicationAllOf`

NewOpenIdConnectApplicationAllOfWithDefaults instantiates a new OpenIdConnectApplicationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *OpenIdConnectApplicationAllOf) GetCredentials() OAuthApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OpenIdConnectApplicationAllOf) GetCredentialsOk() (*OAuthApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OpenIdConnectApplicationAllOf) SetCredentials(v OAuthApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OpenIdConnectApplicationAllOf) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *OpenIdConnectApplicationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenIdConnectApplicationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenIdConnectApplicationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenIdConnectApplicationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *OpenIdConnectApplicationAllOf) GetSettings() OpenIdConnectApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *OpenIdConnectApplicationAllOf) GetSettingsOk() (*OpenIdConnectApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *OpenIdConnectApplicationAllOf) SetSettings(v OpenIdConnectApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *OpenIdConnectApplicationAllOf) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


