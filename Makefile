
CURRENT_DATE := $(shell date +"%Y%m%d.%H%M%S")


# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build -trimpath
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Binary name
BINARY_NAME = playground

all: clean build


build:
		rm -rf main;
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64		$(GOBUILD) main.go

clean:
		$(GOCLEAN)
		rm -rf $(OUT_DIR)

test:
		$(GOTEST) -v ./...

run:
		$(GOBUILD) cmd/$(BINARY_NAME).go
		mv ./$(BINARY_NAME) $(OUT_DIR)
		export GIN_MODE=release
		./out/$(BINARY_NAME)

deps:
		$(GOGET) "github.com/gin-gonic/gin"
		$(GOGET) "github.com/docker/docker/api/types"
		$(GOGET) "github.com/spf13/viper"
		$(GOGET) "gorm.io/gorm"
		$(GOGET) "gorm.io/driver/mysql"
		$(GOGET) "github.com/streadway/amqp"
		$(GOGET) "go.etcd.io/etcd/clientv3"
		$(GOGET) "gorm.io/driver/mysql"


