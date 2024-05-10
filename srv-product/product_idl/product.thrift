namespace go product

struct Product {
    1: i64 id,
    2: string name,
    3: optional i64 price,
    4: optional i64 stock,
    5: optional string image,
    6: optional string description,
    7: optional string create_time,
}

service ProductService{
    Product getProduct(Product p)
}