# Hello world 

![setup](/go-training-beginner/modulo-1/1-hello_world/img/init.webp)

Empecemos escribiendo nuestro "hello world" 馃ケ. 

Primero vamos a inicializar un modulo ( *module* ). Para eso podemos usar el comando `go mod`


```
$ go mod init hello-world
```

>Pero que es *module*? 馃 No nos preocupemos por eso ahora,  lo vamos a tratar mas adelante. Por ahora lo pensemos como una colecci贸n de Go packages, como un gestor de dependencias.

Vamos a crear un archivo **main.go** y escribir un c贸digo que simplemente imprima un "hello world"

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

Analicemos r谩pidamente la estructura de este programa

Primero definimos un paquete como `main`
```go
package main
```

Despu茅s tenemos un `import`.
>"fmt" is una parte de la librer铆a est谩ndar de Go. Que es un conjunto de paquetes proporcionados por el lenguaje

```go
import "fmt"
```

Por ultimo tenemos nuestra funci贸n principal que act煤a como punto de entrada para nuestra aplicaci贸n, igual que otros lenguajes como C, C# o Java.

```go
func main() {
    ...
}
```

>De nuevo, el objetivo de este "hello world" es crear una introducci贸n al curso, durante el curso vamos a aprender sobre funciones, imports, paquetes y otras cosas!

Por 煤ltimo ejecutamos el c贸digo con el comando `go run`

```
$ go run main.go

Hello world!
```

Eso es todo, sigamos con el siguiente tema **insertar link**
