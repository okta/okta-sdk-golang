# SecurityEventsProviderSettingsNonSSFCompliant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | **string** | Issuer URL | 
**JwksUrl** | **string** | The public URL where the JWKS public key is uploaded | 

## Methods

### NewSecurityEventsProviderSettingsNonSSFCompliant

`func NewSecurityEventsProviderSettingsNonSSFCompliant(issuer string, jwksUrl string, ) *SecurityEventsProviderSettingsNonSSFCompliant`

NewSecurityEventsProviderSettingsNonSSFCompliant instantiates a new SecurityEventsProviderSettingsNonSSFCompliant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventsProviderSettingsNonSSFCompliantWithDefaults

`func NewSecurityEventsProviderSettingsNonSSFCompliantWithDefaults() *SecurityEventsProviderSettingsNonSSFCompliant`

NewSecurityEventsProviderSettingsNonSSFCompliantWithDefaults instantiates a new SecurityEventsProviderSettingsNonSSFCompliant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *SecurityEventsProviderSettingsNonSSFCompliant) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SecurityEventsProviderSettingsNonSSFCompliant) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SecurityEventsProviderSettingsNonSSFCompliant) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetJwksUrl

`func (o *SecurityEventsProviderSettingsNonSSFCompliant) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *SecurityEventsProviderSettingsNonSSFCompliant) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *SecurityEventsProviderSettingsNonSSFCompliant) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


