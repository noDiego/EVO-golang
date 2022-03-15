# OrderInitialResponseObjectOrderInitialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | Pointer to **string** | Tipo de Operación que se requirió, contendrá el mismo valor que se recibió en el requerimiento, sobre formatos que no soportan elementos complejos o compuestos. | [optional] 
**ResponseActions** | **[]string** | Acciones a realizar por parte del POS y/o PINPAD en base al resultado de la operación que ha sido procesada. Cada uno de estos actions o acciones están concatenadas por comas. Los posibles actions son OK, Approve, Refuse, IssuerCall, Tickets, WithHold, GetCard, UseTerminalToAuthorize, ConfigurationError, SystemError, ResourceError, ProcessError, Completed. | 
**ResponseMessage** | **string** | Descripción del resultado del proceso del requerimiento recibido. Esta descripción es generada por la plataforma, no por el Host que termine resolviendo la transacción. | 
**SystemIdentification** | Pointer to **string** | ID que identifica el sistema desde donde proviene la petición. | [optional] 
**CompanyIdentification** | Pointer to **string** | ID que identifica la companía desde donde proviene la petición. | [optional] 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**ForeignResponseCode** | Pointer to **string** | Código de respuesta para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF. | [optional] 
**ResponseCode** | **int32** | Código de Respuesta Interno de la plataforma, el POS debe actuar por lo que indican las acciones especificadas en ResponseActions y no por el código de respuesta informado en este campo o elemento, pero es una buena práctica que sea persistido por el mismo. | 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**AdditionalInformation** | Pointer to [**[]SaleResponseObjectSaleResponseAdditionalInformation**](SaleResponseObjectSaleResponseAdditionalInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**RequiredInformation** | Pointer to [**[]DebtPaymentObjectDebtPaymentRequiredInformation**](DebtPaymentObjectDebtPaymentRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**InitialIdentification** | Pointer to **string** | ID privado del contexto de pago generado. Este elemento es equivalente al llamado RequestKey. | [optional] 
**InitialToken** | Pointer to **string** | ID publico del contexto de pago generado, mediante el cual se lo podra referenciar desde el formulario de pago | [optional] 
**CustomerRedirectAddress** | Pointer to **string** | Elementos a adiciónar a la URL a la cual se redigira para mostrar el formulario de pago | [optional] 
**Configuration** | Pointer to [**SaleResponseObjectSaleResponseConfiguration**](SaleResponseObjectSaleResponseConfiguration.md) |  | [optional] 

## Methods

### NewOrderInitialResponseObjectOrderInitialResponse

`func NewOrderInitialResponseObjectOrderInitialResponse(responseActions []string, responseMessage string, responseCode int32, ) *OrderInitialResponseObjectOrderInitialResponse`

NewOrderInitialResponseObjectOrderInitialResponse instantiates a new OrderInitialResponseObjectOrderInitialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderInitialResponseObjectOrderInitialResponseWithDefaults

`func NewOrderInitialResponseObjectOrderInitialResponseWithDefaults() *OrderInitialResponseObjectOrderInitialResponse`

NewOrderInitialResponseObjectOrderInitialResponseWithDefaults instantiates a new OrderInitialResponseObjectOrderInitialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetResponseActions

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetResponseActions() []string`

GetResponseActions returns the ResponseActions field if non-nil, zero value otherwise.

### GetResponseActionsOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetResponseActionsOk() (*[]string, bool)`

GetResponseActionsOk returns a tuple with the ResponseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseActions

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetResponseActions(v []string)`

SetResponseActions sets ResponseActions field to given value.


### GetResponseMessage

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetSystemIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetForeignResponseCode

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetForeignResponseCode() string`

GetForeignResponseCode returns the ForeignResponseCode field if non-nil, zero value otherwise.

### GetForeignResponseCodeOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetForeignResponseCodeOk() (*string, bool)`

GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignResponseCode

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetForeignResponseCode(v string)`

SetForeignResponseCode sets ForeignResponseCode field to given value.

### HasForeignResponseCode

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasForeignResponseCode() bool`

HasForeignResponseCode returns a boolean if a field has been set.

### GetResponseCode

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetServiceVersion

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetAdditionalInformationOk() (*[]SaleResponseObjectSaleResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetRequiredInformationOk() (*[]DebtPaymentObjectDebtPaymentRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetInitialIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetInitialIdentification() string`

GetInitialIdentification returns the InitialIdentification field if non-nil, zero value otherwise.

### GetInitialIdentificationOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetInitialIdentificationOk() (*string, bool)`

GetInitialIdentificationOk returns a tuple with the InitialIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetInitialIdentification(v string)`

SetInitialIdentification sets InitialIdentification field to given value.

### HasInitialIdentification

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasInitialIdentification() bool`

HasInitialIdentification returns a boolean if a field has been set.

### GetInitialToken

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetInitialToken() string`

GetInitialToken returns the InitialToken field if non-nil, zero value otherwise.

### GetInitialTokenOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetInitialTokenOk() (*string, bool)`

GetInitialTokenOk returns a tuple with the InitialToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialToken

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetInitialToken(v string)`

SetInitialToken sets InitialToken field to given value.

### HasInitialToken

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasInitialToken() bool`

HasInitialToken returns a boolean if a field has been set.

### GetCustomerRedirectAddress

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetCustomerRedirectAddress() string`

GetCustomerRedirectAddress returns the CustomerRedirectAddress field if non-nil, zero value otherwise.

### GetCustomerRedirectAddressOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetCustomerRedirectAddressOk() (*string, bool)`

GetCustomerRedirectAddressOk returns a tuple with the CustomerRedirectAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRedirectAddress

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetCustomerRedirectAddress(v string)`

SetCustomerRedirectAddress sets CustomerRedirectAddress field to given value.

### HasCustomerRedirectAddress

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasCustomerRedirectAddress() bool`

HasCustomerRedirectAddress returns a boolean if a field has been set.

### GetConfiguration

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *OrderInitialResponseObjectOrderInitialResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *OrderInitialResponseObjectOrderInitialResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *OrderInitialResponseObjectOrderInitialResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


