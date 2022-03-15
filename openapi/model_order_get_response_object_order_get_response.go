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

// OrderGetResponseObjectOrderGetResponse struct for OrderGetResponseObjectOrderGetResponse
type OrderGetResponseObjectOrderGetResponse struct {
	// Tipo de Operación que se requirió, contendrá el mismo valor que se recibió en el requerimiento, sobre formatos que no soportan elementos complejos o compuestos.
	RequestType *string `json:"RequestType,omitempty"`
	// Acciones a realizar por parte del POS y/o PINPAD en base al resultado de la operación que ha sido procesada. Cada uno de estos actions o acciones están concatenadas por comas. Los posibles actions son OK, Approve, Refuse, IssuerCall, Tickets, WithHold, GetCard, UseTerminalToAuthorize, ConfigurationError, SystemError, ResourceError, ProcessError, Completed.
	ResponseActions []string `json:"ResponseActions"`
	// Descripción del resultado del proceso del requerimiento recibido. Esta descripción es generada por la plataforma, no por el Host que termine resolviendo la transacción.
	ResponseMessage string `json:"ResponseMessage"`
	// ID que identifica la companía desde donde proviene la petición.
	CompanyIdentification *string `json:"CompanyIdentification,omitempty"`
	// ID que identifica el sistema desde donde proviene la petición.
	SystemIdentification *string `json:"SystemIdentification,omitempty"`
	// ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía.
	BranchIdentification *string `json:"BranchIdentification,omitempty"`
	// ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía.
	POSIdentification *string `json:"POSIdentification,omitempty"`
	// Código de respuesta para el sistema externo, es decir, para la aplicación cliente que se comunica con el TEF.
	ForeignResponseCode *string `json:"ForeignResponseCode,omitempty"`
	// Código de Respuesta Interno de la plataforma, el POS debe actuar por lo que indican las acciones especificadas en ResponseActions y no por el código de respuesta informado en este campo o elemento, pero es una buena práctica que sea persistido por el mismo.
	ResponseCode int32 `json:"ResponseCode"`
	// Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible.
	ServiceVersion *string `json:"ServiceVersion,omitempty"`
	// Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma.
	Sequence *string `json:"Sequence,omitempty"`
	// Datos asociados a la seguridad de la transacción o de elementos sensibles.
	Security []SaleObjectSaleSecurity `json:"Security,omitempty"`
	// En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente.
	AdditionalInformation []SaleResponseObjectSaleResponseAdditionalInformation `json:"AdditionalInformation,omitempty"`
	// En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente.
	RequiredInformation []DebtPaymentObjectDebtPaymentRequiredInformation `json:"RequiredInformation,omitempty"`
	// ID privado del contexto de pago generado. Este elemento es equivalente al llamado RequestKey.
	InitialIdentification *string `json:"InitialIdentification,omitempty"`
	// ID publico del contexto de pago generado, mediante el cual se lo podra referenciar desde el formulario de pago
	InitialToken *string `json:"InitialToken,omitempty"`
	// Elementos a adiciónar a la URL a la cual se redigira para mostrar el formulario de pago
	CustomerRedirectAddress *string `json:"CustomerRedirectAddress,omitempty"`
	// Fecha y Hora de la transacción generada por el Punto de Venta - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14
	DateTime *time.Time `json:"DateTime,omitempty"`
	// ID de la operación a realizar, generado por el sistema externo
	TransactionIdentification *string `json:"TransactionIdentification,omitempty"`
	// Descripción del tipo de operación que se realizará
	TransactionDescription *string `json:"TransactionDescription,omitempty"`
	// ID del comercio que el para el cual se estara realizando la operación. Este valor puede ser enviado en lugar del SystemIdentification y CompanyIdentification, para luego ser traducido por el propio Plataforma a los valores configurados para ello. Ademas, puede relaciónarse el valor de BranchIdentification y/o POSIdentification de la misma forma. Si ello no se realiza, estos elementos tomaran valores genericos (por default) igual a 0.
	MerchantIdentification *string `json:"MerchantIdentification,omitempty"`
	// Cantidad de cuotas que permite este plan
	Installments *int32 `json:"Installments,omitempty"`
	// Tipo de plan utilizado para para realizar la operación
	FacilityType *int32 `json:"FacilityType,omitempty"`
	// código de Moneda - ISO 4217 <https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica <br />   * Num   - Alpha - Description <br />   * '032' - 'ARS' - Pesos Argentinos <br />   * '152' - 'CLP' - Pesos Chilenos <br/>   * '484' - 'MXN' - Pesos Mexicanos <br/>   * '840' - 'USD' - dólares Americanos <br/>   * '878' - 'EUR' - Euros <br/>   * '858' - 'UYU' - Pesos Uruguayos <br/>   * '878' - 'EUR' - Euros <br/>   * '986' - 'BRL' - Real Brasileño
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
	// Monto con la que se realizó transacción. Si este valor es recibido, la búsqueda de los planes será limitada con este criterio.
	Amount *float32 `json:"Amount,omitempty"`
	// .
	NetAmount *string `json:"NetAmount,omitempty"`
	Payer *SaleObjectSalePayer `json:"Payer,omitempty"`
	Customer *SaleObjectSaleCustomer `json:"Customer,omitempty"`
	Seller *SaleObjectSaleSeller `json:"Seller,omitempty"`
	// .
	TaxRefundType *string `json:"TaxRefundType,omitempty"`
	// .
	MerchantRedirectURL *string `json:"MerchantRedirectURL,omitempty"`
	// Fecha y Hora de fin de validez de La transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14
	TransactionValidThru *time.Time `json:"TransactionValidThru,omitempty"`
	// Detalle de Productos de la Operación.
	Products []SaleResponseObjectSaleResponseProducts `json:"Products,omitempty"`
	Configuration *SaleResponseObjectSaleResponseConfiguration `json:"Configuration,omitempty"`
}

