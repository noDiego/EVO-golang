# SettlementObjectSettlement

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
**Ticket** | Pointer to **string** | Ticket Digitalizado ( Total o parte del mismo por ejemplo la Firma Digitalizada )    codificado en Base64. | [optional] 
**TicketAnswerKey** | Pointer to **string** | Identificador Unívoco de la Transacción que se quiere Referenciar a la cual pertenece el Ticket Digitalizado. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. Si la firma fue capturada previamente y se envía en el requerimiento de la misma Operación Sale, Authorize*, Void, Return, Adjustment, DebtPayment, VoidDebtPayment o Enrollment no es necesario que se envíe este elemento o campo. | [optional] 
**Timeout** | Pointer to **float32** | Tiempo de espera que el POS espera al PINPAD para obtener la respuesta al requerimiento.  | [optional] 
**MerchantNotifyURL** | Pointer to **string** | URL para notificación del comercio | [optional] 
**IsReverse** | Pointer to **float32** | Identificador de Reversa. | [optional] 
**ReverseReason** | Pointer to **string** | Motivo por el cual se requiere generar la reversa. | [optional] 
**ReasonSequenceBreak** | Pointer to **string** | Motivo por el cual se requiere romper la secuencia. | [optional] 
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
**Amount** | Pointer to **float32** | Importe o Monto de la Transacción. | [optional] 
**AmountCharged** | Pointer to **float32** | Importe o Monto de la Transacción que efectivamente se cobro , si se envía en Void o Return en lugar de Amount, se genera un Ajuste si el Host lo soporta. | [optional] 
**CardReadMode** | Pointer to **string** | Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes ) | [optional] 
**CardGetMode** | Pointer to **string** | Indican por cada elemento que contiene los datos sensibles, si están encriptados  y también el algoritmo usado. En caso de no estar especificado se asume PLAIN. | [optional] 
**CardNumber** | Pointer to **string** | Número de Tarjeta, en el caso de las respuestas el mismo estará enmascarado. | [optional] 
**CardNumberMasked** | Pointer to **string** | Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón.     | [optional] 
**CardNumberEncrypted** | Pointer to **string** | Número de tarjeta encriptado.                       | [optional] 
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
**OrigAnswerKey** | Pointer to **string** | Identificador Unívoco de la Transacción que se quiere Referenciar, usado en Deposit, Void, Return, etc. O sea en las transacciones que hacen referencia a una Transacción previa. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada.  | [optional] 
**OrigBlock** | Pointer to **string** | Identificador de Bloque que se quiere recuperar la información de su composición con la operación GetBlock, Void, Return. | [optional] 
**OrigForeignBlock** | Pointer to **string** | Identificador de Bloque que se quiere recuperar la información de su composición con la operación GetBlock, Void, Return. | [optional] 
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relacionado a cada uno de los planes disponibles. | [optional] 
**TerminalID** | Pointer to **string** | Identificador de Terminal por el cual se envía la Transacción al Host. | [optional] 
**TerminalTrace** | Pointer to **int32** | Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID. | [optional] 
**SettlementBatchNumber** | Pointer to **int32** | Para aquellos host que exista el concepto de lote, es el número de lote al cual pertenece la transacción. | [optional] 

## Methods

### NewSettlementObjectSettlement

`func NewSettlementObjectSettlement(companyIdentification string, systemIdentification string, ) *SettlementObjectSettlement`

NewSettlementObjectSettlement instantiates a new SettlementObjectSettlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementObjectSettlementWithDefaults

`func NewSettlementObjectSettlementWithDefaults() *SettlementObjectSettlement`

NewSettlementObjectSettlementWithDefaults instantiates a new SettlementObjectSettlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *SettlementObjectSettlement) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *SettlementObjectSettlement) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *SettlementObjectSettlement) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetSystemIdentification

`func (o *SettlementObjectSettlement) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *SettlementObjectSettlement) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *SettlementObjectSettlement) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetBranchIdentification

`func (o *SettlementObjectSettlement) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *SettlementObjectSettlement) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *SettlementObjectSettlement) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *SettlementObjectSettlement) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *SettlementObjectSettlement) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *SettlementObjectSettlement) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *SettlementObjectSettlement) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *SettlementObjectSettlement) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetRequestType

