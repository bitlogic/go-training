# Slices

## ¿Qué es un slice?

La solución a la inflexibilidad de los arrays debido a su tamaño fijo.

Entonces, un slice es parte de un array. Un array que tiene "propiedades" específicas.

![slice](/go-training-beginner/modulo-2/5-slices/img/slice.png)

Una slice se forma de tres partes:

- Una puntero a una array.
- El tamaño (len) de ese array.
- La capacidad, que es el tamaño máximo hasta el que puede crecer el segmento.

Al igual que la función `len()`, podemos determinar la capacidad de un segmento utilizando la función  `cap` integrada. Ejemplo,

```go
package main

import "fmt"

func main() {
	array := [5]int{20, 15, 5, 30, 25}

	slice := array[1:4]

	fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", array, len(array), cap(array))

	fmt.Printf("Slice: %v, Length: %d, Capacity: %d", slice, len(slice), cap(slice))
}
```
Veamos que esta pasando en ese código

## Declaración

Para declarar un slice lo hacemos de la siguiente manera

```go
var s []T
```

Como podemos ver, no necesitamos especificar ninguna longitud.

Declaremos un `slice` de enteros.

```go
func main() {
	var s []string

	fmt.Println(s)
	fmt.Println(s == nil)
}
```
```cmd
$ go run slices.go

[]
true
```

>Entonces, a diferencia de los arreglos, el valor cero de un slice es `nil`

## Inicialización

Hay varias formas de inicializar nuestros slice. Una forma es usar la función `make.

```go
make([]T, len, cap) []T
```
```go
func main() {
var s = make([]string, 0, 0)

	fmt.Println(s)
}
```
```cmd
$ go run slices.go

[]
```

Similar a los array, podemos usar la declaración literal inicializar nuestro segmento

```go
func main() {
	var s = []string{"Go", "TypeScript"}

	fmt.Println(s)
}
```
```cmd
$ go run slices.go

[Go TypeScript]
```

Otra forma es crear un slice a partir de una array. Dado que un slice es parte de un array, podemos crear un segmento de índice `low` a `high` de la siguiente manera

```go
a[low:high]
```
```go
func main() {
	var a = [4]string{
		"C++",
		"Go",
		"Java",
		"TypeScript",
	}

	s1 := a[0:2] // Select from 0 to 2
	s2 := a[:3]  // Select first 3
	s3 := a[2:]  // Select last 2

	fmt.Println("Array:", a)
	fmt.Println("Slice 1:", s1)
	fmt.Println("Slice 2:", s2)
	fmt.Println("Slice 3:", s3)
}
```
```cmd
$ go run slices.go

Array: [C++ Go Java TypeScript]
Slice 1: [C++ Go]
Slice 2: [C++ Go Java]
Slice 3: [Java TypeScript]
```

>Cuando evitamos el uso del índice `low` implica que usamos 0 y si evitamos el uso del índice `high` implica que usamos `len(a)`

## Iteración

Podemos iterar sobre un segmento de la misma manera que iteramos un array, usando el bucle `for` con la función `len()` o usando `range`.

## Funciones

- copy
    
    La función copy() copia elementos de un slice a otro. Se necesitan 2 slices, un destino y una origen.Devuelve el número de elementos copiados.

    ```go
    func copy(dst, src []T) int
    ```
    Veamos cómo usarla.

    ```go
    func main() {
        s4 := []string{"a", "b", "c", "d"}
        s5 := make([]string, 4)

        e := copy(s5, s4)

        fmt.Println("Src:", s4)
        fmt.Println("Dst:", s5)
        fmt.Println("Elements:", e)
    }
    ```
    ```cmd
    $ go run slices.go

    Src: [a b c d]
    Dst: [a b c d]
    Elements: 4
    ```

- append

    Para agregar elementos a nuestro slice usamos la función `append()` que agrega nuevos elementos al final de un slice determinado.

    Toma un slice y un número variable de argumentos. Luego devuelve un nuevo slice que contiene todos los elementos.

    ```go
    append(slice []T, elems ...T) []T
    ```

    Veamos un ejemplo

    ```go
    func main() {
        s6 := []string{"a", "b", "c", "d"}

        s7 := append(s6, "e", "f")

        fmt.Println("s6:", s6)
        fmt.Println("s7:", s7)
    }
    ```
    ```cmd
    $ go run slices.go

    a1: [a b c d]
    a2: [a b c d e f]
    ```


>Si el slice dado no tiene capacidad suficiente para los nuevos elementos, entonces se asigna automáticamente un nuevo slice con una capacidad mayor.

## Propiedades
Ahora analicemos algunas propiedades de los slice.

Los slices son tipos de referencia, a diferencia de los array que son tipo de valor.

Esto significa que modificar los elementos de un slice modificará los elementos correspondientes en el array a la que se hace referencia.

```go
package main

import "fmt"

func main() {
	array2 := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	slice2 := array2[0:2]

	fmt.Println(array2)
	fmt.Println(slice2)

	slice2[0] = "Sun"

	fmt.Println(slice2)
}
```
```cmd
[Mon Tue Wed Thu Fri Sat Sun]
[Mon Tue]
[Sun Tue]
```
Los slices también se pueden usar con funciones variádicas.

```go
package main

import "fmt"

func main() {
	values := []int{1, 2, 3}
	sum := add(values...)
	fmt.Println(sum)
}

func add(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}

	return sum
}
```
