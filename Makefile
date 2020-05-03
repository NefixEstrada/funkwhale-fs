FUNKWHALE_TARGET = 0.21

WHOAMI = $(shell whoami)

.PHONY: openapi
openapi:
	rm -rf openapi
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
	--skip-validate-spec \
    -i https://dev.funkwhale.audio/funkwhale/funkwhale/-/raw/$(FUNKWHALE_TARGET)/docs/swagger.yml \
    -g go \
    -o /local/openapi
	sudo chown -R $(WHOAMI):$(WHOAMI) openapi