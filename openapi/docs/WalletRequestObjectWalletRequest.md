# WalletRequestObjectWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemIdentification** | **string** | Identificación  del sistema o Aplicativo. | 
**CompanyIdentification** | **string** | Identificador de la Compañía para la Plataforma de Integración. | 
**BranchIdentification** | Pointer to **string** | ID de la sucursal, en caso de no existir ese concepto y ser un dispotivo que tiene posiciónamiento, deberá ser el mismo en caso contrario el mismo valor que el elemento &lt;b&gt;CompanyIdentification&lt;/b&gt; | [optional] 
**POSIdentification** | Pointer to **string** | ID de la caja perteneciente a la sucursal indicada, en caso de no existir ese concepto y ser un dispotivo que tiene un identificador unico como por ejemplo el IMEI deberá viajar este. | [optional] 
**RequestType** | Pointer to **string** | Tipo de Requerimiento. Solo requerido para protocolos de transporte donde el tipo no se este especificando o en el PATH o en el tipo complejo que contiene al mismo | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**Block** | Pointer to **string** | ID que identifica a un grupo de transacciones que serán confirmadas o canceladas | [optional] 
**Ticket** | Pointer to **string** | Ticket Digitalizado ( Total o parte del mismo por ejemplo la Firma Digitalizada )    codificado en Base64. | [optional] 
**TicketAnswerKey** | Pointer to **string** | Identificador Unívoco de la Transacción que se quiere Referenciar a la cual pertenece el Ticket Digitalizado. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. Si firma fue capturada previamente y se envía en el requerimiento de la misma Operación Sale, Authorize*, Void, Return, Adjustment, DebtPayment o VoidDebtPayment no es necesario que se envíe este elemento o campo. | [optional] 
**RequiredInformation** | Pointer to [**[]DebtPaymentObjectDebtPaymentRequiredInformation**](DebtPaymentObjectDebtPaymentRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**AdditionalInformation** | Pointer to [**[]SaleResponseObjectSaleResponseAdditionalInformation**](SaleResponseObjectSaleResponseAdditionalInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**RequestKey** | Pointer to **string** | Identificador Privado para ser identificar la transacción a realizar, normalmente es un código a Presentar, solo sera enviado si ya fue creado por un WalletRequest previo y solamente queremos obtener los datos que pudo el Pagador Selecciónar, como el Medio de Pago para poder aplicar una promoción | [optional] 
**MerchantNotifyURL** | Pointer to **string** | URL para notificación del comercio | [optional] 
**Reference** | Pointer to **string** | Referencia de la transacción para el punto de venta | [optional] 
**TransactionType** | Pointer to **string** | Tipo de Transacción (Sale, Void, Return, Authorize,...) por la cual se está realizado el requerimiento (Usado en WalletRequest) | [optional] 
**TransactionDescription** | Pointer to **string** | Descripción del tipo de operación que se realizará | [optional] 
**TransactionIdentification** | Pointer to **string** | ID de La transacción UNIVOCO para el Punto de venta | [optional] 
**TrasactionDateTime** | Pointer to **string** | Fecha y Hora de la transacción generada por el Punto de Venta - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**TransactionTimeout** | Pointer to **float32** | Tiempo en segúndos que la transacción permanecera vigente | [optional] 
**POSType** | Pointer to **string** | Tipo de punto de venta. | [optional] 
**POSVersion** | Pointer to **string** | Versión del Aplicativo del punto de Venta. | [optional] 
**POSAddress** | Pointer to **string** | Dirección IP de la Caja o POS. | [optional] 
**POSSerial** | Pointer to **string** | Número de serie o identificador unívoco del punto de venta. | [optional] 
**POSGEO** | Pointer to **string** | Coordenadas Geográficas del aplicativo de punto de venta | [optional] 
**ReadingDeviceType** | Pointer to **string** | Tipo de dispositivo utilizado para ingresar los datos de la tarjeta. En función al dispositivo usado, serán realizadas ciertas verificaciones, por lo que ciertos datos serán requeridos. CustomerKeyboard, utilizado para ingreso manual de tarjeta a través de un portal web, por ejemplo - Keyboard, utilizado para ingreso manual de la tarjeta por parte del vendedor - MagneticStripReader, lector de banda de tarjetas por emulación de teclado, u otro valor configurado  que indentifique el dispositivo que se esta utilizando. | [optional] 
**ReadingDeviceOperatingFrom** | Pointer to **time.Time** | Indica desde cuando se encuentra operativo o encendido el dispositivo | [optional] 
**ReadingDeviceVersion** | Pointer to **string** | Versión del dispositivo. | [optional] 
**ReadingDeviceAddress** | Pointer to **string** | Dirección IP o MAC Address del dispositivo. | [optional] 
**ReadingDeviceSerial** | Pointer to **string** | Número de serie o identificador unívoco del dispositivo. | [optional] 
**ReadingDeviceGEO** | Pointer to **string** | Coordenadas Geograficas del dispositivo de lectura | [optional] 
**UserID** | Pointer to **string** | Identificador del usuario que está realizando la Transacción. | [optional] 
**UserName** | Pointer to **string** | Nombre del usuario que está realizando la Transacción. | [optional] 
**CurrencyCode** | Pointer to **string** | código de Moneda - ISO 4217 &lt;https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica &lt;br /&gt;   * Num   - Alpha - Description &lt;br /&gt;   * &#39;032&#39; - &#39;ARS&#39; - Pesos Argentinos &lt;br /&gt;   * &#39;152&#39; - &#39;CLP&#39; - Pesos Chilenos &lt;br/&gt;   * &#39;484&#39; - &#39;MXN&#39; - Pesos Mexicanos &lt;br/&gt;   * &#39;840&#39; - &#39;USD&#39; - dólares Americanos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;858&#39; - &#39;UYU&#39; - Pesos Uruguayos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;986&#39; - &#39;BRL&#39; - Real Brasileño | [optional] 
**Amount** | Pointer to **float32** | Monto de la transacción a ser enviada | [optional] 
**CashbackAmount** | Pointer to **float32** | Monto del dinero en efectivo (cashback). | [optional] 
**TipAmount** | Pointer to **float32** | Importe o Monto de la Propina. | [optional] 
**WalletIdentification** | Pointer to **string** | Identificador del Wallet Retornado por la Operación Wallets | [optional] 
**Payer** | Pointer to [**SaleObjectSalePayer**](SaleObjectSalePayer.md) |  | [optional] 
**Customer** | Pointer to [**SaleObjectSaleCustomer**](SaleObjectSaleCustomer.md) |  | [optional] 
**Seller** | Pointer to [**SaleObjectSaleSeller**](SaleObjectSaleSeller.md) |  | [optional] 
**Products** | Pointer to [**[]SaleResponseObjectSaleResponseProducts**](SaleResponseObjectSaleResponseProducts.md) | Detalle de Productos de la Operación. | [optional] 
**TaxRefundType** | Pointer to **string** | Esquema de Devolución de Impuestos a utilizar en la transacción | [optional] 
**ValidThru** | Pointer to **time.Time** | Fecha y Hora de fin de validez de La transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relacionado a cada uno de los planes disponibles. | [optional] 
**TerminalID** | Pointer to **string** | Identificador de Terminal por el cual se envía la Transacción al Host. | [optional] 
**TerminalTrace** | Pointer to **int32** | Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID. | [optional] 
**SettlementBatchNumber** | Pointer to **int32** | Para aquellos host que exista el concepto de lote, es el número de lote al cual pertenece la transacción. | [optional] 

## Methods

### NewWalletRequestObjectWalletRequest

`func NewWalletRequestObjectWalletRequest(systemIdentification string, companyIdentification string, ) *WalletRequestObjectWalletRequest`

NewWalletRequestObjectWalletRequest instantiates a new WalletRequestObjectWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletRequestObjectWalletRequestWithDefaults

`func NewWalletRequestObjectWalletRequestWithDefaults() *WalletRequestObjectWalletRequest`

NewWalletRequestObjectWalletRequestWithDefaults instantiates a new WalletRequestObjectWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemIdentification

`func (o *WalletRequestObjectWalletRequest) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *WalletRequestObjectWalletRequest) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *WalletRequestObjectWalletRequest) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetCompanyIdentification

`func (o *WalletRequestObjectWalletRequest) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *WalletRequestObjectWalletRequest) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *WalletRequestObjectWalletRequest) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetBranchIdentification

