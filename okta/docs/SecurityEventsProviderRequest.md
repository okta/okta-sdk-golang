# SecurityEventsProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Security Events Provider instance | 
**Settings** | [**SecurityEventsProviderRequestSettings**](SecurityEventsProviderRequestSettings.md) |  | 
**Type** | **string** | The application type of the Security Events Provider | 

## Methods

### NewSecurityEventsProviderRequest

`func NewSecurityEventsProviderRequest(name string, settings SecurityEventsProviderRequestSettings, type_ string, ) *SecurityEventsProviderRequest`

NewSecurityEventsProviderRequest instantiates a new SecurityEventsProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventsProviderRequestWithDefaults

`func NewSecurityEventsProviderRequestWithDefaults() *SecurityEventsProviderRequest`

NewSecurityEventsProviderRequestWithDefaults instantiates a new SecurityEventsProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityEventsProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityEventsProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityEventsProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *SecurityEventsProviderRequest) GetSettings() SecurityEventsProviderRequestSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SecurityEventsProviderRequest) GetSettingsOk() (*SecurityEventsProviderRequestSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SecurityEventsProviderRequest) SetSettings(v SecurityEventsProviderRequestSettings)`

SetSettings sets Settings field to given value.


### GetType

`func (o *SecurityEventsProviderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityEventsProviderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityEventsProviderRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


