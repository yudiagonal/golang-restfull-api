# go-crud-api
version v.0

# gunakan postman

### Get all categories
GET http://localhost:9000/api/categories

### Create new category
POST http://localhost:9000/api/categories
tulis json di body

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