`func (o *WalletRequestObjectWalletRequest) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *WalletRequestObjectWalletRequest) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *WalletRequestObjectWalletRequest) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *WalletRequestObjectWalletRequest) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *WalletRequestObjectWalletRequest) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *WalletRequestObjectWalletRequest) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *WalletRequestObjectWalletRequest) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *WalletRequestObjectWalletRequest) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetRequestType

`func (o *WalletRequestObjectWalletRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *WalletRequestObjectWalletRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *WalletRequestObjectWalletRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *WalletRequestObjectWalletRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *WalletRequestObjectWalletRequest) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *WalletRequestObjectWalletRequest) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *WalletRequestObjectWalletRequest) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *WalletRequestObjectWalletRequest) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *WalletRequestObjectWalletRequest) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *WalletRequestObjectWalletRequest) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *WalletRequestObjectWalletRequest) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *WalletRequestObjectWalletRequest) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *WalletRequestObjectWalletRequest) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *WalletRequestObjectWalletRequest) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *WalletRequestObjectWalletRequest) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *WalletRequestObjectWalletRequest) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *WalletRequestObjectWalletRequest) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *WalletRequestObjectWalletRequest) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *WalletRequestObjectWalletRequest) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *WalletRequestObjectWalletRequest) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetTicket

`func (o *WalletRequestObjectWalletRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *WalletRequestObjectWalletRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *WalletRequestObjectWalletRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *WalletRequestObjectWalletRequest) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTicketAnswerKey

`func (o *WalletRequestObjectWalletRequest) GetTicketAnswerKey() string`

GetTicketAnswerKey returns the TicketAnswerKey field if non-nil, zero value otherwise.

### GetTicketAnswerKeyOk

`func (o *WalletRequestObjectWalletRequest) GetTicketAnswerKeyOk() (*string, bool)`

GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketAnswerKey

`func (o *WalletRequestObjectWalletRequest) SetTicketAnswerKey(v string)`

SetTicketAnswerKey sets TicketAnswerKey field to given value.

### HasTicketAnswerKey

`func (o *WalletRequestObjectWalletRequest) HasTicketAnswerKey() bool`

HasTicketAnswerKey returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *WalletRequestObjectWalletRequest) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *WalletRequestObjectWalletRequest) GetRequiredInformationOk() (*[]DebtPaymentObjectDebtPaymentRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *WalletRequestObjectWalletRequest) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *WalletRequestObjectWalletRequest) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *WalletRequestObjectWalletRequest) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *WalletRequestObjectWalletRequest) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *WalletRequestObjectWalletRequest) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *WalletRequestObjectWalletRequest) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRequestKey

`func (o *WalletRequestObjectWalletRequest) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *WalletRequestObjectWalletRequest) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *WalletRequestObjectWalletRequest) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *WalletRequestObjectWalletRequest) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetMerchantNotifyURL

