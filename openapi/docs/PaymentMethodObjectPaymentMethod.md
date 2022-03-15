# PaymentMethodObjectPaymentMethod

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
**Ticket** | Pointer to **string** | Ticket Digitalizado ( Total o parte del mismo por ejemplo la Firma Digitalizada )    codificado en Base64. | [optional] 
**TicketAnswerKey** | Pointer to **string** | Identificador Unívoco de la Transacción que se quiere Referenciar a la cual pertenece el Ticket Digitalizado. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. Si la firma fue capturada previamente y se envía en el requerimiento de la misma Operación Sale, Authorize*, Void, Return, Adjustment, DebtPayment, VoidDebtPayment o Enrollment no es necesario que se envíe este elemento o campo. | [optional] 
**Timeout** | Pointer to **float32** | Tiempo de espera que el POS espera al PINPAD para obtener la respuesta al requerimiento.  | [optional] 
**RequestKey** | Pointer to **string** | ID generado para la identificación por parte del Plataforma de la información generada en la ejecución de un GetCard o un Payment Method. Será necesario para que un mensaje de Sale, Void, PaymentMethod o Enrollment  identifique el contexto generado y lo utilice para esa operación. | [optional] 
**ReasonSequenceBreak** | Pointer to **string** | Motivo por el cual se requiere romper la secuencia. | [optional] 
**TransactionType** | Pointer to **string** | Tipo de Transacción (Sale, Void, Return, Authorize,...) por la cual se está realizado el requerimiento | [optional] 
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
**AlternativeAmount** | Pointer to **float32** | Monto con la que se realizó transacción. Si este valor es recibido, la búsqueda de los planes será limitada con este criterio. | [optional] 
**CurrencyCode** | Pointer to **string** | Código de Moneda - ISO 4217 &lt;https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica &lt;br /&gt;   * Num   - Alpha - Description &lt;br /&gt;   * &#39;032&#39; - &#39;ARS&#39; - Pesos Argentinos &lt;br /&gt;   * &#39;152&#39; - &#39;CLP&#39; - Pesos Chilenos &lt;br/&gt;   * &#39;484&#39; - &#39;MXN&#39; - Pesos Mexicanos &lt;br/&gt;   * &#39;840&#39; - &#39;USD&#39; - dólares Americanos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;858&#39; - &#39;UYU&#39; - Pesos Uruguayos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;986&#39; - &#39;BRL&#39; - Real Brasileño | [optional] 
**FacilityPayments** | Pointer to **float32** | Cantidad de cuotas en las que será realizada la transacción | [optional] 
**FacilityType** | Pointer to **string** | Tipo de Plan de Financiación | [optional] 
**CardAppName** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppIdentifier** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAppLabel** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthRequestCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardAuthResponseCryptogram** | Pointer to **string** | Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers. | [optional] 
**CardReadMode** | Pointer to **string** | Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes ) | [optional] 
**CardNumber** | Pointer to **string** | Número de Tarjeta, en el caso de las respuestas el mismo estará enmascarado. | [optional] 
**CardNumberMasked** | Pointer to **string** | Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón.             | [optional] 
**PaymentMethodId** | Pointer to **int32** | ID del  metodo de pago que se requerie la información | [optional] 
**PaymentMethodForeignIdentifier** | Pointer to **int32** | ID externo de este metodo de pago, utilizado por el punto de venta para reconocer al medio de pago en su base de datos | [optional] 
**OrigTransactionType** | Pointer to **string** | Tipo de operación financiera a la cual se pretende anular o devolver, para la cual se pretende utilizar esta lectura de tarjeta. Este campo sera de uso obligatorio solo para determinados tipos de dispositivos de captura. | [optional] 
**Issuer** | Pointer to **string** | ID del emisor que se encuentra disponible para este plan | [optional] 
**GetDinamicPlans** | Pointer to **float32** | Flag utilizado para indicar si se quiere conocer los planes dinamicos tambien, como Plan V de Prisma. | [optional] 
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relacionado a cada uno de los planes disponibles. | [optional] 

## Methods

### NewPaymentMethodObjectPaymentMethod

`func NewPaymentMethodObjectPaymentMethod(companyIdentification string, systemIdentification string, ) *PaymentMethodObjectPaymentMethod`

NewPaymentMethodObjectPaymentMethod instantiates a new PaymentMethodObjectPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodObjectPaymentMethodWithDefaults

`func NewPaymentMethodObjectPaymentMethodWithDefaults() *PaymentMethodObjectPaymentMethod`

