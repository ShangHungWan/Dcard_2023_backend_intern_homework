package controllers_test

import (
	"bytes"
	"context"
	"fmt"
	"key-value-system/controllers"
	"key-value-system/helper"
	"key-value-system/requests"
	"key-value-system/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreHead(t *testing.T) {
	helper.ClearDB()

	router := helper.SetUpRouter()
	router.POST("/head", controllers.StoreHead)
	router.GET("/head/:key", controllers.ShowHead)

	type args struct {
		key string
	}
	tests := []struct {
		name   string
		args   args
		status int
	}{
		{
			name: "normal",
			args: args{
				key: "qweasd123qwe123",
			},
			status: http.StatusCreated,
		},
		{
			name: "no input",
			args: args{
				key: "",
			},
			status: http.StatusBadRequest,
		},
		{
			name: "duplicated",
			args: args{
				key: "qweasd123qwe123",
			},
			status: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody := fmt.Sprintf(`{"key": "%s"}`, tt.args.key)

			req, _ := http.NewRequest("POST", "/head", bytes.NewReader([]byte(jsonBody)))
			req.Header.Add("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}

	path := fmt.Sprintf("/head/qweasd123qwe123")
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestStoreNode(t *testing.T) {
	helper.ClearDB()

	// create a head for `prev`
	request := requests.CreateHeadRequest{
		Key: "node1",
	}
	services.StoreHead(request)

	router := helper.SetUpRouter()
	router.POST("/node", controllers.StoreNode)
	router.GET("/node/:key", controllers.ShowNode)

	type args struct {
		key   string
		value string
		prev  string
	}
	tests := []struct {
		name   string
		args   args
		status int
	}{
		{
			name: "normal",
			args: args{
				key:   "qweasd123qwe123",
				value: "qwe123",
				prev:  "node1",
			},
			status: http.StatusCreated,
		},
		{
			name: "no input",
			args: args{
				key:   "",
				value: "",
				prev:  "",
			},
			status: http.StatusBadRequest,
		},
		{
			name: "duplicated",
			args: args{
				key:   "qweasd123qwe123",
				value: "qwe123",
				prev:  "node1",
			},
			status: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody := fmt.Sprintf(`{"key": "%s", "value": "%s", "prev": "%s"}`, tt.args.key, tt.args.value, tt.args.prev)

			req, _ := http.NewRequest("POST", "/node", bytes.NewReader([]byte(jsonBody)))
			req.Header.Add("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}

	path := fmt.Sprintf("/node/qweasd123qwe123")
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestShowHead(t *testing.T) {
	helper.ClearDB()

	// create a head
	request := requests.CreateHeadRequest{
		Key: "head1",
	}
	services.StoreHead(request)

	router := helper.SetUpRouter()
	router.GET("/head/:key", controllers.ShowHead)

	type args struct {
		key string
	}
	tests := []struct {
		name     string
		args     args
		status   int
		response string
	}{
		{
			name: "normal",
			args: args{
				key: "head1",
			},
			status:   http.StatusOK,
			response: "head1",
		},
		{
			name: "not found",
			args: args{
				key: "aaa",
			},
			status:   http.StatusNotFound,
			response: "not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/head/%s", tt.args.key)
			req, _ := http.NewRequest("GET", path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
			assert.Contains(t, w.Body.String(), tt.response)
		})
	}
}

func TestShowNode(t *testing.T) {
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

	router := helper.SetUpRouter()
	router.GET("/node/:key", controllers.ShowNode)

	type args struct {
		key string
	}
	tests := []struct {
		name     string
		args     args
		status   int
		response string
	}{
		{
			name: "normal",
			args: args{
				key: "node1",
			},
			status:   http.StatusOK,
			response: "node1",
		},
		{
			name: "not found",
			args: args{
				key: "aaa",
			},
			status:   http.StatusNotFound,
			response: "not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/node/%s", tt.args.key)
			req, _ := http.NewRequest("GET", path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
			assert.Contains(t, w.Body.String(), tt.response)
		})
	}
}

func TestDeleteHead(t *testing.T) {
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

	router := helper.SetUpRouter()
	router.DELETE("/head/:key", controllers.RemoveHead)
	router.GET("/head/:key", controllers.ShowHead)
	router.GET("/node/:key", controllers.ShowNode)

	type args struct {
		key string
	}
	tests := []struct {
		name   string
		args   args
		status int
	}{
		{
			name: "normal",
			args: args{
				key: "head1",
			},
			status: http.StatusNoContent,
		},
		{
			name: "not found",
			args: args{
				key: "aaa",
			},
			status: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/head/%s", tt.args.key)
			req, _ := http.NewRequest("DELETE", path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}

	path := fmt.Sprintf("/head/head1")
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	path = fmt.Sprintf("/node/node1")
	req, _ = http.NewRequest("GET", path, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteNode(t *testing.T) {
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

	router := helper.SetUpRouter()
	router.DELETE("/node/:key", controllers.RemoveNode)
	router.GET("/node/:key", controllers.ShowNode)

	type args struct {
		key string
	}
	tests := []struct {
		name   string
		args   args
		status int
	}{
		{
			name: "normal",
			args: args{
				key: "node1",
			},
			status: http.StatusNoContent,
		},
		{
			name: "not found",
			args: args{
				key: "aaa",
			},
			status: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/node/%s", tt.args.key)
			req, _ := http.NewRequest("DELETE", path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}

	path := fmt.Sprintf("/node/node1")
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
