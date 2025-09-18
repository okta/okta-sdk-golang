# OrgOktaSupportSettingsObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseNumber** | Pointer to **NullableString** | Support case number for the Okta Support access grant | [optional] [readonly] 
**Expiration** | Pointer to **NullableTime** | Expiration of Okta Support | [optional] [readonly] 
**Support** | Pointer to **string** | Status of Okta Support Settings | [optional] 
**Links** | Pointer to [**OrgOktaSupportSettingsObjLinks**](OrgOktaSupportSettingsObjLinks.md) |  | [optional] 

## Methods

### NewOrgOktaSupportSettingsObj

`func NewOrgOktaSupportSettingsObj() *OrgOktaSupportSettingsObj`

NewOrgOktaSupportSettingsObj instantiates a new OrgOktaSupportSettingsObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgOktaSupportSettingsObjWithDefaults

`func NewOrgOktaSupportSettingsObjWithDefaults() *OrgOktaSupportSettingsObj`

NewOrgOktaSupportSettingsObjWithDefaults instantiates a new OrgOktaSupportSettingsObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseNumber

`func (o *OrgOktaSupportSettingsObj) GetCaseNumber() string`

GetCaseNumber returns the CaseNumber field if non-nil, zero value otherwise.

### GetCaseNumberOk

`func (o *OrgOktaSupportSettingsObj) GetCaseNumberOk() (*string, bool)`

GetCaseNumberOk returns a tuple with the CaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseNumber

`func (o *OrgOktaSupportSettingsObj) SetCaseNumber(v string)`

SetCaseNumber sets CaseNumber field to given value.

### HasCaseNumber

`func (o *OrgOktaSupportSettingsObj) HasCaseNumber() bool`

HasCaseNumber returns a boolean if a field has been set.

### SetCaseNumberNil

`func (o *OrgOktaSupportSettingsObj) SetCaseNumberNil(b bool)`

 SetCaseNumberNil sets the value for CaseNumber to be an explicit nil

### UnsetCaseNumber
`func (o *OrgOktaSupportSettingsObj) UnsetCaseNumber()`

UnsetCaseNumber ensures that no value is present for CaseNumber, not even an explicit nil
### GetExpiration

`func (o *OrgOktaSupportSettingsObj) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *OrgOktaSupportSettingsObj) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *OrgOktaSupportSettingsObj) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *OrgOktaSupportSettingsObj) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *OrgOktaSupportSettingsObj) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *OrgOktaSupportSettingsObj) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetSupport

`func (o *OrgOktaSupportSettingsObj) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *OrgOktaSupportSettingsObj) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *OrgOktaSupportSettingsObj) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *OrgOktaSupportSettingsObj) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetLinks

`func (o *OrgOktaSupportSettingsObj) GetLinks() OrgOktaSupportSettingsObjLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrgOktaSupportSettingsObj) GetLinksOk() (*OrgOktaSupportSettingsObjLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrgOktaSupportSettingsObj) SetLinks(v OrgOktaSupportSettingsObjLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrgOktaSupportSettingsObj) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


