# Arrays

## ¿Qué es un arreglo (array)?

Es una colección de tamaño fijo de elementos del mismo tipo. Estos elementos se almacenan secuencialmente y se puede acceder a ellos usando su `index`.

![array](/go-training-beginner/modulo-2/4-arrays/img/array.png)

## Declaración

Podemos declarar un array de la siguiente manera

```go
var a [n]T
```

Vemos que `n` es la longitud y `T` puede ser de cualquier tipo, como un `int`, un `string`,una `struct`, etc.

Ahora, declaremos un array de enteros (int) con longitud 4.

```go
func main() {
	var arr1 [4]int

	fmt.Println("Array 1: ", arr1)
}
```
```cmd
$ go run arrays.go

Array 1:  [0 0 0 0]
```

Al igual que antes, de forma predeterminada, todos los elementos del array se inicializan con el valor cero (zero type) del tipo de array correspondiente.

## Inicialización

También podemos inicializar un array de manera literal.

```go
var a [n]T = [n]T{V1, V2, ... Vn}

func main() {
	var arr2 = [4]int{1, 2, 3, 4}

	fmt.Println("Array 1: ", arr2)
}
```
```cmd
$ go run arrays.go

Array 2:  [1 2 3 4]
```

También tenemos los shortcuts para los arrays.

```go
...

arr := [4]int{1, 2, 3, 4}
```

## Index

Y al igual que en otros lenguajes, podemos acceder a los elementos usando el `index` que se la asigna a cada elemento cuando se almacenan secuencialmente.

```go
func main() {
	arr := [4]int{1, 2, 3, 4}

	fmt.Println("first element:", arr2[0])
}
```
```cmd
$ go run arrays.go

first element: 1
```

## Iteración

Hay un par de maneras de iterar un array.

El primero es usando el bucle `for` con la función `len` que nos da la longitud del array.

```go
func main() {
	arr := [4]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index: %d, Element: %d\n", i, arr[i])
	}
}
```
```cmd
$ go run arrays.go

Index: 0, Element: 1
Index: 1, Element: 2
Index: 2, Element: 3
Index: 3, Element: 4
```
>Créanme que no van a ver esta forma de iterar un array muy seguido.

Otra forma es usar la palabra reservada `range` con el bucle `for`.

```go
func main() {
	arr := [4]int{1, 2, 3, 4}

	for i, e := range arr {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}
}
```
```cmd
$ go run arrays.go

Index: 0, Element: 1
Index: 1, Element: 2
Index: 2, Element: 3
Index: 3, Element: 4
```

Pero la `range` es bastante versátil y se puede usar de múltiples maneras.

```go
for i, e := range arr {} // Normal usage of range
```
```go
for _, e := range arr {} // Omit index with _ and use element
```
```go
for i := range arr {} // Use index only
```
```go
for range arr {} // Simply loop over the array
```
### Matrices

Una matriz es un array multidimencional. Y en Go se definen de la siguiente manera.

```go
func main() {
	arr := [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for i, e := range arr {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}
}
```
```cmd
$ go run arrays.go

Index: 0, Element: [1 2 3 4]
Index: 1, Element: [5 6 7 8]
```

También podemos dejar que el compilador deduzca la longitud de la matriz usando `...` puntos suspensivos en lugar de la longitud

```go
func main() {
	arr := [...][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for i, e := range arr {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}
}
```
```cmd
$ go run arrays.go

Index: 0, Element: [1 2 3 4]
Index: 1, Element: [5 6 7 8]
```

## Propiedades

Ahora hablemos de algunas propiedades de los arrays.

La longitud de la matriz es parte de su tipo. 

```go
package main

func main() {
	var a = [4]int{1, 2, 3, 4}
	var b [2]int = a
}
```
```cmd
$ go run arrays_properties.go

arrays.go:5:17: cannot use a (variable of type [4]int) as type [2]int in variable declaration

```
Entonces, el array a y b son tipos completamente distintos, y no podemos asignar uno al otro.

>Esto también significa que no podemos cambiar el tamaño de un array, porque cambiar el tamaño de un array significaría cambiar su tipo.

Los arrays en Go son tipos de valor a diferencia de otros lenguajes como C, C++ y Java, donde los arrays son tipos de referencia.

>Esto significa que cuando asignamos un array a una nueva variable o pasamos un array a una función, se copia todo el array.

```go
package main

import "fmt"

func main() {
	var a = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	var b = a // Copy of a is assigned to b

	b[0] = "Monday"

	fmt.Println(a) 
	fmt.Println(b)
}
```
```cmd 
go run arrays_properties.go

[Mon Tue Wed Thu Fri Sat Sun]
[Monday Tue Wed Thu Fri Sat Sun]
```

Entonces, si hacemos algún cambio en esta matriz copiada, la matriz original no se verá afectada y permanecerá sin cambios.
