package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/osramirezdev/scraperPersonas/app/entities"
	"github.com/osramirezdev/scraperPersonas/app/repository"
	"github.com/osramirezdev/scraperPersonas/pkg/utils"
)

type FileServiceI interface {
	SaveToDB(file multipart.File, filename string) (entities.Archivo, error)
}

type fileConstructor struct {
	ArchivoRepo repository.ArchivoRepositoryI
}

var FileService fileConstructor

func init() {
	FileService = fileConstructor{
		ArchivoRepo: repository.ArchivoRepository,
	}
}

func (repo *fileConstructor) SaveToDB(fileForm *multipart.FileHeader, filename string) (entities.Archivo, error) {
	file, errFile := fileForm.Open()
	if errFile != nil {
		fmt.Println("error abrir archivo", errFile)
		return entities.Archivo{}, errFile
	}

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return entities.Archivo{}, err
	}
	fmt.Println(header)
	// antes guardamos en directorio
	nombreArchivoFormateado := utils.ObtenerHora("2006_01_02_15_04_05") + "_" + filename
	fileToWrite, errW := os.Create("./archivos/" + nombreArchivoFormateado)
	if errW != nil {
		return entities.Archivo{}, errW
	}
	_, errWF := io.Copy(fileToWrite, file)
	if errWF != nil {
		return entities.Archivo{}, errW
	}

	archivo := entities.Archivo{
		Nombre:        nombreArchivoFormateado,
		FullDireccion: "./archivos/" + nombreArchivoFormateado,
	}
	newArchivo, errorF := repo.ArchivoRepo.Upsert(archivo)
	if errorF != nil {
		return entities.Archivo{}, errorF
	}
	defer fileToWrite.Close()
	return newArchivo, nil
}
