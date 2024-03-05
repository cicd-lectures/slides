package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// tag::main_start[]
func main() {
	// end::main_start[]

	// tag::variables[]
	// Declare une variable de type string et assigne à la valeur par defaut du type.
	// "" pour string.
	var message string
	// Assigne (copie) la valeur "Hello ENSG!" a la variable message.
	message = "Hello ENSG!"

	// Est équivalent a:
	var message string = "Hello ENSG!"

	// Est encore équivalent a... avec une syntaxe compacte.
	// Ici le compilateur devine le type de la variable en fonction de la valeur assignée.
	message := "Hello ENSG!"

	// Affiche la valeur de la variable message dans la sortie standard.
	fmt.Println(message)
	// end::variables[]

	// tag::variables_bad[]
	// Déclare et initialise une variable message de type string.
	message := "Hello ENSG!"
	// Déclare et initialise une variable age de type int.
	age := 43

	// Assigne la valeur age dans la variable message.
	// ❌ Ne compile pas!
	// message: cannot use age (variable of type int) as string value in assignment.
	message = age
	// end::variables_bad[]

	// tag::pointers_ref[]
	// On déclare et initialise une variable. Cela aloue de la mémoire sur la pile.
	var message string = "Hello ENSG!"
	// On copie l'adresse memoire de cette variable dans une nouvelle variable.
	// Pour cela on utilise l'opérateur & (réference).
	var pointerToMessage *string = &message

	// Affiche: message address in memory is: 0xc000014070
	fmt.Println("message address in memory is:", pointerToMessage)
	// end::pointers_ref[]

	// tag::pointers_deref[]
	var message string = "Hello ENSG!"
	var pointerToMessage *string = &message

	// Affiche: message is: Hello ENSG!"
	fmt.Println("message is:", *pointerToMessage)
	// end::pointers_deref[]

	// tag::pointers_panic[]
	// Le `= nil` est optionel ici: la valeur par défaut d'un pointeur est nil.
	var nilPointer *string = nil

	fmt.Println("address is:", nilPointer)
	// A votre avis: que fait cette ligne?
	fmt.Println("message is:", *nilPointer)
	// end::pointers_panic[]
	// tag::main_end[]
}

// end::main_end[]

// tag::package_visibility[]

// fonction privée du package courant. Ne peut pas être utilisée a l'extérieur.
func privateFunc() {
	// appelle la fonction `readContent` du package os.
	// ❌ Ne compile pas: `readContent` n'est pas exportée.
	content := os.readContent("/some/file")
}

// fonction publique du package courant.
// Peut être appellée depuis un autre package.
func PublicFunc() {
	// appelle la fonction `OpenFile` du package os.
	// ✅ compile, `OpenFile` est exportée!
	file, _ := os.OpenFile("/some/file")
}

// Variable de package publique.
var PublicVar int = 12

// end::package_visibility[]

// tag::var_scopes[]
func doSomething() {
	var age int64

	age = readAge()

	// ❌ Ne compile pas: newAge n'est pas définie dans ce scope.
	newAge = 45

	fmt.Println(age)
}

func readAge() int64 {
	// ❌ Ne compile pas. age n'est pas définie dans ce scope.
	age = 42

	newAge := readAgeFile()

	return newAge
}

// end::var_scopes[]

// tag::func_examples_1[]
// Un fonction qui accepte une string et ne retourne rien.
func sayHello(message string) {
	fmt.Println("Hello:", message)
}

// Une fonction qui accepte deux entiers et qui retourne un float64 et une erreur.
func divide(numerator, denominator int) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("can't divide by 0")
	}

	return numerator / denominator, nil
}

// end::func_examples_1[]

// tag::func_defer[]
func faireDeLaPolitique() {
	rendreLargent := func() {
		fmt.Println("Argent rendu")
	}

	// Quoi qu'il advienne, l'argent sera rendu.
	// Peu importe le résultat des élections.
	defer rendreLargent()

	if elu := elections(); elu {
		fmt.Println("Je suis élu")
		return
	}

	fmt.Printn("Je ne suis pas élu :(")
}

