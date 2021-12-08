import {
  CHAIN_ID_POLYGON,
  CHAIN_ID_ETH,
  CHAIN_ID_SOLANA,
  attestFromEth,
  parseSequenceFromLogEth,
  getEmitterAddressEth,
  getSignedVAA,
  redeemOnEth,
} from "@certusone/wormhole-sdk";
import { Wallet } from "ethers";
import { waffle } from "hardhat";

describe("susy bridge testing in mainnet", () => {
  it("evm to evm", async () => {
    console.log({
      CHAIN_ID_POLYGON,
      CHAIN_ID_ETH,
      CHAIN_ID_SOLANA,
    });

    const ETH_BRIDGE_ADDRESS = "";
    const ETH_TOKEN_BRIDGE_ADDRESS = "";
    const WORMHOLE_RPC_HOST = "";
    const tokenAddress = "";

    const signer = Wallet.createRandom();

    const receipt = await attestFromEth(ETH_TOKEN_BRIDGE_ADDRESS, signer, tokenAddress);
    // Get the sequence number and emitter address required to fetch the signedVAA of our message
    const sequence = parseSequenceFromLogEth(receipt, ETH_BRIDGE_ADDRESS);
    const emitterAddress = getEmitterAddressEth(ETH_TOKEN_BRIDGE_ADDRESS);
    // Fetch the signedVAA from the Wormhole Network (this may require retries while you wait for confirmation)
    const { signedVAA } = await getSignedVAA(WORMHOLE_RPC_HOST, CHAIN_ID_ETH, emitterAddress, sequence);

    await redeemOnEth(ETH_TOKEN_BRIDGE_ADDRESS, signer, signedVAA);
  });
});
