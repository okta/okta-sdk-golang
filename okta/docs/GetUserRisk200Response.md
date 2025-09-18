# GetUserRisk200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskLevel** | Pointer to **string** | The risk level associated with the user | [optional] 
**Links** | Pointer to [**UserRiskGetResponseLinks**](UserRiskGetResponseLinks.md) |  | [optional] 

## Methods

### NewGetUserRisk200Response

`func NewGetUserRisk200Response() *GetUserRisk200Response`

NewGetUserRisk200Response instantiates a new GetUserRisk200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserRisk200ResponseWithDefaults

`func NewGetUserRisk200ResponseWithDefaults() *GetUserRisk200Response`

NewGetUserRisk200ResponseWithDefaults instantiates a new GetUserRisk200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskLevel

`func (o *GetUserRisk200Response) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *GetUserRisk200Response) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *GetUserRisk200Response) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *GetUserRisk200Response) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetLinks

`func (o *GetUserRisk200Response) GetLinks() UserRiskGetResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUserRisk200Response) GetLinksOk() (*UserRiskGetResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUserRisk200Response) SetLinks(v UserRiskGetResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetUserRisk200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


