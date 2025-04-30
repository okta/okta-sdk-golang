# ProvisioningConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | Pointer to **string** | A token is used to authenticate with the app. This property is only returned for the &#x60;TOKEN&#x60; authentication scheme. | [optional] 
**BaseUrl** | Pointer to **string** | Base URL | [optional] 
**Profile** | [**ProvisioningConnectionResponseProfile**](ProvisioningConnectionResponseProfile.md) |  | 
**Status** | **string** | Provisioning Connection status | [default to "DISABLED"]
**Links** | Pointer to [**LinksSelfLifecycleAndAuthorize**](LinksSelfLifecycleAndAuthorize.md) |  | [optional] 

## Methods

### NewProvisioningConnectionResponse

`func NewProvisioningConnectionResponse(profile ProvisioningConnectionResponseProfile, status string, ) *ProvisioningConnectionResponse`

NewProvisioningConnectionResponse instantiates a new ProvisioningConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConnectionResponseWithDefaults

`func NewProvisioningConnectionResponseWithDefaults() *ProvisioningConnectionResponse`

NewProvisioningConnectionResponseWithDefaults instantiates a new ProvisioningConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *ProvisioningConnectionResponse) GetAuthScheme() string`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ProvisioningConnectionResponse) GetAuthSchemeOk() (*string, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ProvisioningConnectionResponse) SetAuthScheme(v string)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *ProvisioningConnectionResponse) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### GetBaseUrl

`func (o *ProvisioningConnectionResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ProvisioningConnectionResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ProvisioningConnectionResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ProvisioningConnectionResponse) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetProfile

`func (o *ProvisioningConnectionResponse) GetProfile() ProvisioningConnectionResponseProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ProvisioningConnectionResponse) GetProfileOk() (*ProvisioningConnectionResponseProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ProvisioningConnectionResponse) SetProfile(v ProvisioningConnectionResponseProfile)`

SetProfile sets Profile field to given value.


### GetStatus

`func (o *ProvisioningConnectionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisioningConnectionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisioningConnectionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLinks

`func (o *ProvisioningConnectionResponse) GetLinks() LinksSelfLifecycleAndAuthorize`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProvisioningConnectionResponse) GetLinksOk() (*LinksSelfLifecycleAndAuthorize, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProvisioningConnectionResponse) SetLinks(v LinksSelfLifecycleAndAuthorize)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProvisioningConnectionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