// NewOrderGetResponseObjectOrderGetResponse instantiates a new OrderGetResponseObjectOrderGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderGetResponseObjectOrderGetResponse(responseActions []string, responseMessage string, responseCode int32) *OrderGetResponseObjectOrderGetResponse {
	this := OrderGetResponseObjectOrderGetResponse{}
	this.ResponseActions = responseActions
	this.ResponseMessage = responseMessage
	this.ResponseCode = responseCode
	return &this
}

// NewOrderGetResponseObjectOrderGetResponseWithDefaults instantiates a new OrderGetResponseObjectOrderGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderGetResponseObjectOrderGetResponseWithDefaults() *OrderGetResponseObjectOrderGetResponse {
	this := OrderGetResponseObjectOrderGetResponse{}
	return &this
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *OrderGetResponseObjectOrderGetResponse) SetRequestType(v string) {
	o.RequestType = &v
}

// GetResponseActions returns the ResponseActions field value
func (o *OrderGetResponseObjectOrderGetResponse) GetResponseActions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResponseActions
}

// GetResponseActionsOk returns a tuple with the ResponseActions field value
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetResponseActionsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseActions, true
}

// SetResponseActions sets field value
func (o *OrderGetResponseObjectOrderGetResponse) SetResponseActions(v []string) {
	o.ResponseActions = v
}

// GetResponseMessage returns the ResponseMessage field value
func (o *OrderGetResponseObjectOrderGetResponse) GetResponseMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetResponseMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseMessage, true
}

// SetResponseMessage sets field value
func (o *OrderGetResponseObjectOrderGetResponse) SetResponseMessage(v string) {
	o.ResponseMessage = v
}

// GetCompanyIdentification returns the CompanyIdentification field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetCompanyIdentification() string {
	if o == nil || o.CompanyIdentification == nil {
		var ret string
		return ret
	}
	return *o.CompanyIdentification
}

// GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetCompanyIdentificationOk() (*string, bool) {
	if o == nil || o.CompanyIdentification == nil {
		return nil, false
	}
	return o.CompanyIdentification, true
}

// HasCompanyIdentification returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasCompanyIdentification() bool {
	if o != nil && o.CompanyIdentification != nil {
		return true
	}

	return false
}

// SetCompanyIdentification gets a reference to the given string and assigns it to the CompanyIdentification field.
func (o *OrderGetResponseObjectOrderGetResponse) SetCompanyIdentification(v string) {
	o.CompanyIdentification = &v
}

// GetSystemIdentification returns the SystemIdentification field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetSystemIdentification() string {
	if o == nil || o.SystemIdentification == nil {
		var ret string
		return ret
	}
	return *o.SystemIdentification
}

// GetSystemIdentificationOk returns a tuple with the SystemIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetSystemIdentificationOk() (*string, bool) {
	if o == nil || o.SystemIdentification == nil {
		return nil, false
	}
	return o.SystemIdentification, true
}

// HasSystemIdentification returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasSystemIdentification() bool {
	if o != nil && o.SystemIdentification != nil {
		return true
	}

	return false
}

// SetSystemIdentification gets a reference to the given string and assigns it to the SystemIdentification field.
func (o *OrderGetResponseObjectOrderGetResponse) SetSystemIdentification(v string) {
	o.SystemIdentification = &v
}

// GetBranchIdentification returns the BranchIdentification field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetBranchIdentification() string {
	if o == nil || o.BranchIdentification == nil {
		var ret string
		return ret
	}
	return *o.BranchIdentification
}

// GetBranchIdentificationOk returns a tuple with the BranchIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetBranchIdentificationOk() (*string, bool) {
	if o == nil || o.BranchIdentification == nil {
		return nil, false
	}
	return o.BranchIdentification, true
}

// HasBranchIdentification returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasBranchIdentification() bool {
	if o != nil && o.BranchIdentification != nil {
		return true
	}

	return false
}

// SetBranchIdentification gets a reference to the given string and assigns it to the BranchIdentification field.
func (o *OrderGetResponseObjectOrderGetResponse) SetBranchIdentification(v string) {
	o.BranchIdentification = &v
}

// GetPOSIdentification returns the POSIdentification field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetPOSIdentification() string {
	if o == nil || o.POSIdentification == nil {
		var ret string
		return ret
	}
	return *o.POSIdentification
}

// GetPOSIdentificationOk returns a tuple with the POSIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetPOSIdentificationOk() (*string, bool) {
	if o == nil || o.POSIdentification == nil {
		return nil, false
	}
	return o.POSIdentification, true
}

// HasPOSIdentification returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasPOSIdentification() bool {
	if o != nil && o.POSIdentification != nil {
		return true
	}

	return false
}

// SetPOSIdentification gets a reference to the given string and assigns it to the POSIdentification field.
func (o *OrderGetResponseObjectOrderGetResponse) SetPOSIdentification(v string) {
	o.POSIdentification = &v
}

// GetForeignResponseCode returns the ForeignResponseCode field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetForeignResponseCode() string {
	if o == nil || o.ForeignResponseCode == nil {
		var ret string
		return ret
	}
	return *o.ForeignResponseCode
}

// GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetForeignResponseCodeOk() (*string, bool) {
	if o == nil || o.ForeignResponseCode == nil {
		return nil, false
	}
	return o.ForeignResponseCode, true
}

// HasForeignResponseCode returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasForeignResponseCode() bool {
	if o != nil && o.ForeignResponseCode != nil {
		return true
	}

	return false
}

// SetForeignResponseCode gets a reference to the given string and assigns it to the ForeignResponseCode field.
func (o *OrderGetResponseObjectOrderGetResponse) SetForeignResponseCode(v string) {
	o.ForeignResponseCode = &v
}

// GetResponseCode returns the ResponseCode field value
func (o *OrderGetResponseObjectOrderGetResponse) GetResponseCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetResponseCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseCode, true
}

