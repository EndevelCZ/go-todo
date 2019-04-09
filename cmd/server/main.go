package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	datastoredb "github.com/adamplansky/todo/internal/database/datastore"
	"github.com/adamplansky/todo/internal/todo"
	"github.com/adamplansky/todo/pb"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"cloud.google.com/go/datastore"
)

func main() {

	grpcPtr := flag.Bool("grpc", false, "grpc server")
	httpPtr := flag.Bool("http", false, "http server")
	port := flag.String("port", "5000", "port")
	flag.Parse()
	datastoreURL := os.Getenv("DATASTORE_EMULATOR_HOST")
	logrus.Println("DATASTORE_EMULATOR_HOST: ", datastoreURL)
	logrus.Println("listening on PORT: ", *port)
	var todoRepository todo.TodoRepository

	projectID := "silent-turbine-233205"
	client, err := newDatastoreClient(projectID)
	if err != nil {

		logrus.Fatalf("unable to connect to database %v", err)
	}
	defer client.Close()
	todoRepository = datastoredb.NewDataStoreTodoRepository(client)
	todoService := todo.NewTodoService(todoRepository)
	switch {
	case *grpcPtr:
		todoGrpcHandler := todo.NewTodoGrpcHandler(todoService)

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
		if err != nil {
			logrus.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterTodosServer(s, todoGrpcHandler)
		if err := s.Serve(lis); err != nil {
			logrus.Fatalf("failed to serve: %v", err)
		}
	case *httpPtr:
		todoHandler := todo.NewTodoHandler(todoService)

		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
		router.HandleFunc("/todos", todoHandler.Get).Methods("GET")

		// http.Handle("/", accessControl(middleware.Authenticate(router)))
		// http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
		logrus.Println(fmt.Sprintf(":%s", *port))
		logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), router))
	default:
		fmt.Println("invalid flag")
	}

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
	logrus.Println("logrus: newDatastoreClient", projectID)
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	logrus.Printf("logrus: newDatastoreClient %v", client)
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
