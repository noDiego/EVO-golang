# AuthorizeSaleObjectAuthorizeSale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyIdentification** | **string** | ID que identifica la companía desde donde proviene la petición. | 
**SystemIdentification** | **string** | ID que identifica el sistema desde donde proviene la petición. | 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**RequestType** | Pointer to **string** | Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**RequestKey** | Pointer to **string** | ID generado para la identificación por parte del Plataforma de la información generada en la ejecución de un GetCard o un Payment Method. Será necesario para que un mensaje de Sale, Void, PaymentMethod o Enrollment  identifique el contexto generado y lo utilice para esa operación. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**Block** | Pointer to **string** | ID que identifica a un grupo de transacciones que serán confirmadas o canceladas. | [optional] 
**RequiredInformation** | Pointer to [**[]SaleObjectSaleRequiredInformation**](SaleObjectSaleRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**Ticket** | Pointer to **string** | Ticket Digitalizado ( Total o parte del mismo por ejemplo la Firma Digitalizada )    codificado en Base64. | [optional] 
**TicketAnswerKey** | Pointer to **string** | Identificador Unívoco de la Transacción que se quiere Referenciar a la cual pertenece el Ticket Digitalizado. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. Si la firma fue capturada previamente y se envía en el requerimiento de la misma Operación Sale, Authorize*, Void, Return, Adjustment, DebtPayment, VoidDebtPayment o Enrollment no es necesario que se envíe este elemento o campo. | [optional] 
**Timeout** | Pointer to **float32** | Tiempo de espera que el POS espera al PINPAD para obtener la respuesta al requerimiento.  | [optional] 
**MerchantNotifyURL** | Pointer to **string** | URL para notificación del comercio | [optional] 
**IsReverse** | Pointer to **float32** | Identificador de Reversa. | [optional] 
**ReasonReverse** | Pointer to **string** | Motivo por el cual se requiere generar la reversa. | [optional] 
**ReasonSequenceBreak** | Pointer to **string** | Motivo por el cual se requiere romper la secuencia. | [optional] 
**Reference** | Pointer to **string** | Referencia de la transacción para el punto de venta | [optional] 
**TransactionDescription** | Pointer to **string** | Descripción del tipo de operación que se realizará | [optional] 
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
**ReadingDeviceGEO** | Pointer to **string** | Coordenadas Geográficas del dispositivo de lectura | [optional] 
**UserID** | Pointer to **string** | Identificador del usuario que está realizando la Transacción. | [optional] 
**UserName** | Pointer to **string** | Nombre del usuario que está realizando la Transacción. | [optional] 
**Amount** | **float32** | Importe o Monto de la Transacción. | 
**AlternativeAmount** | Pointer to **float32** | Monto con la que se realizó transacción. Si este valor es recibido, la búsqueda de los planes será limitada con este criterio. | [optional] 
**CashbackAmount** | Pointer to **float32** | Monto del dinero en efectivo (cashback). | [optional] 
**TipAmount** | Pointer to **float32** | Importe o Monto de la Propina. | [optional] 
**PromotedAmount** | Pointer to **float32** | Monto sujeto a promoción, monto de la operación. | [optional] 
**CurrencyCode** | Pointer to **string** | Código de Moneda - ISO 4217 &lt;https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica &lt;br /&gt;   * Num   - Alpha - Description &lt;br /&gt;   * &#39;032&#39; - &#39;ARS&#39; - Pesos Argentinos &lt;br /&gt;   * &#39;152&#39; - &#39;CLP&#39; - Pesos Chilenos &lt;br/&gt;   * &#39;484&#39; - &#39;MXN&#39; - Pesos Mexicanos &lt;br/&gt;   * &#39;840&#39; - &#39;USD&#39; - dólares Americanos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;858&#39; - &#39;UYU&#39; - Pesos Uruguayos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;986&#39; - &#39;BRL&#39; - Real Brasileño | [optional] 
**FacilityPayments** | Pointer to **float32** | Cantidad de cuotas en las que será realizada la transacción | [optional] 
**FacilityType** | Pointer to **string** | Tipo de Plan de Financiación | [optional] 
**Plan** | Pointer to **string** | Código/ID de Plan ( obtenido por la Operación PaymentMethod ) , en caso de ser enviado no se requiere en la Transacción el envío de CurrencyCode ni FacilityPayments | [optional] 
**CardReadMode** | Pointer to **string** | Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes ) | [optional] 
**CardGetMode** | Pointer to **string** | Indican por cada elemento que contiene los datos sensibles, si están encriptados  y también el algoritmo usado. En caso de no estar especificado se asume PLAIN. | [optional] 
**CardNumber** | Pointer to **string** | Número de Tarjeta, en el caso de las respuestas el mismo estará enmascarado. | [optional] 
**CardNumberMasked** | Pointer to **string** | Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón.    | [optional] 
**CardNumberEncrypted** | Pointer to **string** | Número de tarjeta encriptado.                        | [optional] 
**CardExp** | Pointer to **string** | Fecha de vencimiento de la tarjeta. Este dato será necesario si el modo de ingreso fue manual/digitada. | [optional] 
**CardCryptogram** | Pointer to **string** | Tags EMV en formato TLV Leídos. | [optional] 
**CardAppName** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppIdentifier** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppLabel** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthRequestCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthResponseCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**Track1** | Pointer to **string** | Banda Número 1 de la tarjeta. Este dato será necesario si el modo de ingreso fue por banda (deslizando la banda por el lector). | [optional] 
**Track2** | Pointer to **string** | Banda Número 2 de la tarjeta. Este dato será necesario si el modo de ingreso fue por banda (deslizando la banda por el lector). | [optional] 
**InputTokens** | Pointer to [**[]SaleObjectSaleInputTokens**](SaleObjectSaleInputTokens.md) | Tokens. | [optional] 
**SecurityCode** | Pointer to **string** | Código de seguridad de la tarjeta. | [optional] 
**Pin** | Pointer to **string** | PIN block | [optional] 
**CredentialToken** | Pointer to **string** | Token asociado a la Credencial Enrolada | [optional] 
**CredentialIssuerToken** | Pointer to **string** | Emisor del Token asociado a la credencial enrolada | [optional] 
**Payer** | Pointer to [**SaleObjectSalePayer**](SaleObjectSalePayer.md) |  | [optional] 
**Customer** | Pointer to [**SaleObjectSaleCustomer**](SaleObjectSaleCustomer.md) |  | [optional] 
**Seller** | Pointer to [**SaleObjectSaleSeller**](SaleObjectSaleSeller.md) |  | [optional] 
**Products** | Pointer to [**[]SaleObjectSaleProducts**](SaleObjectSaleProducts.md) | Detalle de Productos de la Operación. | [optional] 
**TaxRefundType** | Pointer to **string** | Esquema de Devolución de Impuestos a utilizar en la transacción | [optional] 
**AuthCode** | Pointer to **string** | Código de autorización retornado por el Host que resuelve la transacción. | [optional] 
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relacionado a cada uno de los planes disponibles. | [optional] 
**TerminalID** | Pointer to **string** | Identificador de Terminal por el cual se envía la Transacción al Host. | [optional] 
**TerminalTrace** | Pointer to **int32** | Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID. | [optional] 
**SettlementBatchNumber** | Pointer to **int32** | Para aquellos host que exista el concepto de lote, es el número de lote al cual pertenece la transacción. | [optional] 