// end::func_defer[]

// tag::func_examples_2[]
// Une fonction qui accepte une chaine de caractères
// ... et qui retourne une fonction qui n'accepte aucun argument
// mais qui retourne une chaine de caractères.
func messWithFuncs(name string) func() string {
	// Les fonctions peuvent etre manipulées comme des valeurs!
	fn := func() string {
		return "Hello " + name
	}

	return fn
}

// end::func_examples_2[]

// tag::func_pass_by_value[]
func main() {
	name := "John"

	addGreeting(name)

	// Affiche "John" et non "Hello John".
	fmt.Println(name)
}

func addGreeting(name string) {
	name = "Hello " + name
}

// end::func_pass_by_value[]

// tag::func_pass_by_reference[]
func main() {
	name := "John"

	// On passe en argument de addGreeting l'adresse de la variable `name`.
	addGreeting(&name)

	fmt.Println(name)
}

func addGreeting(namePtr *string) {
	// La valeur de la variable référencée par namePtr égale à
	// la chaine de caractères "Hello " concaténée avec
	// la valeur de la variable référencée par namePtr.
	*namePtr = "Hello " + *namePtr
}

// end::func_pass_by_reference[]

// tag::flow_control_if_1[]

func main() {
	// if / else classique.
	ok := doSomething()
	if ok {
		fmt.Println("C'est OK!")
	} else {
		fmt.Println("C'est pas OK!")
	}
}

func doSomething() bool {
	return true
}

// end::flow_control_if_1[]

// tag::flow_control_if_2[]

func main() {
	// if / else avec short statement.
	// avantage: ok n'exsite que dans le scope du if.
	if ok := doSomething(); ok {
		fmt.Println("C'est OK!")
	} else {
		fmt.Println("C'est pas OK!")
	}

	// Ne compile pas: ok n'est pas défini.
	ok = true
}

func doSomething() bool {
	return true
}

// end::flow_control_if_2[]

// tag::flow_control_switch[]
func main() {
	// switch
	age := readAge()
	switch age {
	case 10:
		fmt.Println("Hello 10")
	case 42:
		fmt.Println("Hello 42")
	default:
		fmt.Prinln("Hello darkness my old friend")
	}
}

func readAge() int {
	return 42
}

// end::flow_control_switch[]

// tag::flow_control_for[]
func sum0to9() {
	var total int

	for i := 0; i < 10; i++ {
		total += i
	}

	fmt.Println("Total", total)
}

// end::flow_control_for[]

// tag::arrays[]
func main() {
	// Declare et initialise un tableau de 2 entiers.
	var intArray [2]int
	// On peut assigner un élément du tableau en utilisant son index.
	intArray[0] = 1
	intArray[1] = 3
	// On accède a un élément du tableau en utilisant son index.
	fmt.Println(intArray[0], intArray[1])

	anotherArray := [4]int{2, 4, 6, 8}
	// ❌ Ne compile pas: la taille fait partie du type!
	// On assigne un tableau de 4 entrées a un tableau de deux entrées
	intArray = anotherArray
}

// end::arrays[]

// tag::slices[]
func main() {
	anArray := [4]int{2, 4, 6, 8}

	// Declare et initialise une slice référençant les entrées
	// entre l'index 1 et 3 du tableau anArray.
	// Se lit interval [1:4[, du coup 1,2 et 3.
	var aSlice []int = anArray[1:4]

	// ⚠️ Une ecriture écrit une valeur dans le tableau référencé!
	aSlice[0] = 9
	fmt.Println(aSlice)  // [9, 6, 8]
	fmt.Println(anArray) // [2, 9, 6, 8]
}

// end::slices[]

// tag::slices_literals[]
func main() {
	aSlice := []int{2, 4, 6, 8}
	// Sélectionne les entrées entre l'index 2 et 3 de la slice aSlice.
	anotherSlice := aSlice[2:4]
	fmt.Println(aSlice)       // [2, 4, 6, 8]
	fmt.Println(anotherSlice) // [6, 8]

	// Initialise une slice de strings de 3 entrées.
	yetAnotherSlice := make([]string, 3)
	fmt.Println(yetAnotherSlice) // ["", "", ""]
}