NewPaymentMethodObjectPaymentMethodWithDefaults instantiates a new PaymentMethodObjectPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *PaymentMethodObjectPaymentMethod) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *PaymentMethodObjectPaymentMethod) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *PaymentMethodObjectPaymentMethod) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetSystemIdentification

`func (o *PaymentMethodObjectPaymentMethod) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *PaymentMethodObjectPaymentMethod) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *PaymentMethodObjectPaymentMethod) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetBranchIdentification

`func (o *PaymentMethodObjectPaymentMethod) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *PaymentMethodObjectPaymentMethod) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *PaymentMethodObjectPaymentMethod) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *PaymentMethodObjectPaymentMethod) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *PaymentMethodObjectPaymentMethod) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *PaymentMethodObjectPaymentMethod) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *PaymentMethodObjectPaymentMethod) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *PaymentMethodObjectPaymentMethod) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetRequestType

`func (o *PaymentMethodObjectPaymentMethod) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *PaymentMethodObjectPaymentMethod) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *PaymentMethodObjectPaymentMethod) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *PaymentMethodObjectPaymentMethod) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *PaymentMethodObjectPaymentMethod) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *PaymentMethodObjectPaymentMethod) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *PaymentMethodObjectPaymentMethod) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *PaymentMethodObjectPaymentMethod) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *PaymentMethodObjectPaymentMethod) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *PaymentMethodObjectPaymentMethod) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *PaymentMethodObjectPaymentMethod) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *PaymentMethodObjectPaymentMethod) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *PaymentMethodObjectPaymentMethod) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *PaymentMethodObjectPaymentMethod) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *PaymentMethodObjectPaymentMethod) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *PaymentMethodObjectPaymentMethod) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *PaymentMethodObjectPaymentMethod) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *PaymentMethodObjectPaymentMethod) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *PaymentMethodObjectPaymentMethod) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *PaymentMethodObjectPaymentMethod) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetTicket

`func (o *PaymentMethodObjectPaymentMethod) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *PaymentMethodObjectPaymentMethod) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *PaymentMethodObjectPaymentMethod) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *PaymentMethodObjectPaymentMethod) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTicketAnswerKey

`func (o *PaymentMethodObjectPaymentMethod) GetTicketAnswerKey() string`

GetTicketAnswerKey returns the TicketAnswerKey field if non-nil, zero value otherwise.

### GetTicketAnswerKeyOk

`func (o *PaymentMethodObjectPaymentMethod) GetTicketAnswerKeyOk() (*string, bool)`

GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketAnswerKey

`func (o *PaymentMethodObjectPaymentMethod) SetTicketAnswerKey(v string)`

SetTicketAnswerKey sets TicketAnswerKey field to given value.

### HasTicketAnswerKey

`func (o *PaymentMethodObjectPaymentMethod) HasTicketAnswerKey() bool`

HasTicketAnswerKey returns a boolean if a field has been set.

### GetTimeout

`func (o *PaymentMethodObjectPaymentMethod) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PaymentMethodObjectPaymentMethod) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PaymentMethodObjectPaymentMethod) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PaymentMethodObjectPaymentMethod) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRequestKey

`func (o *PaymentMethodObjectPaymentMethod) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *PaymentMethodObjectPaymentMethod) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *PaymentMethodObjectPaymentMethod) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *PaymentMethodObjectPaymentMethod) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetReasonSequenceBreak

`func (o *PaymentMethodObjectPaymentMethod) GetReasonSequenceBreak() string`

GetReasonSequenceBreak returns the ReasonSequenceBreak field if non-nil, zero value otherwise.

### GetReasonSequenceBreakOk

`func (o *PaymentMethodObjectPaymentMethod) GetReasonSequenceBreakOk() (*string, bool)`

GetReasonSequenceBreakOk returns a tuple with the ReasonSequenceBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSequenceBreak

`func (o *PaymentMethodObjectPaymentMethod) SetReasonSequenceBreak(v string)`

SetReasonSequenceBreak sets ReasonSequenceBreak field to given value.

### HasReasonSequenceBreak

`func (o *PaymentMethodObjectPaymentMethod) HasReasonSequenceBreak() bool`

HasReasonSequenceBreak returns a boolean if a field has been set.

### GetTransactionType