// SetResponseCode sets field value
func (o *OrderGetResponseObjectOrderGetResponse) SetResponseCode(v int32) {
	o.ResponseCode = v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *OrderGetResponseObjectOrderGetResponse) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetSequence() string {
	if o == nil || o.Sequence == nil {
		var ret string
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetSequenceOk() (*string, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given string and assigns it to the Sequence field.
func (o *OrderGetResponseObjectOrderGetResponse) SetSequence(v string) {
	o.Sequence = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetSecurity() []SaleObjectSaleSecurity {
	if o == nil || o.Security == nil {
		var ret []SaleObjectSaleSecurity
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetSecurityOk() ([]SaleObjectSaleSecurity, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []SaleObjectSaleSecurity and assigns it to the Security field.
func (o *OrderGetResponseObjectOrderGetResponse) SetSecurity(v []SaleObjectSaleSecurity) {
	o.Security = v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation {
	if o == nil || o.AdditionalInformation == nil {
		var ret []SaleResponseObjectSaleResponseAdditionalInformation
		return ret
	}
	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetAdditionalInformationOk() ([]SaleResponseObjectSaleResponseAdditionalInformation, bool) {
	if o == nil || o.AdditionalInformation == nil {
		return nil, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasAdditionalInformation() bool {
	if o != nil && o.AdditionalInformation != nil {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given []SaleResponseObjectSaleResponseAdditionalInformation and assigns it to the AdditionalInformation field.
func (o *OrderGetResponseObjectOrderGetResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation) {
	o.AdditionalInformation = v
}

// GetRequiredInformation returns the RequiredInformation field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetRequiredInformation() []DebtPaymentObjectDebtPaymentRequiredInformation {
	if o == nil || o.RequiredInformation == nil {
		var ret []DebtPaymentObjectDebtPaymentRequiredInformation
		return ret
	}
	return o.RequiredInformation
}

// GetRequiredInformationOk returns a tuple with the RequiredInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetRequiredInformationOk() ([]DebtPaymentObjectDebtPaymentRequiredInformation, bool) {
	if o == nil || o.RequiredInformation == nil {
		return nil, false
	}
	return o.RequiredInformation, true
}

// HasRequiredInformation returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasRequiredInformation() bool {
	if o != nil && o.RequiredInformation != nil {
		return true
	}

	return false
}

// SetRequiredInformation gets a reference to the given []DebtPaymentObjectDebtPaymentRequiredInformation and assigns it to the RequiredInformation field.
func (o *OrderGetResponseObjectOrderGetResponse) SetRequiredInformation(v []DebtPaymentObjectDebtPaymentRequiredInformation) {
	o.RequiredInformation = v
}

// GetInitialIdentification returns the InitialIdentification field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetInitialIdentification() string {
	if o == nil || o.InitialIdentification == nil {
		var ret string
		return ret
	}
	return *o.InitialIdentification
}

// GetInitialIdentificationOk returns a tuple with the InitialIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetInitialIdentificationOk() (*string, bool) {
	if o == nil || o.InitialIdentification == nil {
		return nil, false
	}
	return o.InitialIdentification, true
}

// HasInitialIdentification returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasInitialIdentification() bool {
	if o != nil && o.InitialIdentification != nil {
		return true
	}

	return false
}

// SetInitialIdentification gets a reference to the given string and assigns it to the InitialIdentification field.
func (o *OrderGetResponseObjectOrderGetResponse) SetInitialIdentification(v string) {
	o.InitialIdentification = &v
}

// GetInitialToken returns the InitialToken field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetInitialToken() string {
	if o == nil || o.InitialToken == nil {
		var ret string
		return ret
	}
	return *o.InitialToken
}

// GetInitialTokenOk returns a tuple with the InitialToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetInitialTokenOk() (*string, bool) {
	if o == nil || o.InitialToken == nil {
		return nil, false
	}
	return o.InitialToken, true
}

// HasInitialToken returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasInitialToken() bool {
	if o != nil && o.InitialToken != nil {
		return true
	}

	return false
}

// SetInitialToken gets a reference to the given string and assigns it to the InitialToken field.
func (o *OrderGetResponseObjectOrderGetResponse) SetInitialToken(v string) {
	o.InitialToken = &v
}

// GetCustomerRedirectAddress returns the CustomerRedirectAddress field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetCustomerRedirectAddress() string {
	if o == nil || o.CustomerRedirectAddress == nil {
		var ret string
		return ret
	}
	return *o.CustomerRedirectAddress
}

// GetCustomerRedirectAddressOk returns a tuple with the CustomerRedirectAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetCustomerRedirectAddressOk() (*string, bool) {
	if o == nil || o.CustomerRedirectAddress == nil {
		return nil, false
	}
	return o.CustomerRedirectAddress, true
}

// HasCustomerRedirectAddress returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasCustomerRedirectAddress() bool {
	if o != nil && o.CustomerRedirectAddress != nil {
		return true
	}

	return false
}

// SetCustomerRedirectAddress gets a reference to the given string and assigns it to the CustomerRedirectAddress field.
func (o *OrderGetResponseObjectOrderGetResponse) SetCustomerRedirectAddress(v string) {
	o.CustomerRedirectAddress = &v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetDateTime() time.Time {
	if o == nil || o.DateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetDateTimeOk() (*time.Time, bool) {
	if o == nil || o.DateTime == nil {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasDateTime() bool {
	if o != nil && o.DateTime != nil {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *OrderGetResponseObjectOrderGetResponse) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetTransactionIdentification returns the TransactionIdentification field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionIdentification() string {
	if o == nil || o.TransactionIdentification == nil {
		var ret string
		return ret
	}
	return *o.TransactionIdentification
}

// GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionIdentificationOk() (*string, bool) {
	if o == nil || o.TransactionIdentification == nil {
		return nil, false
	}
	return o.TransactionIdentification, true
}

// HasTransactionIdentification returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasTransactionIdentification() bool {
	if o != nil && o.TransactionIdentification != nil {
		return true
	}

	return false
}

// SetTransactionIdentification gets a reference to the given string and assigns it to the TransactionIdentification field.
func (o *OrderGetResponseObjectOrderGetResponse) SetTransactionIdentification(v string) {
	o.TransactionIdentification = &v
}

// GetTransactionDescription returns the TransactionDescription field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionDescription() string {
	if o == nil || o.TransactionDescription == nil {
		var ret string
		return ret
	}
	return *o.TransactionDescription
}

// GetTransactionDescriptionOk returns a tuple with the TransactionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionDescriptionOk() (*string, bool) {
	if o == nil || o.TransactionDescription == nil {
		return nil, false
	}
	return o.TransactionDescription, true
}

// HasTransactionDescription returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasTransactionDescription() bool {
	if o != nil && o.TransactionDescription != nil {
		return true
	}

	return false
}

// SetTransactionDescription gets a reference to the given string and assigns it to the TransactionDescription field.
func (o *OrderGetResponseObjectOrderGetResponse) SetTransactionDescription(v string) {
	o.TransactionDescription = &v
}

// GetMerchantIdentification returns the MerchantIdentification field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetMerchantIdentification() string {
	if o == nil || o.MerchantIdentification == nil {
		var ret string
		return ret
	}
	return *o.MerchantIdentification
}