`func (o *WalletRequestObjectWalletRequest) GetMerchantNotifyURL() string`

GetMerchantNotifyURL returns the MerchantNotifyURL field if non-nil, zero value otherwise.

### GetMerchantNotifyURLOk

`func (o *WalletRequestObjectWalletRequest) GetMerchantNotifyURLOk() (*string, bool)`

GetMerchantNotifyURLOk returns a tuple with the MerchantNotifyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantNotifyURL

`func (o *WalletRequestObjectWalletRequest) SetMerchantNotifyURL(v string)`

SetMerchantNotifyURL sets MerchantNotifyURL field to given value.

### HasMerchantNotifyURL

`func (o *WalletRequestObjectWalletRequest) HasMerchantNotifyURL() bool`

HasMerchantNotifyURL returns a boolean if a field has been set.

### GetReference

`func (o *WalletRequestObjectWalletRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *WalletRequestObjectWalletRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *WalletRequestObjectWalletRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *WalletRequestObjectWalletRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTransactionType

`func (o *WalletRequestObjectWalletRequest) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *WalletRequestObjectWalletRequest) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *WalletRequestObjectWalletRequest) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *WalletRequestObjectWalletRequest) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *WalletRequestObjectWalletRequest) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *WalletRequestObjectWalletRequest) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *WalletRequestObjectWalletRequest) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *WalletRequestObjectWalletRequest) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetTransactionIdentification

