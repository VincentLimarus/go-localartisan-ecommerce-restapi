**Local Artisan - Simple E-Commerce**
--------------------------------
Indonesian artisans struggle to expand their market and showcase their craftsmanship beyond local communities. A platform is needed to help them display products, sell online, and promote Indonesia's cultural heritage.
So This Repository is going to answer the challenges by doing backend and database development. I'll use Golang with Gin and GORM alongside PostgreSQL for efficient data management. This approach ensures the E-commerce Platform for Local Artisans tackles the practical issues faced by Indonesian artisans while also preserving and promoting Indonesia’s cultural heritage digitally.

**Main Features**
---
- **CRUD Operations**: Enables users to perform Create, Read, Update, and Delete actions on the platform (nearly all features support these operations).

- **User Registration**: Allows users to create an account on the platform.

- **User Login**: Authenticates users and generates a JWT token, enabling access to authenticated services.

- **Artisan Registration**: Automatically updates the ```IsArtisan``` field in the User table upon successful registration as an artisan.

- **Add Product to Cart**: Adds a product to the user's cart in the database. If the user doesn't have a cart, a new one is created. If the product already exists in the cart, the quantity is updated instead of creating a new entry.

- **Checkout Product**: Enables users to purchase a single product in one order, using the "Buy Now" feature.

- **Checkout Products** in Cart: Allows users to checkout multiple products they have added to their cart.

**Key Takeaways**
---
- Upon login, cookies are generated to enable access to authentication-required services.

- Services are categorized into two main router groups:
  - Base Service: Accessible without needing cookies.
  - Auth Service: Requires cookies to access.

- To ensure security, only authenticated users can access certain services. LocalArtisans uses JWT tokens or cookies for authentication, meaning users must log in to access these services.

- Checkout Product and Add Product to Cart will only generates if the quantity of product is higher or equal the user request.

- Security and Privacy: The use of JWT tokens or cookies for authentication not only secures user data but also helps in maintaining session privacy and integrity across the platform.
  
- The ```IsArtisan``` attribute in the user table is automatically updated when a user registers as an artisan, with the boolean value changing accordingly

**Tech Stack**
---------------
**Server**: Go, Gin-Gonic, GORM, PostgreSQL, Docker

**Helper Stack**: tablePlus, pgAdmin, Postman, Visual Studio Code, Git (bash, etc)

**Static Documentation**
---
Postman Documentation: https://documenter.getpostman.com/view/33434480/2sAXjDfGAW

**Helper Files**
---
Entity Relationship Diagram: https://drawsql.app/teams/sen-2/diagrams/e-commerce-local-artisans

**Env Files**
---
```
APP_NAME = Local Artisans

DB_HOST = localArtisan

DB_PORT = 5432

DB_USER = postgres

DB_PASSWORD = database_password

DB_NAME = database_name

JWT_SECRET_KEY="UR_SECRET_KEY"

TIMEOUT="3s"

```
**Pull / Download in ur Local**
---
Clone this repository into ur Local
```
git clone https://github.com/VincentLimarus/E-Commerce-LocalArtisans.git
```
Go to the Directory
```
cd E-Commerce-LocalArtisans
```
Install Dependency
```
go get .
```
Run Server 
```
go run .
```

**Download from Docker / Pull from Docker**
---
Docker Pull Command
```
docker pull vincentlim27/localartisanv1
```
