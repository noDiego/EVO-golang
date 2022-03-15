# GetBlockResponseObjectGetBlockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | Pointer to **string** | Tipo de Operación que se requirió, contendrá el mismo valor que se recibió en el requerimiento, sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**ResponseActions** | **[]string** | Acciones a realizar por parte del POS y/o PINPAD en base al resultado de la operación que ha sido procesada. Cada uno de estos actions o acciones están concatenadas por comas. Los posibles actions son OK, Approve, Refuse, IssuerCall, Tickets, WithHold, GetCard, UseTerminalToAuthorize, ConfigurationError, SystemError, ResourceError, ProcessError, Completed. | 
**ResponseMessage** | **string** | Descripción del resultado del proceso del requerimiento recibido. Esta descripción es generada por la plataforma, no por el Host que termine resolviendo la transacción. | 
**CompanyIdentification** | Pointer to **string** | ID que identifica la companía desde donde proviene la petición. | [optional] 
**SystemIdentification** | Pointer to **string** | ID que identifica el sistema desde donde proviene la petición. | [optional] 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**ForeignResponseCode** | Pointer to **string** | Código de respuesta para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF. | [optional] 
**ResponseCode** | **int32** | Código de Respuesta Interno de la plataforma, el POS debe actuar por lo que indican las acciones especificadas en ResponseActions y no por el código de respuesta informado en este campo o elemento, pero es una buena práctica que sea persistido por el mismo. | 
**AnswerType** | Pointer to **string** | Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**AnswerKey** | Pointer to **string** | Código de identificación, generado por Plataforma, de la operación realizada | [optional] 
**RequestKey** | Pointer to **string** | ID generado para la identificación por parte del Plataforma de la información generada en la ejecución de un GetCard o un Payment Method. Sera necesario para que un mensaje de Sale, Void o Payment Method identifique el contexto generado y lo utilice para esa operación. | [optional] 
**ServerVersion** | Pointer to **string** | Versión del Servicio de la Plataforma que resolvió el requerimiento. | [optional] 
**ServerAddress** | Pointer to **string** | Dirección IP del Server que atiende el requerimiento. | [optional] 
**ServerInstance** | Pointer to **string** | Instancia de Server que atiende el requerimiento. | [optional] 
**ServerNodeName** | Pointer to **string** | Nombre del Nodo que atendió el requerimiento. | [optional] 
**MessageID** | Pointer to **string** | Identificador Unívoco del Mensaje ( UUID v5 ). | [optional] 
**AdapterInputVersion** | Pointer to **string** | Versión del  Adaptador de Protocolo Entrante que atiende el Requerimiento. | [optional] 
**AdapterInputAddress** | Pointer to **string** | Dirección IP del Adaptador de Protocolo Entrante que atiende el requerimiento. | [optional] 
**AdapterInputNodeName** | Pointer to **string** | Nombre del Nodo del Adaptador de Protocolo Entrante que atiende el requerimiento. | [optional] 
**AdapterOutputVersion** | Pointer to **string** | Versión del  Adaptador de Protocolo Saliente que atiende el Requerimiento. | [optional] 
**AdapterOutputAddress** | Pointer to **string** | Dirección IP  del  Adaptador de Protocolo Saliente que atiende el Requerimiento. | [optional] 
**AdapterOutputNodeName** | Pointer to **string** | Nombre del Nodo  del  Adaptador de Protocolo Saliente que atiende el Requerimiento. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**WasReversePrevious** | Pointer to **int32** | Flag indicador de generación de reversa para la última operación reversable | [optional] 
**WasClosedPreviousBlock** | Pointer to **int32** | Flag indicador de cancelación o confirmación del último bloque de transacciones, previo al nuevo valor recibido | [optional] 
**ReversedAnswerKey** | Pointer to **string** | ID que identifica a la operación que acaba de ser reversada. | [optional] 
**ReversedSequence** | Pointer to **string** | Secuencia de la transacción que fue reversada | [optional] 
**CommittedBlock** | Pointer to **string** | ID del bloque de transacciones que ha sido confirmado de forma automática (es decir, sin recibir un requerimiento de BlockClose). Este escenario se presentará si el Plataforma así se ha configurado para actuar bajo esa circunstancia. | [optional] 
**ReversedBlock** | Pointer to **string** | ID del bloque de transacciones que ha sido cancelado de forma automática (es decir, sin recibir un requerimiento de BlockClose). Este escenario se presentará si el Plataforma así se ha configurado para actuar bajo esa circunstancia. | [optional] 
**RequiredInformation** | Pointer to [**[]DebtPaymentObjectDebtPaymentRequiredInformation**](DebtPaymentObjectDebtPaymentRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**AdditionalInformation** | Pointer to [**[]SaleResponseObjectSaleResponseAdditionalInformation**](SaleResponseObjectSaleResponseAdditionalInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**DisplayResponseMessage** | Pointer to **[]string** | Información adicional/Mensaje promocional/Leyenda de respuesta a mostrar en pantalla en el ticket de la operación. Cada línea de este mensaje será un elemento dentro del array. | [optional] 
**Block** | Pointer to **string** | ID que identifica a un grupo de transacciones que serán confirmadas o canceladas | [optional] 
**TransmitionDateTime** | Pointer to **time.Time** | Fecha y hora de transmision de la operación hacia el host - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**Transactions** | Pointer to **string** | Detalle del bloque, es una array de un resumen de cada transacción que lo compone. | [optional] 
**ForeignBlock** | Pointer to **string** | Identificador Alternativo de un  Bloque Que se informa en las operaciones BlockCancel o BlockClose. | [optional] 
**Configuration** | Pointer to [**SaleResponseObjectSaleResponseConfiguration**](SaleResponseObjectSaleResponseConfiguration.md) |  | [optional] 

## Methods

### NewGetBlockResponseObjectGetBlockResponse

`func NewGetBlockResponseObjectGetBlockResponse(responseActions []string, responseMessage string, responseCode int32, ) *GetBlockResponseObjectGetBlockResponse`

NewGetBlockResponseObjectGetBlockResponse instantiates a new GetBlockResponseObjectGetBlockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockResponseObjectGetBlockResponseWithDefaults

`func NewGetBlockResponseObjectGetBlockResponseWithDefaults() *GetBlockResponseObjectGetBlockResponse`

NewGetBlockResponseObjectGetBlockResponseWithDefaults instantiates a new GetBlockResponseObjectGetBlockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *GetBlockResponseObjectGetBlockResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *GetBlockResponseObjectGetBlockResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *GetBlockResponseObjectGetBlockResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetResponseActions

`func (o *GetBlockResponseObjectGetBlockResponse) GetResponseActions() []string`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetResponseActionsOk() (*[]string, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *GetBlockResponseObjectGetBlockResponse) SetResponseActions(v []string)`

SetResponseActions sets ResponseActions field to given value.


### GetResponseMessage

`func (o *GetBlockResponseObjectGetBlockResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *GetBlockResponseObjectGetBlockResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetCompanyIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetSystemIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *GetBlockResponseObjectGetBlockResponse) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetForeignResponseCode

`func (o *GetBlockResponseObjectGetBlockResponse) GetForeignResponseCode() string`

GetForeignResponseCode returns the ForeignResponseCode field if non-nil, zero value otherwise.

### GetForeignResponseCodeOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetForeignResponseCodeOk() (*string, bool)`

GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignResponseCode

`func (o *GetBlockResponseObjectGetBlockResponse) SetForeignResponseCode(v string)`

SetForeignResponseCode sets ForeignResponseCode field to given value.

### HasForeignResponseCode

`func (o *GetBlockResponseObjectGetBlockResponse) HasForeignResponseCode() bool`

HasForeignResponseCode returns a boolean if a field has been set.

### GetResponseCode

`func (o *GetBlockResponseObjectGetBlockResponse) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *GetBlockResponseObjectGetBlockResponse) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetAnswerType

`func (o *GetBlockResponseObjectGetBlockResponse) GetAnswerType() string`

GetAnswerType returns the AnswerType field if non-nil, zero value otherwise.

### GetAnswerTypeOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetAnswerTypeOk() (*string, bool)`

GetAnswerTypeOk returns a tuple with the AnswerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerType

`func (o *GetBlockResponseObjectGetBlockResponse) SetAnswerType(v string)`

SetAnswerType sets AnswerType field to given value.

### HasAnswerType

`func (o *GetBlockResponseObjectGetBlockResponse) HasAnswerType() bool`

HasAnswerType returns a boolean if a field has been set.

### GetAnswerKey

`func (o *GetBlockResponseObjectGetBlockResponse) GetAnswerKey() string`

GetAnswerKey returns the AnswerKey field if non-nil, zero value otherwise.

### GetAnswerKeyOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetAnswerKeyOk() (*string, bool)`

GetAnswerKeyOk returns a tuple with the AnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerKey

`func (o *GetBlockResponseObjectGetBlockResponse) SetAnswerKey(v string)`

SetAnswerKey sets AnswerKey field to given value.

### HasAnswerKey

`func (o *GetBlockResponseObjectGetBlockResponse) HasAnswerKey() bool`

HasAnswerKey returns a boolean if a field has been set.

### GetRequestKey

`func (o *GetBlockResponseObjectGetBlockResponse) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *GetBlockResponseObjectGetBlockResponse) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.

### HasRequestKey

`func (o *GetBlockResponseObjectGetBlockResponse) HasRequestKey() bool`

HasRequestKey returns a boolean if a field has been set.

### GetServerVersion

`func (o *GetBlockResponseObjectGetBlockResponse) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *GetBlockResponseObjectGetBlockResponse) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *GetBlockResponseObjectGetBlockResponse) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetServerAddress

