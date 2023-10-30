# ProvisioningConnectionProfileToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | [**ProvisioningConnectionAuthScheme**](ProvisioningConnectionAuthScheme.md) |  | 
**Token** | **string** | Token used to authenticate with the app | 

## Methods

### NewProvisioningConnectionProfileToken

`func NewProvisioningConnectionProfileToken(authScheme ProvisioningConnectionAuthScheme, token string, ) *ProvisioningConnectionProfileToken`

NewProvisioningConnectionProfileToken instantiates a new ProvisioningConnectionProfileToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConnectionProfileTokenWithDefaults

`func NewProvisioningConnectionProfileTokenWithDefaults() *ProvisioningConnectionProfileToken`

NewProvisioningConnectionProfileTokenWithDefaults instantiates a new ProvisioningConnectionProfileToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *ProvisioningConnectionProfileToken) GetAuthScheme() ProvisioningConnectionAuthScheme`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ProvisioningConnectionProfileToken) GetAuthSchemeOk() (*ProvisioningConnectionAuthScheme, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ProvisioningConnectionProfileToken) SetAuthScheme(v ProvisioningConnectionAuthScheme)`

SetAuthScheme sets AuthScheme field to given value.


### GetToken

`func (o *ProvisioningConnectionProfileToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProvisioningConnectionProfileToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProvisioningConnectionProfileToken) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


