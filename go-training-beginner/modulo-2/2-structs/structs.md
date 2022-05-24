# Structs

![structs](/go-training-beginner/modulo-2/structs/img/structs.png)

## ¿Qué es una estructura (struct)?

Es un tipo de dato definido por el usuario (programador) que contiene una serie de campos con sus respectivos nombres. Entonces estas estructuras son útiles para agrupar datos relacionados bajo una misma entidad o unidad.

> En este punto, si venidos de la programación orientada a objetos, podemos comparar a una estructura con una pequeña clase que puede extender de otra pero que no soporta herencia.

## Definición

```go
type Person struct {}
```

`type` y `struct` son palabras reservadas del lenguaje

Ahora le agreguemos "campos con nombre"

```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() { ... }
```
 Si los campos son del mismo tipo, podemos agruparlos. 

 ```go
 type Person struct {
	FirstName, LastName string
	Age       int
}

func main() { ... }
```

>No es muy limpio, a la hora de leer código, agrupar los campos.

## Declaración e inicialización

Ahora que tenemos definida nuestra estructura (Person) podemos declararla como cualquier otro tipo de dato.

```go
func main() {
	var p1 Person

	fmt.Println("Person 1:", p1)
}
```
```cmd
$ go run structs.go

Person 1: {  0}
```

Como podemos ver, todos los campos de la estructura se inicializan con sus `zero values`. Por eso `FirstName` y `LastName` son un string vacío `""` y `Age` es `0`.

Podemos inicializar una estructura en forma "literal"

```go
func main() {
	var p1 Person

	fmt.Println("Person 1:", p1)

	var p2 = Person{
        FirstName: "Peter", 
        LastName: "Parker", 
        Age: 22,
    }

	fmt.Println("Person 2:", p2)
}
```

```cmd 
$ go run structs.go

Person 1: {  0}
Person 2: {Peter Parker 22}
```

O solo inicializar algunos campos

```go
func main() {
	var p1 Person

	fmt.Println("Person 1:", p1)

	var p2 = Person{
		FirstName: "Peter",
		LastName:  "Parker",
		Age:       22,
	}

	fmt.Println("Person 2:", p2)

	var p3 = Person{
		FirstName: "Tony",
		LastName:  "Stark",
	}

	fmt.Println("Person 3:", p3)
}
```

```cmd
$ go run structs.go

Person 1: {  0}
Person 2: {Peter Parker 22}
Person 3: {Tony Stark 0}
```
>Vemos como el campo edad de persona 3 tiene por defecto 0

También podemos inicializar una estructura sin los nombres de los campos

```go
func main() {
	var p1 Person

	fmt.Println("Person 1:", p1)

	var p2 = Person{
		FirstName: "Peter",
		LastName:  "Parker",
		Age:       22,
	}

	fmt.Println("Person 2:", p2)

	var p3 = Person{
		FirstName: "Tony",
		LastName:  "Stark",
	}

	fmt.Println("Person 3:", p3)

	var p4 = Person{"Bruce", "Wayne"}

	fmt.Println("Person 4:", p4)
}
```

`Pero...`
```cmd
$ go run structs.go

./structs.go:31:34: too few values in Person{...}
```

>En este caso debemos darle valor a todos los campos de la estructura

```go
	var p4 = Person{"Bruce", "Wayne", 40}

	fmt.Println("Person 4:", p4)
```

Por último, podemos declarar una estructura anónima

```go
func main() {
	var p1 Person

	fmt.Println("Person 1:", p1)

	var p2 = Person{
		FirstName: "Peter",
		LastName:  "Parker",
		Age:       22,
	}

	fmt.Println("Person 2:", p2)

	var p3 = Person{
		FirstName: "Tony",
		LastName:  "Stark",
	}

	fmt.Println("Person 3:", p3)

	var p4 = Person{"Bruce", "Wayne", 40}

	fmt.Println("Person 4:", p4)

	var a = struct {
		Name     string
		IsGreat bool
	}{"Golang", true}

	b := struct {
		Name     string
		IsGreat bool
	}{"PHP", false}

	fmt.Println("Anonymous:", a)
	fmt.Println("Anonymous:", b)


	fmt.Println("Anonymous:", a)
}
```