`func (o *GetBlockResponseObjectGetBlockResponse) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *GetBlockResponseObjectGetBlockResponse) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *GetBlockResponseObjectGetBlockResponse) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetServerInstance

`func (o *GetBlockResponseObjectGetBlockResponse) GetServerInstance() string`

GetServerInstance returns the ServerInstance field if non-nil, zero value otherwise.

### GetServerInstanceOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetServerInstanceOk() (*string, bool)`

GetServerInstanceOk returns a tuple with the ServerInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstance

`func (o *GetBlockResponseObjectGetBlockResponse) SetServerInstance(v string)`

SetServerInstance sets ServerInstance field to given value.

### HasServerInstance

`func (o *GetBlockResponseObjectGetBlockResponse) HasServerInstance() bool`

HasServerInstance returns a boolean if a field has been set.

### GetServerNodeName

`func (o *GetBlockResponseObjectGetBlockResponse) GetServerNodeName() string`

GetServerNodeName returns the ServerNodeName field if non-nil, zero value otherwise.

### GetServerNodeNameOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetServerNodeNameOk() (*string, bool)`

GetServerNodeNameOk returns a tuple with the ServerNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNodeName

`func (o *GetBlockResponseObjectGetBlockResponse) SetServerNodeName(v string)`

SetServerNodeName sets ServerNodeName field to given value.