## Methods

### NewAuthorizeSaleObjectAuthorizeSale

`func NewAuthorizeSaleObjectAuthorizeSale(companyIdentification string, systemIdentification string, amount float32, ) *AuthorizeSaleObjectAuthorizeSale`

NewAuthorizeSaleObjectAuthorizeSale instantiates a new AuthorizeSaleObjectAuthorizeSale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeSaleObjectAuthorizeSaleWithDefaults

`func NewAuthorizeSaleObjectAuthorizeSaleWithDefaults() *AuthorizeSaleObjectAuthorizeSale`

NewAuthorizeSaleObjectAuthorizeSaleWithDefaults instantiates a new AuthorizeSaleObjectAuthorizeSale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetSystemIdentification

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *AuthorizeSaleObjectAuthorizeSale) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetBranchIdentification

`func (o *AuthorizeSaleObjectAuthorizeSale) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *AuthorizeSaleObjectAuthorizeSale) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *AuthorizeSaleObjectAuthorizeSale) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetRequestType

`func (o *AuthorizeSaleObjectAuthorizeSale) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *AuthorizeSaleObjectAuthorizeSale) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *AuthorizeSaleObjectAuthorizeSale) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetRequestKey

`func (o *AuthorizeSaleObjectAuthorizeSale) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *AuthorizeSaleObjectAuthorizeSale) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *AuthorizeSaleObjectAuthorizeSale) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetServiceVersion

`func (o *AuthorizeSaleObjectAuthorizeSale) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *AuthorizeSaleObjectAuthorizeSale) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *AuthorizeSaleObjectAuthorizeSale) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *AuthorizeSaleObjectAuthorizeSale) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *AuthorizeSaleObjectAuthorizeSale) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *AuthorizeSaleObjectAuthorizeSale) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *AuthorizeSaleObjectAuthorizeSale) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *AuthorizeSaleObjectAuthorizeSale) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *AuthorizeSaleObjectAuthorizeSale) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *AuthorizeSaleObjectAuthorizeSale) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *AuthorizeSaleObjectAuthorizeSale) GetRequiredInformation() []SaleObjectSaleRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetRequiredInformationOk() (*[]SaleObjectSaleRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *AuthorizeSaleObjectAuthorizeSale) SetRequiredInformation(v []SaleObjectSaleRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *AuthorizeSaleObjectAuthorizeSale) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetTicket

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *AuthorizeSaleObjectAuthorizeSale) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *AuthorizeSaleObjectAuthorizeSale) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTicketAnswerKey

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTicketAnswerKey() string`

GetTicketAnswerKey returns the TicketAnswerKey field if non-nil, zero value otherwise.

### GetTicketAnswerKeyOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTicketAnswerKeyOk() (*string, bool)`

GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketAnswerKey

`func (o *AuthorizeSaleObjectAuthorizeSale) SetTicketAnswerKey(v string)`

SetTicketAnswerKey sets TicketAnswerKey field to given value.

### HasTicketAnswerKey

`func (o *AuthorizeSaleObjectAuthorizeSale) HasTicketAnswerKey() bool`

HasTicketAnswerKey returns a boolean if a field has been set.

### GetTimeout

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AuthorizeSaleObjectAuthorizeSale) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *AuthorizeSaleObjectAuthorizeSale) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetMerchantNotifyURL

`func (o *AuthorizeSaleObjectAuthorizeSale) GetMerchantNotifyURL() string`

GetMerchantNotifyURL returns the MerchantNotifyURL field if non-nil, zero value otherwise.

### GetMerchantNotifyURLOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetMerchantNotifyURLOk() (*string, bool)`

GetMerchantNotifyURLOk returns a tuple with the MerchantNotifyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantNotifyURL

`func (o *AuthorizeSaleObjectAuthorizeSale) SetMerchantNotifyURL(v string)`

SetMerchantNotifyURL sets MerchantNotifyURL field to given value.

### HasMerchantNotifyURL

`func (o *AuthorizeSaleObjectAuthorizeSale) HasMerchantNotifyURL() bool`

HasMerchantNotifyURL returns a boolean if a field has been set.

### GetIsReverse

`func (o *AuthorizeSaleObjectAuthorizeSale) GetIsReverse() float32`

GetIsReverse returns the IsReverse field if non-nil, zero value otherwise.

### GetIsReverseOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetIsReverseOk() (*float32, bool)`

GetIsReverseOk returns a tuple with the IsReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReverse

`func (o *AuthorizeSaleObjectAuthorizeSale) SetIsReverse(v float32)`

SetIsReverse sets IsReverse field to given value.

### HasIsReverse

`func (o *AuthorizeSaleObjectAuthorizeSale) HasIsReverse() bool`

HasIsReverse returns a boolean if a field has been set.

### GetReasonReverse

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReasonReverse() string`

GetReasonReverse returns the ReasonReverse field if non-nil, zero value otherwise.

### GetReasonReverseOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReasonReverseOk() (*string, bool)`

GetReasonReverseOk returns a tuple with the ReasonReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonReverse

`func (o *AuthorizeSaleObjectAuthorizeSale) SetReasonReverse(v string)`

SetReasonReverse sets ReasonReverse field to given value.

### HasReasonReverse

