# SubmissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**[]SubmissionResponseConfigInner**](SubmissionResponseConfigInner.md) | List of org-level variables for the customer per-tenant configuration. For example, a &#x60;subdomain&#x60; variable can be used in the ACS URL: &#x60;https://${org.subdomain}.example.com/saml/login&#x60; | [optional] 
**Description** | Pointer to **string** | A general description of your application and the benefits provided to your customers | [optional] 
**Id** | Pointer to **string** | OIN Integration ID | [optional] [readonly] 
**LastPublished** | Pointer to **string** | Timestamp when the OIN Integration was last published | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp when the OIN Integration instance was last updated | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** | ID of the user who made the last update | [optional] [readonly] 
**Logo** | Pointer to **string** | URL to an uploaded application logo. This logo appears next to your app integration name in the OIN catalog. You must first [Upload an OIN Integration logo](/openapi/okta-management/management/tag/YourOinIntegrations/#tag/YourOinIntegrations/operation/uploadSubmissionLogo) to obtain the logo URL before you can specify this value. | [optional] 
**Name** | Pointer to **string** | The app integration name. This is the main title used for your integration in the OIN catalog. | [optional] 
**Sso** | Pointer to [**Sso**](Sso.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the OIN Integration submission | [optional] [readonly] 

## Methods

### NewSubmissionResponse

`func NewSubmissionResponse() *SubmissionResponse`

NewSubmissionResponse instantiates a new SubmissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionResponseWithDefaults

`func NewSubmissionResponseWithDefaults() *SubmissionResponse`

NewSubmissionResponseWithDefaults instantiates a new SubmissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SubmissionResponse) GetConfig() []SubmissionResponseConfigInner`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SubmissionResponse) GetConfigOk() (*[]SubmissionResponseConfigInner, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SubmissionResponse) SetConfig(v []SubmissionResponseConfigInner)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SubmissionResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *SubmissionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubmissionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubmissionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubmissionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SubmissionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubmissionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastPublished

`func (o *SubmissionResponse) GetLastPublished() string`

GetLastPublished returns the LastPublished field if non-nil, zero value otherwise.

### GetLastPublishedOk

`func (o *SubmissionResponse) GetLastPublishedOk() (*string, bool)`

GetLastPublishedOk returns a tuple with the LastPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPublished

`func (o *SubmissionResponse) SetLastPublished(v string)`

SetLastPublished sets LastPublished field to given value.

### HasLastPublished

`func (o *SubmissionResponse) HasLastPublished() bool`

HasLastPublished returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SubmissionResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SubmissionResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SubmissionResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SubmissionResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *SubmissionResponse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *SubmissionResponse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *SubmissionResponse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *SubmissionResponse) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLogo

`func (o *SubmissionResponse) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *SubmissionResponse) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *SubmissionResponse) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *SubmissionResponse) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *SubmissionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubmissionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubmissionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubmissionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSso

`func (o *SubmissionResponse) GetSso() Sso`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *SubmissionResponse) GetSsoOk() (*Sso, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *SubmissionResponse) SetSso(v Sso)`

SetSso sets Sso field to given value.

### HasSso

`func (o *SubmissionResponse) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetStatus

`func (o *SubmissionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmissionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmissionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmissionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