// end::slices_literals[]

// tag::slices_append[]
func main() {
	// On ajoute l'entrée 10 a la slice  aSlice
	aSlice := []int{2, 4, 6, 8}
	aSlice = append(aSlice, 10)
	fmt.Println(aSlice) // [2, 4, 6, 8, 10]

	// On ajoute tous les items de la `anotherSlice` a la slice `aSlice`
	// Et on assigne le résultat à la variable yetAnotherSlice
	// Notez les  "..."
	anotherSlice := []int{10, 12, 14, 16}
	yetAnotherSlice := append(aSlice, anotherSlice...)
	fmt.Println(yetAnotherSlice) // [2, 4, 6, 8, 10, 12, 14, 16]
}

// end::slices_append[]

// tag::slices_len_cap[]
func main() {
	sliceOne := []int{0, 1, 2, 3}
	sliceTwo := sliceOne[0:2]
	// Affiche "Length: 2 Capacity: 4"
	fmt.Println("Length: ", len(sliceTwo), "Capacity: ", cap(sliceTwo))
}

// end::slices_len_cap[]

// tag::slices_nil[]
func main() {
	var nilSlice []string

	// panic!: on accède a un tableau qui n'existe pas.
	v := nilSlice[0] // 💥

	fmt.Println(len(nilSlice), cap(nilSlice)) // 0, 0

	nilSlice = append(nilSlice, "foo", "bar", "biz")
	fmt.Println(nilSlice) // ["foo", "bar", "biz"]
}

// end::slices_nil[]

// tag::slices_range[]
func main() {
	slice := []int{2, 4, 6, 8}

	// Affiche:
	// Index: 0 Value: 2
	// Index: 1 Value: 4
	// Index: 2 Value: 6
	// Index: 3 Value: 8
	for index, value := range slice {
		fmt.Println("Index: ", index, "Value: ", value)
	}
}

// end::slices_range[]

// tag::slices_int_to_string[]
func main() {
	input := []int{1, 2, 3, 4}
	output := toStringSlice(input)
	fmt.Println(output)
}

func toStringSlice(input []int) []string {
	// On alloue une slice de string de la taille de la slice d'entiers donnée en paramètre.
	result := make([]string, len(input))

	// Pour chaque entrée de la slice input...
	for i, v := range input {
		// On ecrit le resultat de la conversion
		// dans la slice de resultat a l'index courant.
		result[i] = strconv.Itoa(v)
	}

	return result
}

// end::slices_int_to_string[]

// tag::maps[]
func main() {
	// Déclaration et initialisation d'une map de façon littérale.
	mapAges := map[string]int{
		"Julien": 35,
		"Damien": 36,
	}

	// Déclaration et initialisation d'une map de taille 2.
	mapVilles := make(map[string]string, 2)
	// Ecritures des valeurs dans la map.
	mapVilles["Julien"] = "Lyon"
	mapVilles["Damien"] = "St-Etienne"

	var nilMap map[int]int
	nilMap[21] = 42 // panic! écriture dans une map qui n'est pas instanciée

	// On peut suprimer une entrée d'une map
	delete(mapVilles, "Julien")

	// Affiche 2, 1, 0.
	fmt.Println(len(mapAges), len(mapVilles), len(nilMap))
}

// end::maps[]

// tag::maps_reads[]
func main() {
	// Déclaration et initialisation d'une map de façon littérale.
	mapAges := map[string]int{
		"Julien": 35,
		"Damien": 36,
	}

	// Lecture sans vérification.
	// Si la clé existe, retourne la valeur associée.
	// Si la clé n'existe pas, retoure la valeur par défaut du type de la valeur.
	ageJulien := mapAges["Julien"]

	fmt.Println("Age de Julien", ageJulien)

	// Lecture avec vérification.
	// Si la clé existe, la valeur sera retournée, et ok sera a true
	// Si la clé n'existe pas, ok sera false.
	ageMichel, ok := mapAges["Michel"]
	if !ok {
		fmt.Println("Pas d'age pour Michel")
	} else {
		fmt.Println("Age de Michel", ageMichel)
	}
}

