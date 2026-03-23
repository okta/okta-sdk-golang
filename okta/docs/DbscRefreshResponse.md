# DbscRefreshResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**[]DbscCredential**](DbscCredential.md) |  | [readonly] 
**RefreshUrl** | **string** | URL to call for cookie refresh | [readonly] 
**Scope** | [**DbscScope**](DbscScope.md) |  | 
**SessionIdentifier** | **string** | The session identifier for this DBSC binding | [readonly] 

## Methods

### NewDbscRefreshResponse

`func NewDbscRefreshResponse(credentials []DbscCredential, refreshUrl string, scope DbscScope, sessionIdentifier string, ) *DbscRefreshResponse`

NewDbscRefreshResponse instantiates a new DbscRefreshResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbscRefreshResponseWithDefaults

`func NewDbscRefreshResponseWithDefaults() *DbscRefreshResponse`

NewDbscRefreshResponseWithDefaults instantiates a new DbscRefreshResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *DbscRefreshResponse) GetCredentials() []DbscCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DbscRefreshResponse) GetCredentialsOk() (*[]DbscCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DbscRefreshResponse) SetCredentials(v []DbscCredential)`

SetCredentials sets Credentials field to given value.


### GetRefreshUrl

`func (o *DbscRefreshResponse) GetRefreshUrl() string`

GetRefreshUrl returns the RefreshUrl field if non-nil, zero value otherwise.

### GetRefreshUrlOk

`func (o *DbscRefreshResponse) GetRefreshUrlOk() (*string, bool)`

GetRefreshUrlOk returns a tuple with the RefreshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshUrl

`func (o *DbscRefreshResponse) SetRefreshUrl(v string)`

SetRefreshUrl sets RefreshUrl field to given value.


### GetScope

`func (o *DbscRefreshResponse) GetScope() DbscScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DbscRefreshResponse) GetScopeOk() (*DbscScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DbscRefreshResponse) SetScope(v DbscScope)`

SetScope sets Scope field to given value.


### GetSessionIdentifier

`func (o *DbscRefreshResponse) GetSessionIdentifier() string`

GetSessionIdentifier returns the SessionIdentifier field if non-nil, zero value otherwise.

### GetSessionIdentifierOk

`func (o *DbscRefreshResponse) GetSessionIdentifierOk() (*string, bool)`

GetSessionIdentifierOk returns a tuple with the SessionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIdentifier

`func (o *DbscRefreshResponse) SetSessionIdentifier(v string)`

SetSessionIdentifier sets SessionIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


