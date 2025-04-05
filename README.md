This is an extremely simple project that I created in my journey to learning Golang.
The project's purpose is to generate random slugs for a shortened URL.

## Steps to run:


Initialize go modules
```
go mod init linkshorty
```

Create the executable:
```
go build -o linkshorty.exe

```
Run the commands:
```
.\linkshorty.exe add -url https://example.com
```
    - choose an url to create the slug for (persisting in the file)
    
```
.\linkshorty.exe list
```
    - list all the existing slugs and their respective URLs
```
.\linkshorty.exe get -slug {slug value}
```
    - get the URL of a specific slug