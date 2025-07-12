CRUD ke-3

POST
curl -X POST http://localhost:8080/customers \
-H "Content-Type: application/json" \
-d '{"name":"Ahmad", "email":"Ahmad@gmail.com", "phone":"08123456789"}'

GET
curl http://localhost:8080/customers

curl http://localhost:8080/customers/1

PUT
curl -X PUT http://localhost:8080/customers/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Ahmad S.",
  "email": "ahmad_updated@example.com",
  "phone": "089998888777"
}'