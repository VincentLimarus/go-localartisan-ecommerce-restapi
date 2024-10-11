**LocalArtisans - Rest API E-Commerce**
--------------------------------
LocalArtisans is a backend service project that centers on building a robust RESTful API for an e-commerce platform aimed at helping Indonesian artisans. The project focuses on efficient communication between services, ensuring that every user request is handled securely with proper authentication.

The backend is built using Golang with the Gin framework and GORM for handling database operations. All incoming API requests are managed and stored in a PostgreSQL database, ensuring data integrity and reliability. The goal is to provide a seamless and scalable API structure that supports the platform’s functionality while maintaining strong security measures.

This project focuses solely on backend development, ensuring that the services provided to artisans are well-supported and secure, allowing their products and craftsmanship to be efficiently managed and showcased.

**Main Features**
---
- **CRUD Operations**: Allows users and artisans to perform Create, Read, Update, and Delete actions on the platform. Nearly all features support these operations; for example, artisans can create or add their items to the platform, and they can also update or delete them as needed.

- **User Registration**: Allows users to create an account on the platform.

- **User Login**: Authenticates users and generates a JWT token, enabling access to authenticated services.

- **Artisan Registration**: Automatically updates the ```IsArtisan``` field in the User table upon successful registration as an artisan.

- **Add Product to Cart**: Adds a product to the user's cart in the database. If the user doesn't have a cart, a new one is created. If the product already exists in the cart, the quantity is updated instead of creating a new entry.

- **Checkout Product**: Enables users to purchase a single product in one order, using the "Buy Now" feature.

- **Checkout Products in Cart**: Allows users to checkout multiple products they have added to their cart.

- **Pay Order**: Allow the user to pay for an order they placed earlier. Once the order is paid, the quantity of the product will be decreased according to the user’s request.

- **Cancel Order**: Enable the system to cancel an order if there is a discrepancy between the user and the artisan during the ordering process. The product quantity will be restored to its original amount, and the order will be marked as canceled.

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

**Helper Stack**: tablePlus, pgAdmin, Postman, Visual Studio Code, DrawSQL, Git (bash, etc)

**Static Documentation**
---
Postman Documentation: https://documenter.getpostman.com/view/33434480/2sAXjDfGAW

**Helper Files**
---
Entity Relationship Diagram: https://drawsql.app/teams/sen-2/diagrams/e-commerce-local-artisans

**Env Files**
---
```
APP_NAME = "LocalArtisans"

DB_HOST = "localArtisans"

DB_PORT = "5432"

DB_USER = "postgres"

DB_PASSWORD = "database_password"

DB_NAME = "database_name"

JWT_SECRET_KEY="UR_SECRET_KEY"

TIMEOUT="3s"
```
**Pull / Download in ur Local**
---
Clone this repository into ur Local
```
git clone https://github.com/VincentLimarus/go-localartisan-ecommerce-restapi.git
```
Go to the Directory
```
cd LocalArtisan-RestAPI
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
