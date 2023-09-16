# Configuración inical para la asignatura

## Creacion de claves privada/pública

En primer lugar he creado un par de claves (pública/privada) para usar con GitHub. He usado los siguientes comandos.

```terminal

$ ssh-keygen -t ed25519 -C "pablols114@gmail.com"

$ eval "$(ssh-agent -s)"

$ ssh-add ~/.ssh/id_ed25519

```

Por último desde el navegador web he añadido la clave pública a mi perfil de GitHub.

## Configuracion del usuario de Git

```terminal

$ git config user.email "pablols114@gmail.com"

$ git config user.name "Pablo Luque Salguero"

```

## Configuración del perfil de GitHub

[Imagen con la configuración del perfil](https://imgur.com/a/TNxhQPB) 
