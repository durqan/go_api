package main

import (
	"log"
	"test/database"

	_ "test/docs"

	_ "github.com/lib/pq"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}

	db := database.GetDB()
	err = database.AutoMigrateAll(db)
	if err != nil {
		log.Fatal("Ошибка миграций:", err)
	}
	log.Println("✅ Миграции выполнены успешно")

	startRouter()
}
