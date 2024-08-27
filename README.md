**Local Artisan - Simple E-Commerce**
--------------------------------
Indonesian artisans struggle to expand their market and showcase their craftsmanship beyond local communities. A platform is needed to help them display products, sell online, and promote Indonesia's cultural heritage.
So This Repository is going to answer the challenges by doing backend and database development. I'll use Golang with Gin and GORM alongside PostgreSQL for efficient data management. This approach ensures the E-commerce Platform for Local Artisans tackles the practical issues faced by Indonesian artisans while also preserving and promoting Indonesia’s cultural heritage digitally.

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