`func (o *AuthorizeSaleObjectAuthorizeSale) HasReasonReverse() bool`

HasReasonReverse returns a boolean if a field has been set.

### GetReasonSequenceBreak

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReasonSequenceBreak() string`

GetReasonSequenceBreak returns the ReasonSequenceBreak field if non-nil, zero value otherwise.

### GetReasonSequenceBreakOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReasonSequenceBreakOk() (*string, bool)`

GetReasonSequenceBreakOk returns a tuple with the ReasonSequenceBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSequenceBreak

`func (o *AuthorizeSaleObjectAuthorizeSale) SetReasonSequenceBreak(v string)`

SetReasonSequenceBreak sets ReasonSequenceBreak field to given value.

### HasReasonSequenceBreak

`func (o *AuthorizeSaleObjectAuthorizeSale) HasReasonSequenceBreak() bool`

HasReasonSequenceBreak returns a boolean if a field has been set.

### GetReference

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AuthorizeSaleObjectAuthorizeSale) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AuthorizeSaleObjectAuthorizeSale) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *AuthorizeSaleObjectAuthorizeSale) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *AuthorizeSaleObjectAuthorizeSale) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetPOSType

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSType() string`

GetPOSType returns the POSType field if non-nil, zero value otherwise.

### GetPOSTypeOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSTypeOk() (*string, bool)`

GetPOSTypeOk returns a tuple with the POSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSType

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSType(v string)`

SetPOSType sets POSType field to given value.

### HasPOSType

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSType() bool`

HasPOSType returns a boolean if a field has been set.

### GetPOSVersion

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSVersion() string`

GetPOSVersion returns the POSVersion field if non-nil, zero value otherwise.

### GetPOSVersionOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSVersionOk() (*string, bool)`

GetPOSVersionOk returns a tuple with the POSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSVersion

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSVersion(v string)`

SetPOSVersion sets POSVersion field to given value.

### HasPOSVersion

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSVersion() bool`

HasPOSVersion returns a boolean if a field has been set.

### GetPOSAddress

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSAddress() string`

GetPOSAddress returns the POSAddress field if non-nil, zero value otherwise.

### GetPOSAddressOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSAddressOk() (*string, bool)`

GetPOSAddressOk returns a tuple with the POSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSAddress

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSAddress(v string)`

SetPOSAddress sets POSAddress field to given value.

### HasPOSAddress

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSAddress() bool`

HasPOSAddress returns a boolean if a field has been set.

### GetPOSSerial

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSSerial() string`

GetPOSSerial returns the POSSerial field if non-nil, zero value otherwise.

### GetPOSSerialOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSSerialOk() (*string, bool)`

GetPOSSerialOk returns a tuple with the POSSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSSerial

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSSerial(v string)`

SetPOSSerial sets POSSerial field to given value.

### HasPOSSerial

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSSerial() bool`

HasPOSSerial returns a boolean if a field has been set.

### GetPOSGEO

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSGEO() string`

GetPOSGEO returns the POSGEO field if non-nil, zero value otherwise.

### GetPOSGEOOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSGEOOk() (*string, bool)`

GetPOSGEOOk returns a tuple with the POSGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSGEO

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSGEO(v string)`

SetPOSGEO sets POSGEO field to given value.

### HasPOSGEO

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSGEO() bool`

HasPOSGEO returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetReadingDeviceVersion

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceVersion() string`

GetReadingDeviceVersion returns the ReadingDeviceVersion field if non-nil, zero value otherwise.

### GetReadingDeviceVersionOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceVersionOk() (*string, bool)`

GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceVersion

`func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceVersion(v string)`

SetReadingDeviceVersion sets ReadingDeviceVersion field to given value.

### HasReadingDeviceVersion

`func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceVersion() bool`

HasReadingDeviceVersion returns a boolean if a field has been set.

### GetReadingDeviceAddress

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceAddress() string`

GetReadingDeviceAddress returns the ReadingDeviceAddress field if non-nil, zero value otherwise.

### GetReadingDeviceAddressOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceAddressOk() (*string, bool)`

GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceAddress

`func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceAddress(v string)`

SetReadingDeviceAddress sets ReadingDeviceAddress field to given value.

### HasReadingDeviceAddress