`func (o *PaymentMethodObjectPaymentMethod) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *PaymentMethodObjectPaymentMethod) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *PaymentMethodObjectPaymentMethod) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *PaymentMethodObjectPaymentMethod) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetPOSType

`func (o *PaymentMethodObjectPaymentMethod) GetPOSType() string`

GetPOSType returns the POSType field if non-nil, zero value otherwise.

### GetPOSTypeOk

`func (o *PaymentMethodObjectPaymentMethod) GetPOSTypeOk() (*string, bool)`

GetPOSTypeOk returns a tuple with the POSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSType

`func (o *PaymentMethodObjectPaymentMethod) SetPOSType(v string)`

SetPOSType sets POSType field to given value.

### HasPOSType

`func (o *PaymentMethodObjectPaymentMethod) HasPOSType() bool`

HasPOSType returns a boolean if a field has been set.

### GetPOSVersion

`func (o *PaymentMethodObjectPaymentMethod) GetPOSVersion() string`

GetPOSVersion returns the POSVersion field if non-nil, zero value otherwise.

### GetPOSVersionOk

`func (o *PaymentMethodObjectPaymentMethod) GetPOSVersionOk() (*string, bool)`

GetPOSVersionOk returns a tuple with the POSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSVersion

`func (o *PaymentMethodObjectPaymentMethod) SetPOSVersion(v string)`

SetPOSVersion sets POSVersion field to given value.

### HasPOSVersion

`func (o *PaymentMethodObjectPaymentMethod) HasPOSVersion() bool`

HasPOSVersion returns a boolean if a field has been set.

### GetPOSAddress

`func (o *PaymentMethodObjectPaymentMethod) GetPOSAddress() string`

GetPOSAddress returns the POSAddress field if non-nil, zero value otherwise.

### GetPOSAddressOk

`func (o *PaymentMethodObjectPaymentMethod) GetPOSAddressOk() (*string, bool)`

GetPOSAddressOk returns a tuple with the POSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSAddress

`func (o *PaymentMethodObjectPaymentMethod) SetPOSAddress(v string)`

SetPOSAddress sets POSAddress field to given value.

### HasPOSAddress

`func (o *PaymentMethodObjectPaymentMethod) HasPOSAddress() bool`

HasPOSAddress returns a boolean if a field has been set.

### GetPOSSerial

`func (o *PaymentMethodObjectPaymentMethod) GetPOSSerial() string`

GetPOSSerial returns the POSSerial field if non-nil, zero value otherwise.

### GetPOSSerialOk

`func (o *PaymentMethodObjectPaymentMethod) GetPOSSerialOk() (*string, bool)`

GetPOSSerialOk returns a tuple with the POSSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSSerial

`func (o *PaymentMethodObjectPaymentMethod) SetPOSSerial(v string)`

SetPOSSerial sets POSSerial field to given value.

### HasPOSSerial

`func (o *PaymentMethodObjectPaymentMethod) HasPOSSerial() bool`

HasPOSSerial returns a boolean if a field has been set.

### GetPOSGEO

`func (o *PaymentMethodObjectPaymentMethod) GetPOSGEO() string`

GetPOSGEO returns the POSGEO field if non-nil, zero value otherwise.

### GetPOSGEOOk

`func (o *PaymentMethodObjectPaymentMethod) GetPOSGEOOk() (*string, bool)`

GetPOSGEOOk returns a tuple with the POSGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSGEO

`func (o *PaymentMethodObjectPaymentMethod) SetPOSGEO(v string)`

SetPOSGEO sets POSGEO field to given value.

### HasPOSGEO

`func (o *PaymentMethodObjectPaymentMethod) HasPOSGEO() bool`

HasPOSGEO returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *PaymentMethodObjectPaymentMethod) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *PaymentMethodObjectPaymentMethod) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *PaymentMethodObjectPaymentMethod) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *PaymentMethodObjectPaymentMethod) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetReadingDeviceVersion

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceVersion() string`

GetReadingDeviceVersion returns the ReadingDeviceVersion field if non-nil, zero value otherwise.

### GetReadingDeviceVersionOk

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceVersionOk() (*string, bool)`

GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceVersion

`func (o *PaymentMethodObjectPaymentMethod) SetReadingDeviceVersion(v string)`

SetReadingDeviceVersion sets ReadingDeviceVersion field to given value.

### HasReadingDeviceVersion

`func (o *PaymentMethodObjectPaymentMethod) HasReadingDeviceVersion() bool`

HasReadingDeviceVersion returns a boolean if a field has been set.

### GetReadingDeviceAddress

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceAddress() string`

GetReadingDeviceAddress returns the ReadingDeviceAddress field if non-nil, zero value otherwise.

