rm -rf Chain
# 启动链
echo 123456|./ontology --testmode > /dev/null 2>&1 &

echo 123456|./ontology asset transfer --from=1 --to=AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7 --amount=1000  --asset=ong