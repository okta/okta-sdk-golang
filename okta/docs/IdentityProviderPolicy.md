# IdentityProviderPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountLink** | Pointer to [**PolicyAccountLink**](PolicyAccountLink.md) |  | [optional] 
**Conditions** | Pointer to [**PolicyRuleConditions**](PolicyRuleConditions.md) |  | [optional] 
**MaxClockSkew** | Pointer to **int32** |  | [optional] 
**Provisioning** | Pointer to [**Provisioning**](Provisioning.md) |  | [optional] 
**Subject** | Pointer to [**PolicySubject**](PolicySubject.md) |  | [optional] 

## Methods

### NewIdentityProviderPolicy

`func NewIdentityProviderPolicy() *IdentityProviderPolicy`

NewIdentityProviderPolicy instantiates a new IdentityProviderPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderPolicyWithDefaults

`func NewIdentityProviderPolicyWithDefaults() *IdentityProviderPolicy`

NewIdentityProviderPolicyWithDefaults instantiates a new IdentityProviderPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountLink

`func (o *IdentityProviderPolicy) GetAccountLink() PolicyAccountLink`

GetAccountLink returns the AccountLink field if non-nil, zero value otherwise.

### GetAccountLinkOk

`func (o *IdentityProviderPolicy) GetAccountLinkOk() (*PolicyAccountLink, bool)`

GetAccountLinkOk returns a tuple with the AccountLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLink

`func (o *IdentityProviderPolicy) SetAccountLink(v PolicyAccountLink)`

SetAccountLink sets AccountLink field to given value.

### HasAccountLink

`func (o *IdentityProviderPolicy) HasAccountLink() bool`

HasAccountLink returns a boolean if a field has been set.

### GetConditions

`func (o *IdentityProviderPolicy) GetConditions() PolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IdentityProviderPolicy) GetConditionsOk() (*PolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IdentityProviderPolicy) SetConditions(v PolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IdentityProviderPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetMaxClockSkew

`func (o *IdentityProviderPolicy) GetMaxClockSkew() int32`

GetMaxClockSkew returns the MaxClockSkew field if non-nil, zero value otherwise.

### GetMaxClockSkewOk

`func (o *IdentityProviderPolicy) GetMaxClockSkewOk() (*int32, bool)`

GetMaxClockSkewOk returns a tuple with the MaxClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClockSkew

`func (o *IdentityProviderPolicy) SetMaxClockSkew(v int32)`

SetMaxClockSkew sets MaxClockSkew field to given value.

### HasMaxClockSkew

`func (o *IdentityProviderPolicy) HasMaxClockSkew() bool`

HasMaxClockSkew returns a boolean if a field has been set.

### GetProvisioning

`func (o *IdentityProviderPolicy) GetProvisioning() Provisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *IdentityProviderPolicy) GetProvisioningOk() (*Provisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *IdentityProviderPolicy) SetProvisioning(v Provisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *IdentityProviderPolicy) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetSubject

`func (o *IdentityProviderPolicy) GetSubject() PolicySubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IdentityProviderPolicy) GetSubjectOk() (*PolicySubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IdentityProviderPolicy) SetSubject(v PolicySubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IdentityProviderPolicy) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


