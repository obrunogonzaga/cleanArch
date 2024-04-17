# Clean Architecture - Pos GO Expert - Full Cycle

## Introduction
This project is a challenge designed to test your skills in Go, Docker, and various types of APIs. The goal is to create a system for managing orders, with a REST API, a gRPC service, and a GraphQL server. You will also need to set up a database using Docker and create the necessary database migrations.  

The challenge is divided into several parts:  
- Setting up the project and the Dockerized database.
- Creating the REST API endpoint for listing orders.
- Implementing the ListOrders service in gRPC.
- Creating the ListOrders query in GraphQL.
- Writing the necessary database migrations.
- Documenting your work in a README file, including the GraphQL queries.


## Project Setup
This project is a Go application that provides a REST API, a gRPC service, and a GraphQL server for managing orders. It uses Docker for setting up the database.

### Prerequisites
- Docker
- Go 1.16 or later

### Getting Started
1. Clone the repository:
```bash
git clone https://github.com/obrunogonzaga/cleanArch.git
cd cleanArch
```

2. Build and run the Docker containers:
```bash
docker-compose up -d
```

| This will start the PostgreSQL database.

3. Run the database migrations:
```bash
make migrateup
```

4. Start the Go application:
```bash
make start
```

The application will start and listen on the following ports:  
* REST API: Port defined in configs.WebServerPort
* gRPC Service: Port defined in configs.GRPCServerPort
* GraphQL Server: Port defined in configs.GraphQLServerPort

## Usage
### REST API
To list all orders, send a GET request to `/orders`.

### gRPC Service
The gRPC service provides a `ListOrders` method in the `OrderService`.  

### GraphQL
The GraphQL server provides a `listOrders` query and a `createOrder` mutation.

**listOrders Query**
```graphql
query queryOrders {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

**createOrder Mutation**
```graphql
mutation createOrder {
  createOrder(input: {id: "eeeeee", Price: 12.2, Tax: 1.0}) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

## Testing
To test the application, you can use the provided `api.http` file, which contains HTTP requests for creating and listing orders.

## License
This project is licensed under the MIT License.