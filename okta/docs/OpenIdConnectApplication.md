# OpenIdConnectApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**OAuthApplicationCredentials**](OAuthApplicationCredentials.md) |  | 
**Name** | **string** | &#x60;oidc_client&#x60; is the key name for an OIDC app instance | 
**Settings** | [**OpenIdConnectApplicationSettings**](OpenIdConnectApplicationSettings.md) |  | 

## Methods

### NewOpenIdConnectApplication

`func NewOpenIdConnectApplication(credentials OAuthApplicationCredentials, name string, settings OpenIdConnectApplicationSettings, ) *OpenIdConnectApplication`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


