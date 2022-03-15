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

// VoidResponseObjectVoidResponse struct for VoidResponseObjectVoidResponse
type VoidResponseObjectVoidResponse struct {
	// Código de Respuesta Interno de la plataforma, el POS debe actuar por lo que indican las acciones especificadas en ResponseActions y no por el código de respuesta informado en este campo o elemento, pero es una buena práctica que sea persistido por el mismo.
	ResponseCode int32 `json:"ResponseCode"`
	// Acciones a realizar por parte del POS y/o PINPAD en base al resultado de la operación que ha sido procesada. Cada uno de estos actions o acciones están concatenadas por comas. Los posibles actions son OK, Approve, Refuse, IssuerCall, Tickets, WithHold, GetCard, UseTerminalToAuthorize, ConfigurationError, SystemError, ResourceError, ProcessError, Completed, Configure, Display, EnableService y Print.
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
	// Tipo de Operación que se requirió, contendrá el mismo valor que se recibió en el requerimiento, sobre formatos que no soportan elementos complejos o compuestos.
	RequestType *string `json:"RequestType,omitempty"`
	// Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible.
	ServiceVersion *string `json:"ServiceVersion,omitempty"`
	// Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma.
	Sequence *string `json:"Sequence,omitempty"`
	// Datos asociados a la seguridad de la transacción o de elementos sensibles.
	Security []SaleObjectSaleSecurity `json:"Security,omitempty"`
	// ID que identifica a un grupo de transacciones que serán confirmadas o canceladas
	Block *string `json:"Block,omitempty"`
	// ID generado para la identificación por parte del Plataforma de la información generada en la ejecución de un GetCard o un Payment Method. Será necesario para que un mensaje de Sale, Void, Payment Method o Enrollment identifique el contexto generado y lo utilice para esa operación.
	RequestKey *string `json:"RequestKey,omitempty"`
	// Nueva forma de enviar las llaves disponibles para esta caja.
	WorkingKeys []SaleResponseObjectSaleResponseWorkingKeys `json:"WorkingKeys,omitempty"`
	// En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente.
	RequiredInformation []SaleResponseObjectSaleResponseRequiredInformation `json:"RequiredInformation,omitempty"`
	// Tipo de Operación que se está requiriendo, solo necesario sobre formatos que no soportan elementos complejos o compuestos.
	AnswerType *string `json:"AnswerType,omitempty"`
	// Código de identificación, generado por Plataforma, de la operación realizada'.
	AnswerKey *string `json:"AnswerKey,omitempty"`
	// Código público de identificación, generado por Plataforma, de la operación realizada'.
	PublicAnswerKey *string `json:"PublicAnswerKey,omitempty"`
	// Flag indicador de generación de reversa para la última operación reversable.
	WasReversePrevious *int32 `json:"WasReversePrevious,omitempty"`
	// ID que identifica a la operación que acaba de ser reversada.
	ReversedAnswerKey *string `json:"ReversedAnswerKey,omitempty"`
	// Secuencia de la transacción que fue reversada.
	ReversedSequence *string `json:"ReversedSequence,omitempty"`
	// ID del bloque de transacciones que ha sido confirmado de forma automática (es decir, sin recibir un requerimiento de BlockClose). Este escenario se presentará si el Plataforma así se ha configurado para actuar bajo esa circunstancia.
	CommittedBlock *string `json:"CommittedBlock,omitempty"`
	// ID del bloque de transacciones que ha sido cancelado de forma automática (es decir, sin recibir un requerimiento de BlockClose). Este escenario se presentara si el Plataforma así se ha configurado para actuar bajo esa circunstancia.
	ReversedBlock *string `json:"ReversedBlock,omitempty"`
	// Identificador Unívoco del Mensaje ( UUID v5 ).
	MessageID *string `json:"MessageID,omitempty"`
	// Dirección IP del Server que atiende el requerimiento.
	ServerAddress *string `json:"ServerAddress,omitempty"`
	// Instancia de Server que atiende el requerimiento.
	ServerInstance *string `json:"ServerInstance,omitempty"`
	// Nombre del Nodo que atendió el requerimiento.
	ServerNodeName *string `json:"ServerNodeName,omitempty"`
	// Versión del  Adaptador de Protocolo Entrante que atiende el Requerimiento.
	AdapterInputVersion *string `json:"AdapterInputVersion,omitempty"`
	// Dirección IP del Adaptador de Protocolo Entrante que atiende el requerimiento.
	AdapterInputAddress *string `json:"AdapterInputAddress,omitempty"`
	// Nombre del Nodo del Adaptador de Protocolo Entrante que atiende el requerimiento.
	AdapterInputNodeName *string `json:"AdapterInputNodeName,omitempty"`
	// Versión del  Adaptador de Protocolo Saliente que atiende el Requerimiento.
	AdapterOutputVersion *string `json:"AdapterOutputVersion,omitempty"`
	// Dirección IP  del  Adaptador de Protocolo Saliente que atiende el Requerimiento.
	AdapterOutputAddress *string `json:"AdapterOutputAddress,omitempty"`
	// Nombre del Nodo  del  Adaptador de Protocolo Saliente que atiende el Requerimiento.
	AdapterOutputNodeName *string `json:"AdapterOutputNodeName,omitempty"`
	// En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente.
	AdditionalInformation []SaleResponseObjectSaleResponseAdditionalInformation `json:"AdditionalInformation,omitempty"`
	// Información adicional/Mensaje promocional/Leyenda de respuesta a imprimir en el ticket de la operación. Cada línea de este mensaje sera un elemento dentro del array.
	PrinterResponseMessage []string `json:"PrinterResponseMessage,omitempty"`
	// Información adicional/Mensaje promocional/Leyenda de respuesta a mostrar en pantalla en el ticket de la operación. Cada línea de este mensaje será un elemento dentro del array.
	DisplayResponseMessage []string `json:"DisplayResponseMessage,omitempty"`
	// código de Moneda - ISO 4217 <https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica <br />   * Num   - Alpha - Description <br />   * '032' - 'ARS' - Pesos Argentinos <br />   * '152' - 'CLP' - Pesos Chilenos <br/>   * '484' - 'MXN' - Pesos Mexicanos <br/>   * '840' - 'USD' - dólares Americanos <br/>   * '878' - 'EUR' - Euros <br/>   * '858' - 'UYU' - Pesos Uruguayos <br/>   * '878' - 'EUR' - Euros <br/>   * '986' - 'BRL' - Real Brasileño
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
	// Descripción del tipo de cambio utilizado en la transacción.
	CurrencyDescription *string `json:"CurrencyDescription,omitempty"`
	// Tipo de transacción
	TransactionDescription *string `json:"TransactionDescription,omitempty"`
	// Modo en que fue realizada la transacción
	TransactionResolutionMode *string `json:"TransactionResolutionMode,omitempty"`
	// Símbolo monetario del tipo de cambio utilizado en la transacción.
	CurrencySymbol *string `json:"CurrencySymbol,omitempty"`
	// Cantidad de cuotas utilizadas para realizar la operación
	FacilityPayments *int32 `json:"FacilityPayments,omitempty"`
	// Tipo de plan utilizado para para realizar la operación
	FacilityType *int32 `json:"FacilityType,omitempty"`
	// Monto final a pagar en cada una de las cuotas en las que se divida la compra
	ValueFacilityPayments *float32 `json:"ValueFacilityPayments,omitempty"`
	// Importe o Monto de la Transacción.
	Amount *float32 `json:"Amount,omitempty"`
	// Monto con la que se realizó transacción. Si este valor es recibido, la búsqueda de los planes será limitada con este criterio.
	AlternativeAmount *float32 `json:"AlternativeAmount,omitempty"`
	// Importe o Monto de la Transacción que efectivamente se cobró.  Si se envía en Void o Return en lugar de Amount, se genera un Ajuste si el Host lo soporta.
	AmountCharged *float32 `json:"AmountCharged,omitempty"`
	// Importe o Monto de la Transacción a aplicar.
	AmountToApply *float32 `json:"AmountToApply,omitempty"`
	// Monto del dinero en efectivo (cashback).
	CashbackAmount *float32 `json:"CashbackAmount,omitempty"`
	// Importe o Monto de la Propina.
	TipAmount *float32 `json:"TipAmount,omitempty"`
	// Monto libre de costos financerios e impuestos por el que la venta se realizó. El monto cobrado realmente no es este, dado que no incluye las tasas e impuestos
	RequestAmount *float32 `json:"RequestAmount,omitempty"`
	// Monto Sujeto a Promoción monto de la operación
	PromotedAmount *float32 `json:"PromotedAmount,omitempty"`
	// Token asociado a la Credencial Enrolada
	CredentialToken *string `json:"CredentialToken,omitempty"`
	// Emisor del Token asociado a la credencial enrolada
	CredentialIssuerToken *string `json:"CredentialIssuerToken,omitempty"`
	// Tokens.
	InputTokens []SaleObjectSaleInputTokens `json:"InputTokens,omitempty"`
	// Tokens.
	OutputTokens []SaleObjectSaleInputTokens `json:"OutputTokens,omitempty"`
	// Monto del recargo impositivo aplicado al costo financiero que la transacción tiene
	TaxFinancialCostAmount *float32 `json:"TaxFinancialCostAmount,omitempty"`
	// Porcentaje de recargo impositivo a aplicar sobre el monto del costo financiero
	TaxFinancialCostPercentage *float32 `json:"TaxFinancialCostPercentage,omitempty"`
	// Monto del Costo financiero total generado en base al plan elegido
	FinancialCostAmount *float32 `json:"FinancialCostAmount,omitempty"`
	// Porcentaje del costo financiero a aplicar al monto de la transacción
	FinancialCostPercentage *float32 `json:"FinancialCostPercentage,omitempty"`
	// Código de Resultado retornado por el Host Adquirente.
	HostResultCode *int32 `json:"HostResultCode,omitempty"`
	// Mensaje Retornado por el Host Adquirente, normalmente asociado al valor de HostResultCode
	HostMessage *string `json:"HostMessage,omitempty"`
	// Código de autorización retornado por el Host que resuelve la transacción.
	HostCode *string `json:"HostCode,omitempty"`
	// Número de identificación del host al cual fue enviada la petición, y por el cual fue finalmente procesada.
	HostID *int32 `json:"HostID,omitempty"`
	// Identificador del usuario que está realizando la Transacción.
	UserID *string `json:"UserID,omitempty"`
	// Nombre del Emisor de la Credencial o Tarjeta que se usó en la transacción.
	IssuerName *string `json:"IssuerName,omitempty"`
	// Fecha y Hora de la transacción retornada por el Host que resuelve la Transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14
	HostDateTime *time.Time `json:"HostDateTime,omitempty"`
	// Fecha y hora de transmisión de la operación hacia el host - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14
	TransmitionDateTime *time.Time `json:"TransmitionDateTime,omitempty"`
	// Código de Resultado retornado por el Host Adquirente.
	AuthResultCode *string `json:"AuthResultCode,omitempty"`
	// Código de procesamiento de la operación enviada al host. Elemento 3 del protocolo de comunicaciones ISO 8583, formato de mensajes utilizado por los hosts para realizar operaciones financieras.
	AuthHostProcessCode *string `json:"AuthHostProcessCode,omitempty"`
	// Elemento 0 del protocolo de comunicaciones ISO 8583, formato de mensajes utilizado por los hosts para realizar operaciones financieras. El valor de este campo es el que devuelve el host en una respuesta a una petición.
	AuthHostMsgType *string `json:"AuthHostMsgType,omitempty"`
	// Mensaje Retornado por el Host Adquirente, normalmente asociado al valor de AuthResultCode
	AuthHostMessage *string `json:"AuthHostMessage,omitempty"`
	// Elemento 0 del protocolo de comunicaciones ISO 8583, formato de mensajes utilizado por los hosts para realizar operaciones financieras. El valor de este campo es el que se envio al host en el envio de una petición.
	HostMsgType *string `json:"HostMsgType,omitempty"`
	// Código de autorización retornado por el Host que resuelve la transacción.
	AuthCode *string `json:"AuthCode,omitempty"`
	// Fecha y Hora de la transacción retornada por el Host que resuelve la Transacción - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14
	AuthDateTime *time.Time `json:"AuthDateTime,omitempty"`
	// Número Ticket  o Voucher Generado para la Plataforma.
	AuthTicket *int32 `json:"AuthTicket,omitempty"`
	// Número de identificación de la transacción, utilizado por la mayoría de los hosts para realizar anulaciones y devoluciones.
	AuthRRN *string `json:"AuthRRN,omitempty"`
	// Tipo de autenticación
	TransactionAuthenticationType *string `json:"TransactionAuthenticationType,omitempty"`
	// Identificador que genera el Host Adquirente para la Transacción. En algunos podrá ser igual al AuthRRN.
	IdentifierForTheAdquirer *string `json:"IdentifierForTheAdquirer,omitempty"`
	// Identificador que genera el Host o Plataforma que resuelve la transacción.
	IdentifierForTheResolutor *string `json:"IdentifierForTheResolutor,omitempty"`
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
	// Modo de ingreso de los datos de la tarjeta. Los posibles valores significan: C - EMV Chip / B - Banda magnética / L - Contactless Chip / S - Contactless Banda / M - Manual (Tarjeta Presente) / T - Digitada (Tarjeta no Presente) / E - ECOMMERCE (Ventas por Internet)  / F - FALLBACK (Banda por falla en Chip) / K - TOKEN / R - Recurring ( Pagos Recurrentes )
	CardReadMode *string `json:"CardReadMode,omitempty"`
	// Descripción del modo de ingreso con el que fue leída la tarjeta
	CardReadModeDescription *string `json:"CardReadModeDescription,omitempty"`
	// Nombre de la Tarjeta que se usó en la transacción, usado para la impresión del voucher.
	CardDescription *string `json:"CardDescription,omitempty"`
	// Descripción del modo de ingreso utilizado para capturar los datos de la tarjeta.
	CardTypeDescription *string `json:"CardTypeDescription,omitempty"`
	CardCategory *SaleResponseObjectSaleResponseCardCategory `json:"CardCategory,omitempty"`
	// Número de Tarjeta. En el caso de las respuestas el mismo estará enmascarado.
	CardNumber *string `json:"CardNumber,omitempty"`
	// Número de tarjeta enmascarado, según indica la parametrización en la base de datos. Se utilizará para imprimir en el cupón.
	CardNumberMasked *string `json:"CardNumberMasked,omitempty"`
	// Hash de la tarjeta generado por la plataforma.
	CardHashing *string `json:"CardHashing,omitempty"`
	// Fecha de vencimiento de la tarjeta. Este dato sera necesario si el modo de ingreso fue manual/digitada.
	CardExp *string `json:"CardExp,omitempty"`
	// Tags EMV en format TLV recibidos desde el Host.
	CardCryptogramResponse *string `json:"CardCryptogramResponse,omitempty"`
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
	Payer *SaleObjectSalePayer `json:"Payer,omitempty"`
	Customer *SaleObjectSaleCustomer `json:"Customer,omitempty"`
	MerchantCategory *SaleResponseObjectSaleResponseMerchantCategory `json:"MerchantCategory,omitempty"`
	// Detalle de Productos de la Operación.
	Products []SaleResponseObjectSaleResponseProducts `json:"Products,omitempty"`
	PaymentMethod *SaleResponseObjectSaleResponsePaymentMethod `json:"PaymentMethod,omitempty"`
	Plans *SaleResponseObjectSaleResponsePlans `json:"Plans,omitempty"`
	// Descripción del plan utilizado para para realizar la operación
	PlanDescription *string `json:"PlanDescription,omitempty"`
	// Identificador de Versión utilizada por Plataforma en la evaluación del Plan
	PlanConfigVersion *int32 `json:"PlanConfigVersion,omitempty"`
	// Elemento Compuesto que indica qué Tickets intervienen en la transacción y si deben ser digitalizados o impresos.
	Tickets []SaleResponseObjectSaleResponseTickets `json:"Tickets,omitempty"`
	Configuration *SaleResponseObjectSaleResponseConfiguration `json:"Configuration,omitempty"`
	// Identificador Unívoco de la Transacción que se quiere Referenciar, usado en Deposit, Void, Return, etc. O sea en las transacciones que hacen referencia a una Transacción previa. El valor fue obtenido en el campo o elemento AnswerKey de la Respuesta de la transacción referenciada. 
	OrigAnswerKey *string `json:"OrigAnswerKey,omitempty"`
	// Fecha y Hora de la Transacción previa que se está referenciando. - RFC3339 https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14
	OrigAuthDateTime *time.Time `json:"OrigAuthDateTime,omitempty"`
	// Número de Ticket o Voucher de la Transacción Original Referenciada.
	OrigAuthTicket *int32 `json:"OrigAuthTicket,omitempty"`
	// Número de terminal de la Transacción previa que se está referenciando.
	OrigAuthTerminalTrace *int32 `json:"OrigAuthTerminalTrace,omitempty"`
	// Código de autorización de la transacción original.
	OrigAuthCode *string `json:"OrigAuthCode,omitempty"`
	// ID de la marca con la cual la tarjeta fue identificada
	PaymentMethodId *int32 `json:"PaymentMethodId,omitempty"`
	// Descripción o nombre de la marca con la cual la tarjeta fue identificada
	PaymentMethodDescription *string `json:"PaymentMethodDescription,omitempty"`
}

