#!/usr/bin/env bash
rm -rf Chain
# 启动链
echo 123456|./ontology --testmode > /dev/null 2>&1 &
sleep 6
echo 123456|./ontology asset transfer --from=1 --to=AMoNdAWjcrFj5t1GyAx2vKiM6BkCmHaRS7 --amount=100000  --asset=ong
sleep 6
