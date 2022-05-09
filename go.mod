module github.com/go-kratos/kratos/v2

go 1.16

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-kratos/aegis v0.1.1
	github.com/go-playground/form/v4 v4.2.0
	github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/imdario/mergo v0.3.12
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/otel v1.0.0
	go.opentelemetry.io/otel/sdk v1.0.0
	go.opentelemetry.io/otel/trace v1.0.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20210805201207-89edb61ffb67
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace (
	github.com/go-logr/logr => github.com/go-logr/logr v1.2.3
	google.golang.org/grpc => google.golang.org/grpc v1.44.0
)
