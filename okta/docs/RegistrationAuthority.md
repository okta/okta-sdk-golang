# RegistrationAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aki** | Pointer to **string** | The authority key identifier of the RA certificate — the &#x60;ski&#x60; of the CA certificate that signed it. After a rollover or a parallel renewal, this is the renewed CA&#39;s &#x60;ski&#x60;. | [optional] [readonly] 
**AuthorityInstanceId** | Pointer to **string** | The stable identifier of the certificate authority this configuration belongs to | [optional] [readonly] 
**CaExpirationDate** | Pointer to **time.Time** | Timestamp when the CA certificate bound to this configuration expires. It&#39;s resolved from the configuration&#39;s own &#x60;aki&#x60;, so configurations under one authority can report different expiries after a renewal. | [optional] [readonly] 
**ChallengeType** | Pointer to **string** | How the configuration validates an enrollment challenge. &#x60;STATIC&#x60; uses a fixed shared secret, as Jamf Pro does. &#x60;DYNAMIC&#x60; uses a per-enrollment secret that Okta issues. &#x60;DELEGATED&#x60; hands validation to Microsoft Intune. | [optional] [readonly] 
**ConfigInfo** | Pointer to [**RegistrationAuthorityConfigInfo**](RegistrationAuthorityConfigInfo.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the configuration. It&#39;s the &#x60;raId&#x60; segment of the SCEP enrollment URL, and the value you pass to a renewal. | [optional] [readonly] 
**LastUpdatedDate** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the configuration | [optional] 
**Protocol** | Pointer to **string** | The enrollment protocol the configuration serves | [optional] 
**ScepUrl** | Pointer to **string** | The SCEP enrollment URL to put into your MDM software&#39;s profile. It&#39;s built from &#x60;authorityInstanceId&#x60; rather than from a certificate, so a renewal never changes it. | [optional] [readonly] 
**Ski** | Pointer to **string** | The subject key identifier of the RA certificate, in uppercase hexadecimal | [optional] [readonly] 
**SourceConfigId** | Pointer to **string** | The ID of the configuration this one was cloned from, on a configuration created by a parallel renewal. Absent on every other configuration. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the configuration. Only active configurations are returned. | [optional] [readonly] 

## Methods

### NewRegistrationAuthority

`func NewRegistrationAuthority() *RegistrationAuthority`

NewRegistrationAuthority instantiates a new RegistrationAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationAuthorityWithDefaults

`func NewRegistrationAuthorityWithDefaults() *RegistrationAuthority`

NewRegistrationAuthorityWithDefaults instantiates a new RegistrationAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAki

`func (o *RegistrationAuthority) GetAki() string`

GetAki returns the Aki field if non-nil, zero value otherwise.

### GetAkiOk

`func (o *RegistrationAuthority) GetAkiOk() (*string, bool)`

GetAkiOk returns a tuple with the Aki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAki

`func (o *RegistrationAuthority) SetAki(v string)`

SetAki sets Aki field to given value.

### HasAki

`func (o *RegistrationAuthority) HasAki() bool`

HasAki returns a boolean if a field has been set.

### GetAuthorityInstanceId

`func (o *RegistrationAuthority) GetAuthorityInstanceId() string`

GetAuthorityInstanceId returns the AuthorityInstanceId field if non-nil, zero value otherwise.

### GetAuthorityInstanceIdOk

`func (o *RegistrationAuthority) GetAuthorityInstanceIdOk() (*string, bool)`

GetAuthorityInstanceIdOk returns a tuple with the AuthorityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityInstanceId

`func (o *RegistrationAuthority) SetAuthorityInstanceId(v string)`

SetAuthorityInstanceId sets AuthorityInstanceId field to given value.

### HasAuthorityInstanceId

`func (o *RegistrationAuthority) HasAuthorityInstanceId() bool`

HasAuthorityInstanceId returns a boolean if a field has been set.

### GetCaExpirationDate

`func (o *RegistrationAuthority) GetCaExpirationDate() time.Time`

GetCaExpirationDate returns the CaExpirationDate field if non-nil, zero value otherwise.

### GetCaExpirationDateOk

`func (o *RegistrationAuthority) GetCaExpirationDateOk() (*time.Time, bool)`

GetCaExpirationDateOk returns a tuple with the CaExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaExpirationDate

`func (o *RegistrationAuthority) SetCaExpirationDate(v time.Time)`

SetCaExpirationDate sets CaExpirationDate field to given value.

### HasCaExpirationDate

`func (o *RegistrationAuthority) HasCaExpirationDate() bool`

HasCaExpirationDate returns a boolean if a field has been set.

### GetChallengeType

`func (o *RegistrationAuthority) GetChallengeType() string`

GetChallengeType returns the ChallengeType field if non-nil, zero value otherwise.

### GetChallengeTypeOk

`func (o *RegistrationAuthority) GetChallengeTypeOk() (*string, bool)`

GetChallengeTypeOk returns a tuple with the ChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeType

`func (o *RegistrationAuthority) SetChallengeType(v string)`

SetChallengeType sets ChallengeType field to given value.

### HasChallengeType

`func (o *RegistrationAuthority) HasChallengeType() bool`

HasChallengeType returns a boolean if a field has been set.

### GetConfigInfo

`func (o *RegistrationAuthority) GetConfigInfo() RegistrationAuthorityConfigInfo`

GetConfigInfo returns the ConfigInfo field if non-nil, zero value otherwise.

### GetConfigInfoOk

`func (o *RegistrationAuthority) GetConfigInfoOk() (*RegistrationAuthorityConfigInfo, bool)`

GetConfigInfoOk returns a tuple with the ConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigInfo

`func (o *RegistrationAuthority) SetConfigInfo(v RegistrationAuthorityConfigInfo)`

SetConfigInfo sets ConfigInfo field to given value.

### HasConfigInfo

`func (o *RegistrationAuthority) HasConfigInfo() bool`

HasConfigInfo returns a boolean if a field has been set.

### GetCreatedDate

`func (o *RegistrationAuthority) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *RegistrationAuthority) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *RegistrationAuthority) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *RegistrationAuthority) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetId

`func (o *RegistrationAuthority) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistrationAuthority) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistrationAuthority) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistrationAuthority) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *RegistrationAuthority) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *RegistrationAuthority) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *RegistrationAuthority) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *RegistrationAuthority) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetName

