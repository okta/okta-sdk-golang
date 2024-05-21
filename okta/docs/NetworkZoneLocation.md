# NetworkZoneLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Format of the country value: length 2 [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code. Do not use continent codes as they are treated as generic codes for undesignated countries. | [optional] 
**Region** | Pointer to **string** | Format of the region value (optional): region code [ISO-3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) appended to country code (&#x60;countryCode-regionCode&#x60;), or &#x60;null&#x60; if empty. Do not use continent codes as they are treated as generic codes for undesignated regions. | [optional] 

## Methods

### NewNetworkZoneLocation

`func NewNetworkZoneLocation() *NetworkZoneLocation`

NewNetworkZoneLocation instantiates a new NetworkZoneLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkZoneLocationWithDefaults

`func NewNetworkZoneLocationWithDefaults() *NetworkZoneLocation`

NewNetworkZoneLocationWithDefaults instantiates a new NetworkZoneLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *NetworkZoneLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *NetworkZoneLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *NetworkZoneLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *NetworkZoneLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRegion

`func (o *NetworkZoneLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkZoneLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkZoneLocation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NetworkZoneLocation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


