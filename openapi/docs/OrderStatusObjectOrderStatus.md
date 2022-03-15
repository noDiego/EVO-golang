# OrderStatusObjectOrderStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyIdentification** | **string** | ID que identifica la companía desde donde proviene la petición. | 
**SystemIdentification** | **string** | ID que identifica el sistema desde donde proviene la petición. | 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**RequestType** | Pointer to **string** | Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**Block** | Pointer to **string** | ID que identifica a un grupo de transacciones que serán confirmadas o canceladas. | [optional] 
**RequiredInformation** | Pointer to [**[]SaleObjectSaleRequiredInformation**](SaleObjectSaleRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**InputTokens** | Pointer to [**[]SaleObjectSaleInputTokens**](SaleObjectSaleInputTokens.md) | Tokens. | [optional] 
**Ticket** | Pointer to **string** | Ticket Digitalizado ( Total o parte del mismo por ejemplo la Firma Digitalizada )    codificado en Base64. | [optional] 
**TicketAnswerKey** | Pointer to **string** | Identificador Unívoco de la Transacción que se quiere Referenciar a la cual pertenece el Ticket Digitalizado. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. Si la firma fue capturada previamente y se envía en el requerimiento de la misma Operación Sale, Authorize*, Void, Return, Adjustment, DebtPayment, VoidDebtPayment o Enrollment no es necesario que se envíe este elemento o campo. | [optional] 
**Timeout** | Pointer to **float32** | Tiempo de espera que el POS espera al PINPAD para obtener la respuesta al requerimiento.  | [optional] 
**RequestKey** | Pointer to **string** | ID generado para la identificación por parte del Plataforma de la información generada en la ejecución de un GetCard o un Payment Method. Será necesario para que un mensaje de Sale, Void, PaymentMethod o Enrollment  identifique el contexto generado y lo utilice para esa operación. | [optional] 
**MerchantNotifyURL** | Pointer to **string** | URL para notificación del comercio | [optional] 
**IsReverse** | Pointer to **float32** | Identificador de Reversa. | [optional] 
**ReverseReason** | Pointer to **string** | Motivo por el cual se requiere generar la reversa. | [optional] 
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
**CardReadMode** | **string** | Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes ) | 
**CardNumber** | Pointer to **string** | Número de Tarjeta, en el caso de las respuestas el mismo estará enmascarado. | [optional] 
**CardNumberMasked** | Pointer to **string** | Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón.    | [optional] 
**CardNumberEncrypted** | Pointer to **string** | Número de tarjeta encriptado. | [optional] 
**CardExp** | Pointer to **string** | Fecha de vencimiento de la tarjeta. Este dato será necesario si el modo de ingreso fue manual/digitada. | [optional] 
**Track1** | Pointer to **string** | Banda Número 1 de la tarjeta. Este dato será necesario si el modo de ingreso fue por banda (deslizando la banda por el lector). | [optional] 
**Track2** | Pointer to **string** | Banda Número 2 de la tarjeta. Este dato será necesario si el modo de ingreso fue por banda (deslizando la banda por el lector). | [optional] 
**SecurityCode** | Pointer to **string** | Código de seguridad de la tarjeta. | [optional] 
**Pin** | Pointer to **string** | PIN block | [optional] 
**CardCryptogram** | Pointer to **string** | Tags EMV en formato TLV Leídos. | [optional] 
**CredentialToken** | Pointer to **string** | Token asociado a la Credencial Enrolada | [optional] 
**CredentialIssuerToken** | Pointer to **string** | Emisor del Token asociado a la credencial enrolada | [optional] 
**CardAppName** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppIdentifier** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppLabel** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthRequestCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthResponseCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**Payer** | Pointer to [**SaleObjectSalePayer**](SaleObjectSalePayer.md) |  | [optional] 
**Customer** | Pointer to [**SaleObjectSaleCustomer**](SaleObjectSaleCustomer.md) |  | [optional] 
**Seller** | Pointer to [**SaleObjectSaleSeller**](SaleObjectSaleSeller.md) |  | [optional] 
**Products** | Pointer to [**[]SaleObjectSaleProducts**](SaleObjectSaleProducts.md) | Detalle de Productos de la Operación. | [optional] 
**TaxRefundType** | Pointer to **string** | Esquema de Devolución de Impuestos a utilizar en la transacción | [optional] 
**AuthCode** | Pointer to **string** | Código de autorización retornado por el Host que resuelve la transacción. | [optional] 

## Methods

### NewOrderStatusObjectOrderStatus

