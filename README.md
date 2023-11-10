# **C**ofre **S**agrado **V**irtual

## Descripción del problema

Trabajar con ficheros CSV puede suponer una tarea tediosa y complicada. Un valor
que falta en una columna supone que todas las filas posteriores no se pueden
leer, un cambio en el orden de las columnas o en el tipo de datos, pueden
existir datos con formatos incorrectos, etc. Por ello es necesario disponer de
alguna herramienta que sea capaz de detectar errores en los ficheros CSV y que
permita arreglarlos de forma sencilla, normalizando los datos si es necesario
para que elementos de una misma columna tengan el mismo formato o puedan
añadirse datos que faltan.

## Configuración inicial del proyecto

La configuración inicial de git y Github se describe
[aquí](doc/configuracion-inicial.md).

## Historias de usuario

Las historias de usuario se pueden consultar [aquí](doc/user-stories.md).

## Milestones

Las milestones se pueden encontrar [aquí](doc/milestones.md).

## Clase TipoDatosCliente

Esta clase representa la entidad principal del proyecto que contendrá la lógica
de negocio necesaria para la aplicación.

### Comprobación de sintaxis

Con el fin de verificar la precisión de la sintaxis en nuestra entidad y en las
distintas clases de nuestro proyecto, se puede emplear el siguiente comando:

```bash
go run ./build/ check
```

### Validación de la entidad

Para comprobar la validación de la entidad, se puede emplear el siguiente
comando que ejecuta los tests de la entidad:

```bash
go run ./build/ test
```

## Herramientas de desarrollo

Las diferentes herramientas de desarrollo utilizadas en este proyecto se pueden
encontrar [aquí](doc/tools.md).
