# SecurityEventsProviderSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** | Issuer URL | [optional] 
**JwksUrl** | Pointer to **string** | The public URL where the JWKS public key is uploaded | [optional] 
**WellKnownUrl** | Pointer to **NullableString** | The well-known URL of the Security Events Provider (the SSF transmitter) | [optional] 

## Methods

### NewSecurityEventsProviderSettingsResponse

`func NewSecurityEventsProviderSettingsResponse() *SecurityEventsProviderSettingsResponse`

NewSecurityEventsProviderSettingsResponse instantiates a new SecurityEventsProviderSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventsProviderSettingsResponseWithDefaults

`func NewSecurityEventsProviderSettingsResponseWithDefaults() *SecurityEventsProviderSettingsResponse`

NewSecurityEventsProviderSettingsResponseWithDefaults instantiates a new SecurityEventsProviderSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *SecurityEventsProviderSettingsResponse) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SecurityEventsProviderSettingsResponse) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SecurityEventsProviderSettingsResponse) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SecurityEventsProviderSettingsResponse) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJwksUrl

`func (o *SecurityEventsProviderSettingsResponse) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *SecurityEventsProviderSettingsResponse) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *SecurityEventsProviderSettingsResponse) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.

### HasJwksUrl

`func (o *SecurityEventsProviderSettingsResponse) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.

### GetWellKnownUrl

`func (o *SecurityEventsProviderSettingsResponse) GetWellKnownUrl() string`

GetWellKnownUrl returns the WellKnownUrl field if non-nil, zero value otherwise.

### GetWellKnownUrlOk

`func (o *SecurityEventsProviderSettingsResponse) GetWellKnownUrlOk() (*string, bool)`

GetWellKnownUrlOk returns a tuple with the WellKnownUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWellKnownUrl

`func (o *SecurityEventsProviderSettingsResponse) SetWellKnownUrl(v string)`

SetWellKnownUrl sets WellKnownUrl field to given value.

### HasWellKnownUrl

`func (o *SecurityEventsProviderSettingsResponse) HasWellKnownUrl() bool`

HasWellKnownUrl returns a boolean if a field has been set.

### SetWellKnownUrlNil

`func (o *SecurityEventsProviderSettingsResponse) SetWellKnownUrlNil(b bool)`

 SetWellKnownUrlNil sets the value for WellKnownUrl to be an explicit nil

### UnsetWellKnownUrl
`func (o *SecurityEventsProviderSettingsResponse) UnsetWellKnownUrl()`

UnsetWellKnownUrl ensures that no value is present for WellKnownUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


