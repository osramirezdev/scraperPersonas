<p align="center">
  <a href="http://yocreativo.com:5100/swagger">
    <picture>
      <source height="350" media="(prefers-color-scheme: dark)" srcset="scraperlogo.png">
      <img height="350" alt="Scraper Perosnas" src="scraperlogo.png">
    </picture>
  </a>
  <br>
</p>
<p align="left">
  <b>ScraperPersonas</b> provee un RESTfulapi para obtener datos de cédulas, haciendo scraping de los siguientes sitios:
  <br>
  <ol>
    <li><a target="_blank" href="https://ruc.com.py">Datos de RUC</a></li>
    <li><a target="_blank" href="https://servicios.ips.gov.py/consulta_asegurado/comprobacion_de_derecho_externo.php">Datos de IPS</a></li>
    <li><a target="_blank" href="https://datos.sfp.gov.py/data/funcionarios">Datos de Funcionarios Publicos</a></li>
  </ol>
</p>

## [API Demo](http://yocreativo.com:5100/swagger)
<br>
## ⚙️ Installation

**Para correrlo localmente**

  - Tener intalado Go (<a target="_blank" style="color: blue" href="https://go.dev/dl/">download</a>). Version `1.17` o superior.
  - Tener intalado PostgreSQL (<a target="_blank" style="color: blue" href="https://www.postgresql.org/">download</a>). Version `13` o superior y tener creada la base de datos configurada en el .env file.

```bash
go mod download # download modules to local cache
go mod tidy # add missing and remove unused modules
go run main.go
```


## 🚀 Usage

```bash
cp .env.example .env
```

Ingresar las credenciales en el archivo `.env`, y setear `STAGE_STATUS` con `prod`.

Buildear con [docker-compose](https://docs.docker.com/compose/install/)
```bash
sudo docker-compose build
sudo docker-compose up -d
```
## 👨‍💻 Author

#### [Oscar Ramirez](https://yocreativo.com)

## ✍️ License

[MIT](https://choosealicense.com/licenses/mit/)
