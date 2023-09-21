# SlackApplicationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to "slack"]
**SignOnMode** | Pointer to **interface{}** |  | [optional] [default to SAML_2_0]
**Settings** | [**SlackApplicationSettings**](SlackApplicationSettings.md) |  | 

## Methods

### NewSlackApplicationAllOf

`func NewSlackApplicationAllOf(name string, settings SlackApplicationSettings, ) *SlackApplicationAllOf`

NewSlackApplicationAllOf instantiates a new SlackApplicationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackApplicationAllOfWithDefaults

`func NewSlackApplicationAllOfWithDefaults() *SlackApplicationAllOf`

NewSlackApplicationAllOfWithDefaults instantiates a new SlackApplicationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SlackApplicationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlackApplicationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlackApplicationAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetSignOnMode

`func (o *SlackApplicationAllOf) GetSignOnMode() interface{}`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *SlackApplicationAllOf) GetSignOnModeOk() (*interface{}, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *SlackApplicationAllOf) SetSignOnMode(v interface{})`

SetSignOnMode sets SignOnMode field to given value.

### HasSignOnMode

`func (o *SlackApplicationAllOf) HasSignOnMode() bool`

HasSignOnMode returns a boolean if a field has been set.

### SetSignOnModeNil

`func (o *SlackApplicationAllOf) SetSignOnModeNil(b bool)`

 SetSignOnModeNil sets the value for SignOnMode to be an explicit nil

### UnsetSignOnMode
`func (o *SlackApplicationAllOf) UnsetSignOnMode()`

UnsetSignOnMode ensures that no value is present for SignOnMode, not even an explicit nil
### GetSettings

`func (o *SlackApplicationAllOf) GetSettings() SlackApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SlackApplicationAllOf) GetSettingsOk() (*SlackApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SlackApplicationAllOf) SetSettings(v SlackApplicationSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


