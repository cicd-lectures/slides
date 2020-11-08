CURRENT_UID = $(shell id -u):$(shell id -g)
DIST_DIR ?= $(CURDIR)/dist
REPOSITORY_NAME ?= slides
REPOSITORY_OWNER ?= cicd-lectures
REPOSITORY_BASE_URL ?= https://github.com/$(REPOSITORY_OWNER)/$(REPOSITORY_NAME)

REPOSITORY_URL = $(REPOSITORY_BASE_URL)
PRESENTATION_URL = https://$(REPOSITORY_OWNER).github.io/$(REPOSITORY_NAME)

ifdef GIT_TAG
REPOSITORY_URL = $(REPOSITORY_BASE_URL)/tree/$(GIT_TAG)
PRESENTATION_URL = https://$(REPOSITORY_OWNER).github.io/$(REPOSITORY_NAME)/$(GIT_TAG)
else
ifdef GIT_BRANCH
ifneq ($(GIT_BRANCH), main)
REPOSITORY_URL = $(REPOSITORY_BASE_URL)/tree/$(GIT_BRANCH)
PRESENTATION_URL = https://$(REPOSITORY_OWNER).github.io/$(REPOSITORY_NAME)/$(GIT_BRANCH)
endif
endif
endif
export PRESENTATION_URL CURRENT_UID REPOSITORY_URL REPOSITORY_BASE_URL

## Docker Buildkit is enabled for faster build and caching of images
DOCKER_BUILDKIT ?= 1
COMPOSE_DOCKER_CLI_BUILD ?= 1
export DOCKER_BUILDKIT COMPOSE_DOCKER_CLI_BUILD

all: clean build verify

# Generate documents inside a container, all *.adoc in parallel
build: clean $(DIST_DIR)
	@docker-compose up \
		--build \
		--force-recreate \
		--exit-code-from build \
	build

$(DIST_DIR):
	mkdir -p $(DIST_DIR)

verify:
	@echo "Verify disabled"

serve: clean $(DIST_DIR)
	@docker-compose up --build --force-recreate serve

shell: $(DIST_DIR)
	@CURRENT_UID=0 docker-compose run --entrypoint=sh --rm serve

$(DIST_DIR)/index.html: build

pdf: $(DIST_DIR)/index.html
	@docker run --rm -t \
		-v $(DIST_DIR):/slides \
		--user $(CURRENT_UID) \
		astefanutti/decktape:2.9 \
		/slides/index.html \
		/slides/slides.pdf \
		--size='2048x1536'

clean:
	@docker-compose down -v --remove-orphans
	@rm -rf $(DIST_DIR)

qrcode:
	@docker-compose run --entrypoint=/app/node_modules/.bin/qrcode --rm serve -t png -o /app/content/images/qrcode.png $(PRESENTATION_URL)

.PHONY: all build verify serve qrcode pdf
