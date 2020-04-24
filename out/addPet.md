---
 title: addPet
---
```
POST /pets
```

Tags: pet

Creates a new pet in the store. Duplicates are allowed

## Request Body

**Required**

Pet to add to the store

`application/json`
```
// PLEASE RENDER A MEDIA TYPE
```

## Responses

### 200

pet response

`application/json`
```
// PLEASE RENDER A MEDIA TYPE
```
#### Headers

|Name|Type|Description|
|----|---|-----------|
| **X-Rate-Limit-Limit**   | integer | *The number of allowed requests in the current period* |
| **X-Rate-Limit-Remaining**   | integer | *The number of remaining requests in the current period* |
| **X-Rate-Limit-Reset**   | integer | *The number of seconds left in the current period* |

### default

unexpected error

`application/json`
```
// PLEASE RENDER A MEDIA TYPE
```
#### Authentication
|Name|Scopes|
|----|------|
| api_key |  |
| petstore_auth | write:pets, read:pets |

