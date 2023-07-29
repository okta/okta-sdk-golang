# SamlApplicationSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**SamlApplicationSettingsApplication**](SamlApplicationSettingsApplication.md) |  | [optional] 
**SignOn** | Pointer to [**SamlApplicationSettingsSignOn**](SamlApplicationSettingsSignOn.md) |  | [optional] 

## Methods

### NewSamlApplicationSettingsAllOf

`func NewSamlApplicationSettingsAllOf() *SamlApplicationSettingsAllOf`

NewSamlApplicationSettingsAllOf instantiates a new SamlApplicationSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlApplicationSettingsAllOfWithDefaults

`func NewSamlApplicationSettingsAllOfWithDefaults() *SamlApplicationSettingsAllOf`

NewSamlApplicationSettingsAllOfWithDefaults instantiates a new SamlApplicationSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *SamlApplicationSettingsAllOf) GetApp() SamlApplicationSettingsApplication`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *SamlApplicationSettingsAllOf) GetAppOk() (*SamlApplicationSettingsApplication, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *SamlApplicationSettingsAllOf) SetApp(v SamlApplicationSettingsApplication)`

SetApp sets App field to given value.

### HasApp

`func (o *SamlApplicationSettingsAllOf) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetSignOn

`func (o *SamlApplicationSettingsAllOf) GetSignOn() SamlApplicationSettingsSignOn`

GetSignOn returns the SignOn field if non-nil, zero value otherwise.

### GetSignOnOk

`func (o *SamlApplicationSettingsAllOf) GetSignOnOk() (*SamlApplicationSettingsSignOn, bool)`

GetSignOnOk returns a tuple with the SignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOn

`func (o *SamlApplicationSettingsAllOf) SetSignOn(v SamlApplicationSettingsSignOn)`

SetSignOn sets SignOn field to given value.

### HasSignOn

`func (o *SamlApplicationSettingsAllOf) HasSignOn() bool`

HasSignOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


