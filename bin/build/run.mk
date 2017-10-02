debug-octopus: all
	$(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/octopus

debug-fakedemand: all
	PORT=6000 $(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/fakedemand

debug-fakesupplier: all
	PORT=6000 $(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/supplier

run-octopus: all
	$(BIN)/octopus

run-fakedemand: all
	PORT=9898 $(BIN)/fakedemand

run-fakesupplier: all
	$(BIN)/supplier

install-debugger:
	$(GO) get -v github.com/derekparker/delve/cmd/dlv
	$(GO) install -v github.com/derekparker/delve/cmd/dlv