`func NewOrderStatusObjectOrderStatus(companyIdentification string, systemIdentification string, amount float32, cardReadMode string, ) *OrderStatusObjectOrderStatus`

NewOrderStatusObjectOrderStatus instantiates a new OrderStatusObjectOrderStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusObjectOrderStatusWithDefaults

`func NewOrderStatusObjectOrderStatusWithDefaults() *OrderStatusObjectOrderStatus`

NewOrderStatusObjectOrderStatusWithDefaults instantiates a new OrderStatusObjectOrderStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *OrderStatusObjectOrderStatus) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *OrderStatusObjectOrderStatus) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *OrderStatusObjectOrderStatus) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetSystemIdentification

`func (o *OrderStatusObjectOrderStatus) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *OrderStatusObjectOrderStatus) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *OrderStatusObjectOrderStatus) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetBranchIdentification

`func (o *OrderStatusObjectOrderStatus) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *OrderStatusObjectOrderStatus) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *OrderStatusObjectOrderStatus) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *OrderStatusObjectOrderStatus) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *OrderStatusObjectOrderStatus) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *OrderStatusObjectOrderStatus) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *OrderStatusObjectOrderStatus) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *OrderStatusObjectOrderStatus) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetRequestType

`func (o *OrderStatusObjectOrderStatus) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *OrderStatusObjectOrderStatus) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *OrderStatusObjectOrderStatus) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *OrderStatusObjectOrderStatus) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *OrderStatusObjectOrderStatus) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *OrderStatusObjectOrderStatus) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *OrderStatusObjectOrderStatus) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *OrderStatusObjectOrderStatus) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *OrderStatusObjectOrderStatus) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *OrderStatusObjectOrderStatus) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *OrderStatusObjectOrderStatus) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *OrderStatusObjectOrderStatus) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *OrderStatusObjectOrderStatus) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OrderStatusObjectOrderStatus) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OrderStatusObjectOrderStatus) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OrderStatusObjectOrderStatus) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *OrderStatusObjectOrderStatus) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *OrderStatusObjectOrderStatus) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *OrderStatusObjectOrderStatus) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *OrderStatusObjectOrderStatus) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *OrderStatusObjectOrderStatus) GetRequiredInformation() []SaleObjectSaleRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *OrderStatusObjectOrderStatus) GetRequiredInformationOk() (*[]SaleObjectSaleRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *OrderStatusObjectOrderStatus) SetRequiredInformation(v []SaleObjectSaleRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *OrderStatusObjectOrderStatus) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetInputTokens

`func (o *OrderStatusObjectOrderStatus) GetInputTokens() []SaleObjectSaleInputTokens`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *OrderStatusObjectOrderStatus) GetInputTokensOk() (*[]SaleObjectSaleInputTokens, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *OrderStatusObjectOrderStatus) SetInputTokens(v []SaleObjectSaleInputTokens)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *OrderStatusObjectOrderStatus) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetTicket

`func (o *OrderStatusObjectOrderStatus) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *OrderStatusObjectOrderStatus) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *OrderStatusObjectOrderStatus) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *OrderStatusObjectOrderStatus) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTicketAnswerKey

`func (o *OrderStatusObjectOrderStatus) GetTicketAnswerKey() string`

GetTicketAnswerKey returns the TicketAnswerKey field if non-nil, zero value otherwise.

### GetTicketAnswerKeyOk

`func (o *OrderStatusObjectOrderStatus) GetTicketAnswerKeyOk() (*string, bool)`

GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketAnswerKey

`func (o *OrderStatusObjectOrderStatus) SetTicketAnswerKey(v string)`

SetTicketAnswerKey sets TicketAnswerKey field to given value.

### HasTicketAnswerKey

`func (o *OrderStatusObjectOrderStatus) HasTicketAnswerKey() bool`

HasTicketAnswerKey returns a boolean if a field has been set.

### GetTimeout

`func (o *OrderStatusObjectOrderStatus) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *OrderStatusObjectOrderStatus) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *OrderStatusObjectOrderStatus) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *OrderStatusObjectOrderStatus) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRequestKey

`func (o *OrderStatusObjectOrderStatus) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *OrderStatusObjectOrderStatus) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *OrderStatusObjectOrderStatus) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *OrderStatusObjectOrderStatus) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetMerchantNotifyURL

`func (o *OrderStatusObjectOrderStatus) GetMerchantNotifyURL() string`

GetMerchantNotifyURL returns the MerchantNotifyURL field if non-nil, zero value otherwise.

### GetMerchantNotifyURLOk

