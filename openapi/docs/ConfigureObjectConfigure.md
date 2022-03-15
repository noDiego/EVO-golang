# ConfigureObjectConfigure

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
**POSOrDeviceActions** | Pointer to **[]string** | Lista de Acciones que debe ejecutar el POS o el Dispositvo para el caso que este Plan sea seleccionado.  Acciones para el Device &lt;b&gt;RequestPIN&lt;/b&gt; | [optional] 
**OperationMode** | Pointer to **string** | Diferentes modos de ingresos para los cuales el plan puede ser utilizado | [optional] 
**OperationModeDescription** | Pointer to **string** | Descripción del campo OperationMode | [optional] 
**Operations** | Pointer to [**[]ConfigureObjectConfigureOperations**](ConfigureObjectConfigureOperations.md) | Operaciones y su configuración. | [optional] 
**Tables** | Pointer to [**[]ConfigureObjectConfigureTables**](ConfigureObjectConfigureTables.md) | Tablas de  configuración. | [optional] 
**Files** | Pointer to [**[]ConfigureObjectConfigureFiles**](ConfigureObjectConfigureFiles.md) | Archivos de  configuración, Contenido Multimedia, Aplicaciones, Firmware. | [optional] 

## Methods

### NewConfigureObjectConfigure

`func NewConfigureObjectConfigure(companyIdentification string, systemIdentification string, ) *ConfigureObjectConfigure`

NewConfigureObjectConfigure instantiates a new ConfigureObjectConfigure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureObjectConfigureWithDefaults

`func NewConfigureObjectConfigureWithDefaults() *ConfigureObjectConfigure`

NewConfigureObjectConfigureWithDefaults instantiates a new ConfigureObjectConfigure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *ConfigureObjectConfigure) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *ConfigureObjectConfigure) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *ConfigureObjectConfigure) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetSystemIdentification

`func (o *ConfigureObjectConfigure) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *ConfigureObjectConfigure) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *ConfigureObjectConfigure) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetBranchIdentification

`func (o *ConfigureObjectConfigure) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *ConfigureObjectConfigure) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *ConfigureObjectConfigure) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *ConfigureObjectConfigure) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *ConfigureObjectConfigure) GetPOSIdentification() interface{}`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *ConfigureObjectConfigure) GetPOSIdentificationOk() (*interface{}, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *ConfigureObjectConfigure) SetPOSIdentification(v interface{})`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *ConfigureObjectConfigure) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### SetPOSIdentificationNil

`func (o *ConfigureObjectConfigure) SetPOSIdentificationNil(b bool)`

 SetPOSIdentificationNil sets the value for POSIdentification to be an explicit nil

### UnsetPOSIdentification
`func (o *ConfigureObjectConfigure) UnsetPOSIdentification()`

UnsetPOSIdentification ensures that no value is present for POSIdentification, not even an explicit nil
### GetRequestType

`func (o *ConfigureObjectConfigure) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ConfigureObjectConfigure) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ConfigureObjectConfigure) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ConfigureObjectConfigure) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetServiceVersion

`func (o *ConfigureObjectConfigure) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ConfigureObjectConfigure) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ConfigureObjectConfigure) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ConfigureObjectConfigure) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *ConfigureObjectConfigure) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ConfigureObjectConfigure) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ConfigureObjectConfigure) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *ConfigureObjectConfigure) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *ConfigureObjectConfigure) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ConfigureObjectConfigure) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ConfigureObjectConfigure) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ConfigureObjectConfigure) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetBlock

`func (o *ConfigureObjectConfigure) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *ConfigureObjectConfigure) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *ConfigureObjectConfigure) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *ConfigureObjectConfigure) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetTicket

`func (o *ConfigureObjectConfigure) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *ConfigureObjectConfigure) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *ConfigureObjectConfigure) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *ConfigureObjectConfigure) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTicketAnswerKey

`func (o *ConfigureObjectConfigure) GetTicketAnswerKey() string`

