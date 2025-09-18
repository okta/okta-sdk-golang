# TokenPayLoadData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**TokenPayLoadDataContext**](TokenPayLoadDataContext.md) |  | [optional] 
**Identity** | Pointer to [**BaseToken**](BaseToken.md) | Provides information on the properties of the ID token that Okta has generated, including the existing claims that it contains | [optional] 
**Access** | Pointer to [**TokenPayLoadDataAccess**](TokenPayLoadDataAccess.md) |  | [optional] 
**RefreshToken** | Pointer to [**RefreshToken**](RefreshToken.md) |  | [optional] 

## Methods

### NewTokenPayLoadData

`func NewTokenPayLoadData() *TokenPayLoadData`

NewTokenPayLoadData instantiates a new TokenPayLoadData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPayLoadDataWithDefaults

`func NewTokenPayLoadDataWithDefaults() *TokenPayLoadData`

NewTokenPayLoadDataWithDefaults instantiates a new TokenPayLoadData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *TokenPayLoadData) GetContext() TokenPayLoadDataContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TokenPayLoadData) GetContextOk() (*TokenPayLoadDataContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TokenPayLoadData) SetContext(v TokenPayLoadDataContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TokenPayLoadData) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetIdentity

`func (o *TokenPayLoadData) GetIdentity() BaseToken`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *TokenPayLoadData) GetIdentityOk() (*BaseToken, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *TokenPayLoadData) SetIdentity(v BaseToken)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *TokenPayLoadData) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetAccess

`func (o *TokenPayLoadData) GetAccess() TokenPayLoadDataAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *TokenPayLoadData) GetAccessOk() (*TokenPayLoadDataAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *TokenPayLoadData) SetAccess(v TokenPayLoadDataAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *TokenPayLoadData) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetRefreshToken

`func (o *TokenPayLoadData) GetRefreshToken() RefreshToken`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenPayLoadData) GetRefreshTokenOk() (*RefreshToken, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenPayLoadData) SetRefreshToken(v RefreshToken)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenPayLoadData) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


