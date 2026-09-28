# AccountLinkedEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the enrollment was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier of the enrollment | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the enrollment was last updated | [optional] [readonly] 
**Profile** | Pointer to [**AccountLinkedEnrollmentProfile**](AccountLinkedEnrollmentProfile.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the enrollment. Possible values depend on &#x60;type&#x60;. For &#x60;platform_sso&#x60; - &#x60;ACTIVE&#x60;, &#x60;SUSPENDED&#x60;, &#x60;REVOKED&#x60;. For &#x60;desktop_mfa&#x60; - &#x60;ACTIVE&#x60;, &#x60;DELETED&#x60;. | [optional] 
**Type** | Pointer to **string** | Type of the linked enrollment | [optional] 
**Embedded** | Pointer to [**AccountLinkedEnrollmentEmbedded**](AccountLinkedEnrollmentEmbedded.md) |  | [optional] 

## Methods

### NewAccountLinkedEnrollment

`func NewAccountLinkedEnrollment() *AccountLinkedEnrollment`

NewAccountLinkedEnrollment instantiates a new AccountLinkedEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountLinkedEnrollmentWithDefaults

`func NewAccountLinkedEnrollmentWithDefaults() *AccountLinkedEnrollment`

NewAccountLinkedEnrollmentWithDefaults instantiates a new AccountLinkedEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AccountLinkedEnrollment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountLinkedEnrollment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountLinkedEnrollment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccountLinkedEnrollment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AccountLinkedEnrollment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountLinkedEnrollment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountLinkedEnrollment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountLinkedEnrollment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AccountLinkedEnrollment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AccountLinkedEnrollment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AccountLinkedEnrollment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AccountLinkedEnrollment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *AccountLinkedEnrollment) GetProfile() AccountLinkedEnrollmentProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AccountLinkedEnrollment) GetProfileOk() (*AccountLinkedEnrollmentProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AccountLinkedEnrollment) SetProfile(v AccountLinkedEnrollmentProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AccountLinkedEnrollment) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStatus

`func (o *AccountLinkedEnrollment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountLinkedEnrollment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountLinkedEnrollment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountLinkedEnrollment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *AccountLinkedEnrollment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountLinkedEnrollment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountLinkedEnrollment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountLinkedEnrollment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *AccountLinkedEnrollment) GetEmbedded() AccountLinkedEnrollmentEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AccountLinkedEnrollment) GetEmbeddedOk() (*AccountLinkedEnrollmentEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AccountLinkedEnrollment) SetEmbedded(v AccountLinkedEnrollmentEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AccountLinkedEnrollment) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


