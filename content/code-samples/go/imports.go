// tag::import[]
package main

import (
	// Imports de la librairie standard.
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	// Import du projet courant.
	"mondomaine.com/monprojet/pkg/helpers"

	// Imports de librairie externes (dépendances).
	"github.com/prometheus/client_golang/prometheus"
)

// end::import[]
func main() {

	io.Reader
	fmt.Println

	http.Handler
	prometheus.NewLogger
	tls.Config
	helpers.Foo
}
