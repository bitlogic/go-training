# ¿Es Go un lenguaje orientado a objetos?

![poo](/go-training-beginner/modulo-3/go_vs_poo/img/poo.png)

Para empezar este articulo quiero aclarar que se podría hablar mucho respecto al tema, pero creo con no es el momento ni lugar. Por lo que propongo ver algunas definiciones de los referentes y una pequeña comparación de lo que hay en Go con respecto a la programación orientada a objetos. 

## Definiciones de expertos

Alan Key (uno de los pionero en la programación orientada a objetos) basó el término `orientación a objeto`s en los siguientes conceptos:

- Mensajería. (Posible en Go gracias a los canales)
- Posibilidad de aislar, proteger o esconder el estado. (Podemos definir métodos o propiedades, como exportadas o no exportadas en Go)
- Asignación tardía de todas las cosas. (Esto es posible gracias a las funciones de orden superior, es decir funciones que reciben funciones o devuelven funciones, o bien mediante interfaces)

Como vemos Go cumple sobradamente las tres reglas de la orientación a objetos establecida por Alan Kay, así que podemos afirmar que Go ***puede ser*** un lenguaje orientado a objetos.

Por otro lado, sus creadores dicen los siguiente...

>¿Es Go un lenguaje orientado a objetos? Sí y no. A pesar de que Go tiene tipos y métodos y permite la programación orientada a objetos, no hay jerarquía de tipos. El concepto de “interfaz” en Go tiene una aproximación que creemos que es mucho más sencilla de usar y en algunos casos más general. También hay maneras de embeber tipos en otros tipos proporcionando algo similar – pero no idéntico – a la subclases. Además, los métodos en Go son más generals que en C++ o Java: se pueden definir para cualquier tipo de datos, incluso tipos integrados como enteros simples. No están restringidos a estructuras (clases). Además, la falta de jerarquía de tipos hace que los “objetos” en Go parezcan mucho más ligeros que en lenguajes como C++ o Java.

### Conclusión

Podemos decidir enfocar y utilizar Go como lenguaje OO pero tenemos que tener en cuenta que será al **estilo Go** y debemos aprender a convivir con ello.

## Vamos a ver sus principales diferencias.

### **Go no tiene clases**

Pero en cambio sí tiene `structs`, que aunque parezcan una clase no se comportan igual ni son lo mismo. Lo bueno es que podemos tomarlas como tal, ya que el struct solo contiene el estado y no el comportamiento pero podemos crear comportamientos para cambiar estos estados.

```go
type Employee struct {
	ID int
	Name string
}
```

### **Go no tiene herencia**

Pero sí tiene composición, que es patrón que nos permite generar bases con pequeñas piezas, un buen ejemplo es pensarlo como una receta, que dependiendo de cada ingrediente se puede crear un nuevo plato. Ahora esto no significa que con Go no se pueda lograr un comportamiento de herencia, en realidad si se puede.

La herencia sería algo como...

![herencia](/go-training-beginner/modulo-3/go_vs_poo/img/herencia.jpeg)

La composición sería...

![composi](/go-training-beginner/modulo-3/go_vs_poo/img/composi.png)

### **Encapsulamiento en Go**

La encapsulación en Go funciona a nivel de paquetes y además se realiza de manera implícita dependiendo de la primera letra del método o atributo. Así que si declaramos un método o atributo con la primera letra en mayúscula, quiere decir que está exportado, es decir, es público fuera del paquete. Go no implemente `protected` ya que no hay herencia.

Métodos públicos

```go
func NewPublicEmploye(id int, name string) *Employee {
	return &Employee{
		ID: id,
		Name: name,
	}
}
```

Métodos privados

func newPublicEmploye(id int, name string) *Employee {
	return &Employee{
		ID: id,
		Name: name,
	}
}

### **Interfaces**

Las interfaces en Go son implícitas, es decir si tienes los métodos de los que se compone la interfaz, implementas la interfaz. Nada de código extra!

Podemos seguir hablando de más diferencias en el tema de programación orientada a objetos en Go pero para empezar con esto ya tenemos una buena idea.
