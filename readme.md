# Golang Bookshelf  Rest API
A server for storing books usings rest architecture. It currently uses a slice to imitate a server.
The server supports cors requests

## Instructions for testing server (MacOS)
1. Git clone repository 
2. Execute the restapi executable or if you have go installed on your computer you can use the the method "go run main.go"inside the folder
3. The server will run on Port 8000.

## Data Schema
1. Book 
    Title: String
    ISBN: String
    ID: String
    Description: String
    Author *Author
    
2. Author
    FirstName: String
    LastName: String


        
    