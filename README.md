# dockerdeamon
Docker Machine Client serving web page GUI front end.

An Https server for the docker deamon enabling local or remote secure login.

####Installation.

Clone the git directory to your local machine

In order to edit and compile the source some dependancies are required
#####go socket-io
>user@machine:~**go get github.com/googollee/socket-io**

#####Compile the go executable:
>user@machine:~**go make dockerdeamon**

#####Install:
>user@machine:~**go install dockerdeamon**

Copy the directory DockerClient from the src directory to the same directory as your dockerdeamon executable or to a destination of your choosing.

#####It is necessary to hard code the location of the web pages directory [DockerClient] in the main.go source code line 36.

>const webPath string = "/home/user/DockerClient"

alter the path string to point to the location of your webpages.
