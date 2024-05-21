# GetDefaultProvisioningConnectionForApplication200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | **string** | Defines the method of authentication | 
**Status** | **string** | Provisioning connection status | [default to "DISABLED"]
**Links** | Pointer to [**LinksSelfLifecycleAndAuthorize**](LinksSelfLifecycleAndAuthorize.md) |  | [optional] 
**Profile** | Pointer to [**ProvisioningConnectionProfileUnknown**](ProvisioningConnectionProfileUnknown.md) |  | [optional] 

## Methods

### NewGetDefaultProvisioningConnectionForApplication200Response

`func NewGetDefaultProvisioningConnectionForApplication200Response(authScheme string, status string, ) *GetDefaultProvisioningConnectionForApplication200Response`

NewGetDefaultProvisioningConnectionForApplication200Response instantiates a new GetDefaultProvisioningConnectionForApplication200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDefaultProvisioningConnectionForApplication200ResponseWithDefaults

`func NewGetDefaultProvisioningConnectionForApplication200ResponseWithDefaults() *GetDefaultProvisioningConnectionForApplication200Response`

NewGetDefaultProvisioningConnectionForApplication200ResponseWithDefaults instantiates a new GetDefaultProvisioningConnectionForApplication200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *GetDefaultProvisioningConnectionForApplication200Response) GetAuthScheme() string`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *GetDefaultProvisioningConnectionForApplication200Response) GetAuthSchemeOk() (*string, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *GetDefaultProvisioningConnectionForApplication200Response) SetAuthScheme(v string)`

SetAuthScheme sets AuthScheme field to given value.


### GetStatus

`func (o *GetDefaultProvisioningConnectionForApplication200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDefaultProvisioningConnectionForApplication200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDefaultProvisioningConnectionForApplication200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLinks

`func (o *GetDefaultProvisioningConnectionForApplication200Response) GetLinks() LinksSelfLifecycleAndAuthorize`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDefaultProvisioningConnectionForApplication200Response) GetLinksOk() (*LinksSelfLifecycleAndAuthorize, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDefaultProvisioningConnectionForApplication200Response) SetLinks(v LinksSelfLifecycleAndAuthorize)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDefaultProvisioningConnectionForApplication200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetProfile

`func (o *GetDefaultProvisioningConnectionForApplication200Response) GetProfile() ProvisioningConnectionProfileUnknown`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetDefaultProvisioningConnectionForApplication200Response) GetProfileOk() (*ProvisioningConnectionProfileUnknown, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetDefaultProvisioningConnectionForApplication200Response) SetProfile(v ProvisioningConnectionProfileUnknown)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GetDefaultProvisioningConnectionForApplication200Response) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


