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

// GetTransactionResponseObjectGetTransactionResponseTransaction struct for GetTransactionResponseObjectGetTransactionResponseTransaction
type GetTransactionResponseObjectGetTransactionResponseTransaction struct {
	// Acciones a realizar por parte del POS y/o PINPAD en base al resultado de la operación que ha sido procesada. Cada uno de estos actions o acciones están concatenadas por comas. Los posibles actions son OK, OK, NotExist, ConfigurationError, SystemError, ResourceError, ProcessError, Configure.
	ResponseActions []string `json:"ResponseActions,omitempty"`
	// Descripción del resultado del proceso del requerimiento recibido. Esta descripción es generada por la plataforma, no por el Host que termine resolviendo la transacción.
	ResponseMessage *string `json:"ResponseMessage,omitempty"`
	// Código de respuesta para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF.
	ForeignResponseCode *string `json:"ForeignResponseCode,omitempty"`
	// Código de Respuesta Interno de la plataforma, el POS debe actuar por lo que indican las acciones especificadas en ResponseActions y no por el código de respuesta informado en este campo o elemento, pero es una buena práctica que sea persistido por el mismo.
	ResponseCode *int32 `json:"ResponseCode,omitempty"`
	// Monto del recargo impositivo aplicado al costo financiero que la transacción tiene
	TaxFinancialCostAmount *float32 `json:"TaxFinancialCostAmount,omitempty"`
	// Porcentaje de recargo impositivo a aplicar sobre el monto del costo financiero
	TaxFinancialCostPercentage *float32 `json:"TaxFinancialCostPercentage,omitempty"`
	// Monto del Costo financiero total generado en base al plan elegido
	FinancialCostAmount *float32 `json:"FinancialCostAmount,omitempty"`
	// Porcentaje del costo financiero a aplicar al monto de la transacción
	FinancialCostPercentage *float32 `json:"FinancialCostPercentage,omitempty"`
	// Monto libre de costos financerios e impuestos por el que la venta se realizó. El monto cobrado realmente no es este, dado que no incluye las tasas e impuestos
	RequestAmount *float32 `json:"RequestAmount,omitempty"`
	// Importe o Monto de la Transacción.
	Amount *float32 `json:"Amount,omitempty"`
	// Monto con la que se realizó transacción. Si este valor es recibido, la búsqueda de los planes será limitada con este criterio.
	AlternativeAmount *float32 `json:"AlternativeAmount,omitempty"`
	// Código de Resultado retornado por el Host Adquirente.
	HostResultCode *int32 `json:"HostResultCode,omitempty"`
	// Mensaje Retornado por el Host Adquirente, normalmente asociado al valor de HostResultCode
	HostMessage *string `json:"HostMessage,omitempty"`
	// código de autorización retornado por el Host que resuelve la transacción.
	HostCode *string `json:"HostCode,omitempty"`
	// Fecha y Hora de la transacción retornada por el Host que resuelve la Transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14
	HostDateTime *time.Time `json:"HostDateTime,omitempty"`
	// Fecha y hora de transmision de la operación hacia el host - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14
	TransmitionDateTime *time.Time `json:"TransmitionDateTime,omitempty"`
	// Número Ticket  o Voucher Generado para la Plataforma.
	AuthTicket *int32 `json:"AuthTicket,omitempty"`
	// Número de identificación de la transacción, utilizado por la mayoría de los hosts para realizar anulaciones y devoluciones
	AuthRRN *string `json:"AuthRRN,omitempty"`
	// Identificador que genera el Host Adquirente para la Transacción en algunos podrá ser igual al AuthRRN
	IdentifierForTheAdquirer *string `json:"IdentifierForTheAdquirer,omitempty"`
	// Identificador de facilitador de pagos o Payfac.
	PaymentFacilitatorID *string `json:"PaymentFacilitatorID,omitempty"`
	// Número de comercio utilizado para realizar la transacción. Este Número es asignado por el host, y parametrizado en la BD, relaciónado a cada uno de los planes disponibles.
	MerchantID *string `json:"MerchantID,omitempty"`
	// Identificador de Terminal por el cual se envía la Transacción al Host.
	TerminalID *int32 `json:"TerminalID,omitempty"`
	// Número de Trace/Secuencia que genera la plataforma para la transacción asociado al TerminalID.
	TerminalTrace *int32 `json:"TerminalTrace,omitempty"`
	// Para aquellos host que exista el concepto de Lote, es el número de lote al cual pertenece la transacción.
	SettlementBatchNumber *int32 `json:"SettlementBatchNumber,omitempty"`
	// Número de identificación del host al cual fue enviada la petición, y por el cual fue finalmente procesada
	HostID *int32 `json:"HostID,omitempty"`
	// Nombre del Emisor de la Credencial o Tarjeta que se usó en la transacción.
	IssuerName *string `json:"IssuerName,omitempty"`
	// Nombre de la Tarjeta que se usó en la transacción, usado para la impresión del voucher.
	CardDescription *string `json:"CardDescription,omitempty"`
	// Descripción o nombre de la marca con la cual la tarjeta fue identificada
	PaymentMethodDescription *string `json:"PaymentMethodDescription,omitempty"`
	// Descripción del plan utilizado para para realizar la operación
	PlanDescription *string `json:"PlanDescription,omitempty"`
	// Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes )
	CardReadMode *string `json:"CardReadMode,omitempty"`
	// Número de Tarjeta, en el caso de las respuestas el mismo estará enmascarado.
	CardNumber *string `json:"CardNumber,omitempty"`
	// Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón
	CardNumberMasked *string `json:"CardNumberMasked,omitempty"`
	// Hash de la tarjeta generado por la plataforma.
	CardHashing *string `json:"CardHashing,omitempty"`
	// Descripción del tipo de cambio utilizado en la transacción
	CurrencyDescription *string `json:"CurrencyDescription,omitempty"`
	// Simbolo monetario del tipo de cambio utilizado en la transacción
	CurrencySymbol *string `json:"CurrencySymbol,omitempty"`
	// Tags EMV en format TLV recibidos desde el Host.
	CardCryptogramResponse *string `json:"CardCryptogramResponse,omitempty"`
	// Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers.
	CardAppName *string `json:"CardAppName,omitempty"`
	// Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers.
	CardAppIdentifier *string `json:"CardAppIdentifier,omitempty"`
	// Disponible solo con Tarjetas Chip (Incluye Contacless Chip), se debe imprimir en los Tickets/Vouchers.
	CardAppLabel *string `json:"CardAppLabel,omitempty"`
}

