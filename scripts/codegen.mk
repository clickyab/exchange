tools-codegen:
	$(BUILD) clickyab.com/exchange/commands/codegen

octopus-user: tools-codegen
	$(BIN)/codegen -p clickyab.com/exchange/octopus/console/user/aaa
	$(BIN)/codegen -p clickyab.com/exchange/octopus/console/user/routes

octopus-report: tools-codegen
	$(BIN)/codegen -p clickyab.com/exchange/octopus/console/report/routes

codegen: $(BIN)/IP-COUNTRY-REGION-CITY-ISP.BIN migration octopus-user octopus-report
