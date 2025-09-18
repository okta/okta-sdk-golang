# CredentialSyncInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | The error code for the type of error | [optional] [readonly] 
**ErrorReason** | Pointer to **string** | A short description of the error | [optional] [readonly] 
**SecretVersionId** | Pointer to **string** | The version ID of the password secret from the OPA vault. | [optional] 
**SyncState** | Pointer to **string** | Current credential sync status of the privileged resource | [optional] [readonly] 
**SyncTime** | Pointer to **time.Time** | Timestamp when the credential was changed | [optional] [readonly] 

## Methods

### NewCredentialSyncInfo

`func NewCredentialSyncInfo() *CredentialSyncInfo`

NewCredentialSyncInfo instantiates a new CredentialSyncInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialSyncInfoWithDefaults

`func NewCredentialSyncInfoWithDefaults() *CredentialSyncInfo`

NewCredentialSyncInfoWithDefaults instantiates a new CredentialSyncInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *CredentialSyncInfo) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CredentialSyncInfo) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CredentialSyncInfo) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CredentialSyncInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorReason

`func (o *CredentialSyncInfo) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *CredentialSyncInfo) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *CredentialSyncInfo) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.

### HasErrorReason

`func (o *CredentialSyncInfo) HasErrorReason() bool`

HasErrorReason returns a boolean if a field has been set.

### GetSecretVersionId

`func (o *CredentialSyncInfo) GetSecretVersionId() string`

GetSecretVersionId returns the SecretVersionId field if non-nil, zero value otherwise.

### GetSecretVersionIdOk

`func (o *CredentialSyncInfo) GetSecretVersionIdOk() (*string, bool)`

GetSecretVersionIdOk returns a tuple with the SecretVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionId

`func (o *CredentialSyncInfo) SetSecretVersionId(v string)`

SetSecretVersionId sets SecretVersionId field to given value.

### HasSecretVersionId

`func (o *CredentialSyncInfo) HasSecretVersionId() bool`

HasSecretVersionId returns a boolean if a field has been set.

### GetSyncState

`func (o *CredentialSyncInfo) GetSyncState() string`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *CredentialSyncInfo) GetSyncStateOk() (*string, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *CredentialSyncInfo) SetSyncState(v string)`

SetSyncState sets SyncState field to given value.

### HasSyncState

`func (o *CredentialSyncInfo) HasSyncState() bool`

HasSyncState returns a boolean if a field has been set.

### GetSyncTime

`func (o *CredentialSyncInfo) GetSyncTime() time.Time`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *CredentialSyncInfo) GetSyncTimeOk() (*time.Time, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *CredentialSyncInfo) SetSyncTime(v time.Time)`

SetSyncTime sets SyncTime field to given value.

### HasSyncTime

`func (o *CredentialSyncInfo) HasSyncTime() bool`

HasSyncTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


