package main

import (
    "fmt"
    "github.com/joho/godotenv"
    "database/sql"
    "os"
    dbsql "github.com/databricks/databricks-sql-go"
)

func main(){
    fmt.Println("Testing")

    godotenv.Load()

    connector, err := dbsql.NewConnector(
        dbsql.WithAccessToken(os.Getenv("DATABRICKS_TOKEN")),
        dbsql.WithServerHostname(os.Getenv("SERVER_HOSTNAME")),
        dbsql.WithPort(443),
        dbsql.WithHTTPPath(os.Getenv("HTTP_PATH")),
    )
    if err != nil {
        panic(fmt.Errorf("Error gertting connector: %s", err))
    }

    db := sql.OpenDB(connector)
    defer db.Close()

    if err := db.Ping(); err != nil {
        panic(err)
    }

}
