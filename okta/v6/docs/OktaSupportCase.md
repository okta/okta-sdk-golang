# OktaSupportCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseNumber** | Pointer to **string** | Okta Support case number | [optional] [readonly] 
**Impersonation** | Pointer to [**OktaSupportCaseImpersonation**](OktaSupportCaseImpersonation.md) |  | [optional] 
**SelfAssigned** | Pointer to [**OktaSupportCaseSelfAssigned**](OktaSupportCaseSelfAssigned.md) |  | [optional] 
**Subject** | Pointer to **string** | Subject of the support case | [optional] [readonly] 

## Methods

### NewOktaSupportCase

`func NewOktaSupportCase() *OktaSupportCase`

NewOktaSupportCase instantiates a new OktaSupportCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaSupportCaseWithDefaults

`func NewOktaSupportCaseWithDefaults() *OktaSupportCase`

NewOktaSupportCaseWithDefaults instantiates a new OktaSupportCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseNumber

`func (o *OktaSupportCase) GetCaseNumber() string`

GetCaseNumber returns the CaseNumber field if non-nil, zero value otherwise.

### GetCaseNumberOk

`func (o *OktaSupportCase) GetCaseNumberOk() (*string, bool)`

GetCaseNumberOk returns a tuple with the CaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseNumber

`func (o *OktaSupportCase) SetCaseNumber(v string)`

SetCaseNumber sets CaseNumber field to given value.

### HasCaseNumber

`func (o *OktaSupportCase) HasCaseNumber() bool`

HasCaseNumber returns a boolean if a field has been set.

### GetImpersonation

`func (o *OktaSupportCase) GetImpersonation() OktaSupportCaseImpersonation`

GetImpersonation returns the Impersonation field if non-nil, zero value otherwise.

### GetImpersonationOk

`func (o *OktaSupportCase) GetImpersonationOk() (*OktaSupportCaseImpersonation, bool)`

GetImpersonationOk returns a tuple with the Impersonation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonation

`func (o *OktaSupportCase) SetImpersonation(v OktaSupportCaseImpersonation)`

SetImpersonation sets Impersonation field to given value.

### HasImpersonation

`func (o *OktaSupportCase) HasImpersonation() bool`

HasImpersonation returns a boolean if a field has been set.

### GetSelfAssigned

`func (o *OktaSupportCase) GetSelfAssigned() OktaSupportCaseSelfAssigned`

GetSelfAssigned returns the SelfAssigned field if non-nil, zero value otherwise.

### GetSelfAssignedOk

`func (o *OktaSupportCase) GetSelfAssignedOk() (*OktaSupportCaseSelfAssigned, bool)`

GetSelfAssignedOk returns a tuple with the SelfAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfAssigned

`func (o *OktaSupportCase) SetSelfAssigned(v OktaSupportCaseSelfAssigned)`

SetSelfAssigned sets SelfAssigned field to given value.

### HasSelfAssigned

`func (o *OktaSupportCase) HasSelfAssigned() bool`

HasSelfAssigned returns a boolean if a field has been set.

### GetSubject

`func (o *OktaSupportCase) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OktaSupportCase) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OktaSupportCase) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OktaSupportCase) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


