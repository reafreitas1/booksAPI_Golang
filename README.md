![image](https://user-images.githubusercontent.com/79333175/123786148-45751e80-d8d1-11eb-95ad-9a626d617ef9.png)

**Launch and connect the Go app with GitHub - generate: go.mod**
`go mod init github.com/reafreitas1/booksAPI_Golang`

**Imported libraries**
`go get github.com/gin-gonic/gin` (basically assists in server connection, routes, controllers and returning user response)
`go get gorm.io/gorm` (storing "temporary" instance of your query from your database)
`go get gorm.io/driver/postgres`

**Compile and generate single executable binary file**
`go build main.go`

**In a terminal:**
`docker-compose up --build` (database)

**In other terminal to execute:**
`go run main.go`

**Resources:** 
- Go version go1.16.5 linux/amd64
- Docker version 20.10.7
- NodeJs 14.16.0
- Psql (13.3 (Debian 13.3-1.pgdg100+1))
- Visual Studio Code 1.57
- Oracle VM Virtual Box 6.1






