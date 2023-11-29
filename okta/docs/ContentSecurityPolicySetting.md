# ContentSecurityPolicySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |  | [optional] 
**ReportUri** | Pointer to **string** |  | [optional] 
**SrcList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewContentSecurityPolicySetting

`func NewContentSecurityPolicySetting() *ContentSecurityPolicySetting`

NewContentSecurityPolicySetting instantiates a new ContentSecurityPolicySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSecurityPolicySettingWithDefaults

`func NewContentSecurityPolicySettingWithDefaults() *ContentSecurityPolicySetting`

NewContentSecurityPolicySettingWithDefaults instantiates a new ContentSecurityPolicySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ContentSecurityPolicySetting) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ContentSecurityPolicySetting) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ContentSecurityPolicySetting) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ContentSecurityPolicySetting) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetReportUri

`func (o *ContentSecurityPolicySetting) GetReportUri() string`

GetReportUri returns the ReportUri field if non-nil, zero value otherwise.

### GetReportUriOk

`func (o *ContentSecurityPolicySetting) GetReportUriOk() (*string, bool)`

GetReportUriOk returns a tuple with the ReportUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUri

`func (o *ContentSecurityPolicySetting) SetReportUri(v string)`

SetReportUri sets ReportUri field to given value.

### HasReportUri

`func (o *ContentSecurityPolicySetting) HasReportUri() bool`

HasReportUri returns a boolean if a field has been set.

### GetSrcList

`func (o *ContentSecurityPolicySetting) GetSrcList() []string`

GetSrcList returns the SrcList field if non-nil, zero value otherwise.

### GetSrcListOk

`func (o *ContentSecurityPolicySetting) GetSrcListOk() (*[]string, bool)`

GetSrcListOk returns a tuple with the SrcList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcList

`func (o *ContentSecurityPolicySetting) SetSrcList(v []string)`

SetSrcList sets SrcList field to given value.

### HasSrcList

`func (o *ContentSecurityPolicySetting) HasSrcList() bool`

HasSrcList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