`func (o *SettlementObjectSettlement) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SettlementObjectSettlement) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SettlementObjectSettlement) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *SettlementObjectSettlement) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *SettlementObjectSettlement) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *SettlementObjectSettlement) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *SettlementObjectSettlement) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *SettlementObjectSettlement) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *SettlementObjectSettlement) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *SettlementObjectSettlement) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *SettlementObjectSettlement) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *SettlementObjectSettlement) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *SettlementObjectSettlement) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SettlementObjectSettlement) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SettlementObjectSettlement) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SettlementObjectSettlement) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *SettlementObjectSettlement) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *SettlementObjectSettlement) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *SettlementObjectSettlement) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *SettlementObjectSettlement) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *SettlementObjectSettlement) GetRequiredInformation() []SaleObjectSaleRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *SettlementObjectSettlement) GetRequiredInformationOk() (*[]SaleObjectSaleRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *SettlementObjectSettlement) SetRequiredInformation(v []SaleObjectSaleRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *SettlementObjectSettlement) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetTicket

`func (o *SettlementObjectSettlement) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *SettlementObjectSettlement) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *SettlementObjectSettlement) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *SettlementObjectSettlement) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTicketAnswerKey

`func (o *SettlementObjectSettlement) GetTicketAnswerKey() string`

GetTicketAnswerKey returns the TicketAnswerKey field if non-nil, zero value otherwise.

### GetTicketAnswerKeyOk

`func (o *SettlementObjectSettlement) GetTicketAnswerKeyOk() (*string, bool)`

GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketAnswerKey

`func (o *SettlementObjectSettlement) SetTicketAnswerKey(v string)`

SetTicketAnswerKey sets TicketAnswerKey field to given value.

### HasTicketAnswerKey

`func (o *SettlementObjectSettlement) HasTicketAnswerKey() bool`

HasTicketAnswerKey returns a boolean if a field has been set.

### GetTimeout

