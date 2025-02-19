module github.com/achmad-dev/simple-ecommerce/gateway

go 1.22.4

require (
	github.com/achmad-dev/simple-ecommerce/pkg v0.0.0-00010101000000-000000000000
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/labstack/echo/v4 v4.12.0
	github.com/lib/pq v1.10.9
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/crypto v0.25.0
)

require golang.org/x/time v0.5.0 // indirect

replace github.com/achmad-dev/simple-ecommerce/pkg => ../pkg

require (
	github.com/joho/godotenv v1.5.1
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
)
