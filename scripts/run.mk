debug-octopus: all
	$(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/octopus

debug-fakedemand: all
	PORT=6000 $(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/fakedemand

debug-fakesupplier: all
	PORT=6000 $(BIN)/dlv --listen=:5000 --headless=true --api-version=2 exec $(BIN)/supplier

run-octopus: all
	$(BIN)/octopus

run-fakedemand: all
	$(BIN)/rtb-demand

run-fakesupplier: all
	$(BIN)/rtb-supplier

run-winworker: all
	PORT=4444 $(BIN)/octopus-winner-worker

install-debugger:
	$(GO) get -v github.com/derekparker/delve/cmd/dlv
	$(GO) install -v github.com/derekparker/delve/cmd/dlv
