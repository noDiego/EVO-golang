# GetBlockObjectGetBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyIdentification** | **string** | Identificador de la Compañía para la Plataforma de Integración. | 
**SystemIdentification** | **string** | Identificador de Aplicativo para la Plataforma de Integración que usa la Compañía especificada. | 
**BranchIdentification** | Pointer to **string** | Identificador de Aplicativo para la Plataforma de Integración que usa la Compañía especificada. | [optional] 
**POSIdentification** | Pointer to **interface{}** | Identificador del Punto de Venta/POS/Caja que pertenece a la sucursal y companía especificada. | [optional] 
**RequestType** | Pointer to **string** | Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**Block** | Pointer to **string** | ID que identifica a un grupo de transacciones que serán confirmadas o canceladas | [optional] 
**Ticket** | Pointer to **string** | Ticket Digitalizado ( Total o parte del mismo por ejemplo la Firma Digitalizada )    codificado en Base64. | [optional] 
**TicketAnswerKey** | Pointer to **string** | Identificador Unívoco de la Transacción que se quiere Referenciar a la cual pertenece el Ticket Digitalizado. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. Si firma fue capturada previamente y se envía en el requerimiento de la misma Operación Sale, Authorize*, Void, Return, Adjustment, DebtPayment o VoidDebtPayment no es necesario que se envíe este elemento o campo. | [optional] 
**RequestKey** | Pointer to **string** | Identificador del Requermiento obtenido en la respuesta de la operación WalletRequest.             | [optional] 
**ReasonSequenceBreak** | Pointer to **string** | Motivo por el cual se requiere romper la secuencia. | [optional] 
**POSType** | Pointer to **string** | Tipo de punto de venta. | [optional] 
**POSVersion** | Pointer to **string** | Versión del Aplicativo del punto de Venta. | [optional] 
**POSAddress** | Pointer to **string** | Dirección IP de la Caja o POS. | [optional] 
**POSSerial** | Pointer to **string** | Número de serie o identificador unívoco del punto de venta. | [optional] 
**POSGEO** | Pointer to **string** | Coordenadas Geográficas del aplicativo de punto de venta | [optional] 
**ReadingDeviceVersion** | Pointer to **string** | Versión del dispositivo. | [optional] 
**ReadingDeviceType** | Pointer to **string** | Tipo de dispositivo utilizado para ingresar los datos de la tarjeta. En función al dispositivo usado, serán realizadas ciertas verificaciones, por lo que ciertos datos serán requeridos. CustomerKeyboard, utilizado para ingreso manual de tarjeta a través de un portal web, por ejemplo - Keyboard, utilizado para ingreso manual de la tarjeta por parte del vendedor - MagneticStripReader, lector de banda de tarjetas por emulación de teclado, u otro valor configurado  que indentifique el dispositivo que se esta utilizando. | [optional] 
**ReadingDeviceOperatingFrom** | Pointer to **time.Time** | Indica desde cuando se encuentra operativo o encendido el dispositivo | [optional] 
**ReadingDeviceSerial** | Pointer to **string** | Número de serie o identificador unívoco del dispositivo. | [optional] 
**ReadingDeviceGEO** | Pointer to **string** | Coordenadas Geograficas del dispositivo de lectura | [optional] 
**ReadingDeviceAddress** | Pointer to **string** | Dirección IP o MAC Address del dispositivo. | [optional] 
**UserID** | Pointer to **string** | Identificador del usuario que está realizando la Transacción. | [optional] 
**UserName** | Pointer to **string** | Nombre del usuario que está realizando la Transacción. | [optional] 
**OrigBlock** | Pointer to **string** | Identificador de Bloque que se quiere recuperar la información de su composición con la operación GetBlock, Void, Return. | [optional] 
**OrigForeignBlock** | Pointer to **string** | Identificador de Bloque que se quiere recuperar la información de su composición con la operación GetBlock, Void, Return. | [optional] 
**TransactionsRequired** | Pointer to [**[]GetBlockObjectGetBlockTransactionsRequired**](GetBlockObjectGetBlockTransactionsRequired.md) | Permite especificarLas transacciones requeridas y su estado | [optional] 

## Methods

### NewGetBlockObjectGetBlock

`func NewGetBlockObjectGetBlock(companyIdentification string, systemIdentification string, ) *GetBlockObjectGetBlock`

NewGetBlockObjectGetBlock instantiates a new GetBlockObjectGetBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockObjectGetBlockWithDefaults

