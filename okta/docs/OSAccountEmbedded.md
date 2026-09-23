# OSAccountEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountLinkedEnrollments** | Pointer to [**[]AccountLinkedEnrollment**](AccountLinkedEnrollment.md) | Enrollments linked to this OS account | [optional] 
**Users** | Pointer to [**[]User**](User.md) | Users associated with this OS account | [optional] 

## Methods

### NewOSAccountEmbedded

`func NewOSAccountEmbedded() *OSAccountEmbedded`

NewOSAccountEmbedded instantiates a new OSAccountEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSAccountEmbeddedWithDefaults

`func NewOSAccountEmbeddedWithDefaults() *OSAccountEmbedded`

NewOSAccountEmbeddedWithDefaults instantiates a new OSAccountEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountLinkedEnrollments

`func (o *OSAccountEmbedded) GetAccountLinkedEnrollments() []AccountLinkedEnrollment`

GetAccountLinkedEnrollments returns the AccountLinkedEnrollments field if non-nil, zero value otherwise.

### GetAccountLinkedEnrollmentsOk

`func (o *OSAccountEmbedded) GetAccountLinkedEnrollmentsOk() (*[]AccountLinkedEnrollment, bool)`

GetAccountLinkedEnrollmentsOk returns a tuple with the AccountLinkedEnrollments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLinkedEnrollments

`func (o *OSAccountEmbedded) SetAccountLinkedEnrollments(v []AccountLinkedEnrollment)`

SetAccountLinkedEnrollments sets AccountLinkedEnrollments field to given value.

### HasAccountLinkedEnrollments

`func (o *OSAccountEmbedded) HasAccountLinkedEnrollments() bool`

HasAccountLinkedEnrollments returns a boolean if a field has been set.

### GetUsers

`func (o *OSAccountEmbedded) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *OSAccountEmbedded) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *OSAccountEmbedded) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *OSAccountEmbedded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


