# LogClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** |  | [optional] [readonly] 
**GeographicalContext** | Pointer to [**LogGeographicalContext**](LogGeographicalContext.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IpAddress** | Pointer to **string** |  | [optional] [readonly] 
**UserAgent** | Pointer to [**LogUserAgent**](LogUserAgent.md) |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLogClient

`func NewLogClient() *LogClient`

NewLogClient instantiates a new LogClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogClientWithDefaults

`func NewLogClientWithDefaults() *LogClient`

NewLogClientWithDefaults instantiates a new LogClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *LogClient) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *LogClient) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *LogClient) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *LogClient) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetGeographicalContext

`func (o *LogClient) GetGeographicalContext() LogGeographicalContext`

GetGeographicalContext returns the GeographicalContext field if non-nil, zero value otherwise.

### GetGeographicalContextOk

`func (o *LogClient) GetGeographicalContextOk() (*LogGeographicalContext, bool)`

GetGeographicalContextOk returns a tuple with the GeographicalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicalContext

`func (o *LogClient) SetGeographicalContext(v LogGeographicalContext)`

SetGeographicalContext sets GeographicalContext field to given value.

### HasGeographicalContext

`func (o *LogClient) HasGeographicalContext() bool`

HasGeographicalContext returns a boolean if a field has been set.

### GetId

`func (o *LogClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogClient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *LogClient) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LogClient) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LogClient) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *LogClient) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetUserAgent

`func (o *LogClient) GetUserAgent() LogUserAgent`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *LogClient) GetUserAgentOk() (*LogUserAgent, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *LogClient) SetUserAgent(v LogUserAgent)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *LogClient) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetZone

`func (o *LogClient) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *LogClient) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *LogClient) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *LogClient) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


