package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinacials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

type Human struct {
	Name string `required max:"100"`
}

type Doctor struct {
	Human
	Number     int
	Companions []string
}

type Writer interface {
	Write([]byte) (int, error)
}
type Closer interface {
	Close() error
}

// Embed interfaces
type WriterCloser interface {
	Writer
	Closer
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 0 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// Method of the struct, using pointer is an option
func (d *Doctor) Hello() (string, error) {
	if d.Name == "" {
		return "", fmt.Errorf("name is empty")
	}
	fmt.Println("Hello", d.Name)
	d.Name = d.Name + " new"
	fmt.Println("Hello", d.Name)
	fmt.Println()
	return "Hello " + d.Name, nil
}

func BasicTypes() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
	fmt.Printf("%v\n", KB)
	fmt.Printf("%v\n", MB)
	fmt.Printf("%v\n\n", GB)

	var roles byte = isAdmin | canSeeFinacials | canSeeEurope
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
	fmt.Printf("Is Admin & can see financials? %v\n", canSeeFinacials&roles == canSeeFinacials)
	fmt.Printf("%v, %v, %v, %v\n\n", roles, isAdmin, isAdmin&roles, roles&isAdmin)

	grades := [...]int{97, 85, 93, 44}
	// grades := [3]int{97, 85, 93}
	// grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v, %v\n\n", grades, len(grades))

	var matrix [3][3]int
	matrix[0] = [3]int{3, 4, 5}
	fmt.Printf("Grades: %v, %v\n\n", matrix, len(matrix))

	va1 := [...]int{1, 2, 3}
	// va2 := va1
	va2 := &va1
	va2[1] = 5
	fmt.Println(va1)
	fmt.Println(va2)
	fmt.Println()

	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice2[1] = 5
	fmt.Printf("%v, %v, length: %v, capacity: %v\n\n", slice1, slice2, len(slice1), cap(slice1))

	// sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceA := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceB := sliceA[:]
	sliceC := sliceA[3:]
	sliceD := sliceA[:6]
	sliceE := sliceA[3:6]
	sliceA[5] = 42
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n\n", sliceA, sliceB, sliceC, sliceD, sliceE)

	// exA := make([]int, 3)
	// fmt.Println(exA)
	// fmt.Printf("Length: %v\n", len(exA))
	// fmt.Printf("Campacity: %v\n\n", cap(exA))

	// slices always address reference
	exA := []int{}
	fmt.Println(exA)
	fmt.Printf("Length: %v\n", len(exA))
	fmt.Printf("Campacity: %v\n", cap(exA))
	exA = append(exA, 1, 2, 3)
	exA = append(exA, []int{1, 2, 3}...)
	fmt.Println(exA)
	fmt.Printf("Length: %v\n", len(exA))
	fmt.Printf("Campacity: %v\n", cap(exA))
}

func Map() {
	// Manipulate 1 cause changes in the initial map
	// mapA := make(map[string]int)
	mapA := map[string]int{
		"California": 111,
		"Texas":      222,
	}
	mapA["Georgia"] = 12233
	delete(mapA, "Georgia")
	if pop, found := mapA["Georgia"]; found {
		fmt.Println(pop)
	}
	fmt.Printf("%v\n\n", mapA)

	mapTest := mapA
	delete(mapTest, "Texas")
	fmt.Printf("%v\n\n", mapA)

	// mapB := map[[3]int]string{}
	// fmt.Printf("%v\n\n", mapB)
}

func Struct() {
	t := reflect.TypeOf(Human{})
	field, _ := t.FieldByName("Name")
	fmt.Printf("%v\n\n", field.Tag)

	aDoctor := Doctor{
		Number: 3,
		Human:  Human{Name: "John"},
		Companions: []string{
			"liz",
			"Jo",
			"sarah",
		},
	}
	fmt.Printf("%v\n\n", aDoctor)

	bDoctor := struct{ name string }{name: "zek"}
	fmt.Printf("%v\n\n", bDoctor.name)

	cDoctor := aDoctor
	cDoctor.Name = "Zek"
	fmt.Printf("Value remains in the initial struct: %v\n\n", aDoctor.Name)
}

func IfSwitch() {
	switch i := 2 + 3; i {
	case 1, 2, 4, 5:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}
	fmt.Println()

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than 10")
		fallthrough // no mater true or false, next case must be executed
	case i >= 20:
		fmt.Println("more than 20")
	default:
		fmt.Println("defaults")
	}
	fmt.Println()

	var j interface{} = [3]int{1, 2, 3}
	switch j.(type) {
	case int:
		fmt.Println("i is int")
	case float32:
		fmt.Println("i is float32")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("none")
	}
	fmt.Println()
}

func Loop() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
}

func panicker() {
	fmt.Println("About to panic")

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// Continue to panic
			// panic(err)
		}
	}()

	// panic() executes after defer()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func ControlFlow() {
	// defer run FILO
	defer fmt.Println("defer Start")
	defer fmt.Println("defer Middle")
	defer fmt.Println("defer End")
	defer fmt.Println()

	a := "start"
	defer fmt.Println(a)
	a = "end"

	panicker()
	fmt.Println("Panic ended")
}

func HttpExample() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // This will be executed before HttpExample close up
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func WebServiceListener() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func VarPointer() {
	var a int = 43
	var b *int = &a
	fmt.Println(a, b, *b)
	fmt.Println()

	var ms *Human = new(Human)
	ms.Name = "ZEE" // (*ms).Name = "ZEE"
	fmt.Println(ms.Name)
}

func BufferExample() {
	var wc WriterCloser = &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
	bwc := wc.(*BufferedWriterCloser) // Type conversion
	fmt.Println(bwc)
	wc.Write([]byte("Hello Youtube listeners, this is a test11"))
	wc.Close()

}