// GetMerchantIdentificationOk returns a tuple with the MerchantIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetMerchantIdentificationOk() (*string, bool) {
	if o == nil || o.MerchantIdentification == nil {
		return nil, false
	}
	return o.MerchantIdentification, true
}

// HasMerchantIdentification returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasMerchantIdentification() bool {
	if o != nil && o.MerchantIdentification != nil {
		return true
	}

	return false
}

// SetMerchantIdentification gets a reference to the given string and assigns it to the MerchantIdentification field.
func (o *OrderGetResponseObjectOrderGetResponse) SetMerchantIdentification(v string) {
	o.MerchantIdentification = &v
}

// GetInstallments returns the Installments field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetInstallments() int32 {
	if o == nil || o.Installments == nil {
		var ret int32
		return ret
	}
	return *o.Installments
}

// GetInstallmentsOk returns a tuple with the Installments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetInstallmentsOk() (*int32, bool) {
	if o == nil || o.Installments == nil {
		return nil, false
	}
	return o.Installments, true
}

// HasInstallments returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasInstallments() bool {
	if o != nil && o.Installments != nil {
		return true
	}

	return false
}

// SetInstallments gets a reference to the given int32 and assigns it to the Installments field.
func (o *OrderGetResponseObjectOrderGetResponse) SetInstallments(v int32) {
	o.Installments = &v
}

