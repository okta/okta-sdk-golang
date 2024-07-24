# ProvisioningConnectionTokenRequestProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | **string** | A token is used to authenticate with the app. This property is only returned for the &#x60;TOKEN&#x60; authentication scheme. | 
**Token** | Pointer to **string** | Token used to authenticate with the app | [optional] 

## Methods

### NewProvisioningConnectionTokenRequestProfile

`func NewProvisioningConnectionTokenRequestProfile(authScheme string, ) *ProvisioningConnectionTokenRequestProfile`

NewProvisioningConnectionTokenRequestProfile instantiates a new ProvisioningConnectionTokenRequestProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConnectionTokenRequestProfileWithDefaults

`func NewProvisioningConnectionTokenRequestProfileWithDefaults() *ProvisioningConnectionTokenRequestProfile`

NewProvisioningConnectionTokenRequestProfileWithDefaults instantiates a new ProvisioningConnectionTokenRequestProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *ProvisioningConnectionTokenRequestProfile) GetAuthScheme() string`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ProvisioningConnectionTokenRequestProfile) GetAuthSchemeOk() (*string, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ProvisioningConnectionTokenRequestProfile) SetAuthScheme(v string)`

SetAuthScheme sets AuthScheme field to given value.


### GetToken

`func (o *ProvisioningConnectionTokenRequestProfile) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProvisioningConnectionTokenRequestProfile) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProvisioningConnectionTokenRequestProfile) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProvisioningConnectionTokenRequestProfile) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


