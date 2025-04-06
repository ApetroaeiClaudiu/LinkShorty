This is an extremely simple project that I created in my journey to learning Golang.
The project's purpose is to generate random slugs for a shortened URL.

## Steps to run:

Download Go from their website:
```
https://go.dev/dl/
```

Clone the Repository:
```
git clone https://github.com/ApetroaeiClaudiu/LinkShorty.git
```

Enter the project directory
```
cd LinkShorty
```

Initialize go modules
```
go mod init linkshorty
```

Create the executable:
```
go build -o linkshorty.exe
```

Run the server:
```
go run main.go storage.go
```


Test the app:
```
Go to http://localhost:8080
```

Use cases:
```
Enter a URL to shorten it
```
```
View the list of all shortened urls and be able to directly access them.
```
```
Validation for invalid URLs and duplicate URLs
```
