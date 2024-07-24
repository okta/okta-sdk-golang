# BehaviorRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Link** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewBehaviorRule

`func NewBehaviorRule(name string, type_ string, ) *BehaviorRule`

NewBehaviorRule instantiates a new BehaviorRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorRuleWithDefaults

`func NewBehaviorRuleWithDefaults() *BehaviorRule`

NewBehaviorRuleWithDefaults instantiates a new BehaviorRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *BehaviorRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BehaviorRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BehaviorRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BehaviorRule) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *BehaviorRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BehaviorRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BehaviorRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BehaviorRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *BehaviorRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *BehaviorRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *BehaviorRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *BehaviorRule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *BehaviorRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BehaviorRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BehaviorRule) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *BehaviorRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BehaviorRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BehaviorRule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BehaviorRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *BehaviorRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorRule) SetType(v string)`

SetType sets Type field to given value.


### GetLink

`func (o *BehaviorRule) GetLink() LinksSelf`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *BehaviorRule) GetLinkOk() (*LinksSelf, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *BehaviorRule) SetLink(v LinksSelf)`

SetLink sets Link field to given value.

### HasLink

`func (o *BehaviorRule) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


