# grpc-go-course

gRPC w/ Golang - couse from Udemy

## HTTP 2

Is more secure, uses less bandwidth and has less chatter than HTTP 1:

- Most of the web works with HTTP 1
  - It's slower (opens TCP connection for each request)
  - Doesn't compress headers
  - Only accepts Request/Response
- HTTP 2
  - One TCP connection
  - Supports server push
    - Server can push multiple messages
  - Supports multiplexing
    - Can push multiple messages in parallel over the same TCP connection
    - Processes Request/Response faster
  - SSL connection by default
  - **Headers and data are compressed into binary** -> relates to protobuf

## Types of API in gRPC

### Unary

- Closest from REST
- Request/Response

In Protobuf:

```go
rpc Greet(GreetRequest) returns (GreetResponse) {};
```

### Server Streaming

- Client sends 1 request
- Server returns one or more responses
- Interesting when client wants to get updated in real time

In Protobuf:

```go
rpc GreetManyTimes(GreetRequest) returns (stream GreetResponse) {};
```

### Client Streaming

- Client sends 1 or multiple requests
- Server returns 1 response
- Interesting for uploading/updating data

In Protobuf:

```go
rpc LongGreet(stream GreetRequest) returns (GreetResponse) {};
```

### Bi direcional Streaming

- Client and Server can send multiple requests/responses
- Responses can arrive in any order desired

In Protobuf:

```go
rpc GreetEveryone(stream GreetRequest) returns (stream GreetResponse) {};
```

## Scalability

- Server is async
- Client can be async or blocking
- Example: Google claims 10 B requests/sec

## Security

- Schema-based serialization (not human-readable)
- Easy SSL certificates initialization
- Interceptors for Auth

## gRPC vs REST

|   gRPC        |   REST            |
|---------------|-------------------|
|Protobuf       |JSON               |
|HTTP 2         |HTTP 1             |
|Streaming      |Unary              |
|Bi directional |Client -> Server   |
|Free API design|GET/POST/UPDATE/...|

## Evans - gRPC client CLI for debugging

[Reference](https://github.com/ktr0731/evans)
