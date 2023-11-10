## Gestor de dependencias

En Go actualmente no hay mucha discusión sobre el gestor de dependencias a usar.
Desde la introducción de los módulos en Go 1.11, se ha convertido en el gestor
de facto. Sin embargo, existen otros gestores de dependencias que se pueden
usar, como [dep](https://github.com/golang/dep) o
[glide](https://github.com/Masterminds/glide), pero no se recomienda su uso.

## Gestor de tareas

En cuanto a gestores de tareas, hay una gran variedad de opciones, pero las más
populares son:

- [Make](https://www.gnu.org/software/make/) es el gestor de tareas más popular
  en el mundo de la programación. El 37% de los encuestados en la encuesta de
  [Developer Ecosystem 2022](https://www.jetbrains.com/lp/devecosystem-2022/go/)
  lo usan habitualmente.

- [Task](https://taskfile.dev/#/) es un gestor de tareas escrito en Go que se ha
  vuelto muy popular en los últimos años. Tiene una sintaxis en YAML

- [Mage](https://github.com/magefile/mage) es otro gestor de tareas, pero con la
  diferencia de que las tareas se escriben en Go y Mage las usa objetivos para
  ejecutar. La ventaja que tiene sobre las otras alternativas es que no hace
  falta introducir otro lenguaje para definir las tareas.

Sin embargo la opción por la que me voy a decantar es
[Goyek](https://github.com/goyek/goyek). Goyek es un gestor de tareas escrito en
Go similar a Mage, pero con ventajas a los mencionados anteriormente:

- Es independiente del sistema operativo y del shell.
- No es necesario definir etiquetas de construcción.
- No necesita instalación de un binario.
- Es fácil depurar, ya que es un programa de Go.
- No usa dependencias más allá de la librería estándar de Go.

Aunque Goyek es un proyecto relativamente nuevo, tiene más de 400 estrellas en
GitHub y una comunidad activa de usuarios. Por tanto, es una buena opción para
ser mi gestor de tareas.

## Comprobación de sintaxis

Go es un lenguaje compilado, por lo que la comprobación de sintaxis se hace en
el momento de la compilación. Sin embargo, hay herramientas que permiten
comprobar la sintaxis de los ficheros de Go sin necesidad de compilar el
programa. La herramienta más popular para esta tarea es
[golangci-lint](https://golangci-lint.run/).

Sin embargo, esta herramienta también se usa para comprobar otros aspectos del
código, actuando como *linter*. Por tanto, para la comprobación de sintaxis voy
a usar la herramienta [gofmt](https://golang.org/cmd/gofmt/), que es la que se
usa por defecto en Go y viene ya preinstalada.

## Framework para tests

Es importante que un proyecto realice test para comprobar que el código funciona
correctamente. Go cuenta con una herramienta para definir tests integrada en la
librería estándar, pero su funcionalidad es limitada, ya que no cuenta con
*assertions* y otras funciones. Por ello existen otros frameworks de test que
añaden funcionalidades extra. Los más populares son:

### [Testify](https://github.com/stretchr/testify)

- Proporciona una variedad de funciones de *assertions*, lo que facilita la
  escritura de tests claros y concisos.
- Ofrece soporte para declarar conjuntos de tests y funciones de limpieza.
- Funciona con el paquete de tests nativo de Go, lo que facilita su integración
  con su módulo de tests existente.
- Proporciona mensajes de error útiles que ayudan a identificar y corregir
  rápidamente problemas en el código.
- Fácil de aprender y utilizar.

### [Ginkgo](https://github.com/onsi/ginkgo)

- Proporciona una sintaxis estilo BDD, lo que facilita la escritura de tests que
  son legibles y fáciles de entender.
- Ofrece soporte para tests asíncronos, lo que facilita la evaluación de código
  que involucra concurrencia.
- Incluye soporte para aserciones anidadas, lo que facilita la organización y
  estructuración de los tests.
- Ofrece soporte para reportes personalizados, lo que permite personalizar la
  salida de los tests.

### Decisión para el proyecto

Aunque las librerías mencionadas anteriormente son muy buenas, el criterio que
voy a seguir para este proyecto va a ser el de las mejores prácticas para Go. En
este caso tenemos los [proverbios de Go](https://go-proverbs.github.io). Uno de
ellos dice lo siguiente:

> A little copying is better than a little dependency.

Incluir un framework de tests en el proyecto es una dependencia que no es
estrictamente necesario, ya que el módulo de tests de Go suele ser suficiente
para la mayoría de casos, y muchas funciones como las de *assertions* se pueden
programar manualmente.

Por tanto, no voy a incluir ningún framework de tests en el proyecto.

## Herramienta para ejecutar los tests

Go incluye una herramienta para ejecutar los tests:
[go test](https://pkg.go.dev/cmd/go/internal/test). Es la herramienta por
defecto y forma parte del estado del arte de Go. Por tanto, no es necesario
incluir otra herramienta para ejecutar los tests.
