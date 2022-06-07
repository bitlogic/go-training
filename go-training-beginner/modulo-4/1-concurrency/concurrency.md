# Concurrency

![concurrency](/go-training-beginner/modulo-4/1-concurrency/img/concurrency.png)

Una de las características mas poderosa de Go es la concurrencia. Pero...

## ¿Qué es la concurrencia?

La concurrencia, por definición, es la capacidad de descomponer un programa informático o algoritmo en partes individuales, que pueden ejecutarse de forma independiente.

El resultado final de un programa concurrente es el mismo que el de un programa que ha sido ejecutado secuencialmente.

Usando la concurrencia, podemos lograr los mismos resultados en menos tiempo, aumentando así el rendimiento general y la eficiencia de nuestros programas.

## Concurrencia vs Paralelismo

![c_vs_p](/go-training-beginner/modulo-4/1-concurrency/img/concurrency-vs-parallelism.png)

Mucha gente confunde concurrencia con paralelismo porque ambos implican ejecutar código simultáneamente, pero son dos conceptos completamente diferentes.

**La concurrencia es la tarea de ejecutar y administrar varios cálculos al mismo tiempo, mientras que el paralelismo es la tarea de ejecutar varios cálculos simultáneamente.**

Una cita de *Rob Pike* lo resume bastante bien.

>"La concurrencia se trata de manejar muchas cosas a la vez. El paralelismo se trata de hacer muchas cosas a la vez"

Para aprovechar todo el poder de Go, primero tenemos que entender cómo Go hace la ejecución simultánea de código. Go se basa en un modelo de concurrencia llamado CSP (communicating sequential processes).

La [Comunicación de procesos secuenciales](https://dl.acm.org/doi/10.1145/359576.359585]) es un modelo presentado por *Tony Hoare* en 1978 que describe las interacciones entre procesos concurrentes.

La concurrencia es difícil, pero CSP nos permite dar una mejor estructura a nuestro código concurrente y proporciona un modelo para pensar en la concurrencia de una manera que lo hace un poco más fácil. En Go, los procesos son independientes y se comunican compartiendo canales entre ellos.

![csp](/go-training-beginner/modulo-4/1-concurrency/img/chan.png)

>Aprenderemos cómo lo implementa Golang utilizando goroutines y canales más adelante en el curso.

## Conceptos básicos

Ahora, revisemos algunos conceptos básicos de concurrencia.

### Carrera de datos (Data race)

Un `data race` ocurre cuando los procesos tienen que acceder al mismo recurso al mismo tiempo.

Por ejemplo, un proceso lee mientras otro escribe simultáneamente, exactamente en el mismo recurso.

### Condiciones de carrera (Race conditions) 

Una `race conditions` ocurre cuando el tiempo o el orden de los eventos afecta la ejecución de un fragmento de código.

### Interbloqueos (Deadlocks)

Se produce un `deadlocks` cuando todos los procesos se bloquean mientras se esperan unos a otros y el programa no puede continuar.

#### **Condiciones de Coffman**

Hay cuatro condiciones, conocidas como `condiciones de Coffman`, todas ellas deben cumplirse para que se produzca un `deadlocks`.

- Mutual Exclusion

Un proceso concurrente contiene al menos un recurso en cualquier momento, lo que lo hace no compartible.

En el diagrama ade abajo, hay una sola instancia del recurso 1 y solo está en manos del proceso 1.

![exclusión mutua](/go-training-beginner/modulo-4/1-concurrency/img/mutual-exclusion.png)

- Hold and wait

Un proceso concurrente tiene un recurso y está esperando un recurso adicional.

En el diagrama que se muestra a continuación, el proceso 2 tiene el recurso 2 y el recurso 3 y solicita el recurso 1 que está en manos del proceso 1.

![esperar y esperar](/go-training-beginner/modulo-4/1-concurrency/img/hold-and-wait.png)

- No preemption (Sin preferencia)

El sistema no puede quitar un recurso retenido por un proceso concurrente. Sólo puede ser liberado por el proceso que lo sostiene.

En el siguiente diagrama, el proceso 2 no puede apropiarse del recurso 1 del proceso 1. Solo se liberará cuando el proceso 1 lo abandone voluntariamente después de que se complete su ejecución.

![sin preferencia](/go-training-beginner/modulo-4/1-concurrency/img/no-preemption.png)

- Circular wait (Espera circular)

Un proceso está esperando el recurso que tiene el segundo proceso, el cual está esperando el recurso que tiene el tercer proceso, y así sucesivamente, hasta que el último proceso está esperando un recurso que tiene el primer proceso. Formando una cadena circular.

En el siguiente diagrama, al proceso 1 se le asigna el recurso 2 y solicita el recurso 1. De manera similar, al proceso 2 se le asigna el recurso 1 y solicita el recurso 2. Esto forma un ciclo de espera circular.

![espera circular](/go-training-beginner/modulo-4/1-concurrency/img/circular-wait.png)

### Livelocks

Livelocks son procesos que realizan activamente operaciones concurrentes, pero estas operaciones no hacen nada para hacer avanzar el estado del programa.

### Starvation (Inanición)
Una inanición ocurre cuando un proceso se ve privado de los recursos necesarios y no puede completar su función.

La inanición puede ocurrir debido a deadlocks o algoritmos de programación ineficientes para los procesos. Para resolver una inanición, necesitamos emplear mejores algoritmos de asignación de recursos que aseguren que cada proceso obtenga su parte justa de recursos.