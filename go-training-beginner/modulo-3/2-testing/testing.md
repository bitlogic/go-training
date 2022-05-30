# Testing

![testing](/go-training-beginner/modulo-3/2-testing/img/testing.png)

En este tema vamos a crear un pequeño proyecto para facilitar los ejemplos

Creado un paquete `math` que contiene una función `Add` que, como sugiere el nombre, suma dos números enteros.

```go
package math

func Add(a, b int) int {
	return a + b
}
```

En nuestro paquete `main` la usamos como sigue...

```go
package main

import (
	"example/math"
	"fmt"
)

func main() {
	result := math.Add(2, 2)
	fmt.Println(result)
}
```

Y, si ejecutamos esto, deberíamos ver el resultado.

```go
$ go run main.go

4
```

Ahora, queremos probar nuestra función `Add`. 

En Go, declaramos archivos de test con el sufijo `_test` en el nombre del archivo. Entonces, para nuestro **add.go**, crearemos una archivo de test como **add_test.go**

La estructura de nuestro proyecto debería verse así.
```
.
├── go.mod
├── main.go
└── math
    ├── add.go
    └── add_test.go
```

Comenzaremos usando un paquete `math_test` e importando el paquete `testing` de la biblioteca estándar. 

>¡Así es! Los test están integradas en Go, a diferencia de muchos otros idiomas :sunglasses: .

Pero espere... ¿por qué necesitamos usar math_test como nuestro paquete? ¿No podemos simplemente usar el mismo paquete math?

En realidad, sí, podemos escribir nuestros test en el mismo paquete si queremos, pero personalmente opino que hacer esto en un paquete separado nos ayuda a escribir los test de una manera más desacoplada.

Ahora, podemos crear nuestra función `TestAdd`. Tomará un argumento de tipo testing.T que nos proporcionará métodos útiles.

```
package math_test

import "testing"

func TestAdd(t *testing.T) {}
```

Antes de agregar cualquier lógica de test, intentemos ejecutarla. Pero esta vez, no podemos usar  el comando `go run`, en su lugar, usaremos el comando `go test`.

```cmd
$ go test ./math

ok      example/math 0.429s
```

Aquí, tendremos el nombre de nuestro paquete, que es `math`, pero también podemos usar la ruta relativa `./...` para probar todos los paquetes.

Y si Go no encuentra ningún test en un paquete, nos lo hará saber.

```cmd 
$ go test ./...

?       example [no test files]
ok      example/math 0.348s
```

Perfecto, ahora escribamos algo de código en el test. Para hacer esto, verificaremos nuestro resultado con un valor esperado y, si no coinciden, podemos usar el método `t.Fail` para fallar hacer fallar el test.

```go
package math_test

import "testing"

func TestAdd(t *testing.T) {
	got := math.Add(1, 1)
	expected := 2

	if got != expected {
		t.Fail()
	}
}
```
```cmd
$ go test math

ok      example/math    0.412s
```
¡Excelente! Nuestro test pasó.

También veamos qué sucede si fallamos el test, para eso, podemos cambiar nuestro resultado esperado.

```go
package math_test

import "testing"

func TestAdd(t *testing.T) {
	got := math.Add(1, 1)
	expected := 3

	if got != expected {
		t.Fail()
	}
}
```
```cmd
$ go test ./math

ok      example/math    (cached)
```

Si ves esto, tranquilo. Para la optimización, nuestros test se almacenan en caché. Podemos usar el comando `go clean` para borrar el caché y luego volver a ejecutar el test.

```cmd
$ go clean -testcache
$ go test ./math

--- FAIL: TestAdd (0.00s)
FAIL
FAIL    example/math    0.354s
FAIL
```

Entonces, así es como se verá una falla en el test.

## Test basados en tablas (Table driven tests)

Esto nos lleva a los test basados en tablas. ¿Pero qué son exactamente?

Antes, teníamos argumentos de funciones y variables esperadas que comparamos para determinar si nuestros test pasaban o fallaban. Pero, ¿qué pasa si definimos todo eso en una slice e iteramos sobre eso? Esto hará que nuestros test sean un poco más flexibles y nos ayudará a ejecutar múltiples casos fácilmente.

No te preocupes, veamos un ejemplo para entenderlo. Entonces, comenzaremos definiendo nuestra estructura `addTestCase`.

