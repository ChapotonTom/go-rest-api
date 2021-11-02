# GO-REST-API
A simple REST server with authentification/authorization in Golang programming language

## About the project
  
### Description
The purpose of the rest server : manage users.  
  
A user has a unique username and a password.  
With those credentials, it can authentify to the server.  
  
Each user can have none, one or multiple of these roles : employee & manager  
Depending on its role, a user can do different thing :
- a user can see its personal information and its roles
- a user is an employee if and only if it can view a list of all other users
- a user is a manager if and only if it can create new users and assign new roles
  
### Technologies used
The main technologie used for this project is __Golang__, using the __Gin framework__ for building a RESTful API.  
In order to have a simple setup, __SQLite3__ was used for the creation & handling of the database.


## Installation
After cloning the repository, run this command in order to install all the dependencies of the project :

```
$ go get .
```

_Note : go-sqlite3 is cgo package. In order to build the app using go-sqlite3, you need __gcc__.  
However, after you have built and installed the package, you can build your app without relying on gcc in future._
  
## Run the app
After installation, simply run this command to run the application :

```
$ go run .
```
_Note : the server will run on __localhost:8080__._
  
  
# API endpoints

## Login
### Request
#### `POST /login`  
```
{ "username": "admin", "password": "admin" }
```

_Note : Since only managers can create a user, a user 'admin' is already created by default from which you can login._  
  
### Response
```
{ "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6Ikp..." }
```
_Note : This token will have to be used for every other requests._  


## Get profile
#### `GET /me`  
  
## Get Other Users
#### `GET /users`  

## Create User
#### `POST /users`  
```
{ "name": "user123", "password": "thisisapassword", "roles": ["employee", "manager"] }
```

## Update User
#### `PUT /users/:id`  
```
{ "roles": ["employee", "manager"] }
```

_Note : A manager can only change the roles of a user._  


