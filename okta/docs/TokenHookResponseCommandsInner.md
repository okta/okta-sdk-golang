# TokenHookResponseCommandsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | One of the supported commands:   &#x60;com.okta.identity.patch&#x60;: Modify an ID token   &#x60;com.okta.access.patch&#x60;: Modify an access token &gt; **Note:** The &#x60;commands&#x60; array should only contain commands that can be applied to the requested tokens. For example, if only an ID token is requested, the &#x60;commands&#x60; array shouldn&#39;t contain commands of the type &#x60;com.okta.access.patch&#x60;. | [optional] 
**Value** | Pointer to [**[]TokenHookResponseCommandsInnerValueInner**](TokenHookResponseCommandsInnerValueInner.md) | The &#x60;value&#x60; object is where you specify the operation to perform. It&#39;s an array, which allows you to request more than one operation. | [optional] 

## Methods

### NewTokenHookResponseCommandsInner

`func NewTokenHookResponseCommandsInner() *TokenHookResponseCommandsInner`

NewTokenHookResponseCommandsInner instantiates a new TokenHookResponseCommandsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenHookResponseCommandsInnerWithDefaults

`func NewTokenHookResponseCommandsInnerWithDefaults() *TokenHookResponseCommandsInner`

NewTokenHookResponseCommandsInnerWithDefaults instantiates a new TokenHookResponseCommandsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TokenHookResponseCommandsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenHookResponseCommandsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenHookResponseCommandsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TokenHookResponseCommandsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *TokenHookResponseCommandsInner) GetValue() []TokenHookResponseCommandsInnerValueInner`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TokenHookResponseCommandsInner) GetValueOk() (*[]TokenHookResponseCommandsInnerValueInner, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TokenHookResponseCommandsInner) SetValue(v []TokenHookResponseCommandsInnerValueInner)`

SetValue sets Value field to given value.

### HasValue

`func (o *TokenHookResponseCommandsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