`func (o *RegistrationAuthority) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistrationAuthority) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistrationAuthority) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistrationAuthority) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocol

`func (o *RegistrationAuthority) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RegistrationAuthority) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RegistrationAuthority) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RegistrationAuthority) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetScepUrl

`func (o *RegistrationAuthority) GetScepUrl() string`

GetScepUrl returns the ScepUrl field if non-nil, zero value otherwise.

### GetScepUrlOk

`func (o *RegistrationAuthority) GetScepUrlOk() (*string, bool)`

GetScepUrlOk returns a tuple with the ScepUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScepUrl

`func (o *RegistrationAuthority) SetScepUrl(v string)`

SetScepUrl sets ScepUrl field to given value.

### HasScepUrl

`func (o *RegistrationAuthority) HasScepUrl() bool`

HasScepUrl returns a boolean if a field has been set.

### GetSki

`func (o *RegistrationAuthority) GetSki() string`

GetSki returns the Ski field if non-nil, zero value otherwise.

### GetSkiOk

`func (o *RegistrationAuthority) GetSkiOk() (*string, bool)`

GetSkiOk returns a tuple with the Ski field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSki

`func (o *RegistrationAuthority) SetSki(v string)`

SetSki sets Ski field to given value.

### HasSki

`func (o *RegistrationAuthority) HasSki() bool`

HasSki returns a boolean if a field has been set.

### GetSourceConfigId

`func (o *RegistrationAuthority) GetSourceConfigId() string`

GetSourceConfigId returns the SourceConfigId field if non-nil, zero value otherwise.

### GetSourceConfigIdOk

`func (o *RegistrationAuthority) GetSourceConfigIdOk() (*string, bool)`

GetSourceConfigIdOk returns a tuple with the SourceConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfigId

`func (o *RegistrationAuthority) SetSourceConfigId(v string)`

SetSourceConfigId sets SourceConfigId field to given value.

### HasSourceConfigId

`func (o *RegistrationAuthority) HasSourceConfigId() bool`

HasSourceConfigId returns a boolean if a field has been set.

### GetStatus

`func (o *RegistrationAuthority) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegistrationAuthority) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegistrationAuthority) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegistrationAuthority) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