### HasServerNodeName

`func (o *GetBlockResponseObjectGetBlockResponse) HasServerNodeName() bool`

HasServerNodeName returns a boolean if a field has been set.

### GetMessageID

`func (o *GetBlockResponseObjectGetBlockResponse) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *GetBlockResponseObjectGetBlockResponse) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *GetBlockResponseObjectGetBlockResponse) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetAdapterInputVersion

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterInputVersion() string`

GetAdapterInputVersion returns the AdapterInputVersion field if non-nil, zero value otherwise.

### GetAdapterInputVersionOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterInputVersionOk() (*string, bool)`

GetAdapterInputVersionOk returns a tuple with the AdapterInputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputVersion

`func (o *GetBlockResponseObjectGetBlockResponse) SetAdapterInputVersion(v string)`

SetAdapterInputVersion sets AdapterInputVersion field to given value.

### HasAdapterInputVersion

`func (o *GetBlockResponseObjectGetBlockResponse) HasAdapterInputVersion() bool`

HasAdapterInputVersion returns a boolean if a field has been set.

### GetAdapterInputAddress

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterInputAddress() string`

GetAdapterInputAddress returns the AdapterInputAddress field if non-nil, zero value otherwise.

### GetAdapterInputAddressOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterInputAddressOk() (*string, bool)`

