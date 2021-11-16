import { TokenBridge__factory } from "./../typechain/factories/TokenBridge__factory";
import { ethers, waffle, web3 } from "hardhat";
import BigNumber from "big.js";

import { expect } from "../shared/expect";

// common
import { TestERC20 } from "../typechain/TestERC20";
import { TestERC20__factory } from "./../typechain/factories/TestERC20__factory";

// app
// import { Wormhole } from "../typechain/Wormhole";
import { SuSyBridge } from "../typechain/SuSyBridge";
import { SuSyBridge__factory } from "./../typechain/factories/SuSyBridge__factory";
import { SuSyTokenBridge } from "../typechain/SuSyTokenBridge";
import { SuSyTokenBridge__factory } from "./../typechain/factories/SuSyTokenBridge__factory";

import { Setup } from "../typechain/Setup";
import { Implementation } from "../typechain/Implementation";
import { Wormhole } from "../typechain/Wormhole";
import { Wormhole__factory } from "./../typechain/factories/Wormhole__factory";

// import { ERC20 } from "../typechain/ERC20";

import { BridgeSetup } from "../typechain/BridgeSetup";
import { BridgeImplementation } from "../typechain/BridgeImplementation";
import { TokenImplementation } from "../typechain/TokenImplementation";

import { signAndEncodeVM } from "../shared/bridge";
import { Wallet } from "ethers";