// end::maps_reads[]

// tag::maps_range[]
func main() {
	mapAges := map[string]int{
		"Julien": 35,
		"Damien": 36,
	}

	// Affiche soit:
	// Julien a 35 ans
	// Damien a 35 ans
	// OU
	// Damien a 35 ans
	// Julien a 35 ans
	for name, age := range mapAges {
		fmt.Printf("%s a %d ans\n", name, age)
	}
}

// end::maps_range[]

// tag::word_count[]
func main() {
	input := "The quick quick brown fox jumps over the lazy lazy dog"

	result := WordCount(input)

	fmt.Println(result)
}

func WordCount(input string) map[string]int {
	result := make(map[string]int)

	for _, word := range strings.Fields(input) {
		result[word]++
	}

	return result
}

// end::word_count[]

// tag::structs[]
// Déclaration du type lecture, composé de 3 attributs.
type Lecture struct {
	Topic    string
	Duration time.Duration
	Credits  int
}

func main() {
	// On declare et initialise une nouvelle variable de type Lecture.
	coursCICD := Lecture{
		Topic:    "CICD",
		Duration: 3 * 6 * time.Hour,
		Credits:  2,
	}
	// On prends la référence de la variableCICD
	var ptrVersCoursCICD *Lecture = &coursCICD

	// On accède aux valeurs des membres de la variable coursCICD avec `.`.
	fmt.Println("Sujet du cours", coursCICD.Topic)
	// `.` Fonctionne aussi sur un pointeur!
	fmt.Println("Durée du cours", ptrVersCoursCICD.Duration)
}

// end::structs[]

// tag::struct_defaults[]
// Déclaration du type lecture, composé de 4 attributs.
type Lecture struct {
	Topic    string
	Duration time.Duration
	Credits  int
	// attribut secret, seulement accessible dans le package courant.
	secret string
}

func main() {
	coursVide := Lecture{
		Topic:    "",
		Duration: time.Duration(0),
		Credits:  0,
	}
	coursDéfaut := Lecture{}

	if coursVide == coursDéfaut {
		// Affiche OK.
		fmt.Println("OK")
	}
}

// end::struct_defaults[]

// tag::struct_initialization[]
// Déclaration du type lecture, composé de 4 attributs.
type FileReader struct {
	File *os.File
}

func NewFileReader(path string) (*FileReader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &FileReader{
		File: file,
	}
}

func main() {
	// ❌ instanciation sans utiliser la fonction d'initialisation.
	var reader FileReader

	// ✅
	reader, err := NewFileReader("/some/file")
}

// end::struct_initialization[]

// tag::types[]
// Déclare un type Color représenté par un entier.
type Color int

const (
	ColorBlue  = 0
	ColorGreen = 1
)

// Déclare un type Car représenté par une structure composé de trois attributs.
type Car struct {
	Color   Color
	Engine  Engine
	Battery Battery
}

// Déclare un type Garage, qui est une map entre le nom du propriétaire et la voiture.
type Garage map[string]Car

// Déclare un type CarOption, qui est une fonction qui accepte un pointeur sur Car et retourne une erreur.
type CarOption func(*Car) error

// end::types[]

// tag::methods_1[]
// Définit un type Color représenté en mémoire par un entier.
type Color int

// Attache une méthode String a toute instance de la valeur Color
// qui retourne le nom de la couleur sous forme de chaine de Caractères.
func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "blue"
	case ColorGreen:
		return "green"
	default:
		return "unknown"
	}
}

// Bloc de constantes déclarant les couleurs possibles
const (
	ColorBlue  Color = 1
	ColorGreen Color = 2
)

func main() {
	color := ColorBlue
	fmt.Prntln("La couleur est:", color.String())
}

// end::methods_1[]
// tag::methods_2[]

type Car struct {
	Brand string
	Color Color
}

