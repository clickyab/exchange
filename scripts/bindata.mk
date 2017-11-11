go-bindata:
	GOBIN=$(BIN) $(GO) get -v github.com/jteeuwen/go-bindata/go-bindata
	GOBIN=$(BIN) $(GO) install -v github.com/jteeuwen/go-bindata/go-bindata
	GOBIN=$(BIN) $(GO) get -v golang.org/x/tools/cmd/goimports
	GOBIN=$(BIN) $(GO) install -v golang.org/x/tools/cmd/goimports

$(ROOT)/contrib/IP-COUNTRY-REGION-CITY-ISP.BIN:
	mkdir -p $(ROOT)/contrib
	cd $(ROOT)/contrib && wget -c http://static.clickyab.com/IP-COUNTRY-REGION-CITY-ISP.BIN.gz
	cd $(ROOT)/contrib && gunzip IP-COUNTRY-REGION-CITY-ISP.BIN.gz
	cd $(ROOT)/contrib && rm -f IP-COUNTRY-REGION-CITY-ISP.BIN.md5 && wget -c http://static.clickyab.com/IP-COUNTRY-REGION-CITY-ISP.BIN.md5
	cd $(ROOT)/contrib && md5sum -c IP-COUNTRY-REGION-CITY-ISP.BIN.md5

$(BIN)/IP-COUNTRY-REGION-CITY-ISP.BIN: $(ROOT)/contrib/IP-COUNTRY-REGION-CITY-ISP.BIN
	cp $(ROOT)/contrib/IP-COUNTRY-REGION-CITY-ISP.BIN $(BIN)

prepare: $(ROOT)/contrib/IP-COUNTRY-REGION-CITY-ISP.BIN

codegen-fake: go-bindata
	$(BIN)/go-bindata -o commands/octopus/rtb-supplier/template.gen.go -nomemcopy=true -pkg=main commands/octopus/rtb-supplier/static/template/...
	$(BIN)/goimports -w commands/octopus/rtb-supplier/template.gen.go

.PHONY: go-bindata