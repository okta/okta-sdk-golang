# ServiceAccountDetailsOktaUserAccountSub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**OktaUserServiceAccountCredentials**](OktaUserServiceAccountCredentials.md) |  | [optional] 
**Email** | Pointer to **string** | The email address for the Okta user | [optional] [readonly] 
**OktaUserId** | **string** | The ID of the Okta user to manage as a service account | 

## Methods

### NewServiceAccountDetailsOktaUserAccountSub

`func NewServiceAccountDetailsOktaUserAccountSub(oktaUserId string, ) *ServiceAccountDetailsOktaUserAccountSub`

NewServiceAccountDetailsOktaUserAccountSub instantiates a new ServiceAccountDetailsOktaUserAccountSub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountDetailsOktaUserAccountSubWithDefaults

`func NewServiceAccountDetailsOktaUserAccountSubWithDefaults() *ServiceAccountDetailsOktaUserAccountSub`

NewServiceAccountDetailsOktaUserAccountSubWithDefaults instantiates a new ServiceAccountDetailsOktaUserAccountSub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ServiceAccountDetailsOktaUserAccountSub) GetCredentials() OktaUserServiceAccountCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ServiceAccountDetailsOktaUserAccountSub) GetCredentialsOk() (*OktaUserServiceAccountCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ServiceAccountDetailsOktaUserAccountSub) SetCredentials(v OktaUserServiceAccountCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ServiceAccountDetailsOktaUserAccountSub) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEmail

`func (o *ServiceAccountDetailsOktaUserAccountSub) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServiceAccountDetailsOktaUserAccountSub) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServiceAccountDetailsOktaUserAccountSub) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ServiceAccountDetailsOktaUserAccountSub) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOktaUserId

`func (o *ServiceAccountDetailsOktaUserAccountSub) GetOktaUserId() string`

GetOktaUserId returns the OktaUserId field if non-nil, zero value otherwise.

### GetOktaUserIdOk

`func (o *ServiceAccountDetailsOktaUserAccountSub) GetOktaUserIdOk() (*string, bool)`

GetOktaUserIdOk returns a tuple with the OktaUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaUserId

`func (o *ServiceAccountDetailsOktaUserAccountSub) SetOktaUserId(v string)`

SetOktaUserId sets OktaUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


