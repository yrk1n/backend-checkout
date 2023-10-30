# Checkout-Case


### Environment Configuration:
1. Create an `.env` file at the root of your project.
2. Add the following variables with your details:

DB_USER=yourname
DB_PASSWORD=yourpassword
DB_NAME=dbname


### Initial Setup:
For the initial setup, you can build and run the application using:
docker-compose up --build



## API Endpoints

### 1. Show Cart Items:
- **Method**: GET
- **URL**: /cart

**Example Success Response**:
```json
[
    {
        "item_id": 1001,
        "name": "Laptop",
        "description": "A gaming laptop",
        "price": 1200.5,
        "discounted_price": 1080.45,
        "quantity": 1,
        "seller_id": 987,
        "category_id": 1,
        "cart_id": 1,
        "type": "Default",
        "VasItems": [
            ...
        ]
    }
]
```

### 2. Add Items:
- **Method**: Post
- **URL**: /addItem

**Example Request**
```json
{
  "item_id": 1001,
  "name": "Laptop",
  "description": "A gaming laptop",
  "price": 1200.5,
  "discounted_price": 1100.5,
  "quantity": 1,
  "seller_id": 987,
  "category_id": 1,
  "cart_id": 1,
  "type": "Default",
  "vas_items": []
}
``````

### 3. Add VasItem to Item:
- **Method**: POST
- **URL**: /:itemId/add-vas

**Example Request**
```json
{
    "vas_item_id": 105,
    "name": "beko servis",
    "description": "kurulum",
    "parent_item_id": 1001,
    "category_id": 101,
    "seller_id": 2001,
    "price": 12.99,
    "quantity": 1,
    "cart_id": 1
}
``````

Notes
The responses are currently in JSON format for testing purposes. This will change to boolean values in future iterations.
The promotion logic is still under development and requires further testing.