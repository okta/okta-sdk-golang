# OktaSignOnPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**OktaSignOnPolicyConditions**](OktaSignOnPolicyConditions.md) |  | [optional] 

## Methods

### NewOktaSignOnPolicy

`func NewOktaSignOnPolicy() *OktaSignOnPolicy`

NewOktaSignOnPolicy instantiates a new OktaSignOnPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaSignOnPolicyWithDefaults

`func NewOktaSignOnPolicyWithDefaults() *OktaSignOnPolicy`

NewOktaSignOnPolicyWithDefaults instantiates a new OktaSignOnPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *OktaSignOnPolicy) GetConditions() OktaSignOnPolicyConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *OktaSignOnPolicy) GetConditionsOk() (*OktaSignOnPolicyConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *OktaSignOnPolicy) SetConditions(v OktaSignOnPolicyConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *OktaSignOnPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


