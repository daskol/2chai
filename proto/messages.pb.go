// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Post
	Thread
	Board
	BoardList
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Post struct {
	// Идентификатор нить, к которому принадлежит пост.
	ThreadId uint64 `protobuf:"varint,1,opt,name=thread_id,json=threadId" json:"thread_id,omitempty"`
	// Глобальный идентификатор поста.
	CommentId uint64 `protobuf:"varint,2,opt,name=comment_id,json=commentId" json:"comment_id,omitempty"`
	// Относительный идентификатор поста в данной ните.
	Ordinal uint64 `protobuf:"varint,3,opt,name=ordinal" json:"ordinal,omitempty"`
	// Тема поста.
	Subject string `protobuf:"bytes,4,opt,name=subject" json:"subject,omitempty"`
	// Комментарий.
	Comment string `protobuf:"bytes,5,opt,name=comment" json:"comment,omitempty"`
	// Время, когда был оставлен комментарий.
	Timestamp int64 `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	// Вектор идентификаторов постов, на который отвечает данный коментарий.
	ReplyTo []uint64 `protobuf:"varint,7,rep,packed,name=reply_to,json=replyTo" json:"reply_to,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto1.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Post) GetThreadId() uint64 {
	if m != nil {
		return m.ThreadId
	}
	return 0
}

func (m *Post) GetCommentId() uint64 {
	if m != nil {
		return m.CommentId
	}
	return 0
}

func (m *Post) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *Post) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Post) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Post) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Post) GetReplyTo() []uint64 {
	if m != nil {
		return m.ReplyTo
	}
	return nil
}

type Thread struct {
	// Идентификатор доски, на котором стартовала нить.
	BoardId uint64 `protobuf:"varint,1,opt,name=board_id,json=boardId" json:"board_id,omitempty"`
	// Идентификатор нити(первого поста).
	ThreadId uint64 `protobuf:"varint,2,opt,name=thread_id,json=threadId" json:"thread_id,omitempty"`
	// Ответы, принадлежащией этой ните.
	Posts []*Post `protobuf:"bytes,3,rep,name=posts" json:"posts,omitempty"`
}

func (m *Thread) Reset()                    { *m = Thread{} }
func (m *Thread) String() string            { return proto1.CompactTextString(m) }
func (*Thread) ProtoMessage()               {}
func (*Thread) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Thread) GetBoardId() uint64 {
	if m != nil {
		return m.BoardId
	}
	return 0
}

func (m *Thread) GetThreadId() uint64 {
	if m != nil {
		return m.ThreadId
	}
	return 0
}

func (m *Thread) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type Board struct {
	// Локальный числовой идентификатор доски.
	BoardId uint64 `protobuf:"varint,1,opt,name=board_id,json=boardId" json:"board_id,omitempty"`
	// Название доски.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Board) Reset()                    { *m = Board{} }
func (m *Board) String() string            { return proto1.CompactTextString(m) }
func (*Board) ProtoMessage()               {}
func (*Board) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Board) GetBoardId() uint64 {
	if m != nil {
		return m.BoardId
	}
	return 0
}

func (m *Board) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BoardList struct {
	// Массив c описанием досок.
	Boards []*Board `protobuf:"bytes,1,rep,name=boards" json:"boards,omitempty"`
}

func (m *BoardList) Reset()                    { *m = BoardList{} }
func (m *BoardList) String() string            { return proto1.CompactTextString(m) }
func (*BoardList) ProtoMessage()               {}
func (*BoardList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BoardList) GetBoards() []*Board {
	if m != nil {
		return m.Boards
	}
	return nil
}

func init() {
	proto1.RegisterType((*Post)(nil), "Post")
	proto1.RegisterType((*Thread)(nil), "Thread")
	proto1.RegisterType((*Board)(nil), "Board")
	proto1.RegisterType((*BoardList)(nil), "BoardList")
}

