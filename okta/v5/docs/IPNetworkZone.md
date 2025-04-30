# IPNetworkZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateways** | Pointer to [**[]NetworkZoneAddress**](NetworkZoneAddress.md) | The IP addresses (range or CIDR form) for an IP Network Zone. The maximum array length is 150 entries for admin-created IP zones, 1000 entries for IP blocklist zones, and 5000 entries for the default system IP Zone. | [optional] 
**Proxies** | Pointer to [**[]NetworkZoneAddress**](NetworkZoneAddress.md) | The IP addresses (range or CIDR form) that are allowed to forward a request from gateway addresses for an IP Network Zone. These proxies are automatically trusted by Threat Insights and used to identify the client IP of a request. The maximum array length is 150 entries for admin-created zones and 5000 entries for the default system IP Zone. | [optional] 

## Methods

### NewIPNetworkZone

`func NewIPNetworkZone() *IPNetworkZone`

NewIPNetworkZone instantiates a new IPNetworkZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPNetworkZoneWithDefaults

`func NewIPNetworkZoneWithDefaults() *IPNetworkZone`

NewIPNetworkZoneWithDefaults instantiates a new IPNetworkZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateways

`func (o *IPNetworkZone) GetGateways() []NetworkZoneAddress`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *IPNetworkZone) GetGatewaysOk() (*[]NetworkZoneAddress, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *IPNetworkZone) SetGateways(v []NetworkZoneAddress)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *IPNetworkZone) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetProxies

`func (o *IPNetworkZone) GetProxies() []NetworkZoneAddress`

GetProxies returns the Proxies field if non-nil, zero value otherwise.

### GetProxiesOk

`func (o *IPNetworkZone) GetProxiesOk() (*[]NetworkZoneAddress, bool)`

GetProxiesOk returns a tuple with the Proxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxies

`func (o *IPNetworkZone) SetProxies(v []NetworkZoneAddress)`

SetProxies sets Proxies field to given value.

### HasProxies

`func (o *IPNetworkZone) HasProxies() bool`

HasProxies returns a boolean if a field has been set.

### SetProxiesNil

`func (o *IPNetworkZone) SetProxiesNil(b bool)`

 SetProxiesNil sets the value for Proxies to be an explicit nil

### UnsetProxies
`func (o *IPNetworkZone) UnsetProxies()`

UnsetProxies ensures that no value is present for Proxies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


