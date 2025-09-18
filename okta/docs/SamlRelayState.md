# SamlRelayState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | The format used to generate the &#x60;relayState&#x60; in the SAML request. The &#x60;FROM_URL&#x60; format is used if this value is null. | [optional] 

## Methods

### NewSamlRelayState

`func NewSamlRelayState() *SamlRelayState`

NewSamlRelayState instantiates a new SamlRelayState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlRelayStateWithDefaults

`func NewSamlRelayStateWithDefaults() *SamlRelayState`

NewSamlRelayStateWithDefaults instantiates a new SamlRelayState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *SamlRelayState) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SamlRelayState) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SamlRelayState) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SamlRelayState) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


