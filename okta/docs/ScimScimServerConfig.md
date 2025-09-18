# ScimScimServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Patch** | Pointer to [**ScimScimServerConfigPatch**](ScimScimServerConfigPatch.md) |  | [optional] 
**ChangePassword** | Pointer to [**ScimScimServerConfigChangePassword**](ScimScimServerConfigChangePassword.md) |  | [optional] 

## Methods

### NewScimScimServerConfig

`func NewScimScimServerConfig() *ScimScimServerConfig`

NewScimScimServerConfig instantiates a new ScimScimServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimScimServerConfigWithDefaults

`func NewScimScimServerConfigWithDefaults() *ScimScimServerConfig`

NewScimScimServerConfigWithDefaults instantiates a new ScimScimServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatch

`func (o *ScimScimServerConfig) GetPatch() ScimScimServerConfigPatch`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *ScimScimServerConfig) GetPatchOk() (*ScimScimServerConfigPatch, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *ScimScimServerConfig) SetPatch(v ScimScimServerConfigPatch)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *ScimScimServerConfig) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetChangePassword

`func (o *ScimScimServerConfig) GetChangePassword() ScimScimServerConfigChangePassword`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### GetChangePasswordOk

`func (o *ScimScimServerConfig) GetChangePasswordOk() (*ScimScimServerConfigChangePassword, bool)`

GetChangePasswordOk returns a tuple with the ChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassword

`func (o *ScimScimServerConfig) SetChangePassword(v ScimScimServerConfigChangePassword)`

SetChangePassword sets ChangePassword field to given value.

### HasChangePassword

`func (o *ScimScimServerConfig) HasChangePassword() bool`

HasChangePassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