`func NewGetBlockObjectGetBlockWithDefaults() *GetBlockObjectGetBlock`

NewGetBlockObjectGetBlockWithDefaults instantiates a new GetBlockObjectGetBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *GetBlockObjectGetBlock) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *GetBlockObjectGetBlock) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *GetBlockObjectGetBlock) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetSystemIdentification

`func (o *GetBlockObjectGetBlock) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *GetBlockObjectGetBlock) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *GetBlockObjectGetBlock) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetBranchIdentification

`func (o *GetBlockObjectGetBlock) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *GetBlockObjectGetBlock) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *GetBlockObjectGetBlock) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *GetBlockObjectGetBlock) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *GetBlockObjectGetBlock) GetPOSIdentification() interface{}`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *GetBlockObjectGetBlock) GetPOSIdentificationOk() (*interface{}, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *GetBlockObjectGetBlock) SetPOSIdentification(v interface{})`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *GetBlockObjectGetBlock) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### SetPOSIdentificationNil

`func (o *GetBlockObjectGetBlock) SetPOSIdentificationNil(b bool)`

 SetPOSIdentificationNil sets the value for POSIdentification to be an explicit nil

### UnsetPOSIdentification
`func (o *GetBlockObjectGetBlock) UnsetPOSIdentification()`

UnsetPOSIdentification ensures that no value is present for POSIdentification, not even an explicit nil
### GetRequestType

`func (o *GetBlockObjectGetBlock) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *GetBlockObjectGetBlock) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *GetBlockObjectGetBlock) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *GetBlockObjectGetBlock) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *GetBlockObjectGetBlock) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *GetBlockObjectGetBlock) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *GetBlockObjectGetBlock) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *GetBlockObjectGetBlock) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *GetBlockObjectGetBlock) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetBlockObjectGetBlock) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetBlockObjectGetBlock) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *GetBlockObjectGetBlock) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *GetBlockObjectGetBlock) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *GetBlockObjectGetBlock) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *GetBlockObjectGetBlock) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *GetBlockObjectGetBlock) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *GetBlockObjectGetBlock) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *GetBlockObjectGetBlock) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *GetBlockObjectGetBlock) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *GetBlockObjectGetBlock) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetTicket

`func (o *GetBlockObjectGetBlock) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *GetBlockObjectGetBlock) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *GetBlockObjectGetBlock) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *GetBlockObjectGetBlock) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTicketAnswerKey

`func (o *GetBlockObjectGetBlock) GetTicketAnswerKey() string`

GetTicketAnswerKey returns the TicketAnswerKey field if non-nil, zero value otherwise.

### GetTicketAnswerKeyOk

`func (o *GetBlockObjectGetBlock) GetTicketAnswerKeyOk() (*string, bool)`

GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketAnswerKey

`func (o *GetBlockObjectGetBlock) SetTicketAnswerKey(v string)`

SetTicketAnswerKey sets TicketAnswerKey field to given value.

### HasTicketAnswerKey

`func (o *GetBlockObjectGetBlock) HasTicketAnswerKey() bool`

HasTicketAnswerKey returns a boolean if a field has been set.

### GetRequestKey

