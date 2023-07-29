# ListBehaviorDetectionRules200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Status** | Pointer to [**LifecycleStatus**](LifecycleStatus.md) |  | [optional] 
**Type** | [**BehaviorRuleType**](BehaviorRuleType.md) |  | 
**Link** | Pointer to [**ApiTokenLink**](ApiTokenLink.md) |  | [optional] 
**Settings** | Pointer to [**BehaviorRuleSettingsVelocity**](BehaviorRuleSettingsVelocity.md) |  | [optional] 

## Methods

### NewListBehaviorDetectionRules200ResponseInner

`func NewListBehaviorDetectionRules200ResponseInner(name string, type_ BehaviorRuleType, ) *ListBehaviorDetectionRules200ResponseInner`

NewListBehaviorDetectionRules200ResponseInner instantiates a new ListBehaviorDetectionRules200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBehaviorDetectionRules200ResponseInnerWithDefaults

`func NewListBehaviorDetectionRules200ResponseInnerWithDefaults() *ListBehaviorDetectionRules200ResponseInner`

NewListBehaviorDetectionRules200ResponseInnerWithDefaults instantiates a new ListBehaviorDetectionRules200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ListBehaviorDetectionRules200ResponseInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListBehaviorDetectionRules200ResponseInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListBehaviorDetectionRules200ResponseInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListBehaviorDetectionRules200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ListBehaviorDetectionRules200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBehaviorDetectionRules200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBehaviorDetectionRules200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListBehaviorDetectionRules200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListBehaviorDetectionRules200ResponseInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListBehaviorDetectionRules200ResponseInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListBehaviorDetectionRules200ResponseInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListBehaviorDetectionRules200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ListBehaviorDetectionRules200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBehaviorDetectionRules200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBehaviorDetectionRules200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ListBehaviorDetectionRules200ResponseInner) GetStatus() LifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBehaviorDetectionRules200ResponseInner) GetStatusOk() (*LifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBehaviorDetectionRules200ResponseInner) SetStatus(v LifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListBehaviorDetectionRules200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ListBehaviorDetectionRules200ResponseInner) GetType() BehaviorRuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListBehaviorDetectionRules200ResponseInner) GetTypeOk() (*BehaviorRuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListBehaviorDetectionRules200ResponseInner) SetType(v BehaviorRuleType)`

SetType sets Type field to given value.


### GetLink

`func (o *ListBehaviorDetectionRules200ResponseInner) GetLink() ApiTokenLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ListBehaviorDetectionRules200ResponseInner) GetLinkOk() (*ApiTokenLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ListBehaviorDetectionRules200ResponseInner) SetLink(v ApiTokenLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *ListBehaviorDetectionRules200ResponseInner) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetSettings

`func (o *ListBehaviorDetectionRules200ResponseInner) GetSettings() BehaviorRuleSettingsVelocity`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ListBehaviorDetectionRules200ResponseInner) GetSettingsOk() (*BehaviorRuleSettingsVelocity, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ListBehaviorDetectionRules200ResponseInner) SetSettings(v BehaviorRuleSettingsVelocity)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ListBehaviorDetectionRules200ResponseInner) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