`func (o *SettlementObjectSettlement) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SettlementObjectSettlement) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SettlementObjectSettlement) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SettlementObjectSettlement) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetMerchantNotifyURL

`func (o *SettlementObjectSettlement) GetMerchantNotifyURL() string`

GetMerchantNotifyURL returns the MerchantNotifyURL field if non-nil, zero value otherwise.

### GetMerchantNotifyURLOk

`func (o *SettlementObjectSettlement) GetMerchantNotifyURLOk() (*string, bool)`

GetMerchantNotifyURLOk returns a tuple with the MerchantNotifyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantNotifyURL

`func (o *SettlementObjectSettlement) SetMerchantNotifyURL(v string)`

SetMerchantNotifyURL sets MerchantNotifyURL field to given value.

### HasMerchantNotifyURL

`func (o *SettlementObjectSettlement) HasMerchantNotifyURL() bool`

HasMerchantNotifyURL returns a boolean if a field has been set.

### GetIsReverse

`func (o *SettlementObjectSettlement) GetIsReverse() float32`

GetIsReverse returns the IsReverse field if non-nil, zero value otherwise.

### GetIsReverseOk

`func (o *SettlementObjectSettlement) GetIsReverseOk() (*float32, bool)`

GetIsReverseOk returns a tuple with the IsReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReverse

`func (o *SettlementObjectSettlement) SetIsReverse(v float32)`

SetIsReverse sets IsReverse field to given value.

### HasIsReverse

`func (o *SettlementObjectSettlement) HasIsReverse() bool`

HasIsReverse returns a boolean if a field has been set.

### GetReverseReason

`func (o *SettlementObjectSettlement) GetReverseReason() string`

GetReverseReason returns the ReverseReason field if non-nil, zero value otherwise.

### GetReverseReasonOk

`func (o *SettlementObjectSettlement) GetReverseReasonOk() (*string, bool)`

GetReverseReasonOk returns a tuple with the ReverseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseReason

`func (o *SettlementObjectSettlement) SetReverseReason(v string)`

SetReverseReason sets ReverseReason field to given value.

### HasReverseReason

`func (o *SettlementObjectSettlement) HasReverseReason() bool`

HasReverseReason returns a boolean if a field has been set.

### GetReasonSequenceBreak

`func (o *SettlementObjectSettlement) GetReasonSequenceBreak() string`

GetReasonSequenceBreak returns the ReasonSequenceBreak field if non-nil, zero value otherwise.

### GetReasonSequenceBreakOk

`func (o *SettlementObjectSettlement) GetReasonSequenceBreakOk() (*string, bool)`

GetReasonSequenceBreakOk returns a tuple with the ReasonSequenceBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSequenceBreak

`func (o *SettlementObjectSettlement) SetReasonSequenceBreak(v string)`

SetReasonSequenceBreak sets ReasonSequenceBreak field to given value.

### HasReasonSequenceBreak

`func (o *SettlementObjectSettlement) HasReasonSequenceBreak() bool`

HasReasonSequenceBreak returns a boolean if a field has been set.

### GetPOSType

`func (o *SettlementObjectSettlement) GetPOSType() string`

GetPOSType returns the POSType field if non-nil, zero value otherwise.

### GetPOSTypeOk

`func (o *SettlementObjectSettlement) GetPOSTypeOk() (*string, bool)`

GetPOSTypeOk returns a tuple with the POSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSType

`func (o *SettlementObjectSettlement) SetPOSType(v string)`

SetPOSType sets POSType field to given value.

### HasPOSType

`func (o *SettlementObjectSettlement) HasPOSType() bool`

HasPOSType returns a boolean if a field has been set.

### GetPOSVersion

`func (o *SettlementObjectSettlement) GetPOSVersion() string`

GetPOSVersion returns the POSVersion field if non-nil, zero value otherwise.

### GetPOSVersionOk

`func (o *SettlementObjectSettlement) GetPOSVersionOk() (*string, bool)`

GetPOSVersionOk returns a tuple with the POSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSVersion

`func (o *SettlementObjectSettlement) SetPOSVersion(v string)`

SetPOSVersion sets POSVersion field to given value.

### HasPOSVersion

`func (o *SettlementObjectSettlement) HasPOSVersion() bool`

HasPOSVersion returns a boolean if a field has been set.

### GetPOSAddress

`func (o *SettlementObjectSettlement) GetPOSAddress() string`

GetPOSAddress returns the POSAddress field if non-nil, zero value otherwise.

### GetPOSAddressOk

`func (o *SettlementObjectSettlement) GetPOSAddressOk() (*string, bool)`

GetPOSAddressOk returns a tuple with the POSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSAddress

`func (o *SettlementObjectSettlement) SetPOSAddress(v string)`

SetPOSAddress sets POSAddress field to given value.

### HasPOSAddress

`func (o *SettlementObjectSettlement) HasPOSAddress() bool`

HasPOSAddress returns a boolean if a field has been set.

### GetPOSSerial

`func (o *SettlementObjectSettlement) GetPOSSerial() string`

GetPOSSerial returns the POSSerial field if non-nil, zero value otherwise.

### GetPOSSerialOk

`func (o *SettlementObjectSettlement) GetPOSSerialOk() (*string, bool)`

GetPOSSerialOk returns a tuple with the POSSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSSerial

`func (o *SettlementObjectSettlement) SetPOSSerial(v string)`

SetPOSSerial sets POSSerial field to given value.

### HasPOSSerial

`func (o *SettlementObjectSettlement) HasPOSSerial() bool`

HasPOSSerial returns a boolean if a field has been set.

### GetPOSGEO

`func (o *SettlementObjectSettlement) GetPOSGEO() string`

GetPOSGEO returns the POSGEO field if non-nil, zero value otherwise.

### GetPOSGEOOk

`func (o *SettlementObjectSettlement) GetPOSGEOOk() (*string, bool)`

GetPOSGEOOk returns a tuple with the POSGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSGEO

`func (o *SettlementObjectSettlement) SetPOSGEO(v string)`

SetPOSGEO sets POSGEO field to given value.

### HasPOSGEO

`func (o *SettlementObjectSettlement) HasPOSGEO() bool`

HasPOSGEO returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *SettlementObjectSettlement) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *SettlementObjectSettlement) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *SettlementObjectSettlement) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *SettlementObjectSettlement) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *SettlementObjectSettlement) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *SettlementObjectSettlement) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *SettlementObjectSettlement) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *SettlementObjectSettlement) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetReadingDeviceVersion

`func (o *SettlementObjectSettlement) GetReadingDeviceVersion() string`

GetReadingDeviceVersion returns the ReadingDeviceVersion field if non-nil, zero value otherwise.

### GetReadingDeviceVersionOk

`func (o *SettlementObjectSettlement) GetReadingDeviceVersionOk() (*string, bool)`

GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceVersion

`func (o *SettlementObjectSettlement) SetReadingDeviceVersion(v string)`

SetReadingDeviceVersion sets ReadingDeviceVersion field to given value.

### HasReadingDeviceVersion

`func (o *SettlementObjectSettlement) HasReadingDeviceVersion() bool`

HasReadingDeviceVersion returns a boolean if a field has been set.

### GetReadingDeviceAddress

`func (o *SettlementObjectSettlement) GetReadingDeviceAddress() string`

GetReadingDeviceAddress returns the ReadingDeviceAddress field if non-nil, zero value otherwise.

### GetReadingDeviceAddressOk

`func (o *SettlementObjectSettlement) GetReadingDeviceAddressOk() (*string, bool)`

GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceAddress

`func (o *SettlementObjectSettlement) SetReadingDeviceAddress(v string)`

SetReadingDeviceAddress sets ReadingDeviceAddress field to given value.

### HasReadingDeviceAddress

`func (o *SettlementObjectSettlement) HasReadingDeviceAddress() bool`

HasReadingDeviceAddress returns a boolean if a field has been set.

### GetReadingDeviceSerial

`func (o *SettlementObjectSettlement) GetReadingDeviceSerial() string`

GetReadingDeviceSerial returns the ReadingDeviceSerial field if non-nil, zero value otherwise.

### GetReadingDeviceSerialOk

`func (o *SettlementObjectSettlement) GetReadingDeviceSerialOk() (*string, bool)`

GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceSerial

`func (o *SettlementObjectSettlement) SetReadingDeviceSerial(v string)`

SetReadingDeviceSerial sets ReadingDeviceSerial field to given value.

### HasReadingDeviceSerial

`func (o *SettlementObjectSettlement) HasReadingDeviceSerial() bool`

HasReadingDeviceSerial returns a boolean if a field has been set.

### GetReadingDeviceGEO

`func (o *SettlementObjectSettlement) GetReadingDeviceGEO() string`

GetReadingDeviceGEO returns the ReadingDeviceGEO field if non-nil, zero value otherwise.

### GetReadingDeviceGEOOk

`func (o *SettlementObjectSettlement) GetReadingDeviceGEOOk() (*string, bool)`

GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceGEO

`func (o *SettlementObjectSettlement) SetReadingDeviceGEO(v string)`

SetReadingDeviceGEO sets ReadingDeviceGEO field to given value.

### HasReadingDeviceGEO

`func (o *SettlementObjectSettlement) HasReadingDeviceGEO() bool`

HasReadingDeviceGEO returns a boolean if a field has been set.

### GetUserID

`func (o *SettlementObjectSettlement) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *SettlementObjectSettlement) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *SettlementObjectSettlement) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *SettlementObjectSettlement) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *SettlementObjectSettlement) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SettlementObjectSettlement) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SettlementObjectSettlement) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SettlementObjectSettlement) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetAmount

