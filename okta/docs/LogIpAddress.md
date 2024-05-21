# LogIpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeographicalContext** | Pointer to [**LogGeographicalContext**](LogGeographicalContext.md) |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] [readonly] 
**Source** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLogIpAddress

`func NewLogIpAddress() *LogIpAddress`

NewLogIpAddress instantiates a new LogIpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogIpAddressWithDefaults

`func NewLogIpAddressWithDefaults() *LogIpAddress`

NewLogIpAddressWithDefaults instantiates a new LogIpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeographicalContext

`func (o *LogIpAddress) GetGeographicalContext() LogGeographicalContext`

GetGeographicalContext returns the GeographicalContext field if non-nil, zero value otherwise.

### GetGeographicalContextOk

`func (o *LogIpAddress) GetGeographicalContextOk() (*LogGeographicalContext, bool)`

GetGeographicalContextOk returns a tuple with the GeographicalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicalContext

`func (o *LogIpAddress) SetGeographicalContext(v LogGeographicalContext)`

SetGeographicalContext sets GeographicalContext field to given value.

### HasGeographicalContext

`func (o *LogIpAddress) HasGeographicalContext() bool`

HasGeographicalContext returns a boolean if a field has been set.

### GetIp

`func (o *LogIpAddress) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LogIpAddress) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LogIpAddress) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LogIpAddress) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetSource

`func (o *LogIpAddress) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LogIpAddress) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LogIpAddress) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LogIpAddress) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *LogIpAddress) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LogIpAddress) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LogIpAddress) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LogIpAddress) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


