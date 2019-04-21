module github.com/ubunifupay

go 1.12

require (
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v1.3.1
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ubunifupay/balance v0.0.0-20190421135647-20784d982d9f
	github.com/ubunifupay/db v0.0.0-20190421141246-a67afe691ec9
	github.com/ubunifupay/transaction v0.0.0-20190421140331-9046d432d898
	github.com/ugorji/go v1.1.4 // indirect
	google.golang.org/grpc v1.20.1
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190418165655-df01cb2cc480
	golang.org/x/net => github.com/golang/net v0.0.0-20190420063019-afa5a82059c6
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190412183630-56d357773e84
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190419153524-e8e3143a4f4a
	golang.org/x/text => github.com/golang/text v0.3.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190418145605-e7d98fc518a7
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
)
