---
runme:
  id: 01HX76ZK1SK79VM7CJ415Y58ZT
  version: v3
---

# 🐱EniQilo

EniQilo is an app that allows staff to add product and customer to buy the products from the store.

## 🌟Features

EniQilo offers the following features:

- **Authentication**:
- Staff registration
- Staff login
- **Product Management (CRUD)**:
- Create new product
- View existing product
- Update product
- Delete product
- **Search SKU**:
- Search product by SKU
- **Checkout**:
- Customer registration
- View customers
- Checkout products
- View checkout history

## ⛔️ Requirements

Before running this app make sure you have this installed on your puter :

- [Golang 1.22.0](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [docker](https://docs.docker.com/engine/install/ubuntu/)

## 🎖Prerequisite

To run the application, follow these steps before run the program:

1. Make sure you have Golang, PostgreSQL, Golang Migrate, and Docker installed and configured on your system.
2. Clone this repository:

```bash {"id":"01HXBJ7XEECXDYSM92BBJFY4V5"}

git clone https://github.com/sedelinggam/eniqilo.git

```

3. Navigate to the project directory:

```bash {"id":"01HXBJ7XEECXDYSM92BC18F9P1"}

cd eniqilo

```

4. Run the following command to install dependencies:

```bash {"id":"01HXBJ7XEECXDYSM92BDDP7D4A"}

go mod download

```

5. Run the following command to create environment for the application:

```bash {"id":"01HXBJ7XEECXDYSM92BG3X43GZ"}

mv .env.sample .env

```

## 🚀 Run The Program

1. **Setting Up Environment Variables**

Before starting the application, you need to set up the following environment variables:

- `DB_NAME`: Name of your PostgreSQL database
- `DB_PORT`: Port of your PostgreSQL database (default: 5432)
- `DB_HOST`: Hostname or IP address of your PostgreSQL server
- `DB_USERNAME`: Username for your PostgreSQL database
- `DB_PASSWORD`: Password for your PostgreSQL database
- `DB_PARAMS`: Additional connection parameters for PostgreSQL (e.g., sslmode=disabled)
- `JWT_SECRET`: Secret key used for generating JSON Web Tokens (JWT)
- `BCRYPT_SALT`: Salt for password hashing (use a higher value than 8 in production!)

2. **Database Migrations**

- Apply migrations to the database:

```bash {"id":"01HXBJ7XEECXDYSM92BKBXS47Z"}

make migrate-dev

```

3. **Running the Application**

```bash {"id":"01HXBJ7XEECXDYSM92BNS4FSD8"}

make run

```

You can access the application in your web browser at http://localhost:8080

## 🐋 Build Image

Make sure you already installed Docker on your computer.
Adjust the `.env` file to make sure it's connected to the database and then you can build the Docker image by running:
```bash
make build-image
```

After the image is build, you can run it using:
```bash
docker run -p 8080:8080 eniqilo -d
```