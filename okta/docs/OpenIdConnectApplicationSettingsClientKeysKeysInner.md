# OpenIdConnectApplicationSettingsClientKeysKeysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] 
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAUth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]
**Created** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was created | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of the OAuth Client JSON Web Key | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was updated | [optional] [readonly] 
**Links** | Pointer to [**OAuthClientSecretLinks**](OAuthClientSecretLinks.md) |  | [optional] 

## Methods

### NewOpenIdConnectApplicationSettingsClientKeysKeysInner

`func NewOpenIdConnectApplicationSettingsClientKeysKeysInner() *OpenIdConnectApplicationSettingsClientKeysKeysInner`

NewOpenIdConnectApplicationSettingsClientKeysKeysInner instantiates a new OpenIdConnectApplicationSettingsClientKeysKeysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectApplicationSettingsClientKeysKeysInnerWithDefaults

`func NewOpenIdConnectApplicationSettingsClientKeysKeysInnerWithDefaults() *OpenIdConnectApplicationSettingsClientKeysKeysInner`

NewOpenIdConnectApplicationSettingsClientKeysKeysInnerWithDefaults instantiates a new OpenIdConnectApplicationSettingsClientKeysKeysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKty

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) HasN() bool`

HasN returns a boolean if a field has been set.

### GetUse

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetKid

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetLinks() OAuthClientSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) GetLinksOk() (*OAuthClientSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) SetLinks(v OAuthClientSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OpenIdConnectApplicationSettingsClientKeysKeysInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