`func (o *OrderStatusObjectOrderStatus) GetMerchantNotifyURLOk() (*string, bool)`

GetMerchantNotifyURLOk returns a tuple with the MerchantNotifyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantNotifyURL

`func (o *OrderStatusObjectOrderStatus) SetMerchantNotifyURL(v string)`

SetMerchantNotifyURL sets MerchantNotifyURL field to given value.

### HasMerchantNotifyURL

`func (o *OrderStatusObjectOrderStatus) HasMerchantNotifyURL() bool`

HasMerchantNotifyURL returns a boolean if a field has been set.

### GetIsReverse

`func (o *OrderStatusObjectOrderStatus) GetIsReverse() float32`

GetIsReverse returns the IsReverse field if non-nil, zero value otherwise.

### GetIsReverseOk

`func (o *OrderStatusObjectOrderStatus) GetIsReverseOk() (*float32, bool)`

GetIsReverseOk returns a tuple with the IsReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReverse

`func (o *OrderStatusObjectOrderStatus) SetIsReverse(v float32)`

SetIsReverse sets IsReverse field to given value.

### HasIsReverse

`func (o *OrderStatusObjectOrderStatus) HasIsReverse() bool`

HasIsReverse returns a boolean if a field has been set.

### GetReverseReason

`func (o *OrderStatusObjectOrderStatus) GetReverseReason() string`

GetReverseReason returns the ReverseReason field if non-nil, zero value otherwise.

### GetReverseReasonOk

`func (o *OrderStatusObjectOrderStatus) GetReverseReasonOk() (*string, bool)`

GetReverseReasonOk returns a tuple with the ReverseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseReason

`func (o *OrderStatusObjectOrderStatus) SetReverseReason(v string)`

SetReverseReason sets ReverseReason field to given value.

### HasReverseReason

`func (o *OrderStatusObjectOrderStatus) HasReverseReason() bool`

HasReverseReason returns a boolean if a field has been set.

### GetReasonSequenceBreak

`func (o *OrderStatusObjectOrderStatus) GetReasonSequenceBreak() string`

GetReasonSequenceBreak returns the ReasonSequenceBreak field if non-nil, zero value otherwise.

### GetReasonSequenceBreakOk

`func (o *OrderStatusObjectOrderStatus) GetReasonSequenceBreakOk() (*string, bool)`

GetReasonSequenceBreakOk returns a tuple with the ReasonSequenceBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSequenceBreak

`func (o *OrderStatusObjectOrderStatus) SetReasonSequenceBreak(v string)`

SetReasonSequenceBreak sets ReasonSequenceBreak field to given value.

### HasReasonSequenceBreak

`func (o *OrderStatusObjectOrderStatus) HasReasonSequenceBreak() bool`

HasReasonSequenceBreak returns a boolean if a field has been set.

### GetReference

