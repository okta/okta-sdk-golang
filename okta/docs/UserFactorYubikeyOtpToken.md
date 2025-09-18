# UserFactorYubikeyOtpToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the token was created | [optional] [readonly] 
**Id** | Pointer to **string** | ID of the token | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the token was last updated | [optional] [readonly] 
**LastVerified** | Pointer to **time.Time** | Timestamp when the token was last verified | [optional] [readonly] 
**Profile** | Pointer to **map[string]map[string]interface{}** | Specified profile information for token | [optional] 
**Status** | Pointer to **string** | Token status | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Links** | Pointer to [**UserFactorLinks**](UserFactorLinks.md) |  | [optional] 

## Methods

### NewUserFactorYubikeyOtpToken

`func NewUserFactorYubikeyOtpToken() *UserFactorYubikeyOtpToken`

NewUserFactorYubikeyOtpToken instantiates a new UserFactorYubikeyOtpToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorYubikeyOtpTokenWithDefaults

`func NewUserFactorYubikeyOtpTokenWithDefaults() *UserFactorYubikeyOtpToken`

NewUserFactorYubikeyOtpTokenWithDefaults instantiates a new UserFactorYubikeyOtpToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UserFactorYubikeyOtpToken) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserFactorYubikeyOtpToken) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserFactorYubikeyOtpToken) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserFactorYubikeyOtpToken) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *UserFactorYubikeyOtpToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserFactorYubikeyOtpToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserFactorYubikeyOtpToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserFactorYubikeyOtpToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UserFactorYubikeyOtpToken) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UserFactorYubikeyOtpToken) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UserFactorYubikeyOtpToken) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UserFactorYubikeyOtpToken) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastVerified

`func (o *UserFactorYubikeyOtpToken) GetLastVerified() time.Time`

GetLastVerified returns the LastVerified field if non-nil, zero value otherwise.

### GetLastVerifiedOk

`func (o *UserFactorYubikeyOtpToken) GetLastVerifiedOk() (*time.Time, bool)`

GetLastVerifiedOk returns a tuple with the LastVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerified

`func (o *UserFactorYubikeyOtpToken) SetLastVerified(v time.Time)`

SetLastVerified sets LastVerified field to given value.

### HasLastVerified

`func (o *UserFactorYubikeyOtpToken) HasLastVerified() bool`

HasLastVerified returns a boolean if a field has been set.

### GetProfile

`func (o *UserFactorYubikeyOtpToken) GetProfile() map[string]map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorYubikeyOtpToken) GetProfileOk() (*map[string]map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorYubikeyOtpToken) SetProfile(v map[string]map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorYubikeyOtpToken) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStatus

`func (o *UserFactorYubikeyOtpToken) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserFactorYubikeyOtpToken) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserFactorYubikeyOtpToken) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserFactorYubikeyOtpToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEmbedded

`func (o *UserFactorYubikeyOtpToken) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *UserFactorYubikeyOtpToken) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *UserFactorYubikeyOtpToken) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *UserFactorYubikeyOtpToken) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *UserFactorYubikeyOtpToken) GetLinks() UserFactorLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserFactorYubikeyOtpToken) GetLinksOk() (*UserFactorLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserFactorYubikeyOtpToken) SetLinks(v UserFactorLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserFactorYubikeyOtpToken) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


