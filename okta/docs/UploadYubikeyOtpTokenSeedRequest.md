# UploadYubikeyOtpTokenSeedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **string** | The unique identifier assigned to each YubiKey device | [optional] 
**PublicId** | Pointer to **string** | The YubiKey&#39;s public ID | [optional] 
**PrivateId** | Pointer to **string** | The YubiKey&#39;s private ID | [optional] 
**AesKey** | Pointer to **string** | The cryptographic key used in the AES (Advanced Encryption Standard) algorithm to encrypt and decrypt the YubiKey OTP | [optional] 

## Methods

### NewUploadYubikeyOtpTokenSeedRequest

`func NewUploadYubikeyOtpTokenSeedRequest() *UploadYubikeyOtpTokenSeedRequest`

NewUploadYubikeyOtpTokenSeedRequest instantiates a new UploadYubikeyOtpTokenSeedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadYubikeyOtpTokenSeedRequestWithDefaults

`func NewUploadYubikeyOtpTokenSeedRequestWithDefaults() *UploadYubikeyOtpTokenSeedRequest`

NewUploadYubikeyOtpTokenSeedRequestWithDefaults instantiates a new UploadYubikeyOtpTokenSeedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *UploadYubikeyOtpTokenSeedRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *UploadYubikeyOtpTokenSeedRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *UploadYubikeyOtpTokenSeedRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *UploadYubikeyOtpTokenSeedRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetPublicId

`func (o *UploadYubikeyOtpTokenSeedRequest) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UploadYubikeyOtpTokenSeedRequest) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UploadYubikeyOtpTokenSeedRequest) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UploadYubikeyOtpTokenSeedRequest) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetPrivateId

`func (o *UploadYubikeyOtpTokenSeedRequest) GetPrivateId() string`

GetPrivateId returns the PrivateId field if non-nil, zero value otherwise.

### GetPrivateIdOk

`func (o *UploadYubikeyOtpTokenSeedRequest) GetPrivateIdOk() (*string, bool)`

GetPrivateIdOk returns a tuple with the PrivateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateId

`func (o *UploadYubikeyOtpTokenSeedRequest) SetPrivateId(v string)`

SetPrivateId sets PrivateId field to given value.

### HasPrivateId

`func (o *UploadYubikeyOtpTokenSeedRequest) HasPrivateId() bool`

HasPrivateId returns a boolean if a field has been set.

### GetAesKey

`func (o *UploadYubikeyOtpTokenSeedRequest) GetAesKey() string`

GetAesKey returns the AesKey field if non-nil, zero value otherwise.

### GetAesKeyOk

`func (o *UploadYubikeyOtpTokenSeedRequest) GetAesKeyOk() (*string, bool)`

GetAesKeyOk returns a tuple with the AesKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAesKey

`func (o *UploadYubikeyOtpTokenSeedRequest) SetAesKey(v string)`

SetAesKey sets AesKey field to given value.

### HasAesKey

`func (o *UploadYubikeyOtpTokenSeedRequest) HasAesKey() bool`

HasAesKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


