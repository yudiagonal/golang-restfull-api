# go-crud-api
version v.0

# gunakan postman
# API Spec

## Create Product

Request :
- Method : POST
- Port : `9000`
- Endpoint : `/api/products`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "nama_barang":"buku",
    "harga":"12000",
    "jenis":"ersdf",
    "meta_keyword":"hjtyj"
}
```

Response :

```json 
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": 1,
        "nama_barang": "buku",
        "harga": "12000",
        "jenis": "ersdf",
        "meta_keyword": "hjtyj"
    }
}
```

## FindById

Request :
- Method : GET
- Endpoint : `/api/products/{id_product}`
- Header :
    - Accept: application/json

Response :

```json 
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": 1,
        "nama_barang": "buku",
        "harga": "12000",
        "jenis": "ersdf",
        "meta_keyword": "hjtyj"
    }
}
```

## Update Product

Request :
- Method : PUT
- Endpoint : `/api/products/{id_product}`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "nama_barang":"buku gambar",
    "harga":"15000",
    "jenis":"ersdf",
    "meta_keyword":"hjtyj"
}
```

Response :

```json 
{
   "code": 200,
    "status": "OK",
    "data": {
        "id": 1,
        "nama_barang": "buku gambar",
        "harga": "15000",
        "jenis": "ersdf",
        "meta_keyword": "hjtyj"
    }
}
```

## FindAll Product 

Request :
- Method : GET
- Endpoint : `/api/products`
- Header :
    - Accept: application/json

Response :

```json 
{
    "code": 200,
    "status": "OK",
    "data" : [
        {
              "id": 1,
        "nama_barang": "buku gambar",
        "harga": "15000",
        "jenis": "ersdf",
        "meta_keyword": "hjtyj"
        },
        {
             "id": 2,
             "nama_barang": "pensil",
             "harga": "2000",
             "jenis": "ersdf",
             "meta_keyword": "hjtyj"
         }
    ]
}
```

## Delete Product

Request :
- Method : DELETE
- Endpoint : `/api/products/{id_product}`
- Header :
    - Accept: application/json

Response :

```json 
{
    "code": 200,
    "status": "OK",
}
```

## Category
### Get all categories
GET http://localhost:9000/api/categories

### Create new category
POST http://localhost:9000/api/categories
tulis json di body

- Body :
```json
{
  "name" : "apa aja boleh"
}

```

### Get category by Id
GET http://localhost:9000/api/categories/id


### Update category by id
PUT http://localhost:9000/api/categories/id

```json
{
  "name" : "apa aja boleh"
}

```
### Delete category by id
DELETE http://localhost:9000/api/categories/id


