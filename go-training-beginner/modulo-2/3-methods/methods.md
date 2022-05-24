# Methods

![methods](/go-training-beginner/modulo-2/3-methods/img/methods.png)

Técnicamente, Go no es un lenguaje de programación orientado a objetos. No tiene clases, objetos ni herencia.

Sin embargo, Go tiene tipos (type) y puede definir métodos para esos tipos.

## ¿Qué es un método?

Un método no es más que una función con un argumento receptor especial. Veamos cómo podemos declararlos.

```go
func (variable T) Name(params) (returnTypes) {}
```
El *argumento del receptor tiene un nombre y un tipo*. Aparece entre la palabra reservada `func` y el nombre del método.

Por ejemplo, definamos una estructura `Car`

```go
type Car struct {
	Name string
	Year int
}
```

Ahora, definamos un método `IsLatest` que nos diga si un automóvil se fabricó en los últimos 5 años.

```go
func (c Car) IsLatest() bool {
	return c.Year >= 2018
}
```

Como se puede ver, podemos acceder a la instancia de `Car` usando la variable receptora `c`. 

>Podríamos ver esto como la palabra `this` en la programación orientada a objetos.

Ahora deberíamos poder llamar a este método después de inicializar nuestra estructura, tal como lo hacemos con las clases en otros idiomas.

```go
func main() {
	c := Car{"Tesla", 2021}

	fmt.Println("Is latest", c.IsLatest())
}
```

## Métodos con punteros como receptores 

El ejemplo que vimos anteriormente tenían un receptor de valor.

Con un receptor de valor, el método opera en una copia del valor que se le pasa. Por lo tanto, cualquier modificación realizada en el receptor dentro de los métodos no es visible para la persona que llama.

Por ejemplo, hagamos otro método llamado `UpdateName` que actualizará el nombre de `Car`

```go
func (c Car) UpdateName(name string) {
	c.Name = name
}
```

Ahora, ejecutemos esto

```go
func main() {
	c := Car{"Tesla", 2021}

	c.UpdateName("Toyota")
	fmt.Println("Car:", c)
}
```
```cmd
$ go run methods.go

Car: {Tesla 2021}
```

Como vemos el nombre no se actualizó. Si ahora cambiemos nuestro receptor al tipo puntero e intentemos nuevamente

```go
func (c *Car) UpdateName(name string) {
	c.Name = name
}
```
```cmd
$ go run methods.go

Car: {Toyota 2021}
```

En conclusión, los métodos con punteros como receptores pueden modificar el valor al que apunta el receptor. Por lo que las modificaciones también son visibles para la persona que llama al método.

## Propiedades

Go es lo suficientemente inteligente como para interpretar nuestra llamada de función correctamente y, por lo tanto, y como vimos en el ejemplo anterior la notación siguiente es valida pero no es necesaria.
```go
(&c).UpdateName(...)
```

También podemos omitir la parte variable del receptor si no lo estamos usando.

```go
func (Car) UpdateName(...) {
    fmt.Println("I'm not used the receptor")
}
```

Los métodos no se limitan solo a estructuras, sino que también se pueden usar con tipos (type) que no son estructuras.

```go
package main

import "fmt"

type MyInt int

func (i MyInt) isGreater(value int) bool {
	return i > MyInt(value)
}

func main() {
	i := MyInt(10)

	fmt.Println(i.isGreater(5))
}
```

### ¿Por qué métodos en lugar de funciones?

Como siempre, no hay una respuesta correcta para esto, y de ninguna manera una es mejor que la otra. 

Por ejemplo, los métodos pueden ayudarnos a evitar conflictos de nombres. Como un método está vinculado a un tipo particular, podemos tener los mismos nombres de método para múltiples receptores. *Discutible...*

O podemos usar métodos para definir el alcance de las funciones. *Discutible...*

También, podría decirse que un métodos es mucho mas fácil de leer que una función. *Discutible...*

>En fin, los métodos deben usarse adecuadamente cuando llegue la situación.

