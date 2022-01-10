# Product CRUD Example

## Get started
- Install docker with docker-compose

## Features
-  Get all beers
-  Get a beer by id
-  Create a beer
-  Get price per beer box by currency 

## Run application
```bash
docker-compose up
```

## Examples

###Get all beers
curl --location --request GET 'http://localhost:8080/beers' \
--data-raw ''

### Create a beer
curl --location --request POST 'http://localhost:8080/beers' \
--header 'Content-Type: application/json' \
--data-raw '{
"Id": 8,
"Name": "Quilmes",
"Brewery": "ddfd",
"Country": "Chile",
"Price": 55,
"Currency": "CL"
}'

### Get a beer by id
curl --location --request GET 'http://localhost:8080/beers/<integer>'

### Get price per beer box by currency
curl --location --request GET 'http://localhost:8080/beers/50/boxprice?currency=USD&quantity=6'