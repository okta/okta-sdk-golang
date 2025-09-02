# DynamicNetworkZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asns** | Pointer to **[]string** | An array of ASNs for a Network Zone | [optional] 
**ProxyType** | Pointer to **string** | The proxy type used for a Dynamic Network Zone | [optional] 
**Locations** | Pointer to [**[]NetworkZoneLocation**](NetworkZoneLocation.md) | An array of geolocations for a Dynamic Network Zone | [optional] 

## Methods

### NewDynamicNetworkZone

`func NewDynamicNetworkZone() *DynamicNetworkZone`

NewDynamicNetworkZone instantiates a new DynamicNetworkZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicNetworkZoneWithDefaults

`func NewDynamicNetworkZoneWithDefaults() *DynamicNetworkZone`

NewDynamicNetworkZoneWithDefaults instantiates a new DynamicNetworkZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsns

`func (o *DynamicNetworkZone) GetAsns() []string`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *DynamicNetworkZone) GetAsnsOk() (*[]string, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *DynamicNetworkZone) SetAsns(v []string)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *DynamicNetworkZone) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetProxyType

`func (o *DynamicNetworkZone) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *DynamicNetworkZone) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *DynamicNetworkZone) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *DynamicNetworkZone) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetLocations

`func (o *DynamicNetworkZone) GetLocations() []NetworkZoneLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *DynamicNetworkZone) GetLocationsOk() (*[]NetworkZoneLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *DynamicNetworkZone) SetLocations(v []NetworkZoneLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *DynamicNetworkZone) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *DynamicNetworkZone) SetLocationsNil(b bool)`

 SetLocationsNil sets the value for Locations to be an explicit nil

### UnsetLocations
`func (o *DynamicNetworkZone) UnsetLocations()`

UnsetLocations ensures that no value is present for Locations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