`func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceAddress() bool`

HasReadingDeviceAddress returns a boolean if a field has been set.

### GetReadingDeviceSerial

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceSerial() string`

GetReadingDeviceSerial returns the ReadingDeviceSerial field if non-nil, zero value otherwise.

### GetReadingDeviceSerialOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceSerialOk() (*string, bool)`

GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceSerial

`func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceSerial(v string)`

SetReadingDeviceSerial sets ReadingDeviceSerial field to given value.

### HasReadingDeviceSerial

`func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceSerial() bool`

HasReadingDeviceSerial returns a boolean if a field has been set.

### GetReadingDeviceGEO

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceGEO() string`

GetReadingDeviceGEO returns the ReadingDeviceGEO field if non-nil, zero value otherwise.

### GetReadingDeviceGEOOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceGEOOk() (*string, bool)`

GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceGEO

`func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceGEO(v string)`

SetReadingDeviceGEO sets ReadingDeviceGEO field to given value.

### HasReadingDeviceGEO

`func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceGEO() bool`

HasReadingDeviceGEO returns a boolean if a field has been set.

### GetUserID

`func (o *AuthorizeSaleObjectAuthorizeSale) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AuthorizeSaleObjectAuthorizeSale) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *AuthorizeSaleObjectAuthorizeSale) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *AuthorizeSaleObjectAuthorizeSale) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AuthorizeSaleObjectAuthorizeSale) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AuthorizeSaleObjectAuthorizeSale) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAlternativeAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) GetAlternativeAmount() float32`

GetAlternativeAmount returns the AlternativeAmount field if non-nil, zero value otherwise.

### GetAlternativeAmountOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetAlternativeAmountOk() (*float32, bool)`

GetAlternativeAmountOk returns a tuple with the AlternativeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) SetAlternativeAmount(v float32)`

SetAlternativeAmount sets AlternativeAmount field to given value.

### HasAlternativeAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) HasAlternativeAmount() bool`

HasAlternativeAmount returns a boolean if a field has been set.

### GetCashbackAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCashbackAmount() float32`

GetCashbackAmount returns the CashbackAmount field if non-nil, zero value otherwise.

### GetCashbackAmountOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCashbackAmountOk() (*float32, bool)`

GetCashbackAmountOk returns a tuple with the CashbackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCashbackAmount(v float32)`

SetCashbackAmount sets CashbackAmount field to given value.

### HasCashbackAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCashbackAmount() bool`

HasCashbackAmount returns a boolean if a field has been set.

### GetTipAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTipAmount() float32`

GetTipAmount returns the TipAmount field if non-nil, zero value otherwise.

### GetTipAmountOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTipAmountOk() (*float32, bool)`

GetTipAmountOk returns a tuple with the TipAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) SetTipAmount(v float32)`

SetTipAmount sets TipAmount field to given value.

### HasTipAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) HasTipAmount() bool`

HasTipAmount returns a boolean if a field has been set.

### GetPromotedAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPromotedAmount() float32`

GetPromotedAmount returns the PromotedAmount field if non-nil, zero value otherwise.

### GetPromotedAmountOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPromotedAmountOk() (*float32, bool)`

GetPromotedAmountOk returns a tuple with the PromotedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPromotedAmount(v float32)`

SetPromotedAmount sets PromotedAmount field to given value.

### HasPromotedAmount

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPromotedAmount() bool`

HasPromotedAmount returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetFacilityPayments

`func (o *AuthorizeSaleObjectAuthorizeSale) GetFacilityPayments() float32`

GetFacilityPayments returns the FacilityPayments field if non-nil, zero value otherwise.

### GetFacilityPaymentsOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetFacilityPaymentsOk() (*float32, bool)`

GetFacilityPaymentsOk returns a tuple with the FacilityPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityPayments

`func (o *AuthorizeSaleObjectAuthorizeSale) SetFacilityPayments(v float32)`

SetFacilityPayments sets FacilityPayments field to given value.

### HasFacilityPayments

`func (o *AuthorizeSaleObjectAuthorizeSale) HasFacilityPayments() bool`

HasFacilityPayments returns a boolean if a field has been set.

### GetFacilityType