// GetFacilityType returns the FacilityType field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetFacilityType() int32 {
	if o == nil || o.FacilityType == nil {
		var ret int32
		return ret
	}
	return *o.FacilityType
}

// GetFacilityTypeOk returns a tuple with the FacilityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetFacilityTypeOk() (*int32, bool) {
	if o == nil || o.FacilityType == nil {
		return nil, false
	}
	return o.FacilityType, true
}

// HasFacilityType returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasFacilityType() bool {
	if o != nil && o.FacilityType != nil {
		return true
	}

	return false
}

// SetFacilityType gets a reference to the given int32 and assigns it to the FacilityType field.
func (o *OrderGetResponseObjectOrderGetResponse) SetFacilityType(v int32) {
	o.FacilityType = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *OrderGetResponseObjectOrderGetResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *OrderGetResponseObjectOrderGetResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetNetAmount returns the NetAmount field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetNetAmount() string {
	if o == nil || o.NetAmount == nil {
		var ret string
		return ret
	}
	return *o.NetAmount
}

// GetNetAmountOk returns a tuple with the NetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetNetAmountOk() (*string, bool) {
	if o == nil || o.NetAmount == nil {
		return nil, false
	}
	return o.NetAmount, true
}

// HasNetAmount returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasNetAmount() bool {
	if o != nil && o.NetAmount != nil {
		return true
	}

	return false
}

// SetNetAmount gets a reference to the given string and assigns it to the NetAmount field.
func (o *OrderGetResponseObjectOrderGetResponse) SetNetAmount(v string) {
	o.NetAmount = &v
}

// GetPayer returns the Payer field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetPayer() SaleObjectSalePayer {
	if o == nil || o.Payer == nil {
		var ret SaleObjectSalePayer
		return ret
	}
	return *o.Payer
}

// GetPayerOk returns a tuple with the Payer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetPayerOk() (*SaleObjectSalePayer, bool) {
	if o == nil || o.Payer == nil {
		return nil, false
	}
	return o.Payer, true
}

// HasPayer returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasPayer() bool {
	if o != nil && o.Payer != nil {
		return true
	}

	return false
}

// SetPayer gets a reference to the given SaleObjectSalePayer and assigns it to the Payer field.
func (o *OrderGetResponseObjectOrderGetResponse) SetPayer(v SaleObjectSalePayer) {
	o.Payer = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetCustomer() SaleObjectSaleCustomer {
	if o == nil || o.Customer == nil {
		var ret SaleObjectSaleCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetCustomerOk() (*SaleObjectSaleCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given SaleObjectSaleCustomer and assigns it to the Customer field.
func (o *OrderGetResponseObjectOrderGetResponse) SetCustomer(v SaleObjectSaleCustomer) {
	o.Customer = &v
}

// GetSeller returns the Seller field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetSeller() SaleObjectSaleSeller {
	if o == nil || o.Seller == nil {
		var ret SaleObjectSaleSeller
		return ret
	}
	return *o.Seller
}

// GetSellerOk returns a tuple with the Seller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetSellerOk() (*SaleObjectSaleSeller, bool) {
	if o == nil || o.Seller == nil {
		return nil, false
	}
	return o.Seller, true
}

// HasSeller returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasSeller() bool {
	if o != nil && o.Seller != nil {
		return true
	}

	return false
}

// SetSeller gets a reference to the given SaleObjectSaleSeller and assigns it to the Seller field.
func (o *OrderGetResponseObjectOrderGetResponse) SetSeller(v SaleObjectSaleSeller) {
	o.Seller = &v
}

// GetTaxRefundType returns the TaxRefundType field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetTaxRefundType() string {
	if o == nil || o.TaxRefundType == nil {
		var ret string
		return ret
	}
	return *o.TaxRefundType
}

