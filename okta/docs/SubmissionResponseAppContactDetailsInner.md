# SubmissionResponseAppContactDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactType** | **string** | Type of contact * &#x60;CUSTOMER_SUPPORT&#x60; - Public support contact details visible on your OIN catalog page for end users needing assistance with your integration. * &#x60;ESCALATION_SUPPORT&#x60; - Private support contact used by Okta to reach your organization during emergencies or escalations post-publication of the app (not shared with customers).  | 
**ContactValueType** | **string** | Format of the contact value | 
**Contact** | **string** | The contact value (email, phone, or URL) | 

## Methods

### NewSubmissionResponseAppContactDetailsInner

`func NewSubmissionResponseAppContactDetailsInner(contactType string, contactValueType string, contact string, ) *SubmissionResponseAppContactDetailsInner`

NewSubmissionResponseAppContactDetailsInner instantiates a new SubmissionResponseAppContactDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionResponseAppContactDetailsInnerWithDefaults

`func NewSubmissionResponseAppContactDetailsInnerWithDefaults() *SubmissionResponseAppContactDetailsInner`

NewSubmissionResponseAppContactDetailsInnerWithDefaults instantiates a new SubmissionResponseAppContactDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactType

`func (o *SubmissionResponseAppContactDetailsInner) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *SubmissionResponseAppContactDetailsInner) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *SubmissionResponseAppContactDetailsInner) SetContactType(v string)`

SetContactType sets ContactType field to given value.


### GetContactValueType

`func (o *SubmissionResponseAppContactDetailsInner) GetContactValueType() string`

GetContactValueType returns the ContactValueType field if non-nil, zero value otherwise.

### GetContactValueTypeOk

`func (o *SubmissionResponseAppContactDetailsInner) GetContactValueTypeOk() (*string, bool)`

GetContactValueTypeOk returns a tuple with the ContactValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactValueType

`func (o *SubmissionResponseAppContactDetailsInner) SetContactValueType(v string)`

SetContactValueType sets ContactValueType field to given value.


### GetContact

`func (o *SubmissionResponseAppContactDetailsInner) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SubmissionResponseAppContactDetailsInner) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SubmissionResponseAppContactDetailsInner) SetContact(v string)`

SetContact sets Contact field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


