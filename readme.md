---
runme:
  id: 01HX76ZK1SK79VM7CJ415Y58ZT
  version: v3
---


# 🐱Cats Social

  

Cats Social is an app that allows cat owners to match their cats with each other.

  ## 🌟Features

  

Cats Social offers the following features:

  

-  **Authentication**:

- User registration

- User login

-  **Cat Management (CRUD)**:

- Create new cat profiles

- View existing cat profiles

- Update cat profiles

- Delete cat profiles

-  **Matching**:

- Match your cat with other cats

- View matching cats

- Approve or reject matches

- Delete matches

  


## ⛔️ Requirements

  

Before running this app make sure you have this installed on your puter :



-  [Golang 1.22.0](https://go.dev/dl/)

- [PostgreSQL](https://www.postgresql.org/download/)

- [golang-migrate](https://github.com/golang-migrate/migrate)

  

## 🎖Prerequisite

  

To run the application, follow these steps before run the program:

  

1. Make sure you have Golang, PostgreSQL, and Golang Migrate installed and configured on your system.

  

2. Clone this repository:

  

```bash

git clone https://github.com/sedelinggam/cats-social.git

```

  

3. Navigate to the project directory:

  

```bash

cd cats-social

```

  

4. Run the following command to install dependencies:

```bash

go mod download

```
5. Run the following command to create environment for the application:

```bash

mv .env.sample .env

```

## 🚀 Run The Program

  

1.  **Setting Up Environment Variables**

  

Before starting the application, you need to set up the following environment variables:

  

-  `DB_NAME`: Name of your PostgreSQL database

-  `DB_PORT`: Port of your PostgreSQL database (default: 5432)

-  `DB_HOST`: Hostname or IP address of your PostgreSQL server

-  `DB_USERNAME`: Username for your PostgreSQL database

-  `DB_PASSWORD`: Password for your PostgreSQL database

-  `DB_PARAMS`: Additional connection parameters for PostgreSQL (e.g., sslmode=disabled)

-  `JWT_SECRET`: Secret key used for generating JSON Web Tokens (JWT)

-  `BCRYPT_SALT`: Salt for password hashing (use a higher value than 8 in production!)

  

2.  **Database Migrations**
  
- Apply migrations to the database:

```bash

make migrate-dev

```

  

3.  **Running the Application**


  

```bash

make run

```
You can access the application in your web browser at http://localhost:8080