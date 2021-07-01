rm -rf Chain
# 启动链
sleep 3
echo 123456|./ontology --testmode > /dev/null 2>&1 &
sleep 5
# 往evm case1使用from账户转账
echo 123456|./ontology asset transfer --from=1 --to=AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7 --amount=100000  --asset=ong > /dev/null 2>&1 &
# 往evm case2使用from账户转账
sleep 5
echo 123456|./ontology asset transfer --from=1 --to=ASrpFhwYxvJMXKRPBJ6vHRL6jXFe4trnPX --amount=1  --asset=ong > /dev/null 2>&1 &
# 往evm case3使用from账户转账
sleep 5
echo 123456|./ontology asset transfer --from=1 --to=AZZJgNrvXg9qYqbjHKztmfq7V8wnuK3vVB --amount=100  --asset=ong > /dev/null 2>&1 &
# 往evm case4使用from账户转账
sleep 5
echo 123456|./ontology asset transfer --from=1 --to=AYTtfhEnJjCFAytJnAqeCo6miePjmQCKuR --amount=100  --asset=ong > /dev/null 2>&1 &