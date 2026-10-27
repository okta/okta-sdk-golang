# AccountLinkedEnrollmentEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastAuthenticatedOn** | Pointer to **time.Time** | Timestamp of the last successful authentication | [optional] 
**OfflineFactors** | Pointer to [**[]OfflineFactor**](OfflineFactor.md) | Factors that are enrolled for offline authentication on this enrollment. &#x60;offlineFactors&#x60; is only returned for &#x60;desktop_mfa&#x60; enrollments when &#x60;offlineFactors&#x60; is requested with the &#x60;expand&#x60; query parameter. | [optional] 

## Methods

### NewAccountLinkedEnrollmentEmbedded

`func NewAccountLinkedEnrollmentEmbedded() *AccountLinkedEnrollmentEmbedded`

NewAccountLinkedEnrollmentEmbedded instantiates a new AccountLinkedEnrollmentEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountLinkedEnrollmentEmbeddedWithDefaults

`func NewAccountLinkedEnrollmentEmbeddedWithDefaults() *AccountLinkedEnrollmentEmbedded`

NewAccountLinkedEnrollmentEmbeddedWithDefaults instantiates a new AccountLinkedEnrollmentEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastAuthenticatedOn

`func (o *AccountLinkedEnrollmentEmbedded) GetLastAuthenticatedOn() time.Time`

GetLastAuthenticatedOn returns the LastAuthenticatedOn field if non-nil, zero value otherwise.

### GetLastAuthenticatedOnOk

`func (o *AccountLinkedEnrollmentEmbedded) GetLastAuthenticatedOnOk() (*time.Time, bool)`

GetLastAuthenticatedOnOk returns a tuple with the LastAuthenticatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthenticatedOn

`func (o *AccountLinkedEnrollmentEmbedded) SetLastAuthenticatedOn(v time.Time)`

SetLastAuthenticatedOn sets LastAuthenticatedOn field to given value.

### HasLastAuthenticatedOn

`func (o *AccountLinkedEnrollmentEmbedded) HasLastAuthenticatedOn() bool`

HasLastAuthenticatedOn returns a boolean if a field has been set.

### GetOfflineFactors

`func (o *AccountLinkedEnrollmentEmbedded) GetOfflineFactors() []OfflineFactor`

GetOfflineFactors returns the OfflineFactors field if non-nil, zero value otherwise.

### GetOfflineFactorsOk

`func (o *AccountLinkedEnrollmentEmbedded) GetOfflineFactorsOk() (*[]OfflineFactor, bool)`

GetOfflineFactorsOk returns a tuple with the OfflineFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineFactors

`func (o *AccountLinkedEnrollmentEmbedded) SetOfflineFactors(v []OfflineFactor)`

SetOfflineFactors sets OfflineFactors field to given value.

### HasOfflineFactors

`func (o *AccountLinkedEnrollmentEmbedded) HasOfflineFactors() bool`

HasOfflineFactors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


