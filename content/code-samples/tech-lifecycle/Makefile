# tag::all[]
.PHONY: all
all: clean dist

# tag::build[]
main.html:
	asciidoctor main.adoc

# end::build[]
# tag::dist[]
.PHONY: dist
dist: main.html
	mkdir -p ./dist
	cp ./main.html ./dist/index.html

# end::dist[]
# tag::clean-v1[]
.PHONY: clean
clean:
	rm -f ./main.html

# end::clean-v1[]
# tag::clean-v2[]
.PHONY: clean
clean:
	rm -rf ./dist/ ./main.html

# end::clean-v2[]
# end::all[]
