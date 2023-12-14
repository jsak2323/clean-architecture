# Clean Architecture in Golang
This is a sample repository for a Golang project using the Clean Architecture pattern.

# Project Structure
The project structure follows the Clean Architecture pattern, where the domain layer contains the business logic and the application layer contains the use cases. The infrastructure layer is responsible for the implementation details of the application, such as database access and external APIs.

## Getting Started
Prerequisites
Golang 1.15 or later
PostgreSQL
# InstallationW
## Install Go
    cd /tmp
    wget https://dl.google.com/go/go1.18.2.linux-amd64.tar.gz

    sudo tar -xvf go1.18.2.linux-amd64.tar.gz
    sudo mv go /usr/local

    sudo nano ~/.bashrc
    // add to end of file
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/$GO_WORKSPACE/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

    source ~/.bashrc
    source ~/.profile

    go version
    // should return "go version go1.18.2 linux/amd64"

## Clone App from Git
    cd $GOPATH
    mkdir src && cd src

    git clone git@gitlab.com:saham-rakyat/clean-architecture.git
    cd clean-architecture
    mkdir logs

## Run HTTP App
    go mod init

    // build app
    go build cmd/*.go run-server

    // setup database

    // run app
    ./main


License
This project is licensed under the MIT License - see the LICENSE file for details.


