tools-migrate: migration
	$(BUILD) clickyab.com/exchange/commands/migration

octopus_migup: tools-migrate
	$(BIN)/migration -action=up -app=octopus

octopus_migdown: tools-migrate
	$(BIN)/migration -action=down -app=octopus

octopus_migdown-all: tools-migrate
	$(BIN)/migration -action=down-all -app=octopus

octopus_migredo: tools-migrate
	$(BIN)/migration -action=redo -app=octopus

octopus_miglist: tools-migrate
	$(BIN)/migration -action=list -app=octopus

migcreate:
	@/bin/bash $(BIN)/create_migration.sh

migration: go-bindata
	cd $(ROOT) && $(BIN)/go-bindata -nometadata -o ./commands/migration/migration.gen.go -nomemcopy=true -pkg=main ./db/...
	$(BIN)/goimports -w ./commands/migration/migration.gen.go
