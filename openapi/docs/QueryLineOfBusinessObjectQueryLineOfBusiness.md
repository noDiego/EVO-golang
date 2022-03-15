# QueryLineOfBusinessObjectQueryLineOfBusiness

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
**TransactionIdentification** | Pointer to **string** | ID de La transacción UNIVOCO para el Punto de venta | [optional] 
**TransactionDescription** | Pointer to **string** | Descripción del tipo de operación que se realizará | [optional] 
**TrasactionDateTime** | Pointer to **string** | Fecha y Hora de la transacción generada por el Punto de Venta - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
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
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relacionado a cada uno de los planes disponibles. | [optional] 
**TerminalID** | Pointer to **string** | Identificador de Terminal por el cual se envía la Transacción al Host. | [optional] 
**TerminalTrace** | Pointer to **int32** | Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID. | [optional] 
**SettlementBatchNumber** | Pointer to **int32** | Para aquellos host que exista el concepto de lote, es el número de lote al cual pertenece la transacción. | [optional] 

## Methods

### NewQueryLineOfBusinessObjectQueryLineOfBusiness

`func NewQueryLineOfBusinessObjectQueryLineOfBusiness(companyIdentification string, systemIdentification string, ) *QueryLineOfBusinessObjectQueryLineOfBusiness`

NewQueryLineOfBusinessObjectQueryLineOfBusiness instantiates a new QueryLineOfBusinessObjectQueryLineOfBusiness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLineOfBusinessObjectQueryLineOfBusinessWithDefaults

`func NewQueryLineOfBusinessObjectQueryLineOfBusinessWithDefaults() *QueryLineOfBusinessObjectQueryLineOfBusiness`

