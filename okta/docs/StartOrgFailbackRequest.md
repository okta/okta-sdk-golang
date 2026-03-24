# StartOrgFailbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to **[]string** | The Okta domain to failback | [optional] 

## Methods

### NewStartOrgFailbackRequest

`func NewStartOrgFailbackRequest() *StartOrgFailbackRequest`

NewStartOrgFailbackRequest instantiates a new StartOrgFailbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartOrgFailbackRequestWithDefaults

`func NewStartOrgFailbackRequestWithDefaults() *StartOrgFailbackRequest`

NewStartOrgFailbackRequestWithDefaults instantiates a new StartOrgFailbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *StartOrgFailbackRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *StartOrgFailbackRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *StartOrgFailbackRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *StartOrgFailbackRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


