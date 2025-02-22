package schemas

import 	"gorm.io/gorm"

type Cattle struct {
	gorm.Model
	Name    string
	Earring int16
}
