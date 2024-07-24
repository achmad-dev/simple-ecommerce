## Simple ecommerce

todo:
 - [x] add feature auth, homepage(show product), add product to cart,add payment method, checkout/orders(pay)
 - [ ] check xss injection or others
 - [ ] dockerize
 - [ ] add monitoring
 - [ ] migrate to microservice
 - [ ] integrate stripe

## to run for now
 
run docker compose first to setup infra
```
docker compose up
```

and go to gateway/cmd
```
cd gateway/cmd
```
and run
```
go run .
```

## test with curl

register user
```
curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{
  "email": "user@example.com",
  "password": "whatever"
}'
```

login
```
curl -X GET http://localhost:8080/login -H "Content-Type: application/json" -d '{"email": "user@example.com", "password": "whatever"}'
```
and u can save the token from response
```
TOKEN=<token_from_response>
```

to see all product (home)
```
url -X GET http://localhost:8080/simple-ecommerce/home -H "Authorization: Bearer $TOKEN"
```

to see product with paginate note=offset is the page like page 1 etc

```
curl -X GET "http://localhost:8080/simple-ecommerce/home-paginate?limit=10&offset=1" -H "Authorization: Bearer $TOKEN"
```
to create payment method (just mocking for now)
```
curl -X POST http://localhost:8080/simple-ecommerce/payments -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{ "bank_name": "MAANGEAK"}'
```

add a product to cart
```
curl -X POST http://localhost:8080/simple-ecommerce/cart      -H "Authorization: Bearer $TOKEN"      -H "Content-Type: application/x-www-form-urlencoded"      -d "productID=1"
```
delete product from cart
```
curl -X DELETE "http://localhost:8080/simple-ecommerce/cart/1"      -H "Authorization: Bearer $TOKEN"
```

create order/checkout
```
curl -X POST http://localhost:8080/simple-ecommerce/orders/create \
     -H "Authorization: Bearer $TOKEN"
```

to check order/checkout
```
curl -X GET http://localhost:8080/simple-ecommerce/orders/user \
     -H "Authorization: Bearer $TOKEN"
```

to pay order/checkout
```
curl -X PUT http://localhost:8080/simple-ecommerce/orders/1/pay \
     -H "Authorization: Bearer $TOKEN"
```