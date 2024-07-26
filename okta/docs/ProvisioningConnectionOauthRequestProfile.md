# ProvisioningConnectionOauthRequestProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | **string** | OAuth 2.0 is used to authenticate with the app. | 
**ClientId** | Pointer to **string** | Only used for the Okta Org2Org (&#x60;okta_org2org&#x60;) app. The unique client identifier for the OAuth 2.0 service app from the target org. | [optional] 
**Settings** | Pointer to [**Office365ProvisioningSettings**](Office365ProvisioningSettings.md) |  | [optional] 

## Methods

### NewProvisioningConnectionOauthRequestProfile

`func NewProvisioningConnectionOauthRequestProfile(authScheme string, ) *ProvisioningConnectionOauthRequestProfile`

NewProvisioningConnectionOauthRequestProfile instantiates a new ProvisioningConnectionOauthRequestProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConnectionOauthRequestProfileWithDefaults

`func NewProvisioningConnectionOauthRequestProfileWithDefaults() *ProvisioningConnectionOauthRequestProfile`

NewProvisioningConnectionOauthRequestProfileWithDefaults instantiates a new ProvisioningConnectionOauthRequestProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *ProvisioningConnectionOauthRequestProfile) GetAuthScheme() string`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ProvisioningConnectionOauthRequestProfile) GetAuthSchemeOk() (*string, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ProvisioningConnectionOauthRequestProfile) SetAuthScheme(v string)`

SetAuthScheme sets AuthScheme field to given value.


### GetClientId

`func (o *ProvisioningConnectionOauthRequestProfile) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ProvisioningConnectionOauthRequestProfile) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ProvisioningConnectionOauthRequestProfile) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ProvisioningConnectionOauthRequestProfile) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSettings

`func (o *ProvisioningConnectionOauthRequestProfile) GetSettings() Office365ProvisioningSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ProvisioningConnectionOauthRequestProfile) GetSettingsOk() (*Office365ProvisioningSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ProvisioningConnectionOauthRequestProfile) SetSettings(v Office365ProvisioningSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ProvisioningConnectionOauthRequestProfile) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


