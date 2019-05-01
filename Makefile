BIN_DIR := $(GOPATH)/bin
GO=$(shell which go)

install:
	$(GO) get -u "github.com/sirupsen/logrus"
	$(GO) get -u "github.com/gorilla/mux"
	$(GO) get -u github.com/mongodb/mongo-go-driver

depends:
	dep init
	dep ensure
	# dep ensure -add "go.mongodb.org/mongo-driver/mongo@~1.0.0"
	# dep ensure -add "github.com/gorilla/mux"
	# dep ensure -add "github.com/sirupsen/logrus"



setup: depends