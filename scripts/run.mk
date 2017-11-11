debug-octopus: all
	$(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/octopus

debug-fakedemand: all
	PORT=6000 $(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/fakedemand

debug-fakesupplier: all
	PORT=6000 $(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/supplier

run-octopus: all
	PORT=8091 $(BIN)/octopus

run-fakedemand: all
	PORT=9898 $(BIN)/rtb-demand

run-fakesupplier: all
	PORT=3500 $(BIN)/rtb-supplier

install-debugger:
	$(GO) get -v github.com/derekparker/delve/cmd/dlv
	$(GO) install -v github.com/derekparker/delve/cmd/dlv
