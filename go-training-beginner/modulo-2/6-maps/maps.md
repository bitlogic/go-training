# Maps

![maps-go](/go-training-beginner/modulo-2/6-maps/img/maps-ghoper.png)

## ¿Qué es un map (mapa)?

Bueno, un mapa es una colección desordenada de pares clave-valor. Las claves son únicas dentro de un mapa mientras que los valores pueden no serlo.

![map](/go-training-beginner/modulo-2/6-maps/img/maps.png)

Se utilizan para búsquedas rápidas, recuperación y eliminación de datos basados ​​en claves. Es una de las estructuras de datos que más se usan.

## Declaración

Un mapa se declara usando la siguiente sintaxis

```go
var m map[K]V
```

Donde `K` está el tipo (type) de la clave y `V` el tipo (type) de valor?

Por ejemplo, así es como podemos declarar un map de claves tipo `string` y valores tipo a `int`.

```go
func main() {
	var m map[string]int

	fmt.Println(m)

	if m == nil {
		fmt.Println("nil is zero value to maps")
	}
}
```
```cmd
$ go run maps.go

map[]
nil is zero value to maps
```

Como podemos ver, el valor cero de un mapa es nil.

Un map `nil` no tiene claves. Además, cualquier intento de agregar claves a un map `nil` resultará en un error de tiempo de ejecución.

## Inicialización

Hay varias formas de inicializar un mapa.

### función `make`

Podemos usar la función make, que asigna memoria para los tipos de datos a los que se hace referencia e inicializa sus estructuras de datos subyacentes.

```go
func main() {
	var m = make(map[string]int)

	fmt.Println(m)
}
```
```cmd
$ go run maps.go

map[]
```
### literal map

Otra forma es usar el mapa literal.

```go
func main() {
	var m = map[string]int{
		"a": 0,
         "b": 1,
	}

	fmt.Println(m)
}
```
>La última coma final es necesaria

```cmd
$ go run maps.go

map[a:0 b:1]
```
También podemos usar nuestros tipos personalizados o estructuras.

```go
type Pet struct {
	Name string
}

func main() {
	var m = map[string]Pet{
		"a": Pet{"Gopher"},
		"b": Pet{"Hedwing"},
	}

	fmt.Println(m)
}
```

¡Incluso podemos eliminar el tipo de valor y Go lo resolverá!

```go
var m = map[string]Pet{
	"a": {"Gopher"},
	"b": {"Hedwing"},
}
```
```
$ go run maps.go

map[a:{Gopher} b:{Hedwing}]
```

## Agregar valores

Veamos cómo podemos agregar un valor a nuestro mapa.

```go
func main() {
	var m = map[string]Pet{
		"a": {"Gopher"},
		"b": {"Hedwing"},
	}

	m["c"] = Pet{"Bito"}

	fmt.Println(m)
}
```
```cmd
$ go run iteration.go

map[a:{Gopher} b:{Hedwing} c:{Bito}]
```

## Recuperar valores

También podemos recuperar nuestros valores del mapa usando la clave

```go
...
c := m["c"]
fmt.Println("Key c:", c)
```
```cmd
$ go run iteration.go

key c: {Bito}
```

¿Qué pasa si usamos una clave que no está presente en el mapa?

```go
...
d := m["d"]
fmt.Println("Key d:", d)
```
Obtendremos el valor cero (zero value) del tipo de valor del mapa.

```cmd
$ go run iteration.go

Key c: {Bito}
Key d: {}
```

### Existencia del par clave-valor

Cuando se recupera el valor asignado a una clave determinada, también devuelve un valor booleano adicional. La variable booleana será `true` si existe la clave, y `false` en caso contrario.

Intentemos esto en un ejemplo.

```go
...
c, ok := m["c"]
fmt.Println("Key c:", c, ok)

d, ok := m["d"]
fmt.Println("Key d:", d, ok)
```
```cmd
$ go run iteration.go

Key c: {Bito} Present: true
Key d: {} Present: false
```

## Actualizar un valor

También podemos actualizar el valor de una clave simplemente reasignando una clave.

```go
...
m["a"] = Pet{"Chita"}
```
>En este caso, en que el valor es una estructura o tipo personalizado, si necesitamos especificar el tipo.

```cmd
$ go run iteration.go

map[a:{Chita} b:{Hedwing} c:{Bito}]
```

## Eliminar una clave

Podemos eliminar la clave usando la función `delete` incorporada.

La sintaxis seria algo así

```go
...
delete(m, "b")
```

El primer argumento es el mapa y el segundo es la clave que queremos eliminar.

La función `delete()` no devuelve ningún valor. Además, no hace nada si la clave no existe en el mapa.

```cmd
$ go run iteration.go

map[a:{Chita} c:{Bito}]
```

## Iteración

De manera similar a los arrays los slice, podemos iterar sobre los mapas con la palabra reservada `range`.

```go
package main

import "fmt"

func main() {
	var m = map[string]Pet{
		"a": {"Gopher"},
		"b": {"Hedwing"},
	}

	m["c"] = Pet{"Bito"}

	for key, value := range m {
		fmt.Println("Key: %s, Value: %v", key, value)
	}
}
```
```cmd
$ go run iteration.go

Key: c, Value: {Bito}
Key: a, Value: {Gopher}
Key: b, Value: {Hedwing}
```

>Tengamos en cuenta que un mapa es una colección desordenada y, por lo tanto, no se garantiza que el orden de iteración de un mapa sea el mismo cada vez que iteramos sobre él.

## Propiedades

Por último, hablemos de las propiedades del mapa.

Los mapas son tipos de referencia, lo que significa que cuando asignamos un mapa a una nueva variable, ambos se refieren a la misma estructura de datos subyacente.

Por lo tanto, los cambios realizados por una variable serán visibles para la otra.

```go
package main

import "fmt"

type Pet struct {
	Name string
}

func main() {
	var m = map[string]Pet{
		"a": {"Gopher"},
		"b": {"Hedwing"},
	}

	m2 := m
	m2["c"] = Pet{"Bito"}

	fmt.Println(m) 
	fmt.Println(m2) 
}
```