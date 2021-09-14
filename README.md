# gowoocommerce

## Products

If you want to read out products, then you can do this via the following function. It is important that the page number is stored in the function. There are always 100 products loaded per page. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#list-all-products).

```go
// Define the request
r := &gowoocommerce.Request{
    baseUrl:        "",
    consumerKey:    "",
    consumerSecret: "",
}

// Get all products from the page
products, err := gowoocommerce.Products(1, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(products)
}
```

## Create a product

If you want to create a new product, you can do it as follows. For this, some data is mandatory. The function also returns the error codes from WooCommerce. 

The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-a-product).

```go
// Define the request
r := &gowoocommerce.Request{
baseUrl:        "",
consumerKey:    "",
consumerSecret: "",
}

// Define product body
body := gowoocommerce.ProductsBody{
    Name:              "J&J Interface V2",
    Slug:              "jj-interface",
    Type:              "simple",
    Status:            "publish",
    Featured:          false,
    CatalogVisibility: "visible",
    Description:       "<h1>J&J Interface</h1><p>Das ist eine Beschreibung.</p>",
    ShortDescription:  "Unsere Schnittstelle zwischen Warenwirtschaft und Online Shop.",
    Sku:               "",
    Price:             "2800",
    RegularPrice:      "3000",
    SalePrice:         "",
    DateOnSaleFrom:    nil,
    DateOnSaleFromGmt: nil,
    DateOnSaleTo:      nil,
    DateOnSaleToGmt:   nil,
    OnSale:            false,
    Purchasable:       false,
    TotalSales:        0,
    Virtual:           false,
    Downloadable:      false,
    Downloads:         nil,
    DownloadLimit:     -1,
    DownloadExpiry:    -1,
    ExternalUrl:       "",
    ButtonText:        "",
    TaxStatus:         "taxable",
    TaxClass:          "",
    ManageStock:       true,
    StockQuantity:     120,
    Backorders:        "no",
    BackordersAllowed: false,
    Backordered:       false,
    LowStockAmount:    nil,
    SoldIndividually:  false,
    Weight:            "",
    Dimensions:        gowoocommerce.ProductsBodyDimensions{},
    ShippingRequired:  false,
    ShippingTaxable:   false,
    ShippingClass:     "",
    ShippingClassId:   0,
    ReviewsAllowed:    false,
    AverageRating:     "",
    RatingCount:       0,
    UpsellIds:         nil,
    CrossSellIds:      nil,
    ParentId:          0,
    PurchaseNote:      "",
    Categories:        []gowoocommerce.ProductsBodyCategories{},
    Tags:              nil,
    Images:            []gowoocommerce.ProductsBodyImages{},
    Attributes:        nil,
    DefaultAttributes: nil,
    Variations:        nil,
    GroupedProducts:   nil,
    MenuOrder:         0,
    RelatedIds:        nil,
    MetaData:          nil,
    StockStatus:       "instock",
}

// Add an image
body.Images = append(body.Images, gowoocommerce.ProductsBodyImages{
    Src:  "https://shop.jj-ideenschmiede.de/media/3e/81/e8/1621082068/jj-interface-shop.png",
    Name: "J&J Interface",
    Alt:  "jj-interface",
})

// Create new product
products, err := gowoocommerce.CreateProduct(body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(products)
}
```

## Update a product

If you want to update a product, it works like creating a product. With the difference that you need the ID of the product. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#update-a-product).

