import { TokenImplementation__factory } from "./../typechain/factories/TokenImplementation__factory";
import { BridgeImplementation__factory } from "./../typechain/factories/BridgeImplementation__factory";
import { artifacts, ethers, waffle, web3 } from "hardhat";
import BigNumber from "big.js";

import { expect } from "../shared/expect";

const BridgeSetupArtifact = artifacts.readArtifactSync("BridgeSetup");
const SetupArtifact = artifacts.readArtifactSync("Setup");

// common
import { TestERC20 } from "../typechain/TestERC20";
import { TestERC20__factory } from "./../typechain/factories/TestERC20__factory";

import { BridgeGovernance } from "../typechain/BridgeGovernance";
import { BridgeGovernance__factory } from "./../typechain/factories/BridgeGovernance__factory";

// app
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

  // const testGovernanceChainId = "1";
  // const testChainId = testGovernanceChainId;

  // const testGovernanceContract =
  //   "0x0000000000000000000000000000000000000000000000000000000000000004";

  // const testForeignChainId = "1";
  // const testForeignBridgeContract =
  //   "0x000000000000000000000000000000000000000000000000000000000000ffff";
  // const testBridgedAssetChain = "0001";
  // const testBridgedAssetAddress =
  //   "000000000000000000000000b7a2211e8165943192ad04f5dd21bedc29ff003e";

  type TestProps = {
    governanceChainId: string;
    governanceContract: string;
    chainId: string;
    initSigners: string[]; // address list of guardians

    bridgeChainId: string;
    bridgeGovernanceChainId: string;
    bridgeGovernanceContract: string;
    WETH?: string;
  };

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

    get privateKeysList(): string[] {
      return this.guardians.map((x) => x.privateKey);
    }
  }

  const guardianSet = new GuardianSet(5);

  // # Wormhole Core Migrations
  // INIT_SIGNERS=["0xbeFA429d57cD18b7F8A4d91A2da9AB4AF05d0FBe"]
  // INIT_CHAIN_ID=0x2
  // INIT_GOV_CHAIN_ID=0x1
  // INIT_GOV_CONTRACT=0x0000000000000000000000000000000000000000000000000000000000000004

  // # Bridge Migrations
  // BRIDGE_INIT_CHAIN_ID=0x02
  // BRIDGE_INIT_GOV_CHAIN_ID=0x1
  // BRIDGE_INIT_GOV_CONTRACT=0x0000000000000000000000000000000000000000000000000000000000000004
  // BRIDGE_INIT_WETH=0xDDb64fE46a91D46ee29420539FC25FD07c5FEa3E

  const testProps: TestProps = {
    chainId: "0x2",
    governanceChainId: "0x1",
    governanceContract: "0x0000000000000000000000000000000000000000000000000000000000000004",

    bridgeChainId: "0x02",
    bridgeGovernanceChainId: "0x1",
    bridgeGovernanceContract: "0x0000000000000000000000000000000000000000000000000000000000000004",

    initSigners: guardianSet.addressList,
  };

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
    susyBridge?: SuSyBridge;

    bridgeAsset?: TestERC20;
  } = {};

  function buildOfLen(_input: Buffer | string, n: number): Buffer {
    const input = typeof _input === "string" ? Buffer.from(web3.utils.hexToBytes(_input)) : _input;
    return Buffer.alloc(n).fill(input, n - input.length);
  }

  beforeEach("deploy test contracts", async () => {
    // суета

    // const vanilaGovernanceFactory = (await ethers.getContractFactory("BridgeGovernance")) as BridgeGovernance__factory;

    // const governanceContract = await vanilaGovernanceFactory.deploy();
    // const bridgeGovernanceContract = await vanilaGovernanceFactory.deploy();

    // testProps.governanceContract = governanceContract.address;
    // testProps.bridgeGovernanceContract = bridgeGovernanceContract.address;

    // testProps.governanceContract = "0x" + buildOfLen(governanceContract.address, 32).toString("hex");
    // testProps.bridgeGovernanceContract = "0x" + buildOfLen(bridgeGovernanceContract.address, 32).toString("hex");

    /**
     * Migration N0
     */

    const erc20Factory = (await ethers.getContractFactory("TestERC20")) as TestERC20__factory;
    deployedContracts.WETH = await erc20Factory.deploy(1_000_000);

    const setupFactory = await ethers.getContractFactory("Setup");
    const implFactory = await ethers.getContractFactory("Implementation");
    const wormholeFactory = (await ethers.getContractFactory("Wormhole")) as Wormhole__factory;

    // in order to override

    /**
     * Migration N1
     */
    deployedContracts.setupContract = (await setupFactory.deploy()) as Setup;
    deployedContracts.implementationContract = (await implFactory.deploy()) as Implementation;

    // const initDataEncoded_Wormhole = await deployedContracts.setupContract.setup(
    //   deployedContracts.implementationContract.address,
    //   guardianSet.addressList,
    //   testProps.chainId, // testForeignChainId,
    //   testProps.governanceChainId, // testGovernanceChainId,
    //   testProps.governanceContract // testGovernanceContract
    // );
    const setup = new web3.eth.Contract(SetupArtifact.abi, deployedContracts.setupContract.address);

    const initDataEncoded_Wormhole = setup.methods
      .setup(
        deployedContracts.implementationContract.address,
        guardianSet.addressList,
        testProps.chainId, // testForeignChainId,
        testProps.governanceChainId, // testGovernanceChainId,
        testProps.governanceContract // testGovernanceContract
      )
      .encodeABI();

    console.log({ initDataEncoded_Wormhole });

    deployedContracts.wormholeContract = (await wormholeFactory.deploy(
      deployedContracts.setupContract.address,
      initDataEncoded_Wormhole
    )) as Wormhole;

    /**
     * Migration N2
     */
    const bridgeSetupFactory = await ethers.getContractFactory("BridgeSetup");
    const bridgeImplFactory = (await ethers.getContractFactory("BridgeImplementation")) as BridgeImplementation__factory;
    const tokenImplFactory = (await ethers.getContractFactory("TokenImplementation")) as TokenImplementation__factory;

    deployedContracts.tokenImplementation = (await tokenImplFactory.deploy()) as TokenImplementation;

    deployedContracts.bridgeSetup = (await bridgeSetupFactory.deploy()) as BridgeSetup;
    deployedContracts.bridgeImplementation = (await bridgeImplFactory.deploy()) as BridgeImplementation;

    // encode initialisation data
    // const initDataEncoded_SuSyTokenBridge = await deployedContracts.bridgeSetup.setup(
    //   deployedContracts.bridgeImplementation.address,
    //   testProps.bridgeChainId,
    //   deployedContracts.wormholeContract.address,
    //   testProps.bridgeChainId, // testGovernanceChainId,
    //   testProps.bridgeGovernanceContract, // testGovernanceContract,
    //   deployedContracts.tokenImplementation.address,
    //   deployedContracts.WETH.address
    // );

    console.log({
      args: [
        deployedContracts.bridgeImplementation.address,
        testProps.chainId,
        deployedContracts.wormholeContract.address,
        testProps.bridgeChainId, // testGovernanceChainId,
        testProps.bridgeGovernanceContract, // testGovernanceContract,
        deployedContracts.tokenImplementation.address,
        deployedContracts.WETH.address,
      ],
    });

    const setupWeb3 = new web3.eth.Contract(BridgeSetupArtifact.abi, deployedContracts.bridgeSetup.address);
    const initDataEncoded_SuSyTokenBridge_2 = setupWeb3.methods
      .setup(
        deployedContracts.bridgeImplementation.address,
        testProps.chainId,
        deployedContracts.wormholeContract.address,
        testProps.bridgeChainId, // testGovernanceChainId,
        testProps.bridgeGovernanceContract, // testGovernanceContract,
        deployedContracts.tokenImplementation.address,
        deployedContracts.WETH.address
      )
      .encodeABI();

    console.log({ initDataEncoded_SuSyTokenBridge_2 });
    const susyTokenBridgeFactory = (await ethers.getContractFactory("SuSyTokenBridge")) as SuSyTokenBridge__factory;

    deployedContracts.susyTokenBridge = await susyTokenBridgeFactory.deploy(
      deployedContracts.bridgeSetup.address,
      initDataEncoded_SuSyTokenBridge_2
    );

    console.log("wh: provided addr", { wormhole: deployedContracts.wormholeContract.address });

    // console.log("wh: got addr", { wormhole: await deployedContracts.susyTokenBridge!.wormhole() });
    // console.log("impl", await deployedContracts.susyTokenBridge.resolvedAddress);
  });

  it("should register a foreign bridge implementation correctly", async () => {
    const bridgeImplDeployed = deployedContracts.bridgeImplementation!;

    console.log({ bridge_address: bridgeImplDeployed.address });
    console.log({
      wormhole_contract: deployedContracts.wormholeContract?.address,
    });

    const moduleName = "TokenBridge";
    const moduleBytes = Buffer.from(moduleName, "utf8");

    let moduleNameHex = Buffer.alloc(32)
      .fill(moduleBytes, 32 - moduleBytes.length)
      .toString("hex");

    const registerChainStruct = {};

    console.log({ moduleNameHex, moduleName });
    console.log(
      "000000000000000000000000000000000000000000546f6b656e427269646765",
      web3.utils.hexToString("0x000000000000000000000000000000000000000000546f6b656e427269646765")
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
      web3.eth.abi.encodeParameter("uint16", testProps.bridgeChainId).substring(2 + (64 - 4)),
      web3.eth.abi.encodeParameter("bytes32", testProps.bridgeGovernanceContract).substring(2),
    ];

    const data = dataRaw.join("");

    const vm = await signAndEncodeVM(
      1,
      1,
      testProps.governanceChainId, // testGovernanceChainId,
      testProps.governanceContract, // testGovernanceContract,
      0,
      data,
      guardianSet.privateKeysList,
      0,
      0
    );

    let before = await bridgeImplDeployed.bridgeContracts(testProps.bridgeChainId);
    console.log({ before });
    expect(before, "0x0000000000000000000000000000000000000000000000000000000000000000");

    console.log({ wormhole: await bridgeImplDeployed.wormhole() });
    // console.log({ wormhole: deployedContracts.bridgeSetup? });
    // await bridgeImplDeployed.registerChain("0x" + vm, {
    //   from: guardianSet.addressList[0],
    //   gasLimit: 2000000,
    // });

    let after = await bridgeImplDeployed.bridgeContracts(testProps.bridgeChainId);

    console.log({ after });
    expect(after, testProps.bridgeGovernanceContract);
  });
});