// NewVoidResponseObjectVoidResponse instantiates a new VoidResponseObjectVoidResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoidResponseObjectVoidResponse(responseCode int32, responseActions []string, responseMessage string) *VoidResponseObjectVoidResponse {
	this := VoidResponseObjectVoidResponse{}
	this.ResponseCode = responseCode
	this.ResponseActions = responseActions
	this.ResponseMessage = responseMessage
	return &this
}

// NewVoidResponseObjectVoidResponseWithDefaults instantiates a new VoidResponseObjectVoidResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoidResponseObjectVoidResponseWithDefaults() *VoidResponseObjectVoidResponse {
	this := VoidResponseObjectVoidResponse{}
	return &this
}

// GetResponseCode returns the ResponseCode field value
func (o *VoidResponseObjectVoidResponse) GetResponseCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetResponseCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseCode, true
}

// SetResponseCode sets field value
func (o *VoidResponseObjectVoidResponse) SetResponseCode(v int32) {
	o.ResponseCode = v
}

// GetResponseActions returns the ResponseActions field value
func (o *VoidResponseObjectVoidResponse) GetResponseActions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResponseActions
}

// GetResponseActionsOk returns a tuple with the ResponseActions field value
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetResponseActionsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseActions, true
}

// SetResponseActions sets field value
func (o *VoidResponseObjectVoidResponse) SetResponseActions(v []string) {
	o.ResponseActions = v
}

// GetResponseMessage returns the ResponseMessage field value
func (o *VoidResponseObjectVoidResponse) GetResponseMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetResponseMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseMessage, true
}

// SetResponseMessage sets field value
func (o *VoidResponseObjectVoidResponse) SetResponseMessage(v string) {
	o.ResponseMessage = v
}

// GetCompanyIdentification returns the CompanyIdentification field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCompanyIdentification() string {
	if o == nil || o.CompanyIdentification == nil {
		var ret string
		return ret
	}
	return *o.CompanyIdentification
}

// GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCompanyIdentificationOk() (*string, bool) {
	if o == nil || o.CompanyIdentification == nil {
		return nil, false
	}
	return o.CompanyIdentification, true
}

// HasCompanyIdentification returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCompanyIdentification() bool {
	if o != nil && o.CompanyIdentification != nil {
		return true
	}

	return false
}

// SetCompanyIdentification gets a reference to the given string and assigns it to the CompanyIdentification field.
func (o *VoidResponseObjectVoidResponse) SetCompanyIdentification(v string) {
	o.CompanyIdentification = &v
}

// GetSystemIdentification returns the SystemIdentification field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetSystemIdentification() string {
	if o == nil || o.SystemIdentification == nil {
		var ret string
		return ret
	}
	return *o.SystemIdentification
}

// GetSystemIdentificationOk returns a tuple with the SystemIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetSystemIdentificationOk() (*string, bool) {
	if o == nil || o.SystemIdentification == nil {
		return nil, false
	}
	return o.SystemIdentification, true
}

// HasSystemIdentification returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasSystemIdentification() bool {
	if o != nil && o.SystemIdentification != nil {
		return true
	}

	return false
}

// SetSystemIdentification gets a reference to the given string and assigns it to the SystemIdentification field.
func (o *VoidResponseObjectVoidResponse) SetSystemIdentification(v string) {
	o.SystemIdentification = &v
}

// GetBranchIdentification returns the BranchIdentification field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetBranchIdentification() string {
	if o == nil || o.BranchIdentification == nil {
		var ret string
		return ret
	}
	return *o.BranchIdentification
}

// GetBranchIdentificationOk returns a tuple with the BranchIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetBranchIdentificationOk() (*string, bool) {
	if o == nil || o.BranchIdentification == nil {
		return nil, false
	}
	return o.BranchIdentification, true
}

// HasBranchIdentification returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasBranchIdentification() bool {
	if o != nil && o.BranchIdentification != nil {
		return true
	}

	return false
}

// SetBranchIdentification gets a reference to the given string and assigns it to the BranchIdentification field.
func (o *VoidResponseObjectVoidResponse) SetBranchIdentification(v string) {
	o.BranchIdentification = &v
}

// GetPOSIdentification returns the POSIdentification field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPOSIdentification() string {
	if o == nil || o.POSIdentification == nil {
		var ret string
		return ret
	}
	return *o.POSIdentification
}

// GetPOSIdentificationOk returns a tuple with the POSIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPOSIdentificationOk() (*string, bool) {
	if o == nil || o.POSIdentification == nil {
		return nil, false
	}
	return o.POSIdentification, true
}

// HasPOSIdentification returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPOSIdentification() bool {
	if o != nil && o.POSIdentification != nil {
		return true
	}

	return false
}

// SetPOSIdentification gets a reference to the given string and assigns it to the POSIdentification field.
func (o *VoidResponseObjectVoidResponse) SetPOSIdentification(v string) {
	o.POSIdentification = &v
}

// GetForeignResponseCode returns the ForeignResponseCode field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetForeignResponseCode() string {
	if o == nil || o.ForeignResponseCode == nil {
		var ret string
		return ret
	}
	return *o.ForeignResponseCode
}

// GetForeignResponseCodeOk returns a tuple with the ForeignResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetForeignResponseCodeOk() (*string, bool) {
	if o == nil || o.ForeignResponseCode == nil {
		return nil, false
	}
	return o.ForeignResponseCode, true
}

// HasForeignResponseCode returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasForeignResponseCode() bool {
	if o != nil && o.ForeignResponseCode != nil {
		return true
	}

	return false
}

// SetForeignResponseCode gets a reference to the given string and assigns it to the ForeignResponseCode field.
func (o *VoidResponseObjectVoidResponse) SetForeignResponseCode(v string) {
	o.ForeignResponseCode = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *VoidResponseObjectVoidResponse) SetRequestType(v string) {
	o.RequestType = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *VoidResponseObjectVoidResponse) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetSequence() string {
	if o == nil || o.Sequence == nil {
		var ret string
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetSequenceOk() (*string, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given string and assigns it to the Sequence field.
func (o *VoidResponseObjectVoidResponse) SetSequence(v string) {
	o.Sequence = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetSecurity() []SaleObjectSaleSecurity {
	if o == nil || o.Security == nil {
		var ret []SaleObjectSaleSecurity
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetSecurityOk() ([]SaleObjectSaleSecurity, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []SaleObjectSaleSecurity and assigns it to the Security field.
func (o *VoidResponseObjectVoidResponse) SetSecurity(v []SaleObjectSaleSecurity) {
	o.Security = v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetBlock() string {
	if o == nil || o.Block == nil {
		var ret string
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetBlockOk() (*string, bool) {
	if o == nil || o.Block == nil {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasBlock() bool {
	if o != nil && o.Block != nil {
		return true
	}

	return false
}

// SetBlock gets a reference to the given string and assigns it to the Block field.
func (o *VoidResponseObjectVoidResponse) SetBlock(v string) {
	o.Block = &v
}

// GetRequestKey returns the RequestKey field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetRequestKey() string {
	if o == nil || o.RequestKey == nil {
		var ret string
		return ret
	}
	return *o.RequestKey
}

// GetRequestKeyOk returns a tuple with the RequestKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetRequestKeyOk() (*string, bool) {
	if o == nil || o.RequestKey == nil {
		return nil, false
	}
	return o.RequestKey, true
}

// HasRequestKey returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasRequestKey() bool {
	if o != nil && o.RequestKey != nil {
		return true
	}

	return false
}

// SetRequestKey gets a reference to the given string and assigns it to the RequestKey field.
func (o *VoidResponseObjectVoidResponse) SetRequestKey(v string) {
	o.RequestKey = &v
}

// GetWorkingKeys returns the WorkingKeys field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetWorkingKeys() []SaleResponseObjectSaleResponseWorkingKeys {
	if o == nil || o.WorkingKeys == nil {
		var ret []SaleResponseObjectSaleResponseWorkingKeys
		return ret
	}
	return o.WorkingKeys
}

// GetWorkingKeysOk returns a tuple with the WorkingKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetWorkingKeysOk() ([]SaleResponseObjectSaleResponseWorkingKeys, bool) {
	if o == nil || o.WorkingKeys == nil {
		return nil, false
	}
	return o.WorkingKeys, true
}

// HasWorkingKeys returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasWorkingKeys() bool {
	if o != nil && o.WorkingKeys != nil {
		return true
	}

	return false
}

// SetWorkingKeys gets a reference to the given []SaleResponseObjectSaleResponseWorkingKeys and assigns it to the WorkingKeys field.
func (o *VoidResponseObjectVoidResponse) SetWorkingKeys(v []SaleResponseObjectSaleResponseWorkingKeys) {
	o.WorkingKeys = v
}

// GetRequiredInformation returns the RequiredInformation field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetRequiredInformation() []SaleResponseObjectSaleResponseRequiredInformation {
	if o == nil || o.RequiredInformation == nil {
		var ret []SaleResponseObjectSaleResponseRequiredInformation
		return ret
	}
	return o.RequiredInformation
}

// GetRequiredInformationOk returns a tuple with the RequiredInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetRequiredInformationOk() ([]SaleResponseObjectSaleResponseRequiredInformation, bool) {
	if o == nil || o.RequiredInformation == nil {
		return nil, false
	}
	return o.RequiredInformation, true
}

// HasRequiredInformation returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasRequiredInformation() bool {
	if o != nil && o.RequiredInformation != nil {
		return true
	}

	return false
}

// SetRequiredInformation gets a reference to the given []SaleResponseObjectSaleResponseRequiredInformation and assigns it to the RequiredInformation field.
func (o *VoidResponseObjectVoidResponse) SetRequiredInformation(v []SaleResponseObjectSaleResponseRequiredInformation) {
	o.RequiredInformation = v
}

// GetAnswerType returns the AnswerType field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAnswerType() string {
	if o == nil || o.AnswerType == nil {
		var ret string
		return ret
	}
	return *o.AnswerType
}

// GetAnswerTypeOk returns a tuple with the AnswerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAnswerTypeOk() (*string, bool) {
	if o == nil || o.AnswerType == nil {
		return nil, false
	}
	return o.AnswerType, true
}

// HasAnswerType returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAnswerType() bool {
	if o != nil && o.AnswerType != nil {
		return true
	}

	return false
}

// SetAnswerType gets a reference to the given string and assigns it to the AnswerType field.
func (o *VoidResponseObjectVoidResponse) SetAnswerType(v string) {
	o.AnswerType = &v
}

// GetAnswerKey returns the AnswerKey field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAnswerKey() string {
	if o == nil || o.AnswerKey == nil {
		var ret string
		return ret
	}
	return *o.AnswerKey
}

// GetAnswerKeyOk returns a tuple with the AnswerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAnswerKeyOk() (*string, bool) {
	if o == nil || o.AnswerKey == nil {
		return nil, false
	}
	return o.AnswerKey, true
}

// HasAnswerKey returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAnswerKey() bool {
	if o != nil && o.AnswerKey != nil {
		return true
	}

	return false
}

// SetAnswerKey gets a reference to the given string and assigns it to the AnswerKey field.
func (o *VoidResponseObjectVoidResponse) SetAnswerKey(v string) {
	o.AnswerKey = &v
}

// GetPublicAnswerKey returns the PublicAnswerKey field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPublicAnswerKey() string {
	if o == nil || o.PublicAnswerKey == nil {
		var ret string
		return ret
	}
	return *o.PublicAnswerKey
}

// GetPublicAnswerKeyOk returns a tuple with the PublicAnswerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPublicAnswerKeyOk() (*string, bool) {
	if o == nil || o.PublicAnswerKey == nil {
		return nil, false
	}
	return o.PublicAnswerKey, true
}

// HasPublicAnswerKey returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPublicAnswerKey() bool {
	if o != nil && o.PublicAnswerKey != nil {
		return true
	}

	return false
}

// SetPublicAnswerKey gets a reference to the given string and assigns it to the PublicAnswerKey field.
func (o *VoidResponseObjectVoidResponse) SetPublicAnswerKey(v string) {
	o.PublicAnswerKey = &v
}

