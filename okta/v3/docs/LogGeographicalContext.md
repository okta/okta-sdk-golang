# LogGeographicalContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** |  | [optional] [readonly] 
**Country** | Pointer to **string** |  | [optional] [readonly] 
**Geolocation** | Pointer to [**LogGeolocation**](LogGeolocation.md) |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLogGeographicalContext

`func NewLogGeographicalContext() *LogGeographicalContext`

NewLogGeographicalContext instantiates a new LogGeographicalContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogGeographicalContextWithDefaults

`func NewLogGeographicalContextWithDefaults() *LogGeographicalContext`

NewLogGeographicalContextWithDefaults instantiates a new LogGeographicalContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *LogGeographicalContext) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LogGeographicalContext) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LogGeographicalContext) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *LogGeographicalContext) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *LogGeographicalContext) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LogGeographicalContext) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LogGeographicalContext) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LogGeographicalContext) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetGeolocation

`func (o *LogGeographicalContext) GetGeolocation() LogGeolocation`

GetGeolocation returns the Geolocation field if non-nil, zero value otherwise.

### GetGeolocationOk

`func (o *LogGeographicalContext) GetGeolocationOk() (*LogGeolocation, bool)`

GetGeolocationOk returns a tuple with the Geolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocation

`func (o *LogGeographicalContext) SetGeolocation(v LogGeolocation)`

SetGeolocation sets Geolocation field to given value.

### HasGeolocation

`func (o *LogGeographicalContext) HasGeolocation() bool`

HasGeolocation returns a boolean if a field has been set.

### GetPostalCode

`func (o *LogGeographicalContext) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *LogGeographicalContext) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *LogGeographicalContext) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *LogGeographicalContext) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *LogGeographicalContext) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LogGeographicalContext) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LogGeographicalContext) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LogGeographicalContext) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


