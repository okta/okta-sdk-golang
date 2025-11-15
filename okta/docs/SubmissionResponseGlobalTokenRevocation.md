# SubmissionResponseGlobalTokenRevocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | URL of the authorization server&#39;s global token revocation endpoint | 
**SubjectFormat** | **string** | The format of the subject | 
**AuthMethod** | **string** | Authentication method &lt;br&gt; **Note:** Currently, only the &#x60;SIGNED_JWT&#x60; method is supported. | 
**PartialLogout** | Pointer to **bool** | Allow partial support for Universal Logout | [optional] [default to false]

## Methods

### NewSubmissionResponseGlobalTokenRevocation

`func NewSubmissionResponseGlobalTokenRevocation(endpoint string, subjectFormat string, authMethod string, ) *SubmissionResponseGlobalTokenRevocation`

NewSubmissionResponseGlobalTokenRevocation instantiates a new SubmissionResponseGlobalTokenRevocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionResponseGlobalTokenRevocationWithDefaults

`func NewSubmissionResponseGlobalTokenRevocationWithDefaults() *SubmissionResponseGlobalTokenRevocation`

NewSubmissionResponseGlobalTokenRevocationWithDefaults instantiates a new SubmissionResponseGlobalTokenRevocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *SubmissionResponseGlobalTokenRevocation) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SubmissionResponseGlobalTokenRevocation) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SubmissionResponseGlobalTokenRevocation) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetSubjectFormat

`func (o *SubmissionResponseGlobalTokenRevocation) GetSubjectFormat() string`

GetSubjectFormat returns the SubjectFormat field if non-nil, zero value otherwise.

### GetSubjectFormatOk

`func (o *SubmissionResponseGlobalTokenRevocation) GetSubjectFormatOk() (*string, bool)`

GetSubjectFormatOk returns a tuple with the SubjectFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectFormat

`func (o *SubmissionResponseGlobalTokenRevocation) SetSubjectFormat(v string)`

SetSubjectFormat sets SubjectFormat field to given value.


### GetAuthMethod

`func (o *SubmissionResponseGlobalTokenRevocation) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *SubmissionResponseGlobalTokenRevocation) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *SubmissionResponseGlobalTokenRevocation) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetPartialLogout

`func (o *SubmissionResponseGlobalTokenRevocation) GetPartialLogout() bool`

GetPartialLogout returns the PartialLogout field if non-nil, zero value otherwise.

### GetPartialLogoutOk

`func (o *SubmissionResponseGlobalTokenRevocation) GetPartialLogoutOk() (*bool, bool)`

GetPartialLogoutOk returns a tuple with the PartialLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialLogout

`func (o *SubmissionResponseGlobalTokenRevocation) SetPartialLogout(v bool)`

SetPartialLogout sets PartialLogout field to given value.

### HasPartialLogout

`func (o *SubmissionResponseGlobalTokenRevocation) HasPartialLogout() bool`

HasPartialLogout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


