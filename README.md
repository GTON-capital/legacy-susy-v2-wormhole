# Wormhole v2

This repository contains Certus One's reference node implementation for the Wormhole project.

See [DEVELOP.md](DEVELOP.md) for instructions on how to set up a local devnet, and
[CONTRIBUTING.md](CONTRIBUTING.md) for instructions on how to contribute to this project.

See [docs/operations.md](docs/operations.md) for node operator instructions.

![](docs/images/overview.svg)

⚠ **Wormhole v2 is in active development - see "main" branch for the v1 mainnet version** ⚠

### Audit / Feature Status

⚠ **This software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing permissions and limitations under the License.** Or plainly
spoken - this is a very complex piece of software which targets a bleeding-edge, experimental smart contract runtime.
Mistakes happen, and no matter how hard you try and whether you pay someone to audit it, it may eat your tokens, set
your printer on fire or startle your cat. Cryptocurrencies are a high-risk investment, no matter how fancy.

guardiand node --bootstrap /dns4/127.0.0.1/udp/8999/quic/p2p/12D3KooWBWkwwWbFiporBeebsW9mgkZ1ghhghX5oy1nU58mnkshw \
                --network \ 
                /wormhole/testnet/2 \ 
                --ethContract \
                0x3e2f3BC43F006b773e0a21Ef0FE61ebE02eCb6c2 \
                --solanaContract \
                Brdguy7BmNB4qwEbcqqMbyV5CyJd2sxQNUn6NEpMSsUb \
                --adminSocket \
                ~/wormhole/admin.socket \
                --dataDir \
                ~/wormhole/data \
                --nodeName \
                NodeyMcNodeface \
                --nodeKey \
                ~/wormhole/node.key \
                --guardianKey \
                /home/anaximen/wormhole/wormkey.key \
                --ethRPC \
                https://ropsten.infura.io \
                --solanaRPC \
                https://api.devnet.solana.com \
                --solanaWS \
                ws://api.devnet.solana.com \
                --statusAddr=[::]:6060