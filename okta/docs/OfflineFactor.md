# OfflineFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Human-readable name of the offline factor | [optional] 
**EnrollDate** | Pointer to **time.Time** | Timestamp when the offline factor was enrolled | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier of the offline factor | [optional] [readonly] 
**Profile** | Pointer to [**OfflineFactorProfile**](OfflineFactorProfile.md) |  | [optional] 
**Type** | Pointer to **string** | The type of offline factor. &#x60;TOTP&#x60; is the only supported offline factor for macOS enrollments, currently. | [optional] 

## Methods

### NewOfflineFactor

`func NewOfflineFactor() *OfflineFactor`

NewOfflineFactor instantiates a new OfflineFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfflineFactorWithDefaults

`func NewOfflineFactorWithDefaults() *OfflineFactor`

NewOfflineFactorWithDefaults instantiates a new OfflineFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *OfflineFactor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OfflineFactor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OfflineFactor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OfflineFactor) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnrollDate

`func (o *OfflineFactor) GetEnrollDate() time.Time`

GetEnrollDate returns the EnrollDate field if non-nil, zero value otherwise.

### GetEnrollDateOk

`func (o *OfflineFactor) GetEnrollDateOk() (*time.Time, bool)`

GetEnrollDateOk returns a tuple with the EnrollDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollDate

`func (o *OfflineFactor) SetEnrollDate(v time.Time)`

SetEnrollDate sets EnrollDate field to given value.

### HasEnrollDate

`func (o *OfflineFactor) HasEnrollDate() bool`

HasEnrollDate returns a boolean if a field has been set.

### GetId

`func (o *OfflineFactor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OfflineFactor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OfflineFactor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OfflineFactor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfile

`func (o *OfflineFactor) GetProfile() OfflineFactorProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OfflineFactor) GetProfileOk() (*OfflineFactorProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OfflineFactor) SetProfile(v OfflineFactorProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *OfflineFactor) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetType

`func (o *OfflineFactor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OfflineFactor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OfflineFactor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OfflineFactor) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


