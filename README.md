<p align="center">
  <a href="https://github.com/osramirezdev/scraperPersonas">
    <picture>
      <source height="325" media="(prefers-color-scheme: dark)" srcset="scraperpersona.png">
      <img height="325" alt="Fiber" src="scraperpersona.png">
    </picture>
  </a>
  <br>
</p>
<p align="center">
  <b>ScraperPersonas</b> provee un api para obtener datos de cédulas, haciendo scraping de los siguientes sitios:
  <br>
    * <a href="https://ruc.com.py">Datos de RUC</a><br>
    * <a href="https://servicios.ips.gov.py/consulta_asegurado/comprobacion_de_derecho_externo.php">Datos de IPS</a><br>
    * <a href="https://datos.sfp.gov.py/data/funcionarios">Datos de Funcionarios Publicos</a><br>
    inspired <b>web framework</b> built on top of <a href="https://github.com/valyala/fasthttp">Fasthttp</a>, the <b>fastest</b> HTTP engine for <a href="https://go.dev/doc/">Go</a>. Designed to <b>ease</b> things up for <b>fast</b> development with <b>zero memory allocation</b> and <b>performance</b> in mind.
</p>

It is a simple Go application to scrape the MTESS courses website in the SNPP section.
## 🔧 Requirements

- Go
- Git
- Make
- A mysql database

## 📦 Installation

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
