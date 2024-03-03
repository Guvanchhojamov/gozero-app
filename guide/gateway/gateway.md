### 1. "Auth routes"

1. route definition

- Url: /v1/auth/signUp
- Method: POST
- Request: `signUpReq`
- Response: `signUpResp`

2. request definition



```golang
type SignUpReq struct {
	Login string `json:"login"`
	RoleId uint32 `json:"roleId"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type SignUpResp struct {
	UserId uint32 `json:"userId"`
}
```

### 2. N/A

1. route definition

- Url: /v1/auth/signIn
- Method: POST
- Request: `signInReq`
- Response: `signInResp`

2. request definition



```golang
type SignInReq struct {
	Login string `json:"login"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type SignInResp struct {
	Token string `json:"token"`
}
```

### 3. "DeleteOrder"

1. route definition

- Url: /v1/order/:orderId
- Method: DELETE
- Request: `deleteOrderReq`
- Response: `deleteOrderResp`

2. request definition



```golang
type DeleteOrderReq struct {
	OrderId uint32 `path:"orderId"`
}
```


3. response definition



```golang
type DeleteOrderResp struct {
	StatusCode uint32 `json:"statusCode"`
	Message string `json:"message"`
}
```

### 4. "CreateOrder"

1. route definition

- Url: /v1/order/create
- Method: POST
- Request: `createOrderReq`
- Response: `createOrderResp`

2. request definition



```golang
type CreateOrderReq struct {
	UserId uint32 `json:"userId"`
	ProductId uint32 `json:"productId"`
	Price float32 `json:"price"`
}
```


3. response definition



```golang
type CreateOrderResp struct {
	OrderId uint32 `json:"orderId"`
}
```

### 5. "UpdateOrder"

1. route definition

- Url: /v1/order/update
- Method: PATCH
- Request: `updateOrderReq`
- Response: `updateOrderResp`

2. request definition



```golang
type UpdateOrderReq struct {
	Id uint32 `json:"id"`
	UserId uint32 `json:"userId,optional"`
	ProductId uint32 `json:"productId,optional"`
	Price float32 `json:"price"`
}
```


3. response definition



```golang
type UpdateOrderResp struct {
	OrderId uint32 `json:"orderId"`
}
```

### 6. "Orders routes"

1. route definition

- Url: /v1/orders
- Method: GET
- Request: `getOrdersReq`
- Response: `getOrdersResp`

2. request definition



```golang
type GetOrdersReq struct {
}
```


3. response definition



```golang
type GetOrdersResp struct {
	Orders []Order `json:"orders"`
}
```

### 7. "GetById"

1. route definition

- Url: /v1/orders/:orderId
- Method: GET
- Request: `getOrderByIdReq`
- Response: `getOrderByIdResp`

2. request definition



```golang
type GetOrderByIdReq struct {
	OrderId uint32 `path:"orderId"`
}
```


3. response definition



```golang
type GetOrderByIdResp struct {
	Id uint32 `json:"id"`
	User UserOrder `json:"user"`
	Product Product `json:"product"`
	Price float32 `json:"price"`
}

type UserOrder struct {
	Id uint32 `json:"id"`
	Login string `json:"login"`
}

type Product struct {
	Id uint32 `json:"id"`
	Name string `json:"name"`
	Price float32 `json:"price"`
	CreatedAt uint32 `json:"createdAt"`
}
```

### 8. "Delete Product"

1. route definition

- Url: /v1/product/:productId
- Method: DELETE
- Request: `deleteProductReq`
- Response: `deleteProductResp`

2. request definition



```golang
type DeleteProductReq struct {
	ProductId uint32 `path:"productId"`
}
```


3. response definition



```golang
type DeleteProductResp struct {
	StatusCode uint32 `json:"statusCode"`
	Message string `json:"message"`
}
```

### 9. "Product GetById"

1. route definition

- Url: /v1/product/:productId
- Method: GET
- Request: `getProductByIdReq`
- Response: `getProductByIdResp`

2. request definition



```golang
type GetProductByIdReq struct {
	ProductId uint32 `path:"productId"`
}
```


3. response definition



```golang
type GetProductByIdResp struct {
	Product Product `json:"product"`
}

type Product struct {
	Id uint32 `json:"id"`
	Name string `json:"name"`
	Price float32 `json:"price"`
	CreatedAt uint32 `json:"createdAt"`
}
```

### 10. "Products routes"

1. route definition

- Url: /v1/product/create
- Method: POST
- Request: `createProductReq`
- Response: `createProductResp`

2. request definition



```golang
type CreateProductReq struct {
	Name string `json:"name"`
	Price float32 `json:"price"`
}
```


3. response definition



```golang
type CreateProductResp struct {
	ProductId uint32 `json:"productId"`
}
```

### 11. "Update Product"

1. route definition

- Url: /v1/product/update
- Method: PATCH
- Request: `updateProductReq`
- Response: `updateProductResp`

2. request definition



```golang
type UpdateProductReq struct {
	Id uint32 `json:"id"`
	Name string `json:"name,optional"`
	Price float32 `json:"price,optional"`
}
```


3. response definition



```golang
type UpdateProductResp struct {
	ProductId uint32 `json:"productId"`
}
```

### 12. "Get All Products"

1. route definition

- Url: /v1/products
- Method: GET
- Request: `getProductsReq`
- Response: `getProductsResp`

2. request definition



```golang
type GetProductsReq struct {
}
```


3. response definition



```golang
type GetProductsResp struct {
	Products []Product `json:"products"`
}
```

