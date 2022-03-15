# OrderInitialObjectOrderInitial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemIdentification** | **string** | ID que identifica el sistema desde donde proviene la petición. | 
**CompanyIdentification** | **string** | ID que identifica la companía desde donde proviene la petición. | 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**RequiredInformation** | Pointer to [**[]DebtPaymentObjectDebtPaymentRequiredInformation**](DebtPaymentObjectDebtPaymentRequiredInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**MerchantNotifyURL** | Pointer to **string** | URL para notificación del comercio | [optional] 
**MerchantRedirectURL** | Pointer to **string** | . | [optional] 
**DateTime** | Pointer to **time.Time** | Fecha y Hora de la transacción generada por el Punto de Venta - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 
**TransactionIdentification** | Pointer to **string** | ID de la operación a realizar, generado por el sistema externo | [optional] 
**TransactionDescription** | Pointer to **string** | Descripción del tipo de operación que se realizará | [optional] 
**MerchantIdentification** | Pointer to **string** | ID del comercio que el para el cual se estara realizando la operación. Este valor puede ser enviado en lugar del SystemIdentification y CompanyIdentification, para luego ser traducido por el propio Plataforma a los valores configurados para ello. Ademas, puede relaciónarse el valor de BranchIdentification y/o POSIdentification de la misma forma. Si ello no se realiza, estos elementos tomaran valores genericos (por default) igual a 0. | [optional] 
**Installments** | Pointer to **int32** | Cantidad de cuotas que permite este plan | [optional] 
**FacilityType** | Pointer to **int32** | Tipo de plan utilizado para para realizar la operación | [optional] 
**CurrencyCode** | Pointer to **string** | código de Moneda - ISO 4217 &lt;https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica &lt;br /&gt;   * Num   - Alpha - Description &lt;br /&gt;   * &#39;032&#39; - &#39;ARS&#39; - Pesos Argentinos &lt;br /&gt;   * &#39;152&#39; - &#39;CLP&#39; - Pesos Chilenos &lt;br/&gt;   * &#39;484&#39; - &#39;MXN&#39; - Pesos Mexicanos &lt;br/&gt;   * &#39;840&#39; - &#39;USD&#39; - dólares Americanos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;858&#39; - &#39;UYU&#39; - Pesos Uruguayos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;986&#39; - &#39;BRL&#39; - Real Brasileño | [optional] 
**Amount** | Pointer to **float32** | Monto con la que se realizó transacción. Si este valor es recibido, la búsqueda de los planes será limitada con este criterio. | [optional] 
**NetAmount** | Pointer to **string** | . | [optional] 
**Payer** | Pointer to [**SaleObjectSalePayer**](SaleObjectSalePayer.md) |  | [optional] 
**Customer** | Pointer to [**SaleObjectSaleCustomer**](SaleObjectSaleCustomer.md) |  | [optional] 
**Seller** | Pointer to [**SaleObjectSaleSeller**](SaleObjectSaleSeller.md) |  | [optional] 
**Products** | Pointer to [**[]SaleResponseObjectSaleResponseProducts**](SaleResponseObjectSaleResponseProducts.md) | Detalle de Productos de la Operación. | [optional] 
**TaxRefundType** | Pointer to **string** | Esquema de Devolución de Impuestos a utilizar en la transacción | [optional] 
**ValidThru** | Pointer to **time.Time** | Fecha y Hora de fin de validez de La transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14 | [optional] 

## Methods

### NewOrderInitialObjectOrderInitial

`func NewOrderInitialObjectOrderInitial(systemIdentification string, companyIdentification string, ) *OrderInitialObjectOrderInitial`

NewOrderInitialObjectOrderInitial instantiates a new OrderInitialObjectOrderInitial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderInitialObjectOrderInitialWithDefaults

`func NewOrderInitialObjectOrderInitialWithDefaults() *OrderInitialObjectOrderInitial`

NewOrderInitialObjectOrderInitialWithDefaults instantiates a new OrderInitialObjectOrderInitial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemIdentification

`func (o *OrderInitialObjectOrderInitial) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *OrderInitialObjectOrderInitial) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *OrderInitialObjectOrderInitial) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.


### GetCompanyIdentification

`func (o *OrderInitialObjectOrderInitial) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *OrderInitialObjectOrderInitial) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *OrderInitialObjectOrderInitial) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.


### GetBranchIdentification

