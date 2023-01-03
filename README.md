# transaction-stori
_ Componente encargo de realizar las peticiones a la API Products _
## Comenzando 🚀
_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas._
### 📋 Pre-requisitos
_Para la ejecucion y testeo de este proyecto, son necesarios las siguientes herramientas_
***golang** version 1.14 o superior
***Docker**
***kubectl**
***minikube** o la maquina virutal de **kubernetes** de docker 
***Postman**
## ⚙️ Ejecucion Local
_Para iniciar el microservicio localmente se debe hacer con el siguiente comando_
```
user$ go run main.go
```
## ⚙️ Ejecucion de Pruebas
_Para la ejecucion de pruebas, es necesario que el proyecto no se encuentre corriendo, dado que al ejecutar las pruebas, se volvera a levantar el servidor en el mismo puerto que se especifica en el archivo y dara un error de puerto en uso. Si el proyecto se encuentra corriendo se puede detener con **ctrl+c**/**(^)+c**_
_Se ejecutan las pruebas con el siguiente comando_
```
user$ go test -coverprofile cover.out ./src/...
```
## Versionado 📌
### Version 0.0.1
_Version inicial_
* Se implementa funcionalidad de conexion con API Products
## Autor(es))✒️
* Camilo Fuentes.