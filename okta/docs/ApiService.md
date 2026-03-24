# ApiService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethod** | Pointer to **string** | Specifies the authentication method used by the API service application when requesting an access token&lt;br&gt; **Note:** Only the &#x60;client_secret_basic&#x60; method is supported | [optional] 
**Scopes** | Pointer to **[]string** | A list of Okta OAuth 2.0 scopes required for the API service app to function | [optional] 
**SetupInstructionsUri** | Pointer to **string** | The URL for the API service integration configuration document | [optional] 

## Methods

### NewApiService

`func NewApiService() *ApiService`

NewApiService instantiates a new ApiService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServiceWithDefaults

`func NewApiServiceWithDefaults() *ApiService`

NewApiServiceWithDefaults instantiates a new ApiService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationMethod

`func (o *ApiService) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *ApiService) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *ApiService) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *ApiService) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetScopes

`func (o *ApiService) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiService) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiService) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApiService) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSetupInstructionsUri

`func (o *ApiService) GetSetupInstructionsUri() string`

GetSetupInstructionsUri returns the SetupInstructionsUri field if non-nil, zero value otherwise.

### GetSetupInstructionsUriOk

`func (o *ApiService) GetSetupInstructionsUriOk() (*string, bool)`

GetSetupInstructionsUriOk returns a tuple with the SetupInstructionsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupInstructionsUri

`func (o *ApiService) SetSetupInstructionsUri(v string)`

SetSetupInstructionsUri sets SetupInstructionsUri field to given value.

### HasSetupInstructionsUri

`func (o *ApiService) HasSetupInstructionsUri() bool`

HasSetupInstructionsUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


