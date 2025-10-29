# PolicySubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Optional [regular expression pattern](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Regular_expressions) used to filter untrusted IdP usernames. * As a best security practice, you should define a regular expression pattern to filter untrusted IdP usernames. This is especially important if multiple IdPs are connected to your org. The filter prevents an IdP from issuing an assertion for any user, including partners or directory users in your Okta org. * For example, the filter pattern &#x60;(\\S+@example\\.com)&#x60; allows only Users that have an &#x60;@example.com&#x60; username suffix. It rejects assertions that have any other suffix such as &#x60;@corp.example.com&#x60; or &#x60;@partner.com&#x60;. * Only &#x60;SAML2&#x60; and &#x60;OIDC&#x60; IdP providers support the &#x60;filter&#x60; property. | [optional] 
**MatchAttribute** | Pointer to **string** | Okta user profile attribute for matching a transformed IdP username. Only for matchType &#x60;CUSTOM_ATTRIBUTE&#x60;. The &#x60;matchAttribute&#x60; must be a valid Okta user profile attribute of one of the following types: * String (with no format or &#39;email&#39; format only) * Integer * Number | [optional] 
**MatchType** | Pointer to **string** | Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username | [optional] 
**UserNameTemplate** | Pointer to [**PolicyUserNameTemplate**](PolicyUserNameTemplate.md) |  | [optional] 

## Methods

### NewPolicySubject

`func NewPolicySubject() *PolicySubject`

NewPolicySubject instantiates a new PolicySubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySubjectWithDefaults

`func NewPolicySubjectWithDefaults() *PolicySubject`

NewPolicySubjectWithDefaults instantiates a new PolicySubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *PolicySubject) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PolicySubject) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PolicySubject) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PolicySubject) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMatchAttribute

`func (o *PolicySubject) GetMatchAttribute() string`

GetMatchAttribute returns the MatchAttribute field if non-nil, zero value otherwise.

### GetMatchAttributeOk

`func (o *PolicySubject) GetMatchAttributeOk() (*string, bool)`

GetMatchAttributeOk returns a tuple with the MatchAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAttribute

`func (o *PolicySubject) SetMatchAttribute(v string)`

SetMatchAttribute sets MatchAttribute field to given value.

### HasMatchAttribute

`func (o *PolicySubject) HasMatchAttribute() bool`

HasMatchAttribute returns a boolean if a field has been set.

### GetMatchType

`func (o *PolicySubject) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *PolicySubject) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *PolicySubject) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *PolicySubject) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetUserNameTemplate

`func (o *PolicySubject) GetUserNameTemplate() PolicyUserNameTemplate`

GetUserNameTemplate returns the UserNameTemplate field if non-nil, zero value otherwise.

### GetUserNameTemplateOk

`func (o *PolicySubject) GetUserNameTemplateOk() (*PolicyUserNameTemplate, bool)`

GetUserNameTemplateOk returns a tuple with the UserNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameTemplate

`func (o *PolicySubject) SetUserNameTemplate(v PolicyUserNameTemplate)`

SetUserNameTemplate sets UserNameTemplate field to given value.

### HasUserNameTemplate

`func (o *PolicySubject) HasUserNameTemplate() bool`

HasUserNameTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


