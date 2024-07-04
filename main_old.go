package main

import (
    "database/sql"
    "fmt"
    "time"
    "os"
    "github.com/joho/godotenv"
    _ "github.com/databricks/databricks-sql-go"
    // dbsql "github.com/databricks/databricks-sql-go"
    "github.com/databricks/databricks-sql-go/auth/oauth/u2m"
    // "github.com/databricks/databricks-sdk-go"
    // "github.com/databricks/databricks-sdk-go/service/jobs"
)

func main(){
    fmt.Println("Testing")

    godotenv.Load()

    authenticator,err := u2m.NewAuthenticator(os.Getenv("SERVER_HOSTNAME"), 1 * time.Minute)

    if err != nil {
        panic(fmt.Errorf("Error initializing authenticator in main.go: %s", err))
    }
    fmt.Println("authenticator: ", authenticator)

    connector, err := dbsql.NewConnector(
        dbsql.WithServerHostname(os.Getenv("SERVER_HOSTNAME")),
        dbsql.WithHTTPPath(os.Getenv("HTTP_PATH")),
        dbsql.WithPort(os.Getenv("PORT")),
        dbsql.WithAuthenticator(authenticator),
    )

    if err != nil {
        panic(fmt.Errorf("Error getting connector: %s", err))
    }

    fmt.Println("connector: ", connector)
}
