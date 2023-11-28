# LogEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to [**LogActor**](LogActor.md) |  | [optional] 
**AuthenticationContext** | Pointer to [**LogAuthenticationContext**](LogAuthenticationContext.md) |  | [optional] 
**Client** | Pointer to [**LogClient**](LogClient.md) |  | [optional] 
**DebugContext** | Pointer to [**LogDebugContext**](LogDebugContext.md) |  | [optional] 
**DisplayMessage** | Pointer to **string** |  | [optional] [readonly] 
**EventType** | Pointer to **string** |  | [optional] [readonly] 
**LegacyEventType** | Pointer to **string** |  | [optional] [readonly] 
**Outcome** | Pointer to [**LogOutcome**](LogOutcome.md) |  | [optional] 
**Published** | Pointer to **time.Time** |  | [optional] [readonly] 
**Request** | Pointer to [**LogRequest**](LogRequest.md) |  | [optional] 
**SecurityContext** | Pointer to [**LogSecurityContext**](LogSecurityContext.md) |  | [optional] 
**Severity** | Pointer to [**LogSeverity**](LogSeverity.md) |  | [optional] 
**Target** | Pointer to [**[]LogTarget**](LogTarget.md) |  | [optional] [readonly] 
**Transaction** | Pointer to [**LogTransaction**](LogTransaction.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLogEvent

`func NewLogEvent() *LogEvent`

NewLogEvent instantiates a new LogEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogEventWithDefaults

`func NewLogEventWithDefaults() *LogEvent`

NewLogEventWithDefaults instantiates a new LogEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *LogEvent) GetActor() LogActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *LogEvent) GetActorOk() (*LogActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *LogEvent) SetActor(v LogActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *LogEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAuthenticationContext

`func (o *LogEvent) GetAuthenticationContext() LogAuthenticationContext`

GetAuthenticationContext returns the AuthenticationContext field if non-nil, zero value otherwise.

### GetAuthenticationContextOk

`func (o *LogEvent) GetAuthenticationContextOk() (*LogAuthenticationContext, bool)`

GetAuthenticationContextOk returns a tuple with the AuthenticationContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationContext

`func (o *LogEvent) SetAuthenticationContext(v LogAuthenticationContext)`

SetAuthenticationContext sets AuthenticationContext field to given value.

### HasAuthenticationContext

`func (o *LogEvent) HasAuthenticationContext() bool`

HasAuthenticationContext returns a boolean if a field has been set.

### GetClient

`func (o *LogEvent) GetClient() LogClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *LogEvent) GetClientOk() (*LogClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *LogEvent) SetClient(v LogClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *LogEvent) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetDebugContext

`func (o *LogEvent) GetDebugContext() LogDebugContext`

GetDebugContext returns the DebugContext field if non-nil, zero value otherwise.

### GetDebugContextOk

`func (o *LogEvent) GetDebugContextOk() (*LogDebugContext, bool)`

GetDebugContextOk returns a tuple with the DebugContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugContext

`func (o *LogEvent) SetDebugContext(v LogDebugContext)`

SetDebugContext sets DebugContext field to given value.

### HasDebugContext

`func (o *LogEvent) HasDebugContext() bool`

HasDebugContext returns a boolean if a field has been set.

### GetDisplayMessage

`func (o *LogEvent) GetDisplayMessage() string`

GetDisplayMessage returns the DisplayMessage field if non-nil, zero value otherwise.

### GetDisplayMessageOk

`func (o *LogEvent) GetDisplayMessageOk() (*string, bool)`

GetDisplayMessageOk returns a tuple with the DisplayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMessage

`func (o *LogEvent) SetDisplayMessage(v string)`

SetDisplayMessage sets DisplayMessage field to given value.

### HasDisplayMessage

`func (o *LogEvent) HasDisplayMessage() bool`

HasDisplayMessage returns a boolean if a field has been set.

### GetEventType

`func (o *LogEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *LogEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *LogEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *LogEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetLegacyEventType

`func (o *LogEvent) GetLegacyEventType() string`

GetLegacyEventType returns the LegacyEventType field if non-nil, zero value otherwise.

### GetLegacyEventTypeOk

`func (o *LogEvent) GetLegacyEventTypeOk() (*string, bool)`

GetLegacyEventTypeOk returns a tuple with the LegacyEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyEventType

`func (o *LogEvent) SetLegacyEventType(v string)`

SetLegacyEventType sets LegacyEventType field to given value.

### HasLegacyEventType

`func (o *LogEvent) HasLegacyEventType() bool`

HasLegacyEventType returns a boolean if a field has been set.

### GetOutcome

`func (o *LogEvent) GetOutcome() LogOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *LogEvent) GetOutcomeOk() (*LogOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *LogEvent) SetOutcome(v LogOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *LogEvent) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetPublished

`func (o *LogEvent) GetPublished() time.Time`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *LogEvent) GetPublishedOk() (*time.Time, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *LogEvent) SetPublished(v time.Time)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *LogEvent) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRequest

`func (o *LogEvent) GetRequest() LogRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *LogEvent) GetRequestOk() (*LogRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *LogEvent) SetRequest(v LogRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *LogEvent) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSecurityContext

`func (o *LogEvent) GetSecurityContext() LogSecurityContext`

GetSecurityContext returns the SecurityContext field if non-nil, zero value otherwise.

### GetSecurityContextOk

`func (o *LogEvent) GetSecurityContextOk() (*LogSecurityContext, bool)`

GetSecurityContextOk returns a tuple with the SecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContext

`func (o *LogEvent) SetSecurityContext(v LogSecurityContext)`

SetSecurityContext sets SecurityContext field to given value.

### HasSecurityContext

`func (o *LogEvent) HasSecurityContext() bool`

HasSecurityContext returns a boolean if a field has been set.

### GetSeverity

`func (o *LogEvent) GetSeverity() LogSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *LogEvent) GetSeverityOk() (*LogSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *LogEvent) SetSeverity(v LogSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *LogEvent) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTarget

`func (o *LogEvent) GetTarget() []LogTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogEvent) GetTargetOk() (*[]LogTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LogEvent) SetTarget(v []LogTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *LogEvent) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTransaction

`func (o *LogEvent) GetTransaction() LogTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *LogEvent) GetTransactionOk() (*LogTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *LogEvent) SetTransaction(v LogTransaction)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *LogEvent) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetUuid

`func (o *LogEvent) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LogEvent) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LogEvent) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LogEvent) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *LogEvent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LogEvent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LogEvent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LogEvent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


