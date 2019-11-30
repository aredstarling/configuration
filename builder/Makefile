dependencies:
	go get -u github.com/jteeuwen/go-bindata/...
	go get github.com/cloudfoundry-incubator/spiff

product-research:
	./escape.sh
	./create.sh
	./unescape.sh
	go-bindata -o ../go/research/configuration_data.go -pkg research ../product_research