`func (o *OrderStatusObjectOrderStatus) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *OrderStatusObjectOrderStatus) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *OrderStatusObjectOrderStatus) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *OrderStatusObjectOrderStatus) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *OrderStatusObjectOrderStatus) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *OrderStatusObjectOrderStatus) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *OrderStatusObjectOrderStatus) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *OrderStatusObjectOrderStatus) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetPOSType

`func (o *OrderStatusObjectOrderStatus) GetPOSType() string`

GetPOSType returns the POSType field if non-nil, zero value otherwise.

### GetPOSTypeOk

`func (o *OrderStatusObjectOrderStatus) GetPOSTypeOk() (*string, bool)`

GetPOSTypeOk returns a tuple with the POSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSType

`func (o *OrderStatusObjectOrderStatus) SetPOSType(v string)`

SetPOSType sets POSType field to given value.

### HasPOSType

`func (o *OrderStatusObjectOrderStatus) HasPOSType() bool`

HasPOSType returns a boolean if a field has been set.

### GetPOSVersion

`func (o *OrderStatusObjectOrderStatus) GetPOSVersion() string`

GetPOSVersion returns the POSVersion field if non-nil, zero value otherwise.

### GetPOSVersionOk

`func (o *OrderStatusObjectOrderStatus) GetPOSVersionOk() (*string, bool)`

GetPOSVersionOk returns a tuple with the POSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSVersion

`func (o *OrderStatusObjectOrderStatus) SetPOSVersion(v string)`

SetPOSVersion sets POSVersion field to given value.

### HasPOSVersion

`func (o *OrderStatusObjectOrderStatus) HasPOSVersion() bool`

HasPOSVersion returns a boolean if a field has been set.

### GetPOSAddress

`func (o *OrderStatusObjectOrderStatus) GetPOSAddress() string`

GetPOSAddress returns the POSAddress field if non-nil, zero value otherwise.

### GetPOSAddressOk

`func (o *OrderStatusObjectOrderStatus) GetPOSAddressOk() (*string, bool)`

GetPOSAddressOk returns a tuple with the POSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSAddress

`func (o *OrderStatusObjectOrderStatus) SetPOSAddress(v string)`

SetPOSAddress sets POSAddress field to given value.

### HasPOSAddress

`func (o *OrderStatusObjectOrderStatus) HasPOSAddress() bool`

HasPOSAddress returns a boolean if a field has been set.

### GetPOSSerial

`func (o *OrderStatusObjectOrderStatus) GetPOSSerial() string`

GetPOSSerial returns the POSSerial field if non-nil, zero value otherwise.

### GetPOSSerialOk

`func (o *OrderStatusObjectOrderStatus) GetPOSSerialOk() (*string, bool)`

GetPOSSerialOk returns a tuple with the POSSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSSerial

`func (o *OrderStatusObjectOrderStatus) SetPOSSerial(v string)`

SetPOSSerial sets POSSerial field to given value.

### HasPOSSerial

`func (o *OrderStatusObjectOrderStatus) HasPOSSerial() bool`

HasPOSSerial returns a boolean if a field has been set.

### GetPOSGEO

`func (o *OrderStatusObjectOrderStatus) GetPOSGEO() string`

GetPOSGEO returns the POSGEO field if non-nil, zero value otherwise.

### GetPOSGEOOk

`func (o *OrderStatusObjectOrderStatus) GetPOSGEOOk() (*string, bool)`

GetPOSGEOOk returns a tuple with the POSGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSGEO

`func (o *OrderStatusObjectOrderStatus) SetPOSGEO(v string)`

SetPOSGEO sets POSGEO field to given value.

### HasPOSGEO

`func (o *OrderStatusObjectOrderStatus) HasPOSGEO() bool`

HasPOSGEO returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *OrderStatusObjectOrderStatus) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *OrderStatusObjectOrderStatus) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *OrderStatusObjectOrderStatus) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *OrderStatusObjectOrderStatus) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetReadingDeviceVersion

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceVersion() string`

GetReadingDeviceVersion returns the ReadingDeviceVersion field if non-nil, zero value otherwise.

### GetReadingDeviceVersionOk

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceVersionOk() (*string, bool)`

GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceVersion

`func (o *OrderStatusObjectOrderStatus) SetReadingDeviceVersion(v string)`

SetReadingDeviceVersion sets ReadingDeviceVersion field to given value.

### HasReadingDeviceVersion

`func (o *OrderStatusObjectOrderStatus) HasReadingDeviceVersion() bool`

HasReadingDeviceVersion returns a boolean if a field has been set.

### GetReadingDeviceAddress

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceAddress() string`

GetReadingDeviceAddress returns the ReadingDeviceAddress field if non-nil, zero value otherwise.

### GetReadingDeviceAddressOk

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceAddressOk() (*string, bool)`

GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceAddress

`func (o *OrderStatusObjectOrderStatus) SetReadingDeviceAddress(v string)`

SetReadingDeviceAddress sets ReadingDeviceAddress field to given value.

### HasReadingDeviceAddress

`func (o *OrderStatusObjectOrderStatus) HasReadingDeviceAddress() bool`

HasReadingDeviceAddress returns a boolean if a field has been set.

### GetReadingDeviceSerial

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceSerial() string`

GetReadingDeviceSerial returns the ReadingDeviceSerial field if non-nil, zero value otherwise.

### GetReadingDeviceSerialOk

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceSerialOk() (*string, bool)`

GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceSerial

`func (o *OrderStatusObjectOrderStatus) SetReadingDeviceSerial(v string)`

SetReadingDeviceSerial sets ReadingDeviceSerial field to given value.

### HasReadingDeviceSerial

`func (o *OrderStatusObjectOrderStatus) HasReadingDeviceSerial() bool`

HasReadingDeviceSerial returns a boolean if a field has been set.

### GetReadingDeviceGEO

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceGEO() string`

GetReadingDeviceGEO returns the ReadingDeviceGEO field if non-nil, zero value otherwise.

### GetReadingDeviceGEOOk

`func (o *OrderStatusObjectOrderStatus) GetReadingDeviceGEOOk() (*string, bool)`

GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceGEO

`func (o *OrderStatusObjectOrderStatus) SetReadingDeviceGEO(v string)`

