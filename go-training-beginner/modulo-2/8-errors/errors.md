# Errors

![error](/go-training-beginner/modulo-2/8-errors/img/golang-error.png)


Hablemos de errores y solo de errores. 

Hago esa aclaración para quienes vienen de otros lenguajes en donde en este tema se esperan encontrar con el manejo de excepciones. Pero recordemos que en Go no existen las excepciones. En su lugar podemos usar un tipo de dato incorporado en Go que es `error`

```go
type error interface {
    Error() string
}
```

Regresaremos a esto en breve. Primero, tratemos de entender los conceptos básicos.

Entonces, declaremos una función `Divide` simple que, como su nombre indica, dividirá enteros `a` entre `b`

```go
func Divide(a, b int) int {
	return a/b
}
```

Estupendo. Ahora, queremos devolver un error para evitar la división por cero. Esto nos lleva a la construcción de errores.

## Creación de errores

Hay varias formas de hacer esto, pero veremos las dos más comunes.

### Errors package

La primera es mediante el uso de la función `New()` proporcionada por el paquete `errors`.

```go
package main

import "errors"

func main() {}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a/b, nil
}
```

Notemos cómo devolvemos el error con el resultado. Y si no hay error devolvemos `nil` ya que es el valor cero (zero value) de un error. Porque al fin y al cabo `error` es una interfaz.

### Pero, ¿cómo lo manejamos? 

Para eso, llamemos a la función `Divide` en nuestra función `main`.

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(4, 0)

	if err != nil {
		fmt.Println(err) // Do something with the error
		return
	}

	fmt.Println(result) // Use the result
}

func Divide(a, b int) (int, error) {...}
```

```cmd
$ go run errors.go

cannot divide by zero
```

Simplemente verificamos si el error es `nil` y creamos nuestra lógica en consecuencia. Esto se considera bastante idiomático en Go y veremos que se usa mucho.

### función fmt.Errorf

Otra forma de construir nuestros errores es usando la función `fmt.Errorf()`.

Esta función es similar `fmt.Sprintf()` y nos permite formatear nuestro error. Pero en lugar de devolver una cadena, devuelve un error.

Generalmente se usa para agregar algo de contexto o detalle a nuestros errores.

``` go
...
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}

	return a/b, nil
}
```
```cmd
$ go run errors.go

cannot divide 4 by zero
```
## Sentinel Errores

Otra técnica importante en Go es definir los errores esperados para que puedan verificarse explícitamente en otras partes del código. Estos a veces se denominan errores centinela (sentinel errores).

```go
package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("cannot divide by zero")

func main() {...}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return a/b, nil
}
```

>En Go, es una convención usar el prefijo `Err` en la variable. Por ejemplo, `ErrNotFound`.

Por ejemplo, ahora podemos verificar explícitamente qué error ocurrió usando la función `errors.Is`.

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(4, 0)

	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println(err) // Do something with the error
		default:
			fmt.Println("no idea!")
		}

		return
	}

	fmt.Println(result) // Use the result
}

func Divide(a, b int) (int, error) {...}
```
```cmd 
$ go run sentinel_errors.go

cannot divide by zero
```
## Custom Errors

Esta estrategia cubre la mayoría de los casos de uso del manejo de errores. Pero a veces necesitamos funcionalidades adicionales, como valores dinámicos dentro de nuestros errores.

Anteriormente, vimos que `errors` es solo una interfaz. Básicamente, cualquier cosa puede ser una error siempre que implemente el método `Error()` que devuelve un mensaje de error como una cadena.

Entonces, definamos una estructura `DivisionError` que oficiará como un error personalizado y contendrá un código de error y un mensaje.

```go
package main

import (
	"errors"
	"fmt"
)

type DivisionError struct {
	Code int
	Msg  string
}

func (d DivisionError) Error() string {
	return fmt.Sprintf("code %d: %s", d.Code, d.Msg)
}

func main() {...}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionError{
			Code: 2000,
			Msg:  "cannot divide by zero",
		}
	}

	return a/b, nil
}
```

Usaremos la función `errors.As` en lugar de la función `errors.Is` para convertir el error al tipo correcto.

```go
func main() {
	result, err := Divide(4, 0)

	if err != nil {
		var divErr DivisionError

		switch {
		case errors.As(err, &divErr):
			fmt.Println(divErr) // Do something with the error
		default:
			fmt.Println("no idea!")
		}

		return
	}

	fmt.Println(result) // Use the result
}

func Divide(a, b int) (int, error) {...}
```
```cmd
$ go run custom_error.go

code 2000: cannot divide by zero
```
Pero, ¿cuál es la diferencia entre `errors.Is` y `errors.As`?

La diferencia es que la función  `As()` verifica si el error tiene un tipo específico, a diferencia de `Is()`, que examina si es un objeto de error en particular.

## Type Assertion errors

También podemos usar aserciones de tipo, ***pero no es remondado en Go***

```go
func main() {
	result, err := Divide(4, 0)

	if e, ok := err.(DivisionError); ok {
		fmt.Println(e.Code, e.Msg) // Output: 2000 cannot divide by zero
		return
	}

	fmt.Println(result)
}
```

>En conclusión, vimos que el manejo de errores en Go es bastante diferente en comparación con el `try/catch` de otros lenguajes . Pero es muy poderoso ya que alienta al desarrollador a manejar el error de manera explícita, lo que también mejora la legibilidad.