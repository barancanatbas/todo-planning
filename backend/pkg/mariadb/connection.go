package mariadb

import (
	"fmt"
	"task-manager/internal/model/entity"

	"github.com/rs/zerolog/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMariaDbConnection(config MariadbConfig) *gorm.DB {
	var charSet string

	if config.CharSet == "" {
		charSet = ""
	} else {
		charSet = fmt.Sprintf("&charset=%s", config.CharSet)
	}

	connectionString := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local%s`, config.User, config.Password, config.Host, config.Port, config.DbName, charSet)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		log.Fatal().Err(err).Msg("mariadb connection failed")
	}

	var sum int
	err = db.Raw("select 1").Scan(&sum).Error
	if err != nil {
		log.Fatal().Err(err).Msg("mariadb connection failed")
	}

	if config.Debug == true {
		return db.Debug()
	}

	migrateTables(db)

	return db
}

func migrateTables(db *gorm.DB) error {
	err := db.Migrator().DropTable(&entity.Developer{}, &entity.Task{}, &entity.Week{})
	if err != nil {
		log.Fatal().Err(err)
	}

	err = db.AutoMigrate(&entity.Developer{}, &entity.Task{}, &entity.Week{})
	if err != nil {
		log.Fatal().Err(err)
	}

	developers := []entity.Developer{
		{Name: "DEV1", Capacity: 1 * 45},
		{Name: "DEV2", Capacity: 2 * 45},
		{Name: "DEV3", Capacity: 3 * 45},
		{Name: "DEV4", Capacity: 4 * 45},
		{Name: "DEV5", Capacity: 5 * 45},
	}

	err = db.Create(&developers).Error
	if err != nil {
		log.Fatal().Err(err)
	}

	return err
}
