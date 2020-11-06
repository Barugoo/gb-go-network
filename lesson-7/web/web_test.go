package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	*http.Client
	URL string
}

type mockDB struct {
}

func (mock *mockDB) GetLists() ([]TaskList, error) {
	return []TaskList{
		{
			Name: "dsds",
			Description: "dsds",
		},
	}, nil
}
func (mock *mockDB)  GetList(id string) (TaskList, error) {
	return TaskList{}, nil
}
func (mock *mockDB)  CreateList(list TaskList) (TaskList, error) {
	return TaskList{}, nil
}
func (mock *mockDB) UpdateList(list TaskList) (TaskList, error) {
	return TaskList{}, nil
}

func TestCreteListRequest(t *testing.T) {
	testController := &controller{
		db: &mockDB{},
	}
	srv := httptest.NewServer(setupServer(testController))
	defer srv.Close()

	cli := &Client{
		Client: srv.Client(),
		URL: srv.URL,
	}
	resp, _ := cli.Get(fmt.Sprintf("%s/lists", cli.URL))
	if resp.StatusCode != 200 {
	   t.Errorf("/lists endpoint returned wrong response; expected %d; actual %d", 200, resp.StatusCode)
	}
  }