`func (o *AuthorizeSaleObjectAuthorizeSale) GetFacilityType() string`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetFacilityTypeOk() (*string, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *AuthorizeSaleObjectAuthorizeSale) SetFacilityType(v string)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *AuthorizeSaleObjectAuthorizeSale) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### GetPlan

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetCardReadMode

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardReadMode() string`

GetCardReadMode returns the CardReadMode field if non-nil, zero value otherwise.

### GetCardReadModeOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardReadModeOk() (*string, bool)`

GetCardReadModeOk returns a tuple with the CardReadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadMode

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardReadMode(v string)`

SetCardReadMode sets CardReadMode field to given value.

### HasCardReadMode

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardReadMode() bool`

HasCardReadMode returns a boolean if a field has been set.

### GetCardGetMode

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardGetMode() string`

GetCardGetMode returns the CardGetMode field if non-nil, zero value otherwise.

### GetCardGetModeOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardGetModeOk() (*string, bool)`

GetCardGetModeOk returns a tuple with the CardGetMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardGetMode

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardGetMode(v string)`

SetCardGetMode sets CardGetMode field to given value.

### HasCardGetMode

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardGetMode() bool`

HasCardGetMode returns a boolean if a field has been set.

### GetCardNumber

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNumberMasked

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumberMasked() string`

GetCardNumberMasked returns the CardNumberMasked field if non-nil, zero value otherwise.

### GetCardNumberMaskedOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumberMaskedOk() (*string, bool)`

GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMasked

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardNumberMasked(v string)`

SetCardNumberMasked sets CardNumberMasked field to given value.

### HasCardNumberMasked

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardNumberMasked() bool`

HasCardNumberMasked returns a boolean if a field has been set.

### GetCardNumberEncrypted

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumberEncrypted() string`

GetCardNumberEncrypted returns the CardNumberEncrypted field if non-nil, zero value otherwise.

### GetCardNumberEncryptedOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumberEncryptedOk() (*string, bool)`

GetCardNumberEncryptedOk returns a tuple with the CardNumberEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberEncrypted

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardNumberEncrypted(v string)`

SetCardNumberEncrypted sets CardNumberEncrypted field to given value.

### HasCardNumberEncrypted

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardNumberEncrypted() bool`

HasCardNumberEncrypted returns a boolean if a field has been set.

### GetCardExp

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardExp() string`

GetCardExp returns the CardExp field if non-nil, zero value otherwise.

### GetCardExpOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardExpOk() (*string, bool)`

GetCardExpOk returns a tuple with the CardExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExp

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardExp(v string)`

SetCardExp sets CardExp field to given value.

### HasCardExp

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardExp() bool`

HasCardExp returns a boolean if a field has been set.

### GetCardCryptogram

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardCryptogram() string`

GetCardCryptogram returns the CardCryptogram field if non-nil, zero value otherwise.

### GetCardCryptogramOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardCryptogramOk() (*string, bool)`

GetCardCryptogramOk returns a tuple with the CardCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCryptogram

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardCryptogram(v string)`

SetCardCryptogram sets CardCryptogram field to given value.

### HasCardCryptogram

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardCryptogram() bool`

HasCardCryptogram returns a boolean if a field has been set.

### GetCardAppName

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppName() string`

GetCardAppName returns the CardAppName field if non-nil, zero value otherwise.

### GetCardAppNameOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppNameOk() (*string, bool)`

GetCardAppNameOk returns a tuple with the CardAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppName

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardAppName(v string)`

SetCardAppName sets CardAppName field to given value.

### HasCardAppName

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardAppName() bool`

HasCardAppName returns a boolean if a field has been set.

### GetCardAppIdentifier

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppIdentifier() string`

GetCardAppIdentifier returns the CardAppIdentifier field if non-nil, zero value otherwise.

### GetCardAppIdentifierOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppIdentifierOk() (*string, bool)`

GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppIdentifier

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardAppIdentifier(v string)`

SetCardAppIdentifier sets CardAppIdentifier field to given value.

### HasCardAppIdentifier

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardAppIdentifier() bool`

HasCardAppIdentifier returns a boolean if a field has been set.

