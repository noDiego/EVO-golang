# SaleResponseObjectSaleResponseLayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **[]string** | Posibles atributos de impresión que se aplicarán en el Data. En caso de no estar presente el elemento se imprime el texto en el tamaño predefinido como normal y con justificación a la izquierda (LeftJustified). | [optional] 
**ContentType** | Pointer to **string** | En caso de no estar presente el elemento se asume que el elemento Data contiene texto plano (Text). Para el caso de Template los TAG o  elementos a tratar estan deliminatos por _$ y $_ quedando de la siguiente forma _$ElementName$_. Donde ElementName es el nombre del campo/elemento que deben incluir en el mismo. Ejemplo: Si el dispositivo recibe _$Ticket$_, reemplazara este indicaror por la Firma o Ticket Digitalizado. | [optional] 
**EncodeType** | Pointer to **string** | Tipo de codificación de los datos. En caso de no estar presente el elemento se asume que el elemento Data está codificado como Plain. | [optional] 
**Data** | Pointer to **string** | Dato a imprimir codificado según los valores presentes en ContentType y EncodeType. | [optional] 

## Methods

### NewSaleResponseObjectSaleResponseLayout

`func NewSaleResponseObjectSaleResponseLayout() *SaleResponseObjectSaleResponseLayout`

NewSaleResponseObjectSaleResponseLayout instantiates a new SaleResponseObjectSaleResponseLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleResponseObjectSaleResponseLayoutWithDefaults

`func NewSaleResponseObjectSaleResponseLayoutWithDefaults() *SaleResponseObjectSaleResponseLayout`

NewSaleResponseObjectSaleResponseLayoutWithDefaults instantiates a new SaleResponseObjectSaleResponseLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *SaleResponseObjectSaleResponseLayout) GetAttribute() []string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *SaleResponseObjectSaleResponseLayout) GetAttributeOk() (*[]string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *SaleResponseObjectSaleResponseLayout) SetAttribute(v []string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *SaleResponseObjectSaleResponseLayout) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetContentType

`func (o *SaleResponseObjectSaleResponseLayout) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SaleResponseObjectSaleResponseLayout) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SaleResponseObjectSaleResponseLayout) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SaleResponseObjectSaleResponseLayout) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEncodeType

`func (o *SaleResponseObjectSaleResponseLayout) GetEncodeType() string`

GetEncodeType returns the EncodeType field if non-nil, zero value otherwise.

### GetEncodeTypeOk

`func (o *SaleResponseObjectSaleResponseLayout) GetEncodeTypeOk() (*string, bool)`

GetEncodeTypeOk returns a tuple with the EncodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodeType

`func (o *SaleResponseObjectSaleResponseLayout) SetEncodeType(v string)`

SetEncodeType sets EncodeType field to given value.

### HasEncodeType

`func (o *SaleResponseObjectSaleResponseLayout) HasEncodeType() bool`

HasEncodeType returns a boolean if a field has been set.

### GetData

`func (o *SaleResponseObjectSaleResponseLayout) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SaleResponseObjectSaleResponseLayout) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SaleResponseObjectSaleResponseLayout) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SaleResponseObjectSaleResponseLayout) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


