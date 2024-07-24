# ProvisioningConnectionProfileOauth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | **string** | OAuth 2.0 is used to authenticate with the app. | 
**ClientId** | Pointer to **string** |  | [optional] 

## Methods

### NewProvisioningConnectionProfileOauth

`func NewProvisioningConnectionProfileOauth(authScheme string, ) *ProvisioningConnectionProfileOauth`

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

`func (o *ProvisioningConnectionProfileOauth) GetAuthScheme() string`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ProvisioningConnectionProfileOauth) GetAuthSchemeOk() (*string, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ProvisioningConnectionProfileOauth) SetAuthScheme(v string)`

SetAuthScheme sets AuthScheme field to given value.


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

### HasClientId

`func (o *ProvisioningConnectionProfileOauth) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


