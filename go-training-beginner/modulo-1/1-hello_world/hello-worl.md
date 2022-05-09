# Hello world 

Empecemos escribiendo nuestro "hello world" 🥱. 

Primero vamos a inicializar un modulo ( *module* ). Para eso podemos usar el comando `go mod`


```
$ go mod init hello-world
```

>Pero que es *module*? 🤔 No nos preocupemos por eso ahora,  lo vamos a tratar mas adelante. Por ahora lo pensemos como una colección de Go packages, como un gestor de dependencias.

Vamos a crear un archivo **main.go** y escribir un código que simplemente imprima un "hello world"

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

Analicemos rápidamente la estructura de este programa

Primero definimos un paquete como `main`
```go
package main
```

Después tenemos un `import`.
>"fmt" is una parte de la librería estándar de Go. Que es un conjunto de paquetes proporcionados por el lenguaje

```go
import "fmt"
```

Por ultimo tenemos nuestra función principal que actúa como punto de entrada para nuestra aplicación, igual que otros lenguajes como C, C# o Java.

```go
func main() {
    ...
}
```

>De nuevo, el objetivo de este "hello world" es crear una introducción al curso, durante el curso vamos a aprender sobre funciones, imports, paquetes y otras cosas!

Por último ejecutamos el código con el comando `go run`

```
$ go run main.go

Hello world!
```

Eso es todo, sigamos con el siguiente tema **insertar link**
