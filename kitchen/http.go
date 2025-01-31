package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/youngprinnce/order-management-system/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("POST /create-order", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		var req orders.CreateOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		_, err := c.CreateOrder(ctx, &req)
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		w.WriteHeader(http.StatusCreated)
	})

	router.HandleFunc("/get-orders", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		customerIDStr := r.URL.Query().Get("customer_id")
		if customerIDStr == "" {
			http.Error(w, "Missing customer_id query parameter", http.StatusBadRequest)
			return
		}

		customerID, err := strconv.Atoi(customerIDStr)
		if err != nil {
			http.Error(w, "Invalid customer_id query parameter", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		res, err := c.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerID: int32(customerID),
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, res.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}
	})

	// Get order by ID
	router.HandleFunc("/get-order/{orderID}", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		orderIDStr := r.PathValue("orderID")
		if orderIDStr == "" {
			http.Error(w, "Missing orderID path parameter", http.StatusBadRequest)
			return
		}

		orderID, err := strconv.Atoi(orderIDStr)
		if err != nil {
			http.Error(w, "Invalid orderID path parameter", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		res, err := c.GetOrder(ctx, &orders.GetOrderRequest{
			OrderID: int32(orderID),
		})

		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		t := template.Must(template.New("order").Parse(orderTemplate))

		if err := t.Execute(w, res); err != nil {
			log.Fatalf("template error: %v", err)
		}

	})

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`


var orderTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Single Order</title>
</head>
<body>
	<h1>Order Details</h1>
	<table border="1">
		<tr>
			<th>Order ID</th>
			<th>Customer ID</th>
			<th>Quantity</th>
		</tr>
		<tr>
			<td>{{.OrderID}}</td>
			<td>{{.CustomerID}}</td>
			<td>{{.Quantity}}</td>
		</tr>
	</table>
</body>
</html>`
