# fmt package

![fmt](/go-training-beginner/modulo-1/3-fmt_package/img/fmt.png)

fmt package tiene un montón de funciones dentro pero ahora nos vamos a centrar en las funciones que se usan con mas frecuencias. Como `fmt.Print`.

```go
...

fmt.Print("What", "is", "your", "name?")

fmt.Print("My", "name", "is", "golang")
...
```
```cmd 
$ go run main.go

Whatisyourname?Mynameisgolang
```
Como se puede ver `Print` no formatea el string de ninguna manera, solamente lo imprime.

Ahora, si usamos `Println`, no solo agrega un salto de línea, sino que agrega espacios entre los strings.

```go
...

fmt.Print("What", "is", "your", "name?")

fmt.Print("My", "name", "is", "golang")
...
```
```cmd
$ go run main.go

What is your name?
My name is golang
```

El siguiente es `Printf` conocido como "*print formatter*" que nos permite dar formato a strings, enteros, booleanos, etc.

```go
...
name := "golang"

fmt.Println("What is your name?")

fmt.Printf("My name is %s", name)
...
```

```cmd
$ go run main.go

What is your name?
My name is golang
```

Lo nuevo que vemos es `%s` que es reemplazado por `name`.

>Qué significa `%s`? Se llaman verbos de anotación y le indican a la función cómo formatear los argumentos. Podemos controlar cosas como el ancho, los tipos y la precisión, entre otras cosas y hay muchos de estos verbos. Dejo este [link](https://pkg.go.dev/fmt) a la documentación para ver la lista completa

Veamos alguno ejemplos.

```go
...
percent := 10.0 / 3.0

fmt.Printf("%f", percent)
...
```

```cmd
$ go run main.go

3.333333
```

Si quisiéramos una precisión de 2 decimales y agregar el signo de porcentaje, podemos formatear de la siguiente manera

```go
...
percent := 10.0 / 3.0

fmt.Printf("%.2f %%", percent)
...
```

```cmd 
$ go run main.go

33.33 %
```

Notemos el uso de `.2f` y el `%` para "escapear" del formato

>No te preocupes por aprender todas la posibilidades de formateo ahora, con la practica se va aprendiendo.

Por ultimo pero no menos importante, quizás mas importante, tenemos las funciones de `Sprint`, `Sprintln`, `Sprintf`.

Básicamente funcionas igual que las anteriores, la única diferencia es que devuelven el string formateado en vez de imprimirlo.

```go
...
s := fmt.Sprintf("hex:%x bin:%b", 10 ,10)

fmt.Println(s)
...
```
```cmd
$ go run main.go

hex:a bin:1010
```

Vemos como `Sprintf` formatea neutros enteros hacia un hexadecimal y un binario, y devuelve un string.