SetReadingDeviceGEO sets ReadingDeviceGEO field to given value.

### HasReadingDeviceGEO

`func (o *OrderStatusObjectOrderStatus) HasReadingDeviceGEO() bool`

HasReadingDeviceGEO returns a boolean if a field has been set.

### GetUserID

`func (o *OrderStatusObjectOrderStatus) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *OrderStatusObjectOrderStatus) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *OrderStatusObjectOrderStatus) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *OrderStatusObjectOrderStatus) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *OrderStatusObjectOrderStatus) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OrderStatusObjectOrderStatus) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OrderStatusObjectOrderStatus) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OrderStatusObjectOrderStatus) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetAmount

`func (o *OrderStatusObjectOrderStatus) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderStatusObjectOrderStatus) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderStatusObjectOrderStatus) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAlternativeAmount

`func (o *OrderStatusObjectOrderStatus) GetAlternativeAmount() float32`

GetAlternativeAmount returns the AlternativeAmount field if non-nil, zero value otherwise.

### GetAlternativeAmountOk

`func (o *OrderStatusObjectOrderStatus) GetAlternativeAmountOk() (*float32, bool)`

GetAlternativeAmountOk returns a tuple with the AlternativeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeAmount

`func (o *OrderStatusObjectOrderStatus) SetAlternativeAmount(v float32)`

SetAlternativeAmount sets AlternativeAmount field to given value.

### HasAlternativeAmount

`func (o *OrderStatusObjectOrderStatus) HasAlternativeAmount() bool`

HasAlternativeAmount returns a boolean if a field has been set.

### GetCashbackAmount

`func (o *OrderStatusObjectOrderStatus) GetCashbackAmount() float32`

GetCashbackAmount returns the CashbackAmount field if non-nil, zero value otherwise.

### GetCashbackAmountOk

`func (o *OrderStatusObjectOrderStatus) GetCashbackAmountOk() (*float32, bool)`

GetCashbackAmountOk returns a tuple with the CashbackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackAmount

`func (o *OrderStatusObjectOrderStatus) SetCashbackAmount(v float32)`

SetCashbackAmount sets CashbackAmount field to given value.

### HasCashbackAmount

`func (o *OrderStatusObjectOrderStatus) HasCashbackAmount() bool`

HasCashbackAmount returns a boolean if a field has been set.

### GetTipAmount

`func (o *OrderStatusObjectOrderStatus) GetTipAmount() float32`

GetTipAmount returns the TipAmount field if non-nil, zero value otherwise.

### GetTipAmountOk

`func (o *OrderStatusObjectOrderStatus) GetTipAmountOk() (*float32, bool)`

GetTipAmountOk returns a tuple with the TipAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipAmount

`func (o *OrderStatusObjectOrderStatus) SetTipAmount(v float32)`

SetTipAmount sets TipAmount field to given value.

### HasTipAmount

`func (o *OrderStatusObjectOrderStatus) HasTipAmount() bool`

HasTipAmount returns a boolean if a field has been set.

### GetPromotedAmount

`func (o *OrderStatusObjectOrderStatus) GetPromotedAmount() float32`

GetPromotedAmount returns the PromotedAmount field if non-nil, zero value otherwise.

### GetPromotedAmountOk

`func (o *OrderStatusObjectOrderStatus) GetPromotedAmountOk() (*float32, bool)`

GetPromotedAmountOk returns a tuple with the PromotedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedAmount

`func (o *OrderStatusObjectOrderStatus) SetPromotedAmount(v float32)`

SetPromotedAmount sets PromotedAmount field to given value.

### HasPromotedAmount

`func (o *OrderStatusObjectOrderStatus) HasPromotedAmount() bool`

HasPromotedAmount returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrderStatusObjectOrderStatus) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderStatusObjectOrderStatus) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderStatusObjectOrderStatus) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrderStatusObjectOrderStatus) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetFacilityPayments

`func (o *OrderStatusObjectOrderStatus) GetFacilityPayments() float32`

GetFacilityPayments returns the FacilityPayments field if non-nil, zero value otherwise.

### GetFacilityPaymentsOk

`func (o *OrderStatusObjectOrderStatus) GetFacilityPaymentsOk() (*float32, bool)`

GetFacilityPaymentsOk returns a tuple with the FacilityPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityPayments

`func (o *OrderStatusObjectOrderStatus) SetFacilityPayments(v float32)`

SetFacilityPayments sets FacilityPayments field to given value.

### HasFacilityPayments

`func (o *OrderStatusObjectOrderStatus) HasFacilityPayments() bool`

