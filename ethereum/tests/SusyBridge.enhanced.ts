import { Implementation__factory } from "./../typechain/factories/Implementation__factory";
import { IBeacon } from "./../typechain/IBeacon.d";
import { BridgeToken__factory } from "./../typechain/factories/BridgeToken__factory";
import { assert } from "chai";

import { MockBridgeImplementation__factory } from "./../typechain/factories/MockBridgeImplementation__factory";
import { MockBridgeImplementation } from "./../typechain/MockBridgeImplementation.d";
import { TokenImplementation__factory } from "./../typechain/factories/TokenImplementation__factory";
import { BridgeImplementation__factory } from "./../typechain/factories/BridgeImplementation__factory";
import { artifacts, ethers, waffle, web3 } from "hardhat";
import { Big } from "big.js";

import { expect } from "../shared/expect";

// const BridgeSetupArtifact = artifacts.readArtifactSync("BridgeSetup");
// const SetupArtifact = artifacts.readArtifactSync("Setup");

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
import { BridgeToken } from "./../typechain/BridgeToken.d";

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

  const BridgeStructs_PayloadEnum = {
    Transfer: 1,
    AssetMeta: 2,
    GenericAction: 4,
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
      return this.guardians.map((x) => x.privateKey.slice(2));
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

  // initialize(
  //   name_: string,
  //   symbol_: string,
  //   decimals_: BigNumberish,
  //   sequence_: BigNumberish,
  //   owner_: string,
  //   chainId_: BigNumberish,
  //   nativeContract_: BytesLike,
  type TestTokenProps = {
    name: string;
    symbol: string;
    decimals: number;
    sequence: number;
    owner: string;
    chainId: number;
    nativeContract: string;
  };

  type TestProps = {
    governanceChainId: string;
    governanceContract: string;
    chainId: string;
    initSigners: string[]; // address list of guardians

    bridgeChainId: string;
    bridgeGovernanceChainId: string;
    bridgeGovernanceContract: string;
    WETH?: string;

    testToken?: TestTokenProps;

    emitterChainId: string;
  };

  const testForeignChainId = 33;

  const testProps: TestProps = {
    chainId: "0x1",
    governanceChainId: "0x2",
    governanceContract: "0x0000000000000000000000000000000000000000000000000000000000000004",

    bridgeChainId: "0x01",
    bridgeGovernanceChainId: "0x2",
    bridgeGovernanceContract: "0x0000000000000000000000000000000000000000000000000000000000000004",

    initSigners: guardianSet.addressList,

    testToken: {
      name: "TestToken",
      symbol: "TT",
      decimals: 18,
      sequence: 0,
      owner: guardianSet.addressList[0],
      chainId: testForeignChainId,
      nativeContract: "0x0",
    },

    emitterChainId: "0000",
  };

  testProps.testToken!.nativeContract = testProps.governanceContract;

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

  // суета происходит перед каждым тестом
  // this fn runs before each test (pure state persist)
  beforeEach("deploy test contracts", async () => {
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

    const setupTx = await deployedContracts.setupContract.setup(
      deployedContracts.implementationContract.address,
      guardianSet.addressList,
      testProps.chainId, // testForeignChainId,
      testProps.governanceChainId, // testGovernanceChainId,
      testProps.governanceContract // testGovernanceContract
    );

    deployedContracts.wormholeContract = (await wormholeFactory.deploy(
      deployedContracts.setupContract.address,
      setupTx.data
    )) as Wormhole;

    /**
     * Migration N2
     */
    const bridgeSetupFactory = await ethers.getContractFactory("BridgeSetup");
    const bridgeImplFactory = (await ethers.getContractFactory("BridgeImplementation")) as BridgeImplementation__factory;
    const tokenImplFactory = (await ethers.getContractFactory("TokenImplementation")) as TokenImplementation__factory;

    /** Token Migration */
    const legitOwner = guardianSet.guardians[0];
    const token = await tokenImplFactory.connect(legitOwner).deploy();
    await token
      .connect(legitOwner)
      .initialize(
        testProps.testToken!.name,
        testProps.testToken!.symbol,
        testProps.testToken!.decimals,
        testProps.testToken!.sequence,
        legitOwner.address,
        testProps.testToken!.chainId,
        testProps.testToken!.nativeContract
      );

    deployedContracts.tokenImplementation = token;

    deployedContracts.bridgeSetup = (await bridgeSetupFactory.deploy()) as BridgeSetup;
    deployedContracts.bridgeImplementation = (await bridgeImplFactory.deploy()) as BridgeImplementation;

    const bridgeSetupTx = await deployedContracts.bridgeSetup.setup(
      deployedContracts.bridgeImplementation.address,
      testProps.bridgeChainId,
      deployedContracts.wormholeContract.address,
      testProps.bridgeGovernanceChainId, // testGovernanceChainId,
      testProps.bridgeGovernanceContract, // testGovernanceContract,
      deployedContracts.tokenImplementation.address,
      deployedContracts.WETH.address
    );

    const susyTokenBridgeFactory = (await ethers.getContractFactory("SuSyTokenBridge")) as SuSyTokenBridge__factory;

    deployedContracts.susyTokenBridge = await susyTokenBridgeFactory.deploy(
      deployedContracts.bridgeSetup.address,
      bridgeSetupTx.data
    );

    deployedContracts.bridgeImplementation = await bridgeImplFactory.attach(deployedContracts.susyTokenBridge.address);

    expect(deployedContracts.wormholeContract.address).to.equal(await deployedContracts.bridgeImplementation.wormhole()); // Recommended
  });

  async function runRegisterChainTest() {
    const bridgeImplDeployed = deployedContracts.bridgeImplementation!;

    const moduleName = "TokenBridge";
    const moduleNameHex = buildOfLen(Buffer.from(moduleName, "utf8"), 32).toString("hex");

    console.log(
      "000000000000000000000000000000000000000000546f6b656e427269646765",
      web3.utils.hexToString("0x000000000000000000000000000000000000000000546f6b656e427269646765")
    );

    const dataRaw = [
      "0x",
      moduleNameHex,
      "01", // chain action
      testProps.emitterChainId, // "0000", // chain id
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

    expect(before).to.equal("0x0000000000000000000000000000000000000000000000000000000000000000");

    await bridgeImplDeployed.registerChain("0x" + vm, {
      from: guardianSet.addressList[0],
      gasLimit: 2000000,
    });

    let after = await bridgeImplDeployed.bridgeContracts(testProps.bridgeChainId);

    expect(after).to.equal(testProps.bridgeGovernanceContract);
  }

  it("should register a foreign bridge implementation correctly", async () => {
    await runRegisterChainTest();
  });

  const randomN = (n: number) => Math.ceil(Math.random() * n);
  const randomTwo = (n: number): [number, number] => {
    const a = randomN(n);
    const b = randomN(n);
    return a !== b ? [a, b] : randomTwo(n);
  };

  it("bridged tokens should only be mint- and burn-able by owner", async function () {
    const tokenFactory = (await ethers.getContractFactory("TokenImplementation")) as TokenImplementation__factory;

    const [ownerIndex, nonOwnerIndex] = randomTwo(guardianSet.n - 1);
    const legitOwner = guardianSet.guardians[ownerIndex];
    const nonOwner = guardianSet.guardians[nonOwnerIndex];
    // initialize our template token contract
    // zero init supply

    let token = await tokenFactory.connect(legitOwner).deploy();

    await token
      .connect(legitOwner)
      .initialize(
        testProps.testToken!.name,
        testProps.testToken!.symbol,
        testProps.testToken!.decimals,
        testProps.testToken!.sequence,
        legitOwner.address,
        testProps.testToken!.chainId,
        testProps.testToken!.nativeContract
      );

    await token.connect(legitOwner).mint(guardianSet.addressList[0], 10);
    await token.connect(legitOwner).burn(guardianSet.addressList[0], 5);

    let failed = false;
    try {
      await token.connect(nonOwner).mint(guardianSet.addressList[0], 10);
    } catch (err) {
      failed = true;
    }
    expect(failed).to.equal(true);

    failed = false;
    try {
      await token.connect(nonOwner).burn(guardianSet.addressList[0], 10);
    } catch (err) {
      failed = true;
    }
    expect(failed).to.equal(true);
  });

  it("should attest a token correctly", async function name() {
    const bridgeImpl = deployedContracts.bridgeImplementation!;
    const tokenImpl = deployedContracts.tokenImplementation!;

    await bridgeImpl.attestToken(tokenImpl.address, "234");
  });

  type DeployedTokenResponse = {
    mockData: TestTokenProps;
    wrappedAsset: TokenImplementation;
    tokenOwner: Wallet;
    nameWithPostfix: (x: string) => string;
  };
  async function deployWrappedToken(overrideTokenProps: Partial<TestTokenProps> = {}): Promise<DeployedTokenResponse> {
    const tokenImplFactory = (await ethers.getContractFactory("TokenImplementation")) as TokenImplementation__factory;

    const legitOwner = guardianSet.guardians[0];
    const wrappedAsset = await tokenImplFactory.connect(legitOwner).deploy();

    const mockData: TestTokenProps = {
      name: "Wrapped USDT",
      symbol: "sUSDT",
      decimals: 18,
      sequence: 0,
      owner: guardianSet.addressList[0],
      chainId: testForeignChainId,
      nativeContract: testProps.bridgeGovernanceContract,
      ...overrideTokenProps,
    };

    const nameWithPostfix = (name: string) => {
      const postfix = "(SuSy)";
      return name + " " + postfix;
    };

    await wrappedAsset
      .connect(legitOwner)
      .initialize(
        mockData.name,
        mockData.symbol,
        mockData.decimals,
        mockData.sequence,
        legitOwner.address,
        mockData.chainId,
        mockData.nativeContract
      );

    return { mockData, wrappedAsset, nameWithPostfix, tokenOwner: legitOwner };
  }

  async function runTokenAttestationTest(result: DeployedTokenResponse) {
    const { mockData, wrappedAsset, nameWithPostfix } = result;

    const bridgeImpl = deployedContracts.bridgeImplementation!;

    const web3_wrappedAssetChainId = web3.eth.abi.encodeParameter("uint16", mockData.chainId).substring(2 + (64 - 4));
    const web3_wrappedAssetAddress = buildOfLen(wrappedAsset.address, 32).toString("hex");

    const dataRaw = [
      "0x0" + BridgeStructs_PayloadEnum.AssetMeta,
      // tokenAddress
      web3_wrappedAssetAddress,
      // token chain
      web3_wrappedAssetChainId,
      // decimals
      String(mockData.decimals),
      // symbol
      buildOfLen(Buffer.from(mockData.symbol, "utf8"), 32).toString("hex"),
      // name
      buildOfLen(Buffer.from(mockData.name, "utf8"), 32).toString("hex"),
    ];

    console.log("attest token", {
      dataRaw,
    });
    const data = dataRaw.join("");

    const vm = await signAndEncodeVM(
      0,
      0,
      testProps.bridgeChainId,
      testProps.bridgeGovernanceContract,
      0,
      data,
      guardianSet.privateKeysList,
      0,
      0
    );

    console.log("attest", { web3_wrappedAssetChainId }, { web3_wrappedAssetAddress });

    const tx = await bridgeImpl.createWrapped("0x" + vm);

    const wrappedAssetRetrieved = await bridgeImpl.wrappedAsset("0x" + web3_wrappedAssetChainId, "0x" + web3_wrappedAssetAddress);

    console.log("run attestation", {
      wrappedAssetRetrieved,
      args: ["0x" + web3_wrappedAssetChainId, "0x" + web3_wrappedAssetAddress],
    });
    // assert.isNotNull(wrappedAssetRetrieved, "wrapped asset is ok");

    // const bridgeTokenFactory = await ethers.getContractFactory("TokenImplementation");
    // const bridgedToken = (await bridgeTokenFactory.attach(wrappedAssetRetrieved)) as IBeacon;
    // const initializedWrappedAsset = await wrappedAsset.attach(await bridgedToken.implementation());
    const initializedWrappedAsset = wrappedAsset;

    assert.strictEqual(await initializedWrappedAsset.symbol(), mockData.symbol, "wrapped asset symbol is ok");

    assert.strictEqual(await initializedWrappedAsset.name(), nameWithPostfix(mockData.name), "wrapped asset name is ok");

    assert.strictEqual(await initializedWrappedAsset.decimals(), mockData.decimals, "wrapped asset decimals is ok");

    assert.strictEqual(await initializedWrappedAsset.chainId(), mockData.chainId, "wrapped asset chain id is ok");

    const initializedWrappedAsset_nativeContract = await initializedWrappedAsset.nativeContract();

    assert.strictEqual(initializedWrappedAsset_nativeContract, mockData.nativeContract, "native contracts are ok");
  }

  it("should correctly deploy a wrapped asset for a token attestation", async function name() {
    await runRegisterChainTest();

    const result = await deployWrappedToken();

    await runTokenAttestationTest(result);
  });

  // it("should correctly update a wrapped asset for a token attestation", async () => {
  //   await runRegisterChainTest();

  //   const deployedTokenResponse = await deployWrappedToken();

  //   await runTokenAttestationTest(deployedTokenResponse);

  //   const { mockData, wrappedAsset, nameWithPostfix } = deployedTokenResponse;

  //   const updateToPayload: Partial<TestTokenProps> = {
  //     name: "Wrapped USDC",
  //     symbol: "sUSDC",
  //   };

  //   const web3_wrappedAssetChainId = web3.eth.abi.encodeParameter("uint16", mockData.chainId).substring(2 + (64 - 4));
  //   const web3_wrappedAssetAddress = buildOfLen(wrappedAsset.address, 32).toString("hex");

  //   console.log({ web3_wrappedAssetChainId });
  //   console.log({ web3_wrappedAssetAddress });

  //   const dataRaw = [
  //     "0x0" + BridgeStructs_PayloadEnum.AssetMeta,
  //     // tokenAddress
  //     web3_wrappedAssetAddress,
  //     // token chain
  //     web3_wrappedAssetChainId,
  //     // decimals
  //     String(mockData.decimals),
  //     // symbol
  //     buildOfLen(Buffer.from("aaaaaaaa", "utf8"), 32).toString("hex"),
  //     // name
  //     buildOfLen(Buffer.from("bbbbbbbb", "utf8"), 32).toString("hex"),
  //   ];

  //   console.log({ dataRaw });
  //   const data = dataRaw.join("");

  //   const sequence = 1;

  //   const vm = await signAndEncodeVM(
  //     0,
  //     0,
  //     testProps.bridgeChainId,
  //     testProps.bridgeGovernanceContract,
  //     sequence,
  //     data,
  //     guardianSet.privateKeysList,
  //     0,
  //     0
  //   );

  //   console.log("before", [
  //     wrappedAsset.address,
  //     await wrappedAsset.decimals(),
  //     await wrappedAsset.name(),
  //     await wrappedAsset.symbol(),
  //   ]);

  //   const bridgeImpl = deployedContracts.bridgeImplementation!;

  //   await bridgeImpl.updateWrapped("0x" + vm);

  //   console.log("after", [
  //     wrappedAsset.address,
  //     await wrappedAsset.decimals(),
  //     await wrappedAsset.name(),
  //     await wrappedAsset.symbol(),
  //   ]);

  //   const wrappedAssetRetrieved = await bridgeImpl.wrappedAsset("0x" + web3_wrappedAssetChainId, "0x" + web3_wrappedAssetAddress);

  //   console.log("after run update", {
  //     wrappedAssetRetrieved,
  //     args: ["0x" + web3_wrappedAssetChainId, "0x" + web3_wrappedAssetAddress],
  //   });

  //   // assert.isTrue(await bridgeImpl.isWrappedAsset(wrappedAssetRetrieved), "wrapped asset is ok");

  //   console.log({ wrappedAsset: wrappedAsset.address });

  //   const bridgeTokenFactory = await ethers.getContractFactory("BeaconProxy");
  //   const bridgedToken = (await bridgeTokenFactory.attach(wrappedAssetRetrieved)) as IBeacon;
  //   console.log({ implAddr: await bridgedToken.implementation() });
  //   const updatedWrappedAsset = await wrappedAsset.attach(await bridgedToken.implementation());

  //   console.log([
  //     wrappedAsset.address,
  //     updatedWrappedAsset.address,
  //     wrappedAssetRetrieved,
  //     await updatedWrappedAsset.decimals(),
  //     await updatedWrappedAsset.name(),
  //     await updatedWrappedAsset.symbol(),
  //   ]);

  //   assert.strictEqual(await updatedWrappedAsset.symbol(), updateToPayload.symbol, "wrapped asset symbol is ok");

  //   assert.strictEqual(await updatedWrappedAsset.name(), nameWithPostfix(updateToPayload.name!), "wrapped asset name is ok");

  //   const initializedWrappedAsset_nativeContract = await updatedWrappedAsset.nativeContract();

  //   assert.strictEqual(initializedWrappedAsset_nativeContract, mockData.nativeContract, "native contracts are ok");
  // });

  it("should deposit and log transfers correctly", async function () {
    await runRegisterChainTest();

    const result = await deployWrappedToken();

    await runTokenAttestationTest(result);

    let { wrappedAsset, tokenOwner, mockData } = result;

    // calls from owner
    wrappedAsset = wrappedAsset.connect(tokenOwner);

    const props = { amount: new Big(1).mul(1e18), fee: new Big(0.1).mul(1e18) };

    // mint amount
    await wrappedAsset.mint(tokenOwner.address, props.amount.toString());

    const tokenBridge = deployedContracts.bridgeImplementation!;

    await wrappedAsset.approve(tokenBridge.address, props.amount.toString());

    // deposit tokens

    const accountBalanceBefore = await wrappedAsset.balanceOf(tokenOwner.address);

    const bridgeBalanceBefore = await wrappedAsset.balanceOf(tokenBridge.address);

    assert.equal(accountBalanceBefore.toString(), props.amount.toString());
    assert.equal(bridgeBalanceBefore.toString(), "0");

    const wormhole = deployedContracts.wormholeContract!;

    let lastTransferEvent: { to: string; amount: number; from: string } | undefined | null;
    wrappedAsset.on("Transfer", (to: string, amount: number, from: string) => {
      console.log("Transfer", { to, amount, from });
      lastTransferEvent = { to, amount, from };
    });

    const wormholeImpl = Implementation__factory.connect(wormhole.address, tokenOwner);

    // event LogMessagePublished(address indexed sender, uint64 sequence, uint32 nonce, bytes payload, uint8 consistencyLevel);
    let lastLogMessageEvent:
      | { sender: string; sequence: number; nonce: number; payload: string; consistencyLevel: number }
      | undefined
      | null;
    wormholeImpl.on(
      "LogMessagePublished",
      (sender: string, sequence: number, nonce: number, payload: string, consistencyLevel: number) => {
        lastLogMessageEvent = {
          sender,
          sequence,
          nonce,
          payload,
          consistencyLevel,
        };
        console.log({ lastLogMessageEvent });
      }
    );

    // go deposit
    await tokenBridge.transferTokens(
      wrappedAsset.address,
      props.amount.toString(),
      "10",
      await wrappedAsset.nativeContract(),
      props.fee.toString(), // fee
      "234" // nonce
    );

    await new Promise((resolve) => setTimeout(resolve, 500));

    const accountBalanceAfter = await wrappedAsset.balanceOf(tokenOwner.address);

    const bridgeBalanceAfter = await wrappedAsset.balanceOf(tokenBridge.address);

    assert.equal(accountBalanceAfter.toString(), "0");

    assert.equal(bridgeBalanceAfter.toString(), props.amount.toString());

    assert.equal(lastLogMessageEvent!.sender, tokenBridge.address);

    // payload len
    assert.equal(lastLogMessageEvent!.payload.length - 2, 266);

    // payload id
    assert.equal(lastLogMessageEvent!.payload.substr(2, 2), "01");

    // amount
    assert.equal(
      lastLogMessageEvent!.payload.substr(4, 64),
      web3.eth.abi.encodeParameter("uint256", new Big(props.amount).div(1e10).toString()).substring(2)
    );

    const web3_wrappedAssetAddress = buildOfLen(wrappedAsset.address, 32).toString("hex");

    // token
    assert.equal(
      lastLogMessageEvent!.payload.substr(68, 64),
      web3_wrappedAssetAddress
      // web3.eth.abi.encodeParameter("address", testProps.WETH!).substring(2)
    );

    // chain id
    assert.equal(
      lastLogMessageEvent!.payload.substr(132, 4),
      web3.eth.abi.encodeParameter("uint16", await deployedContracts.bridgeImplementation!.chainId()).substring(2 + 64 - 4)
    );

    // to
    assert.equal("0x" + lastLogMessageEvent!.payload.substr(136, 64), await wrappedAsset.nativeContract());

    // to chain id
    assert.equal(lastLogMessageEvent!.payload.substr(200, 4), web3.eth.abi.encodeParameter("uint16", 10).substring(2 + 64 - 4));

    // fee
    assert.equal(
      lastLogMessageEvent!.payload.substr(204, 64),
      web3.eth.abi.encodeParameter("uint256", new Big(props.fee).div(1e10).toString()).substring(2)
    );
  });
});
