# StartOrgFailoverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to **[]string** | The Okta domain to failover | [optional] 

## Methods

### NewStartOrgFailoverRequest

`func NewStartOrgFailoverRequest() *StartOrgFailoverRequest`

NewStartOrgFailoverRequest instantiates a new StartOrgFailoverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartOrgFailoverRequestWithDefaults

`func NewStartOrgFailoverRequestWithDefaults() *StartOrgFailoverRequest`

NewStartOrgFailoverRequestWithDefaults instantiates a new StartOrgFailoverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *StartOrgFailoverRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *StartOrgFailoverRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *StartOrgFailoverRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *StartOrgFailoverRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