// GetWasReversePrevious returns the WasReversePrevious field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetWasReversePrevious() int32 {
	if o == nil || o.WasReversePrevious == nil {
		var ret int32
		return ret
	}
	return *o.WasReversePrevious
}

// GetWasReversePreviousOk returns a tuple with the WasReversePrevious field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetWasReversePreviousOk() (*int32, bool) {
	if o == nil || o.WasReversePrevious == nil {
		return nil, false
	}
	return o.WasReversePrevious, true
}

// HasWasReversePrevious returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasWasReversePrevious() bool {
	if o != nil && o.WasReversePrevious != nil {
		return true
	}

	return false
}

// SetWasReversePrevious gets a reference to the given int32 and assigns it to the WasReversePrevious field.
func (o *VoidResponseObjectVoidResponse) SetWasReversePrevious(v int32) {
	o.WasReversePrevious = &v
}

// GetReversedAnswerKey returns the ReversedAnswerKey field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetReversedAnswerKey() string {
	if o == nil || o.ReversedAnswerKey == nil {
		var ret string
		return ret
	}
	return *o.ReversedAnswerKey
}

// GetReversedAnswerKeyOk returns a tuple with the ReversedAnswerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetReversedAnswerKeyOk() (*string, bool) {
	if o == nil || o.ReversedAnswerKey == nil {
		return nil, false
	}
	return o.ReversedAnswerKey, true
}

// HasReversedAnswerKey returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasReversedAnswerKey() bool {
	if o != nil && o.ReversedAnswerKey != nil {
		return true
	}

	return false
}

// SetReversedAnswerKey gets a reference to the given string and assigns it to the ReversedAnswerKey field.
func (o *VoidResponseObjectVoidResponse) SetReversedAnswerKey(v string) {
	o.ReversedAnswerKey = &v
}

// GetReversedSequence returns the ReversedSequence field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetReversedSequence() string {
	if o == nil || o.ReversedSequence == nil {
		var ret string
		return ret
	}
	return *o.ReversedSequence
}

// GetReversedSequenceOk returns a tuple with the ReversedSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetReversedSequenceOk() (*string, bool) {
	if o == nil || o.ReversedSequence == nil {
		return nil, false
	}
	return o.ReversedSequence, true
}

// HasReversedSequence returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasReversedSequence() bool {
	if o != nil && o.ReversedSequence != nil {
		return true
	}

	return false
}

// SetReversedSequence gets a reference to the given string and assigns it to the ReversedSequence field.
func (o *VoidResponseObjectVoidResponse) SetReversedSequence(v string) {
	o.ReversedSequence = &v
}

// GetCommittedBlock returns the CommittedBlock field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCommittedBlock() string {
	if o == nil || o.CommittedBlock == nil {
		var ret string
		return ret
	}
	return *o.CommittedBlock
}

// GetCommittedBlockOk returns a tuple with the CommittedBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCommittedBlockOk() (*string, bool) {
	if o == nil || o.CommittedBlock == nil {
		return nil, false
	}
	return o.CommittedBlock, true
}

// HasCommittedBlock returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCommittedBlock() bool {
	if o != nil && o.CommittedBlock != nil {
		return true
	}

	return false
}

// SetCommittedBlock gets a reference to the given string and assigns it to the CommittedBlock field.
func (o *VoidResponseObjectVoidResponse) SetCommittedBlock(v string) {
	o.CommittedBlock = &v
}

// GetReversedBlock returns the ReversedBlock field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetReversedBlock() string {
	if o == nil || o.ReversedBlock == nil {
		var ret string
		return ret
	}
	return *o.ReversedBlock
}

// GetReversedBlockOk returns a tuple with the ReversedBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetReversedBlockOk() (*string, bool) {
	if o == nil || o.ReversedBlock == nil {
		return nil, false
	}
	return o.ReversedBlock, true
}

// HasReversedBlock returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasReversedBlock() bool {
	if o != nil && o.ReversedBlock != nil {
		return true
	}

	return false
}

// SetReversedBlock gets a reference to the given string and assigns it to the ReversedBlock field.
func (o *VoidResponseObjectVoidResponse) SetReversedBlock(v string) {
	o.ReversedBlock = &v
}

// GetMessageID returns the MessageID field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetMessageID() string {
	if o == nil || o.MessageID == nil {
		var ret string
		return ret
	}
	return *o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetMessageIDOk() (*string, bool) {
	if o == nil || o.MessageID == nil {
		return nil, false
	}
	return o.MessageID, true
}

// HasMessageID returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasMessageID() bool {
	if o != nil && o.MessageID != nil {
		return true
	}

	return false
}

// SetMessageID gets a reference to the given string and assigns it to the MessageID field.
func (o *VoidResponseObjectVoidResponse) SetMessageID(v string) {
	o.MessageID = &v
}

// GetServerAddress returns the ServerAddress field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetServerAddress() string {
	if o == nil || o.ServerAddress == nil {
		var ret string
		return ret
	}
	return *o.ServerAddress
}

// GetServerAddressOk returns a tuple with the ServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetServerAddressOk() (*string, bool) {
	if o == nil || o.ServerAddress == nil {
		return nil, false
	}
	return o.ServerAddress, true
}

// HasServerAddress returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasServerAddress() bool {
	if o != nil && o.ServerAddress != nil {
		return true
	}

	return false
}

// SetServerAddress gets a reference to the given string and assigns it to the ServerAddress field.
func (o *VoidResponseObjectVoidResponse) SetServerAddress(v string) {
	o.ServerAddress = &v
}

// GetServerInstance returns the ServerInstance field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetServerInstance() string {
	if o == nil || o.ServerInstance == nil {
		var ret string
		return ret
	}
	return *o.ServerInstance
}

// GetServerInstanceOk returns a tuple with the ServerInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetServerInstanceOk() (*string, bool) {
	if o == nil || o.ServerInstance == nil {
		return nil, false
	}
	return o.ServerInstance, true
}

// HasServerInstance returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasServerInstance() bool {
	if o != nil && o.ServerInstance != nil {
		return true
	}

	return false
}

// SetServerInstance gets a reference to the given string and assigns it to the ServerInstance field.
func (o *VoidResponseObjectVoidResponse) SetServerInstance(v string) {
	o.ServerInstance = &v
}

// GetServerNodeName returns the ServerNodeName field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetServerNodeName() string {
	if o == nil || o.ServerNodeName == nil {
		var ret string
		return ret
	}
	return *o.ServerNodeName
}

// GetServerNodeNameOk returns a tuple with the ServerNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetServerNodeNameOk() (*string, bool) {
	if o == nil || o.ServerNodeName == nil {
		return nil, false
	}
	return o.ServerNodeName, true
}

// HasServerNodeName returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasServerNodeName() bool {
	if o != nil && o.ServerNodeName != nil {
		return true
	}

	return false
}

// SetServerNodeName gets a reference to the given string and assigns it to the ServerNodeName field.
func (o *VoidResponseObjectVoidResponse) SetServerNodeName(v string) {
	o.ServerNodeName = &v
}

// GetAdapterInputVersion returns the AdapterInputVersion field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAdapterInputVersion() string {
	if o == nil || o.AdapterInputVersion == nil {
		var ret string
		return ret
	}
	return *o.AdapterInputVersion
}

// GetAdapterInputVersionOk returns a tuple with the AdapterInputVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAdapterInputVersionOk() (*string, bool) {
	if o == nil || o.AdapterInputVersion == nil {
		return nil, false
	}
	return o.AdapterInputVersion, true
}

// HasAdapterInputVersion returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAdapterInputVersion() bool {
	if o != nil && o.AdapterInputVersion != nil {
		return true
	}

	return false
}

// SetAdapterInputVersion gets a reference to the given string and assigns it to the AdapterInputVersion field.
func (o *VoidResponseObjectVoidResponse) SetAdapterInputVersion(v string) {
	o.AdapterInputVersion = &v
}

// GetAdapterInputAddress returns the AdapterInputAddress field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAdapterInputAddress() string {
	if o == nil || o.AdapterInputAddress == nil {
		var ret string
		return ret
	}
	return *o.AdapterInputAddress
}

// GetAdapterInputAddressOk returns a tuple with the AdapterInputAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAdapterInputAddressOk() (*string, bool) {
	if o == nil || o.AdapterInputAddress == nil {
		return nil, false
	}
	return o.AdapterInputAddress, true
}

// HasAdapterInputAddress returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAdapterInputAddress() bool {
	if o != nil && o.AdapterInputAddress != nil {
		return true
	}

	return false
}

// SetAdapterInputAddress gets a reference to the given string and assigns it to the AdapterInputAddress field.
func (o *VoidResponseObjectVoidResponse) SetAdapterInputAddress(v string) {
	o.AdapterInputAddress = &v
}

// GetAdapterInputNodeName returns the AdapterInputNodeName field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAdapterInputNodeName() string {
	if o == nil || o.AdapterInputNodeName == nil {
		var ret string
		return ret
	}
	return *o.AdapterInputNodeName
}

// GetAdapterInputNodeNameOk returns a tuple with the AdapterInputNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAdapterInputNodeNameOk() (*string, bool) {
	if o == nil || o.AdapterInputNodeName == nil {
		return nil, false
	}
	return o.AdapterInputNodeName, true
}

// HasAdapterInputNodeName returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAdapterInputNodeName() bool {
	if o != nil && o.AdapterInputNodeName != nil {
		return true
	}

	return false
}

// SetAdapterInputNodeName gets a reference to the given string and assigns it to the AdapterInputNodeName field.
func (o *VoidResponseObjectVoidResponse) SetAdapterInputNodeName(v string) {
	o.AdapterInputNodeName = &v
}

// GetAdapterOutputVersion returns the AdapterOutputVersion field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAdapterOutputVersion() string {
	if o == nil || o.AdapterOutputVersion == nil {
		var ret string
		return ret
	}
	return *o.AdapterOutputVersion
}

// GetAdapterOutputVersionOk returns a tuple with the AdapterOutputVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAdapterOutputVersionOk() (*string, bool) {
	if o == nil || o.AdapterOutputVersion == nil {
		return nil, false
	}
	return o.AdapterOutputVersion, true
}

// HasAdapterOutputVersion returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAdapterOutputVersion() bool {
	if o != nil && o.AdapterOutputVersion != nil {
		return true
	}

	return false
}

// SetAdapterOutputVersion gets a reference to the given string and assigns it to the AdapterOutputVersion field.
func (o *VoidResponseObjectVoidResponse) SetAdapterOutputVersion(v string) {
	o.AdapterOutputVersion = &v
}

// GetAdapterOutputAddress returns the AdapterOutputAddress field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAdapterOutputAddress() string {
	if o == nil || o.AdapterOutputAddress == nil {
		var ret string
		return ret
	}
	return *o.AdapterOutputAddress
}

// GetAdapterOutputAddressOk returns a tuple with the AdapterOutputAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAdapterOutputAddressOk() (*string, bool) {
	if o == nil || o.AdapterOutputAddress == nil {
		return nil, false
	}
	return o.AdapterOutputAddress, true
}

// HasAdapterOutputAddress returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAdapterOutputAddress() bool {
	if o != nil && o.AdapterOutputAddress != nil {
		return true
	}

	return false
}

// SetAdapterOutputAddress gets a reference to the given string and assigns it to the AdapterOutputAddress field.
func (o *VoidResponseObjectVoidResponse) SetAdapterOutputAddress(v string) {
	o.AdapterOutputAddress = &v
}

// GetAdapterOutputNodeName returns the AdapterOutputNodeName field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAdapterOutputNodeName() string {
	if o == nil || o.AdapterOutputNodeName == nil {
		var ret string
		return ret
	}
	return *o.AdapterOutputNodeName
}

// GetAdapterOutputNodeNameOk returns a tuple with the AdapterOutputNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAdapterOutputNodeNameOk() (*string, bool) {
	if o == nil || o.AdapterOutputNodeName == nil {
		return nil, false
	}
	return o.AdapterOutputNodeName, true
}

// HasAdapterOutputNodeName returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAdapterOutputNodeName() bool {
	if o != nil && o.AdapterOutputNodeName != nil {
		return true
	}

	return false
}

// SetAdapterOutputNodeName gets a reference to the given string and assigns it to the AdapterOutputNodeName field.
func (o *VoidResponseObjectVoidResponse) SetAdapterOutputNodeName(v string) {
	o.AdapterOutputNodeName = &v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAdditionalInformation() []SaleResponseObjectSaleResponseAdditionalInformation {
	if o == nil || o.AdditionalInformation == nil {
		var ret []SaleResponseObjectSaleResponseAdditionalInformation
		return ret
	}
	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAdditionalInformationOk() ([]SaleResponseObjectSaleResponseAdditionalInformation, bool) {
	if o == nil || o.AdditionalInformation == nil {
		return nil, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAdditionalInformation() bool {
	if o != nil && o.AdditionalInformation != nil {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given []SaleResponseObjectSaleResponseAdditionalInformation and assigns it to the AdditionalInformation field.
func (o *VoidResponseObjectVoidResponse) SetAdditionalInformation(v []SaleResponseObjectSaleResponseAdditionalInformation) {
	o.AdditionalInformation = v
}

// GetPrinterResponseMessage returns the PrinterResponseMessage field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPrinterResponseMessage() []string {
	if o == nil || o.PrinterResponseMessage == nil {
		var ret []string
		return ret
	}
	return o.PrinterResponseMessage
}

// GetPrinterResponseMessageOk returns a tuple with the PrinterResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPrinterResponseMessageOk() ([]string, bool) {
	if o == nil || o.PrinterResponseMessage == nil {
		return nil, false
	}
	return o.PrinterResponseMessage, true
}

// HasPrinterResponseMessage returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPrinterResponseMessage() bool {
	if o != nil && o.PrinterResponseMessage != nil {
		return true
	}

	return false
}

// SetPrinterResponseMessage gets a reference to the given []string and assigns it to the PrinterResponseMessage field.
func (o *VoidResponseObjectVoidResponse) SetPrinterResponseMessage(v []string) {
	o.PrinterResponseMessage = v
}

// GetDisplayResponseMessage returns the DisplayResponseMessage field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetDisplayResponseMessage() []string {
	if o == nil || o.DisplayResponseMessage == nil {
		var ret []string
		return ret
	}
	return o.DisplayResponseMessage
}

// GetDisplayResponseMessageOk returns a tuple with the DisplayResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetDisplayResponseMessageOk() ([]string, bool) {
	if o == nil || o.DisplayResponseMessage == nil {
		return nil, false
	}
	return o.DisplayResponseMessage, true
}

// HasDisplayResponseMessage returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasDisplayResponseMessage() bool {
	if o != nil && o.DisplayResponseMessage != nil {
		return true
	}

	return false
}

