# ApiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** |  | [optional] [readonly] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**TokenWindow** | Pointer to **string** | A time duration specified as an [ISO-8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Link** | Pointer to [**ApiTokenLink**](ApiTokenLink.md) |  | [optional] 

## Methods

### NewApiToken

`func NewApiToken(name string, ) *ApiToken`

NewApiToken instantiates a new ApiToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenWithDefaults

`func NewApiTokenWithDefaults() *ApiToken`

NewApiTokenWithDefaults instantiates a new ApiToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *ApiToken) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ApiToken) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ApiToken) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ApiToken) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetCreated

`func (o *ApiToken) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiToken) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiToken) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiToken) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *ApiToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ApiToken) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ApiToken) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ApiToken) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ApiToken) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ApiToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiToken) SetName(v string)`

SetName sets Name field to given value.


### GetTokenWindow

`func (o *ApiToken) GetTokenWindow() string`

GetTokenWindow returns the TokenWindow field if non-nil, zero value otherwise.

### GetTokenWindowOk

`func (o *ApiToken) GetTokenWindowOk() (*string, bool)`

GetTokenWindowOk returns a tuple with the TokenWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenWindow

`func (o *ApiToken) SetTokenWindow(v string)`

SetTokenWindow sets TokenWindow field to given value.

### HasTokenWindow

`func (o *ApiToken) HasTokenWindow() bool`

HasTokenWindow returns a boolean if a field has been set.

### GetUserId

`func (o *ApiToken) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiToken) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiToken) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiToken) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetLink

`func (o *ApiToken) GetLink() ApiTokenLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ApiToken) GetLinkOk() (*ApiTokenLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ApiToken) SetLink(v ApiTokenLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *ApiToken) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


