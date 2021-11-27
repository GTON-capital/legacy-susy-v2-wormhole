#!/bin/bash

workspace=$WORKSPACE

chmod -R 777 $workspace

nodekey=$([ -z $NODE_KEY ] && echo "node.key" || $NODE_KEY)
guardiankey=$([ -z $GUARDIAN_KEY ] && echo "g.key" || $GUARDIAN_KEY)

susy keygen --desc "Guardian Key" "$workspace/$guardiankey"

PEER_ID=$(sudo susy node-keygen "$workspace/$nodekey" | tail -n 1 | awk '{ print $2 }')

echo $PEER_ID

chmod 666 $workspace/$nodekey
chmod 666 $workspace/$guardiankey

echo $CFG_PATH

su wormhole -c "susy node \
	--bootstrap \
	/dns4/$NETW/udp/$RPC_PORT/quic/p2p/$PEER_ID \
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
	$workspace/$nodekey \
	--guardianKey \
	$workspace/$guardiankey \
	--solanaRPC \
	https://api.devnet.solana.com \
	--solanaWS \
	wss://api.devnet.solana.com \
	--statusAddr=[::]:6060 \
	--publicRPC=[::]:8089 \
	--publicWeb=[::]:8080 \
	--config $CFG_PATH"