// SetDisplayResponseMessage gets a reference to the given []string and assigns it to the DisplayResponseMessage field.
func (o *VoidResponseObjectVoidResponse) SetDisplayResponseMessage(v []string) {
	o.DisplayResponseMessage = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *VoidResponseObjectVoidResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencyDescription returns the CurrencyDescription field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCurrencyDescription() string {
	if o == nil || o.CurrencyDescription == nil {
		var ret string
		return ret
	}
	return *o.CurrencyDescription
}

// GetCurrencyDescriptionOk returns a tuple with the CurrencyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCurrencyDescriptionOk() (*string, bool) {
	if o == nil || o.CurrencyDescription == nil {
		return nil, false
	}
	return o.CurrencyDescription, true
}

// HasCurrencyDescription returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCurrencyDescription() bool {
	if o != nil && o.CurrencyDescription != nil {
		return true
	}

	return false
}

// SetCurrencyDescription gets a reference to the given string and assigns it to the CurrencyDescription field.
func (o *VoidResponseObjectVoidResponse) SetCurrencyDescription(v string) {
	o.CurrencyDescription = &v
}

// GetTransactionDescription returns the TransactionDescription field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetTransactionDescription() string {
	if o == nil || o.TransactionDescription == nil {
		var ret string
		return ret
	}
	return *o.TransactionDescription
}

// GetTransactionDescriptionOk returns a tuple with the TransactionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetTransactionDescriptionOk() (*string, bool) {
	if o == nil || o.TransactionDescription == nil {
		return nil, false
	}
	return o.TransactionDescription, true
}

// HasTransactionDescription returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasTransactionDescription() bool {
	if o != nil && o.TransactionDescription != nil {
		return true
	}

	return false
}

// SetTransactionDescription gets a reference to the given string and assigns it to the TransactionDescription field.
func (o *VoidResponseObjectVoidResponse) SetTransactionDescription(v string) {
	o.TransactionDescription = &v
}

// GetTransactionResolutionMode returns the TransactionResolutionMode field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetTransactionResolutionMode() string {
	if o == nil || o.TransactionResolutionMode == nil {
		var ret string
		return ret
	}
	return *o.TransactionResolutionMode
}

// GetTransactionResolutionModeOk returns a tuple with the TransactionResolutionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetTransactionResolutionModeOk() (*string, bool) {
	if o == nil || o.TransactionResolutionMode == nil {
		return nil, false
	}
	return o.TransactionResolutionMode, true
}

// HasTransactionResolutionMode returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasTransactionResolutionMode() bool {
	if o != nil && o.TransactionResolutionMode != nil {
		return true
	}

	return false
}

// SetTransactionResolutionMode gets a reference to the given string and assigns it to the TransactionResolutionMode field.
func (o *VoidResponseObjectVoidResponse) SetTransactionResolutionMode(v string) {
	o.TransactionResolutionMode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *VoidResponseObjectVoidResponse) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetFacilityPayments returns the FacilityPayments field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetFacilityPayments() int32 {
	if o == nil || o.FacilityPayments == nil {
		var ret int32
		return ret
	}
	return *o.FacilityPayments
}

// GetFacilityPaymentsOk returns a tuple with the FacilityPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetFacilityPaymentsOk() (*int32, bool) {
	if o == nil || o.FacilityPayments == nil {
		return nil, false
	}
	return o.FacilityPayments, true
}

// HasFacilityPayments returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasFacilityPayments() bool {
	if o != nil && o.FacilityPayments != nil {
		return true
	}

	return false
}

// SetFacilityPayments gets a reference to the given int32 and assigns it to the FacilityPayments field.
func (o *VoidResponseObjectVoidResponse) SetFacilityPayments(v int32) {
	o.FacilityPayments = &v
}

// GetFacilityType returns the FacilityType field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetFacilityType() int32 {
	if o == nil || o.FacilityType == nil {
		var ret int32
		return ret
	}
	return *o.FacilityType
}

// GetFacilityTypeOk returns a tuple with the FacilityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetFacilityTypeOk() (*int32, bool) {
	if o == nil || o.FacilityType == nil {
		return nil, false
	}
	return o.FacilityType, true
}

// HasFacilityType returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasFacilityType() bool {
	if o != nil && o.FacilityType != nil {
		return true
	}

	return false
}

// SetFacilityType gets a reference to the given int32 and assigns it to the FacilityType field.
func (o *VoidResponseObjectVoidResponse) SetFacilityType(v int32) {
	o.FacilityType = &v
}

// GetValueFacilityPayments returns the ValueFacilityPayments field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetValueFacilityPayments() float32 {
	if o == nil || o.ValueFacilityPayments == nil {
		var ret float32
		return ret
	}
	return *o.ValueFacilityPayments
}

// GetValueFacilityPaymentsOk returns a tuple with the ValueFacilityPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetValueFacilityPaymentsOk() (*float32, bool) {
	if o == nil || o.ValueFacilityPayments == nil {
		return nil, false
	}
	return o.ValueFacilityPayments, true
}

// HasValueFacilityPayments returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasValueFacilityPayments() bool {
	if o != nil && o.ValueFacilityPayments != nil {
		return true
	}

	return false
}

// SetValueFacilityPayments gets a reference to the given float32 and assigns it to the ValueFacilityPayments field.
func (o *VoidResponseObjectVoidResponse) SetValueFacilityPayments(v float32) {
	o.ValueFacilityPayments = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *VoidResponseObjectVoidResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetAlternativeAmount returns the AlternativeAmount field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAlternativeAmount() float32 {
	if o == nil || o.AlternativeAmount == nil {
		var ret float32
		return ret
	}
	return *o.AlternativeAmount
}

// GetAlternativeAmountOk returns a tuple with the AlternativeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAlternativeAmountOk() (*float32, bool) {
	if o == nil || o.AlternativeAmount == nil {
		return nil, false
	}
	return o.AlternativeAmount, true
}

// HasAlternativeAmount returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAlternativeAmount() bool {
	if o != nil && o.AlternativeAmount != nil {
		return true
	}

	return false
}

// SetAlternativeAmount gets a reference to the given float32 and assigns it to the AlternativeAmount field.
func (o *VoidResponseObjectVoidResponse) SetAlternativeAmount(v float32) {
	o.AlternativeAmount = &v
}

// GetAmountCharged returns the AmountCharged field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAmountCharged() float32 {
	if o == nil || o.AmountCharged == nil {
		var ret float32
		return ret
	}
	return *o.AmountCharged
}

// GetAmountChargedOk returns a tuple with the AmountCharged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAmountChargedOk() (*float32, bool) {
	if o == nil || o.AmountCharged == nil {
		return nil, false
	}
	return o.AmountCharged, true
}

// HasAmountCharged returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAmountCharged() bool {
	if o != nil && o.AmountCharged != nil {
		return true
	}

	return false
}

// SetAmountCharged gets a reference to the given float32 and assigns it to the AmountCharged field.
func (o *VoidResponseObjectVoidResponse) SetAmountCharged(v float32) {
	o.AmountCharged = &v
}

// GetAmountToApply returns the AmountToApply field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAmountToApply() float32 {
	if o == nil || o.AmountToApply == nil {
		var ret float32
		return ret
	}
	return *o.AmountToApply
}

// GetAmountToApplyOk returns a tuple with the AmountToApply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAmountToApplyOk() (*float32, bool) {
	if o == nil || o.AmountToApply == nil {
		return nil, false
	}
	return o.AmountToApply, true
}

// HasAmountToApply returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAmountToApply() bool {
	if o != nil && o.AmountToApply != nil {
		return true
	}

	return false
}

// SetAmountToApply gets a reference to the given float32 and assigns it to the AmountToApply field.
func (o *VoidResponseObjectVoidResponse) SetAmountToApply(v float32) {
	o.AmountToApply = &v
}

// GetCashbackAmount returns the CashbackAmount field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCashbackAmount() float32 {
	if o == nil || o.CashbackAmount == nil {
		var ret float32
		return ret
	}
	return *o.CashbackAmount
}

// GetCashbackAmountOk returns a tuple with the CashbackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCashbackAmountOk() (*float32, bool) {
	if o == nil || o.CashbackAmount == nil {
		return nil, false
	}
	return o.CashbackAmount, true
}

// HasCashbackAmount returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCashbackAmount() bool {
	if o != nil && o.CashbackAmount != nil {
		return true
	}

	return false
}

// SetCashbackAmount gets a reference to the given float32 and assigns it to the CashbackAmount field.
func (o *VoidResponseObjectVoidResponse) SetCashbackAmount(v float32) {
	o.CashbackAmount = &v
}

// GetTipAmount returns the TipAmount field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetTipAmount() float32 {
	if o == nil || o.TipAmount == nil {
		var ret float32
		return ret
	}
	return *o.TipAmount
}

// GetTipAmountOk returns a tuple with the TipAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetTipAmountOk() (*float32, bool) {
	if o == nil || o.TipAmount == nil {
		return nil, false
	}
	return o.TipAmount, true
}

// HasTipAmount returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasTipAmount() bool {
	if o != nil && o.TipAmount != nil {
		return true
	}

	return false
}

// SetTipAmount gets a reference to the given float32 and assigns it to the TipAmount field.
func (o *VoidResponseObjectVoidResponse) SetTipAmount(v float32) {
	o.TipAmount = &v
}

// GetRequestAmount returns the RequestAmount field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetRequestAmount() float32 {
	if o == nil || o.RequestAmount == nil {
		var ret float32
		return ret
	}
	return *o.RequestAmount
}

// GetRequestAmountOk returns a tuple with the RequestAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetRequestAmountOk() (*float32, bool) {
	if o == nil || o.RequestAmount == nil {
		return nil, false
	}
	return o.RequestAmount, true
}

// HasRequestAmount returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasRequestAmount() bool {
	if o != nil && o.RequestAmount != nil {
		return true
	}

	return false
}

// SetRequestAmount gets a reference to the given float32 and assigns it to the RequestAmount field.
func (o *VoidResponseObjectVoidResponse) SetRequestAmount(v float32) {
	o.RequestAmount = &v
}

// GetPromotedAmount returns the PromotedAmount field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPromotedAmount() float32 {
	if o == nil || o.PromotedAmount == nil {
		var ret float32
		return ret
	}
	return *o.PromotedAmount
}

// GetPromotedAmountOk returns a tuple with the PromotedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPromotedAmountOk() (*float32, bool) {
	if o == nil || o.PromotedAmount == nil {
		return nil, false
	}
	return o.PromotedAmount, true
}

// HasPromotedAmount returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPromotedAmount() bool {
	if o != nil && o.PromotedAmount != nil {
		return true
	}

	return false
}

// SetPromotedAmount gets a reference to the given float32 and assigns it to the PromotedAmount field.
func (o *VoidResponseObjectVoidResponse) SetPromotedAmount(v float32) {
	o.PromotedAmount = &v
}

// GetCredentialToken returns the CredentialToken field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCredentialToken() string {
	if o == nil || o.CredentialToken == nil {
		var ret string
		return ret
	}
	return *o.CredentialToken
}

// GetCredentialTokenOk returns a tuple with the CredentialToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCredentialTokenOk() (*string, bool) {
	if o == nil || o.CredentialToken == nil {
		return nil, false
	}
	return o.CredentialToken, true
}

// HasCredentialToken returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCredentialToken() bool {
	if o != nil && o.CredentialToken != nil {
		return true
	}

	return false
}

// SetCredentialToken gets a reference to the given string and assigns it to the CredentialToken field.
func (o *VoidResponseObjectVoidResponse) SetCredentialToken(v string) {
	o.CredentialToken = &v
}

// GetCredentialIssuerToken returns the CredentialIssuerToken field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCredentialIssuerToken() string {
	if o == nil || o.CredentialIssuerToken == nil {
		var ret string
		return ret
	}
	return *o.CredentialIssuerToken
}

// GetCredentialIssuerTokenOk returns a tuple with the CredentialIssuerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCredentialIssuerTokenOk() (*string, bool) {
	if o == nil || o.CredentialIssuerToken == nil {
		return nil, false
	}
	return o.CredentialIssuerToken, true
}

// HasCredentialIssuerToken returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCredentialIssuerToken() bool {
	if o != nil && o.CredentialIssuerToken != nil {
		return true
	}

	return false
}

// SetCredentialIssuerToken gets a reference to the given string and assigns it to the CredentialIssuerToken field.
func (o *VoidResponseObjectVoidResponse) SetCredentialIssuerToken(v string) {
	o.CredentialIssuerToken = &v
}

// GetInputTokens returns the InputTokens field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetInputTokens() []SaleObjectSaleInputTokens {
	if o == nil || o.InputTokens == nil {
		var ret []SaleObjectSaleInputTokens
		return ret
	}
	return o.InputTokens
}

// GetInputTokensOk returns a tuple with the InputTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetInputTokensOk() ([]SaleObjectSaleInputTokens, bool) {
	if o == nil || o.InputTokens == nil {
		return nil, false
	}
	return o.InputTokens, true
}

// HasInputTokens returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasInputTokens() bool {
	if o != nil && o.InputTokens != nil {
		return true
	}

	return false
}

// SetInputTokens gets a reference to the given []SaleObjectSaleInputTokens and assigns it to the InputTokens field.
func (o *VoidResponseObjectVoidResponse) SetInputTokens(v []SaleObjectSaleInputTokens) {
	o.InputTokens = v
}

// GetOutputTokens returns the OutputTokens field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetOutputTokens() []SaleObjectSaleInputTokens {
	if o == nil || o.OutputTokens == nil {
		var ret []SaleObjectSaleInputTokens
		return ret
	}
	return o.OutputTokens
}

// GetOutputTokensOk returns a tuple with the OutputTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetOutputTokensOk() ([]SaleObjectSaleInputTokens, bool) {
	if o == nil || o.OutputTokens == nil {
		return nil, false
	}
	return o.OutputTokens, true
}

// HasOutputTokens returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasOutputTokens() bool {
	if o != nil && o.OutputTokens != nil {
		return true
	}

	return false
}

// SetOutputTokens gets a reference to the given []SaleObjectSaleInputTokens and assigns it to the OutputTokens field.
func (o *VoidResponseObjectVoidResponse) SetOutputTokens(v []SaleObjectSaleInputTokens) {
	o.OutputTokens = v
}

// GetTaxFinancialCostAmount returns the TaxFinancialCostAmount field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetTaxFinancialCostAmount() float32 {
	if o == nil || o.TaxFinancialCostAmount == nil {
		var ret float32
		return ret
	}
	return *o.TaxFinancialCostAmount
}