// GetTaxRefundTypeOk returns a tuple with the TaxRefundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetTaxRefundTypeOk() (*string, bool) {
	if o == nil || o.TaxRefundType == nil {
		return nil, false
	}
	return o.TaxRefundType, true
}

// HasTaxRefundType returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasTaxRefundType() bool {
	if o != nil && o.TaxRefundType != nil {
		return true
	}

	return false
}

// SetTaxRefundType gets a reference to the given string and assigns it to the TaxRefundType field.
func (o *OrderGetResponseObjectOrderGetResponse) SetTaxRefundType(v string) {
	o.TaxRefundType = &v
}

// GetMerchantRedirectURL returns the MerchantRedirectURL field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetMerchantRedirectURL() string {
	if o == nil || o.MerchantRedirectURL == nil {
		var ret string
		return ret
	}
	return *o.MerchantRedirectURL
}

// GetMerchantRedirectURLOk returns a tuple with the MerchantRedirectURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetMerchantRedirectURLOk() (*string, bool) {
	if o == nil || o.MerchantRedirectURL == nil {
		return nil, false
	}
	return o.MerchantRedirectURL, true
}

// HasMerchantRedirectURL returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasMerchantRedirectURL() bool {
	if o != nil && o.MerchantRedirectURL != nil {
		return true
	}

	return false
}

// SetMerchantRedirectURL gets a reference to the given string and assigns it to the MerchantRedirectURL field.
func (o *OrderGetResponseObjectOrderGetResponse) SetMerchantRedirectURL(v string) {
	o.MerchantRedirectURL = &v
}

// GetTransactionValidThru returns the TransactionValidThru field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionValidThru() time.Time {
	if o == nil || o.TransactionValidThru == nil {
		var ret time.Time
		return ret
	}
	return *o.TransactionValidThru
}

// GetTransactionValidThruOk returns a tuple with the TransactionValidThru field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetTransactionValidThruOk() (*time.Time, bool) {
	if o == nil || o.TransactionValidThru == nil {
		return nil, false
	}
	return o.TransactionValidThru, true
}

// HasTransactionValidThru returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasTransactionValidThru() bool {
	if o != nil && o.TransactionValidThru != nil {
		return true
	}

	return false
}

