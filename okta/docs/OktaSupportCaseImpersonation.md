# OktaSupportCaseImpersonation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of Okta Support access | [optional] 
**Expiration** | Pointer to **NullableTime** | Expiration date of Okta Support access | [optional] 

## Methods

### NewOktaSupportCaseImpersonation

`func NewOktaSupportCaseImpersonation() *OktaSupportCaseImpersonation`

NewOktaSupportCaseImpersonation instantiates a new OktaSupportCaseImpersonation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaSupportCaseImpersonationWithDefaults

`func NewOktaSupportCaseImpersonationWithDefaults() *OktaSupportCaseImpersonation`

NewOktaSupportCaseImpersonationWithDefaults instantiates a new OktaSupportCaseImpersonation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *OktaSupportCaseImpersonation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OktaSupportCaseImpersonation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OktaSupportCaseImpersonation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OktaSupportCaseImpersonation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiration

`func (o *OktaSupportCaseImpersonation) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *OktaSupportCaseImpersonation) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *OktaSupportCaseImpersonation) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *OktaSupportCaseImpersonation) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *OktaSupportCaseImpersonation) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *OktaSupportCaseImpersonation) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


