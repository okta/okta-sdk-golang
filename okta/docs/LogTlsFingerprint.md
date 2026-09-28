# LogTlsFingerprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ja4** | Pointer to **NullableString** | The JA4 TLS client fingerprint hash associated with the event&#39;s request | [optional] [readonly] 

## Methods

### NewLogTlsFingerprint

`func NewLogTlsFingerprint() *LogTlsFingerprint`

NewLogTlsFingerprint instantiates a new LogTlsFingerprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogTlsFingerprintWithDefaults

`func NewLogTlsFingerprintWithDefaults() *LogTlsFingerprint`

NewLogTlsFingerprintWithDefaults instantiates a new LogTlsFingerprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJa4

`func (o *LogTlsFingerprint) GetJa4() string`

GetJa4 returns the Ja4 field if non-nil, zero value otherwise.

### GetJa4Ok

`func (o *LogTlsFingerprint) GetJa4Ok() (*string, bool)`

GetJa4Ok returns a tuple with the Ja4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJa4

`func (o *LogTlsFingerprint) SetJa4(v string)`

SetJa4 sets Ja4 field to given value.

### HasJa4

`func (o *LogTlsFingerprint) HasJa4() bool`

HasJa4 returns a boolean if a field has been set.

### SetJa4Nil

`func (o *LogTlsFingerprint) SetJa4Nil(b bool)`

 SetJa4Nil sets the value for Ja4 to be an explicit nil

### UnsetJa4
`func (o *LogTlsFingerprint) UnsetJa4()`

UnsetJa4 ensures that no value is present for Ja4, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