```go
package math_test

import (
	"example/math"
	"testing"
)

type addTestCase struct {
	a, b, expected int
}

var testCases = []addTestCase{
	{1, 1, 3},
	{25, 25, 50},
	{2, 1, 3},
	{1, 10, 11},
}

func TestAddTDT(t *testing.T) {

	for _, tc := range testCases {
		got := math.Add(tc.a, tc.b)

		if got != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, got)
		}
	}
}
```

>Notemos cómo declaramos `addTestCase`con minúsculas porque no queremos exportarlo ya que no es útil fuera de nuestra lógica de tests.

Ejecutemos el test...

```
$ go run main.go

--- FAIL: TestAdd (0.00s)
    add_test.go:25: Expected 3 but got 2
FAIL
FAIL    example/math    0.334s
FAIL
```

¡Fallaron!, actualicemos nuestros `test cases`.

```go
var testCases = []addTestCase{
	{1, 1, 2},
	{25, 25, 50},
	{2, 1, 3},
	{1, 10, 11},
}
```

```cmd
$ go run main.go

ok      example/math    0.589s
```

¡Está funcionando!

## Cobertura de código (coverage)

Finalmente, hablemos de la cobertura del código. Al escribir test, a menudo es importante saber qué parte de tu código cubren los tests. Esto generalmente se conoce como `coverage`.

Para calcular y exportar `coverage` de nuestros test, simplemente podemos usar el argumento `-coverprofile` con el comando `go test`.

```
$ go test ./math -coverprofile=coverage.out

ok      example/math    0.385s  coverage: 100.0% of statements
```

También podemos verificar el informe que nos generó el comando anterior (coverage.out) usando el comando `go tool cover`, esto nos mostrará en detalle el coverage en un html.

```cmd
$ go tool cover -html=coverage.out
```
![cover](/go-training-beginner/modulo-3/2-testing/img/cover.png)

Como podemos ver, este es un formato mucho más legible. Y lo mejor de todo, está integrado directamente en las herramientas estándar.

## Test de fuzz (Fuzz testing)

Por último, echemos veamos los test de fuzz que se introdujeron en la versión 1.18 de Go.

`Fuzzing` es un tipo de test automatizado que manipula continuamente las entradas de un programa para encontrar errores.

Go `fuzzing` utiliza la guía de cobertura para recorrer de forma inteligente el código que se está fuzzing para encontrar e informar fallas al usuario.

Dado que puede llegar a casos extremos que los programadores a menudo pasan por alto, los test de fuzz pueden ser particularmente valiosos para encontrar errores y vulnerabilidades de seguridad.

Probemos un ejemplo,

```go
func FuzzTestAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		math.Add(a , b)
	})
}
```

Si ejecutamos esto, veremos que automáticamente creará casos de test. Debido a que nuestra función `Add` es bastante simple, los test pasarán.

```cmd
$ go test -fuzz FuzzTestAdd example/math

fuzz: elapsed: 0s, gathering baseline coverage: 0/192 completed
fuzz: elapsed: 0s, gathering baseline coverage: 192/192 completed, now fuzzing with 8 workers
fuzz: elapsed: 3s, execs: 325017 (108336/sec), new interesting: 11 (total: 202)
fuzz: elapsed: 6s, execs: 680218 (118402/sec), new interesting: 12 (total: 203)
fuzz: elapsed: 9s, execs: 1039901 (119895/sec), new interesting: 19 (total: 210)
fuzz: elapsed: 12s, execs: 1386684 (115594/sec), new interesting: 21 (total: 212)
PASS
ok      foo 12.692s
```

Pero si actualizamos nuestra función `Add` con un caso de borde aleatorio en el que el programa entrará en pánico si (b + 10) es mayor que a.

```go
func AddMore(a, b int) int {
	if a > b + 10 {
		panic("B must be greater than A")
	}

	return a + b
}
```

Y si volvemos a ejecutar el test, este caso límite será detectado por el test de fuzz.

```go
$ go test -fuzz FuzzTestAddMore example/math

warning: starting with empty corpus
fuzz: elapsed: 0s, execs: 0 (0/sec), new interesting: 0 (total: 0)
fuzz: elapsed: 0s, execs: 1 (25/sec), new interesting: 0 (total: 0)
--- FAIL: FuzzTestAdd (0.04s)
    --- FAIL: FuzzTestAdd (0.00s)
        testing.go:1349: panic: B is greater than A
```

Creo que esta es una característica realmente muy buena de Go 1.18.  Se puede obtener más información sobre los test de fuzz en la [doc oficial de Go](https://go.dev/doc/fuzz/).