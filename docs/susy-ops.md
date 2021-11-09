git clone https://github.com/SuSy-One/susy-v2
cd susy-v2/
make node
cd build/bin

sudo ./guardiand keygen --desc "Testnet key foo" my.key

sudo cat my.key
sudo chmod 666 my.key

```
-----BEGIN WORMHOLE GUARDIAN PRIVATE KEY-----
PublicKey: 0xB26c946F0558dc09B63198c6c83b7AE7b0dEc13e
Description: Testnet key foo

CiAqR26+GpzYSynskJ7IBDcAvgUVTv17UVfuRW8NhwsQ1w==
=J6A2
-----END WORMHOLE GUARDIAN PRIVATE KEY-----
```

создаем файл guardiand.yaml с содержимым
```
evm_watchers:
  - name: ethWatcher
    url: wss://ropsten.infura.io/ws/v3/287b06b76784416b9f230b04235de663
    contract: "0x6a475018119169DcBaC9D0efF9369318A05EC8f2"
    network_name: eth
    readiness: ethSyncing
    chain_id: 2
  - name: bscWatcher
    url: wss://arb:arb@apis.ankr.com/wss/6017ce5a3b86463c8b15279cc25d74b3/d37735e535d9d051230799cae45aeb6a/binance/full/main
    contract: "0xFC8D8166e5143559A6042BaA23174103463d5AD4"
    network_name: bsc
    readiness: bscSyncing
    chain_id: 3
```

запускаем
```
./guardiand node \
--bootstrap \
/dns4/127.0.0.1/udp/8999/quic/p2p/12D3KooWAP46H3uijFX6YcwjNQBixpyKNVp5qcP3e1eBZH13UkK7 \
--network \
/wormhole/testnet/2 \
--solanaContract \
Brdguy7BmNB4qwEbcqqMbyV5CyJd2sxQNUn6NEpMSsUb \
--adminSocket \
/home/alex/experience/wormhole/admin.socket \
--dataDir \
/home/alex/experience/wormhole/data \
--nodeName \
NodeyMcNodeface \
--nodeKey \
/home/alex/experience/wormhole/node.key \
--guardianKey \
/home/alex/experience/wormhole/my.key \
--solanaRPC \
https://api.devnet.solana.com \
--solanaWS \
wss://api.devnet.solana.com \
--statusAddr=[::]:6060 \
--publicRPC=[::]:8089 \
--publicWeb=[::]:8080 \
--config /home/alex/experience/wormhole/guardiand.yaml
```


нода при первом запуске упадет
в логе будут следующие строки
Loaded guardian key     {"address": "0xB26c946F0558dc09B63198c6c83b7AE7b0dEc13e"}
Found existing node key {"path": "/home/alex/experience/wormhole/node.key", "peerID": "12D3KooWEKf9fqxWusuQmJpC564ZGrFbM2yc51Fmqk4HALYPbFvZ"}

guardian key нужно добавить в контракт как нового валидатора
node key нужно записать в строку bootstrap вида /dns4/127.0.0.1/udp/8999/quic/p2p/12D3KooWEKf9fqxWusuQmJpC564ZGrFbM2yc51Fmqk4HALYPbFvZ 
пиры указываются через запятую.

строку пира разослать остальным валидаторам, чтобы они добавили у себя
