package main

import (
    "task-5-pbi-btpns-mdafir/database"
    "task-5-pbi-btpns-mdafir/router"
)

func main() {
    db, err := database.ConnectDB()
    if err != nil {
        panic("Failed to connect to the database")
    }

    database.Migrate(db)
    r := router.SetupRouter()
    r.Run(":8080")
}