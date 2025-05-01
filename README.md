# Socionics app using Golang

This is a simple app to show you what type of character you are!
We are using Golang and PostgreSQL.

To create folders:

```bash
    md "sub\sub dir\other dir"
```

To create multiple **embedded** folders:
```bash
    "entity", "usecases", "repositories", "handlers", "services", "app" | ForEach-Object { md -Force "internal/$_" }
```

To create a file:
```bash
    ni internal/entities/user.go -Type File
```

## Step 1: Create entities (for user and character)

Work with the directory "...\gocionics\internal\entities"

Create there 2 files, **user.go** and **character.go**, and put there your structures.

- What is an "entity" in this project?

- What are the "DTOs" of this project?

## Step 2: Create Use Case
We create a User structure and methods with panic("implement me")

## Step 3: Create Repository
Here we create exact interface to work with our User structure. 
So the structure does exactly the thing it's asked for, realising the interface.

## Step 4: Create a Controller
Create a controller for auth.

## Step 5: Database realisation
In database.go create your db implementation

## Step 6: Connect all layers
In main.go we connect all our layers 

## Step 7: Added Readmes for all folders to show what's there

## Step 8: After working with entities -> repos:
We code a constructor for our usecases

## Step 9: We realise the usecases
Set registration + set character

## Step 10: Handlers
Write code logic for methods and CRUDs