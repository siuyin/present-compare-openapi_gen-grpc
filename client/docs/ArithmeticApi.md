# \ArithmeticApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Sum**](ArithmeticApi.md#Sum) | **Post** /sum | Adds two numbers.



## Sum

> SumResponse Sum(ctx, sumRequest)
Adds two numbers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sumRequest** | [**SumRequest**](SumRequest.md)|  | 

### Return type

[**SumResponse**](SumResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

