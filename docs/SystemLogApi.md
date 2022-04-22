# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogs**](SystemLogApi.md#GetLogs) | **Get** /api/v1/logs | Fetch a list of events from your Okta organization system log.

# **GetLogs**
> []LogEvent GetLogs(ctx, optional)
Fetch a list of events from your Okta organization system log.

The Okta System Log API provides read access to your organization’s system log. This API provides more functionality than the Events API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemLogApiGetLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemLogApiGetLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **optional.Time**|  | 
 **until** | **optional.Time**|  | 
 **filter** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **sortOrder** | **optional.String**|  | [default to ASCENDING]
 **after** | **optional.String**|  | 

### Return type

[**[]LogEvent**](LogEvent.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

