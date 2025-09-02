# SamlSigningCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kid** | Pointer to **string** | IdP key credential reference to the Okta X.509 signature certificate | [optional] 

## Methods

### NewSamlSigningCredentials

`func NewSamlSigningCredentials() *SamlSigningCredentials`

NewSamlSigningCredentials instantiates a new SamlSigningCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlSigningCredentialsWithDefaults

`func NewSamlSigningCredentialsWithDefaults() *SamlSigningCredentials`

NewSamlSigningCredentialsWithDefaults instantiates a new SamlSigningCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKid

`func (o *SamlSigningCredentials) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *SamlSigningCredentials) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *SamlSigningCredentials) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *SamlSigningCredentials) HasKid() bool`

HasKid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


