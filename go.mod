module test_evm

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.25
	github.com/google/uuid v1.1.2 // indirect
	github.com/ontio/ontology v1.14.1-alpha
	github.com/ontio/ontology-go-sdk v1.12.5
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.7.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191029031824-8986dd9e96cf
	golang.org/x/net => github.com/golang/net v0.0.0-20191028085509-fe3aa8a45271
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190412213103-97732733099d
	golang.org/x/text => github.com/golang/text v0.3.0
)
