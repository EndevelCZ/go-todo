package main

import (
	"context"
	"log"
	"net/http"

	datastoredb "github.com/adamplansky/todo/internal/database/datastore"
	"github.com/adamplansky/todo/internal/todo"
	"github.com/gorilla/mux"

	"cloud.google.com/go/datastore"
)

func main() {
	var todoRepository todo.TodoRepository

	projectID := "silent-turbine-233205"
	client, err := newDatastoreClient(projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	todoRepository = datastoredb.NewDataStoreTodoRepository(client)
	todoService := todo.NewTodoService(todoRepository)
	todoHandler := todo.NewTodoHandler(todoService)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", todoHandler.Get).Methods("GET")

	// http.Handle("/", accessControl(middleware.Authenticate(router)))
	// http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
	log.Fatal(http.ListenAndServe(":80", router))
}

// func postgresConnection(database string) *sql.DB {
// 	logrus.Info("Connecting to PostgreSQL DB")
// 	db, err := sql.Open("postgres", database)
// 	if err != nil {
// 		logrus.Fatal(err)
// 		panic(err)
// 	}
// 	return db
// }
// func datastoreConnection(datastore string) {
// 	ctx := context.Background()
// 	client, err := datastore.NewClient(ctx, "project-id")
// 	if err != nil {
// 		// TODO: Handle error.
// 	}
// 	const retries = 3

// }
func newDatastoreClient(projectID string) (*datastore.Client, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// func configureDatastoreDB(projectID string) (BookDatabase, error) {
// 	client, err := newDatastoreClient(projectID)
// 	if err != nil {
// 		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
// 	}
// 	ctx := context.Background()
// 	// Verify that we can communicate and authenticate with the datastore service.
// 	t, err := client.NewTransaction(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
// 	}
// 	if err := t.Rollback(); err != nil {
// 		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
// 	}
// 	return &datastoreDB{
// 		client: client,
// 	}, nil
// }