### GetReadingDeviceAddressOk

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceAddressOk() (*string, bool)`

GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceAddress

`func (o *PaymentMethodObjectPaymentMethod) SetReadingDeviceAddress(v string)`

SetReadingDeviceAddress sets ReadingDeviceAddress field to given value.

### HasReadingDeviceAddress

`func (o *PaymentMethodObjectPaymentMethod) HasReadingDeviceAddress() bool`

HasReadingDeviceAddress returns a boolean if a field has been set.

### GetReadingDeviceSerial

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceSerial() string`

GetReadingDeviceSerial returns the ReadingDeviceSerial field if non-nil, zero value otherwise.

### GetReadingDeviceSerialOk

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceSerialOk() (*string, bool)`

GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceSerial

`func (o *PaymentMethodObjectPaymentMethod) SetReadingDeviceSerial(v string)`

SetReadingDeviceSerial sets ReadingDeviceSerial field to given value.

### HasReadingDeviceSerial

`func (o *PaymentMethodObjectPaymentMethod) HasReadingDeviceSerial() bool`

HasReadingDeviceSerial returns a boolean if a field has been set.

### GetReadingDeviceGEO

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceGEO() string`

GetReadingDeviceGEO returns the ReadingDeviceGEO field if non-nil, zero value otherwise.

### GetReadingDeviceGEOOk

`func (o *PaymentMethodObjectPaymentMethod) GetReadingDeviceGEOOk() (*string, bool)`

GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceGEO

`func (o *PaymentMethodObjectPaymentMethod) SetReadingDeviceGEO(v string)`

SetReadingDeviceGEO sets ReadingDeviceGEO field to given value.

### HasReadingDeviceGEO

`func (o *PaymentMethodObjectPaymentMethod) HasReadingDeviceGEO() bool`

HasReadingDeviceGEO returns a boolean if a field has been set.

### GetUserID

