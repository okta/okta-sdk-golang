# ApplicationVisibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppLinks** | Pointer to **map[string]bool** |  | [optional] 
**AutoLaunch** | Pointer to **bool** |  | [optional] 
**AutoSubmitToolbar** | Pointer to **bool** |  | [optional] 
**Hide** | Pointer to [**ApplicationVisibilityHide**](ApplicationVisibilityHide.md) |  | [optional] 

## Methods

### NewApplicationVisibility

`func NewApplicationVisibility() *ApplicationVisibility`

NewApplicationVisibility instantiates a new ApplicationVisibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationVisibilityWithDefaults

`func NewApplicationVisibilityWithDefaults() *ApplicationVisibility`

NewApplicationVisibilityWithDefaults instantiates a new ApplicationVisibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppLinks

`func (o *ApplicationVisibility) GetAppLinks() map[string]bool`

GetAppLinks returns the AppLinks field if non-nil, zero value otherwise.

### GetAppLinksOk

`func (o *ApplicationVisibility) GetAppLinksOk() (*map[string]bool, bool)`

GetAppLinksOk returns a tuple with the AppLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLinks

`func (o *ApplicationVisibility) SetAppLinks(v map[string]bool)`

SetAppLinks sets AppLinks field to given value.

### HasAppLinks

`func (o *ApplicationVisibility) HasAppLinks() bool`

HasAppLinks returns a boolean if a field has been set.

### GetAutoLaunch

`func (o *ApplicationVisibility) GetAutoLaunch() bool`

GetAutoLaunch returns the AutoLaunch field if non-nil, zero value otherwise.

### GetAutoLaunchOk

`func (o *ApplicationVisibility) GetAutoLaunchOk() (*bool, bool)`

GetAutoLaunchOk returns a tuple with the AutoLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLaunch

`func (o *ApplicationVisibility) SetAutoLaunch(v bool)`

SetAutoLaunch sets AutoLaunch field to given value.

### HasAutoLaunch

`func (o *ApplicationVisibility) HasAutoLaunch() bool`

HasAutoLaunch returns a boolean if a field has been set.

### GetAutoSubmitToolbar

`func (o *ApplicationVisibility) GetAutoSubmitToolbar() bool`

GetAutoSubmitToolbar returns the AutoSubmitToolbar field if non-nil, zero value otherwise.

### GetAutoSubmitToolbarOk

`func (o *ApplicationVisibility) GetAutoSubmitToolbarOk() (*bool, bool)`

GetAutoSubmitToolbarOk returns a tuple with the AutoSubmitToolbar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSubmitToolbar

`func (o *ApplicationVisibility) SetAutoSubmitToolbar(v bool)`

SetAutoSubmitToolbar sets AutoSubmitToolbar field to given value.

### HasAutoSubmitToolbar

`func (o *ApplicationVisibility) HasAutoSubmitToolbar() bool`

HasAutoSubmitToolbar returns a boolean if a field has been set.

### GetHide

`func (o *ApplicationVisibility) GetHide() ApplicationVisibilityHide`

GetHide returns the Hide field if non-nil, zero value otherwise.

### GetHideOk

`func (o *ApplicationVisibility) GetHideOk() (*ApplicationVisibilityHide, bool)`

GetHideOk returns a tuple with the Hide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHide

`func (o *ApplicationVisibility) SetHide(v ApplicationVisibilityHide)`

SetHide sets Hide field to given value.

### HasHide

`func (o *ApplicationVisibility) HasHide() bool`

HasHide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