`func (o *WalletRequestObjectWalletRequest) GetTransactionIdentification() string`

GetTransactionIdentification returns the TransactionIdentification field if non-nil, zero value otherwise.

### GetTransactionIdentificationOk

`func (o *WalletRequestObjectWalletRequest) GetTransactionIdentificationOk() (*string, bool)`

GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentification

`func (o *WalletRequestObjectWalletRequest) SetTransactionIdentification(v string)`

SetTransactionIdentification sets TransactionIdentification field to given value.

### HasTransactionIdentification

`func (o *WalletRequestObjectWalletRequest) HasTransactionIdentification() bool`

HasTransactionIdentification returns a boolean if a field has been set.

### GetTrasactionDateTime

`func (o *WalletRequestObjectWalletRequest) GetTrasactionDateTime() string`

GetTrasactionDateTime returns the TrasactionDateTime field if non-nil, zero value otherwise.

### GetTrasactionDateTimeOk

`func (o *WalletRequestObjectWalletRequest) GetTrasactionDateTimeOk() (*string, bool)`

GetTrasactionDateTimeOk returns a tuple with the TrasactionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrasactionDateTime

`func (o *WalletRequestObjectWalletRequest) SetTrasactionDateTime(v string)`

SetTrasactionDateTime sets TrasactionDateTime field to given value.

### HasTrasactionDateTime

`func (o *WalletRequestObjectWalletRequest) HasTrasactionDateTime() bool`

HasTrasactionDateTime returns a boolean if a field has been set.

### GetTransactionTimeout

`func (o *WalletRequestObjectWalletRequest) GetTransactionTimeout() float32`

GetTransactionTimeout returns the TransactionTimeout field if non-nil, zero value otherwise.

### GetTransactionTimeoutOk

`func (o *WalletRequestObjectWalletRequest) GetTransactionTimeoutOk() (*float32, bool)`

GetTransactionTimeoutOk returns a tuple with the TransactionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTimeout

`func (o *WalletRequestObjectWalletRequest) SetTransactionTimeout(v float32)`

SetTransactionTimeout sets TransactionTimeout field to given value.

### HasTransactionTimeout

`func (o *WalletRequestObjectWalletRequest) HasTransactionTimeout() bool`

HasTransactionTimeout returns a boolean if a field has been set.

### GetPOSType

`func (o *WalletRequestObjectWalletRequest) GetPOSType() string`

GetPOSType returns the POSType field if non-nil, zero value otherwise.

### GetPOSTypeOk

`func (o *WalletRequestObjectWalletRequest) GetPOSTypeOk() (*string, bool)`

GetPOSTypeOk returns a tuple with the POSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSType

`func (o *WalletRequestObjectWalletRequest) SetPOSType(v string)`

SetPOSType sets POSType field to given value.

### HasPOSType

`func (o *WalletRequestObjectWalletRequest) HasPOSType() bool`

HasPOSType returns a boolean if a field has been set.

### GetPOSVersion

`func (o *WalletRequestObjectWalletRequest) GetPOSVersion() string`

GetPOSVersion returns the POSVersion field if non-nil, zero value otherwise.

### GetPOSVersionOk

`func (o *WalletRequestObjectWalletRequest) GetPOSVersionOk() (*string, bool)`

GetPOSVersionOk returns a tuple with the POSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSVersion

`func (o *WalletRequestObjectWalletRequest) SetPOSVersion(v string)`

SetPOSVersion sets POSVersion field to given value.

### HasPOSVersion

`func (o *WalletRequestObjectWalletRequest) HasPOSVersion() bool`

HasPOSVersion returns a boolean if a field has been set.

### GetPOSAddress

`func (o *WalletRequestObjectWalletRequest) GetPOSAddress() string`

GetPOSAddress returns the POSAddress field if non-nil, zero value otherwise.

### GetPOSAddressOk

`func (o *WalletRequestObjectWalletRequest) GetPOSAddressOk() (*string, bool)`

GetPOSAddressOk returns a tuple with the POSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSAddress

`func (o *WalletRequestObjectWalletRequest) SetPOSAddress(v string)`

SetPOSAddress sets POSAddress field to given value.

### HasPOSAddress

`func (o *WalletRequestObjectWalletRequest) HasPOSAddress() bool`

HasPOSAddress returns a boolean if a field has been set.

### GetPOSSerial

`func (o *WalletRequestObjectWalletRequest) GetPOSSerial() string`

GetPOSSerial returns the POSSerial field if non-nil, zero value otherwise.

### GetPOSSerialOk

`func (o *WalletRequestObjectWalletRequest) GetPOSSerialOk() (*string, bool)`

GetPOSSerialOk returns a tuple with the POSSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSSerial

`func (o *WalletRequestObjectWalletRequest) SetPOSSerial(v string)`

SetPOSSerial sets POSSerial field to given value.

### HasPOSSerial

`func (o *WalletRequestObjectWalletRequest) HasPOSSerial() bool`

HasPOSSerial returns a boolean if a field has been set.

### GetPOSGEO

`func (o *WalletRequestObjectWalletRequest) GetPOSGEO() string`

GetPOSGEO returns the POSGEO field if non-nil, zero value otherwise.

### GetPOSGEOOk

`func (o *WalletRequestObjectWalletRequest) GetPOSGEOOk() (*string, bool)`

GetPOSGEOOk returns a tuple with the POSGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSGEO

`func (o *WalletRequestObjectWalletRequest) SetPOSGEO(v string)`

SetPOSGEO sets POSGEO field to given value.

### HasPOSGEO

`func (o *WalletRequestObjectWalletRequest) HasPOSGEO() bool`

HasPOSGEO returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *WalletRequestObjectWalletRequest) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *WalletRequestObjectWalletRequest) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *WalletRequestObjectWalletRequest) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *WalletRequestObjectWalletRequest) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetReadingDeviceVersion

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceVersion() string`

GetReadingDeviceVersion returns the ReadingDeviceVersion field if non-nil, zero value otherwise.

### GetReadingDeviceVersionOk

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceVersionOk() (*string, bool)`

GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceVersion

`func (o *WalletRequestObjectWalletRequest) SetReadingDeviceVersion(v string)`

SetReadingDeviceVersion sets ReadingDeviceVersion field to given value.

### HasReadingDeviceVersion

`func (o *WalletRequestObjectWalletRequest) HasReadingDeviceVersion() bool`

HasReadingDeviceVersion returns a boolean if a field has been set.

### GetReadingDeviceAddress

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceAddress() string`

GetReadingDeviceAddress returns the ReadingDeviceAddress field if non-nil, zero value otherwise.

### GetReadingDeviceAddressOk

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceAddressOk() (*string, bool)`

GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceAddress

`func (o *WalletRequestObjectWalletRequest) SetReadingDeviceAddress(v string)`

SetReadingDeviceAddress sets ReadingDeviceAddress field to given value.

### HasReadingDeviceAddress

`func (o *WalletRequestObjectWalletRequest) HasReadingDeviceAddress() bool`

HasReadingDeviceAddress returns a boolean if a field has been set.

### GetReadingDeviceSerial

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceSerial() string`

GetReadingDeviceSerial returns the ReadingDeviceSerial field if non-nil, zero value otherwise.

### GetReadingDeviceSerialOk

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceSerialOk() (*string, bool)`

GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceSerial

`func (o *WalletRequestObjectWalletRequest) SetReadingDeviceSerial(v string)`

SetReadingDeviceSerial sets ReadingDeviceSerial field to given value.

### HasReadingDeviceSerial

`func (o *WalletRequestObjectWalletRequest) HasReadingDeviceSerial() bool`

HasReadingDeviceSerial returns a boolean if a field has been set.

