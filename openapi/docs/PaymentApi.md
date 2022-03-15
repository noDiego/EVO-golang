# \PaymentApi

All URIs are relative to *https://testcodi.multipay.mx:30801/Payments/Authorize/5.6.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeSalePost**](PaymentApi.md#AuthorizeSalePost) | **Post** /AuthorizeSale | Autorización de Venta sin Captura o Pre Autorización
[**BalanceInquiryPost**](PaymentApi.md#BalanceInquiryPost) | **Post** /BalanceInquiry | Consulta de Saldo/Deuda de cuenta/Credencial
[**BlockCancelPost**](PaymentApi.md#BlockCancelPost) | **Post** /BlockCancel | Cancelación del último bloque de transacciones realizadas.
[**BlockClosePost**](PaymentApi.md#BlockClosePost) | **Post** /BlockClose | Confirmación del último bloque de transacciones realizadas.
[**BlockCreatePost**](PaymentApi.md#BlockCreatePost) | **Post** /BlockCreate | Crea un Identificador de Bloque-Block de transacciones.
[**CancelPost**](PaymentApi.md#CancelPost) | **Post** /Cancel | Cancela una Transacción en Curso, Inciada con un GetCard/WalletRequest.
[**CapturePost**](PaymentApi.md#CapturePost) | **Post** /Capture | Confirmación de un consumo previo.
[**ClosePost**](PaymentApi.md#ClosePost) | **Post** /Close | Utilizada por el POS para indicar que finalizo su sesion.
[**ConfigurePost**](PaymentApi.md#ConfigurePost) | **Post** /Configure | Permite crear configuración en Plataforma
[**ConfirmPost**](PaymentApi.md#ConfirmPost) | **Post** /Confirm | Confirmación de la última operación realizada.
[**DebtInquiryPost**](PaymentApi.md#DebtInquiryPost) | **Post** /DebtInquiry | Consulta de Deuda de cuenta/credencial
[**DebtPaymentPost**](PaymentApi.md#DebtPaymentPost) | **Post** /DebtPayment | Pago de Deuda, Resumen de Cuenta o Saldo.
[**DepositPost**](PaymentApi.md#DepositPost) | **Post** /Deposit | Confirmación de un consumo previo.
[**EnableServicePost**](PaymentApi.md#EnableServicePost) | **Post** /EnableService | Permite Habilitar el uso de un Servicio
[**EnrollmentPost**](PaymentApi.md#EnrollmentPost) | **Post** /Enrollment | Suscripción al servicio de pagos Tokenizados y pagos recurrentes.
[**GetBlockPost**](PaymentApi.md#GetBlockPost) | **Post** /GetBlock | Recupera los identificadores de las transacciones que  lo componen.
[**GetCardPost**](PaymentApi.md#GetCardPost) | **Post** /GetCard | Solicitud de Lectura del Medio de Pago
[**GetTransactionPost**](PaymentApi.md#GetTransactionPost) | **Post** /GetTransaction | Recupera los datos de la transacción especificada. 
[**KeepAlivePost**](PaymentApi.md#KeepAlivePost) | **Post** /KeepAlive | Mensaje que informa si está disponible el Servicio Authorize.
[**KeysRenewalPost**](PaymentApi.md#KeysRenewalPost) | **Post** /KeysRenewal | Renovacion de Llaves
[**OrderFinalPost**](PaymentApi.md#OrderFinalPost) | **Post** /OrderFinal | Reclamar el estatus de la operación de compra.
[**OrderGetPost**](PaymentApi.md#OrderGetPost) | **Post** /OrderGet | Recuperar la operación iniciada por el comercio, para su compra.
[**OrderInitialPost**](PaymentApi.md#OrderInitialPost) | **Post** /OrderInitial | Indica el inicio de una operación de venta.
[**OrderStatusPost**](PaymentApi.md#OrderStatusPost) | **Post** /OrderStatus | Recuperación del Estado de una Transacción Iniciada por el OrderInitial.
[**PaymentMethodPost**](PaymentApi.md#PaymentMethodPost) | **Post** /PaymentMethod | Consulta de  \&quot;planes\&quot; financieros para un Medio de Pago.
[**PaymentMethodsPost**](PaymentApi.md#PaymentMethodsPost) | **Post** /PaymentMethods | Consulta de los Medios de Pago  disponibles.
[**QueryCompaniesPost**](PaymentApi.md#QueryCompaniesPost) | **Post** /QueryCompanies | Consulta de Empresas para el Pago de Servicios o Deuda
[**QueryLineOfBusinessPost**](PaymentApi.md#QueryLineOfBusinessPost) | **Post** /QueryLineOfBusiness | Consulta de Rubros de Empresas para el Pago de Servicios o Deuda
[**ReturnPost**](PaymentApi.md#ReturnPost) | **Post** /Return | Realización de una devolución de operación de compra/autorización.
[**SalePost**](PaymentApi.md#SalePost) | **Post** /Sale | Realización de una compra/Autorización de compra
[**SettlementPost**](PaymentApi.md#SettlementPost) | **Post** /Settlement | Confirmación de un consumo previo.
[**ValidatePost**](PaymentApi.md#ValidatePost) | **Post** /Validate | Realización de una Validación
[**VoidDebtPaymentPost**](PaymentApi.md#VoidDebtPaymentPost) | **Post** /VoidDebtPayment | Cancelación de  Pago de Deuda, Saldo o Resumen.
[**VoidPost**](PaymentApi.md#VoidPost) | **Post** /Void | Operación de Cancelación/Anulación.
[**WalletRequestPost**](PaymentApi.md#WalletRequestPost) | **Post** /WalletRequest | Inicia un transacción contra el Wallet
[**WalletsPost**](PaymentApi.md#WalletsPost) | **Post** /Wallets | Obtiene la Lista de Wallets Disponibles



## AuthorizeSalePost

> AuthorizeSaleResponseObject AuthorizeSalePost(ctx).AuthorizeSaleObject(authorizeSaleObject).Execute()

Autorización de Venta sin Captura o Pre Autorización



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorizeSaleObject := *openapiclient.NewAuthorizeSaleObject() // AuthorizeSaleObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.AuthorizeSalePost(context.Background()).AuthorizeSaleObject(authorizeSaleObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.AuthorizeSalePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizeSalePost`: AuthorizeSaleResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.AuthorizeSalePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeSalePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizeSaleObject** | [**AuthorizeSaleObject**](AuthorizeSaleObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**AuthorizeSaleResponseObject**](AuthorizeSaleResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BalanceInquiryPost

> BalanceInquiryResponseObject BalanceInquiryPost(ctx).BalanceInquiryObject(balanceInquiryObject).Execute()

Consulta de Saldo/Deuda de cuenta/Credencial



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    balanceInquiryObject := *openapiclient.NewBalanceInquiryObject() // BalanceInquiryObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.BalanceInquiryPost(context.Background()).BalanceInquiryObject(balanceInquiryObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.BalanceInquiryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BalanceInquiryPost`: BalanceInquiryResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.BalanceInquiryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBalanceInquiryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **balanceInquiryObject** | [**BalanceInquiryObject**](BalanceInquiryObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**BalanceInquiryResponseObject**](BalanceInquiryResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockCancelPost

> BlockCancelResponseObject BlockCancelPost(ctx).BlockCancelObject(blockCancelObject).Execute()

Cancelación del último bloque de transacciones realizadas.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blockCancelObject := *openapiclient.NewBlockCancelObject() // BlockCancelObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.BlockCancelPost(context.Background()).BlockCancelObject(blockCancelObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.BlockCancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlockCancelPost`: BlockCancelResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.BlockCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockCancelObject** | [**BlockCancelObject**](BlockCancelObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**BlockCancelResponseObject**](BlockCancelResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockClosePost

> BlockCloseResponseObject BlockClosePost(ctx).BlockCloseObject(blockCloseObject).Execute()

Confirmación del último bloque de transacciones realizadas.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blockCloseObject := *openapiclient.NewBlockCloseObject() // BlockCloseObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.BlockClosePost(context.Background()).BlockCloseObject(blockCloseObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.BlockClosePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlockClosePost`: BlockCloseResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.BlockClosePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockClosePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockCloseObject** | [**BlockCloseObject**](BlockCloseObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**BlockCloseResponseObject**](BlockCloseResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockCreatePost

> BlockCreateResponseObject BlockCreatePost(ctx).BlockCreateObject(blockCreateObject).Execute()

Crea un Identificador de Bloque-Block de transacciones.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blockCreateObject := *openapiclient.NewBlockCreateObject() // BlockCreateObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.BlockCreatePost(context.Background()).BlockCreateObject(blockCreateObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.BlockCreatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlockCreatePost`: BlockCreateResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.BlockCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockCreateObject** | [**BlockCreateObject**](BlockCreateObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**BlockCreateResponseObject**](BlockCreateResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelPost

> CancelResponseObject CancelPost(ctx).CancelObject(cancelObject).Execute()

Cancela una Transacción en Curso, Inciada con un GetCard/WalletRequest.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cancelObject := *openapiclient.NewCancelObject() // CancelObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.CancelPost(context.Background()).CancelObject(cancelObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.CancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelPost`: CancelResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.CancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelObject** | [**CancelObject**](CancelObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**CancelResponseObject**](CancelResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapturePost

> CaptureResponseObject CapturePost(ctx).CaptureObject(captureObject).Execute()

Confirmación de un consumo previo.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    captureObject := *openapiclient.NewCaptureObject() // CaptureObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.CapturePost(context.Background()).CaptureObject(captureObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.CapturePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CapturePost`: CaptureResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.CapturePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCapturePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captureObject** | [**CaptureObject**](CaptureObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**CaptureResponseObject**](CaptureResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClosePost

> CloseResponseObject ClosePost(ctx).CloseObject(closeObject).Execute()

Utilizada por el POS para indicar que finalizo su sesion.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    closeObject := *openapiclient.NewCloseObject() // CloseObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.ClosePost(context.Background()).CloseObject(closeObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.ClosePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClosePost`: CloseResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.ClosePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClosePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **closeObject** | [**CloseObject**](CloseObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**CloseResponseObject**](CloseResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurePost

> ConfigureResponseObject ConfigurePost(ctx).ConfigureObject(configureObject).Execute()

Permite crear configuración en Plataforma



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configureObject := *openapiclient.NewConfigureObject() // ConfigureObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.ConfigurePost(context.Background()).ConfigureObject(configureObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.ConfigurePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigurePost`: ConfigureResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.ConfigurePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigurePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureObject** | [**ConfigureObject**](ConfigureObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**ConfigureResponseObject**](ConfigureResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmPost

> ConfirmResponseObject ConfirmPost(ctx).ConfirmObject(confirmObject).Execute()

Confirmación de la última operación realizada.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    confirmObject := *openapiclient.NewConfirmObject() // ConfirmObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.ConfirmPost(context.Background()).ConfirmObject(confirmObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.ConfirmPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmPost`: ConfirmResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.ConfirmPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confirmObject** | [**ConfirmObject**](ConfirmObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**ConfirmResponseObject**](ConfirmResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DebtInquiryPost

> DebtInquiryResponseObject DebtInquiryPost(ctx).DebtInquiryObject(debtInquiryObject).Execute()

Consulta de Deuda de cuenta/credencial



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    debtInquiryObject := *openapiclient.NewDebtInquiryObject() // DebtInquiryObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.DebtInquiryPost(context.Background()).DebtInquiryObject(debtInquiryObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.DebtInquiryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DebtInquiryPost`: DebtInquiryResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.DebtInquiryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDebtInquiryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debtInquiryObject** | [**DebtInquiryObject**](DebtInquiryObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**DebtInquiryResponseObject**](DebtInquiryResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DebtPaymentPost

> DebtPaymentResponseObject DebtPaymentPost(ctx).DebtPaymentObject(debtPaymentObject).Execute()

Pago de Deuda, Resumen de Cuenta o Saldo.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    debtPaymentObject := *openapiclient.NewDebtPaymentObject() // DebtPaymentObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.DebtPaymentPost(context.Background()).DebtPaymentObject(debtPaymentObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.DebtPaymentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DebtPaymentPost`: DebtPaymentResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.DebtPaymentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDebtPaymentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debtPaymentObject** | [**DebtPaymentObject**](DebtPaymentObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**DebtPaymentResponseObject**](DebtPaymentResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositPost

> DepositResponseObject DepositPost(ctx).DepositObject(depositObject).Execute()

Confirmación de un consumo previo.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    depositObject := *openapiclient.NewDepositObject() // DepositObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.DepositPost(context.Background()).DepositObject(depositObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.DepositPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DepositPost`: DepositResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.DepositPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepositPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositObject** | [**DepositObject**](DepositObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**DepositResponseObject**](DepositResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableServicePost

> EnableServiceResponseObject EnableServicePost(ctx).EnableServiceObject(enableServiceObject).Execute()

Permite Habilitar el uso de un Servicio



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enableServiceObject := *openapiclient.NewEnableServiceObject() // EnableServiceObject | Objeto que contendrá los datos del Requerimiento para Habilitar un Servicio.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.EnableServicePost(context.Background()).EnableServiceObject(enableServiceObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.EnableServicePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableServicePost`: EnableServiceResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.EnableServicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableServicePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableServiceObject** | [**EnableServiceObject**](EnableServiceObject.md) | Objeto que contendrá los datos del Requerimiento para Habilitar un Servicio. | 

### Return type

[**EnableServiceResponseObject**](EnableServiceResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollmentPost

> EnrollmentResponseObject EnrollmentPost(ctx).EnrollmentObject(enrollmentObject).Execute()

Suscripción al servicio de pagos Tokenizados y pagos recurrentes.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enrollmentObject := *openapiclient.NewEnrollmentObject() // EnrollmentObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.EnrollmentPost(context.Background()).EnrollmentObject(enrollmentObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.EnrollmentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollmentPost`: EnrollmentResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.EnrollmentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentObject** | [**EnrollmentObject**](EnrollmentObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**EnrollmentResponseObject**](EnrollmentResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockPost

> GetBlockResponseObject GetBlockPost(ctx).GetBlockObject(getBlockObject).Execute()

Recupera los identificadores de las transacciones que  lo componen.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getBlockObject := *openapiclient.NewGetBlockObject() // GetBlockObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.GetBlockPost(context.Background()).GetBlockObject(getBlockObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetBlockPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockPost`: GetBlockResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.GetBlockPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getBlockObject** | [**GetBlockObject**](GetBlockObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**GetBlockResponseObject**](GetBlockResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardPost

> GetCardResponseObject GetCardPost(ctx).GetCardObject(getCardObject).Execute()

Solicitud de Lectura del Medio de Pago



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getCardObject := *openapiclient.NewGetCardObject() // GetCardObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.GetCardPost(context.Background()).GetCardObject(getCardObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetCardPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardPost`: GetCardResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.GetCardPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCardPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCardObject** | [**GetCardObject**](GetCardObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**GetCardResponseObject**](GetCardResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionPost

> GetTransactionResponseObject GetTransactionPost(ctx).GetTransactionObject(getTransactionObject).Execute()

Recupera los datos de la transacción especificada. 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getTransactionObject := *openapiclient.NewGetTransactionObject() // GetTransactionObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.GetTransactionPost(context.Background()).GetTransactionObject(getTransactionObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetTransactionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionPost`: GetTransactionResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.GetTransactionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getTransactionObject** | [**GetTransactionObject**](GetTransactionObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**GetTransactionResponseObject**](GetTransactionResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeepAlivePost

> KeepAliveResponseObject KeepAlivePost(ctx).KeepAliveObject(keepAliveObject).Execute()

Mensaje que informa si está disponible el Servicio Authorize.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keepAliveObject := *openapiclient.NewKeepAliveObject() // KeepAliveObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.KeepAlivePost(context.Background()).KeepAliveObject(keepAliveObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.KeepAlivePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeepAlivePost`: KeepAliveResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.KeepAlivePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeepAlivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keepAliveObject** | [**KeepAliveObject**](KeepAliveObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**KeepAliveResponseObject**](KeepAliveResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysRenewalPost

> KeysRenewalObject KeysRenewalPost(ctx).KeysRenewalObject(keysRenewalObject).Execute()

Renovacion de Llaves



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keysRenewalObject := *openapiclient.NewKeysRenewalObject() // KeysRenewalObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.KeysRenewalPost(context.Background()).KeysRenewalObject(keysRenewalObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.KeysRenewalPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeysRenewalPost`: KeysRenewalObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.KeysRenewalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeysRenewalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keysRenewalObject** | [**KeysRenewalObject**](KeysRenewalObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**KeysRenewalObject**](KeysRenewalObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderFinalPost

> OrderFinalResponseObject OrderFinalPost(ctx).OrderFinalObject(orderFinalObject).Execute()

Reclamar el estatus de la operación de compra.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderFinalObject := *openapiclient.NewOrderFinalObject() // OrderFinalObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.OrderFinalPost(context.Background()).OrderFinalObject(orderFinalObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.OrderFinalPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderFinalPost`: OrderFinalResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.OrderFinalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderFinalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderFinalObject** | [**OrderFinalObject**](OrderFinalObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**OrderFinalResponseObject**](OrderFinalResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetPost

> OrderGetResponseObject OrderGetPost(ctx).OrderGetObject(orderGetObject).Execute()

Recuperar la operación iniciada por el comercio, para su compra.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderGetObject := *openapiclient.NewOrderGetObject() // OrderGetObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.OrderGetPost(context.Background()).OrderGetObject(orderGetObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.OrderGetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderGetPost`: OrderGetResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.OrderGetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderGetObject** | [**OrderGetObject**](OrderGetObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**OrderGetResponseObject**](OrderGetResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderInitialPost

> OrderInitialResponseObject OrderInitialPost(ctx).OrderInitialObject(orderInitialObject).Execute()

Indica el inicio de una operación de venta.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderInitialObject := *openapiclient.NewOrderInitialObject() // OrderInitialObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.OrderInitialPost(context.Background()).OrderInitialObject(orderInitialObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.OrderInitialPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderInitialPost`: OrderInitialResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.OrderInitialPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderInitialPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderInitialObject** | [**OrderInitialObject**](OrderInitialObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**OrderInitialResponseObject**](OrderInitialResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderStatusPost

> OrderStatusResponseObject OrderStatusPost(ctx).OrderStatusObject(orderStatusObject).Execute()

Recuperación del Estado de una Transacción Iniciada por el OrderInitial.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderStatusObject := *openapiclient.NewOrderStatusObject() // OrderStatusObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.OrderStatusPost(context.Background()).OrderStatusObject(orderStatusObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.OrderStatusPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderStatusPost`: OrderStatusResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.OrderStatusPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderStatusObject** | [**OrderStatusObject**](OrderStatusObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**OrderStatusResponseObject**](OrderStatusResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentMethodPost

> PaymentMethodResponseObject PaymentMethodPost(ctx).PaymentMethodObject(paymentMethodObject).Execute()

Consulta de  \"planes\" financieros para un Medio de Pago.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentMethodObject := *openapiclient.NewPaymentMethodObject() // PaymentMethodObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.PaymentMethodPost(context.Background()).PaymentMethodObject(paymentMethodObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.PaymentMethodPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentMethodPost`: PaymentMethodResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.PaymentMethodPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentMethodPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentMethodObject** | [**PaymentMethodObject**](PaymentMethodObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**PaymentMethodResponseObject**](PaymentMethodResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentMethodsPost

> PaymentMethodsResponseObject PaymentMethodsPost(ctx).PaymentMethodsObject(paymentMethodsObject).Execute()

Consulta de los Medios de Pago  disponibles.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentMethodsObject := *openapiclient.NewPaymentMethodsObject() // PaymentMethodsObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.PaymentMethodsPost(context.Background()).PaymentMethodsObject(paymentMethodsObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.PaymentMethodsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentMethodsPost`: PaymentMethodsResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.PaymentMethodsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentMethodsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentMethodsObject** | [**PaymentMethodsObject**](PaymentMethodsObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**PaymentMethodsResponseObject**](PaymentMethodsResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryCompaniesPost

> QueryCompaniesResponseObject QueryCompaniesPost(ctx).QueryCompaniesObject(queryCompaniesObject).Execute()

Consulta de Empresas para el Pago de Servicios o Deuda



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    queryCompaniesObject := *openapiclient.NewQueryCompaniesObject() // QueryCompaniesObject | Objeto que contendrá los datos del Requerimiento para obtener la lista de Empresas.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.QueryCompaniesPost(context.Background()).QueryCompaniesObject(queryCompaniesObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.QueryCompaniesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryCompaniesPost`: QueryCompaniesResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.QueryCompaniesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryCompaniesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryCompaniesObject** | [**QueryCompaniesObject**](QueryCompaniesObject.md) | Objeto que contendrá los datos del Requerimiento para obtener la lista de Empresas. | 

### Return type

[**QueryCompaniesResponseObject**](QueryCompaniesResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryLineOfBusinessPost

> QueryLineOfBusinessResponseObject QueryLineOfBusinessPost(ctx).QueryLineOfBusinessObject(queryLineOfBusinessObject).Execute()

Consulta de Rubros de Empresas para el Pago de Servicios o Deuda



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    queryLineOfBusinessObject := *openapiclient.NewQueryLineOfBusinessObject() // QueryLineOfBusinessObject | Objeto que contendrá los datos del Requerimiento para obtener la lista de Rubros.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.QueryLineOfBusinessPost(context.Background()).QueryLineOfBusinessObject(queryLineOfBusinessObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.QueryLineOfBusinessPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryLineOfBusinessPost`: QueryLineOfBusinessResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.QueryLineOfBusinessPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryLineOfBusinessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryLineOfBusinessObject** | [**QueryLineOfBusinessObject**](QueryLineOfBusinessObject.md) | Objeto que contendrá los datos del Requerimiento para obtener la lista de Rubros. | 

### Return type

[**QueryLineOfBusinessResponseObject**](QueryLineOfBusinessResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnPost

> ReturnResponseObject ReturnPost(ctx).ReturnObject(returnObject).Execute()

Realización de una devolución de operación de compra/autorización.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    returnObject := *openapiclient.NewReturnObject() // ReturnObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.ReturnPost(context.Background()).ReturnObject(returnObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.ReturnPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnPost`: ReturnResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.ReturnPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnObject** | [**ReturnObject**](ReturnObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**ReturnResponseObject**](ReturnResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalePost

> SaleResponseObject SalePost(ctx).SaleObject(saleObject).Execute()

Realización de una compra/Autorización de compra



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    saleObject := *openapiclient.NewSaleObject() // SaleObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.SalePost(context.Background()).SaleObject(saleObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.SalePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SalePost`: SaleResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.SalePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSalePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saleObject** | [**SaleObject**](SaleObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**SaleResponseObject**](SaleResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettlementPost

> SettlementResponseObject SettlementPost(ctx).SettlementObject(settlementObject).Execute()

Confirmación de un consumo previo.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    settlementObject := *openapiclient.NewSettlementObject() // SettlementObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.SettlementPost(context.Background()).SettlementObject(settlementObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.SettlementPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettlementPost`: SettlementResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.SettlementPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettlementPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settlementObject** | [**SettlementObject**](SettlementObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**SettlementResponseObject**](SettlementResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePost

> ValidateResponseObject ValidatePost(ctx).ValidateObject(validateObject).Execute()

Realización de una Validación



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    validateObject := *openapiclient.NewValidateObject() // ValidateObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.ValidatePost(context.Background()).ValidateObject(validateObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.ValidatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidatePost`: ValidateResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.ValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateObject** | [**ValidateObject**](ValidateObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**ValidateResponseObject**](ValidateResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidDebtPaymentPost

> VoidDebtPaymentResponseObject VoidDebtPaymentPost(ctx).VoidDebtPaymentObject(voidDebtPaymentObject).Execute()

Cancelación de  Pago de Deuda, Saldo o Resumen.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    voidDebtPaymentObject := *openapiclient.NewVoidDebtPaymentObject() // VoidDebtPaymentObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.VoidDebtPaymentPost(context.Background()).VoidDebtPaymentObject(voidDebtPaymentObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.VoidDebtPaymentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VoidDebtPaymentPost`: VoidDebtPaymentResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.VoidDebtPaymentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVoidDebtPaymentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voidDebtPaymentObject** | [**VoidDebtPaymentObject**](VoidDebtPaymentObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**VoidDebtPaymentResponseObject**](VoidDebtPaymentResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidPost

> VoidResponseObject VoidPost(ctx).VoidObject(voidObject).Execute()

Operación de Cancelación/Anulación.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    voidObject := *openapiclient.NewVoidObject() // VoidObject | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.VoidPost(context.Background()).VoidObject(voidObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.VoidPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VoidPost`: VoidResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.VoidPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVoidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voidObject** | [**VoidObject**](VoidObject.md) | Elementos o Atributos que componen el requerimiento de una transacción, los atributos condicionales y/o opcionales que no sean requeridos para esta transacción no deberán ser enviados. | 

### Return type

[**VoidResponseObject**](VoidResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletRequestPost

> WalletRequestResponseObject WalletRequestPost(ctx).WalletRequestObject(walletRequestObject).Execute()

Inicia un transacción contra el Wallet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    walletRequestObject := *openapiclient.NewWalletRequestObject() // WalletRequestObject | Objeto que contendrá los datos del Requerimiento para obtener el Requerimeinto del Wallet.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.WalletRequestPost(context.Background()).WalletRequestObject(walletRequestObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.WalletRequestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WalletRequestPost`: WalletRequestResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.WalletRequestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletRequestObject** | [**WalletRequestObject**](WalletRequestObject.md) | Objeto que contendrá los datos del Requerimiento para obtener el Requerimeinto del Wallet. | 

### Return type

[**WalletRequestResponseObject**](WalletRequestResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsPost

> WalletsResponseObject WalletsPost(ctx).WalletsObject(walletsObject).Execute()

Obtiene la Lista de Wallets Disponibles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    walletsObject := *openapiclient.NewWalletsObject() // WalletsObject | Objeto que contendrá los datos del Requerimiento para obtener los Wallets disponibles.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.WalletsPost(context.Background()).WalletsObject(walletsObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.WalletsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WalletsPost`: WalletsResponseObject
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.WalletsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletsObject** | [**WalletsObject**](WalletsObject.md) | Objeto que contendrá los datos del Requerimiento para obtener los Wallets disponibles. | 

### Return type

[**WalletsResponseObject**](WalletsResponseObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

