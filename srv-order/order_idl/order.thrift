namespace go order

struct Order {
    1: string id,
    2: string user_id,
    3: string product_id,
    4: optional string status,
}

struct OrderResponse {
    1: string message
}

service OrderService {
    OrderResponse CreateOrder (Order o),
    OrderResponse DeleteOrder (Order o),
    i64 GetOrderStatus (Order o),
}