### GetReadingDeviceGEO

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceGEO() string`

GetReadingDeviceGEO returns the ReadingDeviceGEO field if non-nil, zero value otherwise.

### GetReadingDeviceGEOOk

`func (o *WalletRequestObjectWalletRequest) GetReadingDeviceGEOOk() (*string, bool)`

GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceGEO

`func (o *WalletRequestObjectWalletRequest) SetReadingDeviceGEO(v string)`

SetReadingDeviceGEO sets ReadingDeviceGEO field to given value.

### HasReadingDeviceGEO

`func (o *WalletRequestObjectWalletRequest) HasReadingDeviceGEO() bool`

HasReadingDeviceGEO returns a boolean if a field has been set.

### GetUserID

`func (o *WalletRequestObjectWalletRequest) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *WalletRequestObjectWalletRequest) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *WalletRequestObjectWalletRequest) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *WalletRequestObjectWalletRequest) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *WalletRequestObjectWalletRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *WalletRequestObjectWalletRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *WalletRequestObjectWalletRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *WalletRequestObjectWalletRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *WalletRequestObjectWalletRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *WalletRequestObjectWalletRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *WalletRequestObjectWalletRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *WalletRequestObjectWalletRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmount

`func (o *WalletRequestObjectWalletRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletRequestObjectWalletRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletRequestObjectWalletRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WalletRequestObjectWalletRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCashbackAmount

`func (o *WalletRequestObjectWalletRequest) GetCashbackAmount() float32`

GetCashbackAmount returns the CashbackAmount field if non-nil, zero value otherwise.

### GetCashbackAmountOk

`func (o *WalletRequestObjectWalletRequest) GetCashbackAmountOk() (*float32, bool)`

GetCashbackAmountOk returns a tuple with the CashbackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackAmount

`func (o *WalletRequestObjectWalletRequest) SetCashbackAmount(v float32)`

SetCashbackAmount sets CashbackAmount field to given value.

### HasCashbackAmount

`func (o *WalletRequestObjectWalletRequest) HasCashbackAmount() bool`

HasCashbackAmount returns a boolean if a field has been set.

### GetTipAmount

`func (o *WalletRequestObjectWalletRequest) GetTipAmount() float32`

GetTipAmount returns the TipAmount field if non-nil, zero value otherwise.

### GetTipAmountOk

`func (o *WalletRequestObjectWalletRequest) GetTipAmountOk() (*float32, bool)`

GetTipAmountOk returns a tuple with the TipAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipAmount

`func (o *WalletRequestObjectWalletRequest) SetTipAmount(v float32)`

SetTipAmount sets TipAmount field to given value.

### HasTipAmount

`func (o *WalletRequestObjectWalletRequest) HasTipAmount() bool`

HasTipAmount returns a boolean if a field has been set.

### GetWalletIdentification

`func (o *WalletRequestObjectWalletRequest) GetWalletIdentification() string`

GetWalletIdentification returns the WalletIdentification field if non-nil, zero value otherwise.

### GetWalletIdentificationOk

`func (o *WalletRequestObjectWalletRequest) GetWalletIdentificationOk() (*string, bool)`

GetWalletIdentificationOk returns a tuple with the WalletIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletIdentification

`func (o *WalletRequestObjectWalletRequest) SetWalletIdentification(v string)`

SetWalletIdentification sets WalletIdentification field to given value.

### HasWalletIdentification

`func (o *WalletRequestObjectWalletRequest) HasWalletIdentification() bool`

HasWalletIdentification returns a boolean if a field has been set.

### GetPayer

`func (o *WalletRequestObjectWalletRequest) GetPayer() SaleObjectSalePayer`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *WalletRequestObjectWalletRequest) GetPayerOk() (*SaleObjectSalePayer, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *WalletRequestObjectWalletRequest) SetPayer(v SaleObjectSalePayer)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *WalletRequestObjectWalletRequest) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetCustomer

`func (o *WalletRequestObjectWalletRequest) GetCustomer() SaleObjectSaleCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *WalletRequestObjectWalletRequest) GetCustomerOk() (*SaleObjectSaleCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *WalletRequestObjectWalletRequest) SetCustomer(v SaleObjectSaleCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *WalletRequestObjectWalletRequest) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSeller

`func (o *WalletRequestObjectWalletRequest) GetSeller() SaleObjectSaleSeller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *WalletRequestObjectWalletRequest) GetSellerOk() (*SaleObjectSaleSeller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *WalletRequestObjectWalletRequest) SetSeller(v SaleObjectSaleSeller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *WalletRequestObjectWalletRequest) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetProducts

`func (o *WalletRequestObjectWalletRequest) GetProducts() []SaleResponseObjectSaleResponseProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *WalletRequestObjectWalletRequest) GetProductsOk() (*[]SaleResponseObjectSaleResponseProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *WalletRequestObjectWalletRequest) SetProducts(v []SaleResponseObjectSaleResponseProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *WalletRequestObjectWalletRequest) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetTaxRefundType

