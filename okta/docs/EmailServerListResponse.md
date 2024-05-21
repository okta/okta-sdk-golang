# EmailServerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailServers** | Pointer to [**[]EmailServerResponse**](EmailServerResponse.md) |  | [optional] 

## Methods

### NewEmailServerListResponse

`func NewEmailServerListResponse() *EmailServerListResponse`

NewEmailServerListResponse instantiates a new EmailServerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailServerListResponseWithDefaults

`func NewEmailServerListResponseWithDefaults() *EmailServerListResponse`

NewEmailServerListResponseWithDefaults instantiates a new EmailServerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailServers

`func (o *EmailServerListResponse) GetEmailServers() []EmailServerResponse`

GetEmailServers returns the EmailServers field if non-nil, zero value otherwise.

### GetEmailServersOk

`func (o *EmailServerListResponse) GetEmailServersOk() (*[]EmailServerResponse, bool)`

GetEmailServersOk returns a tuple with the EmailServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailServers

`func (o *EmailServerListResponse) SetEmailServers(v []EmailServerResponse)`

SetEmailServers sets EmailServers field to given value.

### HasEmailServers

`func (o *EmailServerListResponse) HasEmailServers() bool`

HasEmailServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


