# BreachedCredentialProtectionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionMethod** | **string** | The breach detection method used for breached credential protection | 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewBreachedCredentialProtectionConfiguration

`func NewBreachedCredentialProtectionConfiguration(detectionMethod string, ) *BreachedCredentialProtectionConfiguration`

NewBreachedCredentialProtectionConfiguration instantiates a new BreachedCredentialProtectionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreachedCredentialProtectionConfigurationWithDefaults

`func NewBreachedCredentialProtectionConfigurationWithDefaults() *BreachedCredentialProtectionConfiguration`

NewBreachedCredentialProtectionConfigurationWithDefaults instantiates a new BreachedCredentialProtectionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectionMethod

`func (o *BreachedCredentialProtectionConfiguration) GetDetectionMethod() string`

GetDetectionMethod returns the DetectionMethod field if non-nil, zero value otherwise.

### GetDetectionMethodOk

`func (o *BreachedCredentialProtectionConfiguration) GetDetectionMethodOk() (*string, bool)`

GetDetectionMethodOk returns a tuple with the DetectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMethod

`func (o *BreachedCredentialProtectionConfiguration) SetDetectionMethod(v string)`

SetDetectionMethod sets DetectionMethod field to given value.


### GetLinks

`func (o *BreachedCredentialProtectionConfiguration) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BreachedCredentialProtectionConfiguration) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BreachedCredentialProtectionConfiguration) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BreachedCredentialProtectionConfiguration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


