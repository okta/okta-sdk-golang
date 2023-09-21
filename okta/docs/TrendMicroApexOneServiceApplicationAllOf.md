# TrendMicroApexOneServiceApplicationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to "trendmicroapexoneservice"]
**SignOnMode** | Pointer to **interface{}** |  | [optional] [default to SAML_2_0]
**Settings** | [**TrendMicroApexOneServiceApplicationSettings**](TrendMicroApexOneServiceApplicationSettings.md) |  | 

## Methods

### NewTrendMicroApexOneServiceApplicationAllOf

`func NewTrendMicroApexOneServiceApplicationAllOf(name string, settings TrendMicroApexOneServiceApplicationSettings, ) *TrendMicroApexOneServiceApplicationAllOf`

NewTrendMicroApexOneServiceApplicationAllOf instantiates a new TrendMicroApexOneServiceApplicationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendMicroApexOneServiceApplicationAllOfWithDefaults

`func NewTrendMicroApexOneServiceApplicationAllOfWithDefaults() *TrendMicroApexOneServiceApplicationAllOf`

NewTrendMicroApexOneServiceApplicationAllOfWithDefaults instantiates a new TrendMicroApexOneServiceApplicationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TrendMicroApexOneServiceApplicationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrendMicroApexOneServiceApplicationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrendMicroApexOneServiceApplicationAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetSignOnMode

`func (o *TrendMicroApexOneServiceApplicationAllOf) GetSignOnMode() interface{}`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *TrendMicroApexOneServiceApplicationAllOf) GetSignOnModeOk() (*interface{}, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *TrendMicroApexOneServiceApplicationAllOf) SetSignOnMode(v interface{})`

SetSignOnMode sets SignOnMode field to given value.

### HasSignOnMode

`func (o *TrendMicroApexOneServiceApplicationAllOf) HasSignOnMode() bool`

HasSignOnMode returns a boolean if a field has been set.

### SetSignOnModeNil

`func (o *TrendMicroApexOneServiceApplicationAllOf) SetSignOnModeNil(b bool)`

 SetSignOnModeNil sets the value for SignOnMode to be an explicit nil

### UnsetSignOnMode
`func (o *TrendMicroApexOneServiceApplicationAllOf) UnsetSignOnMode()`

UnsetSignOnMode ensures that no value is present for SignOnMode, not even an explicit nil
### GetSettings

`func (o *TrendMicroApexOneServiceApplicationAllOf) GetSettings() TrendMicroApexOneServiceApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *TrendMicroApexOneServiceApplicationAllOf) GetSettingsOk() (*TrendMicroApexOneServiceApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *TrendMicroApexOneServiceApplicationAllOf) SetSettings(v TrendMicroApexOneServiceApplicationSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


