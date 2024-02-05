
<p align="center">
  <a href="http://nestjs.com/" target="blank"><img src="https://pkg.go.dev/static/shared/logo/go-blue.svg" width="200" alt="Nest Logo" /></a>
</p>
  <h1 align="center">Task Rest API</h1>
  <p align="center">API REST para una aplicación de tareas desarrollada con go y mux (router) para la gestión de tareas y autenticación de usuarios. Integra el uso de JWT para la seguridad de acceso a los endpoints de la aplicación.
</p>

## :ledger: Index

- [Pre-Requisitos](#pre-requisitos-)
- [Instalación](#instalación-)
- [Environment](#environment)
- [Base de datos](#base-de-datos)
- [Construido](#construido-con-%EF%B8%8F)

## Comenzando 🚀

_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas._

## Pre-Requisitos 📋

_Software requerido_

```
Go >= 1.21.6
Docker >= 3.8
PostgreSql >= 15.3
```

_Software opcional_

```
Visual Studio Code ( O el editor de su preferencia)
```

## Instalación 🔧

_Para ejecutar un entorno de desarrollo_

_Previamente ejecutar el comando en la terminal para descargar los "packages" para el funcionamiento del proyecto_

```
 go mod download
```

_Previamente a ejecutar el servidor en desarrollo configurar el archivo .env con las variables de entorno necesarias y levantar el contenedor para la base de datos , ejecutar :_

```
 go run .
```

_Dirigirse a la ruta http://localhost:4000/ donde tendra acceso a la API_

## Environment

_Se tiene el archivo `env.template` , el cual posee un ejemplo de cada valor de las valores de entorno para poder desplegarlas en nuestro propio ambiente local_

```
TOKEN_SECRET_KEY=<YOUR_SECRET_KEY>
```

## Base de datos

_Se tiene el archivo `docker-compose.yml` , el cual posee la configuración necesaria para leventar el motor de base de datos para el uso de la aplicación, ejecutar:_

```
docker-compose up -d
```

_Diagrama de la base de datos_

![Database Diagram](/docs/db/dbdiagram.png)

## Construido con 🛠️

_Las herramientas utilizadas son:_

- [Go](https://go.dev/) - Go es un lenguaje de programación de código abierto que simplifica la creación de sistemas seguros y escalables.
- [Mux](https://github.com/gorilla/mux) - Potente enrutador HTTP y comparador de URL para crear servidores web Go.
- [Docker](https://www.docker.com/) - Para el despliegue de aplicaciones basado en contenedores
- [PostgreSQL](https://www.postgresql.org/) - Motor de base de datos SQL
- [Visual Studio Code](https://code.visualstudio.com/) - Editor de Codigo
- [Prettier](https://jwt.io/) - Formateador de Codigo
- [JWT](https://jwt.io/) - Estándar para la creación de tokens de acceso
- [Dbdiagram](https://dbdiagram.io/) - Una herramienta en línea gratuita para dibujar diagramas de entidad-relación escribiendo código.

## Versionado 📌

Usamos [GIT](https://git-scm.com/) para el versionado.

## Autor ✒️

- **Jean Osco** - _Software Engineer_
- [willjean29](https://github.com/willjean29)
- email : willjean29@gmail.com