// SetTransactionValidThru gets a reference to the given time.Time and assigns it to the TransactionValidThru field.
func (o *OrderGetResponseObjectOrderGetResponse) SetTransactionValidThru(v time.Time) {
	o.TransactionValidThru = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetProducts() []SaleResponseObjectSaleResponseProducts {
	if o == nil || o.Products == nil {
		var ret []SaleResponseObjectSaleResponseProducts
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetProductsOk() ([]SaleResponseObjectSaleResponseProducts, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SaleResponseObjectSaleResponseProducts and assigns it to the Products field.
func (o *OrderGetResponseObjectOrderGetResponse) SetProducts(v []SaleResponseObjectSaleResponseProducts) {
	o.Products = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *OrderGetResponseObjectOrderGetResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration {
	if o == nil || o.Configuration == nil {
		var ret SaleResponseObjectSaleResponseConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetResponseObjectOrderGetResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *OrderGetResponseObjectOrderGetResponse) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given SaleResponseObjectSaleResponseConfiguration and assigns it to the Configuration field.
func (o *OrderGetResponseObjectOrderGetResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration) {
	o.Configuration = &v
}

func (o OrderGetResponseObjectOrderGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestType != nil {
		toSerialize["RequestType"] = o.RequestType
	}
	if true {
		toSerialize["ResponseActions"] = o.ResponseActions
	}
	if true {
		toSerialize["ResponseMessage"] = o.ResponseMessage
	}
	if o.CompanyIdentification != nil {
		toSerialize["CompanyIdentification"] = o.CompanyIdentification
	}
	if o.SystemIdentification != nil {
		toSerialize["SystemIdentification"] = o.SystemIdentification
	}
	if o.BranchIdentification != nil {
		toSerialize["BranchIdentification"] = o.BranchIdentification
	}
	if o.POSIdentification != nil {
		toSerialize["POSIdentification"] = o.POSIdentification
	}
	if o.ForeignResponseCode != nil {
		toSerialize["ForeignResponseCode"] = o.ForeignResponseCode
	}
	if true {
		toSerialize["ResponseCode"] = o.ResponseCode
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
	if o.AdditionalInformation != nil {
		toSerialize["AdditionalInformation"] = o.AdditionalInformation
	}
	if o.RequiredInformation != nil {
		toSerialize["RequiredInformation"] = o.RequiredInformation
	}
	if o.InitialIdentification != nil {
		toSerialize["InitialIdentification"] = o.InitialIdentification
	}
	if o.InitialToken != nil {
		toSerialize["InitialToken"] = o.InitialToken
	}
	if o.CustomerRedirectAddress != nil {
		toSerialize["CustomerRedirectAddress"] = o.CustomerRedirectAddress
	}
	if o.DateTime != nil {
		toSerialize["DateTime"] = o.DateTime
	}
	if o.TransactionIdentification != nil {
		toSerialize["TransactionIdentification"] = o.TransactionIdentification
	}
	if o.TransactionDescription != nil {
		toSerialize["TransactionDescription"] = o.TransactionDescription
	}
	if o.MerchantIdentification != nil {
		toSerialize["MerchantIdentification"] = o.MerchantIdentification
	}
	if o.Installments != nil {
		toSerialize["Installments"] = o.Installments
	}
	if o.FacilityType != nil {
		toSerialize["FacilityType"] = o.FacilityType
	}
	if o.CurrencyCode != nil {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if o.Amount != nil {
		toSerialize["Amount"] = o.Amount
	}
	if o.NetAmount != nil {
		toSerialize["NetAmount"] = o.NetAmount
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
	if o.TaxRefundType != nil {
		toSerialize["TaxRefundType"] = o.TaxRefundType
	}
	if o.MerchantRedirectURL != nil {
		toSerialize["MerchantRedirectURL"] = o.MerchantRedirectURL
	}
	if o.TransactionValidThru != nil {
		toSerialize["TransactionValidThru"] = o.TransactionValidThru
	}
	if o.Products != nil {
		toSerialize["Products"] = o.Products
	}
	if o.Configuration != nil {
		toSerialize["Configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableOrderGetResponseObjectOrderGetResponse struct {
	value *OrderGetResponseObjectOrderGetResponse
	isSet bool
}

func (v NullableOrderGetResponseObjectOrderGetResponse) Get() *OrderGetResponseObjectOrderGetResponse {
	return v.value
}

func (v *NullableOrderGetResponseObjectOrderGetResponse) Set(val *OrderGetResponseObjectOrderGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderGetResponseObjectOrderGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderGetResponseObjectOrderGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderGetResponseObjectOrderGetResponse(val *OrderGetResponseObjectOrderGetResponse) *NullableOrderGetResponseObjectOrderGetResponse {
	return &NullableOrderGetResponseObjectOrderGetResponse{value: val, isSet: true}
}

func (v NullableOrderGetResponseObjectOrderGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderGetResponseObjectOrderGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

