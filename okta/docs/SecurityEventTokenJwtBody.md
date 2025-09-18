# SecurityEventTokenJwtBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aud** | **string** | Audience | 
**Events** | [**SecurityEventTokenJwtEvents**](SecurityEventTokenJwtEvents.md) |  | 
**Iat** | **int64** | Token issue time (UNIX timestamp) | 
**Iss** | **string** | Token issuer | 
**Jti** | **string** | Token ID | 

## Methods

### NewSecurityEventTokenJwtBody

`func NewSecurityEventTokenJwtBody(aud string, events SecurityEventTokenJwtEvents, iat int64, iss string, jti string, ) *SecurityEventTokenJwtBody`

NewSecurityEventTokenJwtBody instantiates a new SecurityEventTokenJwtBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventTokenJwtBodyWithDefaults

`func NewSecurityEventTokenJwtBodyWithDefaults() *SecurityEventTokenJwtBody`

NewSecurityEventTokenJwtBodyWithDefaults instantiates a new SecurityEventTokenJwtBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *SecurityEventTokenJwtBody) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *SecurityEventTokenJwtBody) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *SecurityEventTokenJwtBody) SetAud(v string)`

SetAud sets Aud field to given value.


### GetEvents

`func (o *SecurityEventTokenJwtBody) GetEvents() SecurityEventTokenJwtEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SecurityEventTokenJwtBody) GetEventsOk() (*SecurityEventTokenJwtEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SecurityEventTokenJwtBody) SetEvents(v SecurityEventTokenJwtEvents)`

SetEvents sets Events field to given value.


### GetIat

`func (o *SecurityEventTokenJwtBody) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *SecurityEventTokenJwtBody) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *SecurityEventTokenJwtBody) SetIat(v int64)`

SetIat sets Iat field to given value.


### GetIss

`func (o *SecurityEventTokenJwtBody) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *SecurityEventTokenJwtBody) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *SecurityEventTokenJwtBody) SetIss(v string)`

SetIss sets Iss field to given value.


### GetJti

`func (o *SecurityEventTokenJwtBody) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *SecurityEventTokenJwtBody) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *SecurityEventTokenJwtBody) SetJti(v string)`

SetJti sets Jti field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