```go
// Define the request
r := &gowoocommerce.Request{
baseUrl:        "",
consumerKey:    "",
consumerSecret: "",
}

// Define product body
body := gowoocommerce.ProductsBody{
    Name:              "J&J Interface",
    Slug:              "jj-interface",
    Type:              "simple",
    Status:            "publish",
    Featured:          false,
    CatalogVisibility: "visible",
    Description:       "<h1>J&J Interface</h1><p>Das ist eine Beschreibung.</p>",
    ShortDescription:  "Unsere Schnittstelle zwischen Warenwirtschaft und Online Shop.",
    Sku:               "",
    Price:             "2800",
    RegularPrice:      "3000",
    SalePrice:         "",
    DateOnSaleFrom:    nil,
    DateOnSaleFromGmt: nil,
    DateOnSaleTo:      nil,
    DateOnSaleToGmt:   nil,
    OnSale:            false,
    Purchasable:       false,
    TotalSales:        0,
    Virtual:           false,
    Downloadable:      false,
    Downloads:         nil,
    DownloadLimit:     -1,
    DownloadExpiry:    -1,
    ExternalUrl:       "",
    ButtonText:        "",
    TaxStatus:         "taxable",
    TaxClass:          "",
    ManageStock:       true,
    StockQuantity:     120,
    Backorders:        "no",
    BackordersAllowed: false,
    Backordered:       false,
    LowStockAmount:    nil,
    SoldIndividually:  false,
    Weight:            "",
    Dimensions:        gowoocommerce.ProductsBodyDimensions{},
    ShippingRequired:  false,
    ShippingTaxable:   false,
    ShippingClass:     "",
    ShippingClassId:   0,
    ReviewsAllowed:    false,
    AverageRating:     "",
    RatingCount:       0,
    UpsellIds:         nil,
    CrossSellIds:      nil,
    ParentId:          0,
    PurchaseNote:      "",
    Categories:        []gowoocommerce.ProductsBodyCategories{},
    Tags:              nil,
    Images:            []gowoocommerce.ProductsBodyImages{},
    Attributes:        nil,
    DefaultAttributes: nil,
    Variations:        nil,
    GroupedProducts:   nil,
    MenuOrder:         0,
    RelatedIds:        nil,
    MetaData:          nil,
    StockStatus:       "instock",
}

// Add an image
body.Images = append(body.Images, gowoocommerce.ProductsBodyImages{
    Src:  "https://shop.jj-ideenschmiede.de/media/3e/81/e8/1621082068/jj-interface-shop.png",
    Name: "J&J Interface",
    Alt:  "jj-interface",
})

// Create new product
products, err := gowoocommerce.UpdateProduct(22, body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(products)
}
```

## Delete a product

If you want to remove a product, you need the ID of the product. This must be stored with the function. Also the value force must be answered with true or false. If this is true, then the product is permanently removed. If the value is answered with false, then the product ends up in the recycle bin for the time being.

The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#delete-a-product).

```go
// Define the request
r := &gowoocommerce.Request{
    baseUrl:        "",
    consumerKey:    "",
    consumerSecret: "",
}

// Delete a product
products, err := DeleteProduct(17, false, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(products)
}
```

## Product variations

If you want to read out all variants for a product, you can do this as follows. The ID of the variant product is required. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#list-all-product-variations).

```go
// Define the request
r := &gowoocommerce.Request{
    baseUrl:        "",
    consumerKey:    "",
    consumerSecret: "",
}

// Read all product variations
productVariations, err := gowoocommerce.ProductVariations(33, 1, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(productVariations)
}
```

## Create a product variant

If you want to create a new variant, you can do so using this function. Thereby some values are needed. The library returns these as values if they should occur. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-a-product-variation).

