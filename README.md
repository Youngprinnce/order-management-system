# Order Management System

This is a simple Order Management System built in Go, utilizing gRPC for communication between services and HTTP for external API interactions. The system allows you to create, retrieve, and manage orders through a gRPC service, which is exposed via an HTTP server for easier interaction.

## Project Structure

The project is organized into several directories and files:

- **`common/genproto/orders/`**: Contains the generated Go code from the Protobuf definitions for the `OrderService`.
- **`kitchen/`**: Contains the HTTP server implementation that interacts with the gRPC service.
- **`orders/`**: Contains the gRPC server implementation, service logic, and controllers.
- **`protobuf/`**: Contains the Protobuf definitions for the `OrderService`.

### Key Files

- **`protobuf/orders.proto`**: The Protobuf definition for the `OrderService`.
- **`orders/main.go`**: The entry point for the gRPC server.
- **`orders/service/order.go`**: The service layer that handles the business logic for orders.
- **`orders/controller/grpc.go`**: The gRPC controller that handles incoming gRPC requests.
- **`kitchen/main.go`**: The entry point for the HTTP server.
- **`kitchen/http.go`**: The HTTP server implementation that interacts with the gRPC service.

## Architecture

```
┌───────────┐       HTTP Requests       ┌───────────────┐       gRPC Calls        ┌──────────────────┐
│           │  (POST/GET on port :1000)  │               │  (gRPC on port :9000)   │                  │
│  Client   │ ─────────────────────────> │  HTTP Server  │ ──────────────────────> │  Orders Service  │
│           │ <───────────────────────── │  (Kitchen)    │ <────────────────────── │  (gRPC Server)   │
└───────────┘     HTML/JSON Responses    └───────────────┘    gRPC Responses       └──────────────────┘
                                                                 │
                                                                 │
                                                                 ▼
                                                        ┌──────────────────┐
                                                        │   In-Memory DB   │
                                                        │  (ordersDb slice)│
                                                        └──────────────────┘
```

## Features

- **Create Order**: Allows you to create a new order with a `customerID`, `productID`, and `quantity`.
- **Get Orders**: Retrieves all orders for a specific customer.
- **Get Order**: Retrieves a single order by its `orderID`.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)
- Protocol Buffers Compiler (`protoc`)
- Go plugins for `protoc`:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/youngprinnce/order-management-system.git
   cd order-management-system
   ```

2. **Generate Protobuf code**:
   ```bash
   make gen
   ```

3. **Run the gRPC server**:
   ```bash
   make run-orders
   ```

4. **Run the HTTP server**:
   ```bash
   make run-kitchen
   ```

### Usage

#### Create an Order

Send a POST request to `http://localhost:1000/create-order` with the following JSON payload:

```json
{
  "customerID": 1,
  "productID": 101,
  "quantity": 2
}
```

#### Get Orders by Customer ID

Send a GET request to `http://localhost:1000/get-orders?customer_id=1`.

#### Get Order by ID

Send a GET request to `http://localhost:1000/get-order/{orderID}`.

### Example Requests

#### Create Order

```bash
curl -X POST http://localhost:1000/create-order -d '{"customerID": 1, "productID": 101, "quantity": 2}'
```

#### Get Orders by Customer ID

```bash
curl http://localhost:1000/get-orders?customer_id=1
```

#### Get Order by ID

```bash
curl http://localhost:1000/get-order/1
```

## API Documentation

### gRPC Service

The gRPC service is defined in `protobuf/orders.proto` and provides the following RPC methods:

- **CreateOrder**: Creates a new order.
- **GetOrders**: Retrieves all orders for a specific customer.
- **GetOrder**: Retrieves a single order by its ID.

### HTTP Endpoints

The HTTP server exposes the following endpoints:

- **POST /create-order**: Creates a new order.
- **GET /get-orders**: Retrieves all orders for a specific customer.
- **GET /get-order/{orderID}**: Retrieves a single order by its ID.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [gRPC](https://grpc.io/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [Go](https://golang.org/)

---

This README provides an overview of the project, how to set it up, and how to use it. If you have any questions or run into issues, please feel free to open an issue on GitHub.

