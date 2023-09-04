package repository

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/osramirezdev/scraperPersonas/platform/database"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm/clause"
)

type BaseRepositoryI[T any] interface {
	CrearOnConflict(entity T, columnName string, columnUpdate []string) (T, error)
	CrearManyConflict(entidad []T, columnsNames []clause.Column, columnUpdate []string) ([]T, error)
	GetByAttribute(entidad T, columnName string, value string) (T, error)
	FindOne(entidad T, columnName string, value string) (T, error)
}

type baseRepository[T interface{}] struct {
	Db *database.DBProvider
}

var BaseRepository baseRepository[interface{}]

func init() {
	datab := &database.DBP
	BaseRepository = baseRepository[interface{}]{Db: datab}
}

func (query *baseRepository[T]) FindOne(entidad T, columnName string, value string) (T, error) {
	pluralizar := pluralize.NewClient()
	entityType := reflect.TypeOf(entidad)
	tableName := strings.ToLower(entityType.Name())
	log.Printf("[%vRepository]...Buscando en tabla %v", tableName, pluralizar.Plural(tableName))
	// obtener nombre de entidad\
	q := fmt.Sprintf(`%s = ?`, columnName)
	fmt.Println("find it", q, value, columnName)
	entity := query.Db.First(&entidad, q, value)
	fmt.Println("find it2", entity, entidad)
	if entity.Error != nil {
		return entidad, entity.Error
	}
	return entidad, nil
}

func (query *baseRepository[T]) CrearOnConflict(entidad T, columnName string, columnUpdate []string) (T, error) {
	// obtener nombre de entidad\
	pluralizar := pluralize.NewClient()
	entityType := reflect.TypeOf(entidad)
	tableName := strings.ToLower(entityType.Name())
	caser := cases.Title(language.Spanish)
	v := reflect.ValueOf(entidad)
	f := v.FieldByName(caser.String(strings.ToLower(columnName)))
	log.Printf("[%vRepository]...Creando en tabla %v", tableName, pluralizar.Plural(tableName))
	q := fmt.Sprintf(`%s = ?`, columnName)
	if entity := query.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: columnName}},
		DoUpdates: clause.AssignmentColumns(columnUpdate),
	}).Create(&entidad).First(&entidad, q, f.String()); entity.Error != nil {
		return entidad, entity.Error
	}
	return entidad, nil
}

func (query *baseRepository[T]) CrearManyConflict(entidad []T, columnsNames []clause.Column, columnUpdate []string) ([]T, error) {
	// obtener nombre de entidad\
	pluralizar := pluralize.NewClient()
	entityType := reflect.TypeOf(entidad[0])
	tableName := strings.ToLower(entityType.Name())
	log.Printf("[%vRepository]...Creando en tabla %v", tableName, pluralizar.Plural(tableName))
	if entity := query.Db.Clauses(clause.OnConflict{
		Columns:   columnsNames,
		DoUpdates: clause.AssignmentColumns(columnUpdate),
	}).Create(&entidad).First(&entidad); entity.Error != nil {
		return entidad, entity.Error
	}
	return entidad, nil
}

func (query *baseRepository[T]) GetByAttribute(entidad T, columnName string, value string) (T, error) {
	pluralizar := pluralize.NewClient()
	entityType := reflect.TypeOf(entidad)
	tableName := strings.ToLower(entityType.Name())
	log.Printf("[%vRepository]...Buscando en tabla %v", tableName, pluralizar.Plural(tableName))
	// obtener nombre de entidad\
	q := fmt.Sprintf(`%v = ?`, columnName)
	entity := query.Db.Model(&entidad).Where(q, value).First(&entidad)
	if entity.Error != nil {
		return entidad, entity.Error
	}
	return entidad, nil
}
