## The structure of microservices
* All microservices are stored in the services/ directory
* Protobuf files are stored in the services/proto_source directory
* The generated files are stored in the services/protofiles directory
* The utils file is responsible for the general logic of the services
* Each microservice contains a main.go file. It initializes the corresponding service, and when the file is started, the
  service is also launched.
* The services/<service_name>/internal directory contains files with command handlers. Each file contains the logic
  of one command is stored in each file. The corresponding command is called in the main.go file in the corresponding method
* In general, data processing in services is as follows connect to the database -> get the db.DatabaseOperation{}
  structure for interacting with the selected table -> executing business logic -> returning a response.