// NewGetTransactionResponseObjectGetTransactionResponseTransaction instantiates a new GetTransactionResponseObjectGetTransactionResponseTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionResponseObjectGetTransactionResponseTransaction() *GetTransactionResponseObjectGetTransactionResponseTransaction {
	this := GetTransactionResponseObjectGetTransactionResponseTransaction{}
	return &this
}

// NewGetTransactionResponseObjectGetTransactionResponseTransactionWithDefaults instantiates a new GetTransactionResponseObjectGetTransactionResponseTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionResponseObjectGetTransactionResponseTransactionWithDefaults() *GetTransactionResponseObjectGetTransactionResponseTransaction {
	this := GetTransactionResponseObjectGetTransactionResponseTransaction{}
	return &this
}

// GetResponseActions returns the ResponseActions field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseActions() []string {
	if o == nil || o.ResponseActions == nil {
		var ret []string
		return ret
	}
	return o.ResponseActions
}

// GetResponseActionsOk returns a tuple with the ResponseActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseActionsOk() ([]string, bool) {
	if o == nil || o.ResponseActions == nil {
		return nil, false
	}
	return o.ResponseActions, true
}

// HasResponseActions returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasResponseActions() bool {
	if o != nil && o.ResponseActions != nil {
		return true
	}

	return false
}

// SetResponseActions gets a reference to the given []string and assigns it to the ResponseActions field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetResponseActions(v []string) {
	o.ResponseActions = v
}

// GetResponseMessage returns the ResponseMessage field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseMessage() string {
	if o == nil || o.ResponseMessage == nil {
		var ret string
		return ret
	}
	return *o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseMessageOk() (*string, bool) {
	if o == nil || o.ResponseMessage == nil {
		return nil, false
	}
	return o.ResponseMessage, true
}

// HasResponseMessage returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasResponseMessage() bool {
	if o != nil && o.ResponseMessage != nil {
		return true
	}

	return false
}

// SetResponseMessage gets a reference to the given string and assigns it to the ResponseMessage field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetResponseMessage(v string) {
	o.ResponseMessage = &v
}

// GetForeignResponseCode returns the ForeignResponseCode field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetForeignResponseCode() string {
	if o == nil || o.ForeignResponseCode == nil {
		var ret string
		return ret
	}
	return *o.ForeignResponseCode
}

// GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetForeignResponseCodeOk() (*string, bool) {
	if o == nil || o.ForeignResponseCode == nil {
		return nil, false
	}
	return o.ForeignResponseCode, true
}

// HasForeignResponseCode returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasForeignResponseCode() bool {
	if o != nil && o.ForeignResponseCode != nil {
		return true
	}

	return false
}