```go
// Define the request
r := &gowoocommerce.Request{
baseUrl:        "",
consumerKey:    "",
consumerSecret: "",
}

// Define body
body := gowoocommerce.ProductVariationsBody{
    Description:       "Our Version for Shopify.",
    Permalink:         "",
    Sku:               "",
    Price:             "255",
    RegularPrice:      "",
    SalePrice:         "215",
    DateOnSaleFrom:    nil,
    DateOnSaleFromGmt: nil,
    DateOnSaleTo:      nil,
    DateOnSaleToGmt:   nil,
    OnSale:            false,
    Status:            "publish",
    Purchasable:       false,
    Virtual:           false,
    Downloadable:      false,
    Downloads:         nil,
    DownloadLimit:     -1,
    DownloadExpiry:    -1,
    TaxStatus:         "taxable",
    TaxClass:          "",
    ManageStock:       false,
    StockQuantity:     nil,
    StockStatus:       "instock",
    Backorders:        "no",
    BackordersAllowed: false,
    Backordered:       false,
    Weight:            "",
    Dimensions:        gowoocommerce.ProductVariationsBodyDimensions{},
    ShippingClass:     "",
    ShippingClassId:   0,
    Image: gowoocommerce.ProductVariationsBodyImages{
        Src:  "https://shop.jj-ideenschmiede.de/thumbnail/3e/81/e8/1621082068/jj-interface-shop_1920x1920.png",
        Name: "J&J Interface",
        Alt:  "J&J Interface",
    },
    Attributes: []gowoocommerce.ProductVariationsBodyAttributes{},
    MenuOrder:  0,
    MetaData:   nil,
}

// Add attribute
body.Attributes = append(body.Attributes, gowoocommerce.ProductVariationsBodyAttributes{
    Name:   "Shopsoftware",
    Option: "WooCommerce",
})

// Create a new variant to a product
productVariations, err := gowoocommerce.CreateProductVariations(33, body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(productVariations)
}
```

## Update a product variant

If you want to update a variant, you can do this as follows. Everything is the same as when you create a variant. Only the ID of the variant must be stored. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#update-a-product-variation).

```go
// Define the request
r := &gowoocommerce.Request{
baseUrl:        "",
consumerKey:    "",
consumerSecret: "",
}

// Define body
body := gowoocommerce.ProductVariationsBody{
    Description:       "Our Version for Shopify.",
    Permalink:         "",
    Sku:               "",
    Price:             "255",
    RegularPrice:      "",
    SalePrice:         "215",
    DateOnSaleFrom:    nil,
    DateOnSaleFromGmt: nil,
    DateOnSaleTo:      nil,
    DateOnSaleToGmt:   nil,
    OnSale:            false,
    Status:            "publish",
    Purchasable:       false,
    Virtual:           false,
    Downloadable:      false,
    Downloads:         nil,
    DownloadLimit:     -1,
    DownloadExpiry:    -1,
    TaxStatus:         "taxable",
    TaxClass:          "",
    ManageStock:       false,
    StockQuantity:     nil,
    StockStatus:       "instock",
    Backorders:        "no",
    BackordersAllowed: false,
    Backordered:       false,
    Weight:            "",
    Dimensions:        gowoocommerce.ProductVariationsBodyDimensions{},
    ShippingClass:     "",
    ShippingClassId:   0,
    Image: gowoocommerce.ProductVariationsBodyImages{
        Src:  "https://shop.jj-ideenschmiede.de/thumbnail/3e/81/e8/1621082068/jj-interface-shop_1920x1920.png",
        Name: "J&J Interface",
        Alt:  "J&J Interface",
    },
    Attributes: []gowoocommerce.ProductVariationsBodyAttributes{},
    MenuOrder:  0,
    MetaData:   nil,
}

// Add attribute
body.Attributes = append(body.Attributes, gowoocommerce.ProductVariationsBodyAttributes{
    Name:   "Shopsoftware",
    Option: "WooCommerce",
})

// Update a variant to a product
productVariations, err := gowoocommerce.UpdateProductVariations(33, 45, body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(productVariations)
}
```

## Delete a product variant

If you want to remove a variant, you can do this with the following function. This requires a productID and a variantID. The description of the API endpoint can be found [here](https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#delete-a-product-variation).

```go
// Define the request
r := &gowoocommerce.Request{
baseUrl:        "",
consumerKey:    "",
consumerSecret: "",
}

// Remove a variant from a product
productVariations, err := gowoocommerce.DeleteProductVariations(33, 37, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(productVariations)
}
```

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