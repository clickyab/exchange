debug-octopus: all
	$(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/octopus

debug-fakedemand: all
	PORT=6000 $(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/fakedemand

run-octopus: all
	$(BIN)/octopus

run-fakedemand: all
	$(BIN)/fakedemand


install-debugger:
	$(GO) get -v github.com/derekparker/delve/cmd/dlv
	$(GO) install -v github.com/derekparker/delve/cmd/dlv
