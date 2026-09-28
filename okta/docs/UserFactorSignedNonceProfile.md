# UserFactorSignedNonceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | Pointer to **string** | ID for the factor credential | [optional] 
**DeviceType** | Pointer to **string** | Type of device | [optional] 
**Keys** | Pointer to [**[]UserFactorSignedNonceProfileKey**](UserFactorSignedNonceProfileKey.md) | Cryptographic keys associated with the signed nonce factor | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**Platform** | Pointer to **string** | OS platform of the associated device | [optional] 
**Version** | Pointer to **string** | OS version of the associated device | [optional] 

## Methods

### NewUserFactorSignedNonceProfile

`func NewUserFactorSignedNonceProfile() *UserFactorSignedNonceProfile`

NewUserFactorSignedNonceProfile instantiates a new UserFactorSignedNonceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorSignedNonceProfileWithDefaults

`func NewUserFactorSignedNonceProfileWithDefaults() *UserFactorSignedNonceProfile`

NewUserFactorSignedNonceProfileWithDefaults instantiates a new UserFactorSignedNonceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *UserFactorSignedNonceProfile) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *UserFactorSignedNonceProfile) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *UserFactorSignedNonceProfile) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *UserFactorSignedNonceProfile) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetDeviceType

`func (o *UserFactorSignedNonceProfile) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *UserFactorSignedNonceProfile) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *UserFactorSignedNonceProfile) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *UserFactorSignedNonceProfile) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetKeys

`func (o *UserFactorSignedNonceProfile) GetKeys() []UserFactorSignedNonceProfileKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *UserFactorSignedNonceProfile) GetKeysOk() (*[]UserFactorSignedNonceProfileKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *UserFactorSignedNonceProfile) SetKeys(v []UserFactorSignedNonceProfileKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *UserFactorSignedNonceProfile) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetName

`func (o *UserFactorSignedNonceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserFactorSignedNonceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserFactorSignedNonceProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserFactorSignedNonceProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *UserFactorSignedNonceProfile) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UserFactorSignedNonceProfile) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UserFactorSignedNonceProfile) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UserFactorSignedNonceProfile) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *UserFactorSignedNonceProfile) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserFactorSignedNonceProfile) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserFactorSignedNonceProfile) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UserFactorSignedNonceProfile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


