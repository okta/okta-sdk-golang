# ProvisioningConnectionProfileOauth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | [**ProvisioningConnectionAuthScheme**](ProvisioningConnectionAuthScheme.md) |  | 
**Token** | Pointer to **string** |  | [optional] 
**ClientId** | **string** | Unique client identifier for the OAuth 2.0 service app from the target org | 

## Methods

### NewProvisioningConnectionProfileOauth

`func NewProvisioningConnectionProfileOauth(authScheme ProvisioningConnectionAuthScheme, clientId string, ) *ProvisioningConnectionProfileOauth`

NewProvisioningConnectionProfileOauth instantiates a new ProvisioningConnectionProfileOauth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConnectionProfileOauthWithDefaults

`func NewProvisioningConnectionProfileOauthWithDefaults() *ProvisioningConnectionProfileOauth`

NewProvisioningConnectionProfileOauthWithDefaults instantiates a new ProvisioningConnectionProfileOauth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *ProvisioningConnectionProfileOauth) GetAuthScheme() ProvisioningConnectionAuthScheme`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ProvisioningConnectionProfileOauth) GetAuthSchemeOk() (*ProvisioningConnectionAuthScheme, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ProvisioningConnectionProfileOauth) SetAuthScheme(v ProvisioningConnectionAuthScheme)`

SetAuthScheme sets AuthScheme field to given value.


### GetToken

`func (o *ProvisioningConnectionProfileOauth) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProvisioningConnectionProfileOauth) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProvisioningConnectionProfileOauth) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProvisioningConnectionProfileOauth) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetClientId

`func (o *ProvisioningConnectionProfileOauth) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ProvisioningConnectionProfileOauth) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ProvisioningConnectionProfileOauth) SetClientId(v string)`

SetClientId sets ClientId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


