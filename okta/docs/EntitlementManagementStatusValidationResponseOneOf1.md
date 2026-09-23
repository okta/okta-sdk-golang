# EntitlementManagementStatusValidationResponseOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** |  | 
**ReasonCode** | **string** | The machine-readable reason code for blocking the transition to the Entitlement Management state | 
**Message** | **string** | Human-readable description of the validation failure. | 

## Methods

### NewEntitlementManagementStatusValidationResponseOneOf1

`func NewEntitlementManagementStatusValidationResponseOneOf1(valid bool, reasonCode string, message string, ) *EntitlementManagementStatusValidationResponseOneOf1`

NewEntitlementManagementStatusValidationResponseOneOf1 instantiates a new EntitlementManagementStatusValidationResponseOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementManagementStatusValidationResponseOneOf1WithDefaults

`func NewEntitlementManagementStatusValidationResponseOneOf1WithDefaults() *EntitlementManagementStatusValidationResponseOneOf1`

NewEntitlementManagementStatusValidationResponseOneOf1WithDefaults instantiates a new EntitlementManagementStatusValidationResponseOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *EntitlementManagementStatusValidationResponseOneOf1) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *EntitlementManagementStatusValidationResponseOneOf1) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *EntitlementManagementStatusValidationResponseOneOf1) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetReasonCode

`func (o *EntitlementManagementStatusValidationResponseOneOf1) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *EntitlementManagementStatusValidationResponseOneOf1) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *EntitlementManagementStatusValidationResponseOneOf1) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetMessage

`func (o *EntitlementManagementStatusValidationResponseOneOf1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EntitlementManagementStatusValidationResponseOneOf1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EntitlementManagementStatusValidationResponseOneOf1) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


