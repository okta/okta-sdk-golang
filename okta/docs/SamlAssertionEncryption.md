# SamlAssertionEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether Okta encrypts the assertions that it sends to the Service Provider | [optional] 
**EncryptionAlgorithm** | Pointer to **string** | The encryption algorithm used to encrypt the SAML assertion | [optional] 
**KeyTransportAlgorithm** | Pointer to **string** | The key transport algorithm used to encrypt the SAML assertion | [optional] 
**X5c** | Pointer to **[]string** | A list that contains exactly one x509 encoded certificate which Okta uses to encrypt the SAML assertion with | [optional] 

## Methods

### NewSamlAssertionEncryption

`func NewSamlAssertionEncryption() *SamlAssertionEncryption`

NewSamlAssertionEncryption instantiates a new SamlAssertionEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAssertionEncryptionWithDefaults

`func NewSamlAssertionEncryptionWithDefaults() *SamlAssertionEncryption`

NewSamlAssertionEncryptionWithDefaults instantiates a new SamlAssertionEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SamlAssertionEncryption) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SamlAssertionEncryption) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SamlAssertionEncryption) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SamlAssertionEncryption) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *SamlAssertionEncryption) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *SamlAssertionEncryption) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *SamlAssertionEncryption) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *SamlAssertionEncryption) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetKeyTransportAlgorithm

`func (o *SamlAssertionEncryption) GetKeyTransportAlgorithm() string`

GetKeyTransportAlgorithm returns the KeyTransportAlgorithm field if non-nil, zero value otherwise.

### GetKeyTransportAlgorithmOk

`func (o *SamlAssertionEncryption) GetKeyTransportAlgorithmOk() (*string, bool)`

GetKeyTransportAlgorithmOk returns a tuple with the KeyTransportAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTransportAlgorithm

`func (o *SamlAssertionEncryption) SetKeyTransportAlgorithm(v string)`

SetKeyTransportAlgorithm sets KeyTransportAlgorithm field to given value.

### HasKeyTransportAlgorithm

`func (o *SamlAssertionEncryption) HasKeyTransportAlgorithm() bool`

HasKeyTransportAlgorithm returns a boolean if a field has been set.

### GetX5c

`func (o *SamlAssertionEncryption) GetX5c() []string`

GetX5c returns the X5c field if non-nil, zero value otherwise.

### GetX5cOk

`func (o *SamlAssertionEncryption) GetX5cOk() (*[]string, bool)`

GetX5cOk returns a tuple with the X5c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5c

`func (o *SamlAssertionEncryption) SetX5c(v []string)`

SetX5c sets X5c field to given value.

### HasX5c

`func (o *SamlAssertionEncryption) HasX5c() bool`

HasX5c returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