`func (o *PaymentMethodObjectPaymentMethod) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *PaymentMethodObjectPaymentMethod) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *PaymentMethodObjectPaymentMethod) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *PaymentMethodObjectPaymentMethod) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *PaymentMethodObjectPaymentMethod) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PaymentMethodObjectPaymentMethod) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PaymentMethodObjectPaymentMethod) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *PaymentMethodObjectPaymentMethod) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentMethodObjectPaymentMethod) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentMethodObjectPaymentMethod) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentMethodObjectPaymentMethod) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentMethodObjectPaymentMethod) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAlternativeAmount

`func (o *PaymentMethodObjectPaymentMethod) GetAlternativeAmount() float32`

GetAlternativeAmount returns the AlternativeAmount field if non-nil, zero value otherwise.

### GetAlternativeAmountOk

`func (o *PaymentMethodObjectPaymentMethod) GetAlternativeAmountOk() (*float32, bool)`

GetAlternativeAmountOk returns a tuple with the AlternativeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeAmount

`func (o *PaymentMethodObjectPaymentMethod) SetAlternativeAmount(v float32)`

SetAlternativeAmount sets AlternativeAmount field to given value.

### HasAlternativeAmount

`func (o *PaymentMethodObjectPaymentMethod) HasAlternativeAmount() bool`

HasAlternativeAmount returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PaymentMethodObjectPaymentMethod) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PaymentMethodObjectPaymentMethod) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PaymentMethodObjectPaymentMethod) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PaymentMethodObjectPaymentMethod) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetFacilityPayments

`func (o *PaymentMethodObjectPaymentMethod) GetFacilityPayments() float32`

GetFacilityPayments returns the FacilityPayments field if non-nil, zero value otherwise.

### GetFacilityPaymentsOk

`func (o *PaymentMethodObjectPaymentMethod) GetFacilityPaymentsOk() (*float32, bool)`

GetFacilityPaymentsOk returns a tuple with the FacilityPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityPayments

`func (o *PaymentMethodObjectPaymentMethod) SetFacilityPayments(v float32)`

SetFacilityPayments sets FacilityPayments field to given value.

### HasFacilityPayments

`func (o *PaymentMethodObjectPaymentMethod) HasFacilityPayments() bool`

HasFacilityPayments returns a boolean if a field has been set.

### GetFacilityType

`func (o *PaymentMethodObjectPaymentMethod) GetFacilityType() string`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *PaymentMethodObjectPaymentMethod) GetFacilityTypeOk() (*string, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *PaymentMethodObjectPaymentMethod) SetFacilityType(v string)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *PaymentMethodObjectPaymentMethod) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### GetCardAppName

`func (o *PaymentMethodObjectPaymentMethod) GetCardAppName() string`

GetCardAppName returns the CardAppName field if non-nil, zero value otherwise.

### GetCardAppNameOk

`func (o *PaymentMethodObjectPaymentMethod) GetCardAppNameOk() (*string, bool)`

GetCardAppNameOk returns a tuple with the CardAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppName

`func (o *PaymentMethodObjectPaymentMethod) SetCardAppName(v string)`

SetCardAppName sets CardAppName field to given value.

### HasCardAppName

`func (o *PaymentMethodObjectPaymentMethod) HasCardAppName() bool`

HasCardAppName returns a boolean if a field has been set.

### GetCardAppIdentifier

`func (o *PaymentMethodObjectPaymentMethod) GetCardAppIdentifier() string`

GetCardAppIdentifier returns the CardAppIdentifier field if non-nil, zero value otherwise.

### GetCardAppIdentifierOk

`func (o *PaymentMethodObjectPaymentMethod) GetCardAppIdentifierOk() (*string, bool)`

GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppIdentifier

`func (o *PaymentMethodObjectPaymentMethod) SetCardAppIdentifier(v string)`

SetCardAppIdentifier sets CardAppIdentifier field to given value.

### HasCardAppIdentifier

`func (o *PaymentMethodObjectPaymentMethod) HasCardAppIdentifier() bool`

HasCardAppIdentifier returns a boolean if a field has been set.

### GetCardAppLabel

`func (o *PaymentMethodObjectPaymentMethod) GetCardAppLabel() string`

GetCardAppLabel returns the CardAppLabel field if non-nil, zero value otherwise.

### GetCardAppLabelOk

`func (o *PaymentMethodObjectPaymentMethod) GetCardAppLabelOk() (*string, bool)`

GetCardAppLabelOk returns a tuple with the CardAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAppLabel

`func (o *PaymentMethodObjectPaymentMethod) SetCardAppLabel(v string)`

SetCardAppLabel sets CardAppLabel field to given value.

### HasCardAppLabel

`func (o *PaymentMethodObjectPaymentMethod) HasCardAppLabel() bool`

HasCardAppLabel returns a boolean if a field has been set.

### GetCardAuthRequestCryptogram

`func (o *PaymentMethodObjectPaymentMethod) GetCardAuthRequestCryptogram() string`

GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field if non-nil, zero value otherwise.

### GetCardAuthRequestCryptogramOk

`func (o *PaymentMethodObjectPaymentMethod) GetCardAuthRequestCryptogramOk() (*string, bool)`

GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthRequestCryptogram

`func (o *PaymentMethodObjectPaymentMethod) SetCardAuthRequestCryptogram(v string)`

SetCardAuthRequestCryptogram sets CardAuthRequestCryptogram field to given value.

### HasCardAuthRequestCryptogram

`func (o *PaymentMethodObjectPaymentMethod) HasCardAuthRequestCryptogram() bool`

HasCardAuthRequestCryptogram returns a boolean if a field has been set.

### GetCardAuthResponseCryptogram

`func (o *PaymentMethodObjectPaymentMethod) GetCardAuthResponseCryptogram() string`

GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field if non-nil, zero value otherwise.

### GetCardAuthResponseCryptogramOk

`func (o *PaymentMethodObjectPaymentMethod) GetCardAuthResponseCryptogramOk() (*string, bool)`

GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAuthResponseCryptogram

`func (o *PaymentMethodObjectPaymentMethod) SetCardAuthResponseCryptogram(v string)`

SetCardAuthResponseCryptogram sets CardAuthResponseCryptogram field to given value.

### HasCardAuthResponseCryptogram

`func (o *PaymentMethodObjectPaymentMethod) HasCardAuthResponseCryptogram() bool`

HasCardAuthResponseCryptogram returns a boolean if a field has been set.

### GetCardReadMode

`func (o *PaymentMethodObjectPaymentMethod) GetCardReadMode() string`

GetCardReadMode returns the CardReadMode field if non-nil, zero value otherwise.

### GetCardReadModeOk

`func (o *PaymentMethodObjectPaymentMethod) GetCardReadModeOk() (*string, bool)`

GetCardReadModeOk returns a tuple with the CardReadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardReadMode

`func (o *PaymentMethodObjectPaymentMethod) SetCardReadMode(v string)`

SetCardReadMode sets CardReadMode field to given value.

### HasCardReadMode

`func (o *PaymentMethodObjectPaymentMethod) HasCardReadMode() bool`

HasCardReadMode returns a boolean if a field has been set.

### GetCardNumber

`func (o *PaymentMethodObjectPaymentMethod) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *PaymentMethodObjectPaymentMethod) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *PaymentMethodObjectPaymentMethod) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *PaymentMethodObjectPaymentMethod) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNumberMasked

