# NetworkZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asns** | Pointer to **[]string** | Dynamic network zone property: An array of strings that represent an ASN numeric value | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the network zone was created | [optional] [readonly] 
**Gateways** | Pointer to [**[]NetworkZoneAddress**](NetworkZoneAddress.md) | IP network zone property: the IP addresses (range or CIDR form) of this zone. The maximum array length is 150 entries for admin-created IP zones, 1000 entries for IP blocklist zones, and 5000 entries for the default system IP Zone. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the network zone | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the network zone was last modified | [optional] [readonly] 
**Locations** | Pointer to [**[]NetworkZoneLocation**](NetworkZoneLocation.md) | Dynamic network zone property: an array of geolocations of this network zone | [optional] 
**Name** | Pointer to **string** | Unique name for this network zone. Maximum of 128 characters. | [optional] 
**Proxies** | Pointer to [**[]NetworkZoneAddress**](NetworkZoneAddress.md) | IP network zone property: the IP addresses (range or CIDR form) that are allowed to forward a request from gateway addresses These proxies are automatically trusted by Threat Insights, and used to identify the client IP of a request. The maximum array length is 150 entries for admin-created zones and 5000 entries for the default system IP Zone. | [optional] 
**ProxyType** | Pointer to **string** | Dynamic network zone property: the proxy type used | [optional] 
**Status** | Pointer to **string** | Network zone status | [optional] 
**System** | Pointer to **bool** | Indicates if this is a system network zone. For admin-created zones, this is always &#x60;false&#x60;. The system IP Policy Network Zone (&#x60;LegacyIpZone&#x60;) is included by default in your Okta org. Notice that &#x60;system&#x3D;true&#x60; for the &#x60;LegacyIpZone&#x60; object. Admin users can modify the name of this default system Zone and can add up to 5000 gateway or proxy IP entries. | [optional] 
**Type** | Pointer to **string** | The type of network zone | [optional] 
**Usage** | Pointer to **string** | The usage of the network zone | [optional] 
**Links** | Pointer to [**NetworkZoneLinks**](NetworkZoneLinks.md) |  | [optional] 

## Methods

### NewNetworkZone

`func NewNetworkZone() *NetworkZone`

NewNetworkZone instantiates a new NetworkZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkZoneWithDefaults

`func NewNetworkZoneWithDefaults() *NetworkZone`

NewNetworkZoneWithDefaults instantiates a new NetworkZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsns

`func (o *NetworkZone) GetAsns() []string`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *NetworkZone) GetAsnsOk() (*[]string, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *NetworkZone) SetAsns(v []string)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *NetworkZone) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetCreated

`func (o *NetworkZone) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NetworkZone) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NetworkZone) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NetworkZone) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetGateways

`func (o *NetworkZone) GetGateways() []NetworkZoneAddress`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *NetworkZone) GetGatewaysOk() (*[]NetworkZoneAddress, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *NetworkZone) SetGateways(v []NetworkZoneAddress)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *NetworkZone) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetId

`func (o *NetworkZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NetworkZone) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NetworkZone) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NetworkZone) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NetworkZone) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLocations

`func (o *NetworkZone) GetLocations() []NetworkZoneLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *NetworkZone) GetLocationsOk() (*[]NetworkZoneLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *NetworkZone) SetLocations(v []NetworkZoneLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *NetworkZone) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetName

`func (o *NetworkZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProxies

`func (o *NetworkZone) GetProxies() []NetworkZoneAddress`

GetProxies returns the Proxies field if non-nil, zero value otherwise.

### GetProxiesOk

`func (o *NetworkZone) GetProxiesOk() (*[]NetworkZoneAddress, bool)`

GetProxiesOk returns a tuple with the Proxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxies

`func (o *NetworkZone) SetProxies(v []NetworkZoneAddress)`

SetProxies sets Proxies field to given value.

### HasProxies

`func (o *NetworkZone) HasProxies() bool`

HasProxies returns a boolean if a field has been set.

### SetProxiesNil

`func (o *NetworkZone) SetProxiesNil(b bool)`

 SetProxiesNil sets the value for Proxies to be an explicit nil

### UnsetProxies
`func (o *NetworkZone) UnsetProxies()`

UnsetProxies ensures that no value is present for Proxies, not even an explicit nil
### GetProxyType

`func (o *NetworkZone) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *NetworkZone) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *NetworkZone) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *NetworkZone) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *NetworkZone) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *NetworkZone) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *NetworkZone) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *NetworkZone) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *NetworkZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkZone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkZone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsage

`func (o *NetworkZone) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *NetworkZone) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *NetworkZone) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *NetworkZone) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkZone) GetLinks() NetworkZoneLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkZone) GetLinksOk() (*NetworkZoneLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkZone) SetLinks(v NetworkZoneLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkZone) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


