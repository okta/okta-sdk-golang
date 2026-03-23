# SamlRequestAlgorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;XML Digest Algorithm setting for cryptographically hashing &#x60;&lt;AuthnRequest&gt;&#x60; messages sent to the IdP &gt; **Note:**  This property is ignored when you disable request signatures (&#x60;scope&#x60; set as &#x60;NONE&#x60;). | [optional] 
**Signature** | Pointer to [**SamlRequestSignatureAlgorithm**](SamlRequestSignatureAlgorithm.md) |  | [optional] 

## Methods

### NewSamlRequestAlgorithm

`func NewSamlRequestAlgorithm() *SamlRequestAlgorithm`

NewSamlRequestAlgorithm instantiates a new SamlRequestAlgorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlRequestAlgorithmWithDefaults

`func NewSamlRequestAlgorithmWithDefaults() *SamlRequestAlgorithm`

NewSamlRequestAlgorithmWithDefaults instantiates a new SamlRequestAlgorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *SamlRequestAlgorithm) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *SamlRequestAlgorithm) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *SamlRequestAlgorithm) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *SamlRequestAlgorithm) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetSignature

`func (o *SamlRequestAlgorithm) GetSignature() SamlRequestSignatureAlgorithm`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SamlRequestAlgorithm) GetSignatureOk() (*SamlRequestSignatureAlgorithm, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SamlRequestAlgorithm) SetSignature(v SamlRequestSignatureAlgorithm)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SamlRequestAlgorithm) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