`func (o *SettlementObjectSettlement) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SettlementObjectSettlement) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SettlementObjectSettlement) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SettlementObjectSettlement) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountCharged

`func (o *SettlementObjectSettlement) GetAmountCharged() float32`

GetAmountCharged returns the AmountCharged field if non-nil, zero value otherwise.

### GetAmountChargedOk

`func (o *SettlementObjectSettlement) GetAmountChargedOk() (*float32, bool)`

GetAmountChargedOk returns a tuple with the AmountCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCharged

`func (o *SettlementObjectSettlement) SetAmountCharged(v float32)`

SetAmountCharged sets AmountCharged field to given value.

### HasAmountCharged

`func (o *SettlementObjectSettlement) HasAmountCharged() bool`

HasAmountCharged returns a boolean if a field has been set.

### GetCardReadMode

`func (o *SettlementObjectSettlement) GetCardReadMode() string`

GetCardReadMode returns the CardReadMode field if non-nil, zero value otherwise.

### GetCardReadModeOk

`func (o *SettlementObjectSettlement) GetCardReadModeOk() (*string, bool)`

GetCardReadModeOk returns a tuple with the CardReadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadMode

`func (o *SettlementObjectSettlement) SetCardReadMode(v string)`

SetCardReadMode sets CardReadMode field to given value.

### HasCardReadMode

`func (o *SettlementObjectSettlement) HasCardReadMode() bool`

HasCardReadMode returns a boolean if a field has been set.

### GetCardGetMode

`func (o *SettlementObjectSettlement) GetCardGetMode() string`

GetCardGetMode returns the CardGetMode field if non-nil, zero value otherwise.

### GetCardGetModeOk

`func (o *SettlementObjectSettlement) GetCardGetModeOk() (*string, bool)`

GetCardGetModeOk returns a tuple with the CardGetMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardGetMode

`func (o *SettlementObjectSettlement) SetCardGetMode(v string)`

SetCardGetMode sets CardGetMode field to given value.

### HasCardGetMode

`func (o *SettlementObjectSettlement) HasCardGetMode() bool`

HasCardGetMode returns a boolean if a field has been set.

### GetCardNumber

`func (o *SettlementObjectSettlement) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *SettlementObjectSettlement) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *SettlementObjectSettlement) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *SettlementObjectSettlement) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNumberMasked

