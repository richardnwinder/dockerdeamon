# dockerdeamon
Docker Machine Client serving web page GUI front end.

An Https server for the docker deamon enabling local or remote secure login.

###Installation.

Create a new directory
>user@machine:~**mkdir $GOPATH/src/github.com/user/dockerdeamon**

Pull the master repo from the github repository
>user@machine:~**git pull https://github.com/richardnwinder/dockerdeamon master**

In order to edit and compile the source some dependancies are required
#####go socket-io
>user@machine:~**go get github.com/googollee/socket-io**

#####Compile the go executable:
>user@machine:~**go make dockerdeamon**

#####Install:
>user@machine:~**go install dockerdeamon**

Copy the directory DockerClient from the src directory to the same directory as your dockerdeamon executable or to a destination of your choosing.

#####It is necessary to hard code the location of the web pages directory [DockerClient] in the main.go source code line 36.

>**const webPath string = "/home/user/DockerClient"**

and alter the webPath string to point to the location of your DockerClient web pages.

####Generate private key (.key)

##### Key considerations for algorithm "RSA" ≥ 2048-bit
>user@machine:~**openssl genrsa -out server.key 2048**

##### Key considerations for algorithm "ECDSA" ≥ secp384r1
###### List ECDSA the supported curves (openssl ecparam -list_curves)
>user@machine:~**openssl ecparam -genkey -name secp384r1 -out server.key**

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
>user@machine:~**openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650**

###Running

Normal execution command
>user@machine:~**$GOPATH/bin/dockerdeamon -admin=[admin-name] -password=[pwd]**

replace [admin-name] with char string eg - richard - any length of string is accepted but whitespace is not
replace [pwd] with any char string eg - win36qly - any length of string is accepted but whitespace is not
