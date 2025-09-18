# OpenIdConnectApplicationNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | **string** | The connection type of the network. Can be &#x60;ANYWHERE&#x60; or &#x60;ZONE&#x60;.  | 
**Exclude** | Pointer to **[]string** | If &#x60;ZONE&#x60; is specified as a connection, then specify the excluded IP network zones here. Value can be \&quot;ALL_IP_ZONES\&quot; or an array of zone IDs. | [optional] 
**Include** | Pointer to **[]string** | If &#x60;ZONE&#x60; is specified as a connection, then specify the included IP network zones here. Value can be \&quot;ALL_IP_ZONES\&quot; or an array of zone IDs. | [optional] 

## Methods

### NewOpenIdConnectApplicationNetwork

`func NewOpenIdConnectApplicationNetwork(connection string, ) *OpenIdConnectApplicationNetwork`

NewOpenIdConnectApplicationNetwork instantiates a new OpenIdConnectApplicationNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectApplicationNetworkWithDefaults

`func NewOpenIdConnectApplicationNetworkWithDefaults() *OpenIdConnectApplicationNetwork`

NewOpenIdConnectApplicationNetworkWithDefaults instantiates a new OpenIdConnectApplicationNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *OpenIdConnectApplicationNetwork) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *OpenIdConnectApplicationNetwork) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *OpenIdConnectApplicationNetwork) SetConnection(v string)`

SetConnection sets Connection field to given value.


### GetExclude

`func (o *OpenIdConnectApplicationNetwork) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *OpenIdConnectApplicationNetwork) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *OpenIdConnectApplicationNetwork) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *OpenIdConnectApplicationNetwork) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetInclude

`func (o *OpenIdConnectApplicationNetwork) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *OpenIdConnectApplicationNetwork) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *OpenIdConnectApplicationNetwork) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *OpenIdConnectApplicationNetwork) HasInclude() bool`

HasInclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


