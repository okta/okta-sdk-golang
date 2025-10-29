# UserRiskPutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Describes the risk level for the user | [optional] 
**RiskLevel** | Pointer to **string** | The risk level associated with the user | [optional] 
**Links** | Pointer to [**UserRiskGetResponseLinks**](UserRiskGetResponseLinks.md) |  | [optional] 

## Methods

### NewUserRiskPutResponse

`func NewUserRiskPutResponse() *UserRiskPutResponse`

NewUserRiskPutResponse instantiates a new UserRiskPutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRiskPutResponseWithDefaults

`func NewUserRiskPutResponseWithDefaults() *UserRiskPutResponse`

NewUserRiskPutResponseWithDefaults instantiates a new UserRiskPutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *UserRiskPutResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UserRiskPutResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UserRiskPutResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UserRiskPutResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRiskLevel

`func (o *UserRiskPutResponse) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *UserRiskPutResponse) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *UserRiskPutResponse) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *UserRiskPutResponse) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetLinks

`func (o *UserRiskPutResponse) GetLinks() UserRiskGetResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserRiskPutResponse) GetLinksOk() (*UserRiskGetResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserRiskPutResponse) SetLinks(v UserRiskGetResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserRiskPutResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


