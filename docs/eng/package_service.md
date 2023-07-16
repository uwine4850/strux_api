## Package microservice
This section describes all the information you need about the package microservice.

***
### Uploading a package to the server
**Address:** http://host/upload-pkg/ <br>
**Form fields (case sensitive):**

* _files_info(text)_ - This field uses text in json format as a value. The text describes the structure of the 
directories, ie, name (path), files and subdirectories. These operations can be performed using the package
<u>[github.com/uwine4850/strux_api/pkg/uplutils](https://github.com/uwine4850/strux_api/blob/master/pkg/uplutils/upload_package.go)</u>.
Below is an example of the content of this field
```
{
  "name": "tee",
  "fileNames": ["file0.txt", "file00.txt", "file000.txt"],
  "innerDir": [
  {
    "name": "tee/dir1",
    "fileNames": ["file1.txt", "file11.txt"]
  },
  {
    "name": "tee/dir2",
    "fileNames": [],
    "innerDir": [
        {
            "name": "tee/dir2/innerD",
            "fileNames": ["file1.txt"]
        },
        {
            "name": "tee/dir2/innerD2",
            "fileNames": ["file1.txt", "file11.txt"]
        }
    ]
  }]
}
```
* _username(text)_ - Username
* _password(text)_ - Password
* _version(text)_ - Package version
* _DIR$\<dirname>(files)_ - is a special field because it has a dynamic name. The name of the field corresponds to the name of the directory with
files described in the **files_info** field. The number of these fields should be equal to the number of directories with files.
Based on the example of the package structure above, the field names should be **DIR$tee**, **DIR$tee** ...
Accordingly, the value of the first field should contain the files "file0.txt", "file00.txt", "file000.txt", and the value of the second field
"file1.txt", "file11.txt".<br>
If the directory does not contain files, it does not need to be sent in this format, the information in **files_info** is enough.<br>
For convenience, it is recommended to use the package <u>[github.com/uwine4850/strux_api/pkg/uplutils](https://github.com/uwine4850/strux_api/blob/master/pkg/uplutils/upload_package.go)</u>.

**Response:** <u>baseproto.BaseResponse</u>. The response contains current information about the status of the request.

***
### Checking the availability of the package
**Address:** http://host/exists-package/ <br>
**Form fields (case sensitive):**

* _username(text)_ - Username
* _password(text)_ - Password
* _version(text)_ - Package version

**Response:** <u>baseproto.BaseResponse</u>. The response contains current information about the status of the request.

***
### Downloading a package from the service
**Address:** http://host/download-package/ <br>
**Form fields (case sensitive):**

* _username(text)_ - Username
* _pkgName(text)_ - Password
* _version(text)_ - Package version

**Response:** This command sends a form in response that is formatted as for [uploading a package to the server](#uploading-a-package-to-the-server).
That is, it contains the **files_info** field and the already familiar **DIR$** fields. All the data required to create a package on the client's machine
client's machine is passed in a form, which means that the client will have to perform similar actions as the server when uploading the package to the database.<br>
If the command operation on the server for any reason does not complete as expected, a familiar response will be returned
<u>baseproto.BaseResponse</u>.

***
### Show all versions of the package
**Address:** http://host/package-versions/ <br>
**Form fields (case sensitive):**

* _username(text)_ - Username
* _pkgName(text)_ - Password

**Response:** Microservice returns <u>pkgproto.MutateShowVersionBaseResponse</u>. The api server returns a form with a single text field
**versions**. These operations can be performed using the package <u>[github.com/uwine4850/strux_api/pkg/uplutils](https://github.com/uwine4850/strux_api/blob/master/pkg/uplutils/upload_package.go)</u>.