`func (o *PaymentMethodObjectPaymentMethod) GetCardNumberMasked() string`

GetCardNumberMasked returns the CardNumberMasked field if non-nil, zero value otherwise.

### GetCardNumberMaskedOk

`func (o *PaymentMethodObjectPaymentMethod) GetCardNumberMaskedOk() (*string, bool)`

GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberMasked

`func (o *PaymentMethodObjectPaymentMethod) SetCardNumberMasked(v string)`

SetCardNumberMasked sets CardNumberMasked field to given value.

### HasCardNumberMasked

`func (o *PaymentMethodObjectPaymentMethod) HasCardNumberMasked() bool`

HasCardNumberMasked returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *PaymentMethodObjectPaymentMethod) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PaymentMethodObjectPaymentMethod) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PaymentMethodObjectPaymentMethod) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *PaymentMethodObjectPaymentMethod) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetPaymentMethodForeignIdentifier

`func (o *PaymentMethodObjectPaymentMethod) GetPaymentMethodForeignIdentifier() int32`

GetPaymentMethodForeignIdentifier returns the PaymentMethodForeignIdentifier field if non-nil, zero value otherwise.

### GetPaymentMethodForeignIdentifierOk

`func (o *PaymentMethodObjectPaymentMethod) GetPaymentMethodForeignIdentifierOk() (*int32, bool)`

GetPaymentMethodForeignIdentifierOk returns a tuple with the PaymentMethodForeignIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodForeignIdentifier

`func (o *PaymentMethodObjectPaymentMethod) SetPaymentMethodForeignIdentifier(v int32)`

SetPaymentMethodForeignIdentifier sets PaymentMethodForeignIdentifier field to given value.

### HasPaymentMethodForeignIdentifier

`func (o *PaymentMethodObjectPaymentMethod) HasPaymentMethodForeignIdentifier() bool`

HasPaymentMethodForeignIdentifier returns a boolean if a field has been set.

### GetOrigTransactionType

`func (o *PaymentMethodObjectPaymentMethod) GetOrigTransactionType() string`

GetOrigTransactionType returns the OrigTransactionType field if non-nil, zero value otherwise.

### GetOrigTransactionTypeOk

`func (o *PaymentMethodObjectPaymentMethod) GetOrigTransactionTypeOk() (*string, bool)`

GetOrigTransactionTypeOk returns a tuple with the OrigTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigTransactionType

`func (o *PaymentMethodObjectPaymentMethod) SetOrigTransactionType(v string)`

SetOrigTransactionType sets OrigTransactionType field to given value.

### HasOrigTransactionType

`func (o *PaymentMethodObjectPaymentMethod) HasOrigTransactionType() bool`

HasOrigTransactionType returns a boolean if a field has been set.

### GetIssuer

`func (o *PaymentMethodObjectPaymentMethod) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PaymentMethodObjectPaymentMethod) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PaymentMethodObjectPaymentMethod) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PaymentMethodObjectPaymentMethod) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetGetDinamicPlans

`func (o *PaymentMethodObjectPaymentMethod) GetGetDinamicPlans() float32`

GetGetDinamicPlans returns the GetDinamicPlans field if non-nil, zero value otherwise.

### GetGetDinamicPlansOk

`func (o *PaymentMethodObjectPaymentMethod) GetGetDinamicPlansOk() (*float32, bool)`

GetGetDinamicPlansOk returns a tuple with the GetDinamicPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetDinamicPlans

`func (o *PaymentMethodObjectPaymentMethod) SetGetDinamicPlans(v float32)`

SetGetDinamicPlans sets GetDinamicPlans field to given value.

### HasGetDinamicPlans

`func (o *PaymentMethodObjectPaymentMethod) HasGetDinamicPlans() bool`

HasGetDinamicPlans returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *PaymentMethodObjectPaymentMethod) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *PaymentMethodObjectPaymentMethod) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *PaymentMethodObjectPaymentMethod) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *PaymentMethodObjectPaymentMethod) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *PaymentMethodObjectPaymentMethod) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *PaymentMethodObjectPaymentMethod) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *PaymentMethodObjectPaymentMethod) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *PaymentMethodObjectPaymentMethod) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


