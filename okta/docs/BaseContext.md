# BaseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**InlineHookRequestObject**](InlineHookRequestObject.md) |  | [optional] 
**Session** | Pointer to [**BaseContextSession**](BaseContextSession.md) |  | [optional] 
**User** | Pointer to [**BaseContextUser**](BaseContextUser.md) |  | [optional] 

## Methods

### NewBaseContext

`func NewBaseContext() *BaseContext`

NewBaseContext instantiates a new BaseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseContextWithDefaults

`func NewBaseContextWithDefaults() *BaseContext`

NewBaseContextWithDefaults instantiates a new BaseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *BaseContext) GetRequest() InlineHookRequestObject`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *BaseContext) GetRequestOk() (*InlineHookRequestObject, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *BaseContext) SetRequest(v InlineHookRequestObject)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *BaseContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSession

`func (o *BaseContext) GetSession() BaseContextSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *BaseContext) GetSessionOk() (*BaseContextSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *BaseContext) SetSession(v BaseContextSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *BaseContext) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetUser

`func (o *BaseContext) GetUser() BaseContextUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BaseContext) GetUserOk() (*BaseContextUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BaseContext) SetUser(v BaseContextUser)`

SetUser sets User field to given value.

### HasUser

`func (o *BaseContext) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


