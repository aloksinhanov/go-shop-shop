sudo docker ps -a 
CONTAINER ID        IMAGE                 COMMAND                  CREATED              STATUS                      PORTS                    NAMES
a157691cc601        golang:1.14           "/opt/experiments/a.…"   14 seconds ago       Up 13 seconds               0.0.0.0:8001->8001/tcp   shopshop_review_1
456dc6a251d9        golang:1.14           "/opt/experiments/a.…"   About a minute ago   Up About a minute           0.0.0.0:8002->8002/tcp   shopshop_catalog_1

curl http://localhost:8001/reviews/1234
{"rating":5,"comment":"A very durable product","productId":100}alok@continuum-virtualbox:~/Juno/src/github.com/aloksinhanov/shopshop$ 

curl http://localhost:8002/products/1234
{"id":2,"description":"A matt black kitchen sink","ratingId":5}alok@continuum-virtualbox:~/Juno/src/github.com/aloksinhanov/shopshop$ 

