# PasswordImportRequestDataContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**InlineHookRequestObject**](InlineHookRequestObject.md) |  | [optional] 
**Credential** | Pointer to [**PasswordImportRequestDataContextCredential**](PasswordImportRequestDataContextCredential.md) |  | [optional] 

## Methods

### NewPasswordImportRequestDataContext

`func NewPasswordImportRequestDataContext() *PasswordImportRequestDataContext`

NewPasswordImportRequestDataContext instantiates a new PasswordImportRequestDataContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordImportRequestDataContextWithDefaults

`func NewPasswordImportRequestDataContextWithDefaults() *PasswordImportRequestDataContext`

NewPasswordImportRequestDataContextWithDefaults instantiates a new PasswordImportRequestDataContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *PasswordImportRequestDataContext) GetRequest() InlineHookRequestObject`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *PasswordImportRequestDataContext) GetRequestOk() (*InlineHookRequestObject, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *PasswordImportRequestDataContext) SetRequest(v InlineHookRequestObject)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *PasswordImportRequestDataContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetCredential

`func (o *PasswordImportRequestDataContext) GetCredential() PasswordImportRequestDataContextCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *PasswordImportRequestDataContext) GetCredentialOk() (*PasswordImportRequestDataContextCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *PasswordImportRequestDataContext) SetCredential(v PasswordImportRequestDataContextCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *PasswordImportRequestDataContext) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


