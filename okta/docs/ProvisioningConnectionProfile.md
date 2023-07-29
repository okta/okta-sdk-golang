# ProvisioningConnectionProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | Pointer to [**ProvisioningConnectionAuthScheme**](ProvisioningConnectionAuthScheme.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewProvisioningConnectionProfile

`func NewProvisioningConnectionProfile() *ProvisioningConnectionProfile`

NewProvisioningConnectionProfile instantiates a new ProvisioningConnectionProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConnectionProfileWithDefaults

`func NewProvisioningConnectionProfileWithDefaults() *ProvisioningConnectionProfile`

NewProvisioningConnectionProfileWithDefaults instantiates a new ProvisioningConnectionProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *ProvisioningConnectionProfile) GetAuthScheme() ProvisioningConnectionAuthScheme`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ProvisioningConnectionProfile) GetAuthSchemeOk() (*ProvisioningConnectionAuthScheme, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ProvisioningConnectionProfile) SetAuthScheme(v ProvisioningConnectionAuthScheme)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *ProvisioningConnectionProfile) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### GetToken

`func (o *ProvisioningConnectionProfile) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProvisioningConnectionProfile) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProvisioningConnectionProfile) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProvisioningConnectionProfile) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


