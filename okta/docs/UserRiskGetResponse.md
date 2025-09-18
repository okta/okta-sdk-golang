# UserRiskGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskLevel** | Pointer to **string** | The risk level associated with the user | [optional] 
**Links** | Pointer to [**UserRiskGetResponseLinks**](UserRiskGetResponseLinks.md) |  | [optional] 

## Methods

### NewUserRiskGetResponse

`func NewUserRiskGetResponse() *UserRiskGetResponse`

NewUserRiskGetResponse instantiates a new UserRiskGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRiskGetResponseWithDefaults

`func NewUserRiskGetResponseWithDefaults() *UserRiskGetResponse`

NewUserRiskGetResponseWithDefaults instantiates a new UserRiskGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskLevel

`func (o *UserRiskGetResponse) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *UserRiskGetResponse) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *UserRiskGetResponse) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *UserRiskGetResponse) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetLinks

`func (o *UserRiskGetResponse) GetLinks() UserRiskGetResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserRiskGetResponse) GetLinksOk() (*UserRiskGetResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserRiskGetResponse) SetLinks(v UserRiskGetResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserRiskGetResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