GetAdapterInputAddressOk returns a tuple with the AdapterInputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputAddress

`func (o *GetBlockResponseObjectGetBlockResponse) SetAdapterInputAddress(v string)`

SetAdapterInputAddress sets AdapterInputAddress field to given value.

### HasAdapterInputAddress

`func (o *GetBlockResponseObjectGetBlockResponse) HasAdapterInputAddress() bool`

HasAdapterInputAddress returns a boolean if a field has been set.

### GetAdapterInputNodeName

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterInputNodeName() string`

GetAdapterInputNodeName returns the AdapterInputNodeName field if non-nil, zero value otherwise.

### GetAdapterInputNodeNameOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterInputNodeNameOk() (*string, bool)`

GetAdapterInputNodeNameOk returns a tuple with the AdapterInputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterInputNodeName

`func (o *GetBlockResponseObjectGetBlockResponse) SetAdapterInputNodeName(v string)`

SetAdapterInputNodeName sets AdapterInputNodeName field to given value.

### HasAdapterInputNodeName

`func (o *GetBlockResponseObjectGetBlockResponse) HasAdapterInputNodeName() bool`

HasAdapterInputNodeName returns a boolean if a field has been set.

### GetAdapterOutputVersion

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterOutputVersion() string`

GetAdapterOutputVersion returns the AdapterOutputVersion field if non-nil, zero value otherwise.

### GetAdapterOutputVersionOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterOutputVersionOk() (*string, bool)`

GetAdapterOutputVersionOk returns a tuple with the AdapterOutputVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputVersion

`func (o *GetBlockResponseObjectGetBlockResponse) SetAdapterOutputVersion(v string)`

SetAdapterOutputVersion sets AdapterOutputVersion field to given value.

### HasAdapterOutputVersion

`func (o *GetBlockResponseObjectGetBlockResponse) HasAdapterOutputVersion() bool`

HasAdapterOutputVersion returns a boolean if a field has been set.

### GetAdapterOutputAddress

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterOutputAddress() string`

GetAdapterOutputAddress returns the AdapterOutputAddress field if non-nil, zero value otherwise.

### GetAdapterOutputAddressOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterOutputAddressOk() (*string, bool)`

GetAdapterOutputAddressOk returns a tuple with the AdapterOutputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputAddress

`func (o *GetBlockResponseObjectGetBlockResponse) SetAdapterOutputAddress(v string)`

SetAdapterOutputAddress sets AdapterOutputAddress field to given value.

### HasAdapterOutputAddress

`func (o *GetBlockResponseObjectGetBlockResponse) HasAdapterOutputAddress() bool`

HasAdapterOutputAddress returns a boolean if a field has been set.

### GetAdapterOutputNodeName

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterOutputNodeName() string`

GetAdapterOutputNodeName returns the AdapterOutputNodeName field if non-nil, zero value otherwise.

### GetAdapterOutputNodeNameOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdapterOutputNodeNameOk() (*string, bool)`

GetAdapterOutputNodeNameOk returns a tuple with the AdapterOutputNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOutputNodeName

`func (o *GetBlockResponseObjectGetBlockResponse) SetAdapterOutputNodeName(v string)`

SetAdapterOutputNodeName sets AdapterOutputNodeName field to given value.

### HasAdapterOutputNodeName

`func (o *GetBlockResponseObjectGetBlockResponse) HasAdapterOutputNodeName() bool`

HasAdapterOutputNodeName returns a boolean if a field has been set.

### GetServiceVersion

