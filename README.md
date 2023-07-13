<h1 align="center" style="margin: 0">Strux api</h1>

## Other sections
* [The structure of microservices](https://github.com/uwine4850/strux_api/blob/master/docs/eng/microservices_structure.md)
* [More about the user microservice](https://github.com/uwine4850/strux_api/blob/master/docs/eng/user_service.md)
* [Learn more about package microservice](https://github.com/uwine4850/strux_api/blob/master/docs/eng/package_service.md)

## About the project
Strux api is a project designed for package management (also requires
[server client](https://github.com/uwine4850/strux)). A package is just a directory with a user project. 
Everything in the same directory can be considered a package if properly prepared. For other users to download packages,
they need to be moved to the server using a client.
To create a package, you need to use the [server client](https://github.com/uwine4850/strux).<p></p>

## Getting started
To get started, follow the instructions below.

If this project is used as a server, the repository must be downloaded from this link:<br>
```
https://github.com/uwine4850/strux_api
```
To start the server, you need to start the server. First, you need to start the microservices (main.go) in the services directory:<br>
* Microservice for users
* Package microservice
<p></p>
After that, you need to start the http server located at cmd/main.go.

If this project is used as an additional package, you need to install it using the command:<br>
```
go get github.com/uwine4850/strux_api
```

## Information about the api
The service contains the following commands:
* http://host/create-user/ - user creation. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/eng/user_service.md)
* http://host/user-exist/ - checking the availability of the user. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/eng/user_service.md)
* http://host/user-delete/ - deleting a user. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/eng/user_service.md)
* http://host/user-password-update/ - password update. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/eng/user_service.md)
* http://host/user-log-in/ - user login to the system. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/eng/user_service.md)
* http://host/upload-pkg/ - uploading the package to the server. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/eng/package_service.md)
* http://host/exists-package/ -  checking the availability of the package. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/eng/package_service.md)
* http://host/download-package/ - downloading a package from the server. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/eng/package_service.md)
* http://host/package-versions/ - returns all available versions of the package. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/eng/package_service.md)

### Briefly about the work algorithm
To execute the command, you need to send a form to one of the http addresses. For example, to send a form using golang
you can use <u>http.NewRequest</u>. The server accepts only data in the form of a form, the fields can contain text or file data.
After processing the form, the server returns either text data (text/plain) or a form (multipart/form-data). To use the
api, it is better to use a client, but it is not required.