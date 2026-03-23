# DbscStartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**[]DbscCredential**](DbscCredential.md) |  | [readonly] 
**RefreshUrl** | **string** | URL to call for cookie refresh | [readonly] 
**Scope** | [**DbscScope**](DbscScope.md) |  | 
**SessionIdentifier** | **string** | The session identifier for this DBSC binding. Use this value in the &#x60;Sec-Secure-Session-Id&#x60; header for refresh requests. | [readonly] 

## Methods

### NewDbscStartResponse

`func NewDbscStartResponse(credentials []DbscCredential, refreshUrl string, scope DbscScope, sessionIdentifier string, ) *DbscStartResponse`

NewDbscStartResponse instantiates a new DbscStartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbscStartResponseWithDefaults

`func NewDbscStartResponseWithDefaults() *DbscStartResponse`

NewDbscStartResponseWithDefaults instantiates a new DbscStartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *DbscStartResponse) GetCredentials() []DbscCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DbscStartResponse) GetCredentialsOk() (*[]DbscCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DbscStartResponse) SetCredentials(v []DbscCredential)`

SetCredentials sets Credentials field to given value.


### GetRefreshUrl

`func (o *DbscStartResponse) GetRefreshUrl() string`

GetRefreshUrl returns the RefreshUrl field if non-nil, zero value otherwise.

### GetRefreshUrlOk

`func (o *DbscStartResponse) GetRefreshUrlOk() (*string, bool)`

GetRefreshUrlOk returns a tuple with the RefreshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshUrl

`func (o *DbscStartResponse) SetRefreshUrl(v string)`

SetRefreshUrl sets RefreshUrl field to given value.


### GetScope

`func (o *DbscStartResponse) GetScope() DbscScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DbscStartResponse) GetScopeOk() (*DbscScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DbscStartResponse) SetScope(v DbscScope)`

SetScope sets Scope field to given value.


### GetSessionIdentifier

`func (o *DbscStartResponse) GetSessionIdentifier() string`

GetSessionIdentifier returns the SessionIdentifier field if non-nil, zero value otherwise.

### GetSessionIdentifierOk

`func (o *DbscStartResponse) GetSessionIdentifierOk() (*string, bool)`

GetSessionIdentifierOk returns a tuple with the SessionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIdentifier

`func (o *DbscStartResponse) SetSessionIdentifier(v string)`

SetSessionIdentifier sets SessionIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


