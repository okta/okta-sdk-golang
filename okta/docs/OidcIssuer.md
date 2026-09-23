# OidcIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Issuer URL identifying the external IdP (the expected &#x60;iss&#x60; value in OIDC tokens) | [optional] 

## Methods

### NewOidcIssuer

`func NewOidcIssuer() *OidcIssuer`

NewOidcIssuer instantiates a new OidcIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcIssuerWithDefaults

`func NewOidcIssuerWithDefaults() *OidcIssuer`

NewOidcIssuerWithDefaults instantiates a new OidcIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OidcIssuer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OidcIssuer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OidcIssuer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OidcIssuer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