// SetForeignResponseCode gets a reference to the given string and assigns it to the ForeignResponseCode field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetForeignResponseCode(v string) {
	o.ForeignResponseCode = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseCode() int32 {
	if o == nil || o.ResponseCode == nil {
		var ret int32
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetResponseCodeOk() (*int32, bool) {
	if o == nil || o.ResponseCode == nil {
		return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasResponseCode() bool {
	if o != nil && o.ResponseCode != nil {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given int32 and assigns it to the ResponseCode field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetResponseCode(v int32) {
	o.ResponseCode = &v
}

// GetTaxFinancialCostAmount returns the TaxFinancialCostAmount field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTaxFinancialCostAmount() float32 {
	if o == nil || o.TaxFinancialCostAmount == nil {
		var ret float32
		return ret
	}
	return *o.TaxFinancialCostAmount
}

// GetTaxFinancialCostAmountOk returns a tuple with the TaxFinancialCostAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTaxFinancialCostAmountOk() (*float32, bool) {
	if o == nil || o.TaxFinancialCostAmount == nil {
		return nil, false
	}
	return o.TaxFinancialCostAmount, true
}

// HasTaxFinancialCostAmount returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasTaxFinancialCostAmount() bool {
	if o != nil && o.TaxFinancialCostAmount != nil {
		return true
	}

	return false
}

// SetTaxFinancialCostAmount gets a reference to the given float32 and assigns it to the TaxFinancialCostAmount field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetTaxFinancialCostAmount(v float32) {
	o.TaxFinancialCostAmount = &v
}

// GetTaxFinancialCostPercentage returns the TaxFinancialCostPercentage field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTaxFinancialCostPercentage() float32 {
	if o == nil || o.TaxFinancialCostPercentage == nil {
		var ret float32
		return ret
	}
	return *o.TaxFinancialCostPercentage
}

// GetTaxFinancialCostPercentageOk returns a tuple with the TaxFinancialCostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTaxFinancialCostPercentageOk() (*float32, bool) {
	if o == nil || o.TaxFinancialCostPercentage == nil {
		return nil, false
	}
	return o.TaxFinancialCostPercentage, true
}

// HasTaxFinancialCostPercentage returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasTaxFinancialCostPercentage() bool {
	if o != nil && o.TaxFinancialCostPercentage != nil {
		return true
	}

	return false
}

// SetTaxFinancialCostPercentage gets a reference to the given float32 and assigns it to the TaxFinancialCostPercentage field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetTaxFinancialCostPercentage(v float32) {
	o.TaxFinancialCostPercentage = &v
}

// GetFinancialCostAmount returns the FinancialCostAmount field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetFinancialCostAmount() float32 {
	if o == nil || o.FinancialCostAmount == nil {
		var ret float32
		return ret
	}
	return *o.FinancialCostAmount
}

// GetFinancialCostAmountOk returns a tuple with the FinancialCostAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetFinancialCostAmountOk() (*float32, bool) {
	if o == nil || o.FinancialCostAmount == nil {
		return nil, false
	}
	return o.FinancialCostAmount, true
}

// HasFinancialCostAmount returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasFinancialCostAmount() bool {
	if o != nil && o.FinancialCostAmount != nil {
		return true
	}

	return false
}

// SetFinancialCostAmount gets a reference to the given float32 and assigns it to the FinancialCostAmount field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetFinancialCostAmount(v float32) {
	o.FinancialCostAmount = &v
}

// GetFinancialCostPercentage returns the FinancialCostPercentage field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetFinancialCostPercentage() float32 {
	if o == nil || o.FinancialCostPercentage == nil {
		var ret float32
		return ret
	}
	return *o.FinancialCostPercentage
}

// GetFinancialCostPercentageOk returns a tuple with the FinancialCostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetFinancialCostPercentageOk() (*float32, bool) {
	if o == nil || o.FinancialCostPercentage == nil {
		return nil, false
	}
	return o.FinancialCostPercentage, true
}

// HasFinancialCostPercentage returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasFinancialCostPercentage() bool {
	if o != nil && o.FinancialCostPercentage != nil {
		return true
	}

	return false
}

// SetFinancialCostPercentage gets a reference to the given float32 and assigns it to the FinancialCostPercentage field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetFinancialCostPercentage(v float32) {
	o.FinancialCostPercentage = &v
}

// GetRequestAmount returns the RequestAmount field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetRequestAmount() float32 {
	if o == nil || o.RequestAmount == nil {
		var ret float32
		return ret
	}
	return *o.RequestAmount
}

// GetRequestAmountOk returns a tuple with the RequestAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetRequestAmountOk() (*float32, bool) {
	if o == nil || o.RequestAmount == nil {
		return nil, false
	}
	return o.RequestAmount, true
}

// HasRequestAmount returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasRequestAmount() bool {
	if o != nil && o.RequestAmount != nil {
		return true
	}

	return false
}

// SetRequestAmount gets a reference to the given float32 and assigns it to the RequestAmount field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetRequestAmount(v float32) {
	o.RequestAmount = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetAmount(v float32) {
	o.Amount = &v
}

// GetAlternativeAmount returns the AlternativeAmount field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAlternativeAmount() float32 {
	if o == nil || o.AlternativeAmount == nil {
		var ret float32
		return ret
	}
	return *o.AlternativeAmount
}

// GetAlternativeAmountOk returns a tuple with the AlternativeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAlternativeAmountOk() (*float32, bool) {
	if o == nil || o.AlternativeAmount == nil {
		return nil, false
	}
	return o.AlternativeAmount, true
}

