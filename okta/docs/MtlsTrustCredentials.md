# MtlsTrustCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** | Not used | [optional] 
**Issuer** | Pointer to **string** | Description of the certificate issuer | [optional] 
**Kid** | Pointer to **string** | IdP key credential reference to the Okta X.509 signature certificate | [optional] 
**Revocation** | Pointer to **string** | Mechanism to validate the certificate  &gt; **Note:** This property isn&#39;t supported. Okta now handles CRL caching automatically. As of October 8, 2025, in Preview orgs, and October 13, 2025, in Production orgs, this property is ignored if it&#39;s specified in any API requests. Specifying the property in your API requests doesn&#39;t cause any errors since the property has no effect. &gt; &gt; See [Deprecation Notice - Smart Card IdP Legacy CRL Cache Setting](https://support.okta.com/help/s/article/deprecation-notice-smart-card-idp-legacy-crl-cache-setting?language&#x3D;en_US). | [optional] 
**RevocationCacheLifetime** | Pointer to **float32** | Time in minutes to cache the certificate revocation information  &gt; **Note:** This property isn&#39;t supported. Okta now handles CRL caching automatically. As of October 8, 2025, in Preview orgs, and October 13, 2025, in Production orgs, this property is ignored if it&#39;s specified in any API requests. Specifying this property in your API requests doesn&#39;t cause errors since the property has no effect. &gt; &gt; See [Deprecation Notice - Smart Card IdP Legacy CRL Cache Setting](https://support.okta.com/help/s/article/deprecation-notice-smart-card-idp-legacy-crl-cache-setting?language&#x3D;en_US). | [optional] 

## Methods

### NewMtlsTrustCredentials

`func NewMtlsTrustCredentials() *MtlsTrustCredentials`

NewMtlsTrustCredentials instantiates a new MtlsTrustCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMtlsTrustCredentialsWithDefaults

`func NewMtlsTrustCredentialsWithDefaults() *MtlsTrustCredentials`

NewMtlsTrustCredentialsWithDefaults instantiates a new MtlsTrustCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *MtlsTrustCredentials) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *MtlsTrustCredentials) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *MtlsTrustCredentials) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *MtlsTrustCredentials) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *MtlsTrustCredentials) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *MtlsTrustCredentials) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *MtlsTrustCredentials) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *MtlsTrustCredentials) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKid

`func (o *MtlsTrustCredentials) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *MtlsTrustCredentials) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *MtlsTrustCredentials) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *MtlsTrustCredentials) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetRevocation

`func (o *MtlsTrustCredentials) GetRevocation() string`

GetRevocation returns the Revocation field if non-nil, zero value otherwise.

### GetRevocationOk

`func (o *MtlsTrustCredentials) GetRevocationOk() (*string, bool)`

GetRevocationOk returns a tuple with the Revocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocation

`func (o *MtlsTrustCredentials) SetRevocation(v string)`

SetRevocation sets Revocation field to given value.

### HasRevocation

`func (o *MtlsTrustCredentials) HasRevocation() bool`

HasRevocation returns a boolean if a field has been set.

### GetRevocationCacheLifetime

`func (o *MtlsTrustCredentials) GetRevocationCacheLifetime() float32`

GetRevocationCacheLifetime returns the RevocationCacheLifetime field if non-nil, zero value otherwise.

### GetRevocationCacheLifetimeOk

`func (o *MtlsTrustCredentials) GetRevocationCacheLifetimeOk() (*float32, bool)`

GetRevocationCacheLifetimeOk returns a tuple with the RevocationCacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationCacheLifetime

`func (o *MtlsTrustCredentials) SetRevocationCacheLifetime(v float32)`

SetRevocationCacheLifetime sets RevocationCacheLifetime field to given value.

### HasRevocationCacheLifetime

`func (o *MtlsTrustCredentials) HasRevocationCacheLifetime() bool`

HasRevocationCacheLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