HasFacilityPayments returns a boolean if a field has been set.

### GetFacilityType

`func (o *OrderStatusObjectOrderStatus) GetFacilityType() string`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *OrderStatusObjectOrderStatus) GetFacilityTypeOk() (*string, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *OrderStatusObjectOrderStatus) SetFacilityType(v string)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *OrderStatusObjectOrderStatus) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### GetPlan

`func (o *OrderStatusObjectOrderStatus) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OrderStatusObjectOrderStatus) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OrderStatusObjectOrderStatus) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *OrderStatusObjectOrderStatus) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetCardReadMode

`func (o *OrderStatusObjectOrderStatus) GetCardReadMode() string`

GetCardReadMode returns the CardReadMode field if non-nil, zero value otherwise.

### GetCardReadModeOk

`func (o *OrderStatusObjectOrderStatus) GetCardReadModeOk() (*string, bool)`

GetCardReadModeOk returns a tuple with the CardReadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadMode

`func (o *OrderStatusObjectOrderStatus) SetCardReadMode(v string)`

SetCardReadMode sets CardReadMode field to given value.


### GetCardNumber

`func (o *OrderStatusObjectOrderStatus) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *OrderStatusObjectOrderStatus) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *OrderStatusObjectOrderStatus) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *OrderStatusObjectOrderStatus) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNumberMasked

`func (o *OrderStatusObjectOrderStatus) GetCardNumberMasked() string`

GetCardNumberMasked returns the CardNumberMasked field if non-nil, zero value otherwise.

### GetCardNumberMaskedOk

`func (o *OrderStatusObjectOrderStatus) GetCardNumberMaskedOk() (*string, bool)`

GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMasked

`func (o *OrderStatusObjectOrderStatus) SetCardNumberMasked(v string)`

SetCardNumberMasked sets CardNumberMasked field to given value.

### HasCardNumberMasked

`func (o *OrderStatusObjectOrderStatus) HasCardNumberMasked() bool`

HasCardNumberMasked returns a boolean if a field has been set.

### GetCardNumberEncrypted

`func (o *OrderStatusObjectOrderStatus) GetCardNumberEncrypted() string`

GetCardNumberEncrypted returns the CardNumberEncrypted field if non-nil, zero value otherwise.

### GetCardNumberEncryptedOk

`func (o *OrderStatusObjectOrderStatus) GetCardNumberEncryptedOk() (*string, bool)`

GetCardNumberEncryptedOk returns a tuple with the CardNumberEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberEncrypted

`func (o *OrderStatusObjectOrderStatus) SetCardNumberEncrypted(v string)`

SetCardNumberEncrypted sets CardNumberEncrypted field to given value.

### HasCardNumberEncrypted

`func (o *OrderStatusObjectOrderStatus) HasCardNumberEncrypted() bool`

HasCardNumberEncrypted returns a boolean if a field has been set.

### GetCardExp

`func (o *OrderStatusObjectOrderStatus) GetCardExp() string`

GetCardExp returns the CardExp field if non-nil, zero value otherwise.

### GetCardExpOk

`func (o *OrderStatusObjectOrderStatus) GetCardExpOk() (*string, bool)`

GetCardExpOk returns a tuple with the CardExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExp

`func (o *OrderStatusObjectOrderStatus) SetCardExp(v string)`

SetCardExp sets CardExp field to given value.

### HasCardExp

`func (o *OrderStatusObjectOrderStatus) HasCardExp() bool`

HasCardExp returns a boolean if a field has been set.

### GetTrack1

`func (o *OrderStatusObjectOrderStatus) GetTrack1() string`

GetTrack1 returns the Track1 field if non-nil, zero value otherwise.

### GetTrack1Ok

`func (o *OrderStatusObjectOrderStatus) GetTrack1Ok() (*string, bool)`

GetTrack1Ok returns a tuple with the Track1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack1

`func (o *OrderStatusObjectOrderStatus) SetTrack1(v string)`

SetTrack1 sets Track1 field to given value.

### HasTrack1

`func (o *OrderStatusObjectOrderStatus) HasTrack1() bool`

HasTrack1 returns a boolean if a field has been set.

### GetTrack2

`func (o *OrderStatusObjectOrderStatus) GetTrack2() string`

GetTrack2 returns the Track2 field if non-nil, zero value otherwise.

### GetTrack2Ok

`func (o *OrderStatusObjectOrderStatus) GetTrack2Ok() (*string, bool)`

GetTrack2Ok returns a tuple with the Track2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack2

