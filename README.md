# transaction-stori
_ Component for send email with transactions to users _
## Starting 🚀
### 📋 Pre-requisitos
_For the execution and testing of this project, the following tools are necessary_
***golang** version 1.16
***Docker**
***Postman**
## ⚙️ Local Execution
_To start the microservice locally it must be done with the following command_
```
user$ go run main.go
```
_Then you need send an HTTP POST to_
```
http://localhost:8080/transaction/send-email
```
## ⚙️ Docker Execution
_To start the microservice in a docker container you need to do the following steps_
_First you build the image_
```
docker build --tag image_name .
```
_Then you run the image inside of a container_
```
docker run --publish 8080:8080 image_name
```
## Author✒️
* Alonso Orellana
