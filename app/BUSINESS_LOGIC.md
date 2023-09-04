# ./app

**Carpeta con lo lógica de negocios**. Este directorio no se precocupa acerca que que database utilicemos, o que driver implementamos, o manejar libreria de terceros.

- `./app/controllers` aquí se gestiona las solicitudes entrantes HTTP, procesamos los DTOS y dirigimos la logica hacia un servicio (usado en routes).
- `./app/dtos` aquí se almacenan los objetos para pasar datos entre las distintas capas de la aplicación.
- `./app/entities` aquí describimos las entidades del proyectos utilizado en nuestra base de datos.
- `./app/scrapers` aquí mantenemos la lógica de scraping, deberia ir a otro folder.
- `./app/services` aquí mantenemos la lógica de negocios para los controladores, ocultamos la implementación y encapsulamos las acciones con los distintos repositorios.
- `./app/repository` carpeta para realizar las operaciones CRUD en la base de datos