// GetTaxFinancialCostAmountOk returns a tuple with the TaxFinancialCostAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetTaxFinancialCostAmountOk() (*float32, bool) {
	if o == nil || o.TaxFinancialCostAmount == nil {
		return nil, false
	}
	return o.TaxFinancialCostAmount, true
}

// HasTaxFinancialCostAmount returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasTaxFinancialCostAmount() bool {
	if o != nil && o.TaxFinancialCostAmount != nil {
		return true
	}

	return false
}

// SetTaxFinancialCostAmount gets a reference to the given float32 and assigns it to the TaxFinancialCostAmount field.
func (o *VoidResponseObjectVoidResponse) SetTaxFinancialCostAmount(v float32) {
	o.TaxFinancialCostAmount = &v
}

// GetTaxFinancialCostPercentage returns the TaxFinancialCostPercentage field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetTaxFinancialCostPercentage() float32 {
	if o == nil || o.TaxFinancialCostPercentage == nil {
		var ret float32
		return ret
	}
	return *o.TaxFinancialCostPercentage
}

// GetTaxFinancialCostPercentageOk returns a tuple with the TaxFinancialCostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetTaxFinancialCostPercentageOk() (*float32, bool) {
	if o == nil || o.TaxFinancialCostPercentage == nil {
		return nil, false
	}
	return o.TaxFinancialCostPercentage, true
}

// HasTaxFinancialCostPercentage returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasTaxFinancialCostPercentage() bool {
	if o != nil && o.TaxFinancialCostPercentage != nil {
		return true
	}

	return false
}

// SetTaxFinancialCostPercentage gets a reference to the given float32 and assigns it to the TaxFinancialCostPercentage field.
func (o *VoidResponseObjectVoidResponse) SetTaxFinancialCostPercentage(v float32) {
	o.TaxFinancialCostPercentage = &v
}

// GetFinancialCostAmount returns the FinancialCostAmount field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetFinancialCostAmount() float32 {
	if o == nil || o.FinancialCostAmount == nil {
		var ret float32
		return ret
	}
	return *o.FinancialCostAmount
}

// GetFinancialCostAmountOk returns a tuple with the FinancialCostAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetFinancialCostAmountOk() (*float32, bool) {
	if o == nil || o.FinancialCostAmount == nil {
		return nil, false
	}
	return o.FinancialCostAmount, true
}

// HasFinancialCostAmount returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasFinancialCostAmount() bool {
	if o != nil && o.FinancialCostAmount != nil {
		return true
	}

	return false
}

// SetFinancialCostAmount gets a reference to the given float32 and assigns it to the FinancialCostAmount field.
func (o *VoidResponseObjectVoidResponse) SetFinancialCostAmount(v float32) {
	o.FinancialCostAmount = &v
}

// GetFinancialCostPercentage returns the FinancialCostPercentage field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetFinancialCostPercentage() float32 {
	if o == nil || o.FinancialCostPercentage == nil {
		var ret float32
		return ret
	}
	return *o.FinancialCostPercentage
}

// GetFinancialCostPercentageOk returns a tuple with the FinancialCostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetFinancialCostPercentageOk() (*float32, bool) {
	if o == nil || o.FinancialCostPercentage == nil {
		return nil, false
	}
	return o.FinancialCostPercentage, true
}

// HasFinancialCostPercentage returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasFinancialCostPercentage() bool {
	if o != nil && o.FinancialCostPercentage != nil {
		return true
	}

	return false
}

// SetFinancialCostPercentage gets a reference to the given float32 and assigns it to the FinancialCostPercentage field.
func (o *VoidResponseObjectVoidResponse) SetFinancialCostPercentage(v float32) {
	o.FinancialCostPercentage = &v
}

// GetHostResultCode returns the HostResultCode field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetHostResultCode() int32 {
	if o == nil || o.HostResultCode == nil {
		var ret int32
		return ret
	}
	return *o.HostResultCode
}

// GetHostResultCodeOk returns a tuple with the HostResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetHostResultCodeOk() (*int32, bool) {
	if o == nil || o.HostResultCode == nil {
		return nil, false
	}
	return o.HostResultCode, true
}

// HasHostResultCode returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasHostResultCode() bool {
	if o != nil && o.HostResultCode != nil {
		return true
	}

	return false
}

// SetHostResultCode gets a reference to the given int32 and assigns it to the HostResultCode field.
func (o *VoidResponseObjectVoidResponse) SetHostResultCode(v int32) {
	o.HostResultCode = &v
}

// GetHostMessage returns the HostMessage field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetHostMessage() string {
	if o == nil || o.HostMessage == nil {
		var ret string
		return ret
	}
	return *o.HostMessage
}

// GetHostMessageOk returns a tuple with the HostMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetHostMessageOk() (*string, bool) {
	if o == nil || o.HostMessage == nil {
		return nil, false
	}
	return o.HostMessage, true
}

// HasHostMessage returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasHostMessage() bool {
	if o != nil && o.HostMessage != nil {
		return true
	}

	return false
}

// SetHostMessage gets a reference to the given string and assigns it to the HostMessage field.
func (o *VoidResponseObjectVoidResponse) SetHostMessage(v string) {
	o.HostMessage = &v
}

// GetHostCode returns the HostCode field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetHostCode() string {
	if o == nil || o.HostCode == nil {
		var ret string
		return ret
	}
	return *o.HostCode
}

// GetHostCodeOk returns a tuple with the HostCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetHostCodeOk() (*string, bool) {
	if o == nil || o.HostCode == nil {
		return nil, false
	}
	return o.HostCode, true
}

// HasHostCode returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasHostCode() bool {
	if o != nil && o.HostCode != nil {
		return true
	}

	return false
}

// SetHostCode gets a reference to the given string and assigns it to the HostCode field.
func (o *VoidResponseObjectVoidResponse) SetHostCode(v string) {
	o.HostCode = &v
}

// GetHostID returns the HostID field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetHostID() int32 {
	if o == nil || o.HostID == nil {
		var ret int32
		return ret
	}
	return *o.HostID
}

// GetHostIDOk returns a tuple with the HostID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetHostIDOk() (*int32, bool) {
	if o == nil || o.HostID == nil {
		return nil, false
	}
	return o.HostID, true
}

// HasHostID returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasHostID() bool {
	if o != nil && o.HostID != nil {
		return true
	}

	return false
}

// SetHostID gets a reference to the given int32 and assigns it to the HostID field.
func (o *VoidResponseObjectVoidResponse) SetHostID(v int32) {
	o.HostID = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetUserID() string {
	if o == nil || o.UserID == nil {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetUserIDOk() (*string, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *VoidResponseObjectVoidResponse) SetUserID(v string) {
	o.UserID = &v
}

// GetIssuerName returns the IssuerName field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetIssuerName() string {
	if o == nil || o.IssuerName == nil {
		var ret string
		return ret
	}
	return *o.IssuerName
}

// GetIssuerNameOk returns a tuple with the IssuerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetIssuerNameOk() (*string, bool) {
	if o == nil || o.IssuerName == nil {
		return nil, false
	}
	return o.IssuerName, true
}

// HasIssuerName returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasIssuerName() bool {
	if o != nil && o.IssuerName != nil {
		return true
	}

	return false
}

// SetIssuerName gets a reference to the given string and assigns it to the IssuerName field.
func (o *VoidResponseObjectVoidResponse) SetIssuerName(v string) {
	o.IssuerName = &v
}

// GetHostDateTime returns the HostDateTime field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetHostDateTime() time.Time {
	if o == nil || o.HostDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.HostDateTime
}

// GetHostDateTimeOk returns a tuple with the HostDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetHostDateTimeOk() (*time.Time, bool) {
	if o == nil || o.HostDateTime == nil {
		return nil, false
	}
	return o.HostDateTime, true
}

// HasHostDateTime returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasHostDateTime() bool {
	if o != nil && o.HostDateTime != nil {
		return true
	}

	return false
}

// SetHostDateTime gets a reference to the given time.Time and assigns it to the HostDateTime field.
func (o *VoidResponseObjectVoidResponse) SetHostDateTime(v time.Time) {
	o.HostDateTime = &v
}

// GetTransmitionDateTime returns the TransmitionDateTime field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetTransmitionDateTime() time.Time {
	if o == nil || o.TransmitionDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.TransmitionDateTime
}

// GetTransmitionDateTimeOk returns a tuple with the TransmitionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetTransmitionDateTimeOk() (*time.Time, bool) {
	if o == nil || o.TransmitionDateTime == nil {
		return nil, false
	}
	return o.TransmitionDateTime, true
}

// HasTransmitionDateTime returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasTransmitionDateTime() bool {
	if o != nil && o.TransmitionDateTime != nil {
		return true
	}

	return false
}

// SetTransmitionDateTime gets a reference to the given time.Time and assigns it to the TransmitionDateTime field.
func (o *VoidResponseObjectVoidResponse) SetTransmitionDateTime(v time.Time) {
	o.TransmitionDateTime = &v
}

// GetAuthResultCode returns the AuthResultCode field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAuthResultCode() string {
	if o == nil || o.AuthResultCode == nil {
		var ret string
		return ret
	}
	return *o.AuthResultCode
}

// GetAuthResultCodeOk returns a tuple with the AuthResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAuthResultCodeOk() (*string, bool) {
	if o == nil || o.AuthResultCode == nil {
		return nil, false
	}
	return o.AuthResultCode, true
}

// HasAuthResultCode returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAuthResultCode() bool {
	if o != nil && o.AuthResultCode != nil {
		return true
	}

	return false
}

// SetAuthResultCode gets a reference to the given string and assigns it to the AuthResultCode field.
func (o *VoidResponseObjectVoidResponse) SetAuthResultCode(v string) {
	o.AuthResultCode = &v
}

// GetAuthHostProcessCode returns the AuthHostProcessCode field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAuthHostProcessCode() string {
	if o == nil || o.AuthHostProcessCode == nil {
		var ret string
		return ret
	}
	return *o.AuthHostProcessCode
}

// GetAuthHostProcessCodeOk returns a tuple with the AuthHostProcessCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAuthHostProcessCodeOk() (*string, bool) {
	if o == nil || o.AuthHostProcessCode == nil {
		return nil, false
	}
	return o.AuthHostProcessCode, true
}

// HasAuthHostProcessCode returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAuthHostProcessCode() bool {
	if o != nil && o.AuthHostProcessCode != nil {
		return true
	}

	return false
}

// SetAuthHostProcessCode gets a reference to the given string and assigns it to the AuthHostProcessCode field.
func (o *VoidResponseObjectVoidResponse) SetAuthHostProcessCode(v string) {
	o.AuthHostProcessCode = &v
}

// GetAuthHostMsgType returns the AuthHostMsgType field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAuthHostMsgType() string {
	if o == nil || o.AuthHostMsgType == nil {
		var ret string
		return ret
	}
	return *o.AuthHostMsgType
}

// GetAuthHostMsgTypeOk returns a tuple with the AuthHostMsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAuthHostMsgTypeOk() (*string, bool) {
	if o == nil || o.AuthHostMsgType == nil {
		return nil, false
	}
	return o.AuthHostMsgType, true
}

// HasAuthHostMsgType returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAuthHostMsgType() bool {
	if o != nil && o.AuthHostMsgType != nil {
		return true
	}

	return false
}

// SetAuthHostMsgType gets a reference to the given string and assigns it to the AuthHostMsgType field.
func (o *VoidResponseObjectVoidResponse) SetAuthHostMsgType(v string) {
	o.AuthHostMsgType = &v
}

// GetAuthHostMessage returns the AuthHostMessage field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAuthHostMessage() string {
	if o == nil || o.AuthHostMessage == nil {
		var ret string
		return ret
	}
	return *o.AuthHostMessage
}

// GetAuthHostMessageOk returns a tuple with the AuthHostMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAuthHostMessageOk() (*string, bool) {
	if o == nil || o.AuthHostMessage == nil {
		return nil, false
	}
	return o.AuthHostMessage, true
}

// HasAuthHostMessage returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAuthHostMessage() bool {
	if o != nil && o.AuthHostMessage != nil {
		return true
	}

	return false
}

// SetAuthHostMessage gets a reference to the given string and assigns it to the AuthHostMessage field.
func (o *VoidResponseObjectVoidResponse) SetAuthHostMessage(v string) {
	o.AuthHostMessage = &v
}

// GetHostMsgType returns the HostMsgType field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetHostMsgType() string {
	if o == nil || o.HostMsgType == nil {
		var ret string
		return ret
	}
	return *o.HostMsgType
}

// GetHostMsgTypeOk returns a tuple with the HostMsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetHostMsgTypeOk() (*string, bool) {
	if o == nil || o.HostMsgType == nil {
		return nil, false
	}
	return o.HostMsgType, true
}

// HasHostMsgType returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasHostMsgType() bool {
	if o != nil && o.HostMsgType != nil {
		return true
	}

	return false
}

// SetHostMsgType gets a reference to the given string and assigns it to the HostMsgType field.
func (o *VoidResponseObjectVoidResponse) SetHostMsgType(v string) {
	o.HostMsgType = &v
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAuthCode() string {
	if o == nil || o.AuthCode == nil {
		var ret string
		return ret
	}
	return *o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAuthCodeOk() (*string, bool) {
	if o == nil || o.AuthCode == nil {
		return nil, false
	}
	return o.AuthCode, true
}

// HasAuthCode returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAuthCode() bool {
	if o != nil && o.AuthCode != nil {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given string and assigns it to the AuthCode field.
func (o *VoidResponseObjectVoidResponse) SetAuthCode(v string) {
	o.AuthCode = &v
}

// GetAuthDateTime returns the AuthDateTime field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAuthDateTime() time.Time {
	if o == nil || o.AuthDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AuthDateTime
}

// GetAuthDateTimeOk returns a tuple with the AuthDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAuthDateTimeOk() (*time.Time, bool) {
	if o == nil || o.AuthDateTime == nil {
		return nil, false
	}
	return o.AuthDateTime, true
}

// HasAuthDateTime returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAuthDateTime() bool {
	if o != nil && o.AuthDateTime != nil {
		return true
	}

	return false
}

// SetAuthDateTime gets a reference to the given time.Time and assigns it to the AuthDateTime field.
func (o *VoidResponseObjectVoidResponse) SetAuthDateTime(v time.Time) {
	o.AuthDateTime = &v
}

// GetAuthTicket returns the AuthTicket field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAuthTicket() int32 {
	if o == nil || o.AuthTicket == nil {
		var ret int32
		return ret
	}
	return *o.AuthTicket
}

