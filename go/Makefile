-include .$(env).sh

dependencies:
	go get github.com/golang/dep/cmd/dep
	dep ensure -v

	go get github.com/alecthomas/gometalinter
	gometalinter --install

spec: test-spec staging-spec production-spec

test-spec:
	APP_ENV=test MARKETPLACES=US,DE,UK FIND_CATEGORY='DE|Küche, Haushalt & Wohnen' $(MAKE) go-test

staging-spec:
	APP_ENV=staging MARKETPLACES=US FIND_CATEGORY='US|Beauty & Personal Care' $(MAKE) go-test

production-spec:
	APP_ENV=production MARKETPLACES=US,DE,UK FIND_CATEGORY='DE|Küche, Haushalt & Wohnen' $(MAKE) go-test

go-test:
	go test -v -tags=spec -race -cover ./...

analysis:
	gometalinter --vendor ./... --disable-all --enable=gofmt --enable=golint --enable=vet --enable=errcheck --exclude research/configuration_data
