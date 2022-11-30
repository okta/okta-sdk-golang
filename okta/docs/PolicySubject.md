# PolicySubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **[]string** |  | [optional] 
**MatchAttribute** | Pointer to **string** |  | [optional] 
**MatchType** | Pointer to [**PolicySubjectMatchType**](PolicySubjectMatchType.md) |  | [optional] 
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

### GetFormat

`func (o *PolicySubject) GetFormat() []string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PolicySubject) GetFormatOk() (*[]string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PolicySubject) SetFormat(v []string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PolicySubject) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

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

`func (o *PolicySubject) GetMatchType() PolicySubjectMatchType`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *PolicySubject) GetMatchTypeOk() (*PolicySubjectMatchType, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *PolicySubject) SetMatchType(v PolicySubjectMatchType)`

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


