# OrgAerialConsentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The unique ID of the Aerial account | 
**GrantedBy** | Pointer to **string** | Principal ID of the user who granted the permission | [optional] 
**GrantedDate** | Pointer to **string** | Date when grant was created | [optional] 
**Links** | Pointer to [**LinksAerialConsentGranted**](LinksAerialConsentGranted.md) |  | [optional] 

## Methods

### NewOrgAerialConsentDetails

`func NewOrgAerialConsentDetails(accountId string, ) *OrgAerialConsentDetails`

NewOrgAerialConsentDetails instantiates a new OrgAerialConsentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgAerialConsentDetailsWithDefaults

`func NewOrgAerialConsentDetailsWithDefaults() *OrgAerialConsentDetails`

NewOrgAerialConsentDetailsWithDefaults instantiates a new OrgAerialConsentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *OrgAerialConsentDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OrgAerialConsentDetails) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OrgAerialConsentDetails) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetGrantedBy

`func (o *OrgAerialConsentDetails) GetGrantedBy() string`

GetGrantedBy returns the GrantedBy field if non-nil, zero value otherwise.

### GetGrantedByOk

`func (o *OrgAerialConsentDetails) GetGrantedByOk() (*string, bool)`

GetGrantedByOk returns a tuple with the GrantedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedBy

`func (o *OrgAerialConsentDetails) SetGrantedBy(v string)`

SetGrantedBy sets GrantedBy field to given value.

### HasGrantedBy

`func (o *OrgAerialConsentDetails) HasGrantedBy() bool`

HasGrantedBy returns a boolean if a field has been set.

### GetGrantedDate

`func (o *OrgAerialConsentDetails) GetGrantedDate() string`

GetGrantedDate returns the GrantedDate field if non-nil, zero value otherwise.

### GetGrantedDateOk

`func (o *OrgAerialConsentDetails) GetGrantedDateOk() (*string, bool)`

GetGrantedDateOk returns a tuple with the GrantedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedDate

`func (o *OrgAerialConsentDetails) SetGrantedDate(v string)`

SetGrantedDate sets GrantedDate field to given value.

### HasGrantedDate

`func (o *OrgAerialConsentDetails) HasGrantedDate() bool`

HasGrantedDate returns a boolean if a field has been set.

### GetLinks

`func (o *OrgAerialConsentDetails) GetLinks() LinksAerialConsentGranted`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrgAerialConsentDetails) GetLinksOk() (*LinksAerialConsentGranted, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrgAerialConsentDetails) SetLinks(v LinksAerialConsentGranted)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrgAerialConsentDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