// GetAuthTicketOk returns a tuple with the AuthTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAuthTicketOk() (*int32, bool) {
	if o == nil || o.AuthTicket == nil {
		return nil, false
	}
	return o.AuthTicket, true
}

// HasAuthTicket returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAuthTicket() bool {
	if o != nil && o.AuthTicket != nil {
		return true
	}

	return false
}

// SetAuthTicket gets a reference to the given int32 and assigns it to the AuthTicket field.
func (o *VoidResponseObjectVoidResponse) SetAuthTicket(v int32) {
	o.AuthTicket = &v
}

// GetAuthRRN returns the AuthRRN field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetAuthRRN() string {
	if o == nil || o.AuthRRN == nil {
		var ret string
		return ret
	}
	return *o.AuthRRN
}

// GetAuthRRNOk returns a tuple with the AuthRRN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetAuthRRNOk() (*string, bool) {
	if o == nil || o.AuthRRN == nil {
		return nil, false
	}
	return o.AuthRRN, true
}

// HasAuthRRN returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasAuthRRN() bool {
	if o != nil && o.AuthRRN != nil {
		return true
	}

	return false
}

// SetAuthRRN gets a reference to the given string and assigns it to the AuthRRN field.
func (o *VoidResponseObjectVoidResponse) SetAuthRRN(v string) {
	o.AuthRRN = &v
}

// GetTransactionAuthenticationType returns the TransactionAuthenticationType field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetTransactionAuthenticationType() string {
	if o == nil || o.TransactionAuthenticationType == nil {
		var ret string
		return ret
	}
	return *o.TransactionAuthenticationType
}

// GetTransactionAuthenticationTypeOk returns a tuple with the TransactionAuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetTransactionAuthenticationTypeOk() (*string, bool) {
	if o == nil || o.TransactionAuthenticationType == nil {
		return nil, false
	}
	return o.TransactionAuthenticationType, true
}

// HasTransactionAuthenticationType returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasTransactionAuthenticationType() bool {
	if o != nil && o.TransactionAuthenticationType != nil {
		return true
	}

	return false
}

// SetTransactionAuthenticationType gets a reference to the given string and assigns it to the TransactionAuthenticationType field.
func (o *VoidResponseObjectVoidResponse) SetTransactionAuthenticationType(v string) {
	o.TransactionAuthenticationType = &v
}

// GetIdentifierForTheAdquirer returns the IdentifierForTheAdquirer field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetIdentifierForTheAdquirer() string {
	if o == nil || o.IdentifierForTheAdquirer == nil {
		var ret string
		return ret
	}
	return *o.IdentifierForTheAdquirer
}

// GetIdentifierForTheAdquirerOk returns a tuple with the IdentifierForTheAdquirer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetIdentifierForTheAdquirerOk() (*string, bool) {
	if o == nil || o.IdentifierForTheAdquirer == nil {
		return nil, false
	}
	return o.IdentifierForTheAdquirer, true
}

// HasIdentifierForTheAdquirer returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasIdentifierForTheAdquirer() bool {
	if o != nil && o.IdentifierForTheAdquirer != nil {
		return true
	}

	return false
}

// SetIdentifierForTheAdquirer gets a reference to the given string and assigns it to the IdentifierForTheAdquirer field.
func (o *VoidResponseObjectVoidResponse) SetIdentifierForTheAdquirer(v string) {
	o.IdentifierForTheAdquirer = &v
}

// GetIdentifierForTheResolutor returns the IdentifierForTheResolutor field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetIdentifierForTheResolutor() string {
	if o == nil || o.IdentifierForTheResolutor == nil {
		var ret string
		return ret
	}
	return *o.IdentifierForTheResolutor
}

// GetIdentifierForTheResolutorOk returns a tuple with the IdentifierForTheResolutor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetIdentifierForTheResolutorOk() (*string, bool) {
	if o == nil || o.IdentifierForTheResolutor == nil {
		return nil, false
	}
	return o.IdentifierForTheResolutor, true
}

// HasIdentifierForTheResolutor returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasIdentifierForTheResolutor() bool {
	if o != nil && o.IdentifierForTheResolutor != nil {
		return true
	}

	return false
}

// SetIdentifierForTheResolutor gets a reference to the given string and assigns it to the IdentifierForTheResolutor field.
func (o *VoidResponseObjectVoidResponse) SetIdentifierForTheResolutor(v string) {
	o.IdentifierForTheResolutor = &v
}

// GetPaymentFacilitatorID returns the PaymentFacilitatorID field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPaymentFacilitatorID() string {
	if o == nil || o.PaymentFacilitatorID == nil {
		var ret string
		return ret
	}
	return *o.PaymentFacilitatorID
}

// GetPaymentFacilitatorIDOk returns a tuple with the PaymentFacilitatorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPaymentFacilitatorIDOk() (*string, bool) {
	if o == nil || o.PaymentFacilitatorID == nil {
		return nil, false
	}
	return o.PaymentFacilitatorID, true
}

// HasPaymentFacilitatorID returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPaymentFacilitatorID() bool {
	if o != nil && o.PaymentFacilitatorID != nil {
		return true
	}

	return false
}

// SetPaymentFacilitatorID gets a reference to the given string and assigns it to the PaymentFacilitatorID field.
func (o *VoidResponseObjectVoidResponse) SetPaymentFacilitatorID(v string) {
	o.PaymentFacilitatorID = &v
}

// GetMerchantID returns the MerchantID field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetMerchantID() string {
	if o == nil || o.MerchantID == nil {
		var ret string
		return ret
	}
	return *o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetMerchantIDOk() (*string, bool) {
	if o == nil || o.MerchantID == nil {
		return nil, false
	}
	return o.MerchantID, true
}

// HasMerchantID returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasMerchantID() bool {
	if o != nil && o.MerchantID != nil {
		return true
	}

	return false
}

// SetMerchantID gets a reference to the given string and assigns it to the MerchantID field.
func (o *VoidResponseObjectVoidResponse) SetMerchantID(v string) {
	o.MerchantID = &v
}

// GetTerminalID returns the TerminalID field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetTerminalID() string {
	if o == nil || o.TerminalID == nil {
		var ret string
		return ret
	}
	return *o.TerminalID
}

// GetTerminalIDOk returns a tuple with the TerminalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetTerminalIDOk() (*string, bool) {
	if o == nil || o.TerminalID == nil {
		return nil, false
	}
	return o.TerminalID, true
}

// HasTerminalID returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasTerminalID() bool {
	if o != nil && o.TerminalID != nil {
		return true
	}

	return false
}

// SetTerminalID gets a reference to the given string and assigns it to the TerminalID field.
func (o *VoidResponseObjectVoidResponse) SetTerminalID(v string) {
	o.TerminalID = &v
}

// GetTerminalTrace returns the TerminalTrace field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetTerminalTrace() int32 {
	if o == nil || o.TerminalTrace == nil {
		var ret int32
		return ret
	}
	return *o.TerminalTrace
}

// GetTerminalTraceOk returns a tuple with the TerminalTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetTerminalTraceOk() (*int32, bool) {
	if o == nil || o.TerminalTrace == nil {
		return nil, false
	}
	return o.TerminalTrace, true
}

// HasTerminalTrace returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasTerminalTrace() bool {
	if o != nil && o.TerminalTrace != nil {
		return true
	}

	return false
}

// SetTerminalTrace gets a reference to the given int32 and assigns it to the TerminalTrace field.
func (o *VoidResponseObjectVoidResponse) SetTerminalTrace(v int32) {
	o.TerminalTrace = &v
}

// GetSettlementBatchNumber returns the SettlementBatchNumber field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetSettlementBatchNumber() int32 {
	if o == nil || o.SettlementBatchNumber == nil {
		var ret int32
		return ret
	}
	return *o.SettlementBatchNumber
}

// GetSettlementBatchNumberOk returns a tuple with the SettlementBatchNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetSettlementBatchNumberOk() (*int32, bool) {
	if o == nil || o.SettlementBatchNumber == nil {
		return nil, false
	}
	return o.SettlementBatchNumber, true
}

// HasSettlementBatchNumber returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasSettlementBatchNumber() bool {
	if o != nil && o.SettlementBatchNumber != nil {
		return true
	}

	return false
}

// SetSettlementBatchNumber gets a reference to the given int32 and assigns it to the SettlementBatchNumber field.
func (o *VoidResponseObjectVoidResponse) SetSettlementBatchNumber(v int32) {
	o.SettlementBatchNumber = &v
}

// GetCardReadMode returns the CardReadMode field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardReadMode() string {
	if o == nil || o.CardReadMode == nil {
		var ret string
		return ret
	}
	return *o.CardReadMode
}

// GetCardReadModeOk returns a tuple with the CardReadMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardReadModeOk() (*string, bool) {
	if o == nil || o.CardReadMode == nil {
		return nil, false
	}
	return o.CardReadMode, true
}

// HasCardReadMode returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardReadMode() bool {
	if o != nil && o.CardReadMode != nil {
		return true
	}

	return false
}

// SetCardReadMode gets a reference to the given string and assigns it to the CardReadMode field.
func (o *VoidResponseObjectVoidResponse) SetCardReadMode(v string) {
	o.CardReadMode = &v
}

// GetCardReadModeDescription returns the CardReadModeDescription field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardReadModeDescription() string {
	if o == nil || o.CardReadModeDescription == nil {
		var ret string
		return ret
	}
	return *o.CardReadModeDescription
}

// GetCardReadModeDescriptionOk returns a tuple with the CardReadModeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardReadModeDescriptionOk() (*string, bool) {
	if o == nil || o.CardReadModeDescription == nil {
		return nil, false
	}
	return o.CardReadModeDescription, true
}

// HasCardReadModeDescription returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardReadModeDescription() bool {
	if o != nil && o.CardReadModeDescription != nil {
		return true
	}

	return false
}

// SetCardReadModeDescription gets a reference to the given string and assigns it to the CardReadModeDescription field.
func (o *VoidResponseObjectVoidResponse) SetCardReadModeDescription(v string) {
	o.CardReadModeDescription = &v
}

// GetCardDescription returns the CardDescription field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardDescription() string {
	if o == nil || o.CardDescription == nil {
		var ret string
		return ret
	}
	return *o.CardDescription
}

// GetCardDescriptionOk returns a tuple with the CardDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardDescriptionOk() (*string, bool) {
	if o == nil || o.CardDescription == nil {
		return nil, false
	}
	return o.CardDescription, true
}

// HasCardDescription returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardDescription() bool {
	if o != nil && o.CardDescription != nil {
		return true
	}

	return false
}

// SetCardDescription gets a reference to the given string and assigns it to the CardDescription field.
func (o *VoidResponseObjectVoidResponse) SetCardDescription(v string) {
	o.CardDescription = &v
}

// GetCardTypeDescription returns the CardTypeDescription field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardTypeDescription() string {
	if o == nil || o.CardTypeDescription == nil {
		var ret string
		return ret
	}
	return *o.CardTypeDescription
}

// GetCardTypeDescriptionOk returns a tuple with the CardTypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardTypeDescriptionOk() (*string, bool) {
	if o == nil || o.CardTypeDescription == nil {
		return nil, false
	}
	return o.CardTypeDescription, true
}

// HasCardTypeDescription returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardTypeDescription() bool {
	if o != nil && o.CardTypeDescription != nil {
		return true
	}

	return false
}

// SetCardTypeDescription gets a reference to the given string and assigns it to the CardTypeDescription field.
func (o *VoidResponseObjectVoidResponse) SetCardTypeDescription(v string) {
	o.CardTypeDescription = &v
}

// GetCardCategory returns the CardCategory field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardCategory() SaleResponseObjectSaleResponseCardCategory {
	if o == nil || o.CardCategory == nil {
		var ret SaleResponseObjectSaleResponseCardCategory
		return ret
	}
	return *o.CardCategory
}

// GetCardCategoryOk returns a tuple with the CardCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardCategoryOk() (*SaleResponseObjectSaleResponseCardCategory, bool) {
	if o == nil || o.CardCategory == nil {
		return nil, false
	}
	return o.CardCategory, true
}

// HasCardCategory returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardCategory() bool {
	if o != nil && o.CardCategory != nil {
		return true
	}

	return false
}

// SetCardCategory gets a reference to the given SaleResponseObjectSaleResponseCardCategory and assigns it to the CardCategory field.
func (o *VoidResponseObjectVoidResponse) SetCardCategory(v SaleResponseObjectSaleResponseCardCategory) {
	o.CardCategory = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardNumber() string {
	if o == nil || o.CardNumber == nil {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardNumberOk() (*string, bool) {
	if o == nil || o.CardNumber == nil {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardNumber() bool {
	if o != nil && o.CardNumber != nil {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *VoidResponseObjectVoidResponse) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetCardNumberMasked returns the CardNumberMasked field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardNumberMasked() string {
	if o == nil || o.CardNumberMasked == nil {
		var ret string
		return ret
	}
	return *o.CardNumberMasked
}

// GetCardNumberMaskedOk returns a tuple with the CardNumberMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardNumberMaskedOk() (*string, bool) {
	if o == nil || o.CardNumberMasked == nil {
		return nil, false
	}
	return o.CardNumberMasked, true
}

// HasCardNumberMasked returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardNumberMasked() bool {
	if o != nil && o.CardNumberMasked != nil {
		return true
	}

	return false
}

// SetCardNumberMasked gets a reference to the given string and assigns it to the CardNumberMasked field.
func (o *VoidResponseObjectVoidResponse) SetCardNumberMasked(v string) {
	o.CardNumberMasked = &v
}

// GetCardHashing returns the CardHashing field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardHashing() string {
	if o == nil || o.CardHashing == nil {
		var ret string
		return ret
	}
	return *o.CardHashing
}

// GetCardHashingOk returns a tuple with the CardHashing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardHashingOk() (*string, bool) {
	if o == nil || o.CardHashing == nil {
		return nil, false
	}
	return o.CardHashing, true
}

// HasCardHashing returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardHashing() bool {
	if o != nil && o.CardHashing != nil {
		return true
	}

	return false
}

// SetCardHashing gets a reference to the given string and assigns it to the CardHashing field.
func (o *VoidResponseObjectVoidResponse) SetCardHashing(v string) {
	o.CardHashing = &v
}

// GetCardExp returns the CardExp field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardExp() string {
	if o == nil || o.CardExp == nil {
		var ret string
		return ret
	}
	return *o.CardExp
}

// GetCardExpOk returns a tuple with the CardExp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardExpOk() (*string, bool) {
	if o == nil || o.CardExp == nil {
		return nil, false
	}
	return o.CardExp, true
}