`func (o *GetBlockResponseObjectGetBlockResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *GetBlockResponseObjectGetBlockResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *GetBlockResponseObjectGetBlockResponse) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *GetBlockResponseObjectGetBlockResponse) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetBlockResponseObjectGetBlockResponse) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *GetBlockResponseObjectGetBlockResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *GetBlockResponseObjectGetBlockResponse) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *GetBlockResponseObjectGetBlockResponse) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *GetBlockResponseObjectGetBlockResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetWasReversePrevious

`func (o *GetBlockResponseObjectGetBlockResponse) GetWasReversePrevious() int32`

GetWasReversePrevious returns the WasReversePrevious field if non-nil, zero value otherwise.

### GetWasReversePreviousOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetWasReversePreviousOk() (*int32, bool)`

GetWasReversePreviousOk returns a tuple with the WasReversePrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasReversePrevious

`func (o *GetBlockResponseObjectGetBlockResponse) SetWasReversePrevious(v int32)`

SetWasReversePrevious sets WasReversePrevious field to given value.

### HasWasReversePrevious

`func (o *GetBlockResponseObjectGetBlockResponse) HasWasReversePrevious() bool`

HasWasReversePrevious returns a boolean if a field has been set.

### GetWasClosedPreviousBlock

`func (o *GetBlockResponseObjectGetBlockResponse) GetWasClosedPreviousBlock() int32`

GetWasClosedPreviousBlock returns the WasClosedPreviousBlock field if non-nil, zero value otherwise.

### GetWasClosedPreviousBlockOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetWasClosedPreviousBlockOk() (*int32, bool)`

GetWasClosedPreviousBlockOk returns a tuple with the WasClosedPreviousBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasClosedPreviousBlock

`func (o *GetBlockResponseObjectGetBlockResponse) SetWasClosedPreviousBlock(v int32)`

SetWasClosedPreviousBlock sets WasClosedPreviousBlock field to given value.

### HasWasClosedPreviousBlock

`func (o *GetBlockResponseObjectGetBlockResponse) HasWasClosedPreviousBlock() bool`

HasWasClosedPreviousBlock returns a boolean if a field has been set.

### GetReversedAnswerKey

`func (o *GetBlockResponseObjectGetBlockResponse) GetReversedAnswerKey() string`

GetReversedAnswerKey returns the ReversedAnswerKey field if non-nil, zero value otherwise.

### GetReversedAnswerKeyOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetReversedAnswerKeyOk() (*string, bool)`

GetReversedAnswerKeyOk returns a tuple with the ReversedAnswerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedAnswerKey

`func (o *GetBlockResponseObjectGetBlockResponse) SetReversedAnswerKey(v string)`

SetReversedAnswerKey sets ReversedAnswerKey field to given value.

### HasReversedAnswerKey

`func (o *GetBlockResponseObjectGetBlockResponse) HasReversedAnswerKey() bool`

HasReversedAnswerKey returns a boolean if a field has been set.

### GetReversedSequence

`func (o *GetBlockResponseObjectGetBlockResponse) GetReversedSequence() string`

GetReversedSequence returns the ReversedSequence field if non-nil, zero value otherwise.

### GetReversedSequenceOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetReversedSequenceOk() (*string, bool)`

GetReversedSequenceOk returns a tuple with the ReversedSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedSequence

`func (o *GetBlockResponseObjectGetBlockResponse) SetReversedSequence(v string)`

SetReversedSequence sets ReversedSequence field to given value.

### HasReversedSequence

`func (o *GetBlockResponseObjectGetBlockResponse) HasReversedSequence() bool`

HasReversedSequence returns a boolean if a field has been set.

### GetCommittedBlock

`func (o *GetBlockResponseObjectGetBlockResponse) GetCommittedBlock() string`

GetCommittedBlock returns the CommittedBlock field if non-nil, zero value otherwise.

### GetCommittedBlockOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetCommittedBlockOk() (*string, bool)`

GetCommittedBlockOk returns a tuple with the CommittedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedBlock

`func (o *GetBlockResponseObjectGetBlockResponse) SetCommittedBlock(v string)`

SetCommittedBlock sets CommittedBlock field to given value.

### HasCommittedBlock

`func (o *GetBlockResponseObjectGetBlockResponse) HasCommittedBlock() bool`

