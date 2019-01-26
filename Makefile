
GOTOOLS = \
	github.com/golang/dep/cmd/dep \
	gopkg.in/alecthomas/gometalinter.v2 \
	zombiezen.com/go/capnproto2/... \

PACKAGES=$(shell go list ./... | grep -v '/vendor/')
SERVER=-tags 'bankserver'
CLIENT=-tags 'bankclient'
LDFLAGS= -ldflags
CAPNPPROTO=$(GOPATH)/src/zombiezen.com/go/capnproto2/std
all: tools deps build install test test_release

########################################
### Tools & dependencies
tools:
	@capnp --version || (echo "Install Capn'p first; https://capnproto.org/install.html"; false)
	@echo "Installing tools"
	go get $(GOTOOLS)
	@gometalinter.v2 --install

deps:
	@echo "Cleaning vendors..."
	rm -rf vendor/
	@echo "Running dep..."
	dep ensure -v

########################################
### Build cryptobank
build:
	go build $(SERVER) -o server/build/bankserver ./server/cmd/
	go build $(CLIENT) -o client/build/bankclient ./client/cmd/

########################################
### Testing
test:
	$go test $(PACKAGES)

capnp:
	capnp compile -I $(CAPNPPROTO) -ogo ./cryptobank/cryptobank.capnp

########################################
### Formatting, linting, and vetting
fmt:
	@go fmt ./...

metalinter:
	@echo "--> Running linter"
	@gometalinter.v2 --vendor --deadline=600s --disable-all  \
		--enable=deadcode \
		--enable=gosimple \
	 	--enable=misspell \
		--enable=safesql \
		./...
		#--enable=gas \
		#--enable=maligned \
		#--enable=dupl \
		#--enable=errcheck \
		#--enable=goconst \
		#--enable=gocyclo \
		#--enable=goimports \
		#--enable=golint \ <== comments on anything exported
		#--enable=gotype \
	 	#--enable=ineffassign \
	   	#--enable=interfacer \
	   	#--enable=megacheck \
	   	#--enable=staticcheck \
	   	#--enable=structcheck \
	   	#--enable=unconvert \
	   	#--enable=unparam \
		#--enable=unused \
	   	#--enable=varcheck \
		#--enable=vet \
		#--enable=vetshadow \



.PHONY: capnp fmt build test
.PHONY: capnp
.PHONY: build
.PHONY: tools deps
.PHONY: fmt metalinter

