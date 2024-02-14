
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
- [Analisis de Código](#analisis-de-código-)
- [Testing](#testing)
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

## Testing

_Dentro del proyecto los archivos releaciones a testing tienen la siguiente nomeclatura_

```
*_test.go
```

_para ejecutar los test dentro de la aplicacion y que se muestre el resumen de coverage en consola ejecutar el comando:_
 
```
go test -coverprofile=coverage.out ./...  
```
_Ejecución en consola_

![Reporte coverage consola](/docs/testing/coverage.png)

_Una vez generado el archivo coverage.out, puedes usar el comando **go tool cover** para analizarlo y generar un informe HTML más detallado._
 
```
go tool cover -html=coverage.out -o coverage.html  
```

_Ejecución en navegador_

![Reporte coverage navegador](/docs/testing/reporte.png)
## Analisis de Código 🔩

_**Pre requisitos**_

_En la raiz del proyecto se tiene el archivo *sonar-project.properties* el cual tiene las propiedades necesarias para ejecutarlo sobre un SonarQube_

```
Sonaqube >= 9.X
```

_**SonarQube Properties**_

```
# Host
sonar.host.url=http://localhost:9000
# Nombre del proyecto
sonar.projectKey=YOUR_PROJECT_KEY
sonar.projectName=YOUR_PROJECT_NAME
sonar.token=YOUR_TOKEN
# Versión del proyecto
sonar.projectVersion=YOUR_PROJECT_VERSION

# Rutas a analizar
sonar.sources=.

sonar.sourceEncoding=UTF-8
sonar.go.coverage.reportPaths=coverage.out
```

_Las pruebas fueron realizas sobre *SonarQube 9.8* para ejecutar el analisis de codigo ejecutar el comando para la instancia local:_

```
sonar-scanner
```

_Para correr el comando anteior es necesario tener instalado el CLI de Sonar Scanner_

_Reporte de SonarQube_

![SonarQube](/docs/sonar/sonarqube.png)

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