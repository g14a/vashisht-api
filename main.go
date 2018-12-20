package main

import (
	"fmt"

	"github.com/vashisht-api/db"
	"github.com/vashisht-api/models"
)

var (
	DATABASE   = "vashisht"
	COLLECTION = "events"
)

func main() {

	e3 := &models.Event{
		EventName: "Technical Talk",
		EventId:   "3",
		Fee:       100,
		TeamSize:  0,
		Category:  "Talk",
		Day:       3,
		StartTime: 1000,
		EndTime:   1200,
	}

	dbinstance := db.GetDbInstance()

	e3 = &models.Event{
		EventName: "Technical Talk",
		EventId:   "3",
		Fee:       100,
		TeamSize:  0,
		Category:  "updated",
		Day:       3,
		StartTime: 1000,
		EndTime:   1200,
	}

	err := models.UpdateEvent(*e3, dbinstance)

	if err != nil {
		fmt.Println(err)
	}

}
