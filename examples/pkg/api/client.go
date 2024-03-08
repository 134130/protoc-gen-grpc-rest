package api

import (
	"context"
	"net/http"
)

func main() {
	client := NewUserServiceRestClient(http.DefaultClient, "http://localhost:8080")

	users, err := client.GetUsers(context.Background())
	user, err := client.GetUser(context.Background())
	user, err := client.CreateUser(context.Background(), &User{})
	err := client.DeleteUser(context.Background())
}
