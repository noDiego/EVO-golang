# VoidDebtPaymentObjectVoidDebtPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyIdentification** | **string** | Identificador de la Compañía para la Plataforma de Integración. | 
**SystemIdentification** | **string** | Identificador de Aplicativo para la Plataforma de Integración que usa la Compañía especificada. | 
**BranchIdentification** | Pointer to **string** | Identificador de Aplicativo para la Plataforma de Integración que usa la Compañía especificada. | [optional] 
**POSIdentification** | Pointer to **string** | Identificador del Punto de Venta/POS/Caja que pertenece a la sucursal y companía especificada. | [optional] 
**RequestType** | Pointer to **string** | Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**Block** | Pointer to **string** | ID que identifica a un grupo de transacciones que serán confirmadas o canceladas | [optional] 
**RequiredInformation** | Pointer to [**[]DebtPaymentObjectDebtPaymentRequiredInformation**](DebtPaymentObjectDebtPaymentRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**Ticket** | Pointer to **string** | Ticket Digitalizado ( Total o parte del mismo por ejemplo la Firma Digitalizada )    codificado en Base64. | [optional] 
**TicketAnswerKey** | Pointer to **string** | Identificador Unívoco de la Transacción que se quiere Referenciar a la cual pertenece el Ticket Digitalizado. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. Si firma fue capturada previamente y se envía en el requerimiento de la misma Operación Sale, Authorize*, Void, Return, Adjustment, DebtPayment o VoidDebtPayment no es necesario que se envíe este elemento o campo. | [optional] 
**Timeout** | Pointer to **float32** | Tiempo de espera que el POS espera al PINPAD para obtener la respuesta al requerimiento.  | [optional] 
**RequestKey** | Pointer to **string** | ID generado para la identificación por parte del Plataforma de la información generada en la ejecución de un GetCard o un Payment Method. Sera necesario para que un mensaje de Sale, Void o Payment Method identifique el contexto generado y lo utilice para esa operación. | [optional] 
**IsReverse** | Pointer to **float32** | Identificador de Reversa. | [optional] 
**ReverseReason** | Pointer to **string** | Motivo por el cual se requiere generar la reversa. | [optional] 
**ReasonSequenceBreak** | Pointer to **string** | Motivo por el cual se requiere romper la secuencia. | [optional] 
**Reference** | Pointer to **string** | Referencia de la transacción para el punto de venta | [optional] 
**OrigReference** | Pointer to **string** | Referencia de la transacción para el punto de venta. Corresponde a la transacción original que se está referenciando. | [optional] 
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
**CardReadMode** | Pointer to **string** | Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes ) | [optional] 
**CardGetMode** | Pointer to **string** | Indican por cada elemento que contiene los datos sensibles, si están encriptados  y también el algoritmo usado. En caso de no estar especificado se asume PLAIN. | [optional] 
**CardNumber** | Pointer to **string** | Número de Tarjeta, en el caso de las respuestas el mismo estará enmascarado. | [optional] 
**CardNumberMasked** | Pointer to **string** | Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón. | [optional] 
**CardNumberEncrypted** | Pointer to **string** | Número de tarjeta encriptado. | [optional] 
**CardExp** | Pointer to **string** | Fecha de vencimiento de la tarjeta. Este dato será necesario si el modo de ingreso fue manual/digitada. | [optional] 
**Track1** | Pointer to **string** | Banda Número 1 de la tarjeta. Este dato será necesario si el modo de ingreso fue por banda (deslizando la banda por el lector). | [optional] 
**Track2** | Pointer to **string** | Banda Número 2 de la tarjeta. Este dato será necesario si el modo de ingreso fue por banda (deslizando la banda por el lector). | [optional] 
**SecurityCode** | Pointer to **string** | Código de seguridad de la tarjeta. | [optional] 
**Pin** | Pointer to **string** | PIN block | [optional] 
**CardCryptogram** | Pointer to **string** | Tags EMV en formato TLV Leídos. | [optional] 
**CardAppName** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppIdentifier** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppLabel** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthRequestCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthResponseCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CredentialToken** | Pointer to **string** | Token asociado a la Credencial Enrolada | [optional] 
**CredentialIssuerToken** | Pointer to **string** | Emisor del Token asociado a la credencial enrolada | [optional] 
**InputTokens** | Pointer to [**[]SaleObjectSaleInputTokens**](SaleObjectSaleInputTokens.md) | Tokens. | [optional] 
**Products** | Pointer to [**[]VoidDebtPaymentObjectVoidDebtPaymentProducts**](VoidDebtPaymentObjectVoidDebtPaymentProducts.md) | Detalle de Productos de la Operación. | [optional] 
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relacionado a cada uno de los planes disponibles. | [optional] 
**TerminalID** | Pointer to **string** | Identificador de Terminal por el cual se envía la Transacción al Host. | [optional] 
**TerminalTrace** | Pointer to **int32** | Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID. | [optional] 
**SettlementBatchNumber** | Pointer to **int32** | Para aquellos host que exista el concepto de lote, es el número de lote al cual pertenece la transacción. | [optional] 

