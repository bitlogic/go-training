# Panic, recover, defer

En la mayoría de los casos usamos los errores (error) proporcionados por Go para manejar las condiciones anormales en nuestros programas. Sin embargo, hal algunas situaciones en las que programa no puede continuar.

En esos casos, podemos usar la la función `panic` proporcionada por Go.

![panic](/go-training-beginner//modulo-3/1-panic_recover_defer/img/panic.jpeg)

## Defer

Empecemos por la palabra reservada `defer` ya que nos servirá mas adelante para entender el resto de los temas.

Defer nos permite posponer la ejecución de una función hasta que regrese la función circundante. Es decir, antes de devolver la llamada se ejecutará `defer`.

```go
func main() {
	defer fmt.Println("I am finished")
	fmt.Println("Doing some work...")
}
```

¿Podemos usar múltiples funciones diferidas? La respuesta es sí y esto nos lleva a lo que se conoce como `defer stack` . Veamos un ejemplo.

```go
func main() {
	defer fmt.Println("I am finished")
	defer fmt.Println("Are you sure?")

	fmt.Println("Doing some work...")
}
```
```cmd
$ go run main.go

Doing some work...
Are you sure?
I am finished
```

Como podemos ver, las sentencias diferidas se apilan y ejecutan en forma de *último en entrar, primero en salir*.

Por lo tanto, `defer` es increíblemente útil y se usa comúnmente para realizar tareas de limpieza o manejo de errores.

## Panic

Ahora si, veamos `panic`.

```go
func panic(interface{})
```

La función `panic` detiene la ejecución normal de la actual goroutine.

Cuando una función llama a `panic`, la ejecución normal de la función se detiene inmediatamente y el control se devuelve a la función que llamó anteriormente. Esto se repite hasta que el programa sale con el mensaje de panic y el seguimiento de la pila (stacktrace).

>Nota: Hablaremos más adelante de las goroutines en el curso, por ahora lo tomemos como una función.

Veamos cómo podemos usar la función `panic`.

```go
package main

func main() {
	WillPanic()
}

func WillPanic() {
	panic("Ahhh!!!!")
}
```

Y si ejecutamos esto, podemos ver a `panic` en acción

```cmd
$ go run panic.go

panic: Ahhh!!!!

goroutine 1 [running]:
main.WillPanic(...)
        .../panic.go:8
main.main()
        .../panic.go:4 +0x38
exit status 2
```

Como era de esperar, nuestro programa imprimió el mensaje de panic, seguido por el stacktrace y luego terminó.

### Entonces, ¿qué hacer cuando ocurre un panic inesperado?

## Recover

Bueno, es posible recuperar el control de un programa en panic utilizando la función `recover` integrada, junto con la palabra reservada `defer`.

```go
func recover() interface{}
```

Probemos un ejemplo creando una función `handlePanic`. Y luego, podemos llamarlo usando `defer`.

```go
package main

import "fmt"

func main() {
	WillPanic()

	fmt.Println("Calm down, everything will be fine")
}

func handlePanic() {
	data := recover()
	fmt.Println("Recovered:", data)
}

func WillPanic() {
	defer handlePanic()

	panic("Ahhh!!!!")
}
```
```
$ go run panic.go

Recovered: Ahhh!!!!
Calm down, everything will be fine
```

Como podemos ver, nuestro `panic` se recuperó y ahora nuestro programa puede continuar su ejecución.

>Importante!. Panic y recover puede considerarse similar al try/catch de otros lenguajes. Pero un factor importante, según la filosofía de Go, es que debemos evitar el panic y el recover y usar errores cuando sea posible.

![exception](/go-training-beginner/modulo-3/1-panic_recover_defer/img/panic-exce.webp)

### Si es así, entonces, ¿cuándo deberíamos usar panic?

## Casos de uso

Hay dos casos de uso válidos para un panic:

- Un error irrecuperable

    Lo cual puede ser una situación en la que el programa no puede simplemente continuar con su ejecución.

    Por ejemplo, leer un archivo de configuración que es importante para iniciar el programa, ya que no hay nada más que hacer si la lectura del archivo falla.

- Error del desarrollador

    Por ejemplo, realizar una aserción confiando en que el resultado será correcto, esto es evitando el uso de la variable de control proporcionada por Go (f := i.(float64)). Esto causará un panic.
