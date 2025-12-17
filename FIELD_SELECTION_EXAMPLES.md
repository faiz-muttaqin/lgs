# Enhanced API Documentation

## Overview
The enhanced table handler now supports field selection, filtering, sorting with detailed error messages for better Developer Experience (DX).

## Usage

### Get All Fields (Default)
```
GET /api/products?start=0&length=10
```
Returns all fields for products.

### Get Specific Fields Only
```
GET /api/products?start=0&length=10&fields=id,name,price
```
Returns only `id`, `name`, and `price` fields.

### Multiple Fields
```
GET /api/products?start=0&length=10&fields=id,name,price,stock,category_id
```
Returns selected fields: id, name, price, stock, and category_id.

### With Filtering and Field Selection
```
GET /api/products?fields=id,name,price&price[gte]=100000&price[lte]=500000
```
Returns products with price between 100k-500k, showing only id, name, and price.

### With Sorting and Field Selection
```
GET /api/products?fields=id,name,price,rating&sort=-rating,name
```
Returns products sorted by rating (desc) then name (asc), showing only selected fields.

## Benefits

1. **Reduced Bandwidth**: Only requested data is transferred
2. **Faster Response**: Less data serialization
3. **Better Performance**: Database only selects needed columns
4. **Flexible Queries**: Combine with filters and sorting

## Notes

- Field names must match the JSON field names from your model struct
- Invalid field names will return a 400 Bad Request error
- The `id` field is always included automatically for reference
- Use commas to separate multiple field names (no spaces recommended)

## Example Responses

### Without field selection:
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "iPhone Pro",
      "price": 15000000,
      "stock": 50,
      "description": "...",
      "category_id": 1,
      "subcategory_id": 1,
      "created_at": "...",
      "updated_at": "..."
    }
  ]
}
```

### With field selection (`fields=id,name,price`):
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "iPhone Pro",
      "price": 15000000
    }
  ]
}
```
