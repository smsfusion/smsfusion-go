# \SMSApi

All URIs are relative to *https://api.smsfusion.com.au/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendSMS**](SMSApi.md#SendSMS) | **Get** /sms/ | Send an SMS


# **SendSMS**
> SmsResult SendSMS($key, $num, $msg, $from, $deliverby, $dlrcb, $replycb, $replyemail, $validity, $cc)

Send an SMS

Send one or more SMS


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string**| API Key as generated from the &lt;a href&#x3D;&#39;https://www.smsfusion.com.au/admin/api/&#39;&gt;admin panel&lt;/a&gt; | 
 **num** | [**[]string**](string.md)| Comma separated list of phone numbers or &lt;a href&#x3D;&#39;https://www.smsfusion.com.au/help/msisdn/&#39;&gt;MSDISDN&lt;/a&gt;&#39;s | 
 **msg** | **string**| Message content to send | 
 **from** | **string**| MSISDN or vanity alphanumeric source number | [optional] 
 **deliverby** | **string**| UTC encoded time to send the SMS | [optional] 
 **dlrcb** | **string**| HTTP or HTTPS callback URL for delivery reports. Timeout for callbacks is set to 30 seconds | [optional] 
 **replycb** | **string**| HTTP or HTTPS callback URL for replies. Timeout for callbacks is set to 30 seconds | [optional] 
 **replyemail** | **string**| Email address to send replies to | [optional] 
 **validity** | **int32**| Time in minutes to keep the SMS valid for | [optional] 
 **cc** | **string**| 2 character country code &lt;a href&#x3D;&#39;https://en.wikipedia.org/wiki/ISO_3166-2&#39;&gt;ISO 3166-2&lt;/a&gt; for formatting local numbers internationally | [optional] 

### Return type

[**SmsResult**](SMS Result.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

