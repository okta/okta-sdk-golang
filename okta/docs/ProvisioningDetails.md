# ProvisioningDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | **[]string** | List of provisioning features supported in this integration | 
**Scim** | [**Scim**](Scim.md) |  | 

## Methods

### NewProvisioningDetails

`func NewProvisioningDetails(features []string, scim Scim, ) *ProvisioningDetails`

NewProvisioningDetails instantiates a new ProvisioningDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningDetailsWithDefaults

`func NewProvisioningDetailsWithDefaults() *ProvisioningDetails`

NewProvisioningDetailsWithDefaults instantiates a new ProvisioningDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *ProvisioningDetails) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ProvisioningDetails) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ProvisioningDetails) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetScim

`func (o *ProvisioningDetails) GetScim() Scim`

GetScim returns the Scim field if non-nil, zero value otherwise.

### GetScimOk

`func (o *ProvisioningDetails) GetScimOk() (*Scim, bool)`

GetScimOk returns a tuple with the Scim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScim

`func (o *ProvisioningDetails) SetScim(v Scim)`

SetScim sets Scim field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


