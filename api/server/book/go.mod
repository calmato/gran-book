module github.com/calmato/gran-book/api/server/book

go 1.15

require (
	cloud.google.com/go/storage v1.14.0
	firebase.google.com/go/v4 v4.2.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/prometheus/client_golang v1.9.0
	go.uber.org/zap v1.16.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/api v0.41.0
	google.golang.org/genproto v0.0.0-20210310155132-4ce2db91004e
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.3
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)
