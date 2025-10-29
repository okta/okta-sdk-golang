# TokenProtocolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The ID of the client associated with the token | [optional] 
**GrantType** | Pointer to **string** | Determines the mechanism Okta uses to authorize the creation of the tokens. | [optional] 
**RedirectUri** | Pointer to **string** | Specifies the callback location where the authorization was sent | [optional] 
**ResponseMode** | Pointer to **string** | The authorization response mode | [optional] 
**ResponseType** | Pointer to **string** | The authorization response type | [optional] 
**Scope** | Pointer to **string** | The scopes requested | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewTokenProtocolRequest

`func NewTokenProtocolRequest() *TokenProtocolRequest`

NewTokenProtocolRequest instantiates a new TokenProtocolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenProtocolRequestWithDefaults

`func NewTokenProtocolRequestWithDefaults() *TokenProtocolRequest`

NewTokenProtocolRequestWithDefaults instantiates a new TokenProtocolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TokenProtocolRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenProtocolRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenProtocolRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TokenProtocolRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetGrantType

`func (o *TokenProtocolRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *TokenProtocolRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *TokenProtocolRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *TokenProtocolRequest) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetRedirectUri

`func (o *TokenProtocolRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *TokenProtocolRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *TokenProtocolRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *TokenProtocolRequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetResponseMode

`func (o *TokenProtocolRequest) GetResponseMode() string`

GetResponseMode returns the ResponseMode field if non-nil, zero value otherwise.

### GetResponseModeOk

`func (o *TokenProtocolRequest) GetResponseModeOk() (*string, bool)`

GetResponseModeOk returns a tuple with the ResponseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMode

`func (o *TokenProtocolRequest) SetResponseMode(v string)`

SetResponseMode sets ResponseMode field to given value.

### HasResponseMode

`func (o *TokenProtocolRequest) HasResponseMode() bool`

HasResponseMode returns a boolean if a field has been set.

### GetResponseType

`func (o *TokenProtocolRequest) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *TokenProtocolRequest) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *TokenProtocolRequest) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *TokenProtocolRequest) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetScope

`func (o *TokenProtocolRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TokenProtocolRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TokenProtocolRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *TokenProtocolRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetState

`func (o *TokenProtocolRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TokenProtocolRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TokenProtocolRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TokenProtocolRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


