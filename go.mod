module test_evm

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.25
	github.com/ontio/ontology v1.13.2
	github.com/ontio/ontology-go-sdk v1.11.8
	github.com/smartystreets/goconvey v1.6.4
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191029031824-8986dd9e96cf
	golang.org/x/net => github.com/golang/net v0.0.0-20191028085509-fe3aa8a45271
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190412213103-97732733099d
	golang.org/x/text => github.com/golang/text v0.3.0
)
