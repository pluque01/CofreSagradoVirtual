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

## Imagen de Docker

### Criterios de elección

- **Huella de memoria pequeña**: la imagen debe ser lo más pequeña posible para
  disminuir el tiempo necesario para descargarla y ejecutarla.

- **Seguridad**: la imagen debe tener el menor número de vulnerabilidades
  posible.

- **Soporte**: la imagen debe estar mantenida y actualizada.

### Opciones a considerar

- [golang](https://hub.docker.com/_/golang): imagen oficial de Docker para Go y
  es mantenida por la comunidad de Docker. Es una imagen muy popular, con más de
  mil millones de descargas. Al ser la imagen oficial de Docker, está muy bien
  mantenida y actualizada. Tiene varias variantes disponibles:

  - `golang`: es la imagen por defecto, basado en Debian.

  - `golang:alpine`: basado en Alpine Linux, una distribución muy ligera de
    Linux que tiene un tamaño de 5MB. Esta imagen tiene un número mínimo de
    herramientas, dejando que sea el usuario el que instale las herramientas que
    necesite.

- [alpine](https://hub.docker.com/_/alpine): imagen de Alpine Linux, una
  distribución muy ligera de Linux que tiene un tamaño de 5MB. Esta imagen, sin
  embargo, tiene un número mínimo de herramientas, dejando que sea el usuario el
  que instale las herramientas que necesite.

- [bitnami](https://hub.docker.com/r/bitnami/golang): imagen de Bitnami para Go.
  Son builds automatizadas con los últimos cambios. Están basadas en Minideb,
  una distribución ligera de Debian.

- [debian](https://hub.docker.com/_/debian): Imagen de la distribución Debian.
  No incluye las herramientas para trabajar con Go por defecto, pero tiene una
  versión `slim` con un tamaño de memoria reducido.

### Decisión para el proyecto

Siguiendo los criterios de elección, la imagen que voy a usar para el proyecto
va a ser la imagen de Debian en su versión `slim`. Esta imagen tiene un tamaño
reducido al igual que la imagen de Alpine, pero la de Alpine tiene alguna
vulnerabilidad crítica, mientras que la de Debian solo tiene
[vulnerabilidades leves](https://hub.docker.com/_/debian/tags?page=1&name=stable-slim).

Sin embargo, la imagen de Debian no incluye las herramientas para trabajar con
Go, por lo que se tendrá que instalar manualmente. Para ello usaré un
[`multi-stage build`](https://docs.docker.com/build/building/multi-stage/) para
copiar las herramientas desde una imagen de Go a la imagen final.

Por ello usaré las siguientes imágenes:

- `golang:latest` como fuente donde copiar las herramientas de Go. Nos interesa
  que sea la versión latest para comprobar nuestro código con la última versión
  de Go disponible.
- `debian:stable-slim` como imagen final.

## Herramientas de integración continua

### Criterios de elección

- **Compatibilidad con el lenguaje**: Necesito asegurarme de que el sistema sea
  compatible con Golang.

- **Integración con herramientas existentes:** El sistema debe integrarse sin
  problemas con las herramientas que uso actualmente, como Github y que tengan
  disponible el uso de la API Checks de Github.

- **Compatibilidad con contenedores:** Dado que estoy utilizando contenedores
  Docker, es esencial que el sistema sea compatible con estas tecnologías.

- **Coste:** Debo considerar el coste asociado con al uso del sistema.

- **Servicio disponible en la nube:** Es mejor que el sistema sea un servicio
  disponible en la nube, ya que no tengo que preocuparme de su mantenimiento ni
  por su infraestructura.

### Opciones a considerar

#### [Jenkins](https://www.jenkins.io)

Jenkins es una opción sólida de integración continua que destaca por su
naturaleza de código abierto y su capacidad altamente extensible. Es compatible
con Golang y ofrece una amplia integración con diferentes herramientas gracias a
su extenso conjunto de plugins. La configuración se simplifica mediante una
interfaz web, facilitando la adaptación a las necesidades del proyecto. Jenkins
cuenta con una comunidad activa y una diversidad de recursos disponibles.

#### [Travis CI](https://travis-ci.com)

Travis CI se destaca por su enfoque en proyectos de código abierto y su
integración directa con GitHub, proporcionando una configuración basada en
archivos YAML. Compatible con Golang, ofrece soporte para varios lenguajes de
programación. Además, su entorno de ejecución en contenedores Docker brinda
coherencia y aislamiento en el proceso de integración continua. Ofrece un
periodo de prueba tras el cual es necesario pagar para seguir usando el sistema.

#### [CircleCI](https://circleci.com)

CircleCI es reconocido por su capacidad para ejecutar pruebas en paralelo,
acelerando significativamente el proceso de integración continua. Compatible con
Golang, se integra con GitHub y Bitbucket. Su configuración se basa en archivos
YAML, y su compatibilidad con contenedores Docker facilitan la personalización y
la integración con entornos basados en Docker. CircleCI ofrece un plan gratuito
con 6000 minutos de ejecución de trabajos al mes, lo que lo hace ideal para
proyectos pequeños.

#### [GitHub Actions](https://docs.github.com/es/actions)

GitHub Actions se destaca por su integración directa con la plataforma GitHub,
proporcionando una solución completa para la integración continua y entrega
continua (CI/CD). Utiliza archivos con formato `yaml` para su configuración por
lo que no es tan visual como otras alternativas. Ofrece una amplia biblioteca de
acciones predefinidas y la posibilidad de crear acciones personalizadas para
adaptarse a necesidades específicas del proyecto. Para repositorios públicos es
gratuito, mientras que para repositorios privados se aplican límites de uso.

#### [Semaphore](https://semaphoreci.com)

Semaphore es reconocido por su compatibilidad con Golang y una variedad de
lenguajes. Se integra con plataformas populares como GitHub y Bitbucket,
permitiendo una configuración sencilla. La ejecución de trabajos en contenedores
Docker facilita la integración con entornos basados en Docker. Tiene disponible
un período de prueba tras el cual hay que abonar un importe.