`func (o *WalletRequestObjectWalletRequest) GetTaxRefundType() string`

GetTaxRefundType returns the TaxRefundType field if non-nil, zero value otherwise.

### GetTaxRefundTypeOk

`func (o *WalletRequestObjectWalletRequest) GetTaxRefundTypeOk() (*string, bool)`

GetTaxRefundTypeOk returns a tuple with the TaxRefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRefundType

`func (o *WalletRequestObjectWalletRequest) SetTaxRefundType(v string)`

SetTaxRefundType sets TaxRefundType field to given value.

### HasTaxRefundType

`func (o *WalletRequestObjectWalletRequest) HasTaxRefundType() bool`

HasTaxRefundType returns a boolean if a field has been set.

### GetValidThru

`func (o *WalletRequestObjectWalletRequest) GetValidThru() time.Time`

GetValidThru returns the ValidThru field if non-nil, zero value otherwise.

### GetValidThruOk

`func (o *WalletRequestObjectWalletRequest) GetValidThruOk() (*time.Time, bool)`

GetValidThruOk returns a tuple with the ValidThru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidThru

`func (o *WalletRequestObjectWalletRequest) SetValidThru(v time.Time)`

SetValidThru sets ValidThru field to given value.

### HasValidThru

`func (o *WalletRequestObjectWalletRequest) HasValidThru() bool`

HasValidThru returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *WalletRequestObjectWalletRequest) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *WalletRequestObjectWalletRequest) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *WalletRequestObjectWalletRequest) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *WalletRequestObjectWalletRequest) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *WalletRequestObjectWalletRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *WalletRequestObjectWalletRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *WalletRequestObjectWalletRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *WalletRequestObjectWalletRequest) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetTerminalID

`func (o *WalletRequestObjectWalletRequest) GetTerminalID() string`

GetTerminalID returns the TerminalID field if non-nil, zero value otherwise.

### GetTerminalIDOk

`func (o *WalletRequestObjectWalletRequest) GetTerminalIDOk() (*string, bool)`

GetTerminalIDOk returns a tuple with the TerminalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalID

`func (o *WalletRequestObjectWalletRequest) SetTerminalID(v string)`

SetTerminalID sets TerminalID field to given value.

### HasTerminalID

`func (o *WalletRequestObjectWalletRequest) HasTerminalID() bool`

HasTerminalID returns a boolean if a field has been set.

### GetTerminalTrace

`func (o *WalletRequestObjectWalletRequest) GetTerminalTrace() int32`

GetTerminalTrace returns the TerminalTrace field if non-nil, zero value otherwise.

### GetTerminalTraceOk

`func (o *WalletRequestObjectWalletRequest) GetTerminalTraceOk() (*int32, bool)`

GetTerminalTraceOk returns a tuple with the TerminalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTrace

`func (o *WalletRequestObjectWalletRequest) SetTerminalTrace(v int32)`

SetTerminalTrace sets TerminalTrace field to given value.

### HasTerminalTrace

`func (o *WalletRequestObjectWalletRequest) HasTerminalTrace() bool`

HasTerminalTrace returns a boolean if a field has been set.

### GetSettlementBatchNumber

`func (o *WalletRequestObjectWalletRequest) GetSettlementBatchNumber() int32`

GetSettlementBatchNumber returns the SettlementBatchNumber field if non-nil, zero value otherwise.

### GetSettlementBatchNumberOk

`func (o *WalletRequestObjectWalletRequest) GetSettlementBatchNumberOk() (*int32, bool)`

GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBatchNumber

`func (o *WalletRequestObjectWalletRequest) SetSettlementBatchNumber(v int32)`

SetSettlementBatchNumber sets SettlementBatchNumber field to given value.

### HasSettlementBatchNumber

`func (o *WalletRequestObjectWalletRequest) HasSettlementBatchNumber() bool`

HasSettlementBatchNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


