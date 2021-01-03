.PHONY: all
all: clean dist test

main.html:
	asciidoctor main.adoc

.PHONY: dist
dist: main.html
	mkdir -p ./dist
	cp ./main.html ./dist/index.html

.PHONY: clean
clean:
	rm -rf ./dist/ ./main.html

.PHONY: test
test: main.html
	linkchecker --check-extern ./main.html
