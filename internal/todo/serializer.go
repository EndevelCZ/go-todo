// Mam dva handlery, grpchandler a handler (httphandler)

// grpc handler mi serializuje ||| func(todo (model)) returns (*pb.Todo) ||| (protobuffer) (metoda todoToPbTodo(todo))
// http handler mi serializuje ||| func(todo (model)) returns ([]byte) ||| (jsonu) (metoda json.Marshal(todo))
// a ted bych potreboval mit neco jako interface Serializer
// a pod nim dva

package todo

import (
	"log"

	"github.com/adamplansky/todo/pb"
	"github.com/golang/protobuf/proto"
)

type TodoPbSerializer interface {
	Serialize(todo *Todo) *pb.Todo
}

type todoPbSerializer struct {
	handler TodoGrpcHandler
}

// func NewTodoPbSerializer(handler TodoGrpcHandler) *todoPbSerializer {
// 	return &todoPbSerializer{
// 		handler,
// 	}
// }

func (s *todoPbSerializer) Serialize(todo *Todo) *pb.Todo {
	return &pb.Todo{
		Id:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}
}

//todo: better to use tags in model??
func TodoToPbTodo(todo *Todo) *pb.Todo {
	return &pb.Todo{
		Id:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}
}

func todoToPbTodo(todo *Todo) *pb.Todo {
	return TodoToPbTodo(todo)
}

func TodoToPbTodo2(todo *Todo) *pb.Todo {
	data, err := proto.Marshal(todo)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	pbtd := &pb.Todo{}
	_ = proto.Unmarshal(data, pbtd)

	return pbtd
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
