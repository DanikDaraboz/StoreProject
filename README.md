# Ecommerce web store (GO, MongoDB)

Run docker container: ```docker compose up --build -d```
You should get this messages:  
✔ Container mongo_local     Started
✔ Container go_ecommerce   Started  

To check whether mongoDB connection is set by http request: ```curl -X GET http://localhost:8080/health```
To check docker logs: 
```docker logs go_ecommerce```
```docker logs mongo_local```


To stop and remove running container: ```docker stop go_ecommerce```