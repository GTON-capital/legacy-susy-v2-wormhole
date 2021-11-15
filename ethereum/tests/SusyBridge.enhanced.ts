import { ethers, waffle, web3 } from "hardhat";
import BigNumber from "big.js";

import { expect } from "../shared/expect";

// common
import { TestERC20 } from "../typechain/TestERC20";

// app
// import { Wormhole } from "../typechain/Wormhole";
import { SuSyBridge } from "../typechain/SuSyBridge";
import { BridgeImplementation } from "../typechain/BridgeImplementation";
import { TokenImplementation } from "../typechain/TokenImplementation";
import { signAndEncodeVM } from "../shared/bridge";

describe("Tests: SuSyBridge", () => {
  let gton: TestERC20;
  // let usdc: TestERC20
  // let usdt: TestERC20
  // let weth: WrappedNative
  let startingBalance: BigNumber;
  let bridgeImplementation: BridgeImplementation;
  let tokenImplementation: TokenImplementation;

  const testChainId = "2";
  const testGovernanceChainId = "1";
  const testGovernanceContract =
    "0x0000000000000000000000000000000000000000000000000000000000000004";

  const testForeignChainId = "1";
  const testForeignBridgeContract =
    "0x000000000000000000000000000000000000000000000000000000000000ffff";
  const testBridgedAssetChain = "0001";
  const testBridgedAssetAddress =
    "000000000000000000000000b7a2211e8165943192ad04f5dd21bedc29ff003e";

  it("should be initialized with the correct signers and values", async () => {
    const bridgeImpl = await ethers.getContractFactory("BridgeImplementation");
    const bridgeImplDeployed = (await bridgeImpl.deploy()) as BridgeImplementation;

    // console.log({ bridgeImplDeployed });
  });

  it("should register a foreign bridge implementation correctly", async () => {
    const bridgeImpl = await ethers.getContractFactory("BridgeImplementation");
    const bridgeImplDeployed = (await bridgeImpl.deploy()) as BridgeImplementation;

    const accounts = waffle.provider.getWallets();
    const [testSigner1PK, testSigner2PK] = accounts;

    let data = [
      "0x",
      "000000000000000000000000000000000000000000546f6b656e427269646765",
      "01",
      "0000",
      web3.eth.abi
        .encodeParameter("uint16", testForeignChainId)
        .substring(2 + (64 - 4)),
      web3.eth.abi
        .encodeParameter("bytes32", testForeignBridgeContract)
        .substring(2),
    ].join("");

    const vm = await signAndEncodeVM(
      1,
      1,
      testGovernanceChainId,
      testGovernanceContract,
      0,
      data,
      [testSigner1PK.privateKey],
      0,
      0
    );

    let before = await bridgeImplDeployed.bridgeContracts(testForeignChainId);

    expect(
      before,
      "0x0000000000000000000000000000000000000000000000000000000000000000"
    );

    await bridgeImplDeployed.registerChain("0x" + vm, {
      from: accounts[0].address,
      gasLimit: 2000000,
    });

    let after = await bridgeImplDeployed.bridgeContracts(testForeignChainId);

    expect(after, testForeignBridgeContract);
  });
});
