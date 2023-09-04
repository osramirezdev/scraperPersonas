<p align="center">
  <a href="https://github.com/osramirezdev/scraperPersonas">
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

## ⚙️ Installation

Para correrlo localmente, tener intalado Go (<a target="_blank" style="color: blue" href="https://go.dev/dl/">download</a>). Version `1.17` o superior.

```bash
go mod download
```


## 🚀 Usage

```bash
cp .env.example .env
```

Ingresar las credenciales en el archivo `.env`, y setear `STAGE_STATUS` con `prod`.

Buildear con docker
```bash
docker build -t fiber .
docker run -d -p 5000:5000 --name fiber fiber
```

Tambien se puede ejecutar via docker-compose [docker-compose](https://docs.docker.com/compose/install/)
```bash
sudo docker-compose build
sudo docker-compose up -d
```
## 👨‍💻 Author

#### [Oscar Ramirez](https://yocreativo.com)

## ✍️ License

[MIT](https://choosealicense.com/licenses/mit/)