describe("Tests: SuSyBridge", () => {
  let gton: TestERC20;
  // let usdc: TestERC20
  // let usdt: TestERC20
  // let weth: WrappedNative
  let startingBalance: BigNumber;
  let bridgeImplementation: BridgeImplementation;
  let tokenImplementation: TokenImplementation;

  const testGovernanceChainId = "1";
  const testChainId = testGovernanceChainId;

  const testGovernanceContract =
    "0x0000000000000000000000000000000000000000000000000000000000000004";

  const testForeignChainId = "1";
  const testForeignBridgeContract =
    "0x000000000000000000000000000000000000000000000000000000000000ffff";
  const testBridgedAssetChain = "0001";
  const testBridgedAssetAddress =
    "000000000000000000000000b7a2211e8165943192ad04f5dd21bedc29ff003e";

  class GuardianSet {
    n: number;
    guardians: Wallet[];

    constructor(n: number) {
      if (n < 1) throw new Error("invalid n");

      this.n = n;

      const accounts = waffle.provider.getWallets();
      this.guardians = accounts.slice(0, n);
    }

    get addressList(): string[] {
      return this.guardians.map((x) => x.address);
    }
  }

  const guardianSet = new GuardianSet(5);

  const deployedContracts: {
    /** Migration N0 */
    WETH?: TestERC20;

    /** Migration N1 */
    setupContract?: Setup;
    implementationContract?: Implementation;
    wormholeContract?: Wormhole;

    /** Migration N2 */
    tokenImplementation?: TokenImplementation;
    bridgeSetup?: BridgeSetup;
    bridgeImplementation?: BridgeImplementation;
    susyTokenBridge?: SuSyTokenBridge;
  } = {};

  beforeEach("deploy test contracts", async () => {
    /**
     * Migration N0
     */

    const erc20Factory = (await ethers.getContractFactory(
      "ERC20"
    )) as TestERC20__factory;
    deployedContracts.WETH = await erc20Factory.deploy(1_000_000);

    const setupFactory = await ethers.getContractFactory("Setup");
    const implFactory = await ethers.getContractFactory("Implementation");
    const wormholeFactory = (await ethers.getContractFactory(
      "Wormhole"
    )) as Wormhole__factory;

    /**
     * Migration N1
     */
    deployedContracts.setupContract = (await setupFactory.deploy()) as Setup;
    deployedContracts.implementationContract =
      (await implFactory.deploy()) as Implementation;

    const initDataEncoded_Wormhole =
      await deployedContracts.setupContract.setup(
        deployedContracts.implementationContract.address,
        guardianSet.addressList,
        testForeignChainId,
        testGovernanceChainId,
        testGovernanceContract
      );

    deployedContracts.wormholeContract = await wormholeFactory.deploy(
      deployedContracts.setupContract.address,
      initDataEncoded_Wormhole.data
    );

    /**
     * Migration N2
     */
    
    const bridgeSetupFactory = await ethers.getContractFactory("BridgeSetup");
    const bridgeImplFactory = await ethers.getContractFactory(
      "BridgeImplementation"
    );
    const tokenImplFactory = await ethers.getContractFactory(
      "TokenImplementation"
    );

    deployedContracts.tokenImplementation =
      (await tokenImplFactory.deploy()) as TokenImplementation;

    deployedContracts.bridgeSetup =
      (await bridgeSetupFactory.deploy()) as BridgeSetup;

    deployedContracts.bridgeImplementation =
      (await bridgeImplFactory.deploy()) as BridgeImplementation;

    // encode initialisation data
    const initDataEncoded_SuSyTokenBridge =
      await deployedContracts.bridgeSetup.setup(
        deployedContracts.bridgeImplementation.address,
        testChainId,
        deployedContracts.wormholeContract.address,
        testGovernanceChainId,
        testGovernanceContract,
        tokenImplementation.address,
        deployedContracts.WETH.address
      );

    const susyBridgeFactory = (await ethers.getContractFactory(
      "SuSyTokenBridge"
    )) as SuSyTokenBridge__factory;

    deployedContracts.susyTokenBridge = await susyBridgeFactory.deploy(
      deployedContracts.bridgeSetup.address,
      initDataEncoded_SuSyTokenBridge.data
    );
  });

  // it("should be initialized with the correct signers and values", async () => {
  //   const bridgeImpl = await ethers.getContractFactory("BridgeImplementation");
  //   const bridgeImplDeployed =
  //     (await bridgeImpl.deploy()) as BridgeImplementation;
  // });

  it("should register a foreign bridge implementation correctly", async () => {
    // const bridgeImpl = await ethers.getContractFactory("BridgeImplementation");
    // const bridgeImplDeployed =
    //   (await bridgeImpl.deploy()) as BridgeImplementation;
    const bridgeImplDeployed = deployedContracts.bridgeImplementation!;

    console.log({ bridge_address: bridgeImplDeployed.address });
    console.log({
      wormhole_contract: deployedContracts.wormholeContract?.address,
    });

    const accounts = waffle.provider.getWallets();
    const [testSigner1PK, testSigner2PK] = accounts;

    const moduleName = "SuSyBridge";
    const moduleBytes = Buffer.from(moduleName, "utf8");

    let moduleNameHex = Buffer.alloc(32)
      .fill(moduleBytes, 32 - moduleBytes.length)
      .toString("hex");

    const registerChainStruct = {};

    console.log({ moduleNameHex, moduleName });
    console.log(
      "000000000000000000000000000000000000000000546f6b656e427269646765",
      web3.utils.hexToString(
        "0x000000000000000000000000000000000000000000546f6b656e427269646765"
      )
      // web3.utils.hexToNumber(
      //   "0x000000000000000000000000000000000000000000546f6b656e427269646765"
      // )
    );

    const dataRaw = [
      "0x",
      // "000000000000000000000000000000000000000000546f6b656e427269646765",
      moduleNameHex,
      // Buffer.alloc(1).fill("1").toString(), // chain action
      // Buffer.alloc(2).fill(testGovernanceChainId).toString(), // chain id
      "01",
      "0000",
      web3.eth.abi
        .encodeParameter("uint16", testForeignChainId)
        .substring(2 + (64 - 4)),
      web3.eth.abi
        .encodeParameter("bytes32", testForeignBridgeContract)
        .substring(2),
    ];

    const data = dataRaw.join("");

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
    console.log({ before });
    // expect(
    //   before,
    //   "0x0000000000000000000000000000000000000000000000000000000000000000"
    // );

    await bridgeImplDeployed.registerChain("0x" + vm, {
      from: accounts[0].address,
      gasLimit: 2000000,
    });

    let after = await bridgeImplDeployed.bridgeContracts(testForeignChainId);

    console.log({ after });
    // expect(after, testForeignBridgeContract);
  });
});
