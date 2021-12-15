#!/bin/bash

workspace=$WORKSPACE

/usr/local/var/www/susy-v2/build/bin/guardian node \
	--bootstrap \
	/dns4/127.0.0.1/udp/8999/quic/p2p/12D3KooWAP46H3uijFX6YcwjNQBixpyKNVp5qcP3e1eBZH13UkK7 \
	--network \
	/wormhole/testnet/2 \
	--solanaContract \
	Brdguy7BmNB4qwEbcqqMbyV5CyJd2sxQNUn6NEpMSsUb \
	--adminSocket \
	$workspace/admin.socket \
	--dataDir \
	$workspace/data \
	--nodeName \
	NodeyMcNodeface \
	--nodeKey \
	$workspace/node.key \
	--guardianKey \
	$workspace/my.key \
	--solanaRPC \
	https://api.devnet.solana.com \
	--solanaWS \
	wss://api.devnet.solana.com \
	--statusAddr=[::]:6060 \
	--publicRPC=[::]:8089 \
	--publicWeb=[::]:8080 \
	--config $PWD/guardian.yaml
