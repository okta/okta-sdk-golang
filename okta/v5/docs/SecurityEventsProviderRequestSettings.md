# SecurityEventsProviderRequestSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WellKnownUrl** | **string** | The published well-known URL of the Security Events Provider (the SSF transmitter) | 
**Issuer** | **string** | Issuer URL | 
**JwksUrl** | **string** | The public URL where the JWKS public key is uploaded | 

## Methods

### NewSecurityEventsProviderRequestSettings

`func NewSecurityEventsProviderRequestSettings(wellKnownUrl string, issuer string, jwksUrl string, ) *SecurityEventsProviderRequestSettings`

NewSecurityEventsProviderRequestSettings instantiates a new SecurityEventsProviderRequestSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventsProviderRequestSettingsWithDefaults

`func NewSecurityEventsProviderRequestSettingsWithDefaults() *SecurityEventsProviderRequestSettings`

NewSecurityEventsProviderRequestSettingsWithDefaults instantiates a new SecurityEventsProviderRequestSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWellKnownUrl

`func (o *SecurityEventsProviderRequestSettings) GetWellKnownUrl() string`

GetWellKnownUrl returns the WellKnownUrl field if non-nil, zero value otherwise.

### GetWellKnownUrlOk

`func (o *SecurityEventsProviderRequestSettings) GetWellKnownUrlOk() (*string, bool)`

GetWellKnownUrlOk returns a tuple with the WellKnownUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWellKnownUrl

`func (o *SecurityEventsProviderRequestSettings) SetWellKnownUrl(v string)`

SetWellKnownUrl sets WellKnownUrl field to given value.


### GetIssuer

`func (o *SecurityEventsProviderRequestSettings) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SecurityEventsProviderRequestSettings) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SecurityEventsProviderRequestSettings) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetJwksUrl

`func (o *SecurityEventsProviderRequestSettings) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *SecurityEventsProviderRequestSettings) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *SecurityEventsProviderRequestSettings) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


