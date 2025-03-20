### Install Requirements

``` bash
go mod init github.com/Sarus1997/golang-api

go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get gorm.io/gorm
go get gorm.io/driver/mysql
```
### Check installed

``` bash
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```

### Create JWT SECRET

  ➡️ Method 1

```bash
  run in powershell

  [System.BitConverter]::ToString((1..32 | ForEach-Object {Get-Random -Minimum 0 -Maximum 256})).Replace("-", "").ToLower()
``` 

  ➡️ Method 2

```bash
  go run secret.go

```

### Example Register

```bash
{
  "username": "test1234",
  "email": "test123@gmail.com",
  "password_hash": "test@1234",
  "f_name": "Test",
  "l_name": "Last",
  "profile_picture": "www.123.com"
}
```
