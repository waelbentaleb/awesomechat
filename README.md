
## About Awesome Chat

This project is POC of chat application that uses gRPC as main communication protocol.
Using Golang for the backend, React for the frontend and Envoy as proxy.
After connecting, users can send one-to-one messages by providing the receiver username.
Users can also create and join groups where messages are broadcasting to all group members.

## Getting Started

These instructions will give you a copy of the project up and running on
your local machine for development and testing purposes.

### Prerequisites

You need docker installed in your machine in order to properly run the project

### Running

Docker compose file is provided to ensure all required dependencies and configurations.
You need to have docker installed in your host and then you can run the following command:

    docker compose up

## Golang backend

The backend app was developed using Golang that implements a gRPC server and exposes those methods:

- `CreateUser` create a new user
- `Connect` connect a user to messages stream
- `SendMessage` send message to a user or a group channel
- `CreateGroup` create a new group with the current user as first member
- `JoinGroup` join created group
- `LeftGroup` leave a joined group
- `ListChannels` list all connected users and available groups

All methods and messages are described in the proto file [here](./golang-awesomechat/contracts/awesomechat.proto).

This server implements gRPC reflections, so you can easily retrieve all the specification using a gRPC client like [evans](https://github.com/ktr0731/evans).

### Unit tests

You can run tests using this command

    go test ./... -cover

## React frontend

The frontend app was developed using React. It's only for testing the backend server and showcasing the potential of [grpc-web](https://github.com/grpc/grpc-web).
[Envoy proxy](https://www.envoyproxy.io/) is required to connect React app to our backend server.

## Demo

Live demonstration of 3 chat users that illustrate:

- One-to-one communication
- One-to-many communication (Group channel)



![](./demo.gif)