HasCommittedBlock returns a boolean if a field has been set.

### GetReversedBlock

`func (o *GetBlockResponseObjectGetBlockResponse) GetReversedBlock() string`

GetReversedBlock returns the ReversedBlock field if non-nil, zero value otherwise.

### GetReversedBlockOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetReversedBlockOk() (*string, bool)`

GetReversedBlockOk returns a tuple with the ReversedBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedBlock

`func (o *GetBlockResponseObjectGetBlockResponse) SetReversedBlock(v string)`

SetReversedBlock sets ReversedBlock field to given value.

### HasReversedBlock

`func (o *GetBlockResponseObjectGetBlockResponse) HasReversedBlock() bool`

HasReversedBlock returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *GetBlockResponseObjectGetBlockResponse) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetRequiredInformationOk() (*[]DebtPaymentObjectDebtPaymentRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *GetBlockResponseObjectGetBlockResponse) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *GetBlockResponseObjectGetBlockResponse) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *GetBlockResponseObjectGetBlockResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *GetBlockResponseObjectGetBlockResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetDisplayResponseMessage

`func (o *GetBlockResponseObjectGetBlockResponse) GetDisplayResponseMessage() []string`

GetDisplayResponseMessage returns the DisplayResponseMessage field if non-nil, zero value otherwise.

### GetDisplayResponseMessageOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetDisplayResponseMessageOk() (*[]string, bool)`

GetDisplayResponseMessageOk returns a tuple with the DisplayResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayResponseMessage

`func (o *GetBlockResponseObjectGetBlockResponse) SetDisplayResponseMessage(v []string)`

SetDisplayResponseMessage sets DisplayResponseMessage field to given value.

### HasDisplayResponseMessage

`func (o *GetBlockResponseObjectGetBlockResponse) HasDisplayResponseMessage() bool`

HasDisplayResponseMessage returns a boolean if a field has been set.

### GetBlock

`func (o *GetBlockResponseObjectGetBlockResponse) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *GetBlockResponseObjectGetBlockResponse) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *GetBlockResponseObjectGetBlockResponse) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetTransmitionDateTime

`func (o *GetBlockResponseObjectGetBlockResponse) GetTransmitionDateTime() time.Time`

GetTransmitionDateTime returns the TransmitionDateTime field if non-nil, zero value otherwise.

### GetTransmitionDateTimeOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetTransmitionDateTimeOk() (*time.Time, bool)`

GetTransmitionDateTimeOk returns a tuple with the TransmitionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitionDateTime

`func (o *GetBlockResponseObjectGetBlockResponse) SetTransmitionDateTime(v time.Time)`

SetTransmitionDateTime sets TransmitionDateTime field to given value.

### HasTransmitionDateTime

`func (o *GetBlockResponseObjectGetBlockResponse) HasTransmitionDateTime() bool`

HasTransmitionDateTime returns a boolean if a field has been set.

### GetTransactions

`func (o *GetBlockResponseObjectGetBlockResponse) GetTransactions() string`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetTransactionsOk() (*string, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GetBlockResponseObjectGetBlockResponse) SetTransactions(v string)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *GetBlockResponseObjectGetBlockResponse) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetForeignBlock

`func (o *GetBlockResponseObjectGetBlockResponse) GetForeignBlock() string`

GetForeignBlock returns the ForeignBlock field if non-nil, zero value otherwise.

### GetForeignBlockOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetForeignBlockOk() (*string, bool)`

GetForeignBlockOk returns a tuple with the ForeignBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignBlock

`func (o *GetBlockResponseObjectGetBlockResponse) SetForeignBlock(v string)`

SetForeignBlock sets ForeignBlock field to given value.

### HasForeignBlock

`func (o *GetBlockResponseObjectGetBlockResponse) HasForeignBlock() bool`

HasForeignBlock returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetBlockResponseObjectGetBlockResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetBlockResponseObjectGetBlockResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetBlockResponseObjectGetBlockResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetBlockResponseObjectGetBlockResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