// HasAlternativeAmount returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasAlternativeAmount() bool {
	if o != nil && o.AlternativeAmount != nil {
		return true
	}

	return false
}

// SetAlternativeAmount gets a reference to the given float32 and assigns it to the AlternativeAmount field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetAlternativeAmount(v float32) {
	o.AlternativeAmount = &v
}

// GetHostResultCode returns the HostResultCode field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostResultCode() int32 {
	if o == nil || o.HostResultCode == nil {
		var ret int32
		return ret
	}
	return *o.HostResultCode
}

// GetHostResultCodeOk returns a tuple with the HostResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostResultCodeOk() (*int32, bool) {
	if o == nil || o.HostResultCode == nil {
		return nil, false
	}
	return o.HostResultCode, true
}

// HasHostResultCode returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasHostResultCode() bool {
	if o != nil && o.HostResultCode != nil {
		return true
	}

	return false
}

// SetHostResultCode gets a reference to the given int32 and assigns it to the HostResultCode field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetHostResultCode(v int32) {
	o.HostResultCode = &v
}

// GetHostMessage returns the HostMessage field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostMessage() string {
	if o == nil || o.HostMessage == nil {
		var ret string
		return ret
	}
	return *o.HostMessage
}

// GetHostMessageOk returns a tuple with the HostMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostMessageOk() (*string, bool) {
	if o == nil || o.HostMessage == nil {
		return nil, false
	}
	return o.HostMessage, true
}

// HasHostMessage returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasHostMessage() bool {
	if o != nil && o.HostMessage != nil {
		return true
	}

	return false
}

// SetHostMessage gets a reference to the given string and assigns it to the HostMessage field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetHostMessage(v string) {
	o.HostMessage = &v
}

// GetHostCode returns the HostCode field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostCode() string {
	if o == nil || o.HostCode == nil {
		var ret string
		return ret
	}
	return *o.HostCode
}

// GetHostCodeOk returns a tuple with the HostCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostCodeOk() (*string, bool) {
	if o == nil || o.HostCode == nil {
		return nil, false
	}
	return o.HostCode, true
}

// HasHostCode returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasHostCode() bool {
	if o != nil && o.HostCode != nil {
		return true
	}

	return false
}

// SetHostCode gets a reference to the given string and assigns it to the HostCode field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetHostCode(v string) {
	o.HostCode = &v
}

// GetHostDateTime returns the HostDateTime field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostDateTime() time.Time {
	if o == nil || o.HostDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.HostDateTime
}

// GetHostDateTimeOk returns a tuple with the HostDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostDateTimeOk() (*time.Time, bool) {
	if o == nil || o.HostDateTime == nil {
		return nil, false
	}
	return o.HostDateTime, true
}

// HasHostDateTime returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasHostDateTime() bool {
	if o != nil && o.HostDateTime != nil {
		return true
	}

	return false
}

// SetHostDateTime gets a reference to the given time.Time and assigns it to the HostDateTime field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetHostDateTime(v time.Time) {
	o.HostDateTime = &v
}

// GetTransmitionDateTime returns the TransmitionDateTime field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTransmitionDateTime() time.Time {
	if o == nil || o.TransmitionDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.TransmitionDateTime
}

// GetTransmitionDateTimeOk returns a tuple with the TransmitionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTransmitionDateTimeOk() (*time.Time, bool) {
	if o == nil || o.TransmitionDateTime == nil {
		return nil, false
	}
	return o.TransmitionDateTime, true
}

// HasTransmitionDateTime returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasTransmitionDateTime() bool {
	if o != nil && o.TransmitionDateTime != nil {
		return true
	}

	return false
}

// SetTransmitionDateTime gets a reference to the given time.Time and assigns it to the TransmitionDateTime field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetTransmitionDateTime(v time.Time) {
	o.TransmitionDateTime = &v
}

// GetAuthTicket returns the AuthTicket field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAuthTicket() int32 {
	if o == nil || o.AuthTicket == nil {
		var ret int32
		return ret
	}
	return *o.AuthTicket
}

// GetAuthTicketOk returns a tuple with the AuthTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAuthTicketOk() (*int32, bool) {
	if o == nil || o.AuthTicket == nil {
		return nil, false
	}
	return o.AuthTicket, true
}

// HasAuthTicket returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasAuthTicket() bool {
	if o != nil && o.AuthTicket != nil {
		return true
	}

	return false
}

// SetAuthTicket gets a reference to the given int32 and assigns it to the AuthTicket field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetAuthTicket(v int32) {
	o.AuthTicket = &v
}

// GetAuthRRN returns the AuthRRN field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAuthRRN() string {
	if o == nil || o.AuthRRN == nil {
		var ret string
		return ret
	}
	return *o.AuthRRN
}

// GetAuthRRNOk returns a tuple with the AuthRRN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetAuthRRNOk() (*string, bool) {
	if o == nil || o.AuthRRN == nil {
		return nil, false
	}
	return o.AuthRRN, true
}