## Methods

### NewVoidDebtPaymentObjectVoidDebtPayment

`func NewVoidDebtPaymentObjectVoidDebtPayment(companyIdentification string, systemIdentification string, ) *VoidDebtPaymentObjectVoidDebtPayment`

NewVoidDebtPaymentObjectVoidDebtPayment instantiates a new VoidDebtPaymentObjectVoidDebtPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidDebtPaymentObjectVoidDebtPaymentWithDefaults

`func NewVoidDebtPaymentObjectVoidDebtPaymentWithDefaults() *VoidDebtPaymentObjectVoidDebtPayment`

NewVoidDebtPaymentObjectVoidDebtPaymentWithDefaults instantiates a new VoidDebtPaymentObjectVoidDebtPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetSystemIdentification

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetBranchIdentification

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetRequestType

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetRequiredInformationOk() (*[]DebtPaymentObjectDebtPaymentRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetTicket

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTicketAnswerKey

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTicketAnswerKey() string`

GetTicketAnswerKey returns the TicketAnswerKey field if non-nil, zero value otherwise.

### GetTicketAnswerKeyOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTicketAnswerKeyOk() (*string, bool)`

GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketAnswerKey

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetTicketAnswerKey(v string)`

SetTicketAnswerKey sets TicketAnswerKey field to given value.

### HasTicketAnswerKey

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasTicketAnswerKey() bool`

HasTicketAnswerKey returns a boolean if a field has been set.

### GetTimeout

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRequestKey

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetIsReverse

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetIsReverse() float32`

GetIsReverse returns the IsReverse field if non-nil, zero value otherwise.

### GetIsReverseOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetIsReverseOk() (*float32, bool)`

GetIsReverseOk returns a tuple with the IsReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReverse

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetIsReverse(v float32)`

SetIsReverse sets IsReverse field to given value.

### HasIsReverse

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasIsReverse() bool`

HasIsReverse returns a boolean if a field has been set.

### GetReverseReason

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReverseReason() string`

GetReverseReason returns the ReverseReason field if non-nil, zero value otherwise.

### GetReverseReasonOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReverseReasonOk() (*string, bool)`

GetReverseReasonOk returns a tuple with the ReverseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseReason

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetReverseReason(v string)`

SetReverseReason sets ReverseReason field to given value.

### HasReverseReason

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasReverseReason() bool`

HasReverseReason returns a boolean if a field has been set.

### GetReasonSequenceBreak

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReasonSequenceBreak() string`

GetReasonSequenceBreak returns the ReasonSequenceBreak field if non-nil, zero value otherwise.

### GetReasonSequenceBreakOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReasonSequenceBreakOk() (*string, bool)`

GetReasonSequenceBreakOk returns a tuple with the ReasonSequenceBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSequenceBreak

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetReasonSequenceBreak(v string)`

SetReasonSequenceBreak sets ReasonSequenceBreak field to given value.

### HasReasonSequenceBreak

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasReasonSequenceBreak() bool`

HasReasonSequenceBreak returns a boolean if a field has been set.

### GetReference

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetOrigReference

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetOrigReference() string`

GetOrigReference returns the OrigReference field if non-nil, zero value otherwise.

### GetOrigReferenceOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetOrigReferenceOk() (*string, bool)`

GetOrigReferenceOk returns a tuple with the OrigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigReference

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetOrigReference(v string)`

SetOrigReference sets OrigReference field to given value.

### HasOrigReference

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasOrigReference() bool`

HasOrigReference returns a boolean if a field has been set.

### GetPOSType

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSType() string`

GetPOSType returns the POSType field if non-nil, zero value otherwise.

### GetPOSTypeOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSTypeOk() (*string, bool)`

GetPOSTypeOk returns a tuple with the POSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSType

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetPOSType(v string)`

SetPOSType sets POSType field to given value.

### HasPOSType

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasPOSType() bool`

HasPOSType returns a boolean if a field has been set.

### GetPOSVersion

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSVersion() string`

GetPOSVersion returns the POSVersion field if non-nil, zero value otherwise.

### GetPOSVersionOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSVersionOk() (*string, bool)`

GetPOSVersionOk returns a tuple with the POSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSVersion

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetPOSVersion(v string)`

SetPOSVersion sets POSVersion field to given value.

### HasPOSVersion

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasPOSVersion() bool`

HasPOSVersion returns a boolean if a field has been set.

### GetPOSAddress

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSAddress() string`

GetPOSAddress returns the POSAddress field if non-nil, zero value otherwise.

### GetPOSAddressOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSAddressOk() (*string, bool)`

GetPOSAddressOk returns a tuple with the POSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSAddress

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetPOSAddress(v string)`

SetPOSAddress sets POSAddress field to given value.

### HasPOSAddress

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasPOSAddress() bool`

HasPOSAddress returns a boolean if a field has been set.

### GetPOSSerial

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSSerial() string`

GetPOSSerial returns the POSSerial field if non-nil, zero value otherwise.

### GetPOSSerialOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSSerialOk() (*string, bool)`

GetPOSSerialOk returns a tuple with the POSSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSSerial

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetPOSSerial(v string)`

SetPOSSerial sets POSSerial field to given value.

### HasPOSSerial

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasPOSSerial() bool`

HasPOSSerial returns a boolean if a field has been set.

### GetPOSGEO

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSGEO() string`

GetPOSGEO returns the POSGEO field if non-nil, zero value otherwise.

### GetPOSGEOOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPOSGEOOk() (*string, bool)`

GetPOSGEOOk returns a tuple with the POSGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSGEO

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetPOSGEO(v string)`

SetPOSGEO sets POSGEO field to given value.

### HasPOSGEO

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasPOSGEO() bool`

HasPOSGEO returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetReadingDeviceVersion

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceVersion() string`

GetReadingDeviceVersion returns the ReadingDeviceVersion field if non-nil, zero value otherwise.

### GetReadingDeviceVersionOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceVersionOk() (*string, bool)`

GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceVersion

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetReadingDeviceVersion(v string)`

SetReadingDeviceVersion sets ReadingDeviceVersion field to given value.

### HasReadingDeviceVersion

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasReadingDeviceVersion() bool`

HasReadingDeviceVersion returns a boolean if a field has been set.

### GetReadingDeviceAddress

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceAddress() string`

GetReadingDeviceAddress returns the ReadingDeviceAddress field if non-nil, zero value otherwise.

### GetReadingDeviceAddressOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceAddressOk() (*string, bool)`

GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceAddress

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetReadingDeviceAddress(v string)`

SetReadingDeviceAddress sets ReadingDeviceAddress field to given value.

### HasReadingDeviceAddress

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasReadingDeviceAddress() bool`

HasReadingDeviceAddress returns a boolean if a field has been set.

### GetReadingDeviceSerial

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceSerial() string`

GetReadingDeviceSerial returns the ReadingDeviceSerial field if non-nil, zero value otherwise.

### GetReadingDeviceSerialOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceSerialOk() (*string, bool)`

GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceSerial

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetReadingDeviceSerial(v string)`

SetReadingDeviceSerial sets ReadingDeviceSerial field to given value.

### HasReadingDeviceSerial

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasReadingDeviceSerial() bool`

HasReadingDeviceSerial returns a boolean if a field has been set.

### GetReadingDeviceGEO

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceGEO() string`

GetReadingDeviceGEO returns the ReadingDeviceGEO field if non-nil, zero value otherwise.

### GetReadingDeviceGEOOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetReadingDeviceGEOOk() (*string, bool)`

GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceGEO

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetReadingDeviceGEO(v string)`

SetReadingDeviceGEO sets ReadingDeviceGEO field to given value.

### HasReadingDeviceGEO

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasReadingDeviceGEO() bool`

HasReadingDeviceGEO returns a boolean if a field has been set.

### GetUserID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetCardReadMode

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardReadMode() string`

GetCardReadMode returns the CardReadMode field if non-nil, zero value otherwise.

### GetCardReadModeOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardReadModeOk() (*string, bool)`

GetCardReadModeOk returns a tuple with the CardReadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadMode

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardReadMode(v string)`

SetCardReadMode sets CardReadMode field to given value.

### HasCardReadMode

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardReadMode() bool`

HasCardReadMode returns a boolean if a field has been set.

### GetCardGetMode

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardGetMode() string`

GetCardGetMode returns the CardGetMode field if non-nil, zero value otherwise.

### GetCardGetModeOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardGetModeOk() (*string, bool)`

GetCardGetModeOk returns a tuple with the CardGetMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardGetMode

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardGetMode(v string)`

SetCardGetMode sets CardGetMode field to given value.

### HasCardGetMode

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardGetMode() bool`

HasCardGetMode returns a boolean if a field has been set.

### GetCardNumber

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNumberMasked

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardNumberMasked() string`

GetCardNumberMasked returns the CardNumberMasked field if non-nil, zero value otherwise.

### GetCardNumberMaskedOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardNumberMaskedOk() (*string, bool)`

GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMasked

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardNumberMasked(v string)`

SetCardNumberMasked sets CardNumberMasked field to given value.

### HasCardNumberMasked

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardNumberMasked() bool`

HasCardNumberMasked returns a boolean if a field has been set.

### GetCardNumberEncrypted

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardNumberEncrypted() string`

GetCardNumberEncrypted returns the CardNumberEncrypted field if non-nil, zero value otherwise.

### GetCardNumberEncryptedOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardNumberEncryptedOk() (*string, bool)`

GetCardNumberEncryptedOk returns a tuple with the CardNumberEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberEncrypted

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardNumberEncrypted(v string)`

SetCardNumberEncrypted sets CardNumberEncrypted field to given value.

### HasCardNumberEncrypted

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardNumberEncrypted() bool`

HasCardNumberEncrypted returns a boolean if a field has been set.

### GetCardExp

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardExp() string`

GetCardExp returns the CardExp field if non-nil, zero value otherwise.

### GetCardExpOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardExpOk() (*string, bool)`

GetCardExpOk returns a tuple with the CardExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExp

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardExp(v string)`

SetCardExp sets CardExp field to given value.

### HasCardExp

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardExp() bool`

HasCardExp returns a boolean if a field has been set.

### GetTrack1

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTrack1() string`

GetTrack1 returns the Track1 field if non-nil, zero value otherwise.

### GetTrack1Ok

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTrack1Ok() (*string, bool)`

GetTrack1Ok returns a tuple with the Track1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack1

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetTrack1(v string)`

SetTrack1 sets Track1 field to given value.

### HasTrack1

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasTrack1() bool`

HasTrack1 returns a boolean if a field has been set.

### GetTrack2

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTrack2() string`

GetTrack2 returns the Track2 field if non-nil, zero value otherwise.

### GetTrack2Ok

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTrack2Ok() (*string, bool)`

GetTrack2Ok returns a tuple with the Track2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack2

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetTrack2(v string)`

SetTrack2 sets Track2 field to given value.

### HasTrack2

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasTrack2() bool`

HasTrack2 returns a boolean if a field has been set.

### GetSecurityCode

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.

### GetPin

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetCardCryptogram

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardCryptogram() string`

GetCardCryptogram returns the CardCryptogram field if non-nil, zero value otherwise.

### GetCardCryptogramOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardCryptogramOk() (*string, bool)`

GetCardCryptogramOk returns a tuple with the CardCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCryptogram

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardCryptogram(v string)`

SetCardCryptogram sets CardCryptogram field to given value.

### HasCardCryptogram

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardCryptogram() bool`

HasCardCryptogram returns a boolean if a field has been set.

### GetCardAppName

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardAppName() string`

GetCardAppName returns the CardAppName field if non-nil, zero value otherwise.

### GetCardAppNameOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardAppNameOk() (*string, bool)`

GetCardAppNameOk returns a tuple with the CardAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppName

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardAppName(v string)`

SetCardAppName sets CardAppName field to given value.

### HasCardAppName

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardAppName() bool`

HasCardAppName returns a boolean if a field has been set.

### GetCardAppIdentifier

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardAppIdentifier() string`

GetCardAppIdentifier returns the CardAppIdentifier field if non-nil, zero value otherwise.

### GetCardAppIdentifierOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardAppIdentifierOk() (*string, bool)`

GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppIdentifier

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardAppIdentifier(v string)`

SetCardAppIdentifier sets CardAppIdentifier field to given value.

### HasCardAppIdentifier

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardAppIdentifier() bool`

HasCardAppIdentifier returns a boolean if a field has been set.

### GetCardAppLabel

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardAppLabel() string`

GetCardAppLabel returns the CardAppLabel field if non-nil, zero value otherwise.

### GetCardAppLabelOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardAppLabelOk() (*string, bool)`

GetCardAppLabelOk returns a tuple with the CardAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppLabel

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardAppLabel(v string)`

SetCardAppLabel sets CardAppLabel field to given value.

### HasCardAppLabel

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardAppLabel() bool`

HasCardAppLabel returns a boolean if a field has been set.

### GetCardAuthRequestCryptogram

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardAuthRequestCryptogram() string`

GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field if non-nil, zero value otherwise.

### GetCardAuthRequestCryptogramOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardAuthRequestCryptogramOk() (*string, bool)`

GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthRequestCryptogram

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardAuthRequestCryptogram(v string)`

SetCardAuthRequestCryptogram sets CardAuthRequestCryptogram field to given value.

### HasCardAuthRequestCryptogram

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardAuthRequestCryptogram() bool`

HasCardAuthRequestCryptogram returns a boolean if a field has been set.

### GetCardAuthResponseCryptogram

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardAuthResponseCryptogram() string`

GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field if non-nil, zero value otherwise.

### GetCardAuthResponseCryptogramOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCardAuthResponseCryptogramOk() (*string, bool)`

GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthResponseCryptogram

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCardAuthResponseCryptogram(v string)`

SetCardAuthResponseCryptogram sets CardAuthResponseCryptogram field to given value.

### HasCardAuthResponseCryptogram

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCardAuthResponseCryptogram() bool`

HasCardAuthResponseCryptogram returns a boolean if a field has been set.

### GetCredentialToken

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCredentialToken() string`

GetCredentialToken returns the CredentialToken field if non-nil, zero value otherwise.

### GetCredentialTokenOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCredentialTokenOk() (*string, bool)`

GetCredentialTokenOk returns a tuple with the CredentialToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialToken

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCredentialToken(v string)`

SetCredentialToken sets CredentialToken field to given value.

### HasCredentialToken

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCredentialToken() bool`

HasCredentialToken returns a boolean if a field has been set.

### GetCredentialIssuerToken

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCredentialIssuerToken() string`

GetCredentialIssuerToken returns the CredentialIssuerToken field if non-nil, zero value otherwise.

### GetCredentialIssuerTokenOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetCredentialIssuerTokenOk() (*string, bool)`

GetCredentialIssuerTokenOk returns a tuple with the CredentialIssuerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuerToken

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetCredentialIssuerToken(v string)`

SetCredentialIssuerToken sets CredentialIssuerToken field to given value.

### HasCredentialIssuerToken

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasCredentialIssuerToken() bool`

HasCredentialIssuerToken returns a boolean if a field has been set.

### GetInputTokens

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetInputTokens() []SaleObjectSaleInputTokens`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetInputTokensOk() (*[]SaleObjectSaleInputTokens, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetInputTokens(v []SaleObjectSaleInputTokens)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetProducts

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetProducts() []VoidDebtPaymentObjectVoidDebtPaymentProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetProductsOk() (*[]VoidDebtPaymentObjectVoidDebtPaymentProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetProducts(v []VoidDebtPaymentObjectVoidDebtPaymentProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetTerminalID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTerminalID() string`

GetTerminalID returns the TerminalID field if non-nil, zero value otherwise.

### GetTerminalIDOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTerminalIDOk() (*string, bool)`

GetTerminalIDOk returns a tuple with the TerminalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetTerminalID(v string)`

SetTerminalID sets TerminalID field to given value.

### HasTerminalID

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasTerminalID() bool`

HasTerminalID returns a boolean if a field has been set.

### GetTerminalTrace

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTerminalTrace() int32`

GetTerminalTrace returns the TerminalTrace field if non-nil, zero value otherwise.

### GetTerminalTraceOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetTerminalTraceOk() (*int32, bool)`

GetTerminalTraceOk returns a tuple with the TerminalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTrace

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetTerminalTrace(v int32)`

SetTerminalTrace sets TerminalTrace field to given value.

### HasTerminalTrace

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasTerminalTrace() bool`

HasTerminalTrace returns a boolean if a field has been set.

### GetSettlementBatchNumber

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetSettlementBatchNumber() int32`

GetSettlementBatchNumber returns the SettlementBatchNumber field if non-nil, zero value otherwise.

### GetSettlementBatchNumberOk

`func (o *VoidDebtPaymentObjectVoidDebtPayment) GetSettlementBatchNumberOk() (*int32, bool)`

GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBatchNumber

`func (o *VoidDebtPaymentObjectVoidDebtPayment) SetSettlementBatchNumber(v int32)`

SetSettlementBatchNumber sets SettlementBatchNumber field to given value.

### HasSettlementBatchNumber

`func (o *VoidDebtPaymentObjectVoidDebtPayment) HasSettlementBatchNumber() bool`

HasSettlementBatchNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


