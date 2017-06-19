tools-codegen:
	$(BUILD) clickyab.com/exchange/commands/codegen

octopus-user: tools-codegen
	$(BIN)/codegen -p clickyab.com/exchange/octopus/console/user/aaa
	$(BIN)/codegen -p clickyab.com/exchange/octopus/console/user/routes

crane-user: tools-codegen
	$(BIN)/codegen -p clickyab.com/exchange/crane/models/internal/ad
	$(BIN)/codegen -p clickyab.com/exchange/crane/models/internal/campaign
	$(BIN)/codegen -p clickyab.com/exchange/crane/models/internal/publisher
	$(BIN)/codegen -p clickyab.com/exchange/crane/models/internal/user

codegen: ip2location migration octopus-user crane-user
