# ProvisioningConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | **string** | Defines the method of authentication | 
**Status** | **string** | Provisioning connection status | [default to "DISABLED"]
**Links** | Pointer to [**LinksSelfLifecycleAndAuthorize**](LinksSelfLifecycleAndAuthorize.md) |  | [optional] 

## Methods

### NewProvisioningConnection

`func NewProvisioningConnection(authScheme string, status string, ) *ProvisioningConnection`

NewProvisioningConnection instantiates a new ProvisioningConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConnectionWithDefaults

`func NewProvisioningConnectionWithDefaults() *ProvisioningConnection`

NewProvisioningConnectionWithDefaults instantiates a new ProvisioningConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *ProvisioningConnection) GetAuthScheme() string`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ProvisioningConnection) GetAuthSchemeOk() (*string, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ProvisioningConnection) SetAuthScheme(v string)`

SetAuthScheme sets AuthScheme field to given value.


### GetStatus

`func (o *ProvisioningConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisioningConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisioningConnection) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLinks

`func (o *ProvisioningConnection) GetLinks() LinksSelfLifecycleAndAuthorize`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProvisioningConnection) GetLinksOk() (*LinksSelfLifecycleAndAuthorize, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProvisioningConnection) SetLinks(v LinksSelfLifecycleAndAuthorize)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProvisioningConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