// HasCardExp returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardExp() bool {
	if o != nil && o.CardExp != nil {
		return true
	}

	return false
}

// SetCardExp gets a reference to the given string and assigns it to the CardExp field.
func (o *VoidResponseObjectVoidResponse) SetCardExp(v string) {
	o.CardExp = &v
}

// GetCardCryptogramResponse returns the CardCryptogramResponse field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardCryptogramResponse() string {
	if o == nil || o.CardCryptogramResponse == nil {
		var ret string
		return ret
	}
	return *o.CardCryptogramResponse
}

// GetCardCryptogramResponseOk returns a tuple with the CardCryptogramResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardCryptogramResponseOk() (*string, bool) {
	if o == nil || o.CardCryptogramResponse == nil {
		return nil, false
	}
	return o.CardCryptogramResponse, true
}

// HasCardCryptogramResponse returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardCryptogramResponse() bool {
	if o != nil && o.CardCryptogramResponse != nil {
		return true
	}

	return false
}

// SetCardCryptogramResponse gets a reference to the given string and assigns it to the CardCryptogramResponse field.
func (o *VoidResponseObjectVoidResponse) SetCardCryptogramResponse(v string) {
	o.CardCryptogramResponse = &v
}

// GetCardAppName returns the CardAppName field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardAppName() string {
	if o == nil || o.CardAppName == nil {
		var ret string
		return ret
	}
	return *o.CardAppName
}

// GetCardAppNameOk returns a tuple with the CardAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardAppNameOk() (*string, bool) {
	if o == nil || o.CardAppName == nil {
		return nil, false
	}
	return o.CardAppName, true
}

// HasCardAppName returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardAppName() bool {
	if o != nil && o.CardAppName != nil {
		return true
	}

	return false
}

// SetCardAppName gets a reference to the given string and assigns it to the CardAppName field.
func (o *VoidResponseObjectVoidResponse) SetCardAppName(v string) {
	o.CardAppName = &v
}

// GetCardAppIdentifier returns the CardAppIdentifier field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardAppIdentifier() string {
	if o == nil || o.CardAppIdentifier == nil {
		var ret string
		return ret
	}
	return *o.CardAppIdentifier
}

// GetCardAppIdentifierOk returns a tuple with the CardAppIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardAppIdentifierOk() (*string, bool) {
	if o == nil || o.CardAppIdentifier == nil {
		return nil, false
	}
	return o.CardAppIdentifier, true
}

// HasCardAppIdentifier returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardAppIdentifier() bool {
	if o != nil && o.CardAppIdentifier != nil {
		return true
	}

	return false
}

// SetCardAppIdentifier gets a reference to the given string and assigns it to the CardAppIdentifier field.
func (o *VoidResponseObjectVoidResponse) SetCardAppIdentifier(v string) {
	o.CardAppIdentifier = &v
}

// GetCardAppLabel returns the CardAppLabel field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardAppLabel() string {
	if o == nil || o.CardAppLabel == nil {
		var ret string
		return ret
	}
	return *o.CardAppLabel
}

// GetCardAppLabelOk returns a tuple with the CardAppLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardAppLabelOk() (*string, bool) {
	if o == nil || o.CardAppLabel == nil {
		return nil, false
	}
	return o.CardAppLabel, true
}

// HasCardAppLabel returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardAppLabel() bool {
	if o != nil && o.CardAppLabel != nil {
		return true
	}

	return false
}

// SetCardAppLabel gets a reference to the given string and assigns it to the CardAppLabel field.
func (o *VoidResponseObjectVoidResponse) SetCardAppLabel(v string) {
	o.CardAppLabel = &v
}

// GetCardAuthRequestCryptogram returns the CardAuthRequestCryptogram field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardAuthRequestCryptogram() string {
	if o == nil || o.CardAuthRequestCryptogram == nil {
		var ret string
		return ret
	}
	return *o.CardAuthRequestCryptogram
}

// GetCardAuthRequestCryptogramOk returns a tuple with the CardAuthRequestCryptogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardAuthRequestCryptogramOk() (*string, bool) {
	if o == nil || o.CardAuthRequestCryptogram == nil {
		return nil, false
	}
	return o.CardAuthRequestCryptogram, true
}

// HasCardAuthRequestCryptogram returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardAuthRequestCryptogram() bool {
	if o != nil && o.CardAuthRequestCryptogram != nil {
		return true
	}

	return false
}

// SetCardAuthRequestCryptogram gets a reference to the given string and assigns it to the CardAuthRequestCryptogram field.
func (o *VoidResponseObjectVoidResponse) SetCardAuthRequestCryptogram(v string) {
	o.CardAuthRequestCryptogram = &v
}

// GetCardAuthResponseCryptogram returns the CardAuthResponseCryptogram field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCardAuthResponseCryptogram() string {
	if o == nil || o.CardAuthResponseCryptogram == nil {
		var ret string
		return ret
	}
	return *o.CardAuthResponseCryptogram
}

// GetCardAuthResponseCryptogramOk returns a tuple with the CardAuthResponseCryptogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCardAuthResponseCryptogramOk() (*string, bool) {
	if o == nil || o.CardAuthResponseCryptogram == nil {
		return nil, false
	}
	return o.CardAuthResponseCryptogram, true
}

// HasCardAuthResponseCryptogram returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCardAuthResponseCryptogram() bool {
	if o != nil && o.CardAuthResponseCryptogram != nil {
		return true
	}

	return false
}

// SetCardAuthResponseCryptogram gets a reference to the given string and assigns it to the CardAuthResponseCryptogram field.
func (o *VoidResponseObjectVoidResponse) SetCardAuthResponseCryptogram(v string) {
	o.CardAuthResponseCryptogram = &v
}

// GetPayer returns the Payer field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPayer() SaleObjectSalePayer {
	if o == nil || o.Payer == nil {
		var ret SaleObjectSalePayer
		return ret
	}
	return *o.Payer
}

// GetPayerOk returns a tuple with the Payer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPayerOk() (*SaleObjectSalePayer, bool) {
	if o == nil || o.Payer == nil {
		return nil, false
	}
	return o.Payer, true
}

// HasPayer returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPayer() bool {
	if o != nil && o.Payer != nil {
		return true
	}

	return false
}

// SetPayer gets a reference to the given SaleObjectSalePayer and assigns it to the Payer field.
func (o *VoidResponseObjectVoidResponse) SetPayer(v SaleObjectSalePayer) {
	o.Payer = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetCustomer() SaleObjectSaleCustomer {
	if o == nil || o.Customer == nil {
		var ret SaleObjectSaleCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetCustomerOk() (*SaleObjectSaleCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given SaleObjectSaleCustomer and assigns it to the Customer field.
func (o *VoidResponseObjectVoidResponse) SetCustomer(v SaleObjectSaleCustomer) {
	o.Customer = &v
}

// GetMerchantCategory returns the MerchantCategory field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetMerchantCategory() SaleResponseObjectSaleResponseMerchantCategory {
	if o == nil || o.MerchantCategory == nil {
		var ret SaleResponseObjectSaleResponseMerchantCategory
		return ret
	}
	return *o.MerchantCategory
}

// GetMerchantCategoryOk returns a tuple with the MerchantCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetMerchantCategoryOk() (*SaleResponseObjectSaleResponseMerchantCategory, bool) {
	if o == nil || o.MerchantCategory == nil {
		return nil, false
	}
	return o.MerchantCategory, true
}

// HasMerchantCategory returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasMerchantCategory() bool {
	if o != nil && o.MerchantCategory != nil {
		return true
	}

	return false
}

// SetMerchantCategory gets a reference to the given SaleResponseObjectSaleResponseMerchantCategory and assigns it to the MerchantCategory field.
func (o *VoidResponseObjectVoidResponse) SetMerchantCategory(v SaleResponseObjectSaleResponseMerchantCategory) {
	o.MerchantCategory = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetProducts() []SaleResponseObjectSaleResponseProducts {
	if o == nil || o.Products == nil {
		var ret []SaleResponseObjectSaleResponseProducts
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetProductsOk() ([]SaleResponseObjectSaleResponseProducts, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []SaleResponseObjectSaleResponseProducts and assigns it to the Products field.
func (o *VoidResponseObjectVoidResponse) SetProducts(v []SaleResponseObjectSaleResponseProducts) {
	o.Products = v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPaymentMethod() SaleResponseObjectSaleResponsePaymentMethod {
	if o == nil || o.PaymentMethod == nil {
		var ret SaleResponseObjectSaleResponsePaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPaymentMethodOk() (*SaleResponseObjectSaleResponsePaymentMethod, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given SaleResponseObjectSaleResponsePaymentMethod and assigns it to the PaymentMethod field.
func (o *VoidResponseObjectVoidResponse) SetPaymentMethod(v SaleResponseObjectSaleResponsePaymentMethod) {
	o.PaymentMethod = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPlans() SaleResponseObjectSaleResponsePlans {
	if o == nil || o.Plans == nil {
		var ret SaleResponseObjectSaleResponsePlans
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPlansOk() (*SaleResponseObjectSaleResponsePlans, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given SaleResponseObjectSaleResponsePlans and assigns it to the Plans field.
func (o *VoidResponseObjectVoidResponse) SetPlans(v SaleResponseObjectSaleResponsePlans) {
	o.Plans = &v
}

// GetPlanDescription returns the PlanDescription field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPlanDescription() string {
	if o == nil || o.PlanDescription == nil {
		var ret string
		return ret
	}
	return *o.PlanDescription
}

// GetPlanDescriptionOk returns a tuple with the PlanDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPlanDescriptionOk() (*string, bool) {
	if o == nil || o.PlanDescription == nil {
		return nil, false
	}
	return o.PlanDescription, true
}

// HasPlanDescription returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPlanDescription() bool {
	if o != nil && o.PlanDescription != nil {
		return true
	}

	return false
}

// SetPlanDescription gets a reference to the given string and assigns it to the PlanDescription field.
func (o *VoidResponseObjectVoidResponse) SetPlanDescription(v string) {
	o.PlanDescription = &v
}

// GetPlanConfigVersion returns the PlanConfigVersion field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPlanConfigVersion() int32 {
	if o == nil || o.PlanConfigVersion == nil {
		var ret int32
		return ret
	}
	return *o.PlanConfigVersion
}

// GetPlanConfigVersionOk returns a tuple with the PlanConfigVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPlanConfigVersionOk() (*int32, bool) {
	if o == nil || o.PlanConfigVersion == nil {
		return nil, false
	}
	return o.PlanConfigVersion, true
}

// HasPlanConfigVersion returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPlanConfigVersion() bool {
	if o != nil && o.PlanConfigVersion != nil {
		return true
	}

	return false
}

// SetPlanConfigVersion gets a reference to the given int32 and assigns it to the PlanConfigVersion field.
func (o *VoidResponseObjectVoidResponse) SetPlanConfigVersion(v int32) {
	o.PlanConfigVersion = &v
}

// GetTickets returns the Tickets field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetTickets() []SaleResponseObjectSaleResponseTickets {
	if o == nil || o.Tickets == nil {
		var ret []SaleResponseObjectSaleResponseTickets
		return ret
	}
	return o.Tickets
}

// GetTicketsOk returns a tuple with the Tickets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetTicketsOk() ([]SaleResponseObjectSaleResponseTickets, bool) {
	if o == nil || o.Tickets == nil {
		return nil, false
	}
	return o.Tickets, true
}

// HasTickets returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasTickets() bool {
	if o != nil && o.Tickets != nil {
		return true
	}

	return false
}

// SetTickets gets a reference to the given []SaleResponseObjectSaleResponseTickets and assigns it to the Tickets field.
func (o *VoidResponseObjectVoidResponse) SetTickets(v []SaleResponseObjectSaleResponseTickets) {
	o.Tickets = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetConfiguration() SaleResponseObjectSaleResponseConfiguration {
	if o == nil || o.Configuration == nil {
		var ret SaleResponseObjectSaleResponseConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetConfigurationOk() (*SaleResponseObjectSaleResponseConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given SaleResponseObjectSaleResponseConfiguration and assigns it to the Configuration field.
func (o *VoidResponseObjectVoidResponse) SetConfiguration(v SaleResponseObjectSaleResponseConfiguration) {
	o.Configuration = &v
}

// GetOrigAnswerKey returns the OrigAnswerKey field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetOrigAnswerKey() string {
	if o == nil || o.OrigAnswerKey == nil {
		var ret string
		return ret
	}
	return *o.OrigAnswerKey
}

// GetOrigAnswerKeyOk returns a tuple with the OrigAnswerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetOrigAnswerKeyOk() (*string, bool) {
	if o == nil || o.OrigAnswerKey == nil {
		return nil, false
	}
	return o.OrigAnswerKey, true
}

// HasOrigAnswerKey returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasOrigAnswerKey() bool {
	if o != nil && o.OrigAnswerKey != nil {
		return true
	}

	return false
}

// SetOrigAnswerKey gets a reference to the given string and assigns it to the OrigAnswerKey field.
func (o *VoidResponseObjectVoidResponse) SetOrigAnswerKey(v string) {
	o.OrigAnswerKey = &v
}

// GetOrigAuthDateTime returns the OrigAuthDateTime field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetOrigAuthDateTime() time.Time {
	if o == nil || o.OrigAuthDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.OrigAuthDateTime
}

// GetOrigAuthDateTimeOk returns a tuple with the OrigAuthDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetOrigAuthDateTimeOk() (*time.Time, bool) {
	if o == nil || o.OrigAuthDateTime == nil {
		return nil, false
	}
	return o.OrigAuthDateTime, true
}

// HasOrigAuthDateTime returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasOrigAuthDateTime() bool {
	if o != nil && o.OrigAuthDateTime != nil {
		return true
	}

	return false
}

// SetOrigAuthDateTime gets a reference to the given time.Time and assigns it to the OrigAuthDateTime field.
func (o *VoidResponseObjectVoidResponse) SetOrigAuthDateTime(v time.Time) {
	o.OrigAuthDateTime = &v
}

// GetOrigAuthTicket returns the OrigAuthTicket field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetOrigAuthTicket() int32 {
	if o == nil || o.OrigAuthTicket == nil {
		var ret int32
		return ret
	}
	return *o.OrigAuthTicket
}

// GetOrigAuthTicketOk returns a tuple with the OrigAuthTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetOrigAuthTicketOk() (*int32, bool) {
	if o == nil || o.OrigAuthTicket == nil {
		return nil, false
	}
	return o.OrigAuthTicket, true
}

// HasOrigAuthTicket returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasOrigAuthTicket() bool {
	if o != nil && o.OrigAuthTicket != nil {
		return true
	}

	return false
}

