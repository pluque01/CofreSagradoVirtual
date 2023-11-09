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
