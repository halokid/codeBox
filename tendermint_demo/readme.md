demo practice record
========================================

### kvstore

下面是参照官方demo的一些记录

```shell

abci-cli kvstore

# 不会定期产生空白block， 一般自己测试写demo用这种
tendermint node --consensus.create_empty_blocks=false   


# 会定期产生空白block
tendermint node   



# 储存一个key的value, 如果请求里面只是有 tx="abcd" 就是设置  key, value 的值都是 abcd
curl -s 'localhost:26657/broadcast_tx_commit?tx="abcd"'


# 获取一个key的value
curl -s 'localhost:26657/abci_query?data="abcd"'


# shell验证 base64
cho "amltbXk=" | base64 -D

# value是采用base64加密的， 可以用下面的 python2 程序来验证
"YWJjZA==".decode('base64')

# 验证的  py3 程序
import codecs; codecs.decode("YWJjZA==", 'base64').decode('ascii')


# 设置 key 又设置 value
url -s 'localhost:26657/broadcast_tx_commit?tx="name=satoshi"'

# get
curl -s 'localhost:26657/abci_query?data="name"'

# 显示所有API
http://localhost:26657/


```











