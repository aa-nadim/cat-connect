

    go version

-------

    go install github.com/beego/bee/v2@latest

-------

    echo "export PATH=\$PATH:\$(go env GOPATH)/bin" >> ~/.bashrc
    source ~/.bashrc

-------

    bee version

-------

    mkdir -p ~/go/src/ 
    cd ~/go/src/             
 
-------

    bee new myproject

-------

    cd myproject

-------

    export GO111MODULE=on
    go mod init
    go mod tidy

    bee run

-------

    http://localhost:8080/

------