# 查合约详情
import json
import sys

import requests


def getTransferInfo(hash, iurl):
    jsonData = {
        "jsonrpc": "2.0",
        "method": "eth_getTransactionReceipt",
        "params": [hash],
        "id": 1
    }
    url = "http://127.0.0.1:20339"
    if iurl is not None:
        url = iurl
    response = requests.post(url=url, json=jsonData)
    print(response.status_code)
    res = json.loads(response.text)
    print(res)


if __name__ == "__main__":
    iurl = None
    hash = ""
    try:
        hash = sys.argv[1]
        iurl = sys.argv[2]
    except Exception as e:
        print(e.args[0])
    getTransferInfo(hash, iurl)
