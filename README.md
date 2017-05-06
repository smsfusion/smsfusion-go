# Go API client for swagger

This is the SMS Fusion API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://www.smsfusion.com.au/help/](https://www.smsfusion.com.au/help/)

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.smsfusion.com.au/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*HLRApi* | [**GetHLR**](docs/HLRApi.md#gethlr) | **Get** /hlr/ | HLR number lookup
*HLRApi* | [**GetHLRCallback**](docs/HLRApi.md#gethlrcallback) | **Get** /hlr-callback/ | HLR number lookup with results going to a callback URL
*SMSApi* | [**SendSMS**](docs/SMSApi.md#sendsms) | **Get** /sms/ | Send an SMS


## Documentation For Models

 - [HlrCallback](docs/HlrCallback.md)
 - [HlrError](docs/HlrError.md)
 - [HlrResult](docs/HlrResult.md)
 - [OutOfCredit](docs/OutOfCredit.md)
 - [SmsResult](docs/SmsResult.md)


## Documentation For Authorization


## api_key

- **Type**: API key 
- **API key parameter name**: key
- **Location**: URL query string


## Author

support@smsfusion.com.au

