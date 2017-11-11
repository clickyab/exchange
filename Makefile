export ROOT=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))
include $(ROOT)/scripts/variables.mk
all: codegen
	$(BUILD) ./...
include $(ROOT)/scripts/common.mk
include $(ROOT)/scripts/gb.mk
include $(ROOT)/scripts/linter.mk
include $(ROOT)/scripts/bindata.mk
include $(ROOT)/scripts/migration.mk
include $(ROOT)/scripts/codegen.mk
include $(ROOT)/scripts/services.mk
include $(ROOT)/scripts/cleanup.mk
include $(ROOT)/scripts/test.mk
include $(ROOT)/scripts/run.mk
