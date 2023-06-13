# OrgSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**EndUserSupportHelpURL** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**Subdomain** | Pointer to **string** |  | [optional] [readonly] 
**SupportPhoneNumber** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOrgSetting

`func NewOrgSetting() *OrgSetting`

NewOrgSetting instantiates a new OrgSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingWithDefaults

`func NewOrgSettingWithDefaults() *OrgSetting`

NewOrgSettingWithDefaults instantiates a new OrgSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *OrgSetting) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *OrgSetting) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *OrgSetting) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *OrgSetting) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *OrgSetting) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *OrgSetting) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *OrgSetting) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *OrgSetting) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *OrgSetting) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrgSetting) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrgSetting) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrgSetting) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrgSetting) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrgSetting) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrgSetting) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrgSetting) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *OrgSetting) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrgSetting) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrgSetting) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrgSetting) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreated

`func (o *OrgSetting) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OrgSetting) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OrgSetting) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OrgSetting) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEndUserSupportHelpURL

`func (o *OrgSetting) GetEndUserSupportHelpURL() string`

GetEndUserSupportHelpURL returns the EndUserSupportHelpURL field if non-nil, zero value otherwise.

### GetEndUserSupportHelpURLOk

`func (o *OrgSetting) GetEndUserSupportHelpURLOk() (*string, bool)`

GetEndUserSupportHelpURLOk returns a tuple with the EndUserSupportHelpURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserSupportHelpURL

`func (o *OrgSetting) SetEndUserSupportHelpURL(v string)`

SetEndUserSupportHelpURL sets EndUserSupportHelpURL field to given value.

### HasEndUserSupportHelpURL

`func (o *OrgSetting) HasEndUserSupportHelpURL() bool`

HasEndUserSupportHelpURL returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrgSetting) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrgSetting) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrgSetting) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrgSetting) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *OrgSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgSetting) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrgSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *OrgSetting) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OrgSetting) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OrgSetting) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OrgSetting) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *OrgSetting) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrgSetting) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrgSetting) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrgSetting) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPostalCode

`func (o *OrgSetting) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrgSetting) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrgSetting) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrgSetting) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *OrgSetting) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrgSetting) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrgSetting) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrgSetting) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *OrgSetting) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrgSetting) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrgSetting) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrgSetting) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubdomain

`func (o *OrgSetting) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *OrgSetting) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *OrgSetting) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *OrgSetting) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetSupportPhoneNumber

`func (o *OrgSetting) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *OrgSetting) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *OrgSetting) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.

### HasSupportPhoneNumber

`func (o *OrgSetting) HasSupportPhoneNumber() bool`

HasSupportPhoneNumber returns a boolean if a field has been set.

### GetWebsite

`func (o *OrgSetting) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *OrgSetting) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *OrgSetting) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *OrgSetting) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetLinks

`func (o *OrgSetting) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrgSetting) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrgSetting) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrgSetting) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


