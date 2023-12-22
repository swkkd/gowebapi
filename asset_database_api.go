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

type AssetDatabaseApiService service

/*
	AssetDatabaseApiService Add a reference to an existing element to the specified database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database which the referenced element will be added to.
@param referencedElementWebId The ID of the referenced element. Multiple referenced elements may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "referenceType" (string) The name of the reference type between the parent and the referenced element. This must be a \&quot;strong\&quot; reference type. The default is \&quot;parent-child\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseAddReferencedElement(ctx context.Context, webId string, referencedElementWebId []string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/referencedelements"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["referenceType"], "string", "referenceType"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("referencedElementWebId", parameterToString(referencedElementWebId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["referenceType"].(string); localVarOk {
		localVarQueryParams.Add("referenceType", parameterToString(localVarTempParam, ""))
	}
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create an analysis category at the Asset Database root.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database in which to create the analysis category.
@param analysisCategory The new analysis category definition.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateAnalysisCategory(ctx context.Context, webId string, analysisCategory AnalysisCategory, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/analysiscategories"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	localVarPostBody = &analysisCategory
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create an analysis template at the Asset Database root.

Analyses that are based on an analysis template will inherit characteristics defined in the template.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database in which to create the analysis template.
@param template The new analysis template definition.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateAnalysisTemplate(ctx context.Context, webId string, template AnalysisTemplate, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/analysistemplates"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	localVarPostBody = &template
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create an attribute category at the Asset Database root.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database in which to create the attribute category.
@param attributeCategory The new attribute category definition.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateAttributeCategory(ctx context.Context, webId string, attributeCategory AttributeCategory, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/attributecategories"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	localVarPostBody = &attributeCategory
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create a child element.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the asset database on which to create the element.
@param element The new element definition.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateElement(ctx context.Context, webId string, element Element, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/elements"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	localVarPostBody = &element
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create an element category at the Asset Database root.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database in which to create the element category.
@param elementCategory The new element category definition.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateElementCategory(ctx context.Context, webId string, elementCategory ElementCategory, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/elementcategories"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	localVarPostBody = &elementCategory
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create a template at the Asset Database root. Specify InstanceType of \&quot;Element\&quot; or \&quot;EventFrame\&quot; to create element or event frame template respectively. Only these two types of templates can be created.

Elements and event frames that are based on an element template will inherit characteristics defined in the template.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database in which to create the element template.
@param template The new element template definition.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateElementTemplate(ctx context.Context, webId string, template ElementTemplate, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/elementtemplates"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	localVarPostBody = &template
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create an enumeration set at the Asset Database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database in which to create the enumeration set.
@param enumerationSet The new enumeration set definition.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateEnumerationSet(ctx context.Context, webId string, enumerationSet EnumerationSet, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/enumerationsets"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	localVarPostBody = &enumerationSet
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create an event frame.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database on which to create the event frame.
@param eventFrame The new event frame definition.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateEventFrame(ctx context.Context, webId string, eventFrame EventFrame, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/eventframes"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	localVarPostBody = &eventFrame
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create a security entry owned by the asset database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the asset database where the security entry will be created.
@param securityEntry The new security entry definition. The full list of allow and deny rights must be supplied.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "applyToChildren" (bool) If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change.
	@param "securityItem" (string) The security item of the desired security entries to be created. If the parameter is not specified, security entries of the &#39;Default&#39; security item will be created.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateSecurityEntry(ctx context.Context, webId string, securityEntry SecurityEntry, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/securityentries"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["applyToChildren"], "bool", "applyToChildren"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["securityItem"], "string", "securityItem"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["applyToChildren"].(bool); localVarOk {
		localVarQueryParams.Add("applyToChildren", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["securityItem"].(string); localVarOk {
		localVarQueryParams.Add("securityItem", parameterToString(localVarTempParam, ""))
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
	localVarPostBody = &securityEntry
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create a table on the Asset Database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database in which to create the table.
@param table The new table definition.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateTable(ctx context.Context, webId string, table Table, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/tables"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	localVarPostBody = &table
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Create a table category on the Asset Database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database in which to create the table category.
@param tableCategory The new table category definition.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseCreateTableCategory(ctx context.Context, webId string, tableCategory TableCategory, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/tablecategories"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
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
	localVarPostBody = &tableCategory
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Delete an asset database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database.
@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseDelete(ctx context.Context, webId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Delete a security entry owned by the asset database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param name The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username.
@param webId The ID of the asset database where the security entry will be deleted.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "applyToChildren" (bool) If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change.
	@param "securityItem" (string) The security item of the desired security entries to be deleted. If the parameter is not specified, security entries of the &#39;Default&#39; security item will be deleted.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseDeleteSecurityEntry(ctx context.Context, name string, webId string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/securityentries/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["applyToChildren"], "bool", "applyToChildren"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["securityItem"], "string", "securityItem"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["applyToChildren"].(bool); localVarOk {
		localVarQueryParams.Add("applyToChildren", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["securityItem"].(string); localVarOk {
		localVarQueryParams.Add("securityItem", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Export the asset database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "endTime" (string) The latest ending time for AFEventFrame, AFTransfer, and AFCase objects that may be part of the export. The default is &#39;*&#39;.
	@param "exportMode" ([]string) Indicates the type of export to perform. The default is &#39;StrongReferences&#39;. Multiple export modes may be specified by using multiple instances of exportMode.
	@param "startTime" (string) The earliest starting time for AFEventFrame, AFTransfer, and AFCase objects that may be part of the export. The default is &#39;*-30d&#39;.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseExport(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/export"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["exportMode"].([]string); localVarOk {
		localVarQueryParams.Add("exportMode", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve analyses based on the specified conditions.

Users can search for the analyses based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the analyses that match the default search.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database to search for the analyses.
@param field Specifies which of the object&#39;s properties are searched. Multiple search fields may be specified with multiple instances of the parameter. The default is &#39;Name&#39;.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "maxCount" (int32) The maximum number of objects to be returned per call (page size). The default is 1000.
	@param "query" (string) The query string used for finding analyses. The default is null.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "sortField" (string) The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
	@param "startIndex" (int32) The starting index (zero based) of the items to be returned. The default is 0.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsAnalysis
*/
func (a *AssetDatabaseApiService) AssetDatabaseFindAnalyses(ctx context.Context, webId string, field []string, localVarOptionals map[string]interface{}) (ItemsAnalysis, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsAnalysis
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/analyses"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["query"], "string", "query"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startIndex"], "int32", "startIndex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("field", parameterToString(field, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["query"].(string); localVarOk {
		localVarQueryParams.Add("query", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startIndex"].(int32); localVarOk {
		localVarQueryParams.Add("startIndex", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieves a list of element attributes matching the specified filters from the specified asset database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the asset database to use as the root of the search.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "attributeCategory" (string) Specify that returned attributes must have this category. The default is no filter.
	@param "attributeDescriptionFilter" (string) The attribute description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter.
	@param "attributeNameFilter" (string) The attribute name filter string used for finding objects. The default is no filter.
	@param "attributeType" (string) Specify that returned attributes&#39; value type must be this value type. The default is no filter.
	@param "elementCategory" (string) Specify that the owner of the returned attributes must have this category. The default is no filter.
	@param "elementDescriptionFilter" (string) The element description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter.
	@param "elementNameFilter" (string) The element name filter string used for finding objects. The default is no filter.
	@param "elementTemplate" (string) Specify that the owner of the returned attributes must have this template or a template derived from this template. The default is no filter.
	@param "elementType" (string) Specify that the element of the returned attributes must have this AFElementType. The default is no filter.
	@param "maxCount" (int32) The maximum number of objects to be returned (the page size). The default is 1000.
	@param "searchFullHierarchy" (bool) Specifies if the search should include objects nested further than immediate children of the given resource. The default is &#39;false&#39;.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "sortField" (string) The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
	@param "startIndex" (int32) The starting index (zero based) of the items to be returned. The default is 0.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsAttribute
*/
func (a *AssetDatabaseApiService) AssetDatabaseFindElementAttributes(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsAttribute, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsAttribute
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/elementattributes"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["attributeCategory"], "string", "attributeCategory"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["attributeDescriptionFilter"], "string", "attributeDescriptionFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["attributeNameFilter"], "string", "attributeNameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["attributeType"], "string", "attributeType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["elementCategory"], "string", "elementCategory"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["elementDescriptionFilter"], "string", "elementDescriptionFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["elementNameFilter"], "string", "elementNameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["elementTemplate"], "string", "elementTemplate"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["elementType"], "string", "elementType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startIndex"], "int32", "startIndex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["attributeCategory"].(string); localVarOk {
		localVarQueryParams.Add("attributeCategory", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["attributeDescriptionFilter"].(string); localVarOk {
		localVarQueryParams.Add("attributeDescriptionFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["attributeNameFilter"].(string); localVarOk {
		localVarQueryParams.Add("attributeNameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["attributeType"].(string); localVarOk {
		localVarQueryParams.Add("attributeType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["elementCategory"].(string); localVarOk {
		localVarQueryParams.Add("elementCategory", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["elementDescriptionFilter"].(string); localVarOk {
		localVarQueryParams.Add("elementDescriptionFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["elementNameFilter"].(string); localVarOk {
		localVarQueryParams.Add("elementNameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["elementTemplate"].(string); localVarOk {
		localVarQueryParams.Add("elementTemplate", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["elementType"].(string); localVarOk {
		localVarQueryParams.Add("elementType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startIndex"].(int32); localVarOk {
		localVarQueryParams.Add("startIndex", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieves a list of event frame attributes matching the specified filters from the specified asset database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the asset database to use as the root of the search.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "attributeCategory" (string) Specify that returned attributes must have this category. The default is no filter.
	@param "attributeDescriptionFilter" (string) The attribute description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter.
	@param "attributeNameFilter" (string) The attribute name filter string used for finding objects. The default is no filter.
	@param "attributeType" (string) Specify that returned attributes&#39; value type must be this value type. The default is no filter.
	@param "endTime" (string) A string representing the latest ending time for the event frames to be matched. The endTime must be greater than or equal to the startTime. The default is &#39;*&#39;.
	@param "eventFrameCategory" (string) Specify that the owner of the returned attributes must have this category. The default is no filter.
	@param "eventFrameDescriptionFilter" (string) The event frame description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter.
	@param "eventFrameNameFilter" (string) The event frame name filter string used for finding objects. The default is no filter.
	@param "eventFrameTemplate" (string) Specify that the owner of the returned attributes must have this template or a template derived from this template. The default is no filter.
	@param "maxCount" (int32) The maximum number of objects to be returned (the page size). The default is 1000.
	@param "referencedElementNameFilter" (string) The name query string which must match the name of a referenced element. The default is no filter.
	@param "searchFullHierarchy" (bool) Specifies if the search should include objects nested further than immediate children of the given resource. The default is &#39;false&#39;.
	@param "searchMode" (string) Determines how the startTime and endTime parameters are treated when searching for event frames. The default is &#39;Overlapped&#39;.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "sortField" (string) The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
	@param "startIndex" (int32) The starting index (zero based) of the items to be returned. The default is 0.
	@param "startTime" (string) A string representing the earliest starting time for the event frames to be matched. startTime must be less than or equal to the endTime. The default is &#39;*-8h&#39;.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsAttribute
*/
func (a *AssetDatabaseApiService) AssetDatabaseFindEventFrameAttributes(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsAttribute, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsAttribute
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/eventframeattributes"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["attributeCategory"], "string", "attributeCategory"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["attributeDescriptionFilter"], "string", "attributeDescriptionFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["attributeNameFilter"], "string", "attributeNameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["attributeType"], "string", "attributeType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["eventFrameCategory"], "string", "eventFrameCategory"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["eventFrameDescriptionFilter"], "string", "eventFrameDescriptionFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["eventFrameNameFilter"], "string", "eventFrameNameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["eventFrameTemplate"], "string", "eventFrameTemplate"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["referencedElementNameFilter"], "string", "referencedElementNameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchMode"], "string", "searchMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startIndex"], "int32", "startIndex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["attributeCategory"].(string); localVarOk {
		localVarQueryParams.Add("attributeCategory", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["attributeDescriptionFilter"].(string); localVarOk {
		localVarQueryParams.Add("attributeDescriptionFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["attributeNameFilter"].(string); localVarOk {
		localVarQueryParams.Add("attributeNameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["attributeType"].(string); localVarOk {
		localVarQueryParams.Add("attributeType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["eventFrameCategory"].(string); localVarOk {
		localVarQueryParams.Add("eventFrameCategory", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["eventFrameDescriptionFilter"].(string); localVarOk {
		localVarQueryParams.Add("eventFrameDescriptionFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["eventFrameNameFilter"].(string); localVarOk {
		localVarQueryParams.Add("eventFrameNameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["eventFrameTemplate"].(string); localVarOk {
		localVarQueryParams.Add("eventFrameTemplate", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["referencedElementNameFilter"].(string); localVarOk {
		localVarQueryParams.Add("referencedElementNameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchMode"].(string); localVarOk {
		localVarQueryParams.Add("searchMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startIndex"].(int32); localVarOk {
		localVarQueryParams.Add("startIndex", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve an Asset Database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return AssetDatabase
*/
func (a *AssetDatabaseApiService) AssetDatabaseGet(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (AssetDatabase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     AssetDatabase
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve analysis categories for a given Asset Database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsAnalysisCategory
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetAnalysisCategories(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsAnalysisCategory, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsAnalysisCategory
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/analysiscategories"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve analysis templates based on the specified criteria. By default, all analysis templates in the specified Asset Database are returned.

Users can search for the analysis templates based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the templates that match the default search.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database to search.
@param field Specifies which of the object&#39;s properties are searched. Multiple search fields may be specified with multiple instances of the parameter. The default is &#39;Name&#39;.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "maxCount" (int32) The maximum number of objects to be returned per call (page size). The default is 1000.
	@param "query" (string) The query string used for finding objects. The default is no query string.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "sortField" (string) The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsAnalysisTemplate
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetAnalysisTemplates(ctx context.Context, webId string, field []string, localVarOptionals map[string]interface{}) (ItemsAnalysisTemplate, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsAnalysisTemplate
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/analysistemplates"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["query"], "string", "query"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("field", parameterToString(field, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["query"].(string); localVarOk {
		localVarQueryParams.Add("query", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve attribute categories for a given Asset Database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsAttributeCategory
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetAttributeCategories(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsAttributeCategory, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsAttributeCategory
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/attributecategories"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve an Asset Database by path.

This method returns an asset database based on the hierarchical path associated with it, and should be used when a path has been received from a separate part of the PI System for use in the PI Web API. Users should primarily search with the WebID when available.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param path The path to the database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return AssetDatabase
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetByPath(ctx context.Context, path string, localVarOptionals map[string]interface{}) (AssetDatabase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     AssetDatabase
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("path", parameterToString(path, ""))
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve element categories for a given Asset Database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsElementCategory
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetElementCategories(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsElementCategory, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsElementCategory
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/elementcategories"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve element templates based on the specified criteria. Only templates of instance type \&quot;Element\&quot; and \&quot;EventFrame\&quot; are returned. By default, all element and event frame templates in the specified Asset Database are returned.

Users can search for the element and event frame template based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the templates that match the default search.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database to search.
@param field Specifies which of the object&#39;s properties are searched. Multiple search fields may be specified with multiple instances of the parameter. The default is &#39;Name&#39;.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "maxCount" (int32) The maximum number of objects to be returned per call (page size). The default is 1000.
	@param "query" (string) The query string used for finding objects. The default is no query string.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "sortField" (string) The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsElementTemplate
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetElementTemplates(ctx context.Context, webId string, field []string, localVarOptionals map[string]interface{}) (ItemsElementTemplate, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsElementTemplate
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/elementtemplates"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["query"], "string", "query"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("field", parameterToString(field, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["query"].(string); localVarOk {
		localVarQueryParams.Add("query", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve elements based on the specified conditions. By default, this method selects immediate children of the specified asset database.

Users can search for the elements based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the elements that match the default search.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database to use as the root of the search.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "categoryName" (string) Specify that returned elements must have this category. The default is no category filter.
	@param "descriptionFilter" (string) Specify that returned elements must have this description. The default is no description filter.
	@param "elementType" (string) Specify that returned elements must have this type. The default type is &#39;Any&#39;.
	@param "maxCount" (int32) The maximum number of objects to be returned per call (page size). The default is 1000.
	@param "nameFilter" (string) The name query string used for finding objects. The default is no filter.
	@param "searchFullHierarchy" (bool) Specifies if the search should include objects nested further than the immediate children of the searchRoot. The default is &#39;false&#39;.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "sortField" (string) The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
	@param "startIndex" (int32) The starting index (zero based) of the items to be returned. The default is 0.
	@param "templateName" (string) Specify that returned elements must have this template or a template derived from this template. The default is no template filter.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsElement
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetElements(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsElement, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsElement
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/elements"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["descriptionFilter"], "string", "descriptionFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["elementType"], "string", "elementType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startIndex"], "int32", "startIndex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["descriptionFilter"].(string); localVarOk {
		localVarQueryParams.Add("descriptionFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["elementType"].(string); localVarOk {
		localVarQueryParams.Add("elementType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startIndex"].(int32); localVarOk {
		localVarQueryParams.Add("startIndex", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve enumeration sets for given asset database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsEnumerationSet
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetEnumerationSets(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsEnumerationSet, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsEnumerationSet
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/enumerationsets"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve event frames based on the specified conditions. By default, returns all children of the specified root resource that have been active in the past 8 hours.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the asset database to use as the root of the search.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "canBeAcknowledged" (bool) Specify the returned event frames&#39; canBeAcknowledged property. The default is no canBeAcknowledged filter.
	@param "categoryName" (string) Specify that returned event frames must have this category. The default is no category filter.
	@param "endTime" (string) The ending time for the search. The endTime must be greater than or equal to the startTime. The searchMode parameter will control whether the comparison will be performed against the event frame&#39;s startTime or endTime. The default is &#39;*&#39; if searchMode is not one of the &#39;Backward*&#39; or &#39;Forward*&#39; values.
	@param "isAcknowledged" (bool) Specify the returned event frames&#39; isAcknowledged property. The default no isAcknowledged filter.
	@param "maxCount" (int32) The maximum number of objects to be returned per call (page size). The default is 1000.
	@param "nameFilter" (string) The name query string used for finding event frames. The default is no filter.
	@param "referencedElementNameFilter" (string) The name query string which must match the name of a referenced element. The default is no filter.
	@param "referencedElementTemplateName" (string) Specify that returned event frames must have an element in the event frame&#39;s referenced elements collection that derives from the template. Specify this parameter by name.
	@param "searchFullHierarchy" (bool) Specifies whether the search should include objects nested further than the immediate children of the search root. The default is &#39;false&#39;.
	@param "searchMode" (string) Determines how the startTime and endTime parameters are treated when searching for event frame objects to be included in the returned collection. If this parameter is one of the &#39;Backward*&#39; or &#39;Forward*&#39; values, none of endTime, sortField, or sortOrder may be specified. The default is &#39;Overlapped&#39;.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "severity" ([]string) Specify that returned event frames must have this severity. Multiple severity values may be specified with multiple instances of the parameter. The default is no severity filter.
	@param "sortField" (string) The field or property of the object used to sort the returned collection. The default is &#39;Name&#39; if searchMode is not one of the &#39;Backward*&#39; or &#39;Forward*&#39; values.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39; if searchMode is not one of the &#39;Backward*&#39; or &#39;Forward*&#39; values.
	@param "startIndex" (int32) The starting index (zero based) of the items to be returned. The default is 0.
	@param "startTime" (string) The starting time for the search. startTime must be less than or equal to the endTime. The searchMode parameter will control whether the comparison will be performed against the event frame&#39;s startTime or endTime. The default is &#39;*-8h&#39;.
	@param "templateName" (string) Specify that returned event frames must have this template or a template derived from this template. The default is no template filter. Specify this parameter by name.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsEventFrame
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetEventFrames(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsEventFrame, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsEventFrame
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/eventframes"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["canBeAcknowledged"], "bool", "canBeAcknowledged"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["isAcknowledged"], "bool", "isAcknowledged"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["referencedElementNameFilter"], "string", "referencedElementNameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["referencedElementTemplateName"], "string", "referencedElementTemplateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchMode"], "string", "searchMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startIndex"], "int32", "startIndex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["canBeAcknowledged"].(bool); localVarOk {
		localVarQueryParams.Add("canBeAcknowledged", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["isAcknowledged"].(bool); localVarOk {
		localVarQueryParams.Add("isAcknowledged", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["referencedElementNameFilter"].(string); localVarOk {
		localVarQueryParams.Add("referencedElementNameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["referencedElementTemplateName"].(string); localVarOk {
		localVarQueryParams.Add("referencedElementTemplateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchMode"].(string); localVarOk {
		localVarQueryParams.Add("searchMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["severity"].([]string); localVarOk {
		localVarQueryParams.Add("severity", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startIndex"].(int32); localVarOk {
		localVarQueryParams.Add("startIndex", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve referenced elements based on the specified conditions. By default, this method selects all referenced elements at the root level of the asset database.

Users can search for the referenced elements based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the elements that match the default search.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the resource to use as the root of the search.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "categoryName" (string) Specify that returned elements must have this category. The default is no category filter.
	@param "descriptionFilter" (string) Specify that returned elements must have this description. The default is no description filter.
	@param "elementType" (string) Specify that returned elements must have this type. The default type is &#39;Any&#39;.
	@param "maxCount" (int32) The maximum number of objects to be returned per call (page size). The default is 1000.
	@param "nameFilter" (string) The name query string used for finding objects. The default is no filter.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "sortField" (string) The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;.
	@param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
	@param "startIndex" (int32) The starting index (zero based) of the items to be returned. The default is 0.
	@param "templateName" (string) Specify that returned elements must have this template or a template derived from this template. The default is no template filter.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsElement
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetReferencedElements(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsElement, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsElement
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/referencedelements"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["descriptionFilter"], "string", "descriptionFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["elementType"], "string", "elementType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startIndex"], "int32", "startIndex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["descriptionFilter"].(string); localVarOk {
		localVarQueryParams.Add("descriptionFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["elementType"].(string); localVarOk {
		localVarQueryParams.Add("elementType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startIndex"].(int32); localVarOk {
		localVarQueryParams.Add("startIndex", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Get the security information of the specified security item associated with the asset database for a specified user.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the asset database for the security to be checked.
@param securityItem The security item of the desired security information to be returned. Multiple security items may be specified with multiple instances of the parameter. If the parameter is not specified, only &#39;Default&#39; security item of the security information will be returned.
@param userIdentity The user identity for the security information to be checked. Multiple security identities may be specified with multiple instances of the parameter. If the parameter is not specified, only the current user&#39;s security rights will be returned.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "forceRefresh" (bool) Indicates if the security cache should be refreshed before getting security information. The default is &#39;false&#39;.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsSecurityRights
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetSecurity(ctx context.Context, webId string, securityItem []string, userIdentity []string, localVarOptionals map[string]interface{}) (ItemsSecurityRights, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsSecurityRights
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/security"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["forceRefresh"], "bool", "forceRefresh"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("securityItem", parameterToString(securityItem, "multi"))
	localVarQueryParams.Add("userIdentity", parameterToString(userIdentity, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["forceRefresh"].(bool); localVarOk {
		localVarQueryParams.Add("forceRefresh", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve the security entries of the specified security item associated with the asset database based on the specified criteria. By default, all security entries for this asset database are returned.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the asset database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "nameFilter" (string) The name query string used for filtering security entries. The default is no filter.
	@param "securityItem" (string) The security item of the desired security entries information to be returned. If the parameter is not specified, security entries of the &#39;Default&#39; security item will be returned.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsSecurityEntry
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetSecurityEntries(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsSecurityEntry, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsSecurityEntry
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/securityentries"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["securityItem"], "string", "securityItem"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["securityItem"].(string); localVarOk {
		localVarQueryParams.Add("securityItem", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve the security entry of the specified security item associated with the asset database with the specified name.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param name The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username.
@param webId The ID of the asset database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "securityItem" (string) The security item of the desired security entries information to be returned. If the parameter is not specified, security entries of the &#39;Default&#39; security item will be returned.
	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return SecurityEntry
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetSecurityEntryByName(ctx context.Context, name string, webId string, localVarOptionals map[string]interface{}) (SecurityEntry, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     SecurityEntry
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/securityentries/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["securityItem"], "string", "securityItem"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["securityItem"].(string); localVarOk {
		localVarQueryParams.Add("securityItem", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve table categories for a given Asset Database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsTableCategory
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetTableCategories(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsTableCategory, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsTableCategory
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/tablecategories"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Retrieve tables for given Asset Database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
	@param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.

@return ItemsTable
*/
func (a *AssetDatabaseApiService) AssetDatabaseGetTables(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsTable, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsTable
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/tables"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Import an asset database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the asset database.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "importMode" ([]string) Indicates the type of import to perform. The default is &#39;AllowCreate | AllowUpdate | AutoCheckIn&#39;. Multiple import modes may be specified by using multiple instances of importMode.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseImport(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/import"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarTempParam, localVarOk := localVarOptionals["importMode"].([]string); localVarOk {
		localVarQueryParams.Add("importMode", parameterToString(localVarTempParam, "multi"))
	}
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Remove a reference to an existing element from the specified database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database which the referenced element will be removed from.
@param referencedElementWebId The ID of the referenced element. Multiple referenced elements may be specified with multiple instances of the parameter.
@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseRemoveReferencedElement(ctx context.Context, webId string, referencedElementWebId []string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/referencedelements"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("referencedElementWebId", parameterToString(referencedElementWebId, "multi"))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Update an asset database by replacing items in its definition.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the database.
@param database A partial database containing the desired changes.
@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseUpdate(ctx context.Context, webId string, database AssetDatabase) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

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
	// body params
	localVarPostBody = &database
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	AssetDatabaseApiService Update a security entry owned by the asset database.

* @param ctx context.Context for authentication, logging, tracing, etc.
@param name The name of the security entry.
@param webId The ID of the asset database where the security entry will be updated.
@param securityEntry The new security entry definition. The full list of allow and deny rights must be supplied or they will be removed.
@param optional (nil or map[string]interface{}) with one or more of:

	@param "applyToChildren" (bool) If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change.
	@param "securityItem" (string) The security item of the desired security entries to be updated. If the parameter is not specified, security entries of the &#39;Default&#39; security item will be updated.

@return
*/
func (a *AssetDatabaseApiService) AssetDatabaseUpdateSecurityEntry(ctx context.Context, name string, webId string, securityEntry SecurityEntry, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/assetdatabases/{webId}/securityentries/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["applyToChildren"], "bool", "applyToChildren"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["securityItem"], "string", "securityItem"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["applyToChildren"].(bool); localVarOk {
		localVarQueryParams.Add("applyToChildren", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["securityItem"].(string); localVarOk {
		localVarQueryParams.Add("securityItem", parameterToString(localVarTempParam, ""))
	}
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
	// body params
	localVarPostBody = &securityEntry
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
