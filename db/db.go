package db

import (
	"github.com/nielsdingsbums/anas_epistulae/structs"
	"time"
)

var theTime time.Time = time.Now()

var Messages []structs.Message = []structs.Message{{0, "Hallo", theTime, "lol"}}
var Authors []structs.Author