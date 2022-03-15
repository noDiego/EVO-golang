# OrderGetResponseObjectOrderGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | Pointer to **string** | Tipo de Operación que se requirió, contendrá el mismo valor que se recibió en el requerimiento, sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**ResponseActions** | **[]string** | Acciones a realizar por parte del POS y/o PINPAD en base al resultado de la operación que ha sido procesada. Cada uno de estos actions o acciones están concatenadas por comas. Los posibles actions son OK, Approve, Refuse, IssuerCall, Tickets, WithHold, GetCard, UseTerminalToAuthorize, ConfigurationError, SystemError, ResourceError, ProcessError, Completed. | 
**ResponseMessage** | **string** | Descripción del resultado del proceso del requerimiento recibido. Esta descripción es generada por la plataforma, no por el Host que termine resolviendo la transacción. | 
**CompanyIdentification** | Pointer to **string** | ID que identifica la companía desde donde proviene la petición. | [optional] 
**SystemIdentification** | Pointer to **string** | ID que identifica el sistema desde donde proviene la petición. | [optional] 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**ForeignResponseCode** | Pointer to **string** | Código de respuesta para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF. | [optional] 
**ResponseCode** | **int32** | Código de Respuesta Interno de la plataforma, el POS debe actuar por lo que indican las acciones especificadas en ResponseActions y no por el código de respuesta informado en este campo o elemento, pero es una buena práctica que sea persistido por el mismo. | 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**AdditionalInformation** | Pointer to [**[]SaleResponseObjectSaleResponseAdditionalInformation**](SaleResponseObjectSaleResponseAdditionalInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**RequiredInformation** | Pointer to [**[]DebtPaymentObjectDebtPaymentRequiredInformation**](DebtPaymentObjectDebtPaymentRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**InitialIdentification** | Pointer to **string** | ID privado del contexto de pago generado. Este elemento es equivalente al llamado RequestKey. | [optional] 
**InitialToken** | Pointer to **string** | ID publico del contexto de pago generado, mediante el cual se lo podra referenciar desde el formulario de pago | [optional] 
**CustomerRedirectAddress** | Pointer to **string** | Elementos a adiciónar a la URL a la cual se redigira para mostrar el formulario de pago | [optional] 
**DateTime** | Pointer to **time.Time** | Fecha y Hora de la transacción generada por el Punto de Venta - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**TransactionIdentification** | Pointer to **string** | ID de la operación a realizar, generado por el sistema externo | [optional] 
**TransactionDescription** | Pointer to **string** | Descripción del tipo de operación que se realizará | [optional] 
**MerchantIdentification** | Pointer to **string** | ID del comercio que el para el cual se estara realizando la operación. Este valor puede ser enviado en lugar del SystemIdentification y CompanyIdentification, para luego ser traducido por el propio Plataforma a los valores configurados para ello. Ademas, puede relaciónarse el valor de BranchIdentification y/o POSIdentification de la misma forma. Si ello no se realiza, estos elementos tomaran valores genericos (por default) igual a 0. | [optional] 
**Installments** | Pointer to **int32** | Cantidad de cuotas que permite este plan | [optional] 
**FacilityType** | Pointer to **int32** | Tipo de plan utilizado para para realizar la operación | [optional] 
**CurrencyCode** | Pointer to **string** | código de Moneda - ISO 4217 &lt;https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica &lt;br /&gt;   * Num   - Alpha - Description &lt;br /&gt;   * &#39;032&#39; - &#39;ARS&#39; - Pesos Argentinos &lt;br /&gt;   * &#39;152&#39; - &#39;CLP&#39; - Pesos Chilenos &lt;br/&gt;   * &#39;484&#39; - &#39;MXN&#39; - Pesos Mexicanos &lt;br/&gt;   * &#39;840&#39; - &#39;USD&#39; - dólares Americanos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;858&#39; - &#39;UYU&#39; - Pesos Uruguayos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;986&#39; - &#39;BRL&#39; - Real Brasileño | [optional] 
**Amount** | Pointer to **float32** | Monto con la que se realizó transacción. Si este valor es recibido, la búsqueda de los planes será limitada con este criterio. | [optional] 
**NetAmount** | Pointer to **string** | . | [optional] 
**Payer** | Pointer to [**SaleObjectSalePayer**](SaleObjectSalePayer.md) |  | [optional] 
**Customer** | Pointer to [**SaleObjectSaleCustomer**](SaleObjectSaleCustomer.md) |  | [optional] 
**Seller** | Pointer to [**SaleObjectSaleSeller**](SaleObjectSaleSeller.md) |  | [optional] 
**TaxRefundType** | Pointer to **string** | . | [optional] 
**MerchantRedirectURL** | Pointer to **string** | . | [optional] 
**TransactionValidThru** | Pointer to **time.Time** | Fecha y Hora de fin de validez de La transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**Products** | Pointer to [**[]SaleResponseObjectSaleResponseProducts**](SaleResponseObjectSaleResponseProducts.md) | Detalle de Productos de la Operación. | [optional] 
**Configuration** | Pointer to [**SaleResponseObjectSaleResponseConfiguration**](SaleResponseObjectSaleResponseConfiguration.md) |  | [optional] 

## Methods

### NewOrderGetResponseObjectOrderGetResponse

`func NewOrderGetResponseObjectOrderGetResponse(responseActions []string, responseMessage string, responseCode int32, ) *OrderGetResponseObjectOrderGetResponse`

NewOrderGetResponseObjectOrderGetResponse instantiates a new OrderGetResponseObjectOrderGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderGetResponseObjectOrderGetResponseWithDefaults

`func NewOrderGetResponseObjectOrderGetResponseWithDefaults() *OrderGetResponseObjectOrderGetResponse`

NewOrderGetResponseObjectOrderGetResponseWithDefaults instantiates a new OrderGetResponseObjectOrderGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *OrderGetResponseObjectOrderGetResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *OrderGetResponseObjectOrderGetResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *OrderGetResponseObjectOrderGetResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetResponseActions

`func (o *OrderGetResponseObjectOrderGetResponse) GetResponseActions() []string`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetResponseActionsOk() (*[]string, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *OrderGetResponseObjectOrderGetResponse) SetResponseActions(v []string)`

SetResponseActions sets ResponseActions field to given value.


### GetResponseMessage

`func (o *OrderGetResponseObjectOrderGetResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *OrderGetResponseObjectOrderGetResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetCompanyIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetSystemIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetForeignResponseCode

`func (o *OrderGetResponseObjectOrderGetResponse) GetForeignResponseCode() string`

GetForeignResponseCode returns the ForeignResponseCode field if non-nil, zero value otherwise.

### GetForeignResponseCodeOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetForeignResponseCodeOk() (*string, bool)`

GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignResponseCode

`func (o *OrderGetResponseObjectOrderGetResponse) SetForeignResponseCode(v string)`

SetForeignResponseCode sets ForeignResponseCode field to given value.

### HasForeignResponseCode

`func (o *OrderGetResponseObjectOrderGetResponse) HasForeignResponseCode() bool`

HasForeignResponseCode returns a boolean if a field has been set.

### GetResponseCode

`func (o *OrderGetResponseObjectOrderGetResponse) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *OrderGetResponseObjectOrderGetResponse) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetServiceVersion

`func (o *OrderGetResponseObjectOrderGetResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *OrderGetResponseObjectOrderGetResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *OrderGetResponseObjectOrderGetResponse) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *OrderGetResponseObjectOrderGetResponse) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *OrderGetResponseObjectOrderGetResponse) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *OrderGetResponseObjectOrderGetResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *OrderGetResponseObjectOrderGetResponse) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OrderGetResponseObjectOrderGetResponse) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OrderGetResponseObjectOrderGetResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *OrderGetResponseObjectOrderGetResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *OrderGetResponseObjectOrderGetResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *OrderGetResponseObjectOrderGetResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *OrderGetResponseObjectOrderGetResponse) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetRequiredInformationOk() (*[]DebtPaymentObjectDebtPaymentRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *OrderGetResponseObjectOrderGetResponse) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *OrderGetResponseObjectOrderGetResponse) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetInitialIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) GetInitialIdentification() string`

GetInitialIdentification returns the InitialIdentification field if non-nil, zero value otherwise.

### GetInitialIdentificationOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetInitialIdentificationOk() (*string, bool)`

