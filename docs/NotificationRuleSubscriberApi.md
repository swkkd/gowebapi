# \NotificationRuleSubscriberApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationRuleSubscriberGetNotificationRuleSubscriber**](NotificationRuleSubscriberApi.md#NotificationRuleSubscriberGetNotificationRuleSubscriber) | **Get** /notificationrulesubscribers/{webId} | Retrieve a notification rule subscriber.
[**NotificationRuleSubscriberGetNotificationRuleSubscriberByPath**](NotificationRuleSubscriberApi.md#NotificationRuleSubscriberGetNotificationRuleSubscriberByPath) | **Get** /notificationrulesubscribers | Retrieve a notification rule subscriber by path.
[**NotificationRuleSubscriberGetNotificationRuleSubscribers**](NotificationRuleSubscriberApi.md#NotificationRuleSubscriberGetNotificationRuleSubscribers) | **Get** /notificationrulesubscribers/{webId}/notificationrulesubscribers | Retrieve notification rule subscriber subscribers.


# **NotificationRuleSubscriberGetNotificationRuleSubscriber**
> NotificationRuleSubscriber NotificationRuleSubscriberGetNotificationRuleSubscriber(a.client.ctx, webId, optional)
Retrieve a notification rule subscriber.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the resource to use as the root of the search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the resource to use as the root of the search. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**NotificationRuleSubscriber**](NotificationRuleSubscriber.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationRuleSubscriberGetNotificationRuleSubscriberByPath**
> NotificationRuleSubscriber NotificationRuleSubscriberGetNotificationRuleSubscriberByPath(a.client.ctx, path, optional)
Retrieve a notification rule subscriber by path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| The path to the notification rule subscriber. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| The path to the notification rule subscriber. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**NotificationRuleSubscriber**](NotificationRuleSubscriber.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationRuleSubscriberGetNotificationRuleSubscribers**
> ItemsNotificationRuleSubscriber NotificationRuleSubscriberGetNotificationRuleSubscribers(a.client.ctx, webId, optional)
Retrieve notification rule subscriber subscribers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the resource to use as the root of the search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the resource to use as the root of the search. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsNotificationRuleSubscriber**](Items[NotificationRuleSubscriber].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