GetTicketAnswerKey returns the TicketAnswerKey field if non-nil, zero value otherwise.

### GetTicketAnswerKeyOk

`func (o *ConfigureObjectConfigure) GetTicketAnswerKeyOk() (*string, bool)`

GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketAnswerKey

`func (o *ConfigureObjectConfigure) SetTicketAnswerKey(v string)`

SetTicketAnswerKey sets TicketAnswerKey field to given value.

### HasTicketAnswerKey

`func (o *ConfigureObjectConfigure) HasTicketAnswerKey() bool`

HasTicketAnswerKey returns a boolean if a field has been set.

### GetReasonSequenceBreak

`func (o *ConfigureObjectConfigure) GetReasonSequenceBreak() string`

GetReasonSequenceBreak returns the ReasonSequenceBreak field if non-nil, zero value otherwise.

### GetReasonSequenceBreakOk

`func (o *ConfigureObjectConfigure) GetReasonSequenceBreakOk() (*string, bool)`

GetReasonSequenceBreakOk returns a tuple with the ReasonSequenceBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSequenceBreak

`func (o *ConfigureObjectConfigure) SetReasonSequenceBreak(v string)`

SetReasonSequenceBreak sets ReasonSequenceBreak field to given value.

### HasReasonSequenceBreak

`func (o *ConfigureObjectConfigure) HasReasonSequenceBreak() bool`

HasReasonSequenceBreak returns a boolean if a field has been set.

### GetPOSType

`func (o *ConfigureObjectConfigure) GetPOSType() string`

GetPOSType returns the POSType field if non-nil, zero value otherwise.

### GetPOSTypeOk

`func (o *ConfigureObjectConfigure) GetPOSTypeOk() (*string, bool)`

GetPOSTypeOk returns a tuple with the POSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSType

`func (o *ConfigureObjectConfigure) SetPOSType(v string)`

SetPOSType sets POSType field to given value.

### HasPOSType

`func (o *ConfigureObjectConfigure) HasPOSType() bool`

HasPOSType returns a boolean if a field has been set.

### GetPOSVersion

`func (o *ConfigureObjectConfigure) GetPOSVersion() string`

GetPOSVersion returns the POSVersion field if non-nil, zero value otherwise.

### GetPOSVersionOk

`func (o *ConfigureObjectConfigure) GetPOSVersionOk() (*string, bool)`

GetPOSVersionOk returns a tuple with the POSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSVersion

`func (o *ConfigureObjectConfigure) SetPOSVersion(v string)`

SetPOSVersion sets POSVersion field to given value.

### HasPOSVersion

`func (o *ConfigureObjectConfigure) HasPOSVersion() bool`

HasPOSVersion returns a boolean if a field has been set.

### GetPOSAddress

`func (o *ConfigureObjectConfigure) GetPOSAddress() string`

GetPOSAddress returns the POSAddress field if non-nil, zero value otherwise.

### GetPOSAddressOk

`func (o *ConfigureObjectConfigure) GetPOSAddressOk() (*string, bool)`

GetPOSAddressOk returns a tuple with the POSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSAddress

`func (o *ConfigureObjectConfigure) SetPOSAddress(v string)`

SetPOSAddress sets POSAddress field to given value.

### HasPOSAddress

`func (o *ConfigureObjectConfigure) HasPOSAddress() bool`

HasPOSAddress returns a boolean if a field has been set.

### GetPOSSerial

`func (o *ConfigureObjectConfigure) GetPOSSerial() string`

GetPOSSerial returns the POSSerial field if non-nil, zero value otherwise.

### GetPOSSerialOk

`func (o *ConfigureObjectConfigure) GetPOSSerialOk() (*string, bool)`

GetPOSSerialOk returns a tuple with the POSSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSSerial

`func (o *ConfigureObjectConfigure) SetPOSSerial(v string)`

SetPOSSerial sets POSSerial field to given value.