// HasAuthRRN returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasAuthRRN() bool {
	if o != nil && o.AuthRRN != nil {
		return true
	}

	return false
}

// SetAuthRRN gets a reference to the given string and assigns it to the AuthRRN field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetAuthRRN(v string) {
	o.AuthRRN = &v
}

// GetIdentifierForTheAdquirer returns the IdentifierForTheAdquirer field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetIdentifierForTheAdquirer() string {
	if o == nil || o.IdentifierForTheAdquirer == nil {
		var ret string
		return ret
	}
	return *o.IdentifierForTheAdquirer
}

// GetIdentifierForTheAdquirerOk returns a tuple with the IdentifierForTheAdquirer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetIdentifierForTheAdquirerOk() (*string, bool) {
	if o == nil || o.IdentifierForTheAdquirer == nil {
		return nil, false
	}
	return o.IdentifierForTheAdquirer, true
}

// HasIdentifierForTheAdquirer returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasIdentifierForTheAdquirer() bool {
	if o != nil && o.IdentifierForTheAdquirer != nil {
		return true
	}

	return false
}

// SetIdentifierForTheAdquirer gets a reference to the given string and assigns it to the IdentifierForTheAdquirer field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetIdentifierForTheAdquirer(v string) {
	o.IdentifierForTheAdquirer = &v
}

// GetPaymentFacilitatorID returns the PaymentFacilitatorID field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPaymentFacilitatorID() string {
	if o == nil || o.PaymentFacilitatorID == nil {
		var ret string
		return ret
	}
	return *o.PaymentFacilitatorID
}

// GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPaymentFacilitatorIDOk() (*string, bool) {
	if o == nil || o.PaymentFacilitatorID == nil {
		return nil, false
	}
	return o.PaymentFacilitatorID, true
}

// HasPaymentFacilitatorID returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasPaymentFacilitatorID() bool {
	if o != nil && o.PaymentFacilitatorID != nil {
		return true
	}

	return false
}

// SetPaymentFacilitatorID gets a reference to the given string and assigns it to the PaymentFacilitatorID field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetPaymentFacilitatorID(v string) {
	o.PaymentFacilitatorID = &v
}

// GetMerchantID returns the MerchantID field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetMerchantID() string {
	if o == nil || o.MerchantID == nil {
		var ret string
		return ret
	}
	return *o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetMerchantIDOk() (*string, bool) {
	if o == nil || o.MerchantID == nil {
		return nil, false
	}
	return o.MerchantID, true
}

// HasMerchantID returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasMerchantID() bool {
	if o != nil && o.MerchantID != nil {
		return true
	}

	return false
}

// SetMerchantID gets a reference to the given string and assigns it to the MerchantID field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetMerchantID(v string) {
	o.MerchantID = &v
}

// GetTerminalID returns the TerminalID field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTerminalID() int32 {
	if o == nil || o.TerminalID == nil {
		var ret int32
		return ret
	}
	return *o.TerminalID
}

// GetTerminalIDOk returns a tuple with the TerminalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTerminalIDOk() (*int32, bool) {
	if o == nil || o.TerminalID == nil {
		return nil, false
	}
	return o.TerminalID, true
}

// HasTerminalID returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasTerminalID() bool {
	if o != nil && o.TerminalID != nil {
		return true
	}

	return false
}

// SetTerminalID gets a reference to the given int32 and assigns it to the TerminalID field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetTerminalID(v int32) {
	o.TerminalID = &v
}

// GetTerminalTrace returns the TerminalTrace field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTerminalTrace() int32 {
	if o == nil || o.TerminalTrace == nil {
		var ret int32
		return ret
	}
	return *o.TerminalTrace
}

// GetTerminalTraceOk returns a tuple with the TerminalTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetTerminalTraceOk() (*int32, bool) {
	if o == nil || o.TerminalTrace == nil {
		return nil, false
	}
	return o.TerminalTrace, true
}

// HasTerminalTrace returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasTerminalTrace() bool {
	if o != nil && o.TerminalTrace != nil {
		return true
	}

	return false
}

// SetTerminalTrace gets a reference to the given int32 and assigns it to the TerminalTrace field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetTerminalTrace(v int32) {
	o.TerminalTrace = &v
}

// GetSettlementBatchNumber returns the SettlementBatchNumber field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetSettlementBatchNumber() int32 {
	if o == nil || o.SettlementBatchNumber == nil {
		var ret int32
		return ret
	}
	return *o.SettlementBatchNumber
}

// GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetSettlementBatchNumberOk() (*int32, bool) {
	if o == nil || o.SettlementBatchNumber == nil {
		return nil, false
	}
	return o.SettlementBatchNumber, true
}

// HasSettlementBatchNumber returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasSettlementBatchNumber() bool {
	if o != nil && o.SettlementBatchNumber != nil {
		return true
	}

	return false
}

