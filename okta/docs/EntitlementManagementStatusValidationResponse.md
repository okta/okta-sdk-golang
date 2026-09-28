# EntitlementManagementStatusValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** |  | 
**ReasonCode** | **string** | The machine-readable reason code for blocking the transition to the Entitlement Management state | 
**Message** | **string** | Human-readable description of the validation failure. | 

## Methods

### NewEntitlementManagementStatusValidationResponse

`func NewEntitlementManagementStatusValidationResponse(valid bool, reasonCode string, message string, ) *EntitlementManagementStatusValidationResponse`

NewEntitlementManagementStatusValidationResponse instantiates a new EntitlementManagementStatusValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementManagementStatusValidationResponseWithDefaults

`func NewEntitlementManagementStatusValidationResponseWithDefaults() *EntitlementManagementStatusValidationResponse`

NewEntitlementManagementStatusValidationResponseWithDefaults instantiates a new EntitlementManagementStatusValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *EntitlementManagementStatusValidationResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *EntitlementManagementStatusValidationResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *EntitlementManagementStatusValidationResponse) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetReasonCode

`func (o *EntitlementManagementStatusValidationResponse) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *EntitlementManagementStatusValidationResponse) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *EntitlementManagementStatusValidationResponse) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetMessage

`func (o *EntitlementManagementStatusValidationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EntitlementManagementStatusValidationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EntitlementManagementStatusValidationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


