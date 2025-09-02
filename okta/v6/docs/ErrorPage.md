# ErrorPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageContent** | Pointer to **string** | The HTML for the page | [optional] 
**ContentSecurityPolicySetting** | Pointer to [**ContentSecurityPolicySetting**](ContentSecurityPolicySetting.md) |  | [optional] 

## Methods

### NewErrorPage

`func NewErrorPage() *ErrorPage`

NewErrorPage instantiates a new ErrorPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorPageWithDefaults

`func NewErrorPageWithDefaults() *ErrorPage`

NewErrorPageWithDefaults instantiates a new ErrorPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageContent

`func (o *ErrorPage) GetPageContent() string`

GetPageContent returns the PageContent field if non-nil, zero value otherwise.

### GetPageContentOk

`func (o *ErrorPage) GetPageContentOk() (*string, bool)`

GetPageContentOk returns a tuple with the PageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageContent

`func (o *ErrorPage) SetPageContent(v string)`

SetPageContent sets PageContent field to given value.

### HasPageContent

`func (o *ErrorPage) HasPageContent() bool`

HasPageContent returns a boolean if a field has been set.

### GetContentSecurityPolicySetting

`func (o *ErrorPage) GetContentSecurityPolicySetting() ContentSecurityPolicySetting`

GetContentSecurityPolicySetting returns the ContentSecurityPolicySetting field if non-nil, zero value otherwise.

### GetContentSecurityPolicySettingOk

`func (o *ErrorPage) GetContentSecurityPolicySettingOk() (*ContentSecurityPolicySetting, bool)`

GetContentSecurityPolicySettingOk returns a tuple with the ContentSecurityPolicySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSecurityPolicySetting

`func (o *ErrorPage) SetContentSecurityPolicySetting(v ContentSecurityPolicySetting)`

SetContentSecurityPolicySetting sets ContentSecurityPolicySetting field to given value.

### HasContentSecurityPolicySetting

`func (o *ErrorPage) HasContentSecurityPolicySetting() bool`

HasContentSecurityPolicySetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