GetInitialIdentificationOk returns a tuple with the InitialIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) SetInitialIdentification(v string)`

SetInitialIdentification sets InitialIdentification field to given value.

### HasInitialIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) HasInitialIdentification() bool`

HasInitialIdentification returns a boolean if a field has been set.

### GetInitialToken

`func (o *OrderGetResponseObjectOrderGetResponse) GetInitialToken() string`

GetInitialToken returns the InitialToken field if non-nil, zero value otherwise.

### GetInitialTokenOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetInitialTokenOk() (*string, bool)`

GetInitialTokenOk returns a tuple with the InitialToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialToken

`func (o *OrderGetResponseObjectOrderGetResponse) SetInitialToken(v string)`

SetInitialToken sets InitialToken field to given value.

### HasInitialToken

`func (o *OrderGetResponseObjectOrderGetResponse) HasInitialToken() bool`

HasInitialToken returns a boolean if a field has been set.

### GetCustomerRedirectAddress

`func (o *OrderGetResponseObjectOrderGetResponse) GetCustomerRedirectAddress() string`

GetCustomerRedirectAddress returns the CustomerRedirectAddress field if non-nil, zero value otherwise.

### GetCustomerRedirectAddressOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetCustomerRedirectAddressOk() (*string, bool)`

GetCustomerRedirectAddressOk returns a tuple with the CustomerRedirectAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRedirectAddress

