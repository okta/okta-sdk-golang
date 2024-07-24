# SecurityEventTokenRequestJwtBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aud** | **string** | Audience | 
**Events** | [**SecurityEventTokenRequestJwtEvents**](SecurityEventTokenRequestJwtEvents.md) |  | 
**Iat** | **int64** | Token issue time (UNIX timestamp) | 
**Iss** | **string** | Token issuer | 
**Jti** | **string** | Token ID | 

## Methods

### NewSecurityEventTokenRequestJwtBody

`func NewSecurityEventTokenRequestJwtBody(aud string, events SecurityEventTokenRequestJwtEvents, iat int64, iss string, jti string, ) *SecurityEventTokenRequestJwtBody`

NewSecurityEventTokenRequestJwtBody instantiates a new SecurityEventTokenRequestJwtBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventTokenRequestJwtBodyWithDefaults

`func NewSecurityEventTokenRequestJwtBodyWithDefaults() *SecurityEventTokenRequestJwtBody`

NewSecurityEventTokenRequestJwtBodyWithDefaults instantiates a new SecurityEventTokenRequestJwtBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *SecurityEventTokenRequestJwtBody) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *SecurityEventTokenRequestJwtBody) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *SecurityEventTokenRequestJwtBody) SetAud(v string)`

SetAud sets Aud field to given value.


### GetEvents

`func (o *SecurityEventTokenRequestJwtBody) GetEvents() SecurityEventTokenRequestJwtEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SecurityEventTokenRequestJwtBody) GetEventsOk() (*SecurityEventTokenRequestJwtEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SecurityEventTokenRequestJwtBody) SetEvents(v SecurityEventTokenRequestJwtEvents)`

SetEvents sets Events field to given value.


### GetIat

`func (o *SecurityEventTokenRequestJwtBody) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *SecurityEventTokenRequestJwtBody) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *SecurityEventTokenRequestJwtBody) SetIat(v int64)`

SetIat sets Iat field to given value.


### GetIss

`func (o *SecurityEventTokenRequestJwtBody) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *SecurityEventTokenRequestJwtBody) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *SecurityEventTokenRequestJwtBody) SetIss(v string)`

SetIss sets Iss field to given value.


### GetJti

`func (o *SecurityEventTokenRequestJwtBody) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *SecurityEventTokenRequestJwtBody) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *SecurityEventTokenRequestJwtBody) SetJti(v string)`

SetJti sets Jti field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


