# Data-collection-tree

### Running the container

```docker-compose up```

```sh
$ make run
$ chmod +x scripts/curl_insert.sh
$ # script to execute insert using POST API
$ ./curl_insert
$ chmod +x scripts/curl_query.sh
$ # script to execute querying using POST API
$ ./curl_insert
```

### Inserting data:

```
curl --location --request POST 'http://localhost:8080/v1/insert' \
--header 'Content-Type: application/json' \
--data-raw '{
  "dim": [
    {
      "key": "device",
      "val": "Web"
    },
    {
      "key": "country",
      "val": "IN"
    }
  ],
  "metrics": [
    {
      "key": "webreq",
      "val": 50
    },
    {
      "key": "timespent",
      "val": 30
    }
  ]
}'
```

### Querying data:

```
curl --location --request GET 'http://localhost:8080/v1/query' \
--header 'Content-Type: application/json' \
--data-raw '{
  "dim": [
    {
      "key": "country",
      "val": "IN"
    }
  ]
}'
```