`func (o *OrderInitialObjectOrderInitial) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *OrderInitialObjectOrderInitial) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *OrderInitialObjectOrderInitial) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *OrderInitialObjectOrderInitial) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *OrderInitialObjectOrderInitial) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *OrderInitialObjectOrderInitial) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *OrderInitialObjectOrderInitial) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *OrderInitialObjectOrderInitial) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetRequiredInformation

`func (o *OrderInitialObjectOrderInitial) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation`

GetRequiredInformation returns the RequiredInformation field if non-nil, zero value otherwise.

### GetRequiredInformationOk

`func (o *OrderInitialObjectOrderInitial) GetRequiredInformationOk() (*[]DebtPaymentObjectDebtPaymentRequiredInformation, bool)`

GetRequiredInformationOk returns a tuple with the RequiredInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInformation

`func (o *OrderInitialObjectOrderInitial) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation)`

SetRequiredInformation sets RequiredInformation field to given value.

### HasRequiredInformation

`func (o *OrderInitialObjectOrderInitial) HasRequiredInformation() bool`

HasRequiredInformation returns a boolean if a field has been set.

### GetServiceVersion

`func (o *OrderInitialObjectOrderInitial) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *OrderInitialObjectOrderInitial) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *OrderInitialObjectOrderInitial) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *OrderInitialObjectOrderInitial) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *OrderInitialObjectOrderInitial) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *OrderInitialObjectOrderInitial) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *OrderInitialObjectOrderInitial) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *OrderInitialObjectOrderInitial) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *OrderInitialObjectOrderInitial) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OrderInitialObjectOrderInitial) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OrderInitialObjectOrderInitial) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OrderInitialObjectOrderInitial) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetMerchantNotifyURL

`func (o *OrderInitialObjectOrderInitial) GetMerchantNotifyURL() string`

GetMerchantNotifyURL returns the MerchantNotifyURL field if non-nil, zero value otherwise.

### GetMerchantNotifyURLOk

`func (o *OrderInitialObjectOrderInitial) GetMerchantNotifyURLOk() (*string, bool)`

GetMerchantNotifyURLOk returns a tuple with the MerchantNotifyURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantNotifyURL

`func (o *OrderInitialObjectOrderInitial) SetMerchantNotifyURL(v string)`

SetMerchantNotifyURL sets MerchantNotifyURL field to given value.

### HasMerchantNotifyURL

`func (o *OrderInitialObjectOrderInitial) HasMerchantNotifyURL() bool`

HasMerchantNotifyURL returns a boolean if a field has been set.

### GetMerchantRedirectURL

`func (o *OrderInitialObjectOrderInitial) GetMerchantRedirectURL() string`

GetMerchantRedirectURL returns the MerchantRedirectURL field if non-nil, zero value otherwise.

### GetMerchantRedirectURLOk

`func (o *OrderInitialObjectOrderInitial) GetMerchantRedirectURLOk() (*string, bool)`

GetMerchantRedirectURLOk returns a tuple with the MerchantRedirectURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRedirectURL

`func (o *OrderInitialObjectOrderInitial) SetMerchantRedirectURL(v string)`

SetMerchantRedirectURL sets MerchantRedirectURL field to given value.

### HasMerchantRedirectURL

`func (o *OrderInitialObjectOrderInitial) HasMerchantRedirectURL() bool`

HasMerchantRedirectURL returns a boolean if a field has been set.

### GetDateTime