### HasPOSSerial

`func (o *ConfigureObjectConfigure) HasPOSSerial() bool`

HasPOSSerial returns a boolean if a field has been set.

### GetPOSGEO

`func (o *ConfigureObjectConfigure) GetPOSGEO() string`

GetPOSGEO returns the POSGEO field if non-nil, zero value otherwise.

### GetPOSGEOOk

`func (o *ConfigureObjectConfigure) GetPOSGEOOk() (*string, bool)`

GetPOSGEOOk returns a tuple with the POSGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSGEO

`func (o *ConfigureObjectConfigure) SetPOSGEO(v string)`

SetPOSGEO sets POSGEO field to given value.

### HasPOSGEO

`func (o *ConfigureObjectConfigure) HasPOSGEO() bool`

HasPOSGEO returns a boolean if a field has been set.

### GetReadingDeviceVersion

`func (o *ConfigureObjectConfigure) GetReadingDeviceVersion() string`

GetReadingDeviceVersion returns the ReadingDeviceVersion field if non-nil, zero value otherwise.

### GetReadingDeviceVersionOk

`func (o *ConfigureObjectConfigure) GetReadingDeviceVersionOk() (*string, bool)`

GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceVersion

`func (o *ConfigureObjectConfigure) SetReadingDeviceVersion(v string)`

SetReadingDeviceVersion sets ReadingDeviceVersion field to given value.

### HasReadingDeviceVersion

`func (o *ConfigureObjectConfigure) HasReadingDeviceVersion() bool`

HasReadingDeviceVersion returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *ConfigureObjectConfigure) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *ConfigureObjectConfigure) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *ConfigureObjectConfigure) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *ConfigureObjectConfigure) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *ConfigureObjectConfigure) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *ConfigureObjectConfigure) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *ConfigureObjectConfigure) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *ConfigureObjectConfigure) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetReadingDeviceSerial

`func (o *ConfigureObjectConfigure) GetReadingDeviceSerial() string`

GetReadingDeviceSerial returns the ReadingDeviceSerial field if non-nil, zero value otherwise.

### GetReadingDeviceSerialOk

`func (o *ConfigureObjectConfigure) GetReadingDeviceSerialOk() (*string, bool)`

GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceSerial

`func (o *ConfigureObjectConfigure) SetReadingDeviceSerial(v string)`

SetReadingDeviceSerial sets ReadingDeviceSerial field to given value.

### HasReadingDeviceSerial

`func (o *ConfigureObjectConfigure) HasReadingDeviceSerial() bool`

HasReadingDeviceSerial returns a boolean if a field has been set.

### GetReadingDeviceGEO

`func (o *ConfigureObjectConfigure) GetReadingDeviceGEO() string`

GetReadingDeviceGEO returns the ReadingDeviceGEO field if non-nil, zero value otherwise.

### GetReadingDeviceGEOOk

`func (o *ConfigureObjectConfigure) GetReadingDeviceGEOOk() (*string, bool)`

GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceGEO

`func (o *ConfigureObjectConfigure) SetReadingDeviceGEO(v string)`

SetReadingDeviceGEO sets ReadingDeviceGEO field to given value.

### HasReadingDeviceGEO

`func (o *ConfigureObjectConfigure) HasReadingDeviceGEO() bool`

HasReadingDeviceGEO returns a boolean if a field has been set.

### GetReadingDeviceAddress

`func (o *ConfigureObjectConfigure) GetReadingDeviceAddress() string`

GetReadingDeviceAddress returns the ReadingDeviceAddress field if non-nil, zero value otherwise.

### GetReadingDeviceAddressOk

`func (o *ConfigureObjectConfigure) GetReadingDeviceAddressOk() (*string, bool)`

GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceAddress

`func (o *ConfigureObjectConfigure) SetReadingDeviceAddress(v string)`

SetReadingDeviceAddress sets ReadingDeviceAddress field to given value.