`func (o *OrderStatusObjectOrderStatus) SetTrack2(v string)`

SetTrack2 sets Track2 field to given value.

### HasTrack2

`func (o *OrderStatusObjectOrderStatus) HasTrack2() bool`

HasTrack2 returns a boolean if a field has been set.

### GetSecurityCode

`func (o *OrderStatusObjectOrderStatus) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *OrderStatusObjectOrderStatus) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *OrderStatusObjectOrderStatus) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *OrderStatusObjectOrderStatus) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetPin

`func (o *OrderStatusObjectOrderStatus) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *OrderStatusObjectOrderStatus) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *OrderStatusObjectOrderStatus) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *OrderStatusObjectOrderStatus) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetCardCryptogram

`func (o *OrderStatusObjectOrderStatus) GetCardCryptogram() string`

GetCardCryptogram returns the CardCryptogram field if non-nil, zero value otherwise.

### GetCardCryptogramOk

`func (o *OrderStatusObjectOrderStatus) GetCardCryptogramOk() (*string, bool)`

GetCardCryptogramOk returns a tuple with the CardCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCryptogram

`func (o *OrderStatusObjectOrderStatus) SetCardCryptogram(v string)`

SetCardCryptogram sets CardCryptogram field to given value.

### HasCardCryptogram

`func (o *OrderStatusObjectOrderStatus) HasCardCryptogram() bool`

HasCardCryptogram returns a boolean if a field has been set.

### GetCredentialToken

`func (o *OrderStatusObjectOrderStatus) GetCredentialToken() string`

GetCredentialToken returns the CredentialToken field if non-nil, zero value otherwise.

### GetCredentialTokenOk

`func (o *OrderStatusObjectOrderStatus) GetCredentialTokenOk() (*string, bool)`

GetCredentialTokenOk returns a tuple with the CredentialToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialToken

`func (o *OrderStatusObjectOrderStatus) SetCredentialToken(v string)`

SetCredentialToken sets CredentialToken field to given value.

### HasCredentialToken

`func (o *OrderStatusObjectOrderStatus) HasCredentialToken() bool`

HasCredentialToken returns a boolean if a field has been set.

### GetCredentialIssuerToken

`func (o *OrderStatusObjectOrderStatus) GetCredentialIssuerToken() string`

GetCredentialIssuerToken returns the CredentialIssuerToken field if non-nil, zero value otherwise.

### GetCredentialIssuerTokenOk

`func (o *OrderStatusObjectOrderStatus) GetCredentialIssuerTokenOk() (*string, bool)`

GetCredentialIssuerTokenOk returns a tuple with the CredentialIssuerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuerToken

`func (o *OrderStatusObjectOrderStatus) SetCredentialIssuerToken(v string)`

SetCredentialIssuerToken sets CredentialIssuerToken field to given value.

### HasCredentialIssuerToken

`func (o *OrderStatusObjectOrderStatus) HasCredentialIssuerToken() bool`

HasCredentialIssuerToken returns a boolean if a field has been set.

### GetCardAppName

`func (o *OrderStatusObjectOrderStatus) GetCardAppName() string`

GetCardAppName returns the CardAppName field if non-nil, zero value otherwise.

### GetCardAppNameOk

`func (o *OrderStatusObjectOrderStatus) GetCardAppNameOk() (*string, bool)`

GetCardAppNameOk returns a tuple with the CardAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppName

`func (o *OrderStatusObjectOrderStatus) SetCardAppName(v string)`

SetCardAppName sets CardAppName field to given value.

### HasCardAppName

`func (o *OrderStatusObjectOrderStatus) HasCardAppName() bool`

HasCardAppName returns a boolean if a field has been set.

### GetCardAppIdentifier

`func (o *OrderStatusObjectOrderStatus) GetCardAppIdentifier() string`

GetCardAppIdentifier returns the CardAppIdentifier field if non-nil, zero value otherwise.

### GetCardAppIdentifierOk

`func (o *OrderStatusObjectOrderStatus) GetCardAppIdentifierOk() (*string, bool)`

GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppIdentifier

`func (o *OrderStatusObjectOrderStatus) SetCardAppIdentifier(v string)`

SetCardAppIdentifier sets CardAppIdentifier field to given value.

### HasCardAppIdentifier

`func (o *OrderStatusObjectOrderStatus) HasCardAppIdentifier() bool`

HasCardAppIdentifier returns a boolean if a field has been set.

### GetCardAppLabel

`func (o *OrderStatusObjectOrderStatus) GetCardAppLabel() string`

