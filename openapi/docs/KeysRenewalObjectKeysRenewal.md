# KeysRenewalObjectKeysRenewal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyIdentification** | **string** | Identificador de la Compañía para la Plataforma de Integración. | 
**SystemIdentification** | **string** | Identificador de Aplicativo para la Plataforma de Integración que usa la Compañía especificada. | 
**BranchIdentification** | Pointer to **string** | Identificador de Aplicativo para la Plataforma de Integración que usa la Compañía especificada. | [optional] 
**POSIdentification** | Pointer to **string** | Identificador del Punto de Venta/POS/Caja que pertenece a la sucursal y companía especificada. | [optional] 
**RequestType** | Pointer to **string** | Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
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
**Timeout** | Pointer to **float32** | Tiempo de espera que el POS espera al PINPAD para obtener la respuesta al requerimiento.  | [optional] 
**Ticket** | Pointer to **string** | Ticket Digitalizado ( Total o parte del mismo por ejemplo la Firma Digitalizada )    codificado en Base64. | [optional] 
**TicketAnswerKey** | Pointer to **string** | Identificador Unívoco de la Transacción que se quiere Referenciar a la cual pertenece el Ticket Digitalizado. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. Si la firma fue capturada previamente y se envía en el requerimiento de la misma Operación Sale, Authorize*, Void, Return, Adjustment, DebtPayment, VoidDebtPayment o Enrollment no es necesario que se envíe este elemento o campo. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma, si la última transacción que realizó el POS no está confirmada y está aprobada, genera una reversa de la misma. | [optional] 
**RequiredInformation** | Pointer to [**[]SaleObjectSaleRequiredInformation**](SaleObjectSaleRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**Block** | Pointer to **string** | ID que identifica a un grupo de transacciones que serán confirmadas o canceladas. | [optional] 
**Products** | Pointer to [**[]SaleObjectSaleProducts**](SaleObjectSaleProducts.md) | Detalle de Productos de la Operación. | [optional] 
**PaymentFacilitatorID** | Pointer to **string** | Identificador de facilitador de pagos o Payfac. | [optional] 
**MerchantID** | Pointer to **string** | Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relacionado a cada uno de los planes disponibles. | [optional] 
**TerminalID** | Pointer to **string** | Identificador de Terminal por el cual se envía la Transacción al Host. | [optional] 
**TerminalTrace** | Pointer to **int32** | Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID. | [optional] 
**Configuration** | Pointer to [**SaleResponseObjectSaleResponseConfiguration**](SaleResponseObjectSaleResponseConfiguration.md) |  | [optional] 
**SettlementBatchNumber** | Pointer to **int32** | Para aquellos host que exista el concepto de lote, es el número de lote al cual pertenece la transacción. | [optional] 

## Methods

### NewKeysRenewalObjectKeysRenewal

`func NewKeysRenewalObjectKeysRenewal(companyIdentification string, systemIdentification string, ) *KeysRenewalObjectKeysRenewal`

NewKeysRenewalObjectKeysRenewal instantiates a new KeysRenewalObjectKeysRenewal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeysRenewalObjectKeysRenewalWithDefaults

`func NewKeysRenewalObjectKeysRenewalWithDefaults() *KeysRenewalObjectKeysRenewal`

NewKeysRenewalObjectKeysRenewalWithDefaults instantiates a new KeysRenewalObjectKeysRenewal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *KeysRenewalObjectKeysRenewal) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *KeysRenewalObjectKeysRenewal) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *KeysRenewalObjectKeysRenewal) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetSystemIdentification

`func (o *KeysRenewalObjectKeysRenewal) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *KeysRenewalObjectKeysRenewal) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *KeysRenewalObjectKeysRenewal) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetBranchIdentification

`func (o *KeysRenewalObjectKeysRenewal) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *KeysRenewalObjectKeysRenewal) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *KeysRenewalObjectKeysRenewal) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *KeysRenewalObjectKeysRenewal) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *KeysRenewalObjectKeysRenewal) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *KeysRenewalObjectKeysRenewal) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *KeysRenewalObjectKeysRenewal) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *KeysRenewalObjectKeysRenewal) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetRequestType

