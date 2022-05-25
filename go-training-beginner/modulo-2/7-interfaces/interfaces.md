# Interfaces 

![interface](/go-training-beginner//modulo-2/7-interfaces/img/interfaces.jpeg)

## ¿Qué es una *interface*?

Una interfaz es un tipo abstracto que se define usando un conjunto de firmas de métodos. La interface define el **comportamiento**  de "objetos" similares.

El termino **comportamiento** es clave para entender interfaces.

Entendamos esto con un ejemplo de la vida cotidiana.

En el día a día necesitamos conectar diferentes dispositivos al enchufe o toma corriente.

![no-interface](/go-training-beginner/modulo-2/7-interfaces/img/no-interface.png)

En código podemos representar la imagen anterior de la siguiente manera.

```go
type mobile struct {
	brand string
}

type laptop struct {
	cpu string
}

type toaster struct {
	amount int
}

type kettle struct {
	quantity string
}

type socket struct{}
```

Ahora definamos un método (`Draw`) que se va a encargar de crear un enchufe para nuestro dispositivo `mobile`.

```go
func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d", m, m.brand, power)
}
```
Ahora definamos un método (`Plug`) para nuestro enchufe (`socket`), para decirle que acepte nuestro dispositivos `mobile`

```go
func (socket) Plug(device mobile, power int) {
	device.Draw(power)
}
```
Intentemos "conectar" o "enchufar" el `mobile` a nuestro `socket` en la función `main`

```go
package main

import "fmt"

func main() {
	m := mobile{"Apple"}

	s := socket{}
	s.Plug(m, 10)
}
```

Y si ejecutamos el lo anterior

```cmd
$ go run no_interfaces.go

main.mobile -> brand: Apple, power: 10
```

Excelente, pero ahora queremos conectar nuestra `laptop`.

```go
package main

import "fmt"

func main() {
	m := mobile{"Apple"}
	l := laptop{"Intel i9"}

	s := socket{}

	s.Plug(m, 10)
	s.Plug(l, 50) 
}
```
```cmd
go run no_interfaces.go

./interfaces.go:38:9: cannot use l (variable of type laptop) as type mobile in argument to s.Plug
```
Como podemos ver, esto arrojará un error.

Para resolver debemos crear el método `Draw` que se encarga de crear el enchufe necesario para nuestra `laptop`

```go
func (l laptop) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", l, l.cpu, power)
}
```

y también deberíamos agregar el método a nuestro `socket` para decirle que acepte a nuestra `laptop`.

```go
func (socket) PlugLaptop(device laptop, power int) {
	device.Draw(power)
}
```
Con esos dos nuevos métodos resolvemos el problema y ya podemos conectar un `mobile` y una `laptop`.

```cmd
$ go run no_interfaces.go

main.mobile -> brand: Apple, power: 10
main.laptop -> brand: Intel i7, power: 50
```

Ahora sigamos con la tostadora `toaster`. Que ya nos estamos imaginando que necesitamos si o si un método que cree nuestro enchufe `Draw` y otro `PlugToaster` para decirle al enchufe `socket` que acepte nuestra tostadora.

No parece muy difícil, pero... cada vez que agregamos un nuevo tipo de dispositivo, también necesitaremos agregar un nuevo método al tipo de socket y eso no es lo ideal.

Aquí es donde entran las interfaces. Esencialmente, queremos definir un **contrato** que, en el futuro, deba implementarse.

Podemos definir una interfaz como `PowerDrawery` usarla en nuestra función `Plug` para permitir que cualquier dispositivo que satisfaga los criterios, que es que el tipo debe tener un método `Draw` que coincida con la firma que requiere la interfaz.

Y de esta manera, el socket no necesitará saber nada sobre nuestro dispositivo y simplemente puede llamar al método `Draw` de cada tipo.

![interface](/go-training-beginner/modulo-2/7-interfaces/img/interface.png)


Entonces implementemos nuestra interfaz `PowerDrawer` que se verá algo así.

```go
type PowerDrawer interface {
	Draw(power int)
}
```

![implementation](/go-training-beginner/modulo-2/7-interfaces/img/interface-implementation.png)

>La convención es usar "-er" como sufijo en el nombre. Y como discutimos anteriormente, una interfaz solo debe describir el comportamiento esperado . Que en nuestro caso es el método `Draw`.

Ahora, necesitamos actualizar nuestro método`Plug` para aceptar un dispositivo que implemente la interface `PowerDrawer` como argumento.

```go
func (socket) Plug(device PowerDrawer, power int) {
	device.Draw(power)
}
```

Y para satisfacer la interfaz, simplemente podemos agregar métodos `Draw` a todos los tipos de dispositivos.

```go
type mobile struct {
	brand string
}

func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", m, m.brand, power)
}

type laptop struct {
	cpu string
}

func (l laptop) Draw(power int) {
	fmt.Printf("%T -> cpu: %s, power: %d\n", l, l.cpu, power)
}

type toaster struct {
	amount int
}

func (t toaster) Draw(power int) {
	fmt.Printf("%T -> amount: %d, power: %d\n", t, t.amount, power)
}

type kettle struct {
	quantity string
}

func (k kettle) Draw(power int) {
	fmt.Printf("%T -> quantity: %s, power: %d\n", k, k.quantity, power)
}
```

¡Ahora, podemos conectar todos nuestros dispositivos al enchufe con la ayuda de nuestra interfaz!

```go
func main() {
	m := mobile{"Apple"}
	l := laptop{"Intel i9"}
	t := toaster{4}
	k := kettle{"50%"}

	s := socket{}

	s.Plug(m, 10)
	s.Plug(l, 50)
	s.Plug(t, 30)
	s.Plug(k, 25)
}
```

Y tal como esperábamos, funciona.

```cmd
$ go run interface.go

main.mobile -> brand: Apple, power: 10
main.laptop -> cpu: Intel i9, power: 50
main.toaster -> amount: 4, power: 30
main.kettle -> quantity: 50%, power: 25
```

### Pero, ¿por qué se considera que *interface* es un concepto tan poderoso?

Bueno, una interfaz puede ayudarnos a desacoplar nuestros tipos. Por ejemplo, debido a que tenemos la interfaz, no necesitamos actualizar la implementación de `socket`. Simplemente podemos definir un nuevo tipo de dispositivo con un método `Draw`.

>A diferencia de otros lenguajes, las interfaces de Go se implementan implícitamente , por lo que no necesitamos algo como una palabra reservada como *implements*. Esto significa que un tipo satisface una interfaz automáticamente cuando tiene ***todos los métodos*** de la interfaz.

## Interfaz vacía

Una interfaz vacía puede tomar un valor de cualquier tipo.

Así es como lo declaramos.

```go
var x interface{}
```

Pero, ¿para que sirve?

Las interfaces vacías se pueden usar para manejar valores de tipos desconocidos.

Algunos ejemplos son:

- Lectura de datos heterogéneos de una API
- Variables de tipo desconocido, como en la función fmt.Println()

Para usar un valor de una interface{}, podemos usar aserción de tipo o un cambio de tipo para determinar el tipo del valor.

### Aserción de tipo (Type Assertion)

Una aserción de tipo proporciona acceso al valor concreto subyacente de un valor de interfaz.

Por ejemplo

```go
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
}
```

Esta declaración afirma que el valor de la interfaz contiene un tipo concreto y asigna el valor del tipo subyacente a la variable.

También podemos probar si un valor de interfaz contiene un tipo específico.

Una aserción de tipo puede devolver dos valores:

- El primero es el valor subyacente.
- El segundo es un valor booleano que informa si la afirmación tuvo éxito.

```go
s, ok := i.(string)

fmt.Println(s, ok)
```

Esto puede ayudarnos a probar si un valor de interfaz contiene un tipo específico o no.

En cierto modo, esto es similar a cómo leemos los valores de un mapa. Y si este no es el caso, `ok` será falso y el valor será el valor cero del tipo, y no tendremos un error en tiempo de ejecución (panic error).

```go
f, ok := i.(float64)

fmt.Println(f, ok)
```

Pero si la interfaz no contiene el tipo, la declaración desencadenará un pánico.

```go
f = i.(float64)

fmt.Println(f) // Panic!
```
```
$ go run interface_properties.go

hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64
```

## Type switch

`Switch` se puede usar con una declaración para determinar el tipo de una variable de tipo una interface vacía.

```go
var t interface{}
t = "hello"

switch t := t.(type) {
case string:
	fmt.Printf("string: %s\n", t)
case bool:
	fmt.Printf("boolean: %v\n", t)
case int:
	fmt.Printf("integer: %d\n", t)
default:
	fmt.Printf("unexpected: %T\n", t)
}
```

## Propiedades

Analicemos algunas propiedades de las interfaces.

### zero value

El valor cero de una interfaz es `nil`

```go
package main

import "fmt"

type MyInterface interface {
	Method()
}

func main() {
	var i MyInterface

	fmt.Println(i) // Output: <nil>
}

```
### Embedding

Podemos usar interfaces como estructuras.

Por ejemplo

```go
type interface1 interface {
    Method1()
}

type interface2 interface {
    Method2()
}

type interface3 interface {
    interface1
    interface2
}
```

>Es una característica realmente poderosa, pero recuerda: "Cuanto más grande es la interfaz, más débil es la abstracción" - Rob Pike.