GetCardAppLabel returns the CardAppLabel field if non-nil, zero value otherwise.

### GetCardAppLabelOk

`func (o *OrderStatusObjectOrderStatus) GetCardAppLabelOk() (*string, bool)`

GetCardAppLabelOk returns a tuple with the CardAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppLabel

`func (o *OrderStatusObjectOrderStatus) SetCardAppLabel(v string)`

SetCardAppLabel sets CardAppLabel field to given value.

### HasCardAppLabel

`func (o *OrderStatusObjectOrderStatus) HasCardAppLabel() bool`

HasCardAppLabel returns a boolean if a field has been set.

### GetCardAuthRequestCryptogram

`func (o *OrderStatusObjectOrderStatus) GetCardAuthRequestCryptogram() string`

GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field if non-nil, zero value otherwise.

### GetCardAuthRequestCryptogramOk

`func (o *OrderStatusObjectOrderStatus) GetCardAuthRequestCryptogramOk() (*string, bool)`

GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthRequestCryptogram

`func (o *OrderStatusObjectOrderStatus) SetCardAuthRequestCryptogram(v string)`

SetCardAuthRequestCryptogram sets CardAuthRequestCryptogram field to given value.

### HasCardAuthRequestCryptogram

`func (o *OrderStatusObjectOrderStatus) HasCardAuthRequestCryptogram() bool`

HasCardAuthRequestCryptogram returns a boolean if a field has been set.

### GetCardAuthResponseCryptogram

`func (o *OrderStatusObjectOrderStatus) GetCardAuthResponseCryptogram() string`

GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field if non-nil, zero value otherwise.

### GetCardAuthResponseCryptogramOk

`func (o *OrderStatusObjectOrderStatus) GetCardAuthResponseCryptogramOk() (*string, bool)`

GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthResponseCryptogram

`func (o *OrderStatusObjectOrderStatus) SetCardAuthResponseCryptogram(v string)`

SetCardAuthResponseCryptogram sets CardAuthResponseCryptogram field to given value.

### HasCardAuthResponseCryptogram

`func (o *OrderStatusObjectOrderStatus) HasCardAuthResponseCryptogram() bool`

HasCardAuthResponseCryptogram returns a boolean if a field has been set.

### GetPayer

`func (o *OrderStatusObjectOrderStatus) GetPayer() SaleObjectSalePayer`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *OrderStatusObjectOrderStatus) GetPayerOk() (*SaleObjectSalePayer, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *OrderStatusObjectOrderStatus) SetPayer(v SaleObjectSalePayer)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *OrderStatusObjectOrderStatus) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderStatusObjectOrderStatus) GetCustomer() SaleObjectSaleCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderStatusObjectOrderStatus) GetCustomerOk() (*SaleObjectSaleCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderStatusObjectOrderStatus) SetCustomer(v SaleObjectSaleCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderStatusObjectOrderStatus) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSeller

`func (o *OrderStatusObjectOrderStatus) GetSeller() SaleObjectSaleSeller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *OrderStatusObjectOrderStatus) GetSellerOk() (*SaleObjectSaleSeller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *OrderStatusObjectOrderStatus) SetSeller(v SaleObjectSaleSeller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *OrderStatusObjectOrderStatus) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetProducts

`func (o *OrderStatusObjectOrderStatus) GetProducts() []SaleObjectSaleProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *OrderStatusObjectOrderStatus) GetProductsOk() (*[]SaleObjectSaleProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *OrderStatusObjectOrderStatus) SetProducts(v []SaleObjectSaleProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *OrderStatusObjectOrderStatus) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetTaxRefundType

`func (o *OrderStatusObjectOrderStatus) GetTaxRefundType() string`

GetTaxRefundType returns the TaxRefundType field if non-nil, zero value otherwise.

### GetTaxRefundTypeOk

`func (o *OrderStatusObjectOrderStatus) GetTaxRefundTypeOk() (*string, bool)`

GetTaxRefundTypeOk returns a tuple with the TaxRefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRefundType

`func (o *OrderStatusObjectOrderStatus) SetTaxRefundType(v string)`

SetTaxRefundType sets TaxRefundType field to given value.

### HasTaxRefundType

`func (o *OrderStatusObjectOrderStatus) HasTaxRefundType() bool`

HasTaxRefundType returns a boolean if a field has been set.

### GetAuthCode

`func (o *OrderStatusObjectOrderStatus) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *OrderStatusObjectOrderStatus) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *OrderStatusObjectOrderStatus) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *OrderStatusObjectOrderStatus) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