`func (o *SettlementObjectSettlement) GetCardNumberMasked() string`

GetCardNumberMasked returns the CardNumberMasked field if non-nil, zero value otherwise.

### GetCardNumberMaskedOk

`func (o *SettlementObjectSettlement) GetCardNumberMaskedOk() (*string, bool)`

GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMasked

`func (o *SettlementObjectSettlement) SetCardNumberMasked(v string)`

SetCardNumberMasked sets CardNumberMasked field to given value.

### HasCardNumberMasked

`func (o *SettlementObjectSettlement) HasCardNumberMasked() bool`

HasCardNumberMasked returns a boolean if a field has been set.

### GetCardNumberEncrypted

`func (o *SettlementObjectSettlement) GetCardNumberEncrypted() string`

GetCardNumberEncrypted returns the CardNumberEncrypted field if non-nil, zero value otherwise.

### GetCardNumberEncryptedOk

`func (o *SettlementObjectSettlement) GetCardNumberEncryptedOk() (*string, bool)`

GetCardNumberEncryptedOk returns a tuple with the CardNumberEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberEncrypted

`func (o *SettlementObjectSettlement) SetCardNumberEncrypted(v string)`

SetCardNumberEncrypted sets CardNumberEncrypted field to given value.

### HasCardNumberEncrypted

`func (o *SettlementObjectSettlement) HasCardNumberEncrypted() bool`

HasCardNumberEncrypted returns a boolean if a field has been set.

### GetCardExp

`func (o *SettlementObjectSettlement) GetCardExp() string`

GetCardExp returns the CardExp field if non-nil, zero value otherwise.

### GetCardExpOk

`func (o *SettlementObjectSettlement) GetCardExpOk() (*string, bool)`

GetCardExpOk returns a tuple with the CardExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExp

`func (o *SettlementObjectSettlement) SetCardExp(v string)`

SetCardExp sets CardExp field to given value.

### HasCardExp

`func (o *SettlementObjectSettlement) HasCardExp() bool`

HasCardExp returns a boolean if a field has been set.

### GetTrack1

`func (o *SettlementObjectSettlement) GetTrack1() string`

GetTrack1 returns the Track1 field if non-nil, zero value otherwise.

### GetTrack1Ok

`func (o *SettlementObjectSettlement) GetTrack1Ok() (*string, bool)`

GetTrack1Ok returns a tuple with the Track1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack1

`func (o *SettlementObjectSettlement) SetTrack1(v string)`

SetTrack1 sets Track1 field to given value.

### HasTrack1

`func (o *SettlementObjectSettlement) HasTrack1() bool`

HasTrack1 returns a boolean if a field has been set.

### GetTrack2

`func (o *SettlementObjectSettlement) GetTrack2() string`

GetTrack2 returns the Track2 field if non-nil, zero value otherwise.

### GetTrack2Ok

`func (o *SettlementObjectSettlement) GetTrack2Ok() (*string, bool)`

GetTrack2Ok returns a tuple with the Track2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack2