// SetSettlementBatchNumber gets a reference to the given int32 and assigns it to the SettlementBatchNumber field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetSettlementBatchNumber(v int32) {
	o.SettlementBatchNumber = &v
}

// GetHostID returns the HostID field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostID() int32 {
	if o == nil || o.HostID == nil {
		var ret int32
		return ret
	}
	return *o.HostID
}

// GetHostIDOk returns a tuple with the HostID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetHostIDOk() (*int32, bool) {
	if o == nil || o.HostID == nil {
		return nil, false
	}
	return o.HostID, true
}

// HasHostID returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasHostID() bool {
	if o != nil && o.HostID != nil {
		return true
	}

	return false
}

// SetHostID gets a reference to the given int32 and assigns it to the HostID field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetHostID(v int32) {
	o.HostID = &v
}

// GetIssuerName returns the IssuerName field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetIssuerName() string {
	if o == nil || o.IssuerName == nil {
		var ret string
		return ret
	}
	return *o.IssuerName
}

// GetIssuerNameOk returns a tuple with the IssuerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetIssuerNameOk() (*string, bool) {
	if o == nil || o.IssuerName == nil {
		return nil, false
	}
	return o.IssuerName, true
}

// HasIssuerName returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasIssuerName() bool {
	if o != nil && o.IssuerName != nil {
		return true
	}

	return false
}

// SetIssuerName gets a reference to the given string and assigns it to the IssuerName field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetIssuerName(v string) {
	o.IssuerName = &v
}

// GetCardDescription returns the CardDescription field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardDescription() string {
	if o == nil || o.CardDescription == nil {
		var ret string
		return ret
	}
	return *o.CardDescription
}

// GetCardDescriptionOk returns a tuple with the CardDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardDescriptionOk() (*string, bool) {
	if o == nil || o.CardDescription == nil {
		return nil, false
	}
	return o.CardDescription, true
}

// HasCardDescription returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardDescription() bool {
	if o != nil && o.CardDescription != nil {
		return true
	}

	return false
}

// SetCardDescription gets a reference to the given string and assigns it to the CardDescription field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardDescription(v string) {
	o.CardDescription = &v
}

// GetPaymentMethodDescription returns the PaymentMethodDescription field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPaymentMethodDescription() string {
	if o == nil || o.PaymentMethodDescription == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodDescription
}

// GetPaymentMethodDescriptionOk returns a tuple with the PaymentMethodDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPaymentMethodDescriptionOk() (*string, bool) {
	if o == nil || o.PaymentMethodDescription == nil {
		return nil, false
	}
	return o.PaymentMethodDescription, true
}

// HasPaymentMethodDescription returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasPaymentMethodDescription() bool {
	if o != nil && o.PaymentMethodDescription != nil {
		return true
	}

	return false
}

// SetPaymentMethodDescription gets a reference to the given string and assigns it to the PaymentMethodDescription field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetPaymentMethodDescription(v string) {
	o.PaymentMethodDescription = &v
}

// GetPlanDescription returns the PlanDescription field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPlanDescription() string {
	if o == nil || o.PlanDescription == nil {
		var ret string
		return ret
	}
	return *o.PlanDescription
}

// GetPlanDescriptionOk returns a tuple with the PlanDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetPlanDescriptionOk() (*string, bool) {
	if o == nil || o.PlanDescription == nil {
		return nil, false
	}
	return o.PlanDescription, true
}

// HasPlanDescription returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasPlanDescription() bool {
	if o != nil && o.PlanDescription != nil {
		return true
	}

	return false
}

// SetPlanDescription gets a reference to the given string and assigns it to the PlanDescription field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetPlanDescription(v string) {
	o.PlanDescription = &v
}

// GetCardReadMode returns the CardReadMode field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardReadMode() string {
	if o == nil || o.CardReadMode == nil {
		var ret string
		return ret
	}
	return *o.CardReadMode
}

// GetCardReadModeOk returns a tuple with the CardReadMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardReadModeOk() (*string, bool) {
	if o == nil || o.CardReadMode == nil {
		return nil, false
	}
	return o.CardReadMode, true
}

// HasCardReadMode returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardReadMode() bool {
	if o != nil && o.CardReadMode != nil {
		return true
	}

	return false
}

