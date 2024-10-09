#Backend microservices with kubernetes Pods

## Basket service

- using redis as backing store for user id + items
- endpoints
  - GET /my-basket = returns the current basket info
  - POST /my-basket = adds and item to the basket
    - BODY: [{ itemId: string, amount: int }]
      - an array of items
  - DELETE /my-basket = clears my basket
    - in order to delete some item, you must use POST with amount 0

## Catalog service

- postgres as database
- endpoints
  - GET /catalog = returns all items for the shop
    - supports pagination via query params
      - pN = page number
      - pS = page size (max 25 items)

## Identity service

- postgres as database

## Order service

- postgres as database
- endpoints

## API gateway - Ingress
