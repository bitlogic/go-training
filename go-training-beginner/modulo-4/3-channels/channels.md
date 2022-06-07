# Channels

![canal](/go-training-beginner/modulo-4/3-channels/img/chan.jpeg)

## ¿qué son los canales?
Bueno, simplemente definido.

Un canal es una tubería de comunicación entre goroutines. Las cosas entran por un extremo y salen por otro en el mismo orden hasta que se cierra el canal.

![canal](/go-training-beginner/modulo-4/3-channels/img/channel.png)

## Creando un canal

Ahora que entendemos qué son los canales, veamos cómo podemos declararlos.

```go
var ch chan T
```

Anteponemos nuestro tipo de dato, `T`, que es el tipo de datos del valor que queremos enviar y recibir, con la palabra reservada `chan` que representa un canal.

>Por un mismo canal solo puede viajar un tipo de dato. Es decir, por el mismo canal no podemos enviar un int y luego un string, por ejemplo.

Intentemos imprimir el valor del canal `c` de tipo `string`.

```go
func main() {
	var ch chan string

	fmt.Println(ch)
}
```
```cmd
$ go run channels.go

<nil>
```
Como podemos ver, el valor cero de un canal es `nil` y si intentamos enviar datos a través del canal, el programa nos devolverá un `panic`.

Entonces, de manera similar a los *slices*, podemos inicializar el canal usando la función `make()`.

```go
func main() {
	ch := make(chan string)

	fmt.Println(ch)
}
```
Y si ejecutamos esto, podemos ver que el canal fue inicializado correctamente.

```
$ go run channels.go

0x1400010e060
```

## Envío y recepción de datos

Ahora que conocemos la sintaxis básica de los canales, implementemos el [ejemplo anterior](/go-training-beginner/modulo-4/2-goroutines/code/goroutines.go) usando canales y asi aprender cómo usarlos para comunicarnos entre goroutines.

```go
package main

import "fmt"

func speak(arg string, ch chan string) {
	ch <- arg // Send
}

func main() {
	ch := make(chan string)

	go speak("Hello World", ch)

	data := <-ch // Receive
	fmt.Println(data)
}
```

Observemos la sintaxis para enviar datos usando `channel <- data` y recibir datos usando la `data := <-channel`.

Y si ejecutamos esto

```cmd
$ go run main.go

Hello World
```

Perfecto, el programa funcionó como esperábamos.

## Canales almacenados en buffer

También tenemos canales almacenados en buffer que aceptan un número limitado de valores sin un receptor correspondiente para esos valores.

![canal-búfer](/go-training-beginner/modulo-4/3-channels/img/buffered-channel.png)

Esta *longitud o capacidad* del buffer se puede especificar mediante el segundo argumento de la función `make`.

```go
func main() {
	ch := make(chan string, 2)

	go speak("Hello World", ch)
	go speak("Hi again", ch)

	data1 := <-ch
	fmt.Println(data1)

	data2 := <-ch
	fmt.Println(data2)
}
```

Debido a que este canal está almacenado en buffer, podemos enviar estos valores al canal sin una recepción concurrente correspondiente. Esto significa que envía a un bloqueo de canal con buffer solo cuando el buffer está lleno y recibe el bloqueo cuando el buffer está vacío.

De forma predeterminada, un canal no tiene buffer y tiene una capacidad de 0, por lo tanto, omitimos el segundo argumento de la función `make`.


## Canales direccionales

Cuando usamos canales como parámetros de función, podemos especificar si un canal está destinado solo a enviar o recibir valores. Esto aumenta la seguridad de nuestro programa ya que, de forma predeterminada, un canal puede enviar y recibir valores.

![canales-direccionales](/go-training-beginner/modulo-4/3-channels/img/directional-channels.png)

En el ejemplo, podemos actualizar el segundo argumento de la función `speak` para que solo pueda enviar un valor.

```go
func speak(arg string, ch chan<- string) {
	ch <- arg // Send Only
}
```

Acá, `chan<-` solo se puede usar para enviar valores y entrará en `panic` si intentamos recibir valores.

## Cierre de canales

Además, al igual que con cualquier otro recurso, una vez que hayamos terminado con nuestro canal, debemos cerrarlo. Esto se puede lograr usando la función `close`.

Podemos simplemente pasar la función `close` al canal.

```go
func main() {
	ch := make(chan string, 2)

	go speak("Hello World", ch)
	go speak("Hi again", ch)

	data1 := <-ch
	fmt.Println(data1)

	data2 := <-ch
	fmt.Println(data2)

	close(ch)
}
```

Opcionalmente, los receptores pueden comprobar si se ha cerrado un canal asignando un segundo parámetro a la expresión de recepción.

```go
func main() {
	ch := make(chan string, 2)

	go speak("Hello World", ch)
	go speak("Hi again", ch)

	data1 := <-ch
	fmt.Println(data1)

	data2, ok := <-ch
	fmt.Println(data2, ok)

	close(ch)
}
```

Si `ok` es `false`, no hay más valores para recibir y el canal se cierra.

En cierto modo , esto es similar a cómo comprobamos si existe o no una clave en un mapa.

## Propiedades

Por último, analicemos algunas propiedades de los canales.

 - Si enviamos un dato a un canal no inicializados (nil) se bloqueará la ejecución del programa para siempre.

```go
var c chan string

c <- "Hello, World!" // Panic: all goroutines are asleep - deadlock!
```
- Si recibimos un dato de un canal no inicializados (nil) se bloqueará la ejecución del programa para siempre.

```go
var c chan string

fmt.Println(<-c) // Panic: all goroutines are asleep - deadlock!
```
- Si enviamos un dato a un canal cerrado, nos arrojará un panic

```go
var c = make(chan string, 1)

c <- "Hello, World!"
close(c)
c <- "Hello, Panic!" // Panic: send on closed channel
```

- Si recibimos un dato de un canal cerrado, nos devolverá el valor cero inmediatamente

```go
var c = make(chan int, 2)

c <- 5
c <- 4

close(c)

for i := 0; i < 4; i++ {
    fmt.Printf("%d ", <-c) // Output: 5 4 0 0
}
```

- Range en channels

También podemos usar `for` y `range` para iterar sobre los valores recibidos de un canal.

```go
package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	close(ch)

	for data := range ch {
		fmt.Println(data)
	}
}
```