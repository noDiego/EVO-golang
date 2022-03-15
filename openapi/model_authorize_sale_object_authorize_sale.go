/*
EVO Payment API

<h1 style=\"border:4px solid #004785;color:#004785; font-family: 'Open Sans', sans-serif; font-size: 32px;\">API de Pagos</h1> <br /><br /> <h1 style=\"border:4px solid #004785;color:#004785; font-family: 'Open Sans', sans-serif; font-size: 24px;\">Descripción del Servicio</h1> <p style=\"color:#004785;\"><b><u>Documentación en formato OpenAPI 3.0</b></u></p> <br/> Contrato especificado según especificaciones https://www.openapis.org/ y https://swagger.io/.<br /><br />  En el site https://editor.swagger.io/ se dispone de un  Viewer, Editor y  Generar de Código ( SDK ) para varios lenguajes de programación; incluyendo JAVA, C#, C++, Perl, Node.js, GO, PHP, Ruby y otros.<br/><br/> Para <b>ver</b> la documentación o <b>generar</b> código de la librería cliente o SDK  se deberá selecciónar en el menú horizontal  la opción <b>File</b>, en el menú vertical que se depliega la opción <b>Import File</b> y luego se deberá selecciónar el archivo del contrato deseado, ya sea  extensión <b>.json</b> o <b>.yaml</b>. <br/><br/> Además se puede generar el código de la librería cliente desde la línea de comandos a través de la herramienta  <b>CLI</b>  de  <b>OpenAPI Generator</b>. Esta presenta generadores de SDK en mayor variedad de lenguajes de programación.  En el site https://openapi-generator.tech/docs/installation se documenta cómo  <b>instalar</b> la herramienta CLI.<br/><br/> Los clientes generados contienen, adicionalmente al código,  la documentación de uso del mismo en <b>README.md</b>, como también en el subdirectorio <b>docs</b> toda la documentación del API o servicio y sus operaciones, con el detalle de los  campos o elementos y su dominio.<br /><br /><br /><br /> <h1 style=\"border:4px solid #004785;color:#004785; font-family: 'Open Sans', sans-serif; font-size: 24px;\">Notas a tener en cuenta para realizar la Integración</h1><br/> <p style=\"color:#004785;\"><b><u>Conceptos y/ Mecanismos relevantes Soportados por el Protocolo de Integración</u></b></p> <br/><br/> <span>&#8226;</span> <b>Interpretración de las Respuesta</b>,<br /><br/> El único campo que indica si la transacción fue aprobada, rechazada, o tienen algun error, es el elemento de las respuestas llamado <b>ResponseActions</b>, el  cual es un <b>ARRAY</b> de valores. Cada uno de estos indica una acción a realizar. Los elementos <b>ResponseCode</b>  y <b>ResponseMessage</b> son solamente informativos y por lo tanto no deben usarse para tomar acciones y los mismos pueden cambiar en base a la configuración de la Plataforma.<br/><br/> <span>&#8226;</span> <b>Bloque de transacciones</b>, permite Confirmar o Cancelar/Revertir todas las transacciones que forman parte de un bloque. <br/><br/> El POS puede definir un bloque o conjunto de transacciones simplemente indicando en todas ellas el mismo valor en el atributo/elemento/campo opcional llamado <b>Block</b>.<br/>  La operación <b>BlockCancel</b>, permite que el POS pueda solicitar a la plataforma la reversión y/o cancelación de todo el bloque de transacciones .<br/> La operación <b>BlockClose</b>, confirma todas las transacciones que forman parte del bloque especificado.<br/> Si el POS no posee un identificador unívoco de la transacción de venta, al momento de interactuar contra la plataforma podrá obtener uno con la operación  <b>BlockCreate</b>. Si el elemento o campo <b>Block</b>  existe y su contenido es Vacío o Nulo la plataforma realiza un <b>BlockCreate</b> automáticamente.<br /><br/> <span>&#8226;</span> <b>Reversas por Ruptura de Secuencia</b>. Evita la necesidad de persistir datos de la reversa y ahorra una transacción en el flujo.<br/>   El método llamado de ruptura de secuencia es utilizado para detectar los casos en los cuales el POS o Caja no pudo recibir una respuesta del mismo o no pudo procesarla adecuadamente. De esta forma permite a la misma reversar la transacción que no pudo procesar el POS o recibir la respuesta si fuese necesario.     En todo requerimiento el POS debe enviar el campo/elemento Sequence, con el valor recibido en el anterior requerimiento o vacío en el primero.    La plataforma  genera una nueva secuencia solamente cuando el requerimiento realizado es reversible o cuando se produce una ruptura.  Por lo tanto los comandos en los cuales la plataforma  genera una nueva secuencia son <b>Sale</b>, <b>Void</b>, <b>Authorize*</b>, <b>DebtPayment</b>, <b>VoidDebtPayment</b>, <b>Confirm</b>, <b>Close</b> y <b>Cancel</b>.    En caso de que la plataforma reverse el requerimiento previo retornará en la respuesta los siguiente campos o elementos.   <blockquote><b>WasReversePrevious</b>, con valor <b>1</b><br/>   <b>ReversedAnswerKey</b> conteniendo el <b>AnswerKey</b> de la transacción reversada<br/>   <b>ReversedSequence</b> conteniendo el <b>Sequence </b>de la transacción reversada</blockquote>    En caso de que la plataforma no reverse el requerimiento previo retornará los siguientes campos o elementos <blockquote><b>WasReversePrevious</b>, con valor <b>0</b></blockquote> <br/> <span>&#8226;</span> <b>Reversas Tradicionales</b>. El POS debe repetir el mismo requerimiento adicionando el atributo/elemento <b>IsReverse</b> con valor <b>1</b>.  Se debe tener en cuenta que en esta modalidad la plataforma no retorna los siguientes atributos/elementos.    <blockquote>   <b>WasReversePrevious</b><br/>   <b>ReversedAnswerKey</b><br/>   <b>ReversedSequence</b>   </blockquote>    <span>&#8226;</span> <b>Transacción Opcional de Confirmación</b>, ya que el mecanismo anterior permite que cada transacción Reverse o Confirme la anterior.<br/><br/> <span>&#8226;</span> <b>La Plataforma indica siempre las acciones que se deben realizar</b><br/><br/> <span>&#8226;</span><span>&#8226;</span> <b>Solicitar datos adicionales</b> ( <b>RequiredInformation</b> ), indicando no sólo cuáles son, sino también de qué tipo, valor  inicial, patrón de validación, si son mandatorios o no, qué Label se presenta al usuario, qué ayuda se presenta al usuario, etc.<br/> <span>&#8226;</span><span>&#8226;</span> <b>Mostrar Mensajes en Pantalla</b>. <span>&#8226;</span><span>&#8226;</span> <b>Imprimir Tickets</b>, ya sea en papel o capturar digitalmente el mismo, como así también el Layout de los mismos.<br/><br/><br/> <span>&#8226;</span> <b>Compresión de la trama</b> en base a codificación de los campos numéricos, string siempre de longitud variable, uso de sinónimos en los  campos, para que el programador programe usando los nombres largos y en el transporte se usen sus sinónimos cortos. <br/> <br/> <span>&#8226;</span> <b>Seguridad de los Datos Sensibles y de la Transaccion</b>, El elemento <b>Security</b> debera estar presente solo si los datos sensibles <b>CardNumber</b>, <b>ExpDate</b>, <b>PIN</b>, <b>Track1</b>, <b>Track2</b>, <b>SecurityCode</b> y  <b>CardCryptogram</b> deban ser envidos encriptados y por lo tanto este le elemento nos permite indicar el metodo de encriptacion utilizado y los datos adicionales que sean requeridos por la encriptacion. Si por ejemplo fuese el elemento PIN usando DUKPT y el resto de los datos sencible Track1, Track2 y SecurityCode, se deberian enviar  de la siguiente forma: </br>        \"Security” :  [         {           \"Type\": \"PIN\",           \"Values\":  [               {                    \"Name\": \"Method\",                   \"Value\": \"DUKPT\"               },               {                    \"Name\": \"KSN\",                   \"Value\": \"1234567890ABCDEF\"               },               {                    \"Name\": \"CRC32\",                   \"Value\": \"12345\"               },               {                    \"Name\": \"PlainTextLength\",                   \"Value\": \"4\"               },               {                    \"Name\": \"CipherCounter\",                   \"Value\": \"123\"               },               {                    \"Name\": \"ConsecutiveFailedCiphersCounter\",                   \"Value\": \"123\"               },               {                    \"Name\": \"Data\",                   \"Value\": \"01234567890123456\"               }           ]          },         {           \"Type\": \"SensitiveData\",           \"Values\":  [               {                    \"Name\": \"Method\",                   \"Value\": \"DUKPT-eGlobal\"               },               {                    \"Name\": \"KSN\",                   \"Value\": \"1234567890ABCDEF\"               },               {                    \"Name\": \"Track1CRC32\",                   \"Value\": \"12345\"               },               {                    \"Name\": \"Track2CRC32\",                   \"Value\": \"12345\"               },               {                    \"Name\": \"Track1Length\",                   \"Value\": \"79\"               },               {                    \"Name\": \"Track2Length\",                   \"Value\": \"37\"               },               {                    \"Name\": \"SecurityCodeLength\",                   \"Value\": \"3\"               },               {                    \"Name\": \"CipherCounter\",                   \"Value\": \"123\"               },               {                    \"Name\": \"ConsecutiveFailedCiphersCounter\",                   \"Value\": \"123\"               },               {                    \"Name\": \"Data\",                   \"Value\": \"1ahbcd23412345123412b213b1324b1234b2134b2134132b4123b23\"               }           ]          },         {           \"Type\": \"3DSecure\",           \"Values\":  [               {                    \"Name\": \"Method\",                   \"Value\": \"3DS-SNAP\"               },               {                                           \"Name\":  \"TransactionStatus\",                   \"Value\": \"SuccessfullyAuthenticated\"               },               {                                           \"Name\":  \"AuthenticationECI\",                   \"Value\": \"05\"               },               {                                           \"Name\":  \"IsChallengeMandated\",                   \"Value\": \"false\"               },               ...               {                                           \"Name\":  \"AcsReferenceNumber\",                   \"Value\": \"3DS_LOA_ACS_PPFU_020100_00009\"               },               {                    \"Name\":  \"ProcessedAsDataOnly\",                   \"Value\": \"false\"               }           ]          }               ] </br> Para el caso de DUKPT-eGlobal, <b>Track2</b>, <b>SecurityCode</b> y <b>Track1</b> se cifraran formando parte del mismo Bloque, El mismo se debera formar con el Track2 ( reemplazando el signo = por el digito D ) completandolo hasta los 38 digitos con el digito F, luego el  SecurityCode completandolo hasta 10 digitos y por ultimo el Track1 padeado completando el bloque  de los 208 digitos.  </br> Este elemento <b>Security</b> sera utilizado para enviar cualquier dato de autenticacion del pagador por ejemplo 3DSecure, para el caso de que el proveedor de la Autenticacion sea SNAP se deberan contener como valores todos los elementos definidos en el objeto <b>ThreeDSInformation</b>.     </br> Este mecanismo podra utilizarse en el futuro para encriptar otros datos que sean sensibles pero no del medio de pago, si no de las personas. </br> <h1 style=\"border:4px solid #004785;color:#004785; font-family: 'Open Sans', sans-serif; font-size: 24px;\">Log de Cambios</h1></br> <span>&#8226;</span> <b>Versión 5.6.1</b> <span>&#8226;</span><span>&#8226;</span> Se añade el campo <b>MerchantCategory</b> en las respuestas de todas las transacciones. Sólo se enviará en caso de que la categoría de la compañia exista.</br> <span>&#8226;</span> <b>Versión 5.6.0</b> <span>&#8226;</span><span>&#8226;</span> Los campos <b>ResponseCode</b>, <b>ResponseMessage</b> y <b>ResponseActions</b> son <b>obligatorios</b> en las respuestas de todas las transacciones.</br> <span>&#8226;</span> <b>Versión 5.5.7</b> <span>&#8226;</span><span>&#8226;</span> Se añade el elemento <b>Notification</b>. El mismo se encuentra dentro de <b>SaleResponse</b> y <b>AuthorizeSaleResponse</b>. Notificación a generar alertas vía e-mail.</br> <span>&#8226;</span> <b>Versión 5.5.6</b> <span>&#8226;</span><span>&#8226;</span> Se añaden los elementos <b>CardAppLabel</b>, <b>CardAuthRequestCryptogram</b> y <b>CardAuthResponseCryptogram</b>, para facilitar el analisis de los POS y ReadingDevices, el contenido de dichos elementos se encontraba en Tag de los elementos CardCryptogram y CardCryptogramResponse.</br> <span>&#8226;</span>  <b>Versión 5.5.5</b> <span>&#8226;</span><span>&#8226;</span> Se modifican los elementos <b>AuthorizeSale</b> y <b>AuthorizeSaleResponse</b> para su correcta documentación. Además, se añade el campo <b>ReadingDeviceOperatingFrom</b> el cual indica desde cuando se encuentra operativo o encendido el dispositivo</br> <span>&#8226;</span> <b>Versión 5.5.4</b> <span>&#8226;</span><span>&#8226;</span> Se renombra el atributo <b>ReasonReverse</b> a <b>ReverseReason</b>. Dicho campo permite notificar en las Reversas la razón por la cual fue necesario generarla.</br> <span>&#8226;</span> <b>Versión 5.5.3</b> <span>&#8226;</span><span>&#8226;</span> Se agregan atributos al elemento <b>Configuration</b> para la operación <b>PaymentMethod</b>. Por otra parte, se añade el mismo en todas las operaciones donde no se encontraba documentado. </br><b>• Versión 5.5.2</b> <span>&#8226;</span><span>&#8226;</span> Se Agrega el elemento <b>Payer</b> con los datos del Pagador. Originalmente hasta esta version se envian los mismos en el elemento <b>Customer</b>, pero desde ahora se permite que se informen personas ( fisicas y juridicas ) como cliente comprador y como pagador. Si el elemento <b>Payer</b> no esta presente se tomaran los datos del elemento <b>Customer</b>. Se da soporte al Tipo de Ticket Payer.</br> <span>&#8226;</span> <b>Versión 5.5.1</b> </br> <span>&#8226;</span><span>&#8226;</span> Se completa la documentacion de los Elementos <b>Seller</b> y <b>Customer</b>, agregandose los atributos <b>City</b> y  <b>AbbreviatedName</b>.<br/>   <span>&#8226;</span><span>&#8226;</span> Se unifica la definicion del Elemento  <b>Customer</b> .<br/>   <span>&#8226;</span><span>&#8226;</span> Se agrega el Elemento <b>PaymentFacilitatorID</b> para indicar el Identificador de Facilitador de pagos o Payfac.</br> <span>&#8226;</span> <b>Versión 5.5.0</b> </br> <span>&#8226;</span><span>&#8226;</span> El elemento <b>ResponseActions</b> y <b>PosOrDeviceAction</b> de todas las operaciones deja de ser una lista.<br/>  de elementos en un string y se convierte en un array de string. Cada valor de la lista anterior está representada por un elemento del array.<br/>   <span>&#8226;</span><span>&#8226;</span> Se agregan los campos <b>ForeignIdentifier</b>, <b>SmallImage</b> y <b>LargeImage</b> en el campo <b>Wallets</b> de la operación <b>WalletsResponse</b>.<br/> <span>&#8226;</span><span>&#8226;</span> En el campo <b>PaymentMethods</b> de la operación <b>PaymentMethodsResponse</b> se agregan las properties <b>Imag</b>, <b>SmallImage</b> y <b>LargeImage</b>. Además se adiciona el campo <b>ID</b> en <b>Category</b> y el campo <b>ForeignIdentifier</b> en <b>Type</b>. <br/> <span>&#8226;</span><span>&#8226;</span> Se agrupan los campos relacionados con los datos del cliente y del vendedor en dos únicos campos de tipo objeto denominados <b>Customer</b> y <b>Seller</b>, respectivamente.<br/> <span>&#8226;</span><span>&#8226;</span> El elemento Layout del campo <b>Tickets</b> se convierte en un array de objetos que contiene elementos que permiten describir, dar formato y codificar los datos a imprimir. <br/> <span>&#8226;</span><span>&#8226;</span> Se documenta la operación <b>OrderStatus</b>.</br> <span>&#8226;</span><span>&#8226;</span> Los campos que refieren a tiempo y fecha se convierten en formato date-time. </br> <span>&#8226;</span><span>&#8226;</span> Se agrega el campo <b>ForeignResponseCode</b> en todas las respuestas de las operaciones, como un código de para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF.</br> <span>&#8226;</span><span>&#8226;</span> Se adiciona el campo <b>CardGetMode</b> que permite indicar por cada elemento que contiene los datos sensibles, si están encriptados y también el algoritmo usado. En caso de no estar especificado se asume PLAIN.</br> <span>&#8226;</span><span>&#8226;</span> Se agrega el campo <b>OrigReference</b> en aquellas operaciones que pueden referenciar a una transacción previa, como <b>Void</b>, <b>Return</b> y <b>GetTransaction</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se cambia la estrutura de la respuesta de la Operacion <b>GetTransacion</b> por errores. </br> <span>&#8226;</span><span>&#8226;</span> Se agregan las acciones Ok, Error y Retry en los campos <b>ResponseActions</b>.</br> <span>&#8226;</span><span>&#8226;</span> En aquellas operaciones financieras en las que se especifica la tarjeta se agrega en el requerimiento el campo <b>Pin</b>y en la respuesta el campo <b>WorkingKeys</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se agrega el campo <b>Security</b> con el objetivo de indicar los datos sensibles de seguridad de una transacción tanto en los requerimientos como en las respuestas de las operaciones disponibles.</br> <span>&#8226;</span><span>&#8226;</span> Se agrega la operacion <b>KeysRenewal</b> Las claves podran ser retornadas en el elemento <b>Security</b> y en caso de obtener como accion de respuesta <b>KeysRenewal</b> se esta indicando que esta nueva operacion debe ser ejecutada.<br/>      <span>&#8226;</span><span>&#8226;</span> Se agrega la opcion <b>Signature</b>  .<br/>     <span>&#8226;</span><span>&#8226;</span> Se agrega el elemento  <b>CategoryCode</b> para especificar el MCC del Vendedor y/o del Cliente  .<br/>     <span>&#8226;</span><span>&#8226;</span> Se agregan los Elementos <b>MerchantID</b>, <b>TerminalID</b>, <b>TraceNumber</b> y <b>SettlementBatchNumber</b> En los requerimientos, en caso que los mismos contengan valor los mismos seran utilizados para enviar al Host Resolutor de la Transaccion.</br>  <span>&#8226;</span><span>&#8226;</span> Se agregan los valores para pagos recurrentes a  los Elementos  <b>CardReadMode</b> y  <b>CardReadModeDescription</b> <span>&#8226;</span> <b>Versión 5.4.0</b> </br> <span>&#8226;</span><span>&#8226;</span> Se cambia la dirección IP por el nombre.</br> <span>&#8226;</span><span>&#8226;</span> Se contemplan los Datos del <b>Vendedor/Seller</b> y del <b>Cliente/Customer</b> en las operaciones  <b>WalletRequest</b>, <b>Sale</b>, <b>AuthorizeSale</b>, <b>DebtPayment</b>,  <b>Deposit</b>,  <b>Settlement</b>,  <b>Capture</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se Agregan los elementos <b>POSGEO</b> y <b>ReadingDeviceGEO</b> para que el dispositivo de lectura y el Punto de venta Notifiquen su coordenadas georefenciales en el momento de que se realiza la transacción.</br> <span>&#8226;</span><span>&#8226;</span> Se unifica y amplía el elemento <b>RequiredInformation</b>  tanto en los requerimientos como en las respuestas</br>  <span>&#8226;</span><span>&#8226;</span> Se cambia el tipo el elemento <b>CurrencyCode</b> a string para permitir cualquieras de la notaciones posibles.</br> <span>&#8226;</span><span>&#8226;</span> Se cambia el  elemento <b>Currency</b> por <b>CurrencyCode</b>  en el elemento <b>Plans</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se contemplan del detalle ( elemento <b>Products</b> ) de la venta en las operaciones  <b>WalletRequest</b>, <b>Sale</b>, <b>Void</b>, <b>Return</b>, <b>AuthorizeSale</b>, <b>DebtPayment</b>,  <b>VoidDebtPayment</b>, <b>Deposit</b>,  <b>Settlement</b>,  <b>Capture</b>.</br> <span>&#8226;</span><span>&#8226;</span> Agregamos la operación <b>DebtInquiry</b> que actua como sinónimo de <b>BalanceInquery</b>, la cual podía ser usada para consulta de Saldo y también de deuda.</br> <span>&#8226;</span><span>&#8226;</span> Se corrigen los tipos de Datos de Varios campos <b>Amount</b> que en lugar de string debían ser number.</br> <span>&#8226;</span><span>&#8226;</span> Se agregan las operaciones <b>QueryCompanies</b> y <b>QueryLineOfBusiness</b> para la consulta de Rubros y Empresas que se pueden utilizar para pagar Servicios/Deuda/Facturas con la operación <b>DebtPayment</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se adiciona el elemento <b>Companies</b> en la Operacion <b>BalanceInquiry</b> para el caso de que existan mas de una Compania para el mismo codigo o identificador de la deuda o factura a pagar y adicionalmente se agrega para ese caso la posibilidad de especificar a que compania corrende el Pago en el elemento <b>DebtCompanyIdentification</b> en la operación <b>DebtPayment</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se adiciona el elemento <b>BaseAmonut</b> en los requerimientos de las operación <b>Return</b>, el elemento <b>Reference</b>  en las operaciones <b>Sale</b>, <b>AuthorizeSale</b>, <b>Void</b>, <b>Return</b>, <b>DebtPayment</b>, <b>VoidDebtPayment</b>, <b>GetTransaction</b> y <b>WalletRequest</b>.  Además, se agregan los elementos <b>TaxFinancialCostAmount</b>, <b>TaxFinancialCostPercentage</b>, <b>FinancialCostAmount</b>, <b>FinancialCostPercentage</b> y <b>RequestAmount</b>  en las respuestas de dichas operaciones.</br> <span>&#8226;</span><span>&#8226;</span> En cada plan que se devuelve a través del <b>PaymentMethodResponse</b> estarán presentes <b>TaxFinancialCostAmount</b>,  <b>TaxFinancialCostPercentage</b>, <b>FinancialCostAmount</b> y <b>FinancialCostPercentage</b>. </br> <span>&#8226;</span><span>&#8226;</span> Se agregan los elementos  <b>CardAppName</b> y <b>CardAppIdentifier</b> en las peticiones de las operaciones <b>Sale</b>, <b>AuthorizeSale</b>,  <b>Void</b>, <b>Return</b>, <b>PaymentMethods</b>, <b>GetCard</b>, <b>Validate</b>, <b>DebtInquiery</b>, <b>BalanceInquiry</b>, <b>DebtPayment</b> y <b>VoidDebtPayment</b>.  Además, se agregan en las respuestas de algunas de ellas.</br> <span>&#8226;</span><span>&#8226;</span> Se cambia la estructura del elemento <b>Tickets</b> de las respuestas donde el elemento <b>Action</b>  hace referencia a las acciones que debe ejecutar el punto de venta, el elemento <b>DeviceAction</b> a las acciones que debe ejecutar el dispositivo y <b>ExecutedAction</b> a las acciones  que ejecutó la plataforma para sus <b>Tickets</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se adicionan los elementos <b>POSOrDeviceAction</b>, <b>OperationMode</b> y <b>OperationModeDescription</b> a la operación <b>Configure</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se agrega el elemento <b>RemainderAmount</b> a la operación <b>GetBlockResponse</b> que hace referencia a la diferencia entre el monto total de la transacción y las devoluciones parciales realizadas.</br> <span>&#8226;</span><span>&#8226;</span> Se corrijen errores en la definición de varios campos, como <b>ReadingDeviceType</b> y <b>CardReadMode</b>.</br> <span>&#8226;</span><span>&#8226;</span> Se reemplaza el campo <b>ApplicationIdentification</b> por <b>SystemIdentification</b> en las operaciones <b>EnableService</b>, <b>Wallets</b>, <b>QueryCompanies</b>,  <b>QueryLineOfBusiness</b> y sus respectivas respuestas. </br> <span>&#8226;</span><span>&#8226;</span> Se agregan el identifidor Tributario en <b>Sale</b>, <b>AuthorizeSale</b>, <b>Void</b>,  <b>Return</b>, <b>DebtPayment</b>, <b>VoidDebtPayment</b> y <b>Debtinquery</b> que permite transacciones con <b>Token</b> ( <b>CredentialToken</b> y  <b>CredentialIssuerToken</b> ). En estas operaciones se elimina de mandatorias al campo <b>BranchIdentification</b> y <b>POSIdentification</b><br/> <span>&#8226;</span><span>&#8226;</span> Se agrega la operación <b>Enrollment</b>, la cual permite transacciones con <b>Token</b> ( <b>CredentialToken</b> y  <b>CredentialIssuerToken</b> ) y pagos recurrentes.</br> <span>&#8226;</span><span>&#8226;</span> El campo <b>ResponseAction</b> deja de ser un enum y se convierte en string. Se indica en la descripción los posibles actions.</br> <span>&#8226;</span><span>&#8226;</span> Se agregan los campos <b>SellerIdentification</b> y <b>SellerIdentificationType</b> en aquellas operaciones en las que se especifican con los datos del vendedor.</br> <span>&#8226;</span><span>&#8226;</span> El campo <b>FacilityPayments</b> deja de ser mandatario en las operaciones <b>Enrollment</b> y <b>Sale</b>. </br> <span>&#8226;</span><span>&#8226;</span> Se elimina la posibilidad de envío en el header HTTP.<br/> <span>&#8226;</span><span>&#8226;</span> Se agregan los campos <b>CashbackAmount</b> y <b>TipAmount</b> en la operación <b>WalletRequest</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Se adiciona en el campo <b>CardReadMode</b> la opción K de Token.<br/> <span>&#8226;</span><span>&#8226;</span> Se corrige el campo <b>Answertype</b> y se modifica por <b>AnswerType</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Se agregan los campos referidos al vendedor en las operaciones <b>Void</b> y <b>Return</b>. <br/> <span>&#8226;</span><span>&#8226;</span> Se crea un primer nivel para cada operación de tipo objeto. <br/>  <span>&#8226;</span><span>&#8226;</span> Se crea el campo <b>InputTokens</b> como un array de objetos que contienen Name y Value como properties en las operaciones <b>Sale</b>, <b>AuthorizeSale</b>, <b>Void</b>, <b>Return</b>, <b>DebtPayment</b>, <b>VoidDebtPayment</b> y <b>DebtInquiry</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Los elementos <b>Action</b>, <b>DeviceAction</b> y <b>ExecutedAction</b> del campo <b>Tickets</b> dejan de ser de tipo string y se convierten en arrays.<br/>     <span>&#8226;</span><span>&#8226;</span> Se agrega el elemento <b>AdditionalInformation</b> en las respuestas de todas las operaciones.<br/>     </br> <span>&#8226;</span> <b>Versión 5.3.0</b> Se amplía la definición de la Operación <b>Configure</b> permitiendo tanto en la respuesta como en el requerimiento los elementos <b>Operations</b>, <b>Tables</b> y <b>Files</b></br></br> Se agregan los elementos <b>VoidSupport</b>, <b>ReturnSupport</b>, <b>WalletUseInVoidTransaction</b> y <b>WalletUseInReturnTransaction</b> en las caracteristicas de un Wallet.<br/><br/> Se agrega el Valor <b>Display</b> en el elemento <b>ResponseActions</b> indicando que se debe mostrar en el Display del Dispositivo o Aplicativo el contenido del elemento <b>DisplayResponseMessage</b>.  En la respuesta de la operación  <b>BalanceInquery</b> se agregan los elementos <b>AmountAvailable</b> y <b>PointsAvailable</b> para indicar los saldos.</br> Se especifica en la documentación que el Cancel puede ser usado para Cancelar un Pago con Wallets en Curso.</br></br> Se agregan elementos en los Requerimientos y en las respuestas opcionales entre los POS* que permiten describir las características del punto de venta, los Device* que permiten especificar las características del Dispositivo de Lectura.<br/>   Se cambió el elemento <b>AnswerIdentification</b> por <b>AnswerKey</b>  para compatibilizar con el servicio de Pagos.<br/><br/>     Se agregaron <b>AccountNumber</b>, <b>AccountType</b> y <b>Balance</b> en las operaciones <b>BalanceInquiry</b> y <b>DebtPayment</b> .<br/><br/>     Se agregaron las Operaciones <b>Confirm</b> y  <b>Cancel</b>, donde la Operación <b>Confirm</b> es usada para confirmar un pago recibido por el POS o Aplicativo del comercio. Existen Wallets en los que la confirmación es automática y se indica en el Elemento  <b>AutoConfirm</b> de la respuesta del comando <b>Wallets</b>. La operación <b>Cancel</b> puede ser utilizada a partir de que la Plataforma retorne la acción <b>PaymentFlowIsCancelable</b> en la respuesta de una operación <b>WalletRequest</b>. El Wallet soporta Cancelación de Requerimiento lo cual está indicado con el Elemento <b>SupportRequestCancel</b> dentro de las propiedades de  los Wallets que son retornados por la Operación <b>Wallets</b>.<br/> Se agregó como carasterística de los Wallets también el elemento <b>SupportValidityOfTheRequest</b> que indica si en el primer requerimiento de la Operación <b>WalletRequest</b> se puede enviar el elemento <b>TransactionTimeout</b> que especifica el tiempo de vida de la intención de pago. Superado ese tiempo no se podrá pagar y el ciclo de reintento será detenido por la plataforma, indicado por las siguientes acciones: <b>Completed</b> y <b>Error</b>.<br/> Se agrega el Elemento <b>Tickets</b> en la respuesta de una Operación <b>WalletRequest</b>. El elemento estará presente si como acción está presente el Valor <b>Tickets</b>, indicando que los mismos deberán ser Impresos, capturados digitalmente, etc. según se indique. <br/><br/> Se permite en la Operación <b>PaymentMethod</b> la búsqueda por el Id o el ForeignIdentification<br/><br/> <span>&#8226;</span> <b>Versión 5.2.6</b> Se cambia el nombre del elemento <b>DateTime</b> por <b>TransactionDateTime</b> en la operación <b>WalletRequest</b>.<br/><br/> <span>&#8226;</span> <b>Versión 5.2.5</b> Se agregan en los Planes el atributo <b>POSOrDeviceActions</b> que permite indecarle al Dispositivo que debe solicitar  PIN para esa transacción y eso lo indica enviando la acción <b>RequestPIN</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Se agrega el <b>ResponseActions</b> <b>Configure</b> que indica que se debe ejecutar una reconfiguración para obtener  parámetros nuevos ya que hay alguna actualización. <span>&#8226;</span><span>&#8226;</span> Se agregan las Operaciones <b>Wallets</b>, <b>WalletRequest</b> y <b>EnableService</b>, las mismas pueden formar parte  de un Block y forman parte de la ruptura de Secuencia.<br/><br/> <span>&#8226;</span> <b>Versión 5.2.4</b> Se agrega el identifidor Tributario en <b>OrderInitial</b>, que permite transacciones con <b>Token</b> ( <b>CredentialToken</b> y  <b>CredentialIssuerToken</b> ).<br/> <span>&#8226;</span><span>&#8226;</span> Se completa el <b>GetCardResponse</b> para que contenga los  elementos <b>PaymentMethod</b> y <b>Plans</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Se completa el <b>PaymentMethodResponse</b> para que contenga los elementos <b>PaymentMethod</b> y <b>Plans</b>.<br/> <span>&#8226;</span><span>&#8226;</span> Se agrega en el <b>GetCard</b>: permite forzar un modo de lectura y permite solicitar los datos leídos al POS <b>CardGetMode</b>. <br/><br/> <span>&#8226;</span><span>&#8226;</span> Se permite el envío de datos del cliente <b>Customer*</b> en las operaciones Financieras.<br/><br/> <span>&#8226;</span> <b>Versión 5.2.3</b> Se cambian los valores posibles para <b>ResponseActionCancel</b> en las operaciones <b>GetBlock</b> y <b>GetTransaction</b>.<br/>   <br/> <span>&#8226;</span> <b>Versión 5.2.2</b> Se agrega el Atributo <b>ReasonReverse</b> que permite notificar en las Reversas la razón por la cual fue necesario  generarla, el atributo <b>ReasonSequenceBreak</b> que permite indicar la razón por la cual se produce la ruptura de secuencia que podrá generar una reversa si  fuese necesario, y el atributo </b>Reason</b> en la operación <b>Cancel</b>.<br/>   <br/> <span>&#8226;</span> <b>Versión 5.2.1</b> Se agrega el Atributo <b>IsReverse</b> en todos las operaciones reversables.<br/>   <br/><br/> <br/><br/> <br/><p style=\"color:Blue;\">&copy;2019-2021 EVO Payments Inc. All rights reserved.</p>The EVO Payments name, logo and related trademarks and service marks, owned<br /> by EVO Payments, are registered and/or used in the<br /> United States and many foreign countries. All other trademarks,<br /> service marks and trade names referenced in this site are the property<br /> of their respective owners.<br /> <br /> <br /> ANY USE, COPYING OR REPRODUCTION OF THE TRADEMARKS, LOGOS, INFORMATION,<br />  IMAGES OR DESIGNS CONTAINED IN THIS SITE IS STRICTLY<br />  PROHIBITED WITHOUT THE PRIOR WRITTEN PERMISSION OF EVO Payments Inc.<br /> <br /> 

API version: 5.6.1
Contact: integrations@evopayments.mx
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// AuthorizeSaleObjectAuthorizeSale struct for AuthorizeSaleObjectAuthorizeSale
type AuthorizeSaleObjectAuthorizeSale struct {
	// ID que identifica la companía desde donde proviene la petición.
	CompanyIdentification string `json:"CompanyIdentification"`
	// ID que identifica el sistema desde donde proviene la petición.
	SystemIdentification string `json:"SystemIdentification"`
	// ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía.
	BranchIdentification *string `json:"BranchIdentification,omitempty"`
	// ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía.
	POSIdentification *string `json:"POSIdentification,omitempty"`
	// Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos.
	RequestType *string `json:"RequestType,omitempty"`
	// ID generado para la identificación por parte del Plataforma de la información generada en la ejecución de un GetCard o un Payment Method. Será necesario para que un mensaje de Sale, Void, PaymentMethod o Enrollment  identifique el contexto generado y lo utilice para esa operación.
	RequestKey *string `json:"RequestKey,omitempty"`
	// Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible.
	ServiceVersion *string `json:"ServiceVersion,omitempty"`
	// Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma.
	Sequence *string `json:"Sequence,omitempty"`
	// Datos asociados a la seguridad de la transacción o de elementos sensibles.
	Security []SaleObjectSaleSecurity `json:"Security,omitempty"`
	// ID que identifica a un grupo de transacciones que serán confirmadas o canceladas.
	Block *string `json:"Block,omitempty"`
	// En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente.
	RequiredInformation []SaleObjectSaleRequiredInformation `json:"RequiredInformation,omitempty"`
	// Ticket Digitalizado ( Total o parte del mismo por ejemplo la Firma Digitalizada )    codificado en Base64.
	Ticket *string `json:"Ticket,omitempty"`
	// Identificador Unívoco de la Transacción que se quiere Referenciar a la cual pertenece el Ticket Digitalizado. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. Si la firma fue capturada previamente y se envía en el requerimiento de la misma Operación Sale, Authorize*, Void, Return, Adjustment, DebtPayment, VoidDebtPayment o Enrollment no es necesario que se envíe este elemento o campo.
	TicketAnswerKey *string `json:"TicketAnswerKey,omitempty"`
	// Tiempo de espera que el POS espera al PINPAD para obtener la respuesta al requerimiento. 
	Timeout *float32 `json:"Timeout,omitempty"`
	// URL para notificación del comercio
	MerchantNotifyURL *string `json:"MerchantNotifyURL,omitempty"`
	// Identificador de Reversa.
	IsReverse *float32 `json:"IsReverse,omitempty"`
	// Motivo por el cual se requiere generar la reversa.
	ReasonReverse *string `json:"ReasonReverse,omitempty"`
	// Motivo por el cual se requiere romper la secuencia.
	ReasonSequenceBreak *string `json:"ReasonSequenceBreak,omitempty"`
	// Referencia de la transacción para el punto de venta
	Reference *string `json:"Reference,omitempty"`
	// Descripción del tipo de operación que se realizará
	TransactionDescription *string `json:"TransactionDescription,omitempty"`
	// Tipo de punto de venta.
	POSType *string `json:"POSType,omitempty"`
	// Versión del Aplicativo del punto de Venta.
	POSVersion *string `json:"POSVersion,omitempty"`
	// Dirección IP de la Caja o POS.
	POSAddress *string `json:"POSAddress,omitempty"`
	// Número de serie o identificador unívoco del punto de venta.
	POSSerial *string `json:"POSSerial,omitempty"`
	// Coordenadas Geográficas del aplicativo de punto de venta
	POSGEO *string `json:"POSGEO,omitempty"`
	// Tipo de dispositivo utilizado para ingresar los datos de la tarjeta. En función al dispositivo usado, serán realizadas ciertas verificaciones, por lo que ciertos datos serán requeridos. CustomerKeyboard, utilizado para ingreso manual de tarjeta a través de un portal web, por ejemplo - Keyboard, utilizado para ingreso manual de la tarjeta por parte del vendedor - MagneticStripReader, lector de banda de tarjetas por emulación de teclado, u otro valor configurado  que indentifique el dispositivo que se esta utilizando.
	ReadingDeviceType *string `json:"ReadingDeviceType,omitempty"`
	// Indica desde cuando se encuentra operativo o encendido el dispositivo
	ReadingDeviceOperatingFrom *time.Time `json:"ReadingDeviceOperatingFrom,omitempty"`
	// Versión del dispositivo.
	ReadingDeviceVersion *string `json:"ReadingDeviceVersion,omitempty"`
	// Dirección IP o MAC Address del dispositivo.
	ReadingDeviceAddress *string `json:"ReadingDeviceAddress,omitempty"`
	// Número de serie o identificador unívoco del dispositivo.
	ReadingDeviceSerial *string `json:"ReadingDeviceSerial,omitempty"`
	// Coordenadas Geográficas del dispositivo de lectura
	ReadingDeviceGEO *string `json:"ReadingDeviceGEO,omitempty"`
	// Identificador del usuario que está realizando la Transacción.
	UserID *string `json:"UserID,omitempty"`
	// Nombre del usuario que está realizando la Transacción.
	UserName *string `json:"UserName,omitempty"`
	// Importe o Monto de la Transacción.
	Amount float32 `json:"Amount"`
	// Monto con la que se realizó transacción. Si este valor es recibido, la búsqueda de los planes será limitada con este criterio.
	AlternativeAmount *float32 `json:"AlternativeAmount,omitempty"`
	// Monto del dinero en efectivo (cashback).
	CashbackAmount *float32 `json:"CashbackAmount,omitempty"`
	// Importe o Monto de la Propina.
	TipAmount *float32 `json:"TipAmount,omitempty"`
	// Monto sujeto a promoción, monto de la operación.
	PromotedAmount *float32 `json:"PromotedAmount,omitempty"`
	// Código de Moneda - ISO 4217 <https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica <br />   * Num   - Alpha - Description <br />   * '032' - 'ARS' - Pesos Argentinos <br />   * '152' - 'CLP' - Pesos Chilenos <br/>   * '484' - 'MXN' - Pesos Mexicanos <br/>   * '840' - 'USD' - dólares Americanos <br/>   * '878' - 'EUR' - Euros <br/>   * '858' - 'UYU' - Pesos Uruguayos <br/>   * '878' - 'EUR' - Euros <br/>   * '986' - 'BRL' - Real Brasileño
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
	// Cantidad de cuotas en las que será realizada la transacción
	FacilityPayments *float32 `json:"FacilityPayments,omitempty"`
	// Tipo de Plan de Financiación
	FacilityType *string `json:"FacilityType,omitempty"`
	// Código/ID de Plan ( obtenido por la Operación PaymentMethod ) , en caso de ser enviado no se requiere en la Transacción el envío de CurrencyCode ni FacilityPayments
	Plan *string `json:"Plan,omitempty"`
	// Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes )
	CardReadMode *string `json:"CardReadMode,omitempty"`
	// Indican por cada elemento que contiene los datos sensibles, si están encriptados  y también el algoritmo usado. En caso de no estar especificado se asume PLAIN.
	CardGetMode *string `json:"CardGetMode,omitempty"`
	// Número de Tarjeta, en el caso de las respuestas el mismo estará enmascarado.
	CardNumber *string `json:"CardNumber,omitempty"`
	// Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón.   
	CardNumberMasked *string `json:"CardNumberMasked,omitempty"`
	// Número de tarjeta encriptado.                       
	CardNumberEncrypted *string `json:"CardNumberEncrypted,omitempty"`
	// Fecha de vencimiento de la tarjeta. Este dato será necesario si el modo de ingreso fue manual/digitada.
	CardExp *string `json:"CardExp,omitempty"`
	// Tags EMV en formato TLV Leídos.
	CardCryptogram *string `json:"CardCryptogram,omitempty"`
	// Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers.
	CardAppName *string `json:"CardAppName,omitempty"`
	// Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers.
	CardAppIdentifier *string `json:"CardAppIdentifier,omitempty"`
	// Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers.
	CardAppLabel *string `json:"CardAppLabel,omitempty"`
	// Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers.
	CardAuthRequestCryptogram *string `json:"CardAuthRequestCryptogram,omitempty"`
	// Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers.
	CardAuthResponseCryptogram *string `json:"CardAuthResponseCryptogram,omitempty"`
	// Banda Número 1 de la tarjeta. Este dato será necesario si el modo de ingreso fue por banda (deslizando la banda por el lector).
	Track1 *string `json:"Track1,omitempty"`
	// Banda Número 2 de la tarjeta. Este dato será necesario si el modo de ingreso fue por banda (deslizando la banda por el lector).
	Track2 *string `json:"Track2,omitempty"`
	// Tokens.
	InputTokens []SaleObjectSaleInputTokens `json:"InputTokens,omitempty"`
	// Código de seguridad de la tarjeta.
	SecurityCode *string `json:"SecurityCode,omitempty"`
	// PIN block
	Pin *string `json:"Pin,omitempty"`
	// Token asociado a la Credencial Enrolada
	CredentialToken *string `json:"CredentialToken,omitempty"`
	// Emisor del Token asociado a la credencial enrolada
	CredentialIssuerToken *string `json:"CredentialIssuerToken,omitempty"`
	Payer *SaleObjectSalePayer `json:"Payer,omitempty"`
	Customer *SaleObjectSaleCustomer `json:"Customer,omitempty"`
	Seller *SaleObjectSaleSeller `json:"Seller,omitempty"`
	// Detalle de Productos de la Operación.
	Products []SaleObjectSaleProducts `json:"Products,omitempty"`
	// Esquema de Devolución de Impuestos a utilizar en la transacción
	TaxRefundType *string `json:"TaxRefundType,omitempty"`
	// Código de autorización retornado por el Host que resuelve la transacción.
	AuthCode *string `json:"AuthCode,omitempty"`
	// Identificador de facilitador de pagos o Payfac.
	PaymentFacilitatorID *string `json:"PaymentFacilitatorID,omitempty"`
	// Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relacionado a cada uno de los planes disponibles.
	MerchantID *string `json:"MerchantID,omitempty"`
	// Identificador de Terminal por el cual se envía la Transacción al Host.
	TerminalID *string `json:"TerminalID,omitempty"`
	// Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID.
	TerminalTrace *int32 `json:"TerminalTrace,omitempty"`
	// Para aquellos host que exista el concepto de lote, es el número de lote al cual pertenece la transacción.
	SettlementBatchNumber *int32 `json:"SettlementBatchNumber,omitempty"`
}

// NewAuthorizeSaleObjectAuthorizeSale instantiates a new AuthorizeSaleObjectAuthorizeSale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeSaleObjectAuthorizeSale(companyIdentification string, systemIdentification string, amount float32) *AuthorizeSaleObjectAuthorizeSale {
	this := AuthorizeSaleObjectAuthorizeSale{}
	this.CompanyIdentification = companyIdentification
	this.SystemIdentification = systemIdentification
	this.Amount = amount
	return &this
}

// NewAuthorizeSaleObjectAuthorizeSaleWithDefaults instantiates a new AuthorizeSaleObjectAuthorizeSale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeSaleObjectAuthorizeSaleWithDefaults() *AuthorizeSaleObjectAuthorizeSale {
	this := AuthorizeSaleObjectAuthorizeSale{}
	return &this
}

// GetCompanyIdentification returns the CompanyIdentification field value
func (o *AuthorizeSaleObjectAuthorizeSale) GetCompanyIdentification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyIdentification
}

// GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field value
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCompanyIdentificationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyIdentification, true
}

// SetCompanyIdentification sets field value
func (o *AuthorizeSaleObjectAuthorizeSale) SetCompanyIdentification(v string) {
	o.CompanyIdentification = v
}

// GetSystemIdentification returns the SystemIdentification field value
func (o *AuthorizeSaleObjectAuthorizeSale) GetSystemIdentification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemIdentification
}

// GetSystemIdentificationOk returns a tuple with the SystemIdentification field value
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSystemIdentificationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SystemIdentification, true
}

// SetSystemIdentification sets field value
func (o *AuthorizeSaleObjectAuthorizeSale) SetSystemIdentification(v string) {
	o.SystemIdentification = v
}

// GetBranchIdentification returns the BranchIdentification field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetBranchIdentification() string {
	if o == nil || o.BranchIdentification == nil {
		var ret string
		return ret
	}
	return *o.BranchIdentification
}

// GetBranchIdentificationOk returns a tuple with the BranchIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetBranchIdentificationOk() (*string, bool) {
	if o == nil || o.BranchIdentification == nil {
		return nil, false
	}
	return o.BranchIdentification, true
}

// HasBranchIdentification returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasBranchIdentification() bool {
	if o != nil && o.BranchIdentification != nil {
		return true
	}

	return false
}

// SetBranchIdentification gets a reference to the given string and assigns it to the BranchIdentification field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetBranchIdentification(v string) {
	o.BranchIdentification = &v
}

// GetPOSIdentification returns the POSIdentification field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSIdentification() string {
	if o == nil || o.POSIdentification == nil {
		var ret string
		return ret
	}
	return *o.POSIdentification
}

// GetPOSIdentificationOk returns a tuple with the POSIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSIdentificationOk() (*string, bool) {
	if o == nil || o.POSIdentification == nil {
		return nil, false
	}
	return o.POSIdentification, true
}

// HasPOSIdentification returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSIdentification() bool {
	if o != nil && o.POSIdentification != nil {
		return true
	}

	return false
}

// SetPOSIdentification gets a reference to the given string and assigns it to the POSIdentification field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSIdentification(v string) {
	o.POSIdentification = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetRequestType(v string) {
	o.RequestType = &v
}

// GetRequestKey returns the RequestKey field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetRequestKey() string {
	if o == nil || o.RequestKey == nil {
		var ret string
		return ret
	}
	return *o.RequestKey
}

// GetRequestKeyOk returns a tuple with the RequestKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetRequestKeyOk() (*string, bool) {
	if o == nil || o.RequestKey == nil {
		return nil, false
	}
	return o.RequestKey, true
}

// HasRequestKey returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasRequestKey() bool {
	if o != nil && o.RequestKey != nil {
		return true
	}

	return false
}

// SetRequestKey gets a reference to the given string and assigns it to the RequestKey field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetRequestKey(v string) {
	o.RequestKey = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSequence() string {
	if o == nil || o.Sequence == nil {
		var ret string
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSequenceOk() (*string, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given string and assigns it to the Sequence field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetSequence(v string) {
	o.Sequence = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSecurity() []SaleObjectSaleSecurity {
	if o == nil || o.Security == nil {
		var ret []SaleObjectSaleSecurity
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSecurityOk() ([]SaleObjectSaleSecurity, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []SaleObjectSaleSecurity and assigns it to the Security field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetSecurity(v []SaleObjectSaleSecurity) {
	o.Security = v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetBlock() string {
	if o == nil || o.Block == nil {
		var ret string
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetBlockOk() (*string, bool) {
	if o == nil || o.Block == nil {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasBlock() bool {
	if o != nil && o.Block != nil {
		return true
	}

	return false
}

// SetBlock gets a reference to the given string and assigns it to the Block field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetBlock(v string) {
	o.Block = &v
}

// GetRequiredInformation returns the RequiredInformation field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetRequiredInformation() []SaleObjectSaleRequiredInformation {
	if o == nil || o.RequiredInformation == nil {
		var ret []SaleObjectSaleRequiredInformation
		return ret
	}
	return o.RequiredInformation
}

// GetRequiredInformationOk returns a tuple with the RequiredInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetRequiredInformationOk() ([]SaleObjectSaleRequiredInformation, bool) {
	if o == nil || o.RequiredInformation == nil {
		return nil, false
	}
	return o.RequiredInformation, true
}

// HasRequiredInformation returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasRequiredInformation() bool {
	if o != nil && o.RequiredInformation != nil {
		return true
	}

	return false
}

// SetRequiredInformation gets a reference to the given []SaleObjectSaleRequiredInformation and assigns it to the RequiredInformation field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetRequiredInformation(v []SaleObjectSaleRequiredInformation) {
	o.RequiredInformation = v
}

// GetTicket returns the Ticket field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTicket() string {
	if o == nil || o.Ticket == nil {
		var ret string
		return ret
	}
	return *o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTicketOk() (*string, bool) {
	if o == nil || o.Ticket == nil {
		return nil, false
	}
	return o.Ticket, true
}

// HasTicket returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasTicket() bool {
	if o != nil && o.Ticket != nil {
		return true
	}

	return false
}

// SetTicket gets a reference to the given string and assigns it to the Ticket field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetTicket(v string) {
	o.Ticket = &v
}

// GetTicketAnswerKey returns the TicketAnswerKey field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTicketAnswerKey() string {
	if o == nil || o.TicketAnswerKey == nil {
		var ret string
		return ret
	}
	return *o.TicketAnswerKey
}

// GetTicketAnswerKeyOk returns a tuple with the TicketAnswerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTicketAnswerKeyOk() (*string, bool) {
	if o == nil || o.TicketAnswerKey == nil {
		return nil, false
	}
	return o.TicketAnswerKey, true
}

// HasTicketAnswerKey returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasTicketAnswerKey() bool {
	if o != nil && o.TicketAnswerKey != nil {
		return true
	}

	return false
}

// SetTicketAnswerKey gets a reference to the given string and assigns it to the TicketAnswerKey field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetTicketAnswerKey(v string) {
	o.TicketAnswerKey = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTimeout() float32 {
	if o == nil || o.Timeout == nil {
		var ret float32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTimeoutOk() (*float32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given float32 and assigns it to the Timeout field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetTimeout(v float32) {
	o.Timeout = &v
}

// GetMerchantNotifyURL returns the MerchantNotifyURL field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetMerchantNotifyURL() string {
	if o == nil || o.MerchantNotifyURL == nil {
		var ret string
		return ret
	}
	return *o.MerchantNotifyURL
}

// GetMerchantNotifyURLOk returns a tuple with the MerchantNotifyURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetMerchantNotifyURLOk() (*string, bool) {
	if o == nil || o.MerchantNotifyURL == nil {
		return nil, false
	}
	return o.MerchantNotifyURL, true
}

// HasMerchantNotifyURL returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasMerchantNotifyURL() bool {
	if o != nil && o.MerchantNotifyURL != nil {
		return true
	}

	return false
}

// SetMerchantNotifyURL gets a reference to the given string and assigns it to the MerchantNotifyURL field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetMerchantNotifyURL(v string) {
	o.MerchantNotifyURL = &v
}

// GetIsReverse returns the IsReverse field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetIsReverse() float32 {
	if o == nil || o.IsReverse == nil {
		var ret float32
		return ret
	}
	return *o.IsReverse
}

// GetIsReverseOk returns a tuple with the IsReverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetIsReverseOk() (*float32, bool) {
	if o == nil || o.IsReverse == nil {
		return nil, false
	}
	return o.IsReverse, true
}

// HasIsReverse returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasIsReverse() bool {
	if o != nil && o.IsReverse != nil {
		return true
	}

	return false
}

// SetIsReverse gets a reference to the given float32 and assigns it to the IsReverse field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetIsReverse(v float32) {
	o.IsReverse = &v
}

// GetReasonReverse returns the ReasonReverse field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReasonReverse() string {
	if o == nil || o.ReasonReverse == nil {
		var ret string
		return ret
	}
	return *o.ReasonReverse
}

// GetReasonReverseOk returns a tuple with the ReasonReverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReasonReverseOk() (*string, bool) {
	if o == nil || o.ReasonReverse == nil {
		return nil, false
	}
	return o.ReasonReverse, true
}

// HasReasonReverse returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasReasonReverse() bool {
	if o != nil && o.ReasonReverse != nil {
		return true
	}

	return false
}

// SetReasonReverse gets a reference to the given string and assigns it to the ReasonReverse field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetReasonReverse(v string) {
	o.ReasonReverse = &v
}

// GetReasonSequenceBreak returns the ReasonSequenceBreak field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReasonSequenceBreak() string {
	if o == nil || o.ReasonSequenceBreak == nil {
		var ret string
		return ret
	}
	return *o.ReasonSequenceBreak
}

// GetReasonSequenceBreakOk returns a tuple with the ReasonSequenceBreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReasonSequenceBreakOk() (*string, bool) {
	if o == nil || o.ReasonSequenceBreak == nil {
		return nil, false
	}
	return o.ReasonSequenceBreak, true
}

// HasReasonSequenceBreak returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasReasonSequenceBreak() bool {
	if o != nil && o.ReasonSequenceBreak != nil {
		return true
	}

	return false
}

// SetReasonSequenceBreak gets a reference to the given string and assigns it to the ReasonSequenceBreak field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetReasonSequenceBreak(v string) {
	o.ReasonSequenceBreak = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetReference(v string) {
	o.Reference = &v
}

// GetTransactionDescription returns the TransactionDescription field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTransactionDescription() string {
	if o == nil || o.TransactionDescription == nil {
		var ret string
		return ret
	}
	return *o.TransactionDescription
}

// GetTransactionDescriptionOk returns a tuple with the TransactionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTransactionDescriptionOk() (*string, bool) {
	if o == nil || o.TransactionDescription == nil {
		return nil, false
	}
	return o.TransactionDescription, true
}

// HasTransactionDescription returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasTransactionDescription() bool {
	if o != nil && o.TransactionDescription != nil {
		return true
	}

	return false
}

// SetTransactionDescription gets a reference to the given string and assigns it to the TransactionDescription field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetTransactionDescription(v string) {
	o.TransactionDescription = &v
}

// GetPOSType returns the POSType field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSType() string {
	if o == nil || o.POSType == nil {
		var ret string
		return ret
	}
	return *o.POSType
}

// GetPOSTypeOk returns a tuple with the POSType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSTypeOk() (*string, bool) {
	if o == nil || o.POSType == nil {
		return nil, false
	}
	return o.POSType, true
}

// HasPOSType returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSType() bool {
	if o != nil && o.POSType != nil {
		return true
	}

	return false
}

// SetPOSType gets a reference to the given string and assigns it to the POSType field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSType(v string) {
	o.POSType = &v
}

// GetPOSVersion returns the POSVersion field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSVersion() string {
	if o == nil || o.POSVersion == nil {
		var ret string
		return ret
	}
	return *o.POSVersion
}

// GetPOSVersionOk returns a tuple with the POSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSVersionOk() (*string, bool) {
	if o == nil || o.POSVersion == nil {
		return nil, false
	}
	return o.POSVersion, true
}

// HasPOSVersion returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSVersion() bool {
	if o != nil && o.POSVersion != nil {
		return true
	}

	return false
}

// SetPOSVersion gets a reference to the given string and assigns it to the POSVersion field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSVersion(v string) {
	o.POSVersion = &v
}

// GetPOSAddress returns the POSAddress field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSAddress() string {
	if o == nil || o.POSAddress == nil {
		var ret string
		return ret
	}
	return *o.POSAddress
}

// GetPOSAddressOk returns a tuple with the POSAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSAddressOk() (*string, bool) {
	if o == nil || o.POSAddress == nil {
		return nil, false
	}
	return o.POSAddress, true
}

// HasPOSAddress returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSAddress() bool {
	if o != nil && o.POSAddress != nil {
		return true
	}

	return false
}

// SetPOSAddress gets a reference to the given string and assigns it to the POSAddress field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSAddress(v string) {
	o.POSAddress = &v
}

// GetPOSSerial returns the POSSerial field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSSerial() string {
	if o == nil || o.POSSerial == nil {
		var ret string
		return ret
	}
	return *o.POSSerial
}

// GetPOSSerialOk returns a tuple with the POSSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSSerialOk() (*string, bool) {
	if o == nil || o.POSSerial == nil {
		return nil, false
	}
	return o.POSSerial, true
}

// HasPOSSerial returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSSerial() bool {
	if o != nil && o.POSSerial != nil {
		return true
	}

	return false
}

// SetPOSSerial gets a reference to the given string and assigns it to the POSSerial field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSSerial(v string) {
	o.POSSerial = &v
}

// GetPOSGEO returns the POSGEO field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSGEO() string {
	if o == nil || o.POSGEO == nil {
		var ret string
		return ret
	}
	return *o.POSGEO
}

// GetPOSGEOOk returns a tuple with the POSGEO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPOSGEOOk() (*string, bool) {
	if o == nil || o.POSGEO == nil {
		return nil, false
	}
	return o.POSGEO, true
}

// HasPOSGEO returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPOSGEO() bool {
	if o != nil && o.POSGEO != nil {
		return true
	}

	return false
}

// SetPOSGEO gets a reference to the given string and assigns it to the POSGEO field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPOSGEO(v string) {
	o.POSGEO = &v
}

// GetReadingDeviceType returns the ReadingDeviceType field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceType() string {
	if o == nil || o.ReadingDeviceType == nil {
		var ret string
		return ret
	}
	return *o.ReadingDeviceType
}

// GetReadingDeviceTypeOk returns a tuple with the ReadingDeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceTypeOk() (*string, bool) {
	if o == nil || o.ReadingDeviceType == nil {
		return nil, false
	}
	return o.ReadingDeviceType, true
}

// HasReadingDeviceType returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceType() bool {
	if o != nil && o.ReadingDeviceType != nil {
		return true
	}

	return false
}

// SetReadingDeviceType gets a reference to the given string and assigns it to the ReadingDeviceType field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceType(v string) {
	o.ReadingDeviceType = &v
}

// GetReadingDeviceOperatingFrom returns the ReadingDeviceOperatingFrom field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceOperatingFrom() time.Time {
	if o == nil || o.ReadingDeviceOperatingFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ReadingDeviceOperatingFrom
}

// GetReadingDeviceOperatingFromOk returns a tuple with the ReadingDeviceOperatingFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceOperatingFromOk() (*time.Time, bool) {
	if o == nil || o.ReadingDeviceOperatingFrom == nil {
		return nil, false
	}
	return o.ReadingDeviceOperatingFrom, true
}

// HasReadingDeviceOperatingFrom returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceOperatingFrom() bool {
	if o != nil && o.ReadingDeviceOperatingFrom != nil {
		return true
	}

	return false
}

// SetReadingDeviceOperatingFrom gets a reference to the given time.Time and assigns it to the ReadingDeviceOperatingFrom field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceOperatingFrom(v time.Time) {
	o.ReadingDeviceOperatingFrom = &v
}

// GetReadingDeviceVersion returns the ReadingDeviceVersion field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceVersion() string {
	if o == nil || o.ReadingDeviceVersion == nil {
		var ret string
		return ret
	}
	return *o.ReadingDeviceVersion
}

// GetReadingDeviceVersionOk returns a tuple with the ReadingDeviceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceVersionOk() (*string, bool) {
	if o == nil || o.ReadingDeviceVersion == nil {
		return nil, false
	}
	return o.ReadingDeviceVersion, true
}

// HasReadingDeviceVersion returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceVersion() bool {
	if o != nil && o.ReadingDeviceVersion != nil {
		return true
	}

	return false
}

// SetReadingDeviceVersion gets a reference to the given string and assigns it to the ReadingDeviceVersion field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceVersion(v string) {
	o.ReadingDeviceVersion = &v
}

// GetReadingDeviceAddress returns the ReadingDeviceAddress field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceAddress() string {
	if o == nil || o.ReadingDeviceAddress == nil {
		var ret string
		return ret
	}
	return *o.ReadingDeviceAddress
}

// GetReadingDeviceAddressOk returns a tuple with the ReadingDeviceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceAddressOk() (*string, bool) {
	if o == nil || o.ReadingDeviceAddress == nil {
		return nil, false
	}
	return o.ReadingDeviceAddress, true
}

// HasReadingDeviceAddress returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceAddress() bool {
	if o != nil && o.ReadingDeviceAddress != nil {
		return true
	}

	return false
}

// SetReadingDeviceAddress gets a reference to the given string and assigns it to the ReadingDeviceAddress field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceAddress(v string) {
	o.ReadingDeviceAddress = &v
}

// GetReadingDeviceSerial returns the ReadingDeviceSerial field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceSerial() string {
	if o == nil || o.ReadingDeviceSerial == nil {
		var ret string
		return ret
	}
	return *o.ReadingDeviceSerial
}

// GetReadingDeviceSerialOk returns a tuple with the ReadingDeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceSerialOk() (*string, bool) {
	if o == nil || o.ReadingDeviceSerial == nil {
		return nil, false
	}
	return o.ReadingDeviceSerial, true
}

// HasReadingDeviceSerial returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceSerial() bool {
	if o != nil && o.ReadingDeviceSerial != nil {
		return true
	}

	return false
}

// SetReadingDeviceSerial gets a reference to the given string and assigns it to the ReadingDeviceSerial field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceSerial(v string) {
	o.ReadingDeviceSerial = &v
}

// GetReadingDeviceGEO returns the ReadingDeviceGEO field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceGEO() string {
	if o == nil || o.ReadingDeviceGEO == nil {
		var ret string
		return ret
	}
	return *o.ReadingDeviceGEO
}

// GetReadingDeviceGEOOk returns a tuple with the ReadingDeviceGEO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetReadingDeviceGEOOk() (*string, bool) {
	if o == nil || o.ReadingDeviceGEO == nil {
		return nil, false
	}
	return o.ReadingDeviceGEO, true
}

// HasReadingDeviceGEO returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasReadingDeviceGEO() bool {
	if o != nil && o.ReadingDeviceGEO != nil {
		return true
	}

	return false
}

// SetReadingDeviceGEO gets a reference to the given string and assigns it to the ReadingDeviceGEO field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetReadingDeviceGEO(v string) {
	o.ReadingDeviceGEO = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetUserID() string {
	if o == nil || o.UserID == nil {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetUserIDOk() (*string, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetUserID(v string) {
	o.UserID = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetUserName(v string) {
	o.UserName = &v
}

// GetAmount returns the Amount field value
func (o *AuthorizeSaleObjectAuthorizeSale) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AuthorizeSaleObjectAuthorizeSale) SetAmount(v float32) {
	o.Amount = v
}

// GetAlternativeAmount returns the AlternativeAmount field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetAlternativeAmount() float32 {
	if o == nil || o.AlternativeAmount == nil {
		var ret float32
		return ret
	}
	return *o.AlternativeAmount
}

// GetAlternativeAmountOk returns a tuple with the AlternativeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetAlternativeAmountOk() (*float32, bool) {
	if o == nil || o.AlternativeAmount == nil {
		return nil, false
	}
	return o.AlternativeAmount, true
}

// HasAlternativeAmount returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasAlternativeAmount() bool {
	if o != nil && o.AlternativeAmount != nil {
		return true
	}

	return false
}

// SetAlternativeAmount gets a reference to the given float32 and assigns it to the AlternativeAmount field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetAlternativeAmount(v float32) {
	o.AlternativeAmount = &v
}

// GetCashbackAmount returns the CashbackAmount field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCashbackAmount() float32 {
	if o == nil || o.CashbackAmount == nil {
		var ret float32
		return ret
	}
	return *o.CashbackAmount
}

// GetCashbackAmountOk returns a tuple with the CashbackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCashbackAmountOk() (*float32, bool) {
	if o == nil || o.CashbackAmount == nil {
		return nil, false
	}
	return o.CashbackAmount, true
}

// HasCashbackAmount returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCashbackAmount() bool {
	if o != nil && o.CashbackAmount != nil {
		return true
	}

	return false
}

// SetCashbackAmount gets a reference to the given float32 and assigns it to the CashbackAmount field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCashbackAmount(v float32) {
	o.CashbackAmount = &v
}

// GetTipAmount returns the TipAmount field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTipAmount() float32 {
	if o == nil || o.TipAmount == nil {
		var ret float32
		return ret
	}
	return *o.TipAmount
}

// GetTipAmountOk returns a tuple with the TipAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTipAmountOk() (*float32, bool) {
	if o == nil || o.TipAmount == nil {
		return nil, false
	}
	return o.TipAmount, true
}

// HasTipAmount returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasTipAmount() bool {
	if o != nil && o.TipAmount != nil {
		return true
	}

	return false
}

// SetTipAmount gets a reference to the given float32 and assigns it to the TipAmount field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetTipAmount(v float32) {
	o.TipAmount = &v
}

// GetPromotedAmount returns the PromotedAmount field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPromotedAmount() float32 {
	if o == nil || o.PromotedAmount == nil {
		var ret float32
		return ret
	}
	return *o.PromotedAmount
}

// GetPromotedAmountOk returns a tuple with the PromotedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPromotedAmountOk() (*float32, bool) {
	if o == nil || o.PromotedAmount == nil {
		return nil, false
	}
	return o.PromotedAmount, true
}

// HasPromotedAmount returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPromotedAmount() bool {
	if o != nil && o.PromotedAmount != nil {
		return true
	}

	return false
}

// SetPromotedAmount gets a reference to the given float32 and assigns it to the PromotedAmount field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPromotedAmount(v float32) {
	o.PromotedAmount = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetFacilityPayments returns the FacilityPayments field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetFacilityPayments() float32 {
	if o == nil || o.FacilityPayments == nil {
		var ret float32
		return ret
	}
	return *o.FacilityPayments
}

// GetFacilityPaymentsOk returns a tuple with the FacilityPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetFacilityPaymentsOk() (*float32, bool) {
	if o == nil || o.FacilityPayments == nil {
		return nil, false
	}
	return o.FacilityPayments, true
}

// HasFacilityPayments returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasFacilityPayments() bool {
	if o != nil && o.FacilityPayments != nil {
		return true
	}

	return false
}

// SetFacilityPayments gets a reference to the given float32 and assigns it to the FacilityPayments field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetFacilityPayments(v float32) {
	o.FacilityPayments = &v
}

// GetFacilityType returns the FacilityType field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetFacilityType() string {
	if o == nil || o.FacilityType == nil {
		var ret string
		return ret
	}
	return *o.FacilityType
}

// GetFacilityTypeOk returns a tuple with the FacilityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetFacilityTypeOk() (*string, bool) {
	if o == nil || o.FacilityType == nil {
		return nil, false
	}
	return o.FacilityType, true
}

// HasFacilityType returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasFacilityType() bool {
	if o != nil && o.FacilityType != nil {
		return true
	}

	return false
}

// SetFacilityType gets a reference to the given string and assigns it to the FacilityType field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetFacilityType(v string) {
	o.FacilityType = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPlan() string {
	if o == nil || o.Plan == nil {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPlanOk() (*string, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPlan(v string) {
	o.Plan = &v
}

// GetCardReadMode returns the CardReadMode field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardReadMode() string {
	if o == nil || o.CardReadMode == nil {
		var ret string
		return ret
	}
	return *o.CardReadMode
}

// GetCardReadModeOk returns a tuple with the CardReadMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardReadModeOk() (*string, bool) {
	if o == nil || o.CardReadMode == nil {
		return nil, false
	}
	return o.CardReadMode, true
}

// HasCardReadMode returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardReadMode() bool {
	if o != nil && o.CardReadMode != nil {
		return true
	}

	return false
}

// SetCardReadMode gets a reference to the given string and assigns it to the CardReadMode field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardReadMode(v string) {
	o.CardReadMode = &v
}

// GetCardGetMode returns the CardGetMode field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardGetMode() string {
	if o == nil || o.CardGetMode == nil {
		var ret string
		return ret
	}
	return *o.CardGetMode
}

// GetCardGetModeOk returns a tuple with the CardGetMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardGetModeOk() (*string, bool) {
	if o == nil || o.CardGetMode == nil {
		return nil, false
	}
	return o.CardGetMode, true
}

// HasCardGetMode returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardGetMode() bool {
	if o != nil && o.CardGetMode != nil {
		return true
	}

	return false
}

// SetCardGetMode gets a reference to the given string and assigns it to the CardGetMode field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardGetMode(v string) {
	o.CardGetMode = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumber() string {
	if o == nil || o.CardNumber == nil {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumberOk() (*string, bool) {
	if o == nil || o.CardNumber == nil {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardNumber() bool {
	if o != nil && o.CardNumber != nil {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetCardNumberMasked returns the CardNumberMasked field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumberMasked() string {
	if o == nil || o.CardNumberMasked == nil {
		var ret string
		return ret
	}
	return *o.CardNumberMasked
}

// GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumberMaskedOk() (*string, bool) {
	if o == nil || o.CardNumberMasked == nil {
		return nil, false
	}
	return o.CardNumberMasked, true
}

// HasCardNumberMasked returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardNumberMasked() bool {
	if o != nil && o.CardNumberMasked != nil {
		return true
	}

	return false
}

// SetCardNumberMasked gets a reference to the given string and assigns it to the CardNumberMasked field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardNumberMasked(v string) {
	o.CardNumberMasked = &v
}

// GetCardNumberEncrypted returns the CardNumberEncrypted field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumberEncrypted() string {
	if o == nil || o.CardNumberEncrypted == nil {
		var ret string
		return ret
	}
	return *o.CardNumberEncrypted
}

// GetCardNumberEncryptedOk returns a tuple with the CardNumberEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardNumberEncryptedOk() (*string, bool) {
	if o == nil || o.CardNumberEncrypted == nil {
		return nil, false
	}
	return o.CardNumberEncrypted, true
}

// HasCardNumberEncrypted returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardNumberEncrypted() bool {
	if o != nil && o.CardNumberEncrypted != nil {
		return true
	}

	return false
}

// SetCardNumberEncrypted gets a reference to the given string and assigns it to the CardNumberEncrypted field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardNumberEncrypted(v string) {
	o.CardNumberEncrypted = &v
}

// GetCardExp returns the CardExp field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardExp() string {
	if o == nil || o.CardExp == nil {
		var ret string
		return ret
	}
	return *o.CardExp
}

// GetCardExpOk returns a tuple with the CardExp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardExpOk() (*string, bool) {
	if o == nil || o.CardExp == nil {
		return nil, false
	}
	return o.CardExp, true
}

// HasCardExp returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardExp() bool {
	if o != nil && o.CardExp != nil {
		return true
	}

	return false
}

// SetCardExp gets a reference to the given string and assigns it to the CardExp field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardExp(v string) {
	o.CardExp = &v
}

// GetCardCryptogram returns the CardCryptogram field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardCryptogram() string {
	if o == nil || o.CardCryptogram == nil {
		var ret string
		return ret
	}
	return *o.CardCryptogram
}

// GetCardCryptogramOk returns a tuple with the CardCryptogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardCryptogramOk() (*string, bool) {
	if o == nil || o.CardCryptogram == nil {
		return nil, false
	}
	return o.CardCryptogram, true
}

// HasCardCryptogram returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardCryptogram() bool {
	if o != nil && o.CardCryptogram != nil {
		return true
	}

	return false
}

// SetCardCryptogram gets a reference to the given string and assigns it to the CardCryptogram field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardCryptogram(v string) {
	o.CardCryptogram = &v
}

// GetCardAppName returns the CardAppName field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppName() string {
	if o == nil || o.CardAppName == nil {
		var ret string
		return ret
	}
	return *o.CardAppName
}

// GetCardAppNameOk returns a tuple with the CardAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppNameOk() (*string, bool) {
	if o == nil || o.CardAppName == nil {
		return nil, false
	}
	return o.CardAppName, true
}

// HasCardAppName returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardAppName() bool {
	if o != nil && o.CardAppName != nil {
		return true
	}

	return false
}

// SetCardAppName gets a reference to the given string and assigns it to the CardAppName field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardAppName(v string) {
	o.CardAppName = &v
}

// GetCardAppIdentifier returns the CardAppIdentifier field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppIdentifier() string {
	if o == nil || o.CardAppIdentifier == nil {
		var ret string
		return ret
	}
	return *o.CardAppIdentifier
}

// GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppIdentifierOk() (*string, bool) {
	if o == nil || o.CardAppIdentifier == nil {
		return nil, false
	}
	return o.CardAppIdentifier, true
}

// HasCardAppIdentifier returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardAppIdentifier() bool {
	if o != nil && o.CardAppIdentifier != nil {
		return true
	}

	return false
}

// SetCardAppIdentifier gets a reference to the given string and assigns it to the CardAppIdentifier field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardAppIdentifier(v string) {
	o.CardAppIdentifier = &v
}

// GetCardAppLabel returns the CardAppLabel field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppLabel() string {
	if o == nil || o.CardAppLabel == nil {
		var ret string
		return ret
	}
	return *o.CardAppLabel
}

// GetCardAppLabelOk returns a tuple with the CardAppLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAppLabelOk() (*string, bool) {
	if o == nil || o.CardAppLabel == nil {
		return nil, false
	}
	return o.CardAppLabel, true
}

// HasCardAppLabel returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardAppLabel() bool {
	if o != nil && o.CardAppLabel != nil {
		return true
	}

	return false
}

// SetCardAppLabel gets a reference to the given string and assigns it to the CardAppLabel field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardAppLabel(v string) {
	o.CardAppLabel = &v
}

// GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAuthRequestCryptogram() string {
	if o == nil || o.CardAuthRequestCryptogram == nil {
		var ret string
		return ret
	}
	return *o.CardAuthRequestCryptogram
}

// GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAuthRequestCryptogramOk() (*string, bool) {
	if o == nil || o.CardAuthRequestCryptogram == nil {
		return nil, false
	}
	return o.CardAuthRequestCryptogram, true
}

// HasCardAuthRequestCryptogram returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardAuthRequestCryptogram() bool {
	if o != nil && o.CardAuthRequestCryptogram != nil {
		return true
	}

	return false
}

// SetCardAuthRequestCryptogram gets a reference to the given string and assigns it to the CardAuthRequestCryptogram field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardAuthRequestCryptogram(v string) {
	o.CardAuthRequestCryptogram = &v
}

// GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAuthResponseCryptogram() string {
	if o == nil || o.CardAuthResponseCryptogram == nil {
		var ret string
		return ret
	}
	return *o.CardAuthResponseCryptogram
}

// GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCardAuthResponseCryptogramOk() (*string, bool) {
	if o == nil || o.CardAuthResponseCryptogram == nil {
		return nil, false
	}
	return o.CardAuthResponseCryptogram, true
}

// HasCardAuthResponseCryptogram returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCardAuthResponseCryptogram() bool {
	if o != nil && o.CardAuthResponseCryptogram != nil {
		return true
	}

	return false
}

// SetCardAuthResponseCryptogram gets a reference to the given string and assigns it to the CardAuthResponseCryptogram field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCardAuthResponseCryptogram(v string) {
	o.CardAuthResponseCryptogram = &v
}

// GetTrack1 returns the Track1 field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTrack1() string {
	if o == nil || o.Track1 == nil {
		var ret string
		return ret
	}
	return *o.Track1
}

// GetTrack1Ok returns a tuple with the Track1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTrack1Ok() (*string, bool) {
	if o == nil || o.Track1 == nil {
		return nil, false
	}
	return o.Track1, true
}

// HasTrack1 returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasTrack1() bool {
	if o != nil && o.Track1 != nil {
		return true
	}

	return false
}

// SetTrack1 gets a reference to the given string and assigns it to the Track1 field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetTrack1(v string) {
	o.Track1 = &v
}

// GetTrack2 returns the Track2 field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTrack2() string {
	if o == nil || o.Track2 == nil {
		var ret string
		return ret
	}
	return *o.Track2
}

// GetTrack2Ok returns a tuple with the Track2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTrack2Ok() (*string, bool) {
	if o == nil || o.Track2 == nil {
		return nil, false
	}
	return o.Track2, true
}

// HasTrack2 returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasTrack2() bool {
	if o != nil && o.Track2 != nil {
		return true
	}

	return false
}

// SetTrack2 gets a reference to the given string and assigns it to the Track2 field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetTrack2(v string) {
	o.Track2 = &v
}

// GetInputTokens returns the InputTokens field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetInputTokens() []SaleObjectSaleInputTokens {
	if o == nil || o.InputTokens == nil {
		var ret []SaleObjectSaleInputTokens
		return ret
	}
	return o.InputTokens
}

// GetInputTokensOk returns a tuple with the InputTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetInputTokensOk() ([]SaleObjectSaleInputTokens, bool) {
	if o == nil || o.InputTokens == nil {
		return nil, false
	}
	return o.InputTokens, true
}

// HasInputTokens returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasInputTokens() bool {
	if o != nil && o.InputTokens != nil {
		return true
	}

	return false
}

// SetInputTokens gets a reference to the given []SaleObjectSaleInputTokens and assigns it to the InputTokens field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetInputTokens(v []SaleObjectSaleInputTokens) {
	o.InputTokens = v
}

// GetSecurityCode returns the SecurityCode field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSecurityCode() string {
	if o == nil || o.SecurityCode == nil {
		var ret string
		return ret
	}
	return *o.SecurityCode
}

// GetSecurityCodeOk returns a tuple with the SecurityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSecurityCodeOk() (*string, bool) {
	if o == nil || o.SecurityCode == nil {
		return nil, false
	}
	return o.SecurityCode, true
}

// HasSecurityCode returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasSecurityCode() bool {
	if o != nil && o.SecurityCode != nil {
		return true
	}

	return false
}

// SetSecurityCode gets a reference to the given string and assigns it to the SecurityCode field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetSecurityCode(v string) {
	o.SecurityCode = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPin() string {
	if o == nil || o.Pin == nil {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPinOk() (*string, bool) {
	if o == nil || o.Pin == nil {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPin() bool {
	if o != nil && o.Pin != nil {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPin(v string) {
	o.Pin = &v
}

// GetCredentialToken returns the CredentialToken field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCredentialToken() string {
	if o == nil || o.CredentialToken == nil {
		var ret string
		return ret
	}
	return *o.CredentialToken
}

// GetCredentialTokenOk returns a tuple with the CredentialToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCredentialTokenOk() (*string, bool) {
	if o == nil || o.CredentialToken == nil {
		return nil, false
	}
	return o.CredentialToken, true
}

// HasCredentialToken returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCredentialToken() bool {
	if o != nil && o.CredentialToken != nil {
		return true
	}

	return false
}

// SetCredentialToken gets a reference to the given string and assigns it to the CredentialToken field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCredentialToken(v string) {
	o.CredentialToken = &v
}

// GetCredentialIssuerToken returns the CredentialIssuerToken field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCredentialIssuerToken() string {
	if o == nil || o.CredentialIssuerToken == nil {
		var ret string
		return ret
	}
	return *o.CredentialIssuerToken
}

// GetCredentialIssuerTokenOk returns a tuple with the CredentialIssuerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCredentialIssuerTokenOk() (*string, bool) {
	if o == nil || o.CredentialIssuerToken == nil {
		return nil, false
	}
	return o.CredentialIssuerToken, true
}

// HasCredentialIssuerToken returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCredentialIssuerToken() bool {
	if o != nil && o.CredentialIssuerToken != nil {
		return true
	}

	return false
}

// SetCredentialIssuerToken gets a reference to the given string and assigns it to the CredentialIssuerToken field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCredentialIssuerToken(v string) {
	o.CredentialIssuerToken = &v
}

// GetPayer returns the Payer field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPayer() SaleObjectSalePayer {
	if o == nil || o.Payer == nil {
		var ret SaleObjectSalePayer
		return ret
	}
	return *o.Payer
}

// GetPayerOk returns a tuple with the Payer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPayerOk() (*SaleObjectSalePayer, bool) {
	if o == nil || o.Payer == nil {
		return nil, false
	}
	return o.Payer, true
}

// HasPayer returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPayer() bool {
	if o != nil && o.Payer != nil {
		return true
	}

	return false
}

// SetPayer gets a reference to the given SaleObjectSalePayer and assigns it to the Payer field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPayer(v SaleObjectSalePayer) {
	o.Payer = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCustomer() SaleObjectSaleCustomer {
	if o == nil || o.Customer == nil {
		var ret SaleObjectSaleCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetCustomerOk() (*SaleObjectSaleCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given SaleObjectSaleCustomer and assigns it to the Customer field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetCustomer(v SaleObjectSaleCustomer) {
	o.Customer = &v
}

// GetSeller returns the Seller field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSeller() SaleObjectSaleSeller {
	if o == nil || o.Seller == nil {
		var ret SaleObjectSaleSeller
		return ret
	}
	return *o.Seller
}

// GetSellerOk returns a tuple with the Seller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSellerOk() (*SaleObjectSaleSeller, bool) {
	if o == nil || o.Seller == nil {
		return nil, false
	}
	return o.Seller, true
}

// HasSeller returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasSeller() bool {
	if o != nil && o.Seller != nil {
		return true
	}

	return false
}

// SetSeller gets a reference to the given SaleObjectSaleSeller and assigns it to the Seller field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetSeller(v SaleObjectSaleSeller) {
	o.Seller = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetProducts() []SaleObjectSaleProducts {
	if o == nil || o.Products == nil {
		var ret []SaleObjectSaleProducts
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetProductsOk() ([]SaleObjectSaleProducts, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SaleObjectSaleProducts and assigns it to the Products field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetProducts(v []SaleObjectSaleProducts) {
	o.Products = v
}

// GetTaxRefundType returns the TaxRefundType field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTaxRefundType() string {
	if o == nil || o.TaxRefundType == nil {
		var ret string
		return ret
	}
	return *o.TaxRefundType
}

// GetTaxRefundTypeOk returns a tuple with the TaxRefundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTaxRefundTypeOk() (*string, bool) {
	if o == nil || o.TaxRefundType == nil {
		return nil, false
	}
	return o.TaxRefundType, true
}

// HasTaxRefundType returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasTaxRefundType() bool {
	if o != nil && o.TaxRefundType != nil {
		return true
	}

	return false
}

// SetTaxRefundType gets a reference to the given string and assigns it to the TaxRefundType field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetTaxRefundType(v string) {
	o.TaxRefundType = &v
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetAuthCode() string {
	if o == nil || o.AuthCode == nil {
		var ret string
		return ret
	}
	return *o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetAuthCodeOk() (*string, bool) {
	if o == nil || o.AuthCode == nil {
		return nil, false
	}
	return o.AuthCode, true
}

// HasAuthCode returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasAuthCode() bool {
	if o != nil && o.AuthCode != nil {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given string and assigns it to the AuthCode field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetAuthCode(v string) {
	o.AuthCode = &v
}

// GetPaymentFacilitatorID returns the PaymentFacilitatorID field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPaymentFacilitatorID() string {
	if o == nil || o.PaymentFacilitatorID == nil {
		var ret string
		return ret
	}
	return *o.PaymentFacilitatorID
}

// GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetPaymentFacilitatorIDOk() (*string, bool) {
	if o == nil || o.PaymentFacilitatorID == nil {
		return nil, false
	}
	return o.PaymentFacilitatorID, true
}

// HasPaymentFacilitatorID returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasPaymentFacilitatorID() bool {
	if o != nil && o.PaymentFacilitatorID != nil {
		return true
	}

	return false
}

// SetPaymentFacilitatorID gets a reference to the given string and assigns it to the PaymentFacilitatorID field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetPaymentFacilitatorID(v string) {
	o.PaymentFacilitatorID = &v
}

// GetMerchantID returns the MerchantID field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetMerchantID() string {
	if o == nil || o.MerchantID == nil {
		var ret string
		return ret
	}
	return *o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetMerchantIDOk() (*string, bool) {
	if o == nil || o.MerchantID == nil {
		return nil, false
	}
	return o.MerchantID, true
}

// HasMerchantID returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasMerchantID() bool {
	if o != nil && o.MerchantID != nil {
		return true
	}

	return false
}

// SetMerchantID gets a reference to the given string and assigns it to the MerchantID field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetMerchantID(v string) {
	o.MerchantID = &v
}

// GetTerminalID returns the TerminalID field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTerminalID() string {
	if o == nil || o.TerminalID == nil {
		var ret string
		return ret
	}
	return *o.TerminalID
}

// GetTerminalIDOk returns a tuple with the TerminalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTerminalIDOk() (*string, bool) {
	if o == nil || o.TerminalID == nil {
		return nil, false
	}
	return o.TerminalID, true
}

// HasTerminalID returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasTerminalID() bool {
	if o != nil && o.TerminalID != nil {
		return true
	}

	return false
}

// SetTerminalID gets a reference to the given string and assigns it to the TerminalID field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetTerminalID(v string) {
	o.TerminalID = &v
}

// GetTerminalTrace returns the TerminalTrace field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTerminalTrace() int32 {
	if o == nil || o.TerminalTrace == nil {
		var ret int32
		return ret
	}
	return *o.TerminalTrace
}

// GetTerminalTraceOk returns a tuple with the TerminalTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetTerminalTraceOk() (*int32, bool) {
	if o == nil || o.TerminalTrace == nil {
		return nil, false
	}
	return o.TerminalTrace, true
}

// HasTerminalTrace returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasTerminalTrace() bool {
	if o != nil && o.TerminalTrace != nil {
		return true
	}

	return false
}

// SetTerminalTrace gets a reference to the given int32 and assigns it to the TerminalTrace field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetTerminalTrace(v int32) {
	o.TerminalTrace = &v
}

// GetSettlementBatchNumber returns the SettlementBatchNumber field value if set, zero value otherwise.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSettlementBatchNumber() int32 {
	if o == nil || o.SettlementBatchNumber == nil {
		var ret int32
		return ret
	}
	return *o.SettlementBatchNumber
}

// GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) GetSettlementBatchNumberOk() (*int32, bool) {
	if o == nil || o.SettlementBatchNumber == nil {
		return nil, false
	}
	return o.SettlementBatchNumber, true
}

// HasSettlementBatchNumber returns a boolean if a field has been set.
func (o *AuthorizeSaleObjectAuthorizeSale) HasSettlementBatchNumber() bool {
	if o != nil && o.SettlementBatchNumber != nil {
		return true
	}

	return false
}

// SetSettlementBatchNumber gets a reference to the given int32 and assigns it to the SettlementBatchNumber field.
func (o *AuthorizeSaleObjectAuthorizeSale) SetSettlementBatchNumber(v int32) {
	o.SettlementBatchNumber = &v
}

func (o AuthorizeSaleObjectAuthorizeSale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["CompanyIdentification"] = o.CompanyIdentification
	}
	if true {
		toSerialize["SystemIdentification"] = o.SystemIdentification
	}
	if o.BranchIdentification != nil {
		toSerialize["BranchIdentification"] = o.BranchIdentification
	}
	if o.POSIdentification != nil {
		toSerialize["POSIdentification"] = o.POSIdentification
	}
	if o.RequestType != nil {
		toSerialize["RequestType"] = o.RequestType
	}
	if o.RequestKey != nil {
		toSerialize["RequestKey"] = o.RequestKey
	}
	if o.ServiceVersion != nil {
		toSerialize["ServiceVersion"] = o.ServiceVersion
	}
	if o.Sequence != nil {
		toSerialize["Sequence"] = o.Sequence
	}
	if o.Security != nil {
		toSerialize["Security"] = o.Security
	}
	if o.Block != nil {
		toSerialize["Block"] = o.Block
	}
	if o.RequiredInformation != nil {
		toSerialize["RequiredInformation"] = o.RequiredInformation
	}
	if o.Ticket != nil {
		toSerialize["Ticket"] = o.Ticket
	}
	if o.TicketAnswerKey != nil {
		toSerialize["TicketAnswerKey"] = o.TicketAnswerKey
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}
	if o.MerchantNotifyURL != nil {
		toSerialize["MerchantNotifyURL"] = o.MerchantNotifyURL
	}
	if o.IsReverse != nil {
		toSerialize["IsReverse"] = o.IsReverse
	}
	if o.ReasonReverse != nil {
		toSerialize["ReasonReverse"] = o.ReasonReverse
	}
	if o.ReasonSequenceBreak != nil {
		toSerialize["ReasonSequenceBreak"] = o.ReasonSequenceBreak
	}
	if o.Reference != nil {
		toSerialize["Reference"] = o.Reference
	}
	if o.TransactionDescription != nil {
		toSerialize["TransactionDescription"] = o.TransactionDescription
	}
	if o.POSType != nil {
		toSerialize["POSType"] = o.POSType
	}
	if o.POSVersion != nil {
		toSerialize["POSVersion"] = o.POSVersion
	}
	if o.POSAddress != nil {
		toSerialize["POSAddress"] = o.POSAddress
	}
	if o.POSSerial != nil {
		toSerialize["POSSerial"] = o.POSSerial
	}
	if o.POSGEO != nil {
		toSerialize["POSGEO"] = o.POSGEO
	}
	if o.ReadingDeviceType != nil {
		toSerialize["ReadingDeviceType"] = o.ReadingDeviceType
	}
	if o.ReadingDeviceOperatingFrom != nil {
		toSerialize["ReadingDeviceOperatingFrom"] = o.ReadingDeviceOperatingFrom
	}
	if o.ReadingDeviceVersion != nil {
		toSerialize["ReadingDeviceVersion"] = o.ReadingDeviceVersion
	}
	if o.ReadingDeviceAddress != nil {
		toSerialize["ReadingDeviceAddress"] = o.ReadingDeviceAddress
	}
	if o.ReadingDeviceSerial != nil {
		toSerialize["ReadingDeviceSerial"] = o.ReadingDeviceSerial
	}
	if o.ReadingDeviceGEO != nil {
		toSerialize["ReadingDeviceGEO"] = o.ReadingDeviceGEO
	}
	if o.UserID != nil {
		toSerialize["UserID"] = o.UserID
	}
	if o.UserName != nil {
		toSerialize["UserName"] = o.UserName
	}
	if true {
		toSerialize["Amount"] = o.Amount
	}
	if o.AlternativeAmount != nil {
		toSerialize["AlternativeAmount"] = o.AlternativeAmount
	}
	if o.CashbackAmount != nil {
		toSerialize["CashbackAmount"] = o.CashbackAmount
	}
	if o.TipAmount != nil {
		toSerialize["TipAmount"] = o.TipAmount
	}
	if o.PromotedAmount != nil {
		toSerialize["PromotedAmount"] = o.PromotedAmount
	}
	if o.CurrencyCode != nil {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if o.FacilityPayments != nil {
		toSerialize["FacilityPayments"] = o.FacilityPayments
	}
	if o.FacilityType != nil {
		toSerialize["FacilityType"] = o.FacilityType
	}
	if o.Plan != nil {
		toSerialize["Plan"] = o.Plan
	}
	if o.CardReadMode != nil {
		toSerialize["CardReadMode"] = o.CardReadMode
	}
	if o.CardGetMode != nil {
		toSerialize["CardGetMode"] = o.CardGetMode
	}
	if o.CardNumber != nil {
		toSerialize["CardNumber"] = o.CardNumber
	}
	if o.CardNumberMasked != nil {
		toSerialize["CardNumberMasked"] = o.CardNumberMasked
	}
	if o.CardNumberEncrypted != nil {
		toSerialize["CardNumberEncrypted"] = o.CardNumberEncrypted
	}
	if o.CardExp != nil {
		toSerialize["CardExp"] = o.CardExp
	}
	if o.CardCryptogram != nil {
		toSerialize["CardCryptogram"] = o.CardCryptogram
	}
	if o.CardAppName != nil {
		toSerialize["CardAppName"] = o.CardAppName
	}
	if o.CardAppIdentifier != nil {
		toSerialize["CardAppIdentifier"] = o.CardAppIdentifier
	}
	if o.CardAppLabel != nil {
		toSerialize["CardAppLabel"] = o.CardAppLabel
	}
	if o.CardAuthRequestCryptogram != nil {
		toSerialize["CardAuthRequestCryptogram"] = o.CardAuthRequestCryptogram
	}
	if o.CardAuthResponseCryptogram != nil {
		toSerialize["CardAuthResponseCryptogram"] = o.CardAuthResponseCryptogram
	}
	if o.Track1 != nil {
		toSerialize["Track1"] = o.Track1
	}
	if o.Track2 != nil {
		toSerialize["Track2"] = o.Track2
	}
	if o.InputTokens != nil {
		toSerialize["InputTokens"] = o.InputTokens
	}
	if o.SecurityCode != nil {
		toSerialize["SecurityCode"] = o.SecurityCode
	}
	if o.Pin != nil {
		toSerialize["Pin"] = o.Pin
	}
	if o.CredentialToken != nil {
		toSerialize["CredentialToken"] = o.CredentialToken
	}
	if o.CredentialIssuerToken != nil {
		toSerialize["CredentialIssuerToken"] = o.CredentialIssuerToken
	}
	if o.Payer != nil {
		toSerialize["Payer"] = o.Payer
	}
	if o.Customer != nil {
		toSerialize["Customer"] = o.Customer
	}
	if o.Seller != nil {
		toSerialize["Seller"] = o.Seller
	}
	if o.Products != nil {
		toSerialize["Products"] = o.Products
	}
	if o.TaxRefundType != nil {
		toSerialize["TaxRefundType"] = o.TaxRefundType
	}
	if o.AuthCode != nil {
		toSerialize["AuthCode"] = o.AuthCode
	}
	if o.PaymentFacilitatorID != nil {
		toSerialize["PaymentFacilitatorID"] = o.PaymentFacilitatorID
	}
	if o.MerchantID != nil {
		toSerialize["MerchantID"] = o.MerchantID
	}
	if o.TerminalID != nil {
		toSerialize["TerminalID"] = o.TerminalID
	}
	if o.TerminalTrace != nil {
		toSerialize["TerminalTrace"] = o.TerminalTrace
	}
	if o.SettlementBatchNumber != nil {
		toSerialize["SettlementBatchNumber"] = o.SettlementBatchNumber
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizeSaleObjectAuthorizeSale struct {
	value *AuthorizeSaleObjectAuthorizeSale
	isSet bool
}

func (v NullableAuthorizeSaleObjectAuthorizeSale) Get() *AuthorizeSaleObjectAuthorizeSale {
	return v.value
}

func (v *NullableAuthorizeSaleObjectAuthorizeSale) Set(val *AuthorizeSaleObjectAuthorizeSale) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeSaleObjectAuthorizeSale) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeSaleObjectAuthorizeSale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeSaleObjectAuthorizeSale(val *AuthorizeSaleObjectAuthorizeSale) *NullableAuthorizeSaleObjectAuthorizeSale {
	return &NullableAuthorizeSaleObjectAuthorizeSale{value: val, isSet: true}
}

func (v NullableAuthorizeSaleObjectAuthorizeSale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeSaleObjectAuthorizeSale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


