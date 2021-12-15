#!/bin/bash

workspace=$WORKSPACE

chmod -R 777 $workspace

nodekey=$([ -z $NODE_KEY ] && echo "node.key" || $NODE_KEY)
guardiankey=$([ -z $GUARDIAN_KEY ] && echo "g.key" || $GUARDIAN_KEY)

susy keygen --desc "Guardian Key" "$workspace/$guardiankey"

PEER_ID=$(sudo susy node-keygen "$workspace/$nodekey" | tail -n 1 | awk '{ print $2 }')

echo "peer id: $PEER_ID";

chmod 666 $workspace/$nodekey
chmod 666 $workspace/$guardiankey

echo "cfg path is: $CFG_PATH";

bootstrap_list=/dns4/$NETW/udp/$RPC_PORT/quic/p2p/$PEER_ID;

if [ $BOOTSTRAP_LIST != '0' ]
then
	echo "concatenating additional peers...":
	bootstrap_list="$bootstrap_list,$BOOTSTRAP_LIST";
fi

echo "BOOTSTRAP_LIST: $BOOTSTRAP_LIST";
echo "bootstrap_list $bootstrap_list";

su wormhole -c "susy node \
	--bootstrap \
	\"$bootstrap_list\" \
	--network \
        \"$NETWORK_IDENTIFIER\" \
	--solanaContract \
	\"$SOLANA_CONTRACT\"\
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