// SetOrigAuthTicket gets a reference to the given int32 and assigns it to the OrigAuthTicket field.
func (o *VoidResponseObjectVoidResponse) SetOrigAuthTicket(v int32) {
	o.OrigAuthTicket = &v
}

// GetOrigAuthTerminalTrace returns the OrigAuthTerminalTrace field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetOrigAuthTerminalTrace() int32 {
	if o == nil || o.OrigAuthTerminalTrace == nil {
		var ret int32
		return ret
	}
	return *o.OrigAuthTerminalTrace
}

// GetOrigAuthTerminalTraceOk returns a tuple with the OrigAuthTerminalTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetOrigAuthTerminalTraceOk() (*int32, bool) {
	if o == nil || o.OrigAuthTerminalTrace == nil {
		return nil, false
	}
	return o.OrigAuthTerminalTrace, true
}

// HasOrigAuthTerminalTrace returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasOrigAuthTerminalTrace() bool {
	if o != nil && o.OrigAuthTerminalTrace != nil {
		return true
	}

	return false
}

// SetOrigAuthTerminalTrace gets a reference to the given int32 and assigns it to the OrigAuthTerminalTrace field.
func (o *VoidResponseObjectVoidResponse) SetOrigAuthTerminalTrace(v int32) {
	o.OrigAuthTerminalTrace = &v
}

// GetOrigAuthCode returns the OrigAuthCode field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetOrigAuthCode() string {
	if o == nil || o.OrigAuthCode == nil {
		var ret string
		return ret
	}
	return *o.OrigAuthCode
}

// GetOrigAuthCodeOk returns a tuple with the OrigAuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetOrigAuthCodeOk() (*string, bool) {
	if o == nil || o.OrigAuthCode == nil {
		return nil, false
	}
	return o.OrigAuthCode, true
}

// HasOrigAuthCode returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasOrigAuthCode() bool {
	if o != nil && o.OrigAuthCode != nil {
		return true
	}

	return false
}

// SetOrigAuthCode gets a reference to the given string and assigns it to the OrigAuthCode field.
func (o *VoidResponseObjectVoidResponse) SetOrigAuthCode(v string) {
	o.OrigAuthCode = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPaymentMethodId() int32 {
	if o == nil || o.PaymentMethodId == nil {
		var ret int32
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPaymentMethodIdOk() (*int32, bool) {
	if o == nil || o.PaymentMethodId == nil {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId != nil {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given int32 and assigns it to the PaymentMethodId field.
func (o *VoidResponseObjectVoidResponse) SetPaymentMethodId(v int32) {
	o.PaymentMethodId = &v
}

// GetPaymentMethodDescription returns the PaymentMethodDescription field value if set, zero value otherwise.
func (o *VoidResponseObjectVoidResponse) GetPaymentMethodDescription() string {
	if o == nil || o.PaymentMethodDescription == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodDescription
}

// GetPaymentMethodDescriptionOk returns a tuple with the PaymentMethodDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseObjectVoidResponse) GetPaymentMethodDescriptionOk() (*string, bool) {
	if o == nil || o.PaymentMethodDescription == nil {
		return nil, false
	}
	return o.PaymentMethodDescription, true
}

// HasPaymentMethodDescription returns a boolean if a field has been set.
func (o *VoidResponseObjectVoidResponse) HasPaymentMethodDescription() bool {
	if o != nil && o.PaymentMethodDescription != nil {
		return true
	}

	return false
}

// SetPaymentMethodDescription gets a reference to the given string and assigns it to the PaymentMethodDescription field.
func (o *VoidResponseObjectVoidResponse) SetPaymentMethodDescription(v string) {
	o.PaymentMethodDescription = &v
}

func (o VoidResponseObjectVoidResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ResponseCode"] = o.ResponseCode
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
	if o.RequestType != nil {
		toSerialize["RequestType"] = o.RequestType
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
	if o.RequestKey != nil {
		toSerialize["RequestKey"] = o.RequestKey
	}
	if o.WorkingKeys != nil {
		toSerialize["WorkingKeys"] = o.WorkingKeys
	}
	if o.RequiredInformation != nil {
		toSerialize["RequiredInformation"] = o.RequiredInformation
	}
	if o.AnswerType != nil {
		toSerialize["AnswerType"] = o.AnswerType
	}
	if o.AnswerKey != nil {
		toSerialize["AnswerKey"] = o.AnswerKey
	}
	if o.PublicAnswerKey != nil {
		toSerialize["PublicAnswerKey"] = o.PublicAnswerKey
	}
	if o.WasReversePrevious != nil {
		toSerialize["WasReversePrevious"] = o.WasReversePrevious
	}
	if o.ReversedAnswerKey != nil {
		toSerialize["ReversedAnswerKey"] = o.ReversedAnswerKey
	}
	if o.ReversedSequence != nil {
		toSerialize["ReversedSequence"] = o.ReversedSequence
	}
	if o.CommittedBlock != nil {
		toSerialize["CommittedBlock"] = o.CommittedBlock
	}
	if o.ReversedBlock != nil {
		toSerialize["ReversedBlock"] = o.ReversedBlock
	}
	if o.MessageID != nil {
		toSerialize["MessageID"] = o.MessageID
	}
	if o.ServerAddress != nil {
		toSerialize["ServerAddress"] = o.ServerAddress
	}
	if o.ServerInstance != nil {
		toSerialize["ServerInstance"] = o.ServerInstance
	}
	if o.ServerNodeName != nil {
		toSerialize["ServerNodeName"] = o.ServerNodeName
	}
	if o.AdapterInputVersion != nil {
		toSerialize["AdapterInputVersion"] = o.AdapterInputVersion
	}
	if o.AdapterInputAddress != nil {
		toSerialize["AdapterInputAddress"] = o.AdapterInputAddress
	}
	if o.AdapterInputNodeName != nil {
		toSerialize["AdapterInputNodeName"] = o.AdapterInputNodeName
	}
	if o.AdapterOutputVersion != nil {
		toSerialize["AdapterOutputVersion"] = o.AdapterOutputVersion
	}
	if o.AdapterOutputAddress != nil {
		toSerialize["AdapterOutputAddress"] = o.AdapterOutputAddress
	}
	if o.AdapterOutputNodeName != nil {
		toSerialize["AdapterOutputNodeName"] = o.AdapterOutputNodeName
	}
	if o.AdditionalInformation != nil {
		toSerialize["AdditionalInformation"] = o.AdditionalInformation
	}
	if o.PrinterResponseMessage != nil {
		toSerialize["PrinterResponseMessage"] = o.PrinterResponseMessage
	}
	if o.DisplayResponseMessage != nil {
		toSerialize["DisplayResponseMessage"] = o.DisplayResponseMessage
	}
	if o.CurrencyCode != nil {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if o.CurrencyDescription != nil {
		toSerialize["CurrencyDescription"] = o.CurrencyDescription
	}
	if o.TransactionDescription != nil {
		toSerialize["TransactionDescription"] = o.TransactionDescription
	}
	if o.TransactionResolutionMode != nil {
		toSerialize["TransactionResolutionMode"] = o.TransactionResolutionMode
	}
	if o.CurrencySymbol != nil {
		toSerialize["CurrencySymbol"] = o.CurrencySymbol
	}
	if o.FacilityPayments != nil {
		toSerialize["FacilityPayments"] = o.FacilityPayments
	}
	if o.FacilityType != nil {
		toSerialize["FacilityType"] = o.FacilityType
	}
	if o.ValueFacilityPayments != nil {
		toSerialize["ValueFacilityPayments"] = o.ValueFacilityPayments
	}
	if o.Amount != nil {
		toSerialize["Amount"] = o.Amount
	}
	if o.AlternativeAmount != nil {
		toSerialize["AlternativeAmount"] = o.AlternativeAmount
	}
	if o.AmountCharged != nil {
		toSerialize["AmountCharged"] = o.AmountCharged
	}
	if o.AmountToApply != nil {
		toSerialize["AmountToApply"] = o.AmountToApply
	}
	if o.CashbackAmount != nil {
		toSerialize["CashbackAmount"] = o.CashbackAmount
	}
	if o.TipAmount != nil {
		toSerialize["TipAmount"] = o.TipAmount
	}
	if o.RequestAmount != nil {
		toSerialize["RequestAmount"] = o.RequestAmount
	}
	if o.PromotedAmount != nil {
		toSerialize["PromotedAmount"] = o.PromotedAmount
	}
	if o.CredentialToken != nil {
		toSerialize["CredentialToken"] = o.CredentialToken
	}
	if o.CredentialIssuerToken != nil {
		toSerialize["CredentialIssuerToken"] = o.CredentialIssuerToken
	}
	if o.InputTokens != nil {
		toSerialize["InputTokens"] = o.InputTokens
	}
	if o.OutputTokens != nil {
		toSerialize["OutputTokens"] = o.OutputTokens
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
	if o.HostResultCode != nil {
		toSerialize["HostResultCode"] = o.HostResultCode
	}
	if o.HostMessage != nil {
		toSerialize["HostMessage"] = o.HostMessage
	}
	if o.HostCode != nil {
		toSerialize["HostCode"] = o.HostCode
	}
	if o.HostID != nil {
		toSerialize["HostID"] = o.HostID
	}
	if o.UserID != nil {
		toSerialize["UserID"] = o.UserID
	}
	if o.IssuerName != nil {
		toSerialize["IssuerName"] = o.IssuerName
	}
	if o.HostDateTime != nil {
		toSerialize["HostDateTime"] = o.HostDateTime
	}
	if o.TransmitionDateTime != nil {
		toSerialize["TransmitionDateTime"] = o.TransmitionDateTime
	}
	if o.AuthResultCode != nil {
		toSerialize["AuthResultCode"] = o.AuthResultCode
	}
	if o.AuthHostProcessCode != nil {
		toSerialize["AuthHostProcessCode"] = o.AuthHostProcessCode
	}
	if o.AuthHostMsgType != nil {
		toSerialize["AuthHostMsgType"] = o.AuthHostMsgType
	}
	if o.AuthHostMessage != nil {
		toSerialize["AuthHostMessage"] = o.AuthHostMessage
	}
	if o.HostMsgType != nil {
		toSerialize["HostMsgType"] = o.HostMsgType
	}
	if o.AuthCode != nil {
		toSerialize["AuthCode"] = o.AuthCode
	}
	if o.AuthDateTime != nil {
		toSerialize["AuthDateTime"] = o.AuthDateTime
	}
	if o.AuthTicket != nil {
		toSerialize["AuthTicket"] = o.AuthTicket
	}
	if o.AuthRRN != nil {
		toSerialize["AuthRRN"] = o.AuthRRN
	}
	if o.TransactionAuthenticationType != nil {
		toSerialize["TransactionAuthenticationType"] = o.TransactionAuthenticationType
	}
	if o.IdentifierForTheAdquirer != nil {
		toSerialize["IdentifierForTheAdquirer"] = o.IdentifierForTheAdquirer
	}
	if o.IdentifierForTheResolutor != nil {
		toSerialize["IdentifierForTheResolutor"] = o.IdentifierForTheResolutor
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
	if o.CardReadMode != nil {
		toSerialize["CardReadMode"] = o.CardReadMode
	}
	if o.CardReadModeDescription != nil {
		toSerialize["CardReadModeDescription"] = o.CardReadModeDescription
	}
	if o.CardDescription != nil {
		toSerialize["CardDescription"] = o.CardDescription
	}
	if o.CardTypeDescription != nil {
		toSerialize["CardTypeDescription"] = o.CardTypeDescription
	}
	if o.CardCategory != nil {
		toSerialize["CardCategory"] = o.CardCategory
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
	if o.CardExp != nil {
		toSerialize["CardExp"] = o.CardExp
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
	if o.CardAuthRequestCryptogram != nil {
		toSerialize["CardAuthRequestCryptogram"] = o.CardAuthRequestCryptogram
	}
	if o.CardAuthResponseCryptogram != nil {
		toSerialize["CardAuthResponseCryptogram"] = o.CardAuthResponseCryptogram
	}
	if o.Payer != nil {
		toSerialize["Payer"] = o.Payer
	}
	if o.Customer != nil {
		toSerialize["Customer"] = o.Customer
	}
	if o.MerchantCategory != nil {
		toSerialize["MerchantCategory"] = o.MerchantCategory
	}
	if o.Products != nil {
		toSerialize["Products"] = o.Products
	}
	if o.PaymentMethod != nil {
		toSerialize["PaymentMethod"] = o.PaymentMethod
	}
	if o.Plans != nil {
		toSerialize["Plans"] = o.Plans
	}
	if o.PlanDescription != nil {
		toSerialize["PlanDescription"] = o.PlanDescription
	}
	if o.PlanConfigVersion != nil {
		toSerialize["PlanConfigVersion"] = o.PlanConfigVersion
	}
	if o.Tickets != nil {
		toSerialize["Tickets"] = o.Tickets
	}
	if o.Configuration != nil {
		toSerialize["Configuration"] = o.Configuration
	}
	if o.OrigAnswerKey != nil {
		toSerialize["OrigAnswerKey"] = o.OrigAnswerKey
	}
	if o.OrigAuthDateTime != nil {
		toSerialize["OrigAuthDateTime"] = o.OrigAuthDateTime
	}
	if o.OrigAuthTicket != nil {
		toSerialize["OrigAuthTicket"] = o.OrigAuthTicket
	}
	if o.OrigAuthTerminalTrace != nil {
		toSerialize["OrigAuthTerminalTrace"] = o.OrigAuthTerminalTrace
	}
	if o.OrigAuthCode != nil {
		toSerialize["OrigAuthCode"] = o.OrigAuthCode
	}
	if o.PaymentMethodId != nil {
		toSerialize["PaymentMethodId"] = o.PaymentMethodId
	}
	if o.PaymentMethodDescription != nil {
		toSerialize["PaymentMethodDescription"] = o.PaymentMethodDescription
	}
	return json.Marshal(toSerialize)
}

type NullableVoidResponseObjectVoidResponse struct {
	value *VoidResponseObjectVoidResponse
	isSet bool
}

func (v NullableVoidResponseObjectVoidResponse) Get() *VoidResponseObjectVoidResponse {
	return v.value
}

func (v *NullableVoidResponseObjectVoidResponse) Set(val *VoidResponseObjectVoidResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVoidResponseObjectVoidResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVoidResponseObjectVoidResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoidResponseObjectVoidResponse(val *VoidResponseObjectVoidResponse) *NullableVoidResponseObjectVoidResponse {
	return &NullableVoidResponseObjectVoidResponse{value: val, isSet: true}
}

func (v NullableVoidResponseObjectVoidResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoidResponseObjectVoidResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