`func (o *GetBlockObjectGetBlock) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *GetBlockObjectGetBlock) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *GetBlockObjectGetBlock) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *GetBlockObjectGetBlock) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetReasonSequenceBreak

`func (o *GetBlockObjectGetBlock) GetReasonSequenceBreak() string`

GetReasonSequenceBreak returns the ReasonSequenceBreak field if non-nil, zero value otherwise.

### GetReasonSequenceBreakOk

`func (o *GetBlockObjectGetBlock) GetReasonSequenceBreakOk() (*string, bool)`

GetReasonSequenceBreakOk returns a tuple with the ReasonSequenceBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSequenceBreak

`func (o *GetBlockObjectGetBlock) SetReasonSequenceBreak(v string)`

SetReasonSequenceBreak sets ReasonSequenceBreak field to given value.

### HasReasonSequenceBreak

`func (o *GetBlockObjectGetBlock) HasReasonSequenceBreak() bool`

HasReasonSequenceBreak returns a boolean if a field has been set.

### GetPOSType

`func (o *GetBlockObjectGetBlock) GetPOSType() string`

GetPOSType returns the POSType field if non-nil, zero value otherwise.

### GetPOSTypeOk

`func (o *GetBlockObjectGetBlock) GetPOSTypeOk() (*string, bool)`

GetPOSTypeOk returns a tuple with the POSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSType

`func (o *GetBlockObjectGetBlock) SetPOSType(v string)`

SetPOSType sets POSType field to given value.

### HasPOSType

`func (o *GetBlockObjectGetBlock) HasPOSType() bool`

HasPOSType returns a boolean if a field has been set.

### GetPOSVersion

`func (o *GetBlockObjectGetBlock) GetPOSVersion() string`

GetPOSVersion returns the POSVersion field if non-nil, zero value otherwise.

### GetPOSVersionOk

`func (o *GetBlockObjectGetBlock) GetPOSVersionOk() (*string, bool)`

GetPOSVersionOk returns a tuple with the POSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSVersion

`func (o *GetBlockObjectGetBlock) SetPOSVersion(v string)`

SetPOSVersion sets POSVersion field to given value.

### HasPOSVersion

`func (o *GetBlockObjectGetBlock) HasPOSVersion() bool`

HasPOSVersion returns a boolean if a field has been set.

### GetPOSAddress

`func (o *GetBlockObjectGetBlock) GetPOSAddress() string`

GetPOSAddress returns the POSAddress field if non-nil, zero value otherwise.

### GetPOSAddressOk

`func (o *GetBlockObjectGetBlock) GetPOSAddressOk() (*string, bool)`

GetPOSAddressOk returns a tuple with the POSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSAddress

`func (o *GetBlockObjectGetBlock) SetPOSAddress(v string)`

SetPOSAddress sets POSAddress field to given value.

### HasPOSAddress

`func (o *GetBlockObjectGetBlock) HasPOSAddress() bool`

HasPOSAddress returns a boolean if a field has been set.

### GetPOSSerial

`func (o *GetBlockObjectGetBlock) GetPOSSerial() string`

GetPOSSerial returns the POSSerial field if non-nil, zero value otherwise.

### GetPOSSerialOk

`func (o *GetBlockObjectGetBlock) GetPOSSerialOk() (*string, bool)`

GetPOSSerialOk returns a tuple with the POSSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSSerial

`func (o *GetBlockObjectGetBlock) SetPOSSerial(v string)`

SetPOSSerial sets POSSerial field to given value.

### HasPOSSerial

`func (o *GetBlockObjectGetBlock) HasPOSSerial() bool`

HasPOSSerial returns a boolean if a field has been set.

### GetPOSGEO

`func (o *GetBlockObjectGetBlock) GetPOSGEO() string`

GetPOSGEO returns the POSGEO field if non-nil, zero value otherwise.

### GetPOSGEOOk

`func (o *GetBlockObjectGetBlock) GetPOSGEOOk() (*string, bool)`

GetPOSGEOOk returns a tuple with the POSGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSGEO

`func (o *GetBlockObjectGetBlock) SetPOSGEO(v string)`

SetPOSGEO sets POSGEO field to given value.

### HasPOSGEO

`func (o *GetBlockObjectGetBlock) HasPOSGEO() bool`

HasPOSGEO returns a boolean if a field has been set.

### GetReadingDeviceVersion

`func (o *GetBlockObjectGetBlock) GetReadingDeviceVersion() string`

GetReadingDeviceVersion returns the ReadingDeviceVersion field if non-nil, zero value otherwise.

### GetReadingDeviceVersionOk

`func (o *GetBlockObjectGetBlock) GetReadingDeviceVersionOk() (*string, bool)`

GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceVersion

`func (o *GetBlockObjectGetBlock) SetReadingDeviceVersion(v string)`

SetReadingDeviceVersion sets ReadingDeviceVersion field to given value.

### HasReadingDeviceVersion

`func (o *GetBlockObjectGetBlock) HasReadingDeviceVersion() bool`

HasReadingDeviceVersion returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *GetBlockObjectGetBlock) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *GetBlockObjectGetBlock) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *GetBlockObjectGetBlock) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *GetBlockObjectGetBlock) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *GetBlockObjectGetBlock) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *GetBlockObjectGetBlock) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *GetBlockObjectGetBlock) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *GetBlockObjectGetBlock) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetReadingDeviceSerial

`func (o *GetBlockObjectGetBlock) GetReadingDeviceSerial() string`

GetReadingDeviceSerial returns the ReadingDeviceSerial field if non-nil, zero value otherwise.

### GetReadingDeviceSerialOk

`func (o *GetBlockObjectGetBlock) GetReadingDeviceSerialOk() (*string, bool)`

GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceSerial

`func (o *GetBlockObjectGetBlock) SetReadingDeviceSerial(v string)`

SetReadingDeviceSerial sets ReadingDeviceSerial field to given value.

### HasReadingDeviceSerial

`func (o *GetBlockObjectGetBlock) HasReadingDeviceSerial() bool`

HasReadingDeviceSerial returns a boolean if a field has been set.

### GetReadingDeviceGEO

`func (o *GetBlockObjectGetBlock) GetReadingDeviceGEO() string`

GetReadingDeviceGEO returns the ReadingDeviceGEO field if non-nil, zero value otherwise.

### GetReadingDeviceGEOOk

`func (o *GetBlockObjectGetBlock) GetReadingDeviceGEOOk() (*string, bool)`

GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceGEO

`func (o *GetBlockObjectGetBlock) SetReadingDeviceGEO(v string)`

SetReadingDeviceGEO sets ReadingDeviceGEO field to given value.

### HasReadingDeviceGEO

`func (o *GetBlockObjectGetBlock) HasReadingDeviceGEO() bool`

HasReadingDeviceGEO returns a boolean if a field has been set.

### GetReadingDeviceAddress

`func (o *GetBlockObjectGetBlock) GetReadingDeviceAddress() string`

GetReadingDeviceAddress returns the ReadingDeviceAddress field if non-nil, zero value otherwise.

### GetReadingDeviceAddressOk

`func (o *GetBlockObjectGetBlock) GetReadingDeviceAddressOk() (*string, bool)`

GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceAddress

`func (o *GetBlockObjectGetBlock) SetReadingDeviceAddress(v string)`

SetReadingDeviceAddress sets ReadingDeviceAddress field to given value.

### HasReadingDeviceAddress

`func (o *GetBlockObjectGetBlock) HasReadingDeviceAddress() bool`

HasReadingDeviceAddress returns a boolean if a field has been set.

### GetUserID

`func (o *GetBlockObjectGetBlock) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *GetBlockObjectGetBlock) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *GetBlockObjectGetBlock) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *GetBlockObjectGetBlock) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *GetBlockObjectGetBlock) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GetBlockObjectGetBlock) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GetBlockObjectGetBlock) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GetBlockObjectGetBlock) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetOrigBlock

`func (o *GetBlockObjectGetBlock) GetOrigBlock() string`

GetOrigBlock returns the OrigBlock field if non-nil, zero value otherwise.

### GetOrigBlockOk

`func (o *GetBlockObjectGetBlock) GetOrigBlockOk() (*string, bool)`

GetOrigBlockOk returns a tuple with the OrigBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigBlock

`func (o *GetBlockObjectGetBlock) SetOrigBlock(v string)`

SetOrigBlock sets OrigBlock field to given value.

### HasOrigBlock

`func (o *GetBlockObjectGetBlock) HasOrigBlock() bool`

HasOrigBlock returns a boolean if a field has been set.

### GetOrigForeignBlock

`func (o *GetBlockObjectGetBlock) GetOrigForeignBlock() string`

GetOrigForeignBlock returns the OrigForeignBlock field if non-nil, zero value otherwise.

### GetOrigForeignBlockOk

`func (o *GetBlockObjectGetBlock) GetOrigForeignBlockOk() (*string, bool)`

GetOrigForeignBlockOk returns a tuple with the OrigForeignBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigForeignBlock

`func (o *GetBlockObjectGetBlock) SetOrigForeignBlock(v string)`

SetOrigForeignBlock sets OrigForeignBlock field to given value.

### HasOrigForeignBlock

`func (o *GetBlockObjectGetBlock) HasOrigForeignBlock() bool`

HasOrigForeignBlock returns a boolean if a field has been set.

### GetTransactionsRequired

`func (o *GetBlockObjectGetBlock) GetTransactionsRequired() []GetBlockObjectGetBlockTransactionsRequired`

GetTransactionsRequired returns the TransactionsRequired field if non-nil, zero value otherwise.

### GetTransactionsRequiredOk

`func (o *GetBlockObjectGetBlock) GetTransactionsRequiredOk() (*[]GetBlockObjectGetBlockTransactionsRequired, bool)`

GetTransactionsRequiredOk returns a tuple with the TransactionsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsRequired

`func (o *GetBlockObjectGetBlock) SetTransactionsRequired(v []GetBlockObjectGetBlockTransactionsRequired)`

SetTransactionsRequired sets TransactionsRequired field to given value.

### HasTransactionsRequired

`func (o *GetBlockObjectGetBlock) HasTransactionsRequired() bool`

HasTransactionsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


