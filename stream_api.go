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

type StreamApiService service

/*
	StreamApiService Opens a channel that will send messages about any value changes for the specified stream.

@param webId The ID of the stream.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "heartbeatRate" (int32) HeartbeatRate is an integer multiple of the Polling Interval. It specifies the rate at which a client will receive an empty message if there are no data updates. It can be used to check that the connection is still alive. Zero/negative values correspond to no heartbeat. By default, no empty messages will be sent to the user.
	@param "includeInitialValues" (bool) Specified if the channel should send a message with the current value of the stream after the connection is opened. The default is &#39;false&#39;.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *StreamApiService) StreamGetChannel(webId string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/channel"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["heartbeatRate"], "int32", "heartbeatRate"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeInitialValues"], "bool", "includeInitialValues"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["heartbeatRate"].(int32); localVarOk {
		localVarQueryParams.Add("heartbeatRate", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeInitialValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeInitialValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/*
	StreamApiService Returns the end-of-stream value of the stream.

@param webId The ID of the stream.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "desiredUnits" (string) The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.

@return TimedValue
*/
func (a *StreamApiService) StreamGetEnd(webId string, localVarOptionals map[string]interface{}) (TimedValue, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     TimedValue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/end"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["desiredUnits"].(string); localVarOk {
		localVarQueryParams.Add("desiredUnits", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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

/*
	StreamApiService Retrieves interpolated values over the specified time range at the specified sampling interval.

Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.

@param webId The ID of the stream.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "desiredUnits" (string) The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure.
	@param "endTime" (string) An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
	@param "filterExpression" (string) An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering.
	@param "includeFilteredValues" (bool) Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted.
	@param "interval" (string) The sampling interval, in AFTimeSpan format.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "startTime" (string) An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set.
	@param "syncTime" (string) An optional start time anchor, in AFTime format. When specified, interpolated data retrieval will use the sync time as the origin for calculating the interval times.
	@param "syncTimeBoundaryType" (string) An optional string specifying the boundary type to use when applying a syncTime. The allowed values are &#39;Inside&#39; and &#39;Outside&#39;. The default is &#39;Inside&#39;.
	@param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.

@return TimedValues
*/
func (a *StreamApiService) StreamGetInterpolated(webId string, localVarOptionals map[string]interface{}) (TimedValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     TimedValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/interpolated"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeFilteredValues"], "bool", "includeFilteredValues"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["interval"], "string", "interval"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["syncTime"], "string", "syncTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["syncTimeBoundaryType"], "string", "syncTimeBoundaryType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
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
	if localVarTempParam, localVarOk := localVarOptionals["interval"].(string); localVarOk {
		localVarQueryParams.Add("interval", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["syncTime"].(string); localVarOk {
		localVarQueryParams.Add("syncTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["syncTimeBoundaryType"].(string); localVarOk {
		localVarQueryParams.Add("syncTimeBoundaryType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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

/*
	StreamApiService Retrieves interpolated values over the specified time range at the specified sampling interval.

Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.

@param webId The ID of the stream.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "desiredUnits" (string) The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure.
	@param "filterExpression" (string) An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering.
	@param "includeFilteredValues" (bool) Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
	@param "time" ([]string) The timestamp at which to retrieve an interpolated value. Multiple timestamps may be specified with multiple instances of the parameter.
	@param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.

@return TimedValues
*/
func (a *StreamApiService) StreamGetInterpolatedAtTimes(webId string, localVarOptionals map[string]interface{}) (TimedValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     TimedValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/interpolatedattimes"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeFilteredValues"], "bool", "includeFilteredValues"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["desiredUnits"].(string); localVarOk {
		localVarQueryParams.Add("desiredUnits", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeFilteredValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeFilteredValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["time"].([]string); localVarOk {
		localVarQueryParams.Add("time", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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

/*
	StreamApiService Retrieves values over the specified time range suitable for plotting over the number of intervals (typically represents pixels).

For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state). Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.

@param webId The ID of the stream.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "desiredUnits" (string) The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure.
	@param "endTime" (string) An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
	@param "intervals" (int32) The number of intervals to plot over. Typically, this would be the number of horizontal pixels in the trend. The default is &#39;24&#39;. For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state).
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "startTime" (string) An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set.
	@param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.

@return TimedValues
*/
func (a *StreamApiService) StreamGetPlot(webId string, localVarOptionals map[string]interface{}) (TimedValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     TimedValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/plot"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["intervals"], "int32", "intervals"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["desiredUnits"].(string); localVarOk {
		localVarQueryParams.Add("desiredUnits", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["intervals"].(int32); localVarOk {
		localVarQueryParams.Add("intervals", parameterToString(localVarTempParam, ""))
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
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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

/*
	StreamApiService Returns a list of compressed values for the requested time range from the source provider.

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

@return ExtendedTimedValues
*/
func (a *StreamApiService) StreamGetRecorded(webId string, localVarOptionals map[string]interface{}) (ExtendedTimedValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ExtendedTimedValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/recorded"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["associations"], "string", "associations"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["boundaryType"], "string", "boundaryType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeFilteredValues"], "bool", "includeFilteredValues"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
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
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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

@return ExtendedTimedValue
*/
func (a *StreamApiService) StreamGetRecordedAtTime(webId string, time string, localVarOptionals map[string]interface{}) (ExtendedTimedValue, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ExtendedTimedValue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/recordedattime"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["associations"], "string", "associations"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["retrievalMode"], "string", "retrievalMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
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
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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

@return ExtendedTimedValues
*/
func (a *StreamApiService) StreamGetRecordedAtTimes(webId string, localVarOptionals map[string]interface{}) (ExtendedTimedValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ExtendedTimedValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/recordedattimes"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["associations"], "string", "associations"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["retrievalMode"], "string", "retrievalMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
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
		localVarQueryParams.Add("time", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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

/*
	StreamApiService Returns a summary over the specified time range for the stream.

Count is the only summary type supported on non-numeric streams. Requesting a summary for any other type will generate an error. Time-weighted totals are computed by integrating the rate tag values over the requested time range. If some of the data are bad in the time range, the calculated total is divided by the fraction of the time period for which there are good values. This approach is equivalent to assuming that during the period of bad data, the tag takes on the average values for the entire calculation time range. The PercentGood summary may be used to determine if the calculation results are suitable for the application&#39;s purposes. For time-weighted totals, if the time unit rate of the stream cannot be determined, then the value will be totaled assuming a unit of \&quot;per day\&quot; and no unit of measure will be assigned to the value. If the measured time component of the tag is not based on a day, the user of the data must convert the totalized value to the correct units.

@param webId The ID of the stream.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "calculationBasis" (string) Specifies the method of evaluating the data over the time range. The default is &#39;TimeWeighted&#39;.
	@param "endTime" (string) An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
	@param "filterExpression" (string) A string containing a filter expression. Expression variables are relative to the attribute. Use &#39;.&#39; to reference the containing attribute.
	@param "sampleInterval" (string) When the sampleType is Interval, sampleInterval specifies how often the filter expression is evaluated when computing the summary for an interval.
	@param "sampleType" (string) Defines the evaluation of an expression over a time range. The default is &#39;ExpressionRecordedValues&#39;.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "startTime" (string) An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set.
	@param "summaryDuration" (string) The duration of each summary interval. If specified in hours, minutes, seconds, or milliseconds, the summary durations will be evenly spaced UTC time intervals. Longer interval types are interpreted using wall clock rules and are time zone dependent.
	@param "summaryType" ([]string) Specifies the kinds of summaries to produce over the range. The default is &#39;Total&#39;. Multiple summary types may be specified by using multiple instances of summaryType.
	@param "timeType" (string) Specifies how to calculate the timestamp for each interval. The default is &#39;Auto&#39;.
	@param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.

@return ItemsSummaryValue
*/
func (a *StreamApiService) StreamGetSummary(webId string, localVarOptionals map[string]interface{}) (ItemsSummaryValue, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsSummaryValue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/summary"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["calculationBasis"], "string", "calculationBasis"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleInterval"], "string", "sampleInterval"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleType"], "string", "sampleType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["summaryDuration"], "string", "summaryDuration"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeType"], "string", "timeType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
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
		localVarQueryParams.Add("summaryType", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeType"].(string); localVarOk {
		localVarQueryParams.Add("timeType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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

/*
	StreamApiService Returns the value of the stream at the specified time. By default, this is usually the current value.

@param webId The ID of the stream.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "desiredUnits" (string) The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "time" (string) An optional time. The default time context is determined from the owning object - for example, the time range of the event frame or transfer which holds this attribute. Otherwise, the implementation of the Data Reference determines the meaning of no context. For Points or simply configured PI Point Data References, this means the snapshot value of the PI Point on the Data Server.
	@param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.

@return TimedValue
*/
func (a *StreamApiService) StreamGetValue(webId string, localVarOptionals map[string]interface{}) (TimedValue, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     TimedValue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["desiredUnits"], "string", "desiredUnits"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["time"], "string", "time"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
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
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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

/*
	StreamApiService Updates a value for the specified stream.

@param webId The ID of the stream.
@param value The value to add or update.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "bufferOption" (string) The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;.
	@param "updateOption" (string) The desired AFUpdateOption. The default is &#39;Replace&#39;. This parameter is ignored if the attribute is a configuration item.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *StreamApiService) StreamUpdateValue(webId string, value TimedValue, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["bufferOption"], "string", "bufferOption"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["updateOption"], "string", "updateOption"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	localVarHeaderParams["x-requested-with"] = "Accept"
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
	// body params
	localVarPostBody = &value
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/*
	StreamApiService Updates multiple values for the specified stream.

@param webId The ID of the stream.
@param values The values to add or update.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "bufferOption" (string) The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;.
	@param "updateOption" (string) The desired AFUpdateOption. The default is &#39;Replace&#39;.

@return ItemsSubstatus
*/
func (a *StreamApiService) StreamUpdateValues(webId string, values []TimedValue, localVarOptionals map[string]interface{}) (ItemsSubstatus, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsSubstatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streams/{webId}/recorded"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["bufferOption"], "string", "bufferOption"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["updateOption"], "string", "updateOption"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["bufferOption"].(string); localVarOk {
		localVarQueryParams.Add("bufferOption", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["updateOption"].(string); localVarOk {
		localVarQueryParams.Add("updateOption", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	localVarHeaderParams["x-requested-with"] = "Accept"
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
	// body params
	localVarPostBody = &values
	r, err := a.client.prepareRequest(a.client.ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
