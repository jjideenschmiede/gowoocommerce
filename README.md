# gowoocommerce

## Product categories

With the following function you can read out the stored categories. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#list-all-product-categories).

```go
// Define the request
r := &gowoocommerce.Request{
    baseUrl:        "",
    consumerKey:    "",
    consumerSecret: "",
}

// Get all product categories
categories, err := gowoocommerce.ProductCategories(r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(categories)
}
```

## Add product category

With the following function you can add the new categories. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-a-product-category).

```go
// Define the request
r := &gowoocommerce.Request{
    baseUrl:        "",
    consumerKey:    "",
    consumerSecret: "",
}

// Define body data
body := gowoocommerce.ProductCategory{
    Name:        "",
    Parent:      0,
    Description: "",
    Display:     "",
    MenuOrder:   0,
    Image: gowoocommerce.ProductCategoryImage{
        Src: "",
    },
}

// Create a new product category
category, err := gowoocommerce.AddProductCategory(body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(category)
}
```

## Update a product category

If you want to renew an existing product category, then this actually works like creating a new category. In the function, however, the ID of the category must be specified. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#update-a-product-category).

```go
// Define the request
r := &gowoocommerce.Request{
    baseUrl:        "",
    consumerKey:    "",
    consumerSecret: "",
}


// Define body data
body := gowoocommerce.ProductCategory{
    Name:        "",
    Parent:      0,
    Description: "",
    Display:     "",
    MenuOrder:   0,
    Image: gowoocommerce.ProductCategoryImage{
        Src: "",
    },
}

// Update a product category
category, err := gowoocommerce.UpdateProductCategory(0, body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(category)
}
```

## Delete a product category

To delete a category, only the ID of the category is needed. You can use this function to remove the category. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#delete-a-product-category).

```go
// Define the request
r := &gowoocommerce.Request{
    baseUrl:        "",
    consumerKey:    "",
    consumerSecret: "",
}

// Remove a product categories
category, err := gowoocommerce.DeleteProductCategory(18, true, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(category)
}
```

## Read all orders

If you want to read out all orders, you can do it page by page using the following function. 100 orders are always read out at once. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#orders).

```go
// Define the request
r := &gowoocommerce.Request{
    baseUrl:        "",
    consumerKey:    "",
    consumerSecret: "",
}

// Get all orders from the page
orders, err := gowoocommerce.Orders(1, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(orders)
}
```