# Modules

![mod](/go-training-beginner/modulo-1/6-modules/img/mod.png)
### Module, que son?

Una definición simple es que un modulo (module) es una colección de paquetes almacenados en un archivo `go.mod` en la raíz del proyecto, siempre que este este afuera de `$GOPATH/src`

A partir de la versión 1.13 de Go, go.mod o el trabajo con módulos es la forma predeterminada de trabajo. Antes debíamos activarlo mediante la variable `GO111MODULE=on`.

### Y el GOPATH?

`GOPATH` es una variable que define la raíz the tu workspace y contiene las siguientes carpetas

- **src**: Contiene el código fuente de Go.
- **pkg**: Contiene los paquetes compilados.
- **bin**: Contiene los binarios y los ejecutables

![gopath](/go-training-beginner/modulo-1/6-modules/img/gopath.png)

### Nuevo proyecto, nuevo go.mod

La forma de empezar un nuevo proyecto es creando un `go.mod` con el comando `go mod init`

```cmd
$ go mod init example
```
>Es importante saber que un Go module puede corresponder a un repositorio de Github. Quiere decir que seguramente vas a subir tu repo y la convención seria iniciarlo de la siguiente manera.

```cmd
$ go mod init github.com/username/example
```
El comando anterior no va a generar un archivo go.mod. Este va a definir la ruta del modulo y las rutas de importación que usa GOPATH y las dependencias del proyecto. Va a tener la siguiente forma

```go
module <name> //github.com/username/example

go <version> // 1.18

require (
	... //dependencies
)
```

Por ejemplo, si quisiéramos agregar un dependencia, podemos usar el comando `go install`.

```cmd
$ go install github.com/rs/zerolog
```

entonces nuestro `go.mod` se vera así.

```go
module github.com/username/example

go 1.18

require (
	"github.com/rs/zerolog"
)
```

Luego de ese comando veremos que se creo un archivo go.sum. Este archivo contiene todos los *hashes* necesarios de las dependencias requeridas en `go.mod`

Por último el comando mas *épico* que vamos a usar es `go mod tidy`. 

```cmd
$ go mod tidy
```

Este nos va a "acomodar" nuestro `go.mod` y `go.sum` dependiendo si estamos o no usando las dependencias a lo largo del proyecto.

>En este tema hay mucho para hablar, como los comandos restantes de go, vendoring, mas sobre gopath, workspace, etc. Creo que no es el scope de este curso, y seguro lo vamos a ir aprendiendo a medida que trabajemos con el lenguaje.



