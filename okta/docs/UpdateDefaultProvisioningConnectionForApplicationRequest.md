# UpdateDefaultProvisioningConnectionForApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | Only used for the Zscaler 2.0 (&#x60;zscalerbyz&#x60;) app. The base URL for the Zscaler 2.0 target app, which also contains the Zscaler ID. | [optional] 
**Profile** | [**ProvisioningConnectionOauthRequestProfile**](ProvisioningConnectionOauthRequestProfile.md) |  | 

## Methods

### NewUpdateDefaultProvisioningConnectionForApplicationRequest

`func NewUpdateDefaultProvisioningConnectionForApplicationRequest(profile ProvisioningConnectionOauthRequestProfile, ) *UpdateDefaultProvisioningConnectionForApplicationRequest`

NewUpdateDefaultProvisioningConnectionForApplicationRequest instantiates a new UpdateDefaultProvisioningConnectionForApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDefaultProvisioningConnectionForApplicationRequestWithDefaults

`func NewUpdateDefaultProvisioningConnectionForApplicationRequestWithDefaults() *UpdateDefaultProvisioningConnectionForApplicationRequest`

NewUpdateDefaultProvisioningConnectionForApplicationRequestWithDefaults instantiates a new UpdateDefaultProvisioningConnectionForApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *UpdateDefaultProvisioningConnectionForApplicationRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *UpdateDefaultProvisioningConnectionForApplicationRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *UpdateDefaultProvisioningConnectionForApplicationRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *UpdateDefaultProvisioningConnectionForApplicationRequest) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetProfile

`func (o *UpdateDefaultProvisioningConnectionForApplicationRequest) GetProfile() ProvisioningConnectionOauthRequestProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UpdateDefaultProvisioningConnectionForApplicationRequest) GetProfileOk() (*ProvisioningConnectionOauthRequestProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UpdateDefaultProvisioningConnectionForApplicationRequest) SetProfile(v ProvisioningConnectionOauthRequestProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


