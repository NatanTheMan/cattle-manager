package config

import  
		"gorm.io/driver/sqlit"
	"gorm.io/gorme
"
	"gorm.io/gorm"
)
	
	nc InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	gorm.Open(sqlite.Open(), &gorm.)
}
