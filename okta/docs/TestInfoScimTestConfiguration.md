# TestInfoScimTestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpecTestResults** | **string** | The Runscope URL to your SCIM server specification test results. See [Test your SCIM API](https://developer.okta.com/docs/guides/build-provisioning-integration/test-scim-api/). | 
**CrudTestResults** | **string** | The Runscope URL to your Okta SCIM CRUD test results. See [Test your Okta SCIM integration](https://developer.okta.com/docs/guides/scim-provisioning-integration-test/main/). | 
**EntitlementsTestResults** | Pointer to **string** | The Runscope URL to your entitlements test results | [optional] 

## Methods

### NewTestInfoScimTestConfiguration

`func NewTestInfoScimTestConfiguration(specTestResults string, crudTestResults string, ) *TestInfoScimTestConfiguration`

NewTestInfoScimTestConfiguration instantiates a new TestInfoScimTestConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestInfoScimTestConfigurationWithDefaults

`func NewTestInfoScimTestConfigurationWithDefaults() *TestInfoScimTestConfiguration`

NewTestInfoScimTestConfigurationWithDefaults instantiates a new TestInfoScimTestConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecTestResults

`func (o *TestInfoScimTestConfiguration) GetSpecTestResults() string`

GetSpecTestResults returns the SpecTestResults field if non-nil, zero value otherwise.

### GetSpecTestResultsOk

`func (o *TestInfoScimTestConfiguration) GetSpecTestResultsOk() (*string, bool)`

GetSpecTestResultsOk returns a tuple with the SpecTestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTestResults

`func (o *TestInfoScimTestConfiguration) SetSpecTestResults(v string)`

SetSpecTestResults sets SpecTestResults field to given value.


### GetCrudTestResults

`func (o *TestInfoScimTestConfiguration) GetCrudTestResults() string`

GetCrudTestResults returns the CrudTestResults field if non-nil, zero value otherwise.

### GetCrudTestResultsOk

`func (o *TestInfoScimTestConfiguration) GetCrudTestResultsOk() (*string, bool)`

GetCrudTestResultsOk returns a tuple with the CrudTestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrudTestResults

`func (o *TestInfoScimTestConfiguration) SetCrudTestResults(v string)`

SetCrudTestResults sets CrudTestResults field to given value.


### GetEntitlementsTestResults

`func (o *TestInfoScimTestConfiguration) GetEntitlementsTestResults() string`

GetEntitlementsTestResults returns the EntitlementsTestResults field if non-nil, zero value otherwise.

### GetEntitlementsTestResultsOk

`func (o *TestInfoScimTestConfiguration) GetEntitlementsTestResultsOk() (*string, bool)`

GetEntitlementsTestResultsOk returns a tuple with the EntitlementsTestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsTestResults

`func (o *TestInfoScimTestConfiguration) SetEntitlementsTestResults(v string)`

SetEntitlementsTestResults sets EntitlementsTestResults field to given value.

### HasEntitlementsTestResults

`func (o *TestInfoScimTestConfiguration) HasEntitlementsTestResults() bool`

HasEntitlementsTestResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


