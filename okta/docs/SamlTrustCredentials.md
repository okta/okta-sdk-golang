# SamlTrustCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalKids** | Pointer to **[]string** | Additional IdP key credential reference to the Okta X.509 signature certificate | [optional] 
**Audience** | Pointer to **string** | URI that identifies the target Okta IdP instance (SP) for an &#x60;&lt;Assertion&gt;&#x60; | [optional] 
**Issuer** | Pointer to **string** | URI that identifies the issuer (IdP) of a &#x60;&lt;SAMLResponse&gt;&#x60; message &#x60;&lt;Assertion&gt;&#x60; element | [optional] 
**Kid** | Pointer to **string** | IdP key credential reference to the Okta X.509 signature certificate | [optional] 

## Methods

### NewSamlTrustCredentials

`func NewSamlTrustCredentials() *SamlTrustCredentials`

NewSamlTrustCredentials instantiates a new SamlTrustCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlTrustCredentialsWithDefaults

`func NewSamlTrustCredentialsWithDefaults() *SamlTrustCredentials`

NewSamlTrustCredentialsWithDefaults instantiates a new SamlTrustCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalKids

`func (o *SamlTrustCredentials) GetAdditionalKids() []string`

GetAdditionalKids returns the AdditionalKids field if non-nil, zero value otherwise.

### GetAdditionalKidsOk

`func (o *SamlTrustCredentials) GetAdditionalKidsOk() (*[]string, bool)`

GetAdditionalKidsOk returns a tuple with the AdditionalKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalKids

`func (o *SamlTrustCredentials) SetAdditionalKids(v []string)`

SetAdditionalKids sets AdditionalKids field to given value.

### HasAdditionalKids

`func (o *SamlTrustCredentials) HasAdditionalKids() bool`

HasAdditionalKids returns a boolean if a field has been set.

### GetAudience

`func (o *SamlTrustCredentials) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *SamlTrustCredentials) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *SamlTrustCredentials) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *SamlTrustCredentials) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *SamlTrustCredentials) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SamlTrustCredentials) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SamlTrustCredentials) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SamlTrustCredentials) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKid

`func (o *SamlTrustCredentials) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *SamlTrustCredentials) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *SamlTrustCredentials) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *SamlTrustCredentials) HasKid() bool`

HasKid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


