package funcs

import (
	"github.com/nielsdingsbums/anas_epistulae/db"
	"github.com/nielsdingsbums/anas_epistulae/structs"
)

func GetDB() []structs.Message {
	return db.Messages
}
