module github.com/MAKOSCAFEE/malengo-pay

go 1.12

require (
	github.com/MAKOSCAFEE/malengo-pay/balance v0.0.0-20190421104451-a23831b15d53
	github.com/MAKOSCAFEE/malengo-pay/transaction v0.0.0-20190421104451-a23831b15d53
	github.com/gin-gonic/gin v1.3.0
	github.com/gocql/gocql v0.0.0-20190418090649-59a610c947c5
	github.com/golang/protobuf v1.3.1
	github.com/scylladb/gocqlx v1.2.2
	google.golang.org/grpc v1.20.1
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