// SetCardReadMode gets a reference to the given string and assigns it to the CardReadMode field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardReadMode(v string) {
	o.CardReadMode = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardNumber() string {
	if o == nil || o.CardNumber == nil {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardNumberOk() (*string, bool) {
	if o == nil || o.CardNumber == nil {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardNumber() bool {
	if o != nil && o.CardNumber != nil {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetCardNumberMasked returns the CardNumberMasked field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardNumberMasked() string {
	if o == nil || o.CardNumberMasked == nil {
		var ret string
		return ret
	}
	return *o.CardNumberMasked
}

// GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardNumberMaskedOk() (*string, bool) {
	if o == nil || o.CardNumberMasked == nil {
		return nil, false
	}
	return o.CardNumberMasked, true
}

// HasCardNumberMasked returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardNumberMasked() bool {
	if o != nil && o.CardNumberMasked != nil {
		return true
	}

	return false
}

// SetCardNumberMasked gets a reference to the given string and assigns it to the CardNumberMasked field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardNumberMasked(v string) {
	o.CardNumberMasked = &v
}

// GetCardHashing returns the CardHashing field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardHashing() string {
	if o == nil || o.CardHashing == nil {
		var ret string
		return ret
	}
	return *o.CardHashing
}

// GetCardHashingOk returns a tuple with the CardHashing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardHashingOk() (*string, bool) {
	if o == nil || o.CardHashing == nil {
		return nil, false
	}
	return o.CardHashing, true
}

// HasCardHashing returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardHashing() bool {
	if o != nil && o.CardHashing != nil {
		return true
	}

	return false
}

// SetCardHashing gets a reference to the given string and assigns it to the CardHashing field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardHashing(v string) {
	o.CardHashing = &v
}

// GetCurrencyDescription returns the CurrencyDescription field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCurrencyDescription() string {
	if o == nil || o.CurrencyDescription == nil {
		var ret string
		return ret
	}
	return *o.CurrencyDescription
}

// GetCurrencyDescriptionOk returns a tuple with the CurrencyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCurrencyDescriptionOk() (*string, bool) {
	if o == nil || o.CurrencyDescription == nil {
		return nil, false
	}
	return o.CurrencyDescription, true
}

// HasCurrencyDescription returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCurrencyDescription() bool {
	if o != nil && o.CurrencyDescription != nil {
		return true
	}

	return false
}

// SetCurrencyDescription gets a reference to the given string and assigns it to the CurrencyDescription field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCurrencyDescription(v string) {
	o.CurrencyDescription = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCardCryptogramResponse returns the CardCryptogramResponse field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardCryptogramResponse() string {
	if o == nil || o.CardCryptogramResponse == nil {
		var ret string
		return ret
	}
	return *o.CardCryptogramResponse
}

// GetCardCryptogramResponseOk returns a tuple with the CardCryptogramResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardCryptogramResponseOk() (*string, bool) {
	if o == nil || o.CardCryptogramResponse == nil {
		return nil, false
	}
	return o.CardCryptogramResponse, true
}

// HasCardCryptogramResponse returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardCryptogramResponse() bool {
	if o != nil && o.CardCryptogramResponse != nil {
		return true
	}

	return false
}

// SetCardCryptogramResponse gets a reference to the given string and assigns it to the CardCryptogramResponse field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardCryptogramResponse(v string) {
	o.CardCryptogramResponse = &v
}

// GetCardAppName returns the CardAppName field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppName() string {
	if o == nil || o.CardAppName == nil {
		var ret string
		return ret
	}
	return *o.CardAppName
}

// GetCardAppNameOk returns a tuple with the CardAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppNameOk() (*string, bool) {
	if o == nil || o.CardAppName == nil {
		return nil, false
	}
	return o.CardAppName, true
}

// HasCardAppName returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardAppName() bool {
	if o != nil && o.CardAppName != nil {
		return true
	}

	return false
}

// SetCardAppName gets a reference to the given string and assigns it to the CardAppName field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardAppName(v string) {
	o.CardAppName = &v
}

// GetCardAppIdentifier returns the CardAppIdentifier field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppIdentifier() string {
	if o == nil || o.CardAppIdentifier == nil {
		var ret string
		return ret
	}
	return *o.CardAppIdentifier
}

// GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppIdentifierOk() (*string, bool) {
	if o == nil || o.CardAppIdentifier == nil {
		return nil, false
	}
	return o.CardAppIdentifier, true
}

// HasCardAppIdentifier returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardAppIdentifier() bool {
	if o != nil && o.CardAppIdentifier != nil {
		return true
	}

	return false
}

// SetCardAppIdentifier gets a reference to the given string and assigns it to the CardAppIdentifier field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardAppIdentifier(v string) {
	o.CardAppIdentifier = &v
}

// GetCardAppLabel returns the CardAppLabel field value if set, zero value otherwise.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppLabel() string {
	if o == nil || o.CardAppLabel == nil {
		var ret string
		return ret
	}
	return *o.CardAppLabel
}

// GetCardAppLabelOk returns a tuple with the CardAppLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) GetCardAppLabelOk() (*string, bool) {
	if o == nil || o.CardAppLabel == nil {
		return nil, false
	}
	return o.CardAppLabel, true
}

// HasCardAppLabel returns a boolean if a field has been set.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) HasCardAppLabel() bool {
	if o != nil && o.CardAppLabel != nil {
		return true
	}

	return false
}

