# ListFactors200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the Factor was enrolled | [optional] [readonly] 
**FactorType** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID of the Factor | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Factor was last updated | [optional] [readonly] 
**Profile** | Pointer to [**UserFactorWebAuthnProfile**](UserFactorWebAuthnProfile.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**VendorName** | Pointer to **string** | Name of the Factor vendor. This is usually the same as the provider except for On-Prem MFA where it depends on administrator settings. | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**LinksSelfAndLifecycle**](LinksSelfAndLifecycle.md) |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Timestamp when the Factor verification attempt expires | [optional] [readonly] 
**FactorResult** | Pointer to **string** | Result of a Factor verification attempt | [optional] [readonly] 
**Verify** | Pointer to [**UserFactorHardwareAllOfVerify**](UserFactorHardwareAllOfVerify.md) |  | [optional] 
**FactorProfileId** | Pointer to **string** | ID of an existing Custom TOTP Factor profile. To create this, see [Custom TOTP Factor](https://help.okta.com/okta_help.htm?id&#x3D;ext-mfa-totp). | [optional] 
**Type** | Pointer to **string** | The type of authenticator method | [optional] 
**Settings** | Pointer to [**AuthenticatorMethodSignedNonceAllOfSettings**](AuthenticatorMethodSignedNonceAllOfSettings.md) |  | [optional] 

## Methods

### NewListFactors200ResponseInner

`func NewListFactors200ResponseInner() *ListFactors200ResponseInner`

NewListFactors200ResponseInner instantiates a new ListFactors200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFactors200ResponseInnerWithDefaults

`func NewListFactors200ResponseInnerWithDefaults() *ListFactors200ResponseInner`

NewListFactors200ResponseInnerWithDefaults instantiates a new ListFactors200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ListFactors200ResponseInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListFactors200ResponseInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListFactors200ResponseInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListFactors200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFactorType

`func (o *ListFactors200ResponseInner) GetFactorType() interface{}`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *ListFactors200ResponseInner) GetFactorTypeOk() (*interface{}, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *ListFactors200ResponseInner) SetFactorType(v interface{})`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *ListFactors200ResponseInner) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *ListFactors200ResponseInner) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *ListFactors200ResponseInner) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetId

`func (o *ListFactors200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListFactors200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListFactors200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListFactors200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListFactors200ResponseInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListFactors200ResponseInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListFactors200ResponseInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListFactors200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *ListFactors200ResponseInner) GetProfile() UserFactorWebAuthnProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ListFactors200ResponseInner) GetProfileOk() (*UserFactorWebAuthnProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ListFactors200ResponseInner) SetProfile(v UserFactorWebAuthnProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ListFactors200ResponseInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *ListFactors200ResponseInner) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ListFactors200ResponseInner) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ListFactors200ResponseInner) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ListFactors200ResponseInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *ListFactors200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListFactors200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListFactors200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListFactors200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVendorName

`func (o *ListFactors200ResponseInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *ListFactors200ResponseInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *ListFactors200ResponseInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *ListFactors200ResponseInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetEmbedded

`func (o *ListFactors200ResponseInner) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListFactors200ResponseInner) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListFactors200ResponseInner) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ListFactors200ResponseInner) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *ListFactors200ResponseInner) GetLinks() LinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListFactors200ResponseInner) GetLinksOk() (*LinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListFactors200ResponseInner) SetLinks(v LinksSelfAndLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListFactors200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ListFactors200ResponseInner) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ListFactors200ResponseInner) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ListFactors200ResponseInner) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ListFactors200ResponseInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFactorResult

`func (o *ListFactors200ResponseInner) GetFactorResult() string`

GetFactorResult returns the FactorResult field if non-nil, zero value otherwise.

### GetFactorResultOk

`func (o *ListFactors200ResponseInner) GetFactorResultOk() (*string, bool)`

GetFactorResultOk returns a tuple with the FactorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorResult

`func (o *ListFactors200ResponseInner) SetFactorResult(v string)`

SetFactorResult sets FactorResult field to given value.

### HasFactorResult

`func (o *ListFactors200ResponseInner) HasFactorResult() bool`

HasFactorResult returns a boolean if a field has been set.

### GetVerify

`func (o *ListFactors200ResponseInner) GetVerify() UserFactorHardwareAllOfVerify`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *ListFactors200ResponseInner) GetVerifyOk() (*UserFactorHardwareAllOfVerify, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *ListFactors200ResponseInner) SetVerify(v UserFactorHardwareAllOfVerify)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *ListFactors200ResponseInner) HasVerify() bool`

HasVerify returns a boolean if a field has been set.

### GetFactorProfileId

`func (o *ListFactors200ResponseInner) GetFactorProfileId() string`

GetFactorProfileId returns the FactorProfileId field if non-nil, zero value otherwise.

### GetFactorProfileIdOk

`func (o *ListFactors200ResponseInner) GetFactorProfileIdOk() (*string, bool)`

GetFactorProfileIdOk returns a tuple with the FactorProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorProfileId

`func (o *ListFactors200ResponseInner) SetFactorProfileId(v string)`

SetFactorProfileId sets FactorProfileId field to given value.

### HasFactorProfileId

`func (o *ListFactors200ResponseInner) HasFactorProfileId() bool`

HasFactorProfileId returns a boolean if a field has been set.

### GetType

`func (o *ListFactors200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListFactors200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListFactors200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListFactors200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSettings

`func (o *ListFactors200ResponseInner) GetSettings() AuthenticatorMethodSignedNonceAllOfSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ListFactors200ResponseInner) GetSettingsOk() (*AuthenticatorMethodSignedNonceAllOfSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ListFactors200ResponseInner) SetSettings(v AuthenticatorMethodSignedNonceAllOfSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ListFactors200ResponseInner) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


