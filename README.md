# dockerdeamon
Docker Machine Client serving web page GUI front end.

An Https server for the docker deamon enabling local or remote secure login.

###Installation.

Create a new directory
>user@machine:~/$ **mkdir -p $GOPATH/src/github.com/richardnwinder/dockerdeamon**

Change to the new directory cteate a ne git repository and pull the master repo from the github repository
>user@machine:~/$ **git init**
>user@machine:~/$ **git pull https://github.com/richardnwinder/dockerdeamon master**

Copy the directory DockerClient from the src directory dockerdeamon to a destination of your choosing = usually something like /home/$USER/www

######It is necessary to hard code the location of the web pages directory [DockerClient] in the main.go source code line 36.

>**const webPath string = "/home/user/www/DockerClient"**

and alter the webPath string to point to the location of your DockerClient web pages.

In order to edit and compile the source some dependancies are required
#####goerror
>user@machine:~/$ **go get github.com/richardnwinder/goerror**
#####go socket-io
>user@machine:~/$ **go get github.com/googollee/go-socket.io**

#####Compile the go executable:
>user@machine:~/$ **go make dockerdeamon**

####Install:
>user@machine:~/$ **go install dockerdeamon**

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
>user@machine:~/$ **sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout $GOPATH/ca/key.pem -out  $GOPATH/ca/cert.pem**

Copy the files key.pem and cert.pem to your $GOPATH/bin directory or the directory where your dockerdeamon executable is, from the directory where you created the certs.

###Running

Normal execution command
>user@machine:~/$ **$GOPATH/bin/dockerdeamon -admin=[admin-name] -password=[pwd]**

replace [admin-name] with char string eg - richard - any length of string is accepted but whitespace is not
replace [pwd] with any char string eg - win36qly - any length of string is accepted but whitespace is not.

###Browser
Goto https://localhost:8081/ and login with the admin-name and password.
