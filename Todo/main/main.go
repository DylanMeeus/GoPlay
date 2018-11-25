// todo application to test a stack
package main

import (
    "fmt"
    "time"
    "../todo"
    "cloud.google.com/go/spanner"
    database "cloud.google.com/go/spanner/admin/database/apiv1"
    "context"
    "google.golang.org/api/iterator"
)

const connectionString = "projects/todo/instances/todospanner/databases/tododb"

func main() {
    t := todo.Todo{
        Id: 0,
        Description: "Hello World",
    }
    ts := todo.TodoService{
        todo.JsonRepo{
            File: "todostore.json",
        },
    }
    ts.Save(t)
}

func testSpanner() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
    defer cancel()
    _, spannerClient := createClients(ctx, connectionString) 
    fmt.Printf("%v\n", spannerClient)
    err := query(ctx, spannerClient)
    if err != nil {
        fmt.Println(err)
    }
}

func createClients(ctx context.Context, db string) (*database.DatabaseAdminClient, *spanner.Client) {
    adminClient, err := database.NewDatabaseAdminClient(ctx)
    if err != nil {
        panic(err)
    }

    dClient, err := spanner.NewClient(ctx, db)
    if err != nil {
        panic(err)
    }
    return adminClient, dClient
}

func query(ctx context.Context, client *spanner.Client) error {
    stmt := spanner.Statement{SQL: `select * from todo`}
    iter := client.Single().Query(ctx, stmt)
    defer iter.Stop()
    for {
        row, err := iter.Next()
        if err == iterator.Done {
            return nil
        }
        if err != nil {
            return err
        }
        var todoID int64
        if err := row.Columns(&todoID); err != nil {
            return err
        }
        fmt.Printf("%v\n", todoID)
    }
}
