<table style="width:100%" align="center" border="0">
  <tr align="center">
    <td><img src=".github/assets/cubes.png" alt="module" width="300"></td>
    <td><h1>🧩 Go Modules API 📦</h1></td>
  </tr>
</table>

<p align="center">
  <strong>An open-source API to manage hub client modules built with Go.</strong>
</p>

<p align="center">
  <img src="https://wakatime.com/badge/user/e61842d0-c588-4586-96a3-f0448a434be4/project/c3a55bc2-b58f-455a-a889-0ee8c1e9ff12.svg" alt="waka" />
  <img src="https://img.shields.io/github/license/gabrielmaialva33/go-modules-api?color=00b8d3?style=flat&logo=appveyor" alt="License" />
  <img src="https://img.shields.io/github/languages/top/gabrielmaialva33/go-modules-api?style=flat&logo=appveyor" alt="GitHub top language" >
  <img src="https://img.shields.io/github/languages/count/gabrielmaialva33/go-modules-api?style=flat&logo=appveyor" alt="GitHub language count" >
  <img src="https://img.shields.io/github/repo-size/gabrielmaialva33/go-modules-api?style=flat&logo=appveyor" alt="Repository size" >
  <a href="https://github.com/gabrielmaialva33/go-modules-api/commits/master">
    <img src="https://img.shields.io/github/last-commit/gabrielmaialva33/go-modules-api?style=flat&logo=appveyor" alt="GitHub last commit" >
    <img src="https://img.shields.io/badge/made%20by-Maia-15c3d6?style=flat&logo=appveyor" alt="Maia" >  
  </a>
</p>

<br>

<p align="center">
  <a href="#bookmark-about">About</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#computer-technologies">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#wrench-tools">Tools</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#package-installation">Installation</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#memo-license">License</a>
</p>

<br>

## :bookmark: About

**GO MODULES API** is a REST API built with **Go**, using **Fiber** as the web framework and **GORM** for database
interactions.  
It provides a modular structure to manage **hub clients**, **modules**, **roles**, and **permissions** efficiently.

<br>

## :computer: Technologies

- **[Go](https://go.dev/)**
- **[Fiber](https://gofiber.io/)**
- **[GORM](https://gorm.io/)**
- **[PostgreSQL](https://www.postgresql.org/)**
- **[Docker](https://www.docker.com/)**
- **[Cobra](https://github.com/spf13/cobra)** (CLI management)
- **[Zap](https://github.com/uber-go/zap)** (Logging)
- **[OpenAPI](https://swagger.io/specification/)** (API documentation)

<br>

## :chart_with_upwards_trend: Database Diagram

<p align="center">
  <img src=".github/assets/modules_graphml.svg" alt="Database Diagram" width="800">
</p>

<br>

## :wrench: Tools

- **[VSCode](https://code.visualstudio.com/)**
- **[GoLand](https://www.jetbrains.com/go/)**

<br>

## :package: Installation

### :gear: **Prerequisites**

Ensure you have the following installed:

- **[Go](https://go.dev/dl/)**
- **[Git](https://git-scm.com/)**
- **[Docker](https://www.docker.com/)**
- **[Docker Compose](https://docs.docker.com/compose/)**

<br>

### :octocat: **Cloning the repository**

```sh
git clone https://github.com/gabrielmaialva33/go-modules-api.git
cd go-modules-api
```

<br>

### :whale: **Running the application with Docker**

```sh
docker-compose up --build
```

The application will be available at `http://localhost:3000`.

<br>

### :computer: **Running the application locally**

```sh
# Copy the .env.example file to .env
cp .env.example .env

# Install dependencies
go mod tidy

# Start the application
go run main.go
```

The application will be available at `http://localhost:3000`.

<br>

## :rocket: **API Documentation**

After running the application, the API documentation (redoc) will be available at:

- 📌 Redoc UI: `http://localhost:3000/docs`
- 📌 OpenAPI File: `http://localhost:3000/openapi.yaml`

<br>

## :memo: License

This project is under the **MIT** license. [MIT](./LICENSE) ❤️

<br>

## :rocket: **Contributors**

| [![Maia](https://avatars.githubusercontent.com/u/26732067?size=100)](https://github.com/gabrielmaialva33) |
|-----------------------------------------------------------------------------------------------------------|
| [Maia](https://github.com/gabrielmaialva33)                                                               |

Made with ❤️ by Maia 👋🏽 [Get in touch!](https://t.me/mrootx)

## :star:

Liked? Leave a little star to help the project ⭐

<br/>
<br/>

<p align="center"><img src="https://raw.githubusercontent.com/gabrielmaialva33/gabrielmaialva33/master/assets/gray0_ctp_on_line.svg?sanitize=true" /></p>
<p align="center">&copy; 2017-present <a href="https://github.com/gabrielmaialva33/" target="_blank">Maia</a>



