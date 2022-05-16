# Functions

![func](/go-training-beginner/modulo-1/5-functions/img/func.png)
## Declaración y ejecución simple

***Declaración***
```go
func myfunction() {}
```
***Ejecución***
```go
...

myfunction()

...
```
 Podemos pasarle un parámetro
 ```go
 func main() {
	myFunction("Hello")
}

func myFunction(p1 string) {
	fmt.Printtln(p1)
}
```
 o muchos parámetros. Si son del mismo tipo, podemos hacerlo de forma consecutiva, como por ejemplo...

 ```go
 func myNextFunction(p1, p2 string) {}
```

## Devolviendo valores

Podemos devolver un valor
```go 
func main() {
	s := myFunction("Hello")
	fmt.Println(s)
}

func myFunction(p1 string) string {
	msg := fmt.Sprintf("%s function", p1)
	return msg
}
```
o muchos valores... una gran ventaja de Go
```go
func main() {
	s, i := myFunction("Hello")
	fmt.Println(s, i)
}

func myFunction(p1 string) (string, int) {
	msg := fmt.Sprintf("%s function", p1)
	return msg, 10
}
```

también podemos darle nombres a eso return.
```go
func myFunction(p1 string) (s string, i int) {
	s = fmt.Sprintf("%s function", p1)
	i = 10

	return
}
```
 Prestemos atención a dos cosas en el código anterior. La primera es como se usa el `return` sin ningún argumento, y la segunda como no se declaran variables nuevas, ya que están declaradas en la misma declaración de la función.

 >Por experiencia, esta funcionalidad parece muy buena, pero se presta a mucha confusión cuando la función es muy grande. Diría que se pierde un poco de legibilidad en esos casos.

 ### Funciones como valores

 En Go la funciones son como valores, por ende podemos usarlas como tal.

 ```go
 func myFunction() {
	fn := func() {
		fmt.Println("inside fn")
	}

	fn()
}
```

también las podemos tratar como funciones anónimas. Si eliminamos el `fn`, lo conseguimos.
```go
func myFunction() {
	func() {
		fmt.Println("inside fn")
	}()
}
```

>Notemos como usamos los `()` para ejecutar las funciones.

## Closures

Una definición simple puede ser que ***closure*** es una función que hace referencia a variables desde fuera de su cuerpo.

```go
func myFunction() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v

		return sum
	}
}
```
```go
...
add := myFunction()

add(5)

fmt.Println(add(10))
...
```
```cmd
$ go run main.go

15
```
Como podemos ver, el resultado es 15 porque la variable `sum` esta vinculada a la función.

### Funciones variádicas (Variadic Functions) 

Una misma función puede recibir muchos parámetros o ninguno usando el operador `...`

```go
func main() {
	sum := add(1, 2, 3, 5)
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

En este caso el parámetro `values` se convierte en un array de int. 
Con `range` podemos recorrer ese array, pero no te preocupes que lo vemos mas adelante.

>La función `fmt.Println` es una *variadic function* ya que puede recibir multiples argumentos.













