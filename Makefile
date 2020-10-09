PACKAGE_LIST:=go list ./... | grep -v /vendor/
FMT_PACKAGE:=$$($(PACKAGE_LIST))

all: go-mockery go-fmt go-swag unittest-all go-run

go-run:
	go run .

go-mockery:
	mockery -dir=interfaces -all -output ./interfaces/mocks

go-fmt:
	go fmt $(FMT_PACKAGE)

unittest-all:
	go test -v -count=1 ./... -coverprofile=c.out

go-swag:
	swag i