NewQueryLineOfBusinessObjectQueryLineOfBusinessWithDefaults instantiates a new QueryLineOfBusinessObjectQueryLineOfBusiness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetSystemIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetBranchIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetRequestType

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetRequiredInformation() []SaleObjectSaleRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetRequiredInformationOk() (*[]SaleObjectSaleRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetRequiredInformation(v []SaleObjectSaleRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetTicket

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTicketAnswerKey

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTicketAnswerKey() string`

GetTicketAnswerKey returns the TicketAnswerKey field if non-nil, zero value otherwise.

### GetTicketAnswerKeyOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTicketAnswerKeyOk() (*string, bool)`

GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketAnswerKey

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetTicketAnswerKey(v string)`

SetTicketAnswerKey sets TicketAnswerKey field to given value.

### HasTicketAnswerKey

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasTicketAnswerKey() bool`

HasTicketAnswerKey returns a boolean if a field has been set.

### GetTransactionIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTransactionIdentification() string`

GetTransactionIdentification returns the TransactionIdentification field if non-nil, zero value otherwise.

### GetTransactionIdentificationOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTransactionIdentificationOk() (*string, bool)`

GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetTransactionIdentification(v string)`

SetTransactionIdentification sets TransactionIdentification field to given value.

### HasTransactionIdentification

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasTransactionIdentification() bool`

HasTransactionIdentification returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetTrasactionDateTime

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTrasactionDateTime() string`

GetTrasactionDateTime returns the TrasactionDateTime field if non-nil, zero value otherwise.

### GetTrasactionDateTimeOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTrasactionDateTimeOk() (*string, bool)`

GetTrasactionDateTimeOk returns a tuple with the TrasactionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrasactionDateTime

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetTrasactionDateTime(v string)`

SetTrasactionDateTime sets TrasactionDateTime field to given value.

### HasTrasactionDateTime

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasTrasactionDateTime() bool`

HasTrasactionDateTime returns a boolean if a field has been set.

### GetPOSType

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSType() string`

GetPOSType returns the POSType field if non-nil, zero value otherwise.

### GetPOSTypeOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSTypeOk() (*string, bool)`

GetPOSTypeOk returns a tuple with the POSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSType

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetPOSType(v string)`

SetPOSType sets POSType field to given value.

### HasPOSType

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasPOSType() bool`

HasPOSType returns a boolean if a field has been set.

### GetPOSVersion

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSVersion() string`

GetPOSVersion returns the POSVersion field if non-nil, zero value otherwise.

### GetPOSVersionOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSVersionOk() (*string, bool)`

GetPOSVersionOk returns a tuple with the POSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSVersion

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetPOSVersion(v string)`

SetPOSVersion sets POSVersion field to given value.

### HasPOSVersion

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasPOSVersion() bool`

HasPOSVersion returns a boolean if a field has been set.

### GetPOSAddress

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSAddress() string`

GetPOSAddress returns the POSAddress field if non-nil, zero value otherwise.

### GetPOSAddressOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSAddressOk() (*string, bool)`

GetPOSAddressOk returns a tuple with the POSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSAddress

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetPOSAddress(v string)`

SetPOSAddress sets POSAddress field to given value.

### HasPOSAddress

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasPOSAddress() bool`

HasPOSAddress returns a boolean if a field has been set.

### GetPOSSerial

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSSerial() string`

GetPOSSerial returns the POSSerial field if non-nil, zero value otherwise.

### GetPOSSerialOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSSerialOk() (*string, bool)`

GetPOSSerialOk returns a tuple with the POSSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSSerial

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetPOSSerial(v string)`

SetPOSSerial sets POSSerial field to given value.

### HasPOSSerial

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasPOSSerial() bool`

HasPOSSerial returns a boolean if a field has been set.

### GetPOSGEO

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSGEO() string`

GetPOSGEO returns the POSGEO field if non-nil, zero value otherwise.

### GetPOSGEOOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPOSGEOOk() (*string, bool)`

GetPOSGEOOk returns a tuple with the POSGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSGEO

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetPOSGEO(v string)`

SetPOSGEO sets POSGEO field to given value.

### HasPOSGEO

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasPOSGEO() bool`

HasPOSGEO returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetReadingDeviceVersion

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceVersion() string`

GetReadingDeviceVersion returns the ReadingDeviceVersion field if non-nil, zero value otherwise.

### GetReadingDeviceVersionOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceVersionOk() (*string, bool)`

GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceVersion

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetReadingDeviceVersion(v string)`

SetReadingDeviceVersion sets ReadingDeviceVersion field to given value.

### HasReadingDeviceVersion

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasReadingDeviceVersion() bool`

HasReadingDeviceVersion returns a boolean if a field has been set.

### GetReadingDeviceAddress

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceAddress() string`

GetReadingDeviceAddress returns the ReadingDeviceAddress field if non-nil, zero value otherwise.

### GetReadingDeviceAddressOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceAddressOk() (*string, bool)`

GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceAddress

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetReadingDeviceAddress(v string)`

SetReadingDeviceAddress sets ReadingDeviceAddress field to given value.

### HasReadingDeviceAddress

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasReadingDeviceAddress() bool`

HasReadingDeviceAddress returns a boolean if a field has been set.

### GetReadingDeviceSerial

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceSerial() string`

GetReadingDeviceSerial returns the ReadingDeviceSerial field if non-nil, zero value otherwise.

### GetReadingDeviceSerialOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceSerialOk() (*string, bool)`

GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceSerial

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetReadingDeviceSerial(v string)`

SetReadingDeviceSerial sets ReadingDeviceSerial field to given value.

### HasReadingDeviceSerial

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasReadingDeviceSerial() bool`

HasReadingDeviceSerial returns a boolean if a field has been set.

### GetReadingDeviceGEO

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceGEO() string`

GetReadingDeviceGEO returns the ReadingDeviceGEO field if non-nil, zero value otherwise.

### GetReadingDeviceGEOOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetReadingDeviceGEOOk() (*string, bool)`

GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceGEO

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetReadingDeviceGEO(v string)`

SetReadingDeviceGEO sets ReadingDeviceGEO field to given value.

### HasReadingDeviceGEO

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasReadingDeviceGEO() bool`

HasReadingDeviceGEO returns a boolean if a field has been set.

### GetUserID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetTerminalID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTerminalID() string`

GetTerminalID returns the TerminalID field if non-nil, zero value otherwise.

### GetTerminalIDOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTerminalIDOk() (*string, bool)`

GetTerminalIDOk returns a tuple with the TerminalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetTerminalID(v string)`

SetTerminalID sets TerminalID field to given value.

### HasTerminalID

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasTerminalID() bool`

HasTerminalID returns a boolean if a field has been set.

### GetTerminalTrace

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTerminalTrace() int32`

GetTerminalTrace returns the TerminalTrace field if non-nil, zero value otherwise.

### GetTerminalTraceOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetTerminalTraceOk() (*int32, bool)`

GetTerminalTraceOk returns a tuple with the TerminalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTrace

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetTerminalTrace(v int32)`

SetTerminalTrace sets TerminalTrace field to given value.

### HasTerminalTrace

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasTerminalTrace() bool`

HasTerminalTrace returns a boolean if a field has been set.

### GetSettlementBatchNumber

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetSettlementBatchNumber() int32`

GetSettlementBatchNumber returns the SettlementBatchNumber field if non-nil, zero value otherwise.

### GetSettlementBatchNumberOk

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) GetSettlementBatchNumberOk() (*int32, bool)`

GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBatchNumber

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) SetSettlementBatchNumber(v int32)`

SetSettlementBatchNumber sets SettlementBatchNumber field to given value.

### HasSettlementBatchNumber

`func (o *QueryLineOfBusinessObjectQueryLineOfBusiness) HasSettlementBatchNumber() bool`

HasSettlementBatchNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


