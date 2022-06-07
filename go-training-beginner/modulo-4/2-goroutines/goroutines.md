# Goroutines

![goroutines](/go-training-beginner/modulo-4/2-goroutines/img/gor.jpeg)

Una de las frases mas conocidas de *Rob Pike* sobre estos temas es:

> "Don't communicate by sharing memory, share memory by communicating." - Rob Pike

## ¿Qué es una goroutine?

Una goroutine es un subproceso ligero de ejecución que es administrado por el tiempo de ejecución de Go y esencialmente nos permite escribir código asíncrono de manera sincrónica.

Es importante saber que no son subprocesos reales del sistema operativo y que la función main en sí se ejecuta como una goroutine.

Un solo subproceso puede ejecutar miles de goroutines mediante el *Go runtime scheduler* (programador de tiempo de ejecución de Go) que utiliza la programación cooperativa. Esto implica que si la goroutine actual está bloqueada o se ha completado, el programador moverá las otras goroutines a otro subproceso del sistema operativo. Por lo tanto, logramos eficiencia en la programación donde ninguna rutina se bloquea para siempre.

Podemos convertir cualquier función en una goroutine simplemente usando la palabra clave go.

```go
go fn(x, y, z)
```

Antes de escribir cualquier código, es importante describir brevemente el modelo de *bifurcación-unión*.

## Fork-Join Model (Modelo de bifurcación-unión)

Go usa la idea del modelo de concurrencia fork-join detrás de las goroutines. Este modelo esencialmente implica que un proceso secundario se separa de su proceso principal para ejecutarse simultáneamente con el proceso principal. Después de completar su ejecución, el proceso secundario se fusiona nuevamente con el proceso principal. El punto donde se une se llama punto de unión.

![Fork-Join](/go-training-beginner/modulo-4/2-goroutines/img/fork-join.png)

Escribamos algo de código y creemos nuestra propia goroutine.

```go
package main

import "fmt"

func speak(arg string) {
	fmt.Println(arg)
}

func main() {
	go speak("Hello World")
}
```

La función `speak` tiene el prefijo de la palabra reservada `go`. Esto permitirá que se ejecute como una goroutine separada. Y eso es todo, acabamos de crear nuestra primera goroutine. ¡Es así de simple!

Pero, ejecutemos esto

```cmd
$ go run goroutines.go

```

Parecería que nuestro programa no se ejecutó completamente porque no vemos la salida. Esto se debe a que nuestra rutina principal (main) salió y no esperó a la rutina que creamos.

¿Qué pasa si hacemos que el programa espere usando la  función `time.Sleep`?

```go
func main() {
	...
	time.Sleep(1 * time.Second)
}
```

Y ahora, si ejecutamos esto

```cmd
$ go run goroutines.go

Hello World
```

Ahora podemos ver la salida completa ahora.

Bien, esto funciona, pero no es lo ideal. Entonces, ¿cómo mejoramos esto?

Bueno, la parte más complicada de usar goroutines es saber cuándo se detendrán. Es importante saber que las goroutines se ejecutan en el mismo espacio de direcciones, por lo que se debe sincronizar el acceso a la memoria compartida.

