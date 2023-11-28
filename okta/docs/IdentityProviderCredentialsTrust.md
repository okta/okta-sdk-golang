# IdentityProviderCredentialsTrust

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Kid** | Pointer to **string** |  | [optional] 
**Revocation** | Pointer to [**IdentityProviderCredentialsTrustRevocation**](IdentityProviderCredentialsTrustRevocation.md) |  | [optional] 
**RevocationCacheLifetime** | Pointer to **int32** |  | [optional] 

## Methods

### NewIdentityProviderCredentialsTrust

`func NewIdentityProviderCredentialsTrust() *IdentityProviderCredentialsTrust`

NewIdentityProviderCredentialsTrust instantiates a new IdentityProviderCredentialsTrust object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderCredentialsTrustWithDefaults

`func NewIdentityProviderCredentialsTrustWithDefaults() *IdentityProviderCredentialsTrust`

NewIdentityProviderCredentialsTrustWithDefaults instantiates a new IdentityProviderCredentialsTrust object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *IdentityProviderCredentialsTrust) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *IdentityProviderCredentialsTrust) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *IdentityProviderCredentialsTrust) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *IdentityProviderCredentialsTrust) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *IdentityProviderCredentialsTrust) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IdentityProviderCredentialsTrust) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IdentityProviderCredentialsTrust) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *IdentityProviderCredentialsTrust) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKid

`func (o *IdentityProviderCredentialsTrust) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *IdentityProviderCredentialsTrust) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *IdentityProviderCredentialsTrust) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *IdentityProviderCredentialsTrust) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetRevocation

`func (o *IdentityProviderCredentialsTrust) GetRevocation() IdentityProviderCredentialsTrustRevocation`

GetRevocation returns the Revocation field if non-nil, zero value otherwise.

### GetRevocationOk

`func (o *IdentityProviderCredentialsTrust) GetRevocationOk() (*IdentityProviderCredentialsTrustRevocation, bool)`

GetRevocationOk returns a tuple with the Revocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocation

`func (o *IdentityProviderCredentialsTrust) SetRevocation(v IdentityProviderCredentialsTrustRevocation)`

SetRevocation sets Revocation field to given value.

### HasRevocation

`func (o *IdentityProviderCredentialsTrust) HasRevocation() bool`

HasRevocation returns a boolean if a field has been set.

### GetRevocationCacheLifetime

`func (o *IdentityProviderCredentialsTrust) GetRevocationCacheLifetime() int32`

GetRevocationCacheLifetime returns the RevocationCacheLifetime field if non-nil, zero value otherwise.

### GetRevocationCacheLifetimeOk

`func (o *IdentityProviderCredentialsTrust) GetRevocationCacheLifetimeOk() (*int32, bool)`

GetRevocationCacheLifetimeOk returns a tuple with the RevocationCacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationCacheLifetime

`func (o *IdentityProviderCredentialsTrust) SetRevocationCacheLifetime(v int32)`

SetRevocationCacheLifetime sets RevocationCacheLifetime field to given value.

### HasRevocationCacheLifetime

`func (o *IdentityProviderCredentialsTrust) HasRevocationCacheLifetime() bool`

HasRevocationCacheLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


