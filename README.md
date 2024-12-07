# KV-Store
### Simple in memory Database for general purposes

#### Setup:

Clone the repo:
```
  git clone https://github.com/mESmaC/kvstore.git
```

To run (In the top level of the project):
```
go run .
```

#### Requests:
```
'/create_db' request-method: 'POST'
'/insert_pair' request-method: 'POST'
'/query_db' request-method: 'GET'
'/delete_key' request-method: 'POST'
'/delete_db' request-method: 'POST'
```

#### Body format (RAW):

CreateDB: 
```
{
  "name": string
}
```

Insert Pair:
```
{
  "name": string
  "key": string
  "value": interface (Any type)
}
```

Insert Pair Examples:
```
{
  "name": "myDatabase",
  "key": "username",
  "value": "john_doe"
}

{
  "name": "myDatabase",
  "key": "user_id",
  "value": 123
}

{
  "name": "myDatabase",
  "key": "is_active",
  "value": true
}

{
  "name": "myDatabase",
  "key": "roles",
  "value": ["admin", "editor", "viewer"]
}

{
  "name": "myDatabase",
  "key": "user_profile",
  "value": {
    "first_name": "John",
    "last_name": "Doe",
    "age": 30
  }
}
```

Query DB:
```
{
  "name": string
  "key": string
}
```

Delete Key:
```
{
  "name": string
  "key": string
}
```

Delete DB:
```
{
  "name": string
}
```
