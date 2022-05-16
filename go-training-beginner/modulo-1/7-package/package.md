# Package

![package](/go-training-beginner/modulo-1/7-package/img/package.png)

Un paquete es un un directorio que contiene uno o mas archivos fuentes de Go u otros paquetes de Go

Esto quiere decir que cada  archivo de código de Go, ejemplo `name.go`, debe pertenecer a un paquete, y la declaración de ese paquete se realiza al comienzo de ese archivo.

```go
package <package_name>
```
En Go todo arranca en el paquete `main` que debe contener una función llamada `main()` que funciona como punto de entrada para nuestro ejecutable. 

Por ahora hemos visto que trabajamos solo sobre el paquete `main` ahora vamos a crear un paquete `custom` donde agregaremos algunos archivos como `code.go`

### Importar y exportar

Go tiene una forma muy simple de indicar que, cualquier variable o función, este disponible o sea visible desde cualquier otro paquete. Y es haciendo que la mayuscula la primer letra del nombre de la variable o función.


```go
package custom

var value int = 10 // Will not be exported, private variable
var Value int = 20 // Will be exported
```

Ahora veamos como importamos la variable `Value` desde `main.go`

```go
package main

import (
	"fmt"

	"github.com/username/example/custom" 
    //"github.com/username/example" is the module path
)

func main() {
	fmt.Println(custom.Value)
}
```
También podemos usar un alias

```go
package main

import (
	"fmt"

	code "github.com/username/example/custom" 
    //"github.com/username/example" is the module path
)

func main() {
	fmt.Println(code.Value)
}
```
La estructura final de nuestro código seria algo como...

```
├── example
│   ├── custom
│   │   ├── code.go
│   ├── main.go
│   ├── go.mod
│   ├── go.sum

```
Por último, aclarar que Go no tiene una convención para estructurar nuestro proyecto, pero siempre se recomienda estructurar de la forma mas simple e intuitiva posible. **Es fácil decirlo pero no hacerlo**






