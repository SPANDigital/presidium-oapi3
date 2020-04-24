---
 title: deletePet
---
```
DELETE /pets/{id}
```



deletes a single pet based on the ID supplied

## Parameters

|Name|In|Type|Description|
|----|---|---|-----------|
| **id** | **path**  <br> *required*   | integer | *ID of pet to delete* |

## Responses

### 204

pet deleted

### default

unexpected error

`application/json`
```
// PLEASE RENDER A MEDIA TYPE
```

