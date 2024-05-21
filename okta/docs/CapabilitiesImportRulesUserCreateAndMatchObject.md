# CapabilitiesImportRulesUserCreateAndMatchObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPartialMatch** | Pointer to **bool** | Allows user import upon partial matching. Partial matching occurs when the first and last names of an imported user match those of an existing Okta user, even if the username or email attributes don&#39;t match. | [optional] 
**AutoActivateNewUsers** | Pointer to **bool** | If set to &#x60;true&#x60;, imported new users are automatically activated. | [optional] 
**AutoConfirmExactMatch** | Pointer to **bool** | If set to &#x60;true&#x60;, exact-matched users are automatically confirmed on activation. If set to &#x60;false&#x60;, exact-matched users need to be confirmed manually. | [optional] 
**AutoConfirmNewUsers** | Pointer to **bool** | If set to &#x60;true&#x60;, imported new users are automatically confirmed on activation. This doesn&#39;t apply to imported users that already exist in Okta. | [optional] 
**AutoConfirmPartialMatch** | Pointer to **bool** | If set to &#x60;true&#x60;, partially matched users are automatically confirmed on activation. If set to &#x60;false&#x60;, partially matched users need to be confirmed manually. | [optional] 
**ExactMatchCriteria** | Pointer to **string** | Determines the attribute to match users | [optional] 

## Methods

### NewCapabilitiesImportRulesUserCreateAndMatchObject

`func NewCapabilitiesImportRulesUserCreateAndMatchObject() *CapabilitiesImportRulesUserCreateAndMatchObject`

NewCapabilitiesImportRulesUserCreateAndMatchObject instantiates a new CapabilitiesImportRulesUserCreateAndMatchObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitiesImportRulesUserCreateAndMatchObjectWithDefaults

`func NewCapabilitiesImportRulesUserCreateAndMatchObjectWithDefaults() *CapabilitiesImportRulesUserCreateAndMatchObject`

NewCapabilitiesImportRulesUserCreateAndMatchObjectWithDefaults instantiates a new CapabilitiesImportRulesUserCreateAndMatchObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPartialMatch

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAllowPartialMatch() bool`

GetAllowPartialMatch returns the AllowPartialMatch field if non-nil, zero value otherwise.

### GetAllowPartialMatchOk

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAllowPartialMatchOk() (*bool, bool)`

GetAllowPartialMatchOk returns a tuple with the AllowPartialMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPartialMatch

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetAllowPartialMatch(v bool)`

SetAllowPartialMatch sets AllowPartialMatch field to given value.

### HasAllowPartialMatch

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAllowPartialMatch() bool`

HasAllowPartialMatch returns a boolean if a field has been set.

### GetAutoActivateNewUsers

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoActivateNewUsers() bool`

GetAutoActivateNewUsers returns the AutoActivateNewUsers field if non-nil, zero value otherwise.

### GetAutoActivateNewUsersOk

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoActivateNewUsersOk() (*bool, bool)`

GetAutoActivateNewUsersOk returns a tuple with the AutoActivateNewUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoActivateNewUsers

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetAutoActivateNewUsers(v bool)`

SetAutoActivateNewUsers sets AutoActivateNewUsers field to given value.

### HasAutoActivateNewUsers

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoActivateNewUsers() bool`

HasAutoActivateNewUsers returns a boolean if a field has been set.

### GetAutoConfirmExactMatch

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmExactMatch() bool`

GetAutoConfirmExactMatch returns the AutoConfirmExactMatch field if non-nil, zero value otherwise.

### GetAutoConfirmExactMatchOk

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmExactMatchOk() (*bool, bool)`

GetAutoConfirmExactMatchOk returns a tuple with the AutoConfirmExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfirmExactMatch

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetAutoConfirmExactMatch(v bool)`

SetAutoConfirmExactMatch sets AutoConfirmExactMatch field to given value.

### HasAutoConfirmExactMatch

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoConfirmExactMatch() bool`

HasAutoConfirmExactMatch returns a boolean if a field has been set.

### GetAutoConfirmNewUsers

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmNewUsers() bool`

GetAutoConfirmNewUsers returns the AutoConfirmNewUsers field if non-nil, zero value otherwise.

### GetAutoConfirmNewUsersOk

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmNewUsersOk() (*bool, bool)`

GetAutoConfirmNewUsersOk returns a tuple with the AutoConfirmNewUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfirmNewUsers

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetAutoConfirmNewUsers(v bool)`

SetAutoConfirmNewUsers sets AutoConfirmNewUsers field to given value.

### HasAutoConfirmNewUsers

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoConfirmNewUsers() bool`

HasAutoConfirmNewUsers returns a boolean if a field has been set.

### GetAutoConfirmPartialMatch

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmPartialMatch() bool`

GetAutoConfirmPartialMatch returns the AutoConfirmPartialMatch field if non-nil, zero value otherwise.

### GetAutoConfirmPartialMatchOk

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetAutoConfirmPartialMatchOk() (*bool, bool)`

GetAutoConfirmPartialMatchOk returns a tuple with the AutoConfirmPartialMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfirmPartialMatch

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetAutoConfirmPartialMatch(v bool)`

SetAutoConfirmPartialMatch sets AutoConfirmPartialMatch field to given value.

### HasAutoConfirmPartialMatch

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasAutoConfirmPartialMatch() bool`

HasAutoConfirmPartialMatch returns a boolean if a field has been set.

### GetExactMatchCriteria

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetExactMatchCriteria() string`

GetExactMatchCriteria returns the ExactMatchCriteria field if non-nil, zero value otherwise.

### GetExactMatchCriteriaOk

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) GetExactMatchCriteriaOk() (*string, bool)`

GetExactMatchCriteriaOk returns a tuple with the ExactMatchCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactMatchCriteria

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) SetExactMatchCriteria(v string)`

SetExactMatchCriteria sets ExactMatchCriteria field to given value.

### HasExactMatchCriteria

`func (o *CapabilitiesImportRulesUserCreateAndMatchObject) HasExactMatchCriteria() bool`

HasExactMatchCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


