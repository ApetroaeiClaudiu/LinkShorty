This is an extremely simple project that I created in my journey to learning Golang.
The project's purpose is to generate random slugs for a shortened URL.

# Steps to run:

## Generate the executable:
go build -o linkshorty.exe

## Run the commands:
.\linkshorty.exe add -url https://example.com
.\linkshorty.exe list
.\linkshorty.exe get -slug {slug value}