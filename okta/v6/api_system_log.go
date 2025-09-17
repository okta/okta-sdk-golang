/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

type SystemLogAPI interface {

	/*
			ListLogEvents List all System Log events

			Lists all System Log events

		See [System Log query](https://developer.okta.com/docs/reference/system-log-query/) for further details and examples, and [System Log filters and search](https://help.okta.com/okta_help.htm?type=oie&id=csh-syslog-filters) for common use cases.

		By default, 100 System Log events are returned. If there are more events, see the [header link](https://developer.okta.com/docs/api/#link-header) for the `next` link,
		or increase the number of returned objects using the `limit` parameter.

		>**Note:** The value of the `clientSecret` property in the System Log is secured by a hashing function, and isn't the value used during authentication.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiListLogEventsRequest
	*/
	ListLogEvents(ctx context.Context) ApiListLogEventsRequest

	// ListLogEventsExecute executes the request
	//  @return []LogEvent
	ListLogEventsExecute(r ApiListLogEventsRequest) ([]LogEvent, *APIResponse, error)
}

// SystemLogAPIService SystemLogAPI service
type SystemLogAPIService service

type ApiListLogEventsRequest struct {
	ctx        context.Context
	ApiService SystemLogAPI
	since      *string
	until      *string
	after      *string
	filter     *string
	q          *string
	limit      *int32
	sortOrder  *string
	retryCount int32
}

// Filters the lower time bound of the log events &#x60;published&#x60; property for bounded queries or persistence time for polling queries
func (r ApiListLogEventsRequest) Since(since string) ApiListLogEventsRequest {
	r.since = &since
	return r
}

// Filters the upper time bound of the log events &#x60;published&#x60; property for bounded queries or persistence time for polling queries.
func (r ApiListLogEventsRequest) Until(until string) ApiListLogEventsRequest {
	r.until = &until
	return r
}

// Retrieves the next page of results. Okta returns a link in the HTTP Header (&#x60;rel&#x3D;next&#x60;) that includes the after query parameter
func (r ApiListLogEventsRequest) After(after string) ApiListLogEventsRequest {
	r.after = &after
	return r
}

// Filter expression that filters the results. All operators except [ ] are supported. See [Filter](https://developer.okta.com/docs/api/#filter) and [Operators](https://developer.okta.com/docs/api/#operators).
func (r ApiListLogEventsRequest) Filter(filter string) ApiListLogEventsRequest {
	r.filter = &filter
	return r
}

// Filters log events results by one or more case insensitive keywords.
func (r ApiListLogEventsRequest) Q(q string) ApiListLogEventsRequest {
	r.q = &q
	return r
}

// Sets the number of results that are returned in the response
func (r ApiListLogEventsRequest) Limit(limit int32) ApiListLogEventsRequest {
	r.limit = &limit
	return r
}

// The order of the returned events that are sorted by the &#x60;published&#x60; property
func (r ApiListLogEventsRequest) SortOrder(sortOrder string) ApiListLogEventsRequest {
	r.sortOrder = &sortOrder
	return r
}

func (r ApiListLogEventsRequest) Execute() ([]LogEvent, *APIResponse, error) {
	return r.ApiService.ListLogEventsExecute(r)
}

/*
ListLogEvents List all System Log events

# Lists all System Log events

See [System Log query](https://developer.okta.com/docs/reference/system-log-query/) for further details and examples, and [System Log filters and search](https://help.okta.com/okta_help.htm?type=oie&id=csh-syslog-filters) for common use cases.

By default, 100 System Log events are returned. If there are more events, see the [header link](https://developer.okta.com/docs/api/#link-header) for the `next` link,
or increase the number of returned objects using the `limit` parameter.

>**Note:** The value of the `clientSecret` property in the System Log is secured by a hashing function, and isn't the value used during authentication.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListLogEventsRequest
*/
func (a *SystemLogAPIService) ListLogEvents(ctx context.Context) ApiListLogEventsRequest {
	return ApiListLogEventsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return []LogEvent
func (a *SystemLogAPIService) ListLogEventsExecute(r ApiListLogEventsRequest) ([]LogEvent, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LogEvent
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemLogAPIService.ListLogEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.since != nil {
		localVarQueryParams.Add("since", parameterToString(*r.since, ""))
	}
	if r.until != nil {
		localVarQueryParams.Add("until", parameterToString(*r.until, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}
