# GlobalTokenRevocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethod** | **string** | Authentication method &lt;br&gt; **Note:** Currently, only the &#x60;SIGNED_JWT&#x60; method is supported | 
**Endpoint** | **string** | URL of the authorization server&#39;s global token revocation endpoint | 
**PartialLogout** | Pointer to **bool** | Allow partial support for Universal Logout | [optional] [default to false]
**SubjectFormat** | **string** | The format of the subject | 

## Methods

### NewGlobalTokenRevocation

`func NewGlobalTokenRevocation(authMethod string, endpoint string, subjectFormat string, ) *GlobalTokenRevocation`

NewGlobalTokenRevocation instantiates a new GlobalTokenRevocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalTokenRevocationWithDefaults

`func NewGlobalTokenRevocationWithDefaults() *GlobalTokenRevocation`

NewGlobalTokenRevocationWithDefaults instantiates a new GlobalTokenRevocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethod

`func (o *GlobalTokenRevocation) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *GlobalTokenRevocation) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *GlobalTokenRevocation) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetEndpoint

`func (o *GlobalTokenRevocation) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GlobalTokenRevocation) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GlobalTokenRevocation) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetPartialLogout

`func (o *GlobalTokenRevocation) GetPartialLogout() bool`

GetPartialLogout returns the PartialLogout field if non-nil, zero value otherwise.

### GetPartialLogoutOk

`func (o *GlobalTokenRevocation) GetPartialLogoutOk() (*bool, bool)`

GetPartialLogoutOk returns a tuple with the PartialLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialLogout

`func (o *GlobalTokenRevocation) SetPartialLogout(v bool)`

SetPartialLogout sets PartialLogout field to given value.

### HasPartialLogout

`func (o *GlobalTokenRevocation) HasPartialLogout() bool`

HasPartialLogout returns a boolean if a field has been set.

### GetSubjectFormat

`func (o *GlobalTokenRevocation) GetSubjectFormat() string`

GetSubjectFormat returns the SubjectFormat field if non-nil, zero value otherwise.

### GetSubjectFormatOk

`func (o *GlobalTokenRevocation) GetSubjectFormatOk() (*string, bool)`

GetSubjectFormatOk returns a tuple with the SubjectFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectFormat

`func (o *GlobalTokenRevocation) SetSubjectFormat(v string)`

SetSubjectFormat sets SubjectFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


