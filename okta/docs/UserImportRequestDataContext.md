# UserImportRequestDataContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conflicts** | Pointer to **[]map[string]interface{}** | An array of user profile attributes that are in conflict | [optional] 
**Application** | Pointer to [**UserImportRequestDataContextApplication**](UserImportRequestDataContextApplication.md) |  | [optional] 
**Job** | Pointer to [**UserImportRequestDataContextJob**](UserImportRequestDataContextJob.md) |  | [optional] 
**Matches** | Pointer to **[]map[string]interface{}** | The list of Okta users currently matched to the app user based on import matching. There can be more than one match. | [optional] 
**Policy** | Pointer to **[]map[string]interface{}** | The list of any policies that apply to the import matching | [optional] 

## Methods

### NewUserImportRequestDataContext

`func NewUserImportRequestDataContext() *UserImportRequestDataContext`

NewUserImportRequestDataContext instantiates a new UserImportRequestDataContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportRequestDataContextWithDefaults

`func NewUserImportRequestDataContextWithDefaults() *UserImportRequestDataContext`

NewUserImportRequestDataContextWithDefaults instantiates a new UserImportRequestDataContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConflicts

`func (o *UserImportRequestDataContext) GetConflicts() []map[string]interface{}`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *UserImportRequestDataContext) GetConflictsOk() (*[]map[string]interface{}, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *UserImportRequestDataContext) SetConflicts(v []map[string]interface{})`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *UserImportRequestDataContext) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetApplication

`func (o *UserImportRequestDataContext) GetApplication() UserImportRequestDataContextApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *UserImportRequestDataContext) GetApplicationOk() (*UserImportRequestDataContextApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *UserImportRequestDataContext) SetApplication(v UserImportRequestDataContextApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *UserImportRequestDataContext) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetJob

`func (o *UserImportRequestDataContext) GetJob() UserImportRequestDataContextJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *UserImportRequestDataContext) GetJobOk() (*UserImportRequestDataContextJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *UserImportRequestDataContext) SetJob(v UserImportRequestDataContextJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *UserImportRequestDataContext) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetMatches

`func (o *UserImportRequestDataContext) GetMatches() []map[string]interface{}`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *UserImportRequestDataContext) GetMatchesOk() (*[]map[string]interface{}, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *UserImportRequestDataContext) SetMatches(v []map[string]interface{})`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *UserImportRequestDataContext) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### GetPolicy

`func (o *UserImportRequestDataContext) GetPolicy() []map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UserImportRequestDataContext) GetPolicyOk() (*[]map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UserImportRequestDataContext) SetPolicy(v []map[string]interface{})`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UserImportRequestDataContext) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


