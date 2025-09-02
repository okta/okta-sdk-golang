# WellKnownSSFMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationSchemes** | Pointer to [**[]WellKnownSSFMetadataSpecUrn**](WellKnownSSFMetadataSpecUrn.md) | An array of JSON objects that specify the authorization scheme properties supported by the transmitter | [optional] 
**ConfigurationEndpoint** | Pointer to **string** | The URL of the SSF Stream configuration endpoint | [optional] 
**DefaultSubjects** | Pointer to **string** | A string that indicates the default behavior of newly created streams | [optional] 
**DeliveryMethodsSupported** | Pointer to **[]string** | An array of supported SET delivery methods | [optional] 
**Issuer** | Pointer to **string** | The issuer used in Security Event Tokens. This value is set as &#x60;iss&#x60; in the claim. | [optional] 
**JwksUri** | Pointer to **string** | The URL of the JSON Web Key Set (JWKS) that contains the signing keys for validating the signatures of Security Event Tokens (SETs) | [optional] 
**SpecVersion** | Pointer to **string** | The version identifying the implementer&#39;s draft or final specification implemented by the transmitter | [optional] 
**VerificationEndpoint** | Pointer to **string** | The URL of the SSF Stream verification endpoint | [optional] 

## Methods

### NewWellKnownSSFMetadata

`func NewWellKnownSSFMetadata() *WellKnownSSFMetadata`

NewWellKnownSSFMetadata instantiates a new WellKnownSSFMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWellKnownSSFMetadataWithDefaults

`func NewWellKnownSSFMetadataWithDefaults() *WellKnownSSFMetadata`

NewWellKnownSSFMetadataWithDefaults instantiates a new WellKnownSSFMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationSchemes

`func (o *WellKnownSSFMetadata) GetAuthorizationSchemes() []WellKnownSSFMetadataSpecUrn`

GetAuthorizationSchemes returns the AuthorizationSchemes field if non-nil, zero value otherwise.

### GetAuthorizationSchemesOk

`func (o *WellKnownSSFMetadata) GetAuthorizationSchemesOk() (*[]WellKnownSSFMetadataSpecUrn, bool)`

GetAuthorizationSchemesOk returns a tuple with the AuthorizationSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationSchemes

`func (o *WellKnownSSFMetadata) SetAuthorizationSchemes(v []WellKnownSSFMetadataSpecUrn)`

SetAuthorizationSchemes sets AuthorizationSchemes field to given value.

### HasAuthorizationSchemes

`func (o *WellKnownSSFMetadata) HasAuthorizationSchemes() bool`

HasAuthorizationSchemes returns a boolean if a field has been set.

### GetConfigurationEndpoint

`func (o *WellKnownSSFMetadata) GetConfigurationEndpoint() string`

GetConfigurationEndpoint returns the ConfigurationEndpoint field if non-nil, zero value otherwise.

### GetConfigurationEndpointOk

`func (o *WellKnownSSFMetadata) GetConfigurationEndpointOk() (*string, bool)`

GetConfigurationEndpointOk returns a tuple with the ConfigurationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationEndpoint

`func (o *WellKnownSSFMetadata) SetConfigurationEndpoint(v string)`

SetConfigurationEndpoint sets ConfigurationEndpoint field to given value.

### HasConfigurationEndpoint

`func (o *WellKnownSSFMetadata) HasConfigurationEndpoint() bool`

HasConfigurationEndpoint returns a boolean if a field has been set.

### GetDefaultSubjects

`func (o *WellKnownSSFMetadata) GetDefaultSubjects() string`

GetDefaultSubjects returns the DefaultSubjects field if non-nil, zero value otherwise.

### GetDefaultSubjectsOk

`func (o *WellKnownSSFMetadata) GetDefaultSubjectsOk() (*string, bool)`

GetDefaultSubjectsOk returns a tuple with the DefaultSubjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubjects

`func (o *WellKnownSSFMetadata) SetDefaultSubjects(v string)`

SetDefaultSubjects sets DefaultSubjects field to given value.

### HasDefaultSubjects

`func (o *WellKnownSSFMetadata) HasDefaultSubjects() bool`

HasDefaultSubjects returns a boolean if a field has been set.

### GetDeliveryMethodsSupported

`func (o *WellKnownSSFMetadata) GetDeliveryMethodsSupported() []string`

GetDeliveryMethodsSupported returns the DeliveryMethodsSupported field if non-nil, zero value otherwise.

### GetDeliveryMethodsSupportedOk

`func (o *WellKnownSSFMetadata) GetDeliveryMethodsSupportedOk() (*[]string, bool)`

GetDeliveryMethodsSupportedOk returns a tuple with the DeliveryMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethodsSupported

`func (o *WellKnownSSFMetadata) SetDeliveryMethodsSupported(v []string)`

SetDeliveryMethodsSupported sets DeliveryMethodsSupported field to given value.

### HasDeliveryMethodsSupported

`func (o *WellKnownSSFMetadata) HasDeliveryMethodsSupported() bool`

HasDeliveryMethodsSupported returns a boolean if a field has been set.

### GetIssuer

`func (o *WellKnownSSFMetadata) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *WellKnownSSFMetadata) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *WellKnownSSFMetadata) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *WellKnownSSFMetadata) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJwksUri

`func (o *WellKnownSSFMetadata) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *WellKnownSSFMetadata) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *WellKnownSSFMetadata) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *WellKnownSSFMetadata) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetSpecVersion

`func (o *WellKnownSSFMetadata) GetSpecVersion() string`

GetSpecVersion returns the SpecVersion field if non-nil, zero value otherwise.

### GetSpecVersionOk

`func (o *WellKnownSSFMetadata) GetSpecVersionOk() (*string, bool)`

GetSpecVersionOk returns a tuple with the SpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecVersion

`func (o *WellKnownSSFMetadata) SetSpecVersion(v string)`

SetSpecVersion sets SpecVersion field to given value.

### HasSpecVersion

`func (o *WellKnownSSFMetadata) HasSpecVersion() bool`

HasSpecVersion returns a boolean if a field has been set.

### GetVerificationEndpoint

`func (o *WellKnownSSFMetadata) GetVerificationEndpoint() string`

GetVerificationEndpoint returns the VerificationEndpoint field if non-nil, zero value otherwise.

### GetVerificationEndpointOk

`func (o *WellKnownSSFMetadata) GetVerificationEndpointOk() (*string, bool)`

GetVerificationEndpointOk returns a tuple with the VerificationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationEndpoint

`func (o *WellKnownSSFMetadata) SetVerificationEndpoint(v string)`

SetVerificationEndpoint sets VerificationEndpoint field to given value.

### HasVerificationEndpoint

`func (o *WellKnownSSFMetadata) HasVerificationEndpoint() bool`

HasVerificationEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


