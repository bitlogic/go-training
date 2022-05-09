# Variables y data types

Vamos a aprender muy rápidamente sobre variables y los diferentes data types que nos provee Go

## Variables

Declaración de variables, hay muchas formas de hacerlo

- Declaración sin inicializar
```go 
var foo string
```

- Declaración inicializando
```go
var foo string = "Gologic"
```

- Declaración multiple
```go
var foo, bar string = "Go", "Logic"
// OR
var (
	foo string = "Go"
	bar string  = "Logic"
)
```

- Declaración sin tipo

Podemos omitir el tipo de variable. En este caso Go va a inferir o suponer el tipo de variable
```go
var foo = "I am a string, I guess..."
```

- Declaración abreviada 

Esta Declaración solo es valida dentro de una función

```go
foo := 100
```

> Debemos tener cuidado con las ultimas dos declaraciones ya que podemos cometer el error de asumir que Go eligió correctamente el tipo.
```go
foo := 10 //variable tipo int

foo := 10.0 //variable tipo float
```

## Constantes

Podemos declarar constantes con la palabra reservada `const`. Estos valores al ser constantes no pueden reasignarse.

```go
const constant = "This is a constant"
```

## Data Types

Los tipos de datos básicos en Go son:

- `string`

	En Go un string es una secuencia de bytes
Se pueden declarar con comillas dobles (" ") o backticks (``) cuando queremos multiples lineas

```go
var name string = "This is BitLogic"

var bio string = `I am boring.
								because this topic is boring.`
```

- `bool`

	Se usa para guardar valores booleanos, como en cualquier otro lenguaje. Sus posibles valores son *true* or *false*

	```go
	var value bool = false
	var isItTrue bool = true
	```

 	*Operadores para operar con tipos booleanos*

	|           |     |      |
	| --------- | --- | ---- |
	| logical   | &&  | \|\| |
	| equality  | ==  | !=   |


- `Numeric types`
	- Signed and Unsigned integers
	
		El tamaño de un `int` o un `uint` depende de la plataforma. Es decir que si estamos en un sistema de 32 bit o 64 bit

		```go
		var i int = 404                     // Platform dependent
		var i8 int8 = 127                   // -128 to 127
		var i16 int16 = 32767               // -2^15 to 2^15 - 1
		var i32 int32 = -2147483647         // -2^31 to 2^31 - 1
		var i64 int64 = 9223372036854775807 // -2^63 to 2^63 - 1
		```

		```go
		var ui uint = 404                     // Platform dependent
		var ui8 uint8 = 255                   // 0 to 255
		var ui16 uint16 = 65535               // 0 to 2^16
		var ui32 uint32 = 2147483647          // 0 to 2^32
		var ui64 uint64 = 9223372036854775807 // 0 to 2^64
		```
		>Cual debo usar? :thinking:	En la mayoría de los casos podemos usar `int` a menos que tengamos una razón especifica para elegir el tamaño o el signo


	- float o punto flotante

		Tenemos dos tipos de punto flotante `float32` y `float64`
	
		>Por default se usa `float64`

		```go
		var f32 float32 = 1.7812 // IEEE-754 32-bit
		var f64 float64 = 3.1415 // IEEE-754 64-bit
		```
	*Algunos operadores para operar con tipos numéricos*

	|           |     |      |           |     |      |
	| --------- | --- | ---- | --------- | --- | ---- |
	| Aritmeticos   | + | - | * | / | %| 
	| Comparación   | ==  | !=  | > |< |<= |>=
	| Incremento/decremento  | ++  | -- |
	| Asignación    | +=  | -=  | *= | /= | %= |

>Existen otros tipo de datos como `complex128`, `complex64`, `rune`, `byte` entre otros que los iremos viendo mas adelante. 

## Zero values

Algo que diferencia a Go de otros lenguajes son los valores por defecto o zero values que Go le asigna a las variables cuando no se las inicializa. En estos caso NO existe el `null`.
Usemos un pedacito de código para ver los zero values o valores por defectos de los tipos de datos básicos.

```go
var i int       // 0
var f float64   // 0
var b bool      // false
var s string    // ""

fmt.Printf("%v %v %v %q\n", i, f, b, s)
```

```cmd
 $ go run main.go

 0 0 false ""
```
> No te preocupes por `Printf` y esos símbolos de porcentajes, lo vamos a ver mas adelante

### Types

Otra belleza de Go es que podemos definir nuestros propios tipos de datos simplemente con la palabra reservada `type`

```go
package main

import "fmt"

type MycustomType string

func main() {
	var myData MycustomType = "I am a custom typedata"

	fmt.Printf("%T - %s", str, str) 
}
```

```cmd
 $ go run main.go

 main.MycustomType - I am a custom typedata
```

>Te prometo que esto cobra mas sentido cuando veamos estructuras en el siguiente modulo










