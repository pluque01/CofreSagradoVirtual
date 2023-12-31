# Herramientas

- [Gestor de dependencias](#gestor-de-dependencias)
- [Gestor de tareas](#gestor-de-tareas)
- [Comprobación de sintaxis](#comprobaci%C3%B3n-de-sintaxis)
- [Framework para tests](#framework-para-tests)
- [Herramienta para ejecutar los tests](#herramienta-para-ejecutar-los-tests)
- [Imagen de Docker](#imagen-de-docker)
- [Herramientas de integración continua](#herramientas-de-integraci%C3%B3n-continua)
- [Herramientas de logs](#herramientas-de-logs)
- [Herramienta de configuración](#herramienta-de-configuraci%C3%B3n)
- [Framework REST](#framework-rest)

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

### Decisión para el proyecto

Las opciones que solo ofrecen periodos de prueba las voy a descartar, ya que no
tengo pensado invertir dinero en este proyecto. Por tanto, las únicas opciones
disponibles son CircleCI y GitHub Actions. Entre estas dos, la opción más
recomendable es GitHub Actions, no por tener un número de minutos ilimitados,
sino porque ya lo estoy usando para construir el contenedor de Docker y subirlo
a Docker Hub.

Sin embargo, voy a probar ambas opciones.

- He intentado configurar CircleCI, pero cuando iba a crear el workflow he ido a
  la documentación para ver como podía integrarlo con la API Checks de Github y
  he visto que no es posible
  [por algún motivo que no logro entender](https://circleci.com/docs/enable-checks/).
  Por tanto, descarto CircleCI.

- Github Actions: Con Github Actions no he tenido ningún problema. He creado un
  workflow que ejecuta los tests en dos versiones de Go, la 1.21 y la 1.20,
  ambas en una máquina con Ubuntu.

- AppVeyor: esta opción me la recomendó @JJ como alternativa a CircleCI, ya que
  yo no la había considerado en un primer momento. Ha sido más complicado de
  configurar que GitHub Actions, al no haber encontrado una documentación tan
  completa. Sin embargo, he conseguido configurarla y ejecuta los tests para las
  versiones 1.21 y 1.20 de Go en Windows.

Nota: al final no he usado el contenedor de Docker, ya que este no me permite
probar las distintas versiones de Go.

## Herramientas de logs

### Criterios de elección

- **Adaptación a las necesidades del proyecto:** La herramienta que elija debe
  de adaptarse a las necesidades presentes y futuras del proyecto.

- **Logs estructurados:** La herramienta debe permitir el uso de logs
  estructurados. Los logs estructurados son aquellos que tienen un formato
  definido, mientras que los logs no estructurados no tienen un formato
  definido. Los logs estructurados se pueden almacenar en ficheros con
  clave-valores como JSON o YAML, mientras que los no estructurados se suelen
  almacenar en ficheros de texto plano.

- **Seguridad:** La herramienta debe tener una puntuación aceptable en
  [Snyk Advisor](https://snyk.io/advisor/).

### Opciones a considerar

- [Módulo log de Go](https://pkg.go.dev/log): el módulo log de Go es un módulo
  de la librería estándar que permite registrar mensajes de log. Para usar logs
  estructurados hay que usar el módulo [slog](https://pkg.go.dev/log/slog), pero
  este solo está disponible en la versión 1.21 de Go.

- [Logrus](https://github.com/sirupsen/logrus) es una librería de Go que permite
  registrar mensajes de log. Es una librería muy popular. Tiene soporte para
  logs estructurados y permite la creación de hooks para enviar los logs a
  distintos destinos. Sin embargo, está en modo de mantenimiento, por lo que no
  recibirá nuevas funcionalidades.

- [Zap](https://pkg.go.dev/go.uber.org/zap): desarrollada por el equipo de Uber,
  ofrece *blazing fast, structured, level logging in go*. Es una librería muy
  conocida y una buena opción para obtener logs estructurados.

- [Zerolog](https://github.com/rs/zerolog): inspirada en Zap, es una librería
  muy popular para logs estructurados. Es una buena alternativa a Zap ya que
  asegura ofrecer mayor rendimiento.

### Decisión para el proyecto

Todas las opciones ofrecen logs estructurados, que creo que es necesario para
que el proyecto sea escalable.

Si el módulo **slog** estuviera disponible en la versión 1.20 de Go, lo usaría
para el proyecto, al ser un módulo de la librería estándar y no requiere
dependencias externas. Sin embargo, tengo que elegir uno entre los otros tres.
**Logrus** lo voy a descartar por estar en modo de mantenimiento. Entre **Zap**
y **Zerolog** voy a usar **Zerolog**, ya que ofrece mayor rendimiento y también
tiene 95 puntos en
[Snyk Advisor](https://snyk.io/advisor/golang/github.com/rs/zerolog).

## Herramienta de configuración

### Criterios de elección

- **Seguridad:** Siendo una herramienta que va a distribuir información
  sensible, es necesario que sea segura.

- **Soporte para configuraciones clave-valor distribuidos:** La herramienta debe
  permitir el uso de servicios de distribución de configuraciones.

### Opciones a considerar

- [Viper](https://github.com/spf13/viper): Suite completa de gestión de
  configuraciones para Go. Herramienta muy popular y soporta todo tipo de
  configuraciones, desde variables de entorno a sistemas de distribución como
  etcd o Consul.

- [etcd/clientv3](https://pkg.go.dev/go.etcd.io/etcd/client/v3@v3.5.10): Cliente
  oficial de etcd3 para Go. No hay mucho más que decir salvo que tiene una
  puntuación de 100/100 en Snyk Advisor. Etcd es el sistema de distribución que
  se usa en
  [Kubernetes](https://etcd.io/docs/v3.3/production-users/#all-kubernetes-users).

- [Consul](https://pkg.go.dev/github.com/hashicorp/consul/api): Cliente oficial
  de Consul para Go. También tiene una puntuación perfecta en Snyk Advisor.
  Actualmente es propiedad de Hashicorp, la empresa que desarrolla Terraform.

- [Koanf](https://github.com/knadh/koanf): Alternativa a Viper con la ventaja de
  que funciona por módulos que son independientes entre sí. Esto permite que se
  puedan importar solo los módulos que se vayan a utilizar sin tener que usar el
  resto de módulos. Al igual que Viper, tiene soporte para sistemas de
  clave-valor distribuidos como etcd3.

### Decisión para el proyecto

Viendo las opciones que tengo disponibles creo que me voy a decantar por
**Koanf** al poder trabajar con diferentes módulos de forma independiente. Tiene
una puntuación más baja en Snyk Advisor, pero no es debido a problemas de
seguridad, sino a que tiene menos contribuciones y menos uso.

## Framework REST

### Criterios de elección

- **Simplicidad:** La aplicación que estamos desarrollando es relativamente
  sencilla, por lo que no necesito un framwork muy complejo. Otro de los motivos
  para elegir un framework simple es el grado de abstracción que conlleva el uso
  de un framework complejo. En este caso, al ser un proyecto de aprendizaje, es
  mejor usar un framework simple para entender mejor como se implementan los
  servidores webs.

- **Seguridad:** Aunque el framework vaya a ser simple, no debe comprometer la
  seguridad de la aplicación.

### Opciones a considerar

- **[Gin](https://github.com/gin-gonic/gin):** 73 mil estrellas en Github (wow).
  Ofrece un rendimiento *blazing-fast ©️*, huella de memoria pequeña, un
  middleware fácil de usar y una documentación extensa. Puntuación de 86 en Snyk
  Advisor.

- **[Echo](https://github.com/labstack/echo):** Aún más rendimiento que Gin
  según sus creadores. Framework de middleware extensible, funciones integradas
  para una variedad de códigos HTTP, certificados TLS automáticos con Let´s
  Encrypt. Tiene una puntuación de 94 en Snyk Advisor.

- **[Chi](https://github.com/go-chi/chi)**: Framework ligero, rápido y 100%
  compatible con el módulo `net/http` de Go. Permite el uso de una gran cantidad
  de middlewares. Está en producción en empresas como Cloudflare o Heroku. Tiene
  una puntuación de 94 en Snyk Advisor.

- **[Módulo `net/http` de Go](https://pkg.go.dev/net/http):** Módulo de la
  biblioteca estándar de Go. Es la base sobre la que parten muchos de los
  módulos mencionados, agregando nuevas funcionalidades. No tiene puntuación en
  Snyk Advisor, pero se presupone su seguridad al ser un módulo de la biblioteca
  estándar.

### Decisión para el proyecto

Siguiendo los criterios de elección, voy a descartar Gin y Echo por su alto
grado de abstracción y su extensa cantidad de funcionalidades. Chi sería una
opción muy buena si el tamaño del proyecto fuese algo mayor, así que siguiendo
los criterios y los [proverbios de Go](https://go-proverbs.github.io) ©️, voy a
utilizar el módulo `net/http` de Go.
