CURRENT_UID = $(shell id -u):$(shell id -g)
DIST_DIR ?= $(CURDIR)/dist

REPOSITORY_URL = https://github.com/slides/cicd-lectures
PRESENTATION_URL = https://gifted-albattani-3bc8a1.netlify.app/

export PRESENTATION_URL CURRENT_UID REPOSITORY_URL

## Docker Buildkit is enabled for faster build and caching of images
DOCKER_BUILDKIT ?= 1
COMPOSE_DOCKER_CLI_BUILD ?= 1
export DOCKER_BUILDKIT COMPOSE_DOCKER_CLI_BUILD

all: clean build verify

# Prepare the Docker environment and any required dev. dependency
prepare:
	@docker-compose build

# Generate documents inside a container, all *.adoc in parallel
build: clean $(DIST_DIR) prepare
	@docker-compose up \
		--force-recreate \
		--exit-code-from build \
	build

$(DIST_DIR):
	mkdir -p $(DIST_DIR)

verify:
	@echo "Verify disabled"

serve: clean $(DIST_DIR) prepare
	@docker-compose up --force-recreate serve

shell: $(DIST_DIR) prepare
	@CURRENT_UID=0 docker-compose run --entrypoint=sh --rm serve

dependencies-lock-update: $(DIST_DIR) prepare
	@CURRENT_UID=0 docker-compose run --entrypoint=npm --rm serve install --package-lock

dependencies-update: $(DIST_DIR) prepare
	@CURRENT_UID=0 docker-compose run --entrypoint=ncu --rm serve -u
	@make -C $(CURDIR) dependencies-lock-update

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
	@docker-compose run --entrypoint=/app/node_modules/.bin/qrcode --rm serve -t png -o /app/content/media/qrcode.png $(PRESENTATION_URL)

.PHONY: all build verify serve qrcode pdf prepare dependencies-update dependencies-lock-update