`func (o *OrderGetResponseObjectOrderGetResponse) SetCustomerRedirectAddress(v string)`

SetCustomerRedirectAddress sets CustomerRedirectAddress field to given value.

### HasCustomerRedirectAddress

`func (o *OrderGetResponseObjectOrderGetResponse) HasCustomerRedirectAddress() bool`

HasCustomerRedirectAddress returns a boolean if a field has been set.

### GetDateTime

`func (o *OrderGetResponseObjectOrderGetResponse) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *OrderGetResponseObjectOrderGetResponse) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *OrderGetResponseObjectOrderGetResponse) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetTransactionIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionIdentification() string`

GetTransactionIdentification returns the TransactionIdentification field if non-nil, zero value otherwise.

### GetTransactionIdentificationOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionIdentificationOk() (*string, bool)`

GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) SetTransactionIdentification(v string)`

SetTransactionIdentification sets TransactionIdentification field to given value.

### HasTransactionIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) HasTransactionIdentification() bool`

HasTransactionIdentification returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *OrderGetResponseObjectOrderGetResponse) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *OrderGetResponseObjectOrderGetResponse) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetMerchantIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) GetMerchantIdentification() string`

GetMerchantIdentification returns the MerchantIdentification field if non-nil, zero value otherwise.

### GetMerchantIdentificationOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetMerchantIdentificationOk() (*string, bool)`

GetMerchantIdentificationOk returns a tuple with the MerchantIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) SetMerchantIdentification(v string)`

SetMerchantIdentification sets MerchantIdentification field to given value.

### HasMerchantIdentification

`func (o *OrderGetResponseObjectOrderGetResponse) HasMerchantIdentification() bool`

HasMerchantIdentification returns a boolean if a field has been set.

### GetInstallments

`func (o *OrderGetResponseObjectOrderGetResponse) GetInstallments() int32`

GetInstallments returns the Installments field if non-nil, zero value otherwise.

### GetInstallmentsOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetInstallmentsOk() (*int32, bool)`

GetInstallmentsOk returns a tuple with the Installments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallments

`func (o *OrderGetResponseObjectOrderGetResponse) SetInstallments(v int32)`

SetInstallments sets Installments field to given value.

### HasInstallments

`func (o *OrderGetResponseObjectOrderGetResponse) HasInstallments() bool`

HasInstallments returns a boolean if a field has been set.

### GetFacilityType

`func (o *OrderGetResponseObjectOrderGetResponse) GetFacilityType() int32`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetFacilityTypeOk() (*int32, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *OrderGetResponseObjectOrderGetResponse) SetFacilityType(v int32)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *OrderGetResponseObjectOrderGetResponse) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrderGetResponseObjectOrderGetResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderGetResponseObjectOrderGetResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrderGetResponseObjectOrderGetResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmount

