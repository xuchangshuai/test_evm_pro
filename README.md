# Ontology支持EVM相关测试



## EIP155交易普通转账交易测试

[comment]: <> (测试case同ONG Native合约)

    case1:  测试EIP155交易普通转账ong
    case2:  测试EIP155交易普通转账ong 255
    case3:  测试EIP155交易普通转账ong 256
    case4:  测试EIP155交易普通转账ong 257
    case5:  测试EIP155交易普通转账ong 并发（ong不足两次操作消耗和发送的数量）
    case6:  测试EIP155交易普通转账ong时，nonce值错误(nonce=nonce-1)
    case7:  测试EIP155交易普通转账ong时，nonce值错误(nonce=nonc+-1)
    case8:  测试EIP155交易普通转账ong时，gaslimit值等于21000
    case9:  测试EIP155交易普通转账ong时，gaslimit值小于21000
    case10: 测试EIP155交易 往合约地址转账ong
    case11: 测试EIP155交易 往合约A地址转账ong(合约A内部往合约B转账)
测试执行步骤

    1.ontology编译后放入该项目根目录 对应钱包文件一并放入
    2.根目录执行./re.sh
    3.cd testcase 执行go test -v 
    4.检查case断言

## 交易手续费测试
1. 正常交易能够扣除用户手续费，手续费打给治理地址；
2. 交易手续费不足时，交易revert，用户的ONG扣光；
  


