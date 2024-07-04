package main

import (
    "fmt"
    "context"
    "github.com/joho/godotenv"
    "database/sql"
    "os"
    dbsql "github.com/databricks/databricks-sql-go"
    "github.com/databricks/databricks-sdk-go"
    "github.com/databricks/databricks-sdk-go/service/jobs"
)

func main(){
    fmt.Println("Testing")

    godotenv.Load()

    connector, err := dbsql.NewConnector(
        dbsql.WithServerHostname(os.Getenv("DATABRICKS_SERVER_HOSTNAME")),
        dbsql.WithHTTPPath(os.Getenv("DATABRICKS_HTTP_PATH")),
        dbsql.WithPort(443),
        dbsql.WithAccessToken(os.Getenv("DATABRICKS_TOKEN")),
    )

    if err != nil {
        panic(fmt.Errorf("Error gertting connector: %s", err))
    }
    fmt.Println("connector: ", connector)

    db := sql.OpenDB(connector)
    defer db.Close()
    //
    // if err := db.Ping(); err != nil {
    //     panic(err)
    // }


    w := databricks.Must(databricks.NewWorkspaceClient())
    ctx := context.Background()

    nt := jobs.NotebookTask{
        NotebookPath: os.Getenv("TEST_NOTEBOOK_PATH"),
    }

    jobToRun, err := w.Jobs.Create(ctx, jobs.CreateJob{
        Name: askFor("Some short name for the job:"),
        Tasks: []jobs.JobTaskSettings{
            {
                Description:       askFor("Some short description for the job:"),
                TaskKey:           askFor("Some key to apply to the job's tasks:"),
                ExistingClusterId: askFor("ID of the existing cluster in the workspace to run the job on:"),
                NotebookTask:      &nt,
            },
        },
    })

    if err != nil {
        panic(err)
    }

    fmt.Printf("Now attempting to run the job at %s/#job/%d, please wait...\n",
    w.Config.Host,
    jobToRun.JobId,
)

runningJob, err := w.Jobs.RunNow(ctx, jobs.RunNow{
    JobId: jobToRun.JobId,
})

if err != nil {
    panic(err)
}

jobRun, err := runningJob.Get()

if err != nil {
    panic(err)
}

fmt.Printf("View the job run results at %s/#job/%d/run/%d\n",
w.Config.Host,
jobRun.JobId,
jobRun.RunId,
  )

  // Output:
  //
  // Now attempting to run the job at <workspace-host>/#job/<job-id>, please wait...
  // View the job run results at <workspace-host>/#job/<job-id>/run/<run-id>
}
}
