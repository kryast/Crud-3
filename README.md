CRUD ke-3

POST
curl -X POST http://localhost:8080/customers \
-H "Content-Type: application/json" \
-d '{"name":"Ahmad", "email":"Ahmad@gmail.com", "phone":"08123456789"}'

GET
curl http://localhost:8080/customers

curl http://localhost:8080/customers/1