`func (o *SettlementObjectSettlement) SetTrack2(v string)`

SetTrack2 sets Track2 field to given value.

### HasTrack2

`func (o *SettlementObjectSettlement) HasTrack2() bool`

HasTrack2 returns a boolean if a field has been set.

### GetSecurityCode

`func (o *SettlementObjectSettlement) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *SettlementObjectSettlement) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *SettlementObjectSettlement) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *SettlementObjectSettlement) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetPin

`func (o *SettlementObjectSettlement) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *SettlementObjectSettlement) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *SettlementObjectSettlement) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *SettlementObjectSettlement) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetCardCryptogram

`func (o *SettlementObjectSettlement) GetCardCryptogram() string`

GetCardCryptogram returns the CardCryptogram field if non-nil, zero value otherwise.

### GetCardCryptogramOk

`func (o *SettlementObjectSettlement) GetCardCryptogramOk() (*string, bool)`

GetCardCryptogramOk returns a tuple with the CardCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCryptogram

`func (o *SettlementObjectSettlement) SetCardCryptogram(v string)`

SetCardCryptogram sets CardCryptogram field to given value.

### HasCardCryptogram

`func (o *SettlementObjectSettlement) HasCardCryptogram() bool`

HasCardCryptogram returns a boolean if a field has been set.

### GetCredentialToken

`func (o *SettlementObjectSettlement) GetCredentialToken() string`

GetCredentialToken returns the CredentialToken field if non-nil, zero value otherwise.

### GetCredentialTokenOk

`func (o *SettlementObjectSettlement) GetCredentialTokenOk() (*string, bool)`

GetCredentialTokenOk returns a tuple with the CredentialToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialToken

`func (o *SettlementObjectSettlement) SetCredentialToken(v string)`

SetCredentialToken sets CredentialToken field to given value.

### HasCredentialToken

`func (o *SettlementObjectSettlement) HasCredentialToken() bool`

HasCredentialToken returns a boolean if a field has been set.

### GetCredentialIssuerToken

`func (o *SettlementObjectSettlement) GetCredentialIssuerToken() string`

GetCredentialIssuerToken returns the CredentialIssuerToken field if non-nil, zero value otherwise.

### GetCredentialIssuerTokenOk

`func (o *SettlementObjectSettlement) GetCredentialIssuerTokenOk() (*string, bool)`

GetCredentialIssuerTokenOk returns a tuple with the CredentialIssuerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuerToken

`func (o *SettlementObjectSettlement) SetCredentialIssuerToken(v string)`

SetCredentialIssuerToken sets CredentialIssuerToken field to given value.

### HasCredentialIssuerToken

`func (o *SettlementObjectSettlement) HasCredentialIssuerToken() bool`

HasCredentialIssuerToken returns a boolean if a field has been set.

### GetCardAppName

`func (o *SettlementObjectSettlement) GetCardAppName() string`

GetCardAppName returns the CardAppName field if non-nil, zero value otherwise.

### GetCardAppNameOk

`func (o *SettlementObjectSettlement) GetCardAppNameOk() (*string, bool)`

GetCardAppNameOk returns a tuple with the CardAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppName

`func (o *SettlementObjectSettlement) SetCardAppName(v string)`

SetCardAppName sets CardAppName field to given value.

### HasCardAppName

`func (o *SettlementObjectSettlement) HasCardAppName() bool`

HasCardAppName returns a boolean if a field has been set.

### GetCardAppIdentifier

`func (o *SettlementObjectSettlement) GetCardAppIdentifier() string`

GetCardAppIdentifier returns the CardAppIdentifier field if non-nil, zero value otherwise.

### GetCardAppIdentifierOk

`func (o *SettlementObjectSettlement) GetCardAppIdentifierOk() (*string, bool)`

GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppIdentifier

`func (o *SettlementObjectSettlement) SetCardAppIdentifier(v string)`

SetCardAppIdentifier sets CardAppIdentifier field to given value.

### HasCardAppIdentifier

`func (o *SettlementObjectSettlement) HasCardAppIdentifier() bool`

HasCardAppIdentifier returns a boolean if a field has been set.

### GetCardAppLabel

`func (o *SettlementObjectSettlement) GetCardAppLabel() string`

GetCardAppLabel returns the CardAppLabel field if non-nil, zero value otherwise.