## Acceso a los campos

Para tener acceso a los campos de forma individual podemos hacerlo de la siguiente manera:

```go
func main() {
	var p = Person{
		FirstName: "Peter",
		LastName:  "Parker",
		Age:       22,
	}

	fmt.Println("FirstName", p.FirstName)
}
```
También podemos crear un puntero a una estructura

```go 
func main() {
	var p = Person{
		FirstName: "Peter",
		LastName:  "Parker",
		Age:       22,
	}

	ptr := &p

	fmt.Println((*ptr).FirstName)
	fmt.Println(ptr.FirstName)
}
```
>Ambas declaraciones son iguales ya que ya que en Go no necesitamos desreferenciar explícitamente un puntero. :exploding_head:

O podemos usar `new` como "constructor"

```go
func main() {
	p := new(Person)

	p.FirstName = "Peter"
	p.LastName = "Parker"
	p.Age = 22

	fmt.Println("Person", p)
}
```

```cmd 
$ go run structs.go

Person &{Peter Parker 22}
```

## Campos exportados

Al igual en las variables y funciones, si un campo se declara con su nombre empezando con minúsculas, no se exportara y solo será visible para el paquete definido.

```go
type Person struct {
	FirstName, LastName  string
	Age                  int
	zipCode              string
}
```
En este caso, `zipCode` es un campo que no se exportará. Lo mismo ocurre con la la estructura misma, si su nombre empieza con minúscula, tampoco se exportara.

```go
type person struct {
	FirstName, LastName  string
	Age                  int
	zipCode              string
}
```

### Embedding and composition

Como se mencionó antes, Go no admite herencia, pero podemos simularla embebiendo una estructura en otra

```go
type Person struct {
	FirstName, LastName  string
	Age                  int
}

type SuperHero struct {
	Person
	Alias string
}
```
De esta manera, nuestra estructura `SuperHero` "hereda" todo de `Person` y debería comportarse como cualquier estructura

```go
func main() {
	s := SuperHero{}

	s.FirstName = "Bruce"
	s.LastName = "Wayne"
	s.Age = 40
	s.Alias = "batman"

	fmt.Println(s)
}
```
```cmd
$ go run structs_embbeding.go

{{Bruce Wayne 40} batman}
```

Sin embargo, esto no es recomendable y en la mayoría de los casos se prefiere lo que se llama en Go, `composición`

```go
type Person struct {
	FirstName, LastName  string
	Age                  int
}

type SuperHero struct {
	RealName Person
	Alias  string
}
```
De esta manera podemos reescribir el código

```go
func main() {
	p := Person{"Bruce", "Wayne", 40}
	s := SuperHero{p, "batman"}

	fmt.Println(s)
}
```
```cmd
$ go run structs_composition.go

{{Bruce Wayne 40} batman}
```

>Hay lindos debates con respecto a este tema. En mi humilde opinión, no deberíamos preocuparnos por si es correcto uno u otro, cada uno puede optar por la solución que le mas útil.

## Struct tags

Un tag es una etiqueta que nos permite adjuntar información de metadatos al campo que puede usarse para un comportamiento específico.

```go
type Animal struct {
	Name    string `key:"value1"`
	Age     int    `key:"value2"`
}
```

Estas etiquetas las vamos a ver cuando usemos paquetes de codificación como JSON, XML, YAML, ORM, etc.

Por ejemplo, si estamos creando una api y exponiendo un endpoint, podríamos estar esperando un request que en su body contenga un JSON. Y si quisiéramos "emparejar" o decodificar ese JSON en nuestra estructura, podríamos hacer algo como

```go
type Animal struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
}
```
>Veremos los casos prácticos de los tags mas adelante.









