// SetCardAppLabel gets a reference to the given string and assigns it to the CardAppLabel field.
func (o *GetTransactionResponseObjectGetTransactionResponseTransaction) SetCardAppLabel(v string) {
	o.CardAppLabel = &v
}

func (o GetTransactionResponseObjectGetTransactionResponseTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseActions != nil {
		toSerialize["ResponseActions"] = o.ResponseActions
	}
	if o.ResponseMessage != nil {
		toSerialize["ResponseMessage"] = o.ResponseMessage
	}
	if o.ForeignResponseCode != nil {
		toSerialize["ForeignResponseCode"] = o.ForeignResponseCode
	}
	if o.ResponseCode != nil {
		toSerialize["ResponseCode"] = o.ResponseCode
	}
	if o.TaxFinancialCostAmount != nil {
		toSerialize["TaxFinancialCostAmount"] = o.TaxFinancialCostAmount
	}
	if o.TaxFinancialCostPercentage != nil {
		toSerialize["TaxFinancialCostPercentage"] = o.TaxFinancialCostPercentage
	}
	if o.FinancialCostAmount != nil {
		toSerialize["FinancialCostAmount"] = o.FinancialCostAmount
	}
	if o.FinancialCostPercentage != nil {
		toSerialize["FinancialCostPercentage"] = o.FinancialCostPercentage
	}
	if o.RequestAmount != nil {
		toSerialize["RequestAmount"] = o.RequestAmount
	}
	if o.Amount != nil {
		toSerialize["Amount"] = o.Amount
	}
	if o.AlternativeAmount != nil {
		toSerialize["AlternativeAmount"] = o.AlternativeAmount
	}
	if o.HostResultCode != nil {
		toSerialize["HostResultCode"] = o.HostResultCode
	}
	if o.HostMessage != nil {
		toSerialize["HostMessage"] = o.HostMessage
	}
	if o.HostCode != nil {
		toSerialize["HostCode"] = o.HostCode
	}
	if o.HostDateTime != nil {
		toSerialize["HostDateTime"] = o.HostDateTime
	}
	if o.TransmitionDateTime != nil {
		toSerialize["TransmitionDateTime"] = o.TransmitionDateTime
	}
	if o.AuthTicket != nil {
		toSerialize["AuthTicket"] = o.AuthTicket
	}
	if o.AuthRRN != nil {
		toSerialize["AuthRRN"] = o.AuthRRN
	}
	if o.IdentifierForTheAdquirer != nil {
		toSerialize["IdentifierForTheAdquirer"] = o.IdentifierForTheAdquirer
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
	if o.HostID != nil {
		toSerialize["HostID"] = o.HostID
	}
	if o.IssuerName != nil {
		toSerialize["IssuerName"] = o.IssuerName
	}
	if o.CardDescription != nil {
		toSerialize["CardDescription"] = o.CardDescription
	}
	if o.PaymentMethodDescription != nil {
		toSerialize["PaymentMethodDescription"] = o.PaymentMethodDescription
	}
	if o.PlanDescription != nil {
		toSerialize["PlanDescription"] = o.PlanDescription
	}
	if o.CardReadMode != nil {
		toSerialize["CardReadMode"] = o.CardReadMode
	}
	if o.CardNumber != nil {
		toSerialize["CardNumber"] = o.CardNumber
	}
	if o.CardNumberMasked != nil {
		toSerialize["CardNumberMasked"] = o.CardNumberMasked
	}
	if o.CardHashing != nil {
		toSerialize["CardHashing"] = o.CardHashing
	}
	if o.CurrencyDescription != nil {
		toSerialize["CurrencyDescription"] = o.CurrencyDescription
	}
	if o.CurrencySymbol != nil {
		toSerialize["CurrencySymbol"] = o.CurrencySymbol
	}
	if o.CardCryptogramResponse != nil {
		toSerialize["CardCryptogramResponse"] = o.CardCryptogramResponse
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
	return json.Marshal(toSerialize)
}

type NullableGetTransactionResponseObjectGetTransactionResponseTransaction struct {
	value *GetTransactionResponseObjectGetTransactionResponseTransaction
	isSet bool
}

func (v NullableGetTransactionResponseObjectGetTransactionResponseTransaction) Get() *GetTransactionResponseObjectGetTransactionResponseTransaction {
	return v.value
}

func (v *NullableGetTransactionResponseObjectGetTransactionResponseTransaction) Set(val *GetTransactionResponseObjectGetTransactionResponseTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionResponseObjectGetTransactionResponseTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionResponseObjectGetTransactionResponseTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionResponseObjectGetTransactionResponseTransaction(val *GetTransactionResponseObjectGetTransactionResponseTransaction) *NullableGetTransactionResponseObjectGetTransactionResponseTransaction {
	return &NullableGetTransactionResponseObjectGetTransactionResponseTransaction{value: val, isSet: true}
}

func (v NullableGetTransactionResponseObjectGetTransactionResponseTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionResponseObjectGetTransactionResponseTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


