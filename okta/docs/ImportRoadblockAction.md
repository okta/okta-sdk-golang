# ImportRoadblockAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to apply:  * &#x60;RESTART&#x60;: Restarts the import that triggered the roadblock, bypassing the roadblock safeguard for that import. * &#x60;CANCEL&#x60;: Cancels the import that triggered the roadblock and resumes the user imports paused by that roadblock. The resumed imports run with the roadblock thresholds still enforced: no safeguard is skipped. | 
**AppInstanceId** | **string** | ID of the app instance whose roadblock to act on | 

## Methods

### NewImportRoadblockAction

`func NewImportRoadblockAction(action string, appInstanceId string, ) *ImportRoadblockAction`

NewImportRoadblockAction instantiates a new ImportRoadblockAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRoadblockActionWithDefaults

`func NewImportRoadblockActionWithDefaults() *ImportRoadblockAction`

NewImportRoadblockActionWithDefaults instantiates a new ImportRoadblockAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ImportRoadblockAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ImportRoadblockAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ImportRoadblockAction) SetAction(v string)`

SetAction sets Action field to given value.


### GetAppInstanceId

`func (o *ImportRoadblockAction) GetAppInstanceId() string`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *ImportRoadblockAction) GetAppInstanceIdOk() (*string, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *ImportRoadblockAction) SetAppInstanceId(v string)`

SetAppInstanceId sets AppInstanceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


