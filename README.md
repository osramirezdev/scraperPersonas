![Scraper Personas](scraperpersona.jpg)

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
