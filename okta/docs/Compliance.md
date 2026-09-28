# Compliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fips** | Pointer to **string** | The FIPS compliance mode.   &gt; &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; **Note:** When the Flexible Okta Verify authenticator configuration feature is enabled, this is a common setting shared across all per-method authenticators (&#x60;okta_verify_totp&#x60;, &#x60;okta_verify_push&#x60;, &#x60;okta_verify_fastpass&#x60;). Updating this value on any one authenticator applies the change to all three. | [optional] 

## Methods

### NewCompliance

`func NewCompliance() *Compliance`

NewCompliance instantiates a new Compliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceWithDefaults

`func NewComplianceWithDefaults() *Compliance`

NewComplianceWithDefaults instantiates a new Compliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFips

`func (o *Compliance) GetFips() string`

GetFips returns the Fips field if non-nil, zero value otherwise.

### GetFipsOk

`func (o *Compliance) GetFipsOk() (*string, bool)`

GetFipsOk returns a tuple with the Fips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFips

`func (o *Compliance) SetFips(v string)`

SetFips sets Fips field to given value.

### HasFips

`func (o *Compliance) HasFips() bool`

HasFips returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


