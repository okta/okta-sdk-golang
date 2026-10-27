# SpiffeTrustDomainRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleUrl** | **string** | HTTPS URL of the SPIFFE bundle endpoint that serves the trust domain&#39;s JWK Set. Okta fetches this URL on demand to resolve the signing keys for incoming SPIFFE JWT-SVID client assertions. | 
**Created** | Pointer to **time.Time** | Timestamp when the registration was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier of the SPIFFE trust domain registration | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the registration was last updated | [optional] [readonly] 
**TrustDomain** | **string** | The trust domain of a workload&#39;s SPIFFE ID. For example, in the SPIFFE ID &#x60;spiffe://example.org/payments-service&#x60;, the trust domain is &#x60;example.org&#x60;. The trust domain is immutable after creation; use the DELETE endpoint and create a new registration if the trust domain changes. | 

## Methods

### NewSpiffeTrustDomainRegistration

`func NewSpiffeTrustDomainRegistration(bundleUrl string, trustDomain string, ) *SpiffeTrustDomainRegistration`

NewSpiffeTrustDomainRegistration instantiates a new SpiffeTrustDomainRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpiffeTrustDomainRegistrationWithDefaults

`func NewSpiffeTrustDomainRegistrationWithDefaults() *SpiffeTrustDomainRegistration`

NewSpiffeTrustDomainRegistrationWithDefaults instantiates a new SpiffeTrustDomainRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleUrl

`func (o *SpiffeTrustDomainRegistration) GetBundleUrl() string`

GetBundleUrl returns the BundleUrl field if non-nil, zero value otherwise.

### GetBundleUrlOk

`func (o *SpiffeTrustDomainRegistration) GetBundleUrlOk() (*string, bool)`

GetBundleUrlOk returns a tuple with the BundleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleUrl

`func (o *SpiffeTrustDomainRegistration) SetBundleUrl(v string)`

SetBundleUrl sets BundleUrl field to given value.


### GetCreated

`func (o *SpiffeTrustDomainRegistration) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SpiffeTrustDomainRegistration) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SpiffeTrustDomainRegistration) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SpiffeTrustDomainRegistration) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *SpiffeTrustDomainRegistration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpiffeTrustDomainRegistration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpiffeTrustDomainRegistration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpiffeTrustDomainRegistration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SpiffeTrustDomainRegistration) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SpiffeTrustDomainRegistration) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SpiffeTrustDomainRegistration) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SpiffeTrustDomainRegistration) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetTrustDomain

`func (o *SpiffeTrustDomainRegistration) GetTrustDomain() string`

GetTrustDomain returns the TrustDomain field if non-nil, zero value otherwise.

### GetTrustDomainOk

`func (o *SpiffeTrustDomainRegistration) GetTrustDomainOk() (*string, bool)`

GetTrustDomainOk returns a tuple with the TrustDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustDomain

`func (o *SpiffeTrustDomainRegistration) SetTrustDomain(v string)`

SetTrustDomain sets TrustDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