### HasReadingDeviceAddress

`func (o *ConfigureObjectConfigure) HasReadingDeviceAddress() bool`

HasReadingDeviceAddress returns a boolean if a field has been set.

### GetUserID

`func (o *ConfigureObjectConfigure) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *ConfigureObjectConfigure) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *ConfigureObjectConfigure) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *ConfigureObjectConfigure) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *ConfigureObjectConfigure) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ConfigureObjectConfigure) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ConfigureObjectConfigure) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ConfigureObjectConfigure) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPOSOrDeviceActions

`func (o *ConfigureObjectConfigure) GetPOSOrDeviceActions() []string`

GetPOSOrDeviceActions returns the POSOrDeviceActions field if non-nil, zero value otherwise.

### GetPOSOrDeviceActionsOk

`func (o *ConfigureObjectConfigure) GetPOSOrDeviceActionsOk() (*[]string, bool)`

GetPOSOrDeviceActionsOk returns a tuple with the POSOrDeviceActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSOrDeviceActions

`func (o *ConfigureObjectConfigure) SetPOSOrDeviceActions(v []string)`

SetPOSOrDeviceActions sets POSOrDeviceActions field to given value.

### HasPOSOrDeviceActions

`func (o *ConfigureObjectConfigure) HasPOSOrDeviceActions() bool`

HasPOSOrDeviceActions returns a boolean if a field has been set.

### GetOperationMode

`func (o *ConfigureObjectConfigure) GetOperationMode() string`

GetOperationMode returns the OperationMode field if non-nil, zero value otherwise.

### GetOperationModeOk

`func (o *ConfigureObjectConfigure) GetOperationModeOk() (*string, bool)`

GetOperationModeOk returns a tuple with the OperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMode

`func (o *ConfigureObjectConfigure) SetOperationMode(v string)`

SetOperationMode sets OperationMode field to given value.

### HasOperationMode

`func (o *ConfigureObjectConfigure) HasOperationMode() bool`

HasOperationMode returns a boolean if a field has been set.

### GetOperationModeDescription

`func (o *ConfigureObjectConfigure) GetOperationModeDescription() string`

GetOperationModeDescription returns the OperationModeDescription field if non-nil, zero value otherwise.

### GetOperationModeDescriptionOk

`func (o *ConfigureObjectConfigure) GetOperationModeDescriptionOk() (*string, bool)`

GetOperationModeDescriptionOk returns a tuple with the OperationModeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationModeDescription

`func (o *ConfigureObjectConfigure) SetOperationModeDescription(v string)`

SetOperationModeDescription sets OperationModeDescription field to given value.

### HasOperationModeDescription

`func (o *ConfigureObjectConfigure) HasOperationModeDescription() bool`

HasOperationModeDescription returns a boolean if a field has been set.

### GetOperations

`func (o *ConfigureObjectConfigure) GetOperations() []ConfigureObjectConfigureOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ConfigureObjectConfigure) GetOperationsOk() (*[]ConfigureObjectConfigureOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ConfigureObjectConfigure) SetOperations(v []ConfigureObjectConfigureOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *ConfigureObjectConfigure) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetTables

`func (o *ConfigureObjectConfigure) GetTables() []ConfigureObjectConfigureTables`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *ConfigureObjectConfigure) GetTablesOk() (*[]ConfigureObjectConfigureTables, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *ConfigureObjectConfigure) SetTables(v []ConfigureObjectConfigureTables)`

SetTables sets Tables field to given value.

### HasTables

`func (o *ConfigureObjectConfigure) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetFiles

`func (o *ConfigureObjectConfigure) GetFiles() []ConfigureObjectConfigureFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ConfigureObjectConfigure) GetFilesOk() (*[]ConfigureObjectConfigureFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ConfigureObjectConfigure) SetFiles(v []ConfigureObjectConfigureFiles)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ConfigureObjectConfigure) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