`func (o *OrderInitialObjectOrderInitial) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *OrderInitialObjectOrderInitial) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *OrderInitialObjectOrderInitial) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *OrderInitialObjectOrderInitial) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetTransactionIdentification

`func (o *OrderInitialObjectOrderInitial) GetTransactionIdentification() string`

GetTransactionIdentification returns the TransactionIdentification field if non-nil, zero value otherwise.

### GetTransactionIdentificationOk

`func (o *OrderInitialObjectOrderInitial) GetTransactionIdentificationOk() (*string, bool)`

GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentification

`func (o *OrderInitialObjectOrderInitial) SetTransactionIdentification(v string)`

SetTransactionIdentification sets TransactionIdentification field to given value.

### HasTransactionIdentification

`func (o *OrderInitialObjectOrderInitial) HasTransactionIdentification() bool`

HasTransactionIdentification returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *OrderInitialObjectOrderInitial) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *OrderInitialObjectOrderInitial) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *OrderInitialObjectOrderInitial) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *OrderInitialObjectOrderInitial) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetMerchantIdentification

`func (o *OrderInitialObjectOrderInitial) GetMerchantIdentification() string`

GetMerchantIdentification returns the MerchantIdentification field if non-nil, zero value otherwise.

### GetMerchantIdentificationOk

`func (o *OrderInitialObjectOrderInitial) GetMerchantIdentificationOk() (*string, bool)`

GetMerchantIdentificationOk returns a tuple with the MerchantIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantIdentification

`func (o *OrderInitialObjectOrderInitial) SetMerchantIdentification(v string)`

SetMerchantIdentification sets MerchantIdentification field to given value.

### HasMerchantIdentification

`func (o *OrderInitialObjectOrderInitial) HasMerchantIdentification() bool`

HasMerchantIdentification returns a boolean if a field has been set.

### GetInstallments

`func (o *OrderInitialObjectOrderInitial) GetInstallments() int32`

GetInstallments returns the Installments field if non-nil, zero value otherwise.

### GetInstallmentsOk

`func (o *OrderInitialObjectOrderInitial) GetInstallmentsOk() (*int32, bool)`

GetInstallmentsOk returns a tuple with the Installments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallments

`func (o *OrderInitialObjectOrderInitial) SetInstallments(v int32)`

SetInstallments sets Installments field to given value.

### HasInstallments

`func (o *OrderInitialObjectOrderInitial) HasInstallments() bool`

HasInstallments returns a boolean if a field has been set.

### GetFacilityType

`func (o *OrderInitialObjectOrderInitial) GetFacilityType() int32`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *OrderInitialObjectOrderInitial) GetFacilityTypeOk() (*int32, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *OrderInitialObjectOrderInitial) SetFacilityType(v int32)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *OrderInitialObjectOrderInitial) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrderInitialObjectOrderInitial) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderInitialObjectOrderInitial) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderInitialObjectOrderInitial) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrderInitialObjectOrderInitial) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmount

`func (o *OrderInitialObjectOrderInitial) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderInitialObjectOrderInitial) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderInitialObjectOrderInitial) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OrderInitialObjectOrderInitial) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetNetAmount

`func (o *OrderInitialObjectOrderInitial) GetNetAmount() string`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *OrderInitialObjectOrderInitial) GetNetAmountOk() (*string, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *OrderInitialObjectOrderInitial) SetNetAmount(v string)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *OrderInitialObjectOrderInitial) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetPayer

`func (o *OrderInitialObjectOrderInitial) GetPayer() SaleObjectSalePayer`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *OrderInitialObjectOrderInitial) GetPayerOk() (*SaleObjectSalePayer, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *OrderInitialObjectOrderInitial) SetPayer(v SaleObjectSalePayer)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *OrderInitialObjectOrderInitial) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderInitialObjectOrderInitial) GetCustomer() SaleObjectSaleCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderInitialObjectOrderInitial) GetCustomerOk() (*SaleObjectSaleCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderInitialObjectOrderInitial) SetCustomer(v SaleObjectSaleCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderInitialObjectOrderInitial) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSeller

`func (o *OrderInitialObjectOrderInitial) GetSeller() SaleObjectSaleSeller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *OrderInitialObjectOrderInitial) GetSellerOk() (*SaleObjectSaleSeller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *OrderInitialObjectOrderInitial) SetSeller(v SaleObjectSaleSeller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *OrderInitialObjectOrderInitial) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetProducts

`func (o *OrderInitialObjectOrderInitial) GetProducts() []SaleResponseObjectSaleResponseProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *OrderInitialObjectOrderInitial) GetProductsOk() (*[]SaleResponseObjectSaleResponseProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *OrderInitialObjectOrderInitial) SetProducts(v []SaleResponseObjectSaleResponseProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *OrderInitialObjectOrderInitial) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetTaxRefundType

`func (o *OrderInitialObjectOrderInitial) GetTaxRefundType() string`

GetTaxRefundType returns the TaxRefundType field if non-nil, zero value otherwise.

### GetTaxRefundTypeOk

`func (o *OrderInitialObjectOrderInitial) GetTaxRefundTypeOk() (*string, bool)`

GetTaxRefundTypeOk returns a tuple with the TaxRefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRefundType

`func (o *OrderInitialObjectOrderInitial) SetTaxRefundType(v string)`

SetTaxRefundType sets TaxRefundType field to given value.

### HasTaxRefundType

`func (o *OrderInitialObjectOrderInitial) HasTaxRefundType() bool`

HasTaxRefundType returns a boolean if a field has been set.

### GetValidThru

`func (o *OrderInitialObjectOrderInitial) GetValidThru() time.Time`

GetValidThru returns the ValidThru field if non-nil, zero value otherwise.

### GetValidThruOk

`func (o *OrderInitialObjectOrderInitial) GetValidThruOk() (*time.Time, bool)`

GetValidThruOk returns a tuple with the ValidThru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidThru

`func (o *OrderInitialObjectOrderInitial) SetValidThru(v time.Time)`

SetValidThru sets ValidThru field to given value.

### HasValidThru

`func (o *OrderInitialObjectOrderInitial) HasValidThru() bool`

HasValidThru returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


