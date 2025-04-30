# Saml11Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ApplicationCredentials**](ApplicationCredentials.md) |  | [optional] 
**Name** | **string** | The key name for the SAML 1.1 app definition. You can&#39;t create a custom SAML 1.1 app integration instance. Only existing OIN SAML 1.1 app integrations are supported. | 
**Settings** | Pointer to [**Saml11ApplicationSettings**](Saml11ApplicationSettings.md) |  | [optional] 

## Methods

### NewSaml11Application

`func NewSaml11Application(name string, ) *Saml11Application`

NewSaml11Application instantiates a new Saml11Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaml11ApplicationWithDefaults

`func NewSaml11ApplicationWithDefaults() *Saml11Application`

NewSaml11ApplicationWithDefaults instantiates a new Saml11Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *Saml11Application) GetCredentials() ApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Saml11Application) GetCredentialsOk() (*ApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Saml11Application) SetCredentials(v ApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Saml11Application) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *Saml11Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Saml11Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Saml11Application) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *Saml11Application) GetSettings() Saml11ApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Saml11Application) GetSettingsOk() (*Saml11ApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Saml11Application) SetSettings(v Saml11ApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Saml11Application) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