`func (o *OrderGetResponseObjectOrderGetResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderGetResponseObjectOrderGetResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OrderGetResponseObjectOrderGetResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetNetAmount

`func (o *OrderGetResponseObjectOrderGetResponse) GetNetAmount() string`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetNetAmountOk() (*string, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *OrderGetResponseObjectOrderGetResponse) SetNetAmount(v string)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *OrderGetResponseObjectOrderGetResponse) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetPayer

`func (o *OrderGetResponseObjectOrderGetResponse) GetPayer() SaleObjectSalePayer`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetPayerOk() (*SaleObjectSalePayer, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *OrderGetResponseObjectOrderGetResponse) SetPayer(v SaleObjectSalePayer)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *OrderGetResponseObjectOrderGetResponse) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderGetResponseObjectOrderGetResponse) GetCustomer() SaleObjectSaleCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetCustomerOk() (*SaleObjectSaleCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderGetResponseObjectOrderGetResponse) SetCustomer(v SaleObjectSaleCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderGetResponseObjectOrderGetResponse) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSeller

`func (o *OrderGetResponseObjectOrderGetResponse) GetSeller() SaleObjectSaleSeller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetSellerOk() (*SaleObjectSaleSeller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *OrderGetResponseObjectOrderGetResponse) SetSeller(v SaleObjectSaleSeller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *OrderGetResponseObjectOrderGetResponse) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetTaxRefundType

`func (o *OrderGetResponseObjectOrderGetResponse) GetTaxRefundType() string`

GetTaxRefundType returns the TaxRefundType field if non-nil, zero value otherwise.

### GetTaxRefundTypeOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetTaxRefundTypeOk() (*string, bool)`

GetTaxRefundTypeOk returns a tuple with the TaxRefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRefundType

`func (o *OrderGetResponseObjectOrderGetResponse) SetTaxRefundType(v string)`

SetTaxRefundType sets TaxRefundType field to given value.

### HasTaxRefundType

`func (o *OrderGetResponseObjectOrderGetResponse) HasTaxRefundType() bool`

HasTaxRefundType returns a boolean if a field has been set.

### GetMerchantRedirectURL

`func (o *OrderGetResponseObjectOrderGetResponse) GetMerchantRedirectURL() string`

GetMerchantRedirectURL returns the MerchantRedirectURL field if non-nil, zero value otherwise.

### GetMerchantRedirectURLOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetMerchantRedirectURLOk() (*string, bool)`

GetMerchantRedirectURLOk returns a tuple with the MerchantRedirectURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRedirectURL

`func (o *OrderGetResponseObjectOrderGetResponse) SetMerchantRedirectURL(v string)`

SetMerchantRedirectURL sets MerchantRedirectURL field to given value.

### HasMerchantRedirectURL

`func (o *OrderGetResponseObjectOrderGetResponse) HasMerchantRedirectURL() bool`

HasMerchantRedirectURL returns a boolean if a field has been set.

### GetTransactionValidThru

`func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionValidThru() time.Time`

GetTransactionValidThru returns the TransactionValidThru field if non-nil, zero value otherwise.

### GetTransactionValidThruOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionValidThruOk() (*time.Time, bool)`

GetTransactionValidThruOk returns a tuple with the TransactionValidThru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionValidThru

`func (o *OrderGetResponseObjectOrderGetResponse) SetTransactionValidThru(v time.Time)`

SetTransactionValidThru sets TransactionValidThru field to given value.

### HasTransactionValidThru

`func (o *OrderGetResponseObjectOrderGetResponse) HasTransactionValidThru() bool`

HasTransactionValidThru returns a boolean if a field has been set.

### GetProducts

`func (o *OrderGetResponseObjectOrderGetResponse) GetProducts() []SaleResponseObjectSaleResponseProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetProductsOk() (*[]SaleResponseObjectSaleResponseProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *OrderGetResponseObjectOrderGetResponse) SetProducts(v []SaleResponseObjectSaleResponseProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *OrderGetResponseObjectOrderGetResponse) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetConfiguration

`func (o *OrderGetResponseObjectOrderGetResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *OrderGetResponseObjectOrderGetResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *OrderGetResponseObjectOrderGetResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *OrderGetResponseObjectOrderGetResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


