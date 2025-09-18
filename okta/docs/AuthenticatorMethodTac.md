# AuthenticatorMethodTac

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type of authenticator method | [optional] 
**Links** | Pointer to [**LinksSelfAndLifecycle**](LinksSelfAndLifecycle.md) |  | [optional] 

## Methods

### NewAuthenticatorMethodTac

`func NewAuthenticatorMethodTac() *AuthenticatorMethodTac`

NewAuthenticatorMethodTac instantiates a new AuthenticatorMethodTac object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorMethodTacWithDefaults

`func NewAuthenticatorMethodTacWithDefaults() *AuthenticatorMethodTac`

NewAuthenticatorMethodTacWithDefaults instantiates a new AuthenticatorMethodTac object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AuthenticatorMethodTac) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthenticatorMethodTac) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthenticatorMethodTac) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthenticatorMethodTac) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *AuthenticatorMethodTac) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorMethodTac) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorMethodTac) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthenticatorMethodTac) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *AuthenticatorMethodTac) GetLinks() LinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthenticatorMethodTac) GetLinksOk() (*LinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthenticatorMethodTac) SetLinks(v LinksSelfAndLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthenticatorMethodTac) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


