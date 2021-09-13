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
    Name: "",
    Image: gowoocommerce.ProductCategoryImage{
        Src: "",
    },
}

// Get all product categories
category, err := gowoocommerce.AddProductCategory(body, r)
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