### GetCardAppLabelOk

`func (o *SettlementObjectSettlement) GetCardAppLabelOk() (*string, bool)`

GetCardAppLabelOk returns a tuple with the CardAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppLabel

`func (o *SettlementObjectSettlement) SetCardAppLabel(v string)`

SetCardAppLabel sets CardAppLabel field to given value.

### HasCardAppLabel

`func (o *SettlementObjectSettlement) HasCardAppLabel() bool`

HasCardAppLabel returns a boolean if a field has been set.

### GetCardAuthRequestCryptogram

`func (o *SettlementObjectSettlement) GetCardAuthRequestCryptogram() string`

GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field if non-nil, zero value otherwise.

### GetCardAuthRequestCryptogramOk

`func (o *SettlementObjectSettlement) GetCardAuthRequestCryptogramOk() (*string, bool)`

GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthRequestCryptogram

`func (o *SettlementObjectSettlement) SetCardAuthRequestCryptogram(v string)`

SetCardAuthRequestCryptogram sets CardAuthRequestCryptogram field to given value.

### HasCardAuthRequestCryptogram

`func (o *SettlementObjectSettlement) HasCardAuthRequestCryptogram() bool`

HasCardAuthRequestCryptogram returns a boolean if a field has been set.

### GetCardAuthResponseCryptogram

`func (o *SettlementObjectSettlement) GetCardAuthResponseCryptogram() string`

GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field if non-nil, zero value otherwise.

### GetCardAuthResponseCryptogramOk

`func (o *SettlementObjectSettlement) GetCardAuthResponseCryptogramOk() (*string, bool)`

GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthResponseCryptogram

`func (o *SettlementObjectSettlement) SetCardAuthResponseCryptogram(v string)`

SetCardAuthResponseCryptogram sets CardAuthResponseCryptogram field to given value.

### HasCardAuthResponseCryptogram

`func (o *SettlementObjectSettlement) HasCardAuthResponseCryptogram() bool`

HasCardAuthResponseCryptogram returns a boolean if a field has been set.

### GetPayer

`func (o *SettlementObjectSettlement) GetPayer() SaleObjectSalePayer`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *SettlementObjectSettlement) GetPayerOk() (*SaleObjectSalePayer, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *SettlementObjectSettlement) SetPayer(v SaleObjectSalePayer)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *SettlementObjectSettlement) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetCustomer

`func (o *SettlementObjectSettlement) GetCustomer() SaleObjectSaleCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *SettlementObjectSettlement) GetCustomerOk() (*SaleObjectSaleCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *SettlementObjectSettlement) SetCustomer(v SaleObjectSaleCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *SettlementObjectSettlement) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSeller

`func (o *SettlementObjectSettlement) GetSeller() SaleObjectSaleSeller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *SettlementObjectSettlement) GetSellerOk() (*SaleObjectSaleSeller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *SettlementObjectSettlement) SetSeller(v SaleObjectSaleSeller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *SettlementObjectSettlement) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetProducts

`func (o *SettlementObjectSettlement) GetProducts() []SaleObjectSaleProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *SettlementObjectSettlement) GetProductsOk() (*[]SaleObjectSaleProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *SettlementObjectSettlement) SetProducts(v []SaleObjectSaleProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *SettlementObjectSettlement) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetTaxRefundType

`func (o *SettlementObjectSettlement) GetTaxRefundType() string`

GetTaxRefundType returns the TaxRefundType field if non-nil, zero value otherwise.

### GetTaxRefundTypeOk

`func (o *SettlementObjectSettlement) GetTaxRefundTypeOk() (*string, bool)`

GetTaxRefundTypeOk returns a tuple with the TaxRefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRefundType

`func (o *SettlementObjectSettlement) SetTaxRefundType(v string)`

SetTaxRefundType sets TaxRefundType field to given value.

### HasTaxRefundType

`func (o *SettlementObjectSettlement) HasTaxRefundType() bool`

HasTaxRefundType returns a boolean if a field has been set.

### GetOrigAnswerKey

`func (o *SettlementObjectSettlement) GetOrigAnswerKey() string`

GetOrigAnswerKey returns the OrigAnswerKey field if non-nil, zero value otherwise.

### GetOrigAnswerKeyOk

`func (o *SettlementObjectSettlement) GetOrigAnswerKeyOk() (*string, bool)`

GetOrigAnswerKeyOk returns a tuple with the OrigAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigAnswerKey

`func (o *SettlementObjectSettlement) SetOrigAnswerKey(v string)`

SetOrigAnswerKey sets OrigAnswerKey field to given value.

### HasOrigAnswerKey

`func (o *SettlementObjectSettlement) HasOrigAnswerKey() bool`

HasOrigAnswerKey returns a boolean if a field has been set.

### GetOrigBlock

`func (o *SettlementObjectSettlement) GetOrigBlock() string`

GetOrigBlock returns the OrigBlock field if non-nil, zero value otherwise.

### GetOrigBlockOk

`func (o *SettlementObjectSettlement) GetOrigBlockOk() (*string, bool)`

GetOrigBlockOk returns a tuple with the OrigBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigBlock

`func (o *SettlementObjectSettlement) SetOrigBlock(v string)`

SetOrigBlock sets OrigBlock field to given value.

### HasOrigBlock

`func (o *SettlementObjectSettlement) HasOrigBlock() bool`

HasOrigBlock returns a boolean if a field has been set.

### GetOrigForeignBlock

`func (o *SettlementObjectSettlement) GetOrigForeignBlock() string`

GetOrigForeignBlock returns the OrigForeignBlock field if non-nil, zero value otherwise.

### GetOrigForeignBlockOk

`func (o *SettlementObjectSettlement) GetOrigForeignBlockOk() (*string, bool)`

GetOrigForeignBlockOk returns a tuple with the OrigForeignBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigForeignBlock

`func (o *SettlementObjectSettlement) SetOrigForeignBlock(v string)`

SetOrigForeignBlock sets OrigForeignBlock field to given value.

### HasOrigForeignBlock

`func (o *SettlementObjectSettlement) HasOrigForeignBlock() bool`

HasOrigForeignBlock returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *SettlementObjectSettlement) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *SettlementObjectSettlement) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *SettlementObjectSettlement) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *SettlementObjectSettlement) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *SettlementObjectSettlement) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *SettlementObjectSettlement) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *SettlementObjectSettlement) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *SettlementObjectSettlement) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetTerminalID

`func (o *SettlementObjectSettlement) GetTerminalID() string`

GetTerminalID returns the TerminalID field if non-nil, zero value otherwise.

### GetTerminalIDOk

`func (o *SettlementObjectSettlement) GetTerminalIDOk() (*string, bool)`

GetTerminalIDOk returns a tuple with the TerminalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalID

`func (o *SettlementObjectSettlement) SetTerminalID(v string)`

SetTerminalID sets TerminalID field to given value.

### HasTerminalID

`func (o *SettlementObjectSettlement) HasTerminalID() bool`

HasTerminalID returns a boolean if a field has been set.

### GetTerminalTrace

`func (o *SettlementObjectSettlement) GetTerminalTrace() int32`

GetTerminalTrace returns the TerminalTrace field if non-nil, zero value otherwise.

### GetTerminalTraceOk

`func (o *SettlementObjectSettlement) GetTerminalTraceOk() (*int32, bool)`

GetTerminalTraceOk returns a tuple with the TerminalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTrace

`func (o *SettlementObjectSettlement) SetTerminalTrace(v int32)`

SetTerminalTrace sets TerminalTrace field to given value.

### HasTerminalTrace

`func (o *SettlementObjectSettlement) HasTerminalTrace() bool`

HasTerminalTrace returns a boolean if a field has been set.

### GetSettlementBatchNumber

`func (o *SettlementObjectSettlement) GetSettlementBatchNumber() int32`

GetSettlementBatchNumber returns the SettlementBatchNumber field if non-nil, zero value otherwise.

### GetSettlementBatchNumberOk

`func (o *SettlementObjectSettlement) GetSettlementBatchNumberOk() (*int32, bool)`

GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBatchNumber

`func (o *SettlementObjectSettlement) SetSettlementBatchNumber(v int32)`

SetSettlementBatchNumber sets SettlementBatchNumber field to given value.

### HasSettlementBatchNumber

`func (o *SettlementObjectSettlement) HasSettlementBatchNumber() bool`

HasSettlementBatchNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


