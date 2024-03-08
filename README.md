<div align="center">
<h1>gRPC-REST</h1>
<p>REST client generator following the gRPC HTTP spec</p>
</div>

---

> [!NOTE]
> Currently only supports unary calls

## Installation

```bash
make build
PATH="$(pwd):$PATH" protoc -I=examples/proto --go_out=examples --grpc-rest_out=examples examples/proto/*.proto
```

```go

client := NewUserServiceClient("http://localhost:8080")

# Supports `google.protobuf.Empty`
users, err := client.GetUsers(context.Background())
user, err := client.GetUser(context.Background())
user, err := client.CreateUser(context.Background(), &User{})
err := client.DeleteUser(context.Background())
```