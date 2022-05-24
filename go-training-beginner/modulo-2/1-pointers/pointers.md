# Pointers (Punteros)

## ¿Qué es un puntero (pointer)?

La definición simple es *un puntero es una variable que es usada para guardar la dirección de memoria de otra variable*.

![pointer](/go-training-beginner/modulo-2/1-pointers/img/pointers.png)

### Declaración de punteros

```go
var p *T
```

donde T puede ser cualquier tipo o data type que vimos anteriormente como `int`, `string`, `float`, etc.

La mejor manera de entender un puntero es con ejemplos.

```go
package main

import "fmt"

func main() {
	var p *int

	fmt.Println(p)
}
```
```cmd
$ go run main.go

nil
```

>¿Qué es `nil`? representa el **zero value** para tipos de datos como punteros, interfaces, channels, maps y slices. Que son tipos de datos que veremos mas adelante.
Para recordar con un ejemplo, el zero value de un `int` es un 0, el de un `bool` false... bueno para un puntero es `nil`.

Ahora asignemos un valor a un puntero.

```go
package main

import "fmt"

func main() {
	a := 10

	var p *int = &a

	fmt.Println("address:", p)
}
```

Usamos `&` para hacer referencia a la dirección de memoria donde esta guardada la variable.

```cmd
$ go run main.go

0xc0000b8000
```

El resultado nos da la dirección de memoria donde esta guardada la variable `a`.

Otra forma de declara e inicializar punteros es 

```go
package main

import "fmt"

func main() {
	p := new(int)
	*p = 100

	fmt.Println("value", *p)
	fmt.Println("address", p)
}
```
```cmd
$ go run main.go

value 100
address 0xc000018030
```

### `*` operador de desreferencia

Ahora nuestra variable puntero `p` tiene guardada la dirección de memoria donde se encuentra la variable `a`.
Si queremos obtener el valor que se encuentra dentro de esa dirección de memoria (es decir el valor de `a`) usamos el operador `*`

```go
package main

import "fmt"

func main() {
	a := 10

	var p *int = &a

	fmt.Println("address:", p)
	fmt.Println("value:", *p)
}
```

```cmd
$ go run main.go

address: 0xc000018030
value: 10
```
 Ademas podemos cambiar el valor de nuestra variable `a` a través de nuestro puntero `p`

```go
package main

import "fmt"

func main() {
	a := 10

	var p *int = &a

	fmt.Println("before", a)
	fmt.Println("address:", p)

	*p = 20
	fmt.Println("after:", a)
    fmt.Println("address:", p)
}
```
```cmd
$ go run main.go

before 10
address: 0xc000192000

after: 20
address: 0xc000192000
```

>Algo importante es notar como se usa `*` tanto para definir el puntero (`var p *int`) como para desreferenciarlos `*p`. Debemos entender esa diferencia.

### Punteros en argumentos de funciones

Recordando las definiciones de pasar valores *por valor* y por *referencia*. Go maneja el mismo concepto y podemos pasar argumentos por referencia como sigue.

```go
myFunction(&a)
...

func myFunction(ptr *int) {}
```

### Nota importante

Go NO soporta operaciones aritméticas como lo hace C o C++.

El siguiente ejemplo No es válido.

```go
	p1 := p * 2 // Compiler Error: invalid operation
```

Pero si podemos comparar punteros con `==`

```go
p := &a
p1 := &a

fmt.Println(p == p1)
```

No compara por tipos o valor, compara su contenido, es decir si apuntan a la misma dirección de memoria.



