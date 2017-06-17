build-workers :
	$(BUILD) clickyab.com/exchange/commands/octopus/octopus-workers/octopus-demand-worker
	$(BUILD) clickyab.com/exchange/commands/octopus/octopus-workers/octopus-impression-worker
	$(BUILD) clickyab.com/exchange/commands/octopus/octopus-workers/octopus-show-worker
	$(BUILD) clickyab.com/exchange/commands/octopus/octopus-workers/octopus-winner-worker

run-worker-demand :
	$(BIN)/octopus-demand-worker

run-worker-impression :
	$(BIN)/octopus-impression-worker

run-worker-show :
	$(BIN)/octopus-show-worker

run-worker-winner :
	$(BIN)/octopus-winner-worker