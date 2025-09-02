# SAMLPayLoadDataContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**InlineHookRequestObject**](InlineHookRequestObject.md) |  | [optional] 
**Session** | Pointer to [**BaseContextSession**](BaseContextSession.md) |  | [optional] 
**User** | Pointer to [**BaseContextUser**](BaseContextUser.md) |  | [optional] 
**Protocol** | Pointer to [**SAMLPayLoadDataContextAllOfProtocol**](SAMLPayLoadDataContextAllOfProtocol.md) |  | [optional] 

## Methods

### NewSAMLPayLoadDataContext

`func NewSAMLPayLoadDataContext() *SAMLPayLoadDataContext`

NewSAMLPayLoadDataContext instantiates a new SAMLPayLoadDataContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPayLoadDataContextWithDefaults

`func NewSAMLPayLoadDataContextWithDefaults() *SAMLPayLoadDataContext`

NewSAMLPayLoadDataContextWithDefaults instantiates a new SAMLPayLoadDataContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *SAMLPayLoadDataContext) GetRequest() InlineHookRequestObject`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *SAMLPayLoadDataContext) GetRequestOk() (*InlineHookRequestObject, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *SAMLPayLoadDataContext) SetRequest(v InlineHookRequestObject)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *SAMLPayLoadDataContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSession

`func (o *SAMLPayLoadDataContext) GetSession() BaseContextSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SAMLPayLoadDataContext) GetSessionOk() (*BaseContextSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SAMLPayLoadDataContext) SetSession(v BaseContextSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *SAMLPayLoadDataContext) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetUser

`func (o *SAMLPayLoadDataContext) GetUser() BaseContextUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SAMLPayLoadDataContext) GetUserOk() (*BaseContextUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SAMLPayLoadDataContext) SetUser(v BaseContextUser)`

SetUser sets User field to given value.

### HasUser

`func (o *SAMLPayLoadDataContext) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetProtocol

`func (o *SAMLPayLoadDataContext) GetProtocol() SAMLPayLoadDataContextAllOfProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SAMLPayLoadDataContext) GetProtocolOk() (*SAMLPayLoadDataContextAllOfProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SAMLPayLoadDataContext) SetProtocol(v SAMLPayLoadDataContextAllOfProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SAMLPayLoadDataContext) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