`func (o *KeysRenewalObjectKeysRenewal) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *KeysRenewalObjectKeysRenewal) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *KeysRenewalObjectKeysRenewal) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *KeysRenewalObjectKeysRenewal) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetPOSType

`func (o *KeysRenewalObjectKeysRenewal) GetPOSType() string`

GetPOSType returns the POSType field if non-nil, zero value otherwise.

### GetPOSTypeOk

`func (o *KeysRenewalObjectKeysRenewal) GetPOSTypeOk() (*string, bool)`

GetPOSTypeOk returns a tuple with the POSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSType

`func (o *KeysRenewalObjectKeysRenewal) SetPOSType(v string)`

SetPOSType sets POSType field to given value.

### HasPOSType

`func (o *KeysRenewalObjectKeysRenewal) HasPOSType() bool`

HasPOSType returns a boolean if a field has been set.

### GetPOSVersion

`func (o *KeysRenewalObjectKeysRenewal) GetPOSVersion() string`

GetPOSVersion returns the POSVersion field if non-nil, zero value otherwise.

### GetPOSVersionOk

`func (o *KeysRenewalObjectKeysRenewal) GetPOSVersionOk() (*string, bool)`

GetPOSVersionOk returns a tuple with the POSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSVersion

`func (o *KeysRenewalObjectKeysRenewal) SetPOSVersion(v string)`

SetPOSVersion sets POSVersion field to given value.

### HasPOSVersion

`func (o *KeysRenewalObjectKeysRenewal) HasPOSVersion() bool`

HasPOSVersion returns a boolean if a field has been set.

### GetPOSAddress

`func (o *KeysRenewalObjectKeysRenewal) GetPOSAddress() string`

GetPOSAddress returns the POSAddress field if non-nil, zero value otherwise.

### GetPOSAddressOk

`func (o *KeysRenewalObjectKeysRenewal) GetPOSAddressOk() (*string, bool)`

GetPOSAddressOk returns a tuple with the POSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSAddress

`func (o *KeysRenewalObjectKeysRenewal) SetPOSAddress(v string)`

SetPOSAddress sets POSAddress field to given value.

### HasPOSAddress

`func (o *KeysRenewalObjectKeysRenewal) HasPOSAddress() bool`

HasPOSAddress returns a boolean if a field has been set.

### GetPOSSerial

`func (o *KeysRenewalObjectKeysRenewal) GetPOSSerial() string`

GetPOSSerial returns the POSSerial field if non-nil, zero value otherwise.

### GetPOSSerialOk

`func (o *KeysRenewalObjectKeysRenewal) GetPOSSerialOk() (*string, bool)`

GetPOSSerialOk returns a tuple with the POSSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSSerial

`func (o *KeysRenewalObjectKeysRenewal) SetPOSSerial(v string)`

SetPOSSerial sets POSSerial field to given value.

### HasPOSSerial

`func (o *KeysRenewalObjectKeysRenewal) HasPOSSerial() bool`

HasPOSSerial returns a boolean if a field has been set.

### GetPOSGEO

`func (o *KeysRenewalObjectKeysRenewal) GetPOSGEO() string`

GetPOSGEO returns the POSGEO field if non-nil, zero value otherwise.

### GetPOSGEOOk

`func (o *KeysRenewalObjectKeysRenewal) GetPOSGEOOk() (*string, bool)`

GetPOSGEOOk returns a tuple with the POSGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSGEO

`func (o *KeysRenewalObjectKeysRenewal) SetPOSGEO(v string)`

SetPOSGEO sets POSGEO field to given value.

### HasPOSGEO

`func (o *KeysRenewalObjectKeysRenewal) HasPOSGEO() bool`

HasPOSGEO returns a boolean if a field has been set.

### GetReadingDeviceType

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceType() string`

GetReadingDeviceType returns the ReadingDeviceType field if non-nil, zero value otherwise.

### GetReadingDeviceTypeOk

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceTypeOk() (*string, bool)`

GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceType

`func (o *KeysRenewalObjectKeysRenewal) SetReadingDeviceType(v string)`

SetReadingDeviceType sets ReadingDeviceType field to given value.

### HasReadingDeviceType

`func (o *KeysRenewalObjectKeysRenewal) HasReadingDeviceType() bool`

HasReadingDeviceType returns a boolean if a field has been set.

### GetReadingDeviceOperatingFrom

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceOperatingFrom() time.Time`

GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field if non-nil, zero value otherwise.

### GetReadingDeviceOperatingFromOk

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceOperatingFromOk() (*time.Time, bool)`

GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceOperatingFrom

`func (o *KeysRenewalObjectKeysRenewal) SetReadingDeviceOperatingFrom(v time.Time)`

SetReadingDeviceOperatingFrom sets ReadingDeviceOperatingFrom field to given value.

### HasReadingDeviceOperatingFrom

`func (o *KeysRenewalObjectKeysRenewal) HasReadingDeviceOperatingFrom() bool`

HasReadingDeviceOperatingFrom returns a boolean if a field has been set.

### GetReadingDeviceVersion

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceVersion() string`

