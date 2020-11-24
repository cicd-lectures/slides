# tag::all[]
.PHONY: all
all: clean example-make.tar.gz

# tag::build[]
index.html:
  asciidoctor index.adoc --out-file=$(CURDIR)/index.html
  # $(CURDIR) est une variable Make désignant le chemin vers le dossier contenant le Makefile

# end::build[]
# tag::package[]
example-make.tar.gz: index.html
  # czf: Create Ze (gZipped) File
  tar czf $(CURDIR)/example-make.tar.gz $(CURDIR)/index.html

# end::package[]
# tag::clean[]
.PHONY: clean
clean:
  rm -f $(CURDIR)/example-make.tar.gz $(CURDIR)/index.html

# end::clean[]
# end::all[]
