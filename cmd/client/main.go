package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adamplansky/todo/pb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	port = "5000"
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(fmt.Sprintf(":%s", port), grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect to backend")
	}
	client := pb.NewTodosClient(conn)
	switch cmd := flag.Arg(0); cmd {
	case "add":
		err = add(context.Background(), client, strings.Join(flag.Args()[1:], " "))
	case "list":
		err = list(context.Background(), client)
	case "update":
		id, _ := strconv.ParseInt(flag.Arg(1), 10, 64)
		err = update(context.Background(), client, id)
	case "delete":
		id, _ := strconv.ParseInt(flag.Arg(1), 10, 64)
		err = delete(context.Background(), client, id)
	default:
		logrus.Fatalf("unknown command")
	}
	if err != nil {
		logrus.Fatalln(os.Stderr, err)
		os.Exit(1)
	}
}
func add(ctx context.Context, client pb.TodosClient, text string) error {
	t := &pb.Text{
		Text: text,
	}
	td, err := client.AddTodo(ctx, t)
	if err != nil {
		return err
	}
	logrus.Printf("%#v", td)
	return nil
}
func list(ctx context.Context, client pb.TodosClient) error {
	t, err := client.ListTodos(ctx, &pb.Void{})
	if err != nil {
		return err
	}
	for _, t := range t.Todos {
		fmt.Printf("%d %s ", t.Id, t.Text)
		if t.Done {
			fmt.Printf("✅")
		} else {
			fmt.Printf("❌")
		}
		fmt.Printf("\n")
	}
	return nil
}
func update(ctx context.Context, client pb.TodosClient, id int64) error {
	tid := &pb.Integer{
		Id: id,
	}
	todo, err := client.CheckTodo(ctx, tid)
	if err != nil {
		return err
	}
	logrus.Printf("%#v", todo)
	return nil
}
func delete(ctx context.Context, client pb.TodosClient, id int64) error {
	tid := &pb.Integer{
		Id: id,
	}
	todo, err := client.DeleteTodo(ctx, tid)
	if err != nil {
		return err
	}
	logrus.Printf("%#v", todo)
	return nil
}