### GetCardAppLabel

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppLabel() string`

GetCardAppLabel returns the CardAppLabel field if non-nil, zero value otherwise.

### GetCardAppLabelOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppLabelOk() (*string, bool)`

GetCardAppLabelOk returns a tuple with the CardAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppLabel

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardAppLabel(v string)`

SetCardAppLabel sets CardAppLabel field to given value.

### HasCardAppLabel

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardAppLabel() bool`

HasCardAppLabel returns a boolean if a field has been set.

### GetCardAuthRequestCryptogram

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAuthRequestCryptogram() string`

GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field if non-nil, zero value otherwise.

### GetCardAuthRequestCryptogramOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAuthRequestCryptogramOk() (*string, bool)`

GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthRequestCryptogram

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardAuthRequestCryptogram(v string)`

SetCardAuthRequestCryptogram sets CardAuthRequestCryptogram field to given value.

### HasCardAuthRequestCryptogram

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardAuthRequestCryptogram() bool`

HasCardAuthRequestCryptogram returns a boolean if a field has been set.

### GetCardAuthResponseCryptogram

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAuthResponseCryptogram() string`

GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field if non-nil, zero value otherwise.

### GetCardAuthResponseCryptogramOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAuthResponseCryptogramOk() (*string, bool)`

GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthResponseCryptogram

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCardAuthResponseCryptogram(v string)`

SetCardAuthResponseCryptogram sets CardAuthResponseCryptogram field to given value.

### HasCardAuthResponseCryptogram

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCardAuthResponseCryptogram() bool`

HasCardAuthResponseCryptogram returns a boolean if a field has been set.

### GetTrack1

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTrack1() string`

GetTrack1 returns the Track1 field if non-nil, zero value otherwise.

### GetTrack1Ok

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTrack1Ok() (*string, bool)`

GetTrack1Ok returns a tuple with the Track1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack1

`func (o *AuthorizeSaleObjectAuthorizeSale) SetTrack1(v string)`

SetTrack1 sets Track1 field to given value.

### HasTrack1

`func (o *AuthorizeSaleObjectAuthorizeSale) HasTrack1() bool`

HasTrack1 returns a boolean if a field has been set.

### GetTrack2

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTrack2() string`

GetTrack2 returns the Track2 field if non-nil, zero value otherwise.

### GetTrack2Ok

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTrack2Ok() (*string, bool)`

GetTrack2Ok returns a tuple with the Track2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack2

`func (o *AuthorizeSaleObjectAuthorizeSale) SetTrack2(v string)`

SetTrack2 sets Track2 field to given value.

### HasTrack2

`func (o *AuthorizeSaleObjectAuthorizeSale) HasTrack2() bool`

HasTrack2 returns a boolean if a field has been set.

### GetInputTokens

`func (o *AuthorizeSaleObjectAuthorizeSale) GetInputTokens() []SaleObjectSaleInputTokens`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetInputTokensOk() (*[]SaleObjectSaleInputTokens, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *AuthorizeSaleObjectAuthorizeSale) SetInputTokens(v []SaleObjectSaleInputTokens)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *AuthorizeSaleObjectAuthorizeSale) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetSecurityCode

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *AuthorizeSaleObjectAuthorizeSale) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *AuthorizeSaleObjectAuthorizeSale) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetPin

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetCredentialToken

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCredentialToken() string`

GetCredentialToken returns the CredentialToken field if non-nil, zero value otherwise.

### GetCredentialTokenOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCredentialTokenOk() (*string, bool)`

GetCredentialTokenOk returns a tuple with the CredentialToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialToken

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCredentialToken(v string)`

SetCredentialToken sets CredentialToken field to given value.

### HasCredentialToken

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCredentialToken() bool`

HasCredentialToken returns a boolean if a field has been set.

### GetCredentialIssuerToken

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCredentialIssuerToken() string`

GetCredentialIssuerToken returns the CredentialIssuerToken field if non-nil, zero value otherwise.

### GetCredentialIssuerTokenOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCredentialIssuerTokenOk() (*string, bool)`

GetCredentialIssuerTokenOk returns a tuple with the CredentialIssuerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuerToken

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCredentialIssuerToken(v string)`

SetCredentialIssuerToken sets CredentialIssuerToken field to given value.

### HasCredentialIssuerToken

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCredentialIssuerToken() bool`