// Attache une méthode a toute instance de type "pointeur sur Car".
// Le premier argument avant le nom de la méthode est appelé "receveur".
func (c *Car) Describe() {
	fmt.Printf("Car brand is: %s, car color is %s\n", c.Brand, c.Color.String())
}

func main() {
	car := Car{
		Brand: "Renault",
		Color: ColorBlue,
	}

	car.Describe()
}

// end::methods_2[]

// tag::polymorphism[]
type Vehicle interface {
	Ride()
}

type Scooter struct{}

func (s *Scooter) Ride() {
	fmt.Println("Riding a Scooter")
}

type Bicycle struct{}

func (b *Bicycle) Ride() {
	fmt.Println("Ride a Bicycle")
}

func main() {
	// La variable vehicle peut recevoir soit un Scooter, soit un Bicycle.
	// Ces deux types satisfont l'interface `Vehicle`.
	var vehicle Vehicle

	vehicle = &Scooter{}
	// Affiche "Riding a Scooter".
	vehicle.Ride()

	vehicle = &Bicycle{}
	// Affiche "Rinding a Bicycle".
	vehicle.Ride()
}

// end::polymorphism[]

// tag::conversions[]
func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
}

// end::conversions[]

// tag::io_example[]
// writeHello écrit hello dans n'importe quelle destination du moment qu'elle satisfait `io.Writer`
func writeHello(dest io.Writer) {
	dest.Write([]byte("hello"))
}

func main() {
	var buf bytes.Buffer

	// Ici on écrit dans un buffer en mémoire.
	writeHello(&buf)

	file, _ := os.Open("./file")
	defer file.Close()

	// Ici on l'écrit dans un ficher.
	writeHello(file)
}

// end::io_example[]

// tag::error_type[]
type error interface {
	Error() string
}

// end::error_type[]

// tag::error_example[]
func main() {
	file, err := os.Open("/super/file")
	if err != nil {
		// Si err est non nil, alors l'opération à échouée,
		fmt.Println("Impossible d'ouvrir le fichier", err)
		return
	}
	// On s'assure de toujours fermer le fichier ouvert.
	defer file.Close()

	// On peut intéragir avec le fichier!
}

// end::error_example[]

// tag::decode_json[]
// tag::struct_annot[]
type City struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// {"id": "1", "name": "Lyon"}

// end::struct_annot[]

const payload = `[{"id":"1","name":"Lyon"},{"id":"2","name":"Paris"}]`

func main() {
	var cities []City

	if err := json.Unmarshal([]byte(payload), &cities); err != nil {
		fmt.Println("cannot unmarshal", err)
		return
	}

	fmt.Println(cities)
}

// end::decode_json[]

// tag::read_buffer[]
func main() {
	var buf bytes.Buffer

	readBytes, err := io.ReadAll(&buf)
	if err != nil {
		// KO.
	}
	// OK!
}

// end::read_buffer[]

// tag::http_request[]
func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		// Handle err...
	}
	if resp.StatusCode != http.StatusOK {
		// Code non OK. Ouch.
	}

	// On va accéder au Body.
	// Il faut s'assurer de toujours le fermer.
	defer resp.Body.Close()
}

// end::http_request[]

// tag::tatooine_climate[]
type Planet struct {
	Climate string `json:"climate"`
}

func main() {
	resp, err := http.Get("https://swapi.dev/api/planets/1")
	if err != nil {
		fmt.Println("Cannot query swapi", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Bad status", resp.Status)
		return
	}

	defer resp.Body.Close()

	readBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Can't read body", err)
		return
	}

	var tatooine Planet
	if err := json.Unmarshal(readBody, &tatooine); err != nil {
		fmt.Println("Can't unmarshal payload", err)
		return
	}

	fmt.Println("Tatooine's climate is", tatooine.Climate)
}

// end::tatooine_climate[]

// tag::consts[]
const apiURL = "https://swapi.dev/vehicles/4"

func main() {
	const anotherConst = 4

	fmt.Println(apiURL, anotherConst)
}

// end::consts[]