func init() { proto1.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcb, 0x6a, 0xf3, 0x30,
	0x10, 0x85, 0x51, 0x7c, 0x49, 0x34, 0x3f, 0xfc, 0x0b, 0xad, 0x54, 0xd2, 0xb4, 0xc6, 0x2b, 0x43,
	0xc1, 0x81, 0x14, 0xfa, 0x00, 0xd9, 0x05, 0xba, 0x28, 0x22, 0xab, 0x42, 0x09, 0xb2, 0x25, 0x12,
	0xb5, 0x51, 0x64, 0x2c, 0x65, 0xd1, 0x47, 0xec, 0x5b, 0x15, 0x8d, 0x1d, 0x7a, 0x59, 0x74, 0x25,
	0x9d, 0xf3, 0x8d, 0xce, 0x68, 0x24, 0xf8, 0x6f, 0xb5, 0xf7, 0x72, 0xaf, 0x7d, 0xdd, 0xf5, 0x2e,
	0xb8, 0xf2, 0x83, 0x40, 0xfa, 0xe4, 0x7c, 0x60, 0x73, 0xa0, 0xe1, 0xd0, 0x6b, 0xa9, 0x76, 0x46,
	0x71, 0x52, 0x90, 0x2a, 0x15, 0xb3, 0xc1, 0xd8, 0x28, 0xb6, 0x00, 0x68, 0x9d, 0xb5, 0xfa, 0x14,
	0x22, 0x9d, 0x20, 0xa5, 0xa3, 0xb3, 0x51, 0x8c, 0xc3, 0xd4, 0xf5, 0xca, 0x9c, 0xe4, 0x91, 0x27,
	0xc8, 0x2e, 0x32, 0x12, 0x7f, 0x6e, 0x5e, 0x75, 0x1b, 0x78, 0x5a, 0x90, 0x8a, 0x8a, 0x8b, 0x8c,
	0x64, 0x0c, 0xe0, 0xd9, 0x40, 0x46, 0xc9, 0xae, 0x81, 0x06, 0x63, 0xb5, 0x0f, 0xd2, 0x76, 0x3c,
	0x2f, 0x48, 0x95, 0x88, 0x2f, 0x83, 0x5d, 0xc1, 0xac, 0xd7, 0xdd, 0xf1, 0x7d, 0x17, 0x1c, 0x9f,
	0x16, 0x49, 0x6c, 0x86, 0x7a, 0xeb, 0xca, 0x17, 0xc8, 0xb7, 0x78, 0xe3, 0x58, 0xd4, 0x38, 0xd9,
	0x7f, 0x9b, 0x65, 0x8a, 0x7a, 0xa3, 0x7e, 0xce, 0x39, 0xf9, 0x35, 0xe7, 0x1c, 0xb2, 0xce, 0xf9,
	0xe0, 0x79, 0x52, 0x24, 0xd5, 0xbf, 0x55, 0x56, 0xc7, 0xa7, 0x11, 0x83, 0x57, 0x3e, 0x40, 0xb6,
	0x8e, 0x21, 0x7f, 0xa5, 0x33, 0x48, 0x4f, 0xd2, 0x6a, 0x0c, 0xa6, 0x02, 0xf7, 0xe5, 0x1d, 0x50,
	0x3c, 0xf7, 0x68, 0x7c, 0x60, 0x37, 0x90, 0x63, 0xad, 0xe7, 0x04, 0x5b, 0xe4, 0x35, 0x32, 0x31,
	0xba, 0xeb, 0xdb, 0xe7, 0xc5, 0xde, 0x84, 0xc3, 0xb9, 0xa9, 0x5b, 0x67, 0x97, 0x4a, 0xfa, 0x37,
	0x77, 0x5c, 0xae, 0xda, 0x83, 0x34, 0x4b, 0xfc, 0xb0, 0x26, 0xc7, 0xe5, 0xfe, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x75, 0x91, 0x9d, 0x41, 0xc9, 0x01, 0x00, 0x00,
}
