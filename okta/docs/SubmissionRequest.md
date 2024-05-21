# SubmissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**[]SubmissionResponseConfigInner**](SubmissionResponseConfigInner.md) | List of org-level variables for the customer per-tenant configuration. For example, a &#x60;subdomain&#x60; variable can be used in the ACS URL: &#x60;https://${org.subdomain}.example.com/saml/login&#x60; | [optional] 
**Description** | **string** | A general description of your application and the benefits provided to your customers | 
**Id** | Pointer to **string** | OIN Integration ID | [optional] [readonly] 
**LastPublished** | Pointer to **string** | Timestamp when the OIN Integration was last published | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp when the OIN Integration instance was last updated | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** | ID of the user who made the last update | [optional] [readonly] 
**Logo** | **string** | URL to an uploaded application logo. This logo appears next to your app integration name in the OIN catalog. You must first [Upload an OIN Integration logo](/openapi/okta-management/management/tag/YourOinIntegrations/#tag/YourOinIntegrations/operation/uploadSubmissionLogo) to obtain the logo URL before you can specify this value. | 
**Name** | **string** | The app integration name. This is the main title used for your integration in the OIN catalog. | 
**Sso** | Pointer to [**Sso**](Sso.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the OIN Integration submission | [optional] [readonly] 

## Methods

### NewSubmissionRequest

`func NewSubmissionRequest(description string, logo string, name string, ) *SubmissionRequest`

NewSubmissionRequest instantiates a new SubmissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionRequestWithDefaults

`func NewSubmissionRequestWithDefaults() *SubmissionRequest`

NewSubmissionRequestWithDefaults instantiates a new SubmissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SubmissionRequest) GetConfig() []SubmissionResponseConfigInner`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SubmissionRequest) GetConfigOk() (*[]SubmissionResponseConfigInner, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SubmissionRequest) SetConfig(v []SubmissionResponseConfigInner)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SubmissionRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *SubmissionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubmissionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubmissionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *SubmissionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubmissionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastPublished

`func (o *SubmissionRequest) GetLastPublished() string`

GetLastPublished returns the LastPublished field if non-nil, zero value otherwise.

### GetLastPublishedOk

`func (o *SubmissionRequest) GetLastPublishedOk() (*string, bool)`

GetLastPublishedOk returns a tuple with the LastPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPublished

`func (o *SubmissionRequest) SetLastPublished(v string)`

SetLastPublished sets LastPublished field to given value.

### HasLastPublished

`func (o *SubmissionRequest) HasLastPublished() bool`

HasLastPublished returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SubmissionRequest) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SubmissionRequest) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SubmissionRequest) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SubmissionRequest) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *SubmissionRequest) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *SubmissionRequest) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *SubmissionRequest) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *SubmissionRequest) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLogo

`func (o *SubmissionRequest) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *SubmissionRequest) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *SubmissionRequest) SetLogo(v string)`

SetLogo sets Logo field to given value.


### GetName

`func (o *SubmissionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubmissionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubmissionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSso

`func (o *SubmissionRequest) GetSso() Sso`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *SubmissionRequest) GetSsoOk() (*Sso, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *SubmissionRequest) SetSso(v Sso)`

SetSso sets Sso field to given value.

### HasSso

`func (o *SubmissionRequest) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetStatus

`func (o *SubmissionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmissionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmissionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmissionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