GetReadingDeviceVersion returns the ReadingDeviceVersion field if non-nil, zero value otherwise.

### GetReadingDeviceVersionOk

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceVersionOk() (*string, bool)`

GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceVersion

`func (o *KeysRenewalObjectKeysRenewal) SetReadingDeviceVersion(v string)`

SetReadingDeviceVersion sets ReadingDeviceVersion field to given value.

### HasReadingDeviceVersion

`func (o *KeysRenewalObjectKeysRenewal) HasReadingDeviceVersion() bool`

HasReadingDeviceVersion returns a boolean if a field has been set.

### GetReadingDeviceAddress

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceAddress() string`

GetReadingDeviceAddress returns the ReadingDeviceAddress field if non-nil, zero value otherwise.

### GetReadingDeviceAddressOk

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceAddressOk() (*string, bool)`

GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceAddress

`func (o *KeysRenewalObjectKeysRenewal) SetReadingDeviceAddress(v string)`

SetReadingDeviceAddress sets ReadingDeviceAddress field to given value.

### HasReadingDeviceAddress

`func (o *KeysRenewalObjectKeysRenewal) HasReadingDeviceAddress() bool`

HasReadingDeviceAddress returns a boolean if a field has been set.

### GetReadingDeviceSerial

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceSerial() string`

GetReadingDeviceSerial returns the ReadingDeviceSerial field if non-nil, zero value otherwise.

### GetReadingDeviceSerialOk

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceSerialOk() (*string, bool)`

GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceSerial

`func (o *KeysRenewalObjectKeysRenewal) SetReadingDeviceSerial(v string)`

SetReadingDeviceSerial sets ReadingDeviceSerial field to given value.

### HasReadingDeviceSerial

`func (o *KeysRenewalObjectKeysRenewal) HasReadingDeviceSerial() bool`

HasReadingDeviceSerial returns a boolean if a field has been set.

### GetReadingDeviceGEO

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceGEO() string`

GetReadingDeviceGEO returns the ReadingDeviceGEO field if non-nil, zero value otherwise.

### GetReadingDeviceGEOOk

`func (o *KeysRenewalObjectKeysRenewal) GetReadingDeviceGEOOk() (*string, bool)`

GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingDeviceGEO

`func (o *KeysRenewalObjectKeysRenewal) SetReadingDeviceGEO(v string)`

SetReadingDeviceGEO sets ReadingDeviceGEO field to given value.

### HasReadingDeviceGEO

`func (o *KeysRenewalObjectKeysRenewal) HasReadingDeviceGEO() bool`

HasReadingDeviceGEO returns a boolean if a field has been set.

### GetUserID

`func (o *KeysRenewalObjectKeysRenewal) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *KeysRenewalObjectKeysRenewal) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *KeysRenewalObjectKeysRenewal) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *KeysRenewalObjectKeysRenewal) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *KeysRenewalObjectKeysRenewal) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *KeysRenewalObjectKeysRenewal) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *KeysRenewalObjectKeysRenewal) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *KeysRenewalObjectKeysRenewal) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetTimeout

`func (o *KeysRenewalObjectKeysRenewal) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *KeysRenewalObjectKeysRenewal) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *KeysRenewalObjectKeysRenewal) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *KeysRenewalObjectKeysRenewal) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTicket

`func (o *KeysRenewalObjectKeysRenewal) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *KeysRenewalObjectKeysRenewal) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *KeysRenewalObjectKeysRenewal) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *KeysRenewalObjectKeysRenewal) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTicketAnswerKey

`func (o *KeysRenewalObjectKeysRenewal) GetTicketAnswerKey() string`

GetTicketAnswerKey returns the TicketAnswerKey field if non-nil, zero value otherwise.

### GetTicketAnswerKeyOk

`func (o *KeysRenewalObjectKeysRenewal) GetTicketAnswerKeyOk() (*string, bool)`

GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketAnswerKey

`func (o *KeysRenewalObjectKeysRenewal) SetTicketAnswerKey(v string)`

SetTicketAnswerKey sets TicketAnswerKey field to given value.

### HasTicketAnswerKey

`func (o *KeysRenewalObjectKeysRenewal) HasTicketAnswerKey() bool`

HasTicketAnswerKey returns a boolean if a field has been set.

### GetServiceVersion

`func (o *KeysRenewalObjectKeysRenewal) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *KeysRenewalObjectKeysRenewal) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *KeysRenewalObjectKeysRenewal) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *KeysRenewalObjectKeysRenewal) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *KeysRenewalObjectKeysRenewal) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *KeysRenewalObjectKeysRenewal) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *KeysRenewalObjectKeysRenewal) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *KeysRenewalObjectKeysRenewal) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *KeysRenewalObjectKeysRenewal) GetRequiredInformation() []SaleObjectSaleRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *KeysRenewalObjectKeysRenewal) GetRequiredInformationOk() (*[]SaleObjectSaleRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *KeysRenewalObjectKeysRenewal) SetRequiredInformation(v []SaleObjectSaleRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *KeysRenewalObjectKeysRenewal) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetBlock

`func (o *KeysRenewalObjectKeysRenewal) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *KeysRenewalObjectKeysRenewal) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *KeysRenewalObjectKeysRenewal) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *KeysRenewalObjectKeysRenewal) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetProducts

`func (o *KeysRenewalObjectKeysRenewal) GetProducts() []SaleObjectSaleProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *KeysRenewalObjectKeysRenewal) GetProductsOk() (*[]SaleObjectSaleProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *KeysRenewalObjectKeysRenewal) SetProducts(v []SaleObjectSaleProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *KeysRenewalObjectKeysRenewal) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPaymentFacilitatorID

`func (o *KeysRenewalObjectKeysRenewal) GetPaymentFacilitatorID() string`

GetPaymentFacilitatorID returns the PaymentFacilitatorID field if non-nil, zero value otherwise.

### GetPaymentFacilitatorIDOk

`func (o *KeysRenewalObjectKeysRenewal) GetPaymentFacilitatorIDOk() (*string, bool)`

GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFacilitatorID

`func (o *KeysRenewalObjectKeysRenewal) SetPaymentFacilitatorID(v string)`

SetPaymentFacilitatorID sets PaymentFacilitatorID field to given value.

### HasPaymentFacilitatorID

`func (o *KeysRenewalObjectKeysRenewal) HasPaymentFacilitatorID() bool`

HasPaymentFacilitatorID returns a boolean if a field has been set.

### GetMerchantID

`func (o *KeysRenewalObjectKeysRenewal) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *KeysRenewalObjectKeysRenewal) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *KeysRenewalObjectKeysRenewal) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *KeysRenewalObjectKeysRenewal) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetTerminalID

`func (o *KeysRenewalObjectKeysRenewal) GetTerminalID() string`

GetTerminalID returns the TerminalID field if non-nil, zero value otherwise.

### GetTerminalIDOk

`func (o *KeysRenewalObjectKeysRenewal) GetTerminalIDOk() (*string, bool)`

GetTerminalIDOk returns a tuple with the TerminalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalID

`func (o *KeysRenewalObjectKeysRenewal) SetTerminalID(v string)`

SetTerminalID sets TerminalID field to given value.

### HasTerminalID

`func (o *KeysRenewalObjectKeysRenewal) HasTerminalID() bool`

HasTerminalID returns a boolean if a field has been set.

### GetTerminalTrace

`func (o *KeysRenewalObjectKeysRenewal) GetTerminalTrace() int32`

GetTerminalTrace returns the TerminalTrace field if non-nil, zero value otherwise.

### GetTerminalTraceOk

`func (o *KeysRenewalObjectKeysRenewal) GetTerminalTraceOk() (*int32, bool)`

GetTerminalTraceOk returns a tuple with the TerminalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalTrace

`func (o *KeysRenewalObjectKeysRenewal) SetTerminalTrace(v int32)`

SetTerminalTrace sets TerminalTrace field to given value.

### HasTerminalTrace

`func (o *KeysRenewalObjectKeysRenewal) HasTerminalTrace() bool`

HasTerminalTrace returns a boolean if a field has been set.

### GetConfiguration

`func (o *KeysRenewalObjectKeysRenewal) GetConfiguration() SaleResponseObjectSaleResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *KeysRenewalObjectKeysRenewal) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *KeysRenewalObjectKeysRenewal) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *KeysRenewalObjectKeysRenewal) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetSettlementBatchNumber

`func (o *KeysRenewalObjectKeysRenewal) GetSettlementBatchNumber() int32`

GetSettlementBatchNumber returns the SettlementBatchNumber field if non-nil, zero value otherwise.

### GetSettlementBatchNumberOk

`func (o *KeysRenewalObjectKeysRenewal) GetSettlementBatchNumberOk() (*int32, bool)`

GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBatchNumber

`func (o *KeysRenewalObjectKeysRenewal) SetSettlementBatchNumber(v int32)`

SetSettlementBatchNumber sets SettlementBatchNumber field to given value.

### HasSettlementBatchNumber

`func (o *KeysRenewalObjectKeysRenewal) HasSettlementBatchNumber() bool`

HasSettlementBatchNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


