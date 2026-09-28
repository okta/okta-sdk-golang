# CredentialVerificationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **time.Time** | Timestamp that indicates when the latest verification job reached a terminal state (&#x60;COMPLETED&#x60; or &#x60;FAILED&#x60;). Absent while &#x60;verificationState&#x60; is set to &#x60;NOT_INITIALIZED&#x60; or &#x60;IN_PROGRESS&#x60;. | [optional] [readonly] 
**CredentialValid** | Pointer to **NullableBool** | Indicates whether the credentials were verified successfully. &#x60;true&#x60; if the credentials are valid against the downstream application and &#x60;false&#x60; if they are invalid. Present only when &#x60;verificationState&#x60; is &#x60;COMPLETED&#x60;. | [optional] [readonly] 
**ErrorCode** | Pointer to **NullableString** | Indicates why credential validation didn&#39;t succeed. Present when &#x60;verificationState&#x60; is &#x60;FAILED&#x60;, or when &#x60;verificationState&#x60; is &#x60;COMPLETED&#x60; and &#x60;credentialValid&#x60; is &#x60;false&#x60;. | [optional] [readonly] 
**ErrorReason** | Pointer to **NullableString** | Description of the error | [optional] [readonly] 
**RequestId** | Pointer to **NullableString** | Unique identifier of the most recent verification job initiated for this privileged resource | [optional] [readonly] 
**VerificationState** | Pointer to **string** | Current state of the latest credential verification job for the privileged resource | [optional] [readonly] 

## Methods

### NewCredentialVerificationInfo

`func NewCredentialVerificationInfo() *CredentialVerificationInfo`

NewCredentialVerificationInfo instantiates a new CredentialVerificationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialVerificationInfoWithDefaults

`func NewCredentialVerificationInfoWithDefaults() *CredentialVerificationInfo`

NewCredentialVerificationInfoWithDefaults instantiates a new CredentialVerificationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *CredentialVerificationInfo) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *CredentialVerificationInfo) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *CredentialVerificationInfo) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *CredentialVerificationInfo) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetCredentialValid

`func (o *CredentialVerificationInfo) GetCredentialValid() bool`

GetCredentialValid returns the CredentialValid field if non-nil, zero value otherwise.

### GetCredentialValidOk

`func (o *CredentialVerificationInfo) GetCredentialValidOk() (*bool, bool)`

GetCredentialValidOk returns a tuple with the CredentialValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialValid

`func (o *CredentialVerificationInfo) SetCredentialValid(v bool)`

SetCredentialValid sets CredentialValid field to given value.

### HasCredentialValid

`func (o *CredentialVerificationInfo) HasCredentialValid() bool`

HasCredentialValid returns a boolean if a field has been set.

### SetCredentialValidNil

`func (o *CredentialVerificationInfo) SetCredentialValidNil(b bool)`

 SetCredentialValidNil sets the value for CredentialValid to be an explicit nil

### UnsetCredentialValid
`func (o *CredentialVerificationInfo) UnsetCredentialValid()`

UnsetCredentialValid ensures that no value is present for CredentialValid, not even an explicit nil
### GetErrorCode

`func (o *CredentialVerificationInfo) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CredentialVerificationInfo) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CredentialVerificationInfo) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CredentialVerificationInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *CredentialVerificationInfo) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *CredentialVerificationInfo) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorReason

`func (o *CredentialVerificationInfo) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *CredentialVerificationInfo) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *CredentialVerificationInfo) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.

### HasErrorReason

`func (o *CredentialVerificationInfo) HasErrorReason() bool`

HasErrorReason returns a boolean if a field has been set.

### SetErrorReasonNil

`func (o *CredentialVerificationInfo) SetErrorReasonNil(b bool)`

 SetErrorReasonNil sets the value for ErrorReason to be an explicit nil

### UnsetErrorReason
`func (o *CredentialVerificationInfo) UnsetErrorReason()`

UnsetErrorReason ensures that no value is present for ErrorReason, not even an explicit nil
### GetRequestId

`func (o *CredentialVerificationInfo) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CredentialVerificationInfo) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CredentialVerificationInfo) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CredentialVerificationInfo) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *CredentialVerificationInfo) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *CredentialVerificationInfo) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetVerificationState

`func (o *CredentialVerificationInfo) GetVerificationState() string`

GetVerificationState returns the VerificationState field if non-nil, zero value otherwise.

### GetVerificationStateOk

`func (o *CredentialVerificationInfo) GetVerificationStateOk() (*string, bool)`

GetVerificationStateOk returns a tuple with the VerificationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationState

`func (o *CredentialVerificationInfo) SetVerificationState(v string)`

SetVerificationState sets VerificationState field to given value.

### HasVerificationState

`func (o *CredentialVerificationInfo) HasVerificationState() bool`

HasVerificationState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


