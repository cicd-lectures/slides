// Module path
module github.com/jlevesy/prometheus-elector

// Minimum go version
go 1.21

// Dépendances directes
require (
	github.com/fsnotify/fsnotify v1.6.0
	github.com/imdario/mergo v0.3.16
	github.com/prometheus/client_golang v1.17.0
	github.com/stretchr/testify v1.8.4
	golang.org/x/net v0.15.0
	golang.org/x/sync v0.3.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.28.2
	k8s.io/client-go v0.28.2
	k8s.io/klog/v2 v2.100.1
)

// Dépendances contraintes indirectement (...)
// Certaines librairies ne supportent pas go modules, ils sont gérés comme
// dépendances directes...
// D'autres cas peuvent mener à l'ajout d'une dépendance indirecte.
require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	// [...]
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
