# ProvisioningConnectionResponseProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | **string** | Defines the method of authentication | 
**Signing** | Pointer to [**Org2OrgProvisioningOAuthSigningSettings**](Org2OrgProvisioningOAuthSigningSettings.md) |  | [optional] 

## Methods

### NewProvisioningConnectionResponseProfile

`func NewProvisioningConnectionResponseProfile(authScheme string, ) *ProvisioningConnectionResponseProfile`

NewProvisioningConnectionResponseProfile instantiates a new ProvisioningConnectionResponseProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningConnectionResponseProfileWithDefaults

`func NewProvisioningConnectionResponseProfileWithDefaults() *ProvisioningConnectionResponseProfile`

NewProvisioningConnectionResponseProfileWithDefaults instantiates a new ProvisioningConnectionResponseProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *ProvisioningConnectionResponseProfile) GetAuthScheme() string`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *ProvisioningConnectionResponseProfile) GetAuthSchemeOk() (*string, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *ProvisioningConnectionResponseProfile) SetAuthScheme(v string)`

SetAuthScheme sets AuthScheme field to given value.


### GetSigning

`func (o *ProvisioningConnectionResponseProfile) GetSigning() Org2OrgProvisioningOAuthSigningSettings`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *ProvisioningConnectionResponseProfile) GetSigningOk() (*Org2OrgProvisioningOAuthSigningSettings, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *ProvisioningConnectionResponseProfile) SetSigning(v Org2OrgProvisioningOAuthSigningSettings)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *ProvisioningConnectionResponseProfile) HasSigning() bool`

HasSigning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


