# AppLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAssignmentId** | Pointer to **string** |  | [optional] [readonly] 
**AppInstanceId** | Pointer to **string** |  | [optional] [readonly] 
**AppName** | Pointer to **string** |  | [optional] [readonly] 
**CredentialsSetup** | Pointer to **bool** |  | [optional] [readonly] 
**Hidden** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Label** | Pointer to **string** |  | [optional] [readonly] 
**LinkUrl** | Pointer to **string** |  | [optional] [readonly] 
**LogoUrl** | Pointer to **string** |  | [optional] [readonly] 
**SortOrder** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewAppLink

`func NewAppLink() *AppLink`

NewAppLink instantiates a new AppLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkWithDefaults

`func NewAppLinkWithDefaults() *AppLink`

NewAppLinkWithDefaults instantiates a new AppLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppAssignmentId

`func (o *AppLink) GetAppAssignmentId() string`

GetAppAssignmentId returns the AppAssignmentId field if non-nil, zero value otherwise.

### GetAppAssignmentIdOk

`func (o *AppLink) GetAppAssignmentIdOk() (*string, bool)`

GetAppAssignmentIdOk returns a tuple with the AppAssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAssignmentId

`func (o *AppLink) SetAppAssignmentId(v string)`

SetAppAssignmentId sets AppAssignmentId field to given value.

### HasAppAssignmentId

`func (o *AppLink) HasAppAssignmentId() bool`

HasAppAssignmentId returns a boolean if a field has been set.

### GetAppInstanceId

`func (o *AppLink) GetAppInstanceId() string`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AppLink) GetAppInstanceIdOk() (*string, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AppLink) SetAppInstanceId(v string)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *AppLink) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.

### GetAppName

`func (o *AppLink) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AppLink) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AppLink) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *AppLink) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetCredentialsSetup

`func (o *AppLink) GetCredentialsSetup() bool`

GetCredentialsSetup returns the CredentialsSetup field if non-nil, zero value otherwise.

### GetCredentialsSetupOk

`func (o *AppLink) GetCredentialsSetupOk() (*bool, bool)`

GetCredentialsSetupOk returns a tuple with the CredentialsSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsSetup

`func (o *AppLink) SetCredentialsSetup(v bool)`

SetCredentialsSetup sets CredentialsSetup field to given value.

### HasCredentialsSetup

`func (o *AppLink) HasCredentialsSetup() bool`

HasCredentialsSetup returns a boolean if a field has been set.

### GetHidden

`func (o *AppLink) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *AppLink) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *AppLink) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *AppLink) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *AppLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *AppLink) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AppLink) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AppLink) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AppLink) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLinkUrl

`func (o *AppLink) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *AppLink) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *AppLink) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *AppLink) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AppLink) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AppLink) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AppLink) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AppLink) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetSortOrder

`func (o *AppLink) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AppLink) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AppLink) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AppLink) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