HasCredentialIssuerToken returns a boolean if a field has been set.

### GetPayer

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPayer() SaleObjectSalePayer`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPayerOk() (*SaleObjectSalePayer, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPayer(v SaleObjectSalePayer)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetCustomer

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCustomer() SaleObjectSaleCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetCustomerOk() (*SaleObjectSaleCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *AuthorizeSaleObjectAuthorizeSale) SetCustomer(v SaleObjectSaleCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *AuthorizeSaleObjectAuthorizeSale) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSeller

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSeller() SaleObjectSaleSeller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSellerOk() (*SaleObjectSaleSeller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AuthorizeSaleObjectAuthorizeSale) SetSeller(v SaleObjectSaleSeller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *AuthorizeSaleObjectAuthorizeSale) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetProducts

`func (o *AuthorizeSaleObjectAuthorizeSale) GetProducts() []SaleObjectSaleProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetProductsOk() (*[]SaleObjectSaleProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AuthorizeSaleObjectAuthorizeSale) SetProducts(v []SaleObjectSaleProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AuthorizeSaleObjectAuthorizeSale) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetTaxRefundType

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTaxRefundType() string`

GetTaxRefundType returns the TaxRefundType field if non-nil, zero value otherwise.

### GetTaxRefundTypeOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTaxRefundTypeOk() (*string, bool)`

GetTaxRefundTypeOk returns a tuple with the TaxRefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRefundType

`func (o *AuthorizeSaleObjectAuthorizeSale) SetTaxRefundType(v string)`

SetTaxRefundType sets TaxRefundType field to given value.

### HasTaxRefundType

`func (o *AuthorizeSaleObjectAuthorizeSale) HasTaxRefundType() bool`

HasTaxRefundType returns a boolean if a field has been set.

### GetAuthCode

`func (o *AuthorizeSaleObjectAuthorizeSale) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *AuthorizeSaleObjectAuthorizeSale) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *AuthorizeSaleObjectAuthorizeSale) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *AuthorizeSaleObjectAuthorizeSale) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *AuthorizeSaleObjectAuthorizeSale) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *AuthorizeSaleObjectAuthorizeSale) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *AuthorizeSaleObjectAuthorizeSale) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *AuthorizeSaleObjectAuthorizeSale) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetTerminalID

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTerminalID() string`

GetTerminalID returns the TerminalID field if non-nil, zero value otherwise.

### GetTerminalIDOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTerminalIDOk() (*string, bool)`

GetTerminalIDOk returns a tuple with the TerminalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalID

`func (o *AuthorizeSaleObjectAuthorizeSale) SetTerminalID(v string)`

SetTerminalID sets TerminalID field to given value.

### HasTerminalID

`func (o *AuthorizeSaleObjectAuthorizeSale) HasTerminalID() bool`

HasTerminalID returns a boolean if a field has been set.

### GetTerminalTrace

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTerminalTrace() int32`

GetTerminalTrace returns the TerminalTrace field if non-nil, zero value otherwise.

### GetTerminalTraceOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetTerminalTraceOk() (*int32, bool)`

GetTerminalTraceOk returns a tuple with the TerminalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTrace

`func (o *AuthorizeSaleObjectAuthorizeSale) SetTerminalTrace(v int32)`

SetTerminalTrace sets TerminalTrace field to given value.

### HasTerminalTrace

`func (o *AuthorizeSaleObjectAuthorizeSale) HasTerminalTrace() bool`

HasTerminalTrace returns a boolean if a field has been set.

### GetSettlementBatchNumber

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSettlementBatchNumber() int32`

GetSettlementBatchNumber returns the SettlementBatchNumber field if non-nil, zero value otherwise.

### GetSettlementBatchNumberOk

`func (o *AuthorizeSaleObjectAuthorizeSale) GetSettlementBatchNumberOk() (*int32, bool)`

GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBatchNumber

`func (o *AuthorizeSaleObjectAuthorizeSale) SetSettlementBatchNumber(v int32)`

SetSettlementBatchNumber sets SettlementBatchNumber field to given value.

### HasSettlementBatchNumber

`func (o *AuthorizeSaleObjectAuthorizeSale) HasSettlementBatchNumber() bool`

HasSettlementBatchNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


