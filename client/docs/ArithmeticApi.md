# \ArithmeticApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SumGet**](ArithmeticApi.md#SumGet) | **Get** /sum | Adds two numbers.



## SumGet

> SumResponse SumGet(ctx, num1, num2)
Adds two numbers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**num1** | **int64**|  | 
**num2** | **int64**|  | 

### Return type

[**SumResponse**](SumResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

