# ApplicationCredentialsSigning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kid** | Pointer to **string** |  | [optional] 
**LastRotated** | Pointer to **time.Time** |  | [optional] [readonly] 
**NextRotation** | Pointer to **time.Time** |  | [optional] [readonly] 
**RotationMode** | Pointer to **string** |  | [optional] 
**Use** | Pointer to [**ApplicationCredentialsSigningUse**](ApplicationCredentialsSigningUse.md) |  | [optional] 

## Methods

### NewApplicationCredentialsSigning

`func NewApplicationCredentialsSigning() *ApplicationCredentialsSigning`

NewApplicationCredentialsSigning instantiates a new ApplicationCredentialsSigning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCredentialsSigningWithDefaults

`func NewApplicationCredentialsSigningWithDefaults() *ApplicationCredentialsSigning`

NewApplicationCredentialsSigningWithDefaults instantiates a new ApplicationCredentialsSigning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKid

`func (o *ApplicationCredentialsSigning) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *ApplicationCredentialsSigning) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *ApplicationCredentialsSigning) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *ApplicationCredentialsSigning) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetLastRotated

`func (o *ApplicationCredentialsSigning) GetLastRotated() time.Time`

GetLastRotated returns the LastRotated field if non-nil, zero value otherwise.

### GetLastRotatedOk

`func (o *ApplicationCredentialsSigning) GetLastRotatedOk() (*time.Time, bool)`

GetLastRotatedOk returns a tuple with the LastRotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotated

`func (o *ApplicationCredentialsSigning) SetLastRotated(v time.Time)`

SetLastRotated sets LastRotated field to given value.

### HasLastRotated

`func (o *ApplicationCredentialsSigning) HasLastRotated() bool`

HasLastRotated returns a boolean if a field has been set.

### GetNextRotation

`func (o *ApplicationCredentialsSigning) GetNextRotation() time.Time`

GetNextRotation returns the NextRotation field if non-nil, zero value otherwise.

### GetNextRotationOk

`func (o *ApplicationCredentialsSigning) GetNextRotationOk() (*time.Time, bool)`

GetNextRotationOk returns a tuple with the NextRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRotation

`func (o *ApplicationCredentialsSigning) SetNextRotation(v time.Time)`

SetNextRotation sets NextRotation field to given value.

### HasNextRotation

`func (o *ApplicationCredentialsSigning) HasNextRotation() bool`

HasNextRotation returns a boolean if a field has been set.

### GetRotationMode

`func (o *ApplicationCredentialsSigning) GetRotationMode() string`

GetRotationMode returns the RotationMode field if non-nil, zero value otherwise.

### GetRotationModeOk

`func (o *ApplicationCredentialsSigning) GetRotationModeOk() (*string, bool)`

GetRotationModeOk returns a tuple with the RotationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationMode

`func (o *ApplicationCredentialsSigning) SetRotationMode(v string)`

SetRotationMode sets RotationMode field to given value.

### HasRotationMode

`func (o *ApplicationCredentialsSigning) HasRotationMode() bool`

HasRotationMode returns a boolean if a field has been set.

### GetUse

`func (o *ApplicationCredentialsSigning) GetUse() ApplicationCredentialsSigningUse`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *ApplicationCredentialsSigning) GetUseOk() (*ApplicationCredentialsSigningUse, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *ApplicationCredentialsSigning) SetUse(v ApplicationCredentialsSigningUse)`

SetUse sets Use field to given value.

### HasUse

`func (o *ApplicationCredentialsSigning) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


