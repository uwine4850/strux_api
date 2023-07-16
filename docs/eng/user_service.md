## Microservice for users
This section describes all the necessary information about the user microservice.

***
### Creating a user
**Address:** http://host/create-user/ <br>
**Form fields (case sensitive):**

* _username(text)_ - Username
* _password(text)_ - Password

**Response:** <u>baseproto.BaseResponse</u>. The response contains current information about the status of the request.

***
### Checking the availability of a user
**Address:** http://host/user-exist/ <br>
**Form fields (case sensitive):**

* _username(text)_ - Username

**Response:** <u>baseproto.BaseResponse</u>. The response contains current information about the status of the request.

***
### Deleting a user
**Address:** http://host/user-delete/ <br>
**Form fields (case sensitive):**
* _username(text)_ - Username
* _password(text)_ - Password

**Response:** <u>baseproto.BaseResponse</u>. The response contains current information about the status of the request.

***
### Update password
**Address:** http://host/user-password-update/ <br>
**Form fields (case sensitive):**

* _username(text)_ - Username
* _password(text)_ - Password
* _newPassword(text)_ - New password

**Response:** <u>baseproto.BaseResponse</u>. The response contains current information about the status of the request.

***
### User login
**Address:** http://host/user-log-in/ <br>
**Form fields (case sensitive):**

* _username(text)_ - Username
* _password(text)_ - Password

**Response:** <u>baseproto.BaseResponse</u>. The response contains current information about the status of the request.
