# ProvisioningConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | Pointer to [**ProvisioningConnectionAuthScheme**](ProvisioningConnectionAuthScheme.md) |  | [optional] 
**Status** | Pointer to [**ProvisioningConnectionStatus**](ProvisioningConnectionStatus.md) |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewProvisioningConnection

`func NewProvisioningConnection() *ProvisioningConnection`

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

`func (o *ProvisioningConnection) GetAuthScheme() ProvisioningConnectionAuthScheme`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ProvisioningConnection) GetAuthSchemeOk() (*ProvisioningConnectionAuthScheme, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ProvisioningConnection) SetAuthScheme(v ProvisioningConnectionAuthScheme)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *ProvisioningConnection) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### GetStatus

`func (o *ProvisioningConnection) GetStatus() ProvisioningConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisioningConnection) GetStatusOk() (*ProvisioningConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisioningConnection) SetStatus(v ProvisioningConnectionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProvisioningConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *ProvisioningConnection) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProvisioningConnection) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProvisioningConnection) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProvisioningConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


