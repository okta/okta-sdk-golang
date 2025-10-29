# OAuth2ClientSecretRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | Pointer to **string** | The OAuth 2.0 client secret string | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 Client Secret | [optional] 

## Methods

### NewOAuth2ClientSecretRequestBody

`func NewOAuth2ClientSecretRequestBody() *OAuth2ClientSecretRequestBody`

NewOAuth2ClientSecretRequestBody instantiates a new OAuth2ClientSecretRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientSecretRequestBodyWithDefaults

`func NewOAuth2ClientSecretRequestBodyWithDefaults() *OAuth2ClientSecretRequestBody`

NewOAuth2ClientSecretRequestBodyWithDefaults instantiates a new OAuth2ClientSecretRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *OAuth2ClientSecretRequestBody) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2ClientSecretRequestBody) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2ClientSecretRequestBody) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuth2ClientSecretRequestBody) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetStatus

`func (o *OAuth2ClientSecretRequestBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ClientSecretRequestBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ClientSecretRequestBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ClientSecretRequestBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


