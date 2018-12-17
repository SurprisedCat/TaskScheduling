#! /bin/bash
GOPATH=`go env GOPATH`
if [ $? -ne 0 ]
then
    echo "go env GOPATH executing FAILED" 
fi
echo "Start to install packges"
echo "GFW"
echo "go get -u github.com/golang/exp/rand"
go get -u github.com/golang/exp/rand
cp -r $GOPATH/src/github.com/golang/exp $GOPATH/src/golang.org/x/
if [ $? -ne 0 ]
then
    echo "cp -r $GOPATH/src/github.com/golang/exp $GOPATH/src/golang.org/x/ FAILED" 
fi
echo "go get -u -t gonum.org/v1/gonum/..."
go get -u -t gonum.org/v1/gonum/...
echo "go test gonum.org/v1/..."
go test gonum.org/v1/...
