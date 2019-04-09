package todo

type Todo struct {
	ID   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Done bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
}
