package services_test

import (
	"context"
	"fmt"
	"key-value-system/db"
	"key-value-system/helper"
	"key-value-system/models"
	"key-value-system/requests"
	"key-value-system/services"
	"reflect"
	"testing"
)

func TestStoreHead(t *testing.T) {
	helper.ClearDB()

	type args struct {
		request requests.CreateHeadRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				request: requests.CreateHeadRequest{
					Key: "head1",
				},
			},
			wantErr: false,
		},
		{
			name: "duplicated",
			args: args{
				request: requests.CreateHeadRequest{
					Key: "head1",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := services.StoreHead(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("StoreHead() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	head, _ := services.GetHead("head1")
	if head == nil {
		t.Errorf("Cannot get Head")
	}
}

func TestStoreNode(t *testing.T) {
	helper.ClearDB()

	ctx, _ := context.WithCancel(context.Background())

	// create a head
	headRequest := requests.CreateHeadRequest{
		Key: "head1",
	}
	services.StoreHead(headRequest)

	type args struct {
		request requests.CreateNodeRequest
		ctx     context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				request: requests.CreateNodeRequest{
					Key: "node1",
				},
				ctx: ctx,
			},
			wantErr: false,
		},
		{
			name: "duplicated",
			args: args{
				request: requests.CreateNodeRequest{
					Key: "node1",
				},
				ctx: ctx,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := services.StoreNode(tt.args.request, tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("StoreNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	node, _ := services.GetNode("node1")
	if node == nil {
		t.Errorf("Cannot get Node")
	}
}

func TestGetHead(t *testing.T) {
	helper.ClearDB()

	// create a head
	request := requests.CreateHeadRequest{
		Key: "head1",
	}
	services.StoreHead(request)

	type args struct {
		key string
	}
	tests := []struct {
		name      string
		args      args
		returnNil bool
		wantErr   bool
	}{
		{
			name: "normal",
			args: args{
				key: "head1",
			},
			returnNil: false,
			wantErr:   false,
		},
		{
			name: "not found",
			args: args{
				key: "head2",
			},
			returnNil: true,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := services.GetHead(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.returnNil && got != nil {
				t.Errorf("GetHead() should return nil but did not")
			}
		})
	}
}

func TestGetNode(t *testing.T) {
	helper.ClearDB()

	// create a head
	request := requests.CreateHeadRequest{
		Key: "head1",
	}
	services.StoreHead(request)

	ctx, _ := context.WithCancel(context.Background())
	nodeRequest := requests.CreateNodeRequest{
		Key:   "node1",
		Value: "qwe123",
		Prev:  "head1",
	}
	services.StoreNode(nodeRequest, ctx)

	type args struct {
		key string
	}
	tests := []struct {
		name      string
		args      args
		returnNil bool
		wantErr   bool
	}{
		{
			name: "normal",
			args: args{
				key: "node1",
			},
			returnNil: false,
			wantErr:   false,
		},
		{
			name: "not found",
			args: args{
				key: "node2",
			},
			returnNil: true,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := services.GetNode(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.returnNil && got != nil {
				t.Errorf("GetNode() should return nil but did not")
			}
		})
	}
}

func TestRemoveHead(t *testing.T) {
	helper.ClearDB()

	// create a head and a node
	headRequest := requests.CreateHeadRequest{
		Key: "head1",
	}
	services.StoreHead(headRequest)
	// create a head without next
	headRequest = requests.CreateHeadRequest{
		Key: "head2",
	}
	services.StoreHead(headRequest)

	ctx, _ := context.WithCancel(context.Background())
	nodeRequest := requests.CreateNodeRequest{
		Key:   "node1",
		Value: "qwe123",
		Prev:  "head1",
	}
	services.StoreNode(nodeRequest, ctx)

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				key: "head1",
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				key: "aaa",
			},
			wantErr: true,
		},
		{
			name: "no next's head",
			args: args{
				key: "head2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := services.RemoveHead(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("RemoveHead() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	head, _ := services.GetHead("head1")
	if head == nil {
		t.Errorf("Cannot get Head1")
	}
	head, _ = services.GetHead("head2")
	if head == nil {
		t.Errorf("Cannot get Head2")
	}

	node, _ := services.GetNode("node1")
	if node != nil {
		t.Errorf("Get Node unexpectedly")
	}
}

func TestRemoveNode(t *testing.T) {
	helper.ClearDB()

	// create a head and a node
	headRequest := requests.CreateHeadRequest{
		Key: "head1",
	}
	services.StoreHead(headRequest)

	ctx, _ := context.WithCancel(context.Background())
	nodeRequest := requests.CreateNodeRequest{
		Key:   "node1",
		Value: "qwe123",
		Prev:  "head1",
	}
	services.StoreNode(nodeRequest, ctx)

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				key: "node1",
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				key: "aaa",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := services.RemoveNode(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("RemoveNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	node, _ := services.GetNode("node1")
	if node != nil {
		t.Errorf("Get Node unexpectedly")
	}
}

func TestGetDynamicParametersSQL(t *testing.T) {
	type args struct {
		originalSql string
		keys        []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{
				originalSql: "DELETE FROM %[1]s WHERE key IN ($1)",
				keys: []string{
					"abc",
					"def",
					"123",
				},
			},
			want: fmt.Sprintf("DELETE FROM %s WHERE key IN ($1,$2,$3)", db.TABLE),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := services.GetDynamicParametersSQL(tt.args.originalSql, tt.args.keys); got != tt.want {
				t.Errorf("GetDeleteNodesSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllKeys(t *testing.T) {
	helper.ClearDB()

	// create a head and a node
	headRequest := requests.CreateHeadRequest{
		Key: "head1",
	}
	services.StoreHead(headRequest)

	ctx, _ := context.WithCancel(context.Background())
	nodeRequest := requests.CreateNodeRequest{
		Key:   "node1",
		Value: "qwe123",
		Prev:  "head1",
	}
	services.StoreNode(nodeRequest, ctx)
	nodeRequest = requests.CreateNodeRequest{
		Key:   "node2",
		Value: "qwe123",
		Prev:  "node1",
	}
	services.StoreNode(nodeRequest, ctx)

	head, _ := services.GetHead("head1")

	type args struct {
		head *models.Head
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				head: head,
			},
			want: []string{
				"node1",
				"node2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := services.GetAllKeys(tt.args.head)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
