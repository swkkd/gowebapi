// ************************************************************************
// * Copyright 2018 OSIsoft, LLC
// * Licensed under the Apache License, Version 2.0 (the "License");
// * you may not use this file except in compliance with the License.
// * You may obtain a copy of the License at
// *
// *   <http://www.apache.org/licenses/LICENSE-2.0>
// *
// * Unless required by applicable law or agreed to in writing, software
// * distributed under the License is distributed on an "AS IS" BASIS,
// * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// * See the License for the specific language governing permissions and
// * limitations under the License.
// ************************************************************************

package gowebapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/context"
)

// Linger please
var (
	_ context.Context
)

// BatchRequest struct for single request in batch
type BatchRequest struct {
	Resource string `json:"Resource,omitempty"`
	Method   string `json:"Method,omitempty"`
	Content  string `json:"Content,omitempty"`
}

type BatchApiService service

func (a *BatchApiService) ConstructBatchRequest(method string, resource string, content string) BatchRequest {
	return BatchRequest{
		Resource: resource,
		Method:   method,
		Content:  content,
	}
}

/*
	BatchApiService Execute a batch of requests against the service. As shown in the Sample Request, the input is a dictionary with IDs as keys and request objects as values. Each request object specifies the HTTP method and the resource and, optionally, the content and a list of parent IDs. The list of parent IDs specifies which other requests must complete before the given request will be executed. The example first creates an element, then gets the element by the response&#39;s Location header, then creates an attribute for the element. Note that the resource can be an absolute URL or a JsonPath that references the response to the parent request. The batch&#39;s response is a dictionary uses keys corresponding those provided in the request, with response objects containing a status code, response headers, and the response body. A request can alternatively specify a request template in place of a resource. In this case, a single JsonPath may select multiple tokens, and a separate subrequest will be made from the template for each token. The responses of these subrequests will returned as the content of a single outer response.

@param batch The batch of requests.
@return map[string]Response
*/
func (a *BatchApiService) BatchExecute(batch interface{}) (map[string]Response, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     map[string]Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/batch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["x-requested-with"] = "Accept"
	// body params
	localVarPostBody = &batch
	r, err := a.client.prepareRequest(
		a.client.ctx,
		localVarPath,
		localVarHttpMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		localVarFileName,
		localVarFileBytes,
	)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
func (a *BatchApiService) StreamGetSummary(webId string, localVarOptionals map[string]interface{}) (BatchRequest, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte

		batch BatchRequest
	)

	// create path and map variables
	linkPath := fmt.Sprintf("/streams/%s/summary", webId)
	localVarPath := a.client.cfg.BasePath + linkPath

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	//TODO: refactor this to use the same method as BatchExecute
	if err := typeCheckParameter(localVarOptionals["calculationBasis"], "string", "calculationBasis"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleInterval"], "string", "sampleInterval"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleType"], "string", "sampleType"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["summaryDuration"], "string", "summaryDuration"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["timeType"], "string", "timeType"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return BatchRequest{}, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["calculationBasis"].(string); localVarOk {
		localVarQueryParams.Add("calculationBasis", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sampleInterval"].(string); localVarOk {
		localVarQueryParams.Add("sampleInterval", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sampleType"].(string); localVarOk {
		localVarQueryParams.Add("sampleType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["summaryDuration"].(string); localVarOk {
		localVarQueryParams.Add("summaryDuration", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["summaryType"].([]string); localVarOk {
		for _, t := range localVarTempParam {
			localVarQueryParams.Add("summaryType", t)
		}
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeType"].(string); localVarOk {
		localVarQueryParams.Add("timeType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}

	r, err := a.client.prepareRequest(
		a.client.ctx,
		localVarPath,
		localVarHttpMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		localVarFileName,
		localVarFileBytes,
	)
	if err != nil {
		return BatchRequest{}, err
	}

	//add URL from request to batch
	batch = BatchRequest{
		Method:   r.Method,
		Resource: r.URL.String(),
	}

	return batch, err
}

/*
	BatchApiService Returns a list of compressed values for the requested time range from the source provider.

Returned times are affected by the specified boundary type. If no values are found for the time range and conditions specified then the HTTP response will be success, with a body containing an empty array of Items. When specifying true for the includeFilteredValues parameter, consecutive filtered events are not returned. The first value that would be filtered out is returned with its time and the enumeration value \&quot;Filtered\&quot;. The next value in the collection will be the next compressed value in the specified direction that passes the filter criteria - if any. When both boundaryType and a filterExpression are specified, the events returned for the boundary condition specified are passed through the filter. If the includeFilteredValues parameter is true, the boundary values will be reported at the proper timestamps with the enumeration value \&quot;Filtered\&quot; when the filter conditions are not met at the boundary time. If the includeFilteredValues parameter is false for this case, no event is returned for the boundary time. Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.   If only recorded values with annotations are desired, the filterExpression parameter should include the filter IsSet(&#39;.&#39;, \&quot;a\&quot;).

@param webId The ID of the stream.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "associations" (string) Associated values to return in the response, separated by semicolons (;). This call supports Annotations to return events with annotation values. If this parameter is not specified, annotation values are not returned.
	@param "boundaryType" (string) An optional value that determines how the times and values of the returned end points are determined. The default is &#39;Inside&#39;.
	@param "desiredUnits" (string) The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure.
	@param "endTime" (string) An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
	@param "filterExpression" (string) An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. The default is no filtering.
	@param "includeFilteredValues" (bool) Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted.
	@param "maxCount" (int32) The maximum number of values to be returned. The default is 1000.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "startTime" (string) An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set.
	@param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.

@return BatchRequest
*/
func (a *BatchApiService) StreamGetRecorded(webId string, localVarOptionals map[string]interface{}) (BatchRequest, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte

		batch BatchRequest
	)

	// create path and map variables
	linkPath := fmt.Sprintf("/streams/%s/recorded", webId)
	localVarPath := a.client.cfg.BasePath + linkPath

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["associations"], "string", "associations"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["boundaryType"], "string", "boundaryType"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["includeFilteredValues"], "bool", "includeFilteredValues"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return BatchRequest{}, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["associations"].(string); localVarOk {
		localVarQueryParams.Add("associations", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["boundaryType"].(string); localVarOk {
		localVarQueryParams.Add("boundaryType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["desiredUnits"].(string); localVarOk {
		localVarQueryParams.Add("desiredUnits", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeFilteredValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeFilteredValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}

	r, err := a.client.prepareRequest(
		a.client.ctx,
		localVarPath,
		localVarHttpMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		localVarFileName,
		localVarFileBytes,
	)
	if err != nil {
		return BatchRequest{}, err
	}

	//add URL from request to batch
	batch = BatchRequest{
		Method:   r.Method,
		Resource: r.URL.String(),
	}

	return batch, err
}

/*
	StreamApiService Returns a single recorded value based on the passed time and retrieval mode from the stream.

If only recorded values with annotations are desired, the filterExpression parameter should include the filter IsSet(&#39;.&#39;, \&quot;a\&quot;).

@param webId The ID of the stream.
@param time The timestamp at which the value is desired.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "associations" (string) Associated values to return in the response, separated by semicolons (;). This call supports Annotations to return events with annotation values. If this parameter is not specified, annotation values are not returned.
	@param "desiredUnits" (string) The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure.
	@param "retrievalMode" (string) An optional value that determines the value to return when a value doesn&#39;t exist at the exact time specified. The default is &#39;Auto&#39;.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.

@return BatchRequest
*/
func (a *BatchApiService) StreamGetRecordedAtTime(webId string, time string, localVarOptionals map[string]interface{}) (BatchRequest, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte

		batch BatchRequest
	)

	// create path and map variables
	linkPath := fmt.Sprintf("/streams/%s/recordedattime", webId)
	localVarPath := a.client.cfg.BasePath + linkPath

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["associations"], "string", "associations"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["retrievalMode"], "string", "retrievalMode"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return BatchRequest{}, err
	}

	localVarQueryParams.Add("time", parameterToString(time, ""))
	if localVarTempParam, localVarOk := localVarOptionals["associations"].(string); localVarOk {
		localVarQueryParams.Add("associations", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["desiredUnits"].(string); localVarOk {
		localVarQueryParams.Add("desiredUnits", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["retrievalMode"].(string); localVarOk {
		localVarQueryParams.Add("retrievalMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}

	r, err := a.client.prepareRequest(
		a.client.ctx,
		localVarPath,
		localVarHttpMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		localVarFileName,
		localVarFileBytes,
	)
	if err != nil {
		return BatchRequest{}, err
	}

	//add URL from request to batch
	batch = BatchRequest{
		Method:   r.Method,
		Resource: r.URL.String(),
	}

	return batch, err
}

/*
	StreamApiService Retrieves recorded values at the specified times.

Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.   If only recorded values with annotations are desired, the filterExpression parameter should include the filter IsSet(&#39;.&#39;, \&quot;a\&quot;).

@param webId The ID of the stream.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "associations" (string) Associated values to return in the response, separated by semicolons (;). This call supports Annotations to return events with annotation values. If this parameter is not specified, annotation values are not returned.
	@param "desiredUnits" (string) The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure.
	@param "retrievalMode" (string) An optional value that determines the value to return when a value doesn&#39;t exist at the exact time specified. The default is &#39;Auto&#39;.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
	@param "time" ([]string) The timestamp at which to retrieve a recorded value. Multiple timestamps may be specified with multiple instances of the parameter.
	@param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.

@return BatchRequest
*/
func (a *BatchApiService) StreamGetRecordedAtTimes(webId string, localVarOptionals map[string]interface{}) (BatchRequest, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		batch              BatchRequest
	)

	// create path and map variables
	linkPath := fmt.Sprintf("/streams/%s/recordedattimes", webId)
	localVarPath := a.client.cfg.BasePath + linkPath

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["associations"], "string", "associations"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["retrievalMode"], "string", "retrievalMode"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return BatchRequest{}, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["associations"].(string); localVarOk {
		localVarQueryParams.Add("associations", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["desiredUnits"].(string); localVarOk {
		localVarQueryParams.Add("desiredUnits", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["retrievalMode"].(string); localVarOk {
		localVarQueryParams.Add("retrievalMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["time"].([]string); localVarOk {
		for _, t := range localVarTempParam {
			localVarQueryParams.Add("time", t)
		}
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}

	r, err := a.client.prepareRequest(
		a.client.ctx,
		localVarPath,
		localVarHttpMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		localVarFileName,
		localVarFileBytes,
	)
	if err != nil {
		return BatchRequest{}, err
	}

	//add URL from request to batch
	batch = BatchRequest{
		Method:   r.Method,
		Resource: r.URL.String(),
	}

	return batch, err
}

/*
	StreamApiService Returns the value of the stream at the specified time. By default, this is usually the current value.

@param webId The ID of the stream.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "desiredUnits" (string) The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "time" (string) An optional time. The default time context is determined from the owning object - for example, the time range of the event frame or transfer which holds this attribute. Otherwise, the implementation of the Data Reference determines the meaning of no context. For Points or simply configured PI Point Data References, this means the snapshot value of the PI Point on the Data Server.
	@param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.

@return BatchRequest
*/
func (a *BatchApiService) StreamGetValue(webId string, localVarOptionals map[string]interface{}) (BatchRequest, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		batch              BatchRequest
	)

	// create path and map variables
	linkPath := fmt.Sprintf("/streams/%s/value", webId)
	localVarPath := a.client.cfg.BasePath + linkPath

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["time"], "string", "time"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return BatchRequest{}, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["desiredUnits"].(string); localVarOk {
		localVarQueryParams.Add("desiredUnits", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["time"].(string); localVarOk {
		localVarQueryParams.Add("time", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}

	r, err := a.client.prepareRequest(
		a.client.ctx,
		localVarPath,
		localVarHttpMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		localVarFileName,
		localVarFileBytes,
	)
	if err != nil {
		return BatchRequest{}, err
	}

	//add URL from request to batch
	batch = BatchRequest{
		Method:   r.Method,
		Resource: r.URL.String(),
	}

	return batch, err
}

/*
	StreamApiService Updates a value for the specified stream.

@param webId The ID of the stream.
@param value The value to add or update.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "bufferOption" (string) The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;.
	@param "updateOption" (string) The desired AFUpdateOption. The default is &#39;Replace&#39;. This parameter is ignored if the attribute is a configuration item.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return BatchRequest
*/
func (a *BatchApiService) StreamUpdateValue(webId string, value TimedValue, localVarOptionals map[string]interface{}) (BatchRequest, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		batch              BatchRequest
	)

	// create path and map variables
	linkPath := fmt.Sprintf("/streams/%s/value", webId)
	localVarPath := a.client.cfg.BasePath + linkPath

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["bufferOption"], "string", "bufferOption"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["updateOption"], "string", "updateOption"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return BatchRequest{}, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["bufferOption"].(string); localVarOk {
		localVarQueryParams.Add("bufferOption", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["updateOption"].(string); localVarOk {
		localVarQueryParams.Add("updateOption", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}

	localVarPostBody = &value

	r, err := a.client.prepareRequest(
		a.client.ctx,
		localVarPath,
		localVarHttpMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		localVarFileName,
		localVarFileBytes,
	)
	//TODO: check if this is the correct way to handle the body

	//add URL from request to batch
	batch = BatchRequest{
		Method:   r.Method,
		Resource: r.URL.String(),
		Content:  string(localVarPostBody.([]byte)),
	}

	return batch, err
}

/*
	StreamApiService Updates multiple values for the specified stream.

@param webId The ID of the stream.
@param values The values to add or update.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "bufferOption" (string) The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;.
	@param "updateOption" (string) The desired AFUpdateOption. The default is &#39;Replace&#39;.

@return BatchRequest
*/
func (a *BatchApiService) StreamUpdateValues(webId string, values []TimedValue, localVarOptionals map[string]interface{}) (BatchRequest, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		batch              BatchRequest
	)

	// create path and map variables
	linkPath := fmt.Sprintf("/streams/%s/recorded", webId)
	localVarPath := a.client.cfg.BasePath + linkPath

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["bufferOption"], "string", "bufferOption"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["updateOption"], "string", "updateOption"); err != nil {
		return BatchRequest{}, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["bufferOption"].(string); localVarOk {
		localVarQueryParams.Add("bufferOption", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["updateOption"].(string); localVarOk {
		localVarQueryParams.Add("updateOption", parameterToString(localVarTempParam, ""))
	}

	// body params
	localVarPostBody = &values
	r, err := a.client.prepareRequest(
		a.client.ctx,
		localVarPath,
		localVarHttpMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		localVarFileName,
		localVarFileBytes,
	)
	//TODO: check if this is the correct way to handle the body

	//add URL from request to batch
	batch = BatchRequest{
		Method:   r.Method,
		Resource: r.URL.String(),
		Content:  string(localVarPostBody.([]byte)),
	}

	return batch, err
}

/*
	CalculationApiService Returns the result of evaluating the expression at the specified timestamps.

@param optional (nil or map[string]interface{}) with one or more of:

	@param "expression" (string) A string containing the expression to be evaluated. The syntax for the expression generally follows the Performance Equation syntax as described in the PI Server documentation, with the exception that expressions which target AF objects use attribute names in place of tag names in the equation.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
	@param "time" ([]string) A list of timestamps at which to calculate the expression.
	@param "webId" (string) The ID of the target object of the expression. A target object can be a Data Server, a database, an element, an event frame or an attribute. References to attributes or points are based on the target. If this parameter is not provided, the target object is set to null.

@return TimedValues
*/
func (a *BatchApiService) CalculationGetAtTimes(localVarOptionals map[string]interface{}) (BatchRequest, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		batch              BatchRequest
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/calculation/times"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["expression"], "string", "expression"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["webId"], "string", "webId"); err != nil {
		return BatchRequest{}, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["expression"].(string); localVarOk {
		localVarQueryParams.Add("expression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["time"].([]string); localVarOk {
		for _, t := range localVarTempParam {
			localVarQueryParams.Add("time", t)
		}
	}
	if localVarTempParam, localVarOk := localVarOptionals["webId"].(string); localVarOk {
		localVarQueryParams.Add("webId", parameterToString(localVarTempParam, ""))
	}

	r, err := a.client.prepareRequest(
		a.client.ctx,
		localVarPath,
		localVarHttpMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		localVarFileName,
		localVarFileBytes,
	)
	if err != nil {
		return BatchRequest{}, err
	}

	//add URL from request to batch
	batch = BatchRequest{
		Method:   r.Method,
		Resource: r.URL.String(),
	}

	return batch, err
}

/*
	CalculationApiService Returns the result of evaluating the expression over the time range from the start time to the end time. The time range is first divided into a number of summary intervals. Then the calculation is performed for the specified summaries over each interval.

@param optional (nil or map[string]interface{}) with one or more of:

	@param "calculationBasis" (string) Specifies the method of evaluating the data over the time range. The default is &#39;TimeWeighted&#39;.
	@param "endTime" (string) An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
	@param "expression" (string) A string containing the expression to be evaluated. The syntax for the expression generally follows the Performance Equation syntax as described in the PI Server documentation, with the exception that expressions which target AF objects use attribute names in place of tag names in the equation.
	@param "sampleInterval" (string) A time span specifies how often the filter expression is evaluated when computing the summary for an interval, if the sampleType is &#39;Interval&#39;.
	@param "sampleType" (string) A flag which specifies one or more summaries to compute for each interval over the time range. The default is &#39;ExpressionRecordedValues&#39;.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "startTime" (string) An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set.
	@param "summaryDuration" (string) The duration of each summary interval.
	@param "summaryType" ([]string) Specifies the kinds of summaries to produce over the range. The default is &#39;Total&#39;. Multiple summary types may be specified by using multiple instances of summaryType.
	@param "timeType" (string) Specifies how to calculate the timestamp for each interval. The default is &#39;Auto&#39;.
	@param "webId" (string) The ID of the target object of the expression. A target object can be a Data Server, a database, an element, an event frame or an attribute. References to attributes or points are based on the target. If this parameter is not provided, the target object is set to null.

@return ItemsSummaryValue
*/
func (a *BatchApiService) CalculationGetSummary(localVarOptionals map[string]interface{}) (BatchRequest, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		batch              BatchRequest
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/calculation/summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["calculationBasis"], "string", "calculationBasis"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["expression"], "string", "expression"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleInterval"], "string", "sampleInterval"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleType"], "string", "sampleType"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["summaryDuration"], "string", "summaryDuration"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["timeType"], "string", "timeType"); err != nil {
		return BatchRequest{}, err
	}
	if err := typeCheckParameter(localVarOptionals["webId"], "string", "webId"); err != nil {
		return BatchRequest{}, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["calculationBasis"].(string); localVarOk {
		localVarQueryParams.Add("calculationBasis", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["expression"].(string); localVarOk {
		localVarQueryParams.Add("expression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sampleInterval"].(string); localVarOk {
		localVarQueryParams.Add("sampleInterval", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sampleType"].(string); localVarOk {
		localVarQueryParams.Add("sampleType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["summaryDuration"].(string); localVarOk {
		localVarQueryParams.Add("summaryDuration", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["summaryType"].([]string); localVarOk {
		for _, t := range localVarTempParam {
			localVarQueryParams.Add("summaryType", t)
		}
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeType"].(string); localVarOk {
		localVarQueryParams.Add("timeType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webId"].(string); localVarOk {
		localVarQueryParams.Add("webId", parameterToString(localVarTempParam, ""))
	}

	r, err := a.client.prepareRequest(
		a.client.ctx,
		localVarPath,
		localVarHttpMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		localVarFileName,
		localVarFileBytes,
	)
	if err != nil {
		return BatchRequest{}, err
	}

	//add URL from request to batch
	batch = BatchRequest{
		Method:   r.Method,
		Resource: r.URL.String(),
	}

	return batch, err
}
