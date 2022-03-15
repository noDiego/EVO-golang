## Implementacion

- El Codigo fue generado con openapi-generator-5.4.0
- Para su instalación se copio el codigo generado a la raiz del proyecto de ejemplo en un directorio de nombre "openapi" dado el package generado
- Se genero un go.mod en la raiz del proyecto que hace referencia al directorio generado
- Se agregaron todas las dependencias necesarias para la ejecucion del codigo generado

# ERRORES:
- El codigo generado con openapi presentó errores en los tipos de datos multiples.
  - Se solucionó editando el yaml para dejar los tipos de datos multiples como string

## Uso

- Se debe ejecutar el archivo index.exe para su ejecucion
- Puede agregarse el parametro "--hostname=" para modificar el hostname (ej: index.exe --hostname=http://localhost:8080)
- En caso de necesitar compilarse debe ejecutarse el comando "go build ./index.go"
