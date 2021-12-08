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
import { BigNumber, providers, Wallet } from "ethers";
import { BridgeToken } from "./../typechain/BridgeToken.d";

describe("Tests: SuSyBridge", () => {
  // let gton: TestERC20;
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

  class VMArgsEncoder {
    static instance = new VMArgsEncoder();

    completeTransferMethodEncode(
      amount: Big,
      wrappedAssetAddr: string,
      wrappedAssetChainId: number,
      receiverAddr: string,
      receiverChainId: number,
      fee: string
    ): string[] {
      return [
        "0x0" + BridgeStructs_PayloadEnum.Transfer,
        // amount
        web3.eth.abi.encodeParameter("uint256", amount.div(1e10).toFixed()).substring(2),
        // tokenaddress
        web3.eth.abi.encodeParameter("address", wrappedAssetAddr).substr(2),
        // tokenchain
        web3.eth.abi.encodeParameter("uint16", wrappedAssetChainId).substring(2 + (64 - 4)),
        // receiver
        web3.eth.abi.encodeParameter("address", receiverAddr).substr(2),
        // receiving chain
        web3.eth.abi.encodeParameter("uint16", receiverChainId).substring(2 + (64 - 4)),
        // fee
        // caseProps.fee ?? "0000000000000000000000000000000000000000000000000000000000000000",
        buildOfLen(fee, 32).toString("hex"),
      ];
    }
  }

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

  // foreign chain is chain id of wrapped asset
  const testForeignChainId = 33;

  const testProps: TestProps = {
    chainId: "0x1",
    governanceChainId: "0x1",
    governanceContract: "0x0000000000000000000000000000000000000000000000000000000000000004",

    bridgeChainId: "0x1",
    bridgeGovernanceChainId: "0x1",
    bridgeGovernanceContract: "0x0000000000000000000000000000000000000000000000000000000000000004",

    emitterChainId: "0000",

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
  };

  testProps.testToken!.nativeContract = testProps.governanceContract;

  type DeployedContracts = {
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
  };
  const deployedContracts: DeployedContracts = {};

  function buildOfLen(_input: Buffer | string, n: number): Buffer {
    const input = typeof _input === "string" ? Buffer.from(web3.utils.hexToBytes(_input)) : _input;
    return Buffer.alloc(n).fill(input, n - input.length);
  }

  async function runInitialMigration(_deployedContracts: DeployedContracts, _props: Partial<TestProps> = {}) {
    const props = {
      ...testProps,
      ..._props,
    };

    // console.log({ _props, testProps, props });
    /**
     * Migration N0
     */

    const erc20Factory = (await ethers.getContractFactory("TestERC20")) as TestERC20__factory;
    _deployedContracts.WETH = await erc20Factory.deploy(1_000_000);

    const setupFactory = await ethers.getContractFactory("Setup");
    const implFactory = await ethers.getContractFactory("Implementation");
    const wormholeFactory = (await ethers.getContractFactory("Wormhole")) as Wormhole__factory;

    // in order to override

    /**
     * Migration N1
     */
    _deployedContracts.setupContract = (await setupFactory.deploy()) as Setup;
    _deployedContracts.implementationContract = (await implFactory.deploy()) as Implementation;

    const setupTx = await _deployedContracts.setupContract.setup(
      _deployedContracts.implementationContract.address,
      guardianSet.addressList,
      props.chainId, // testForeignChainId,
      props.governanceChainId, // testGovernanceChainId,
      props.governanceContract // testGovernanceContract
    );

    _deployedContracts.wormholeContract = (await wormholeFactory.deploy(
      _deployedContracts.setupContract.address,
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
        props.testToken!.name,
        props.testToken!.symbol,
        props.testToken!.decimals,
        props.testToken!.sequence,
        legitOwner.address,
        props.testToken!.chainId,
        props.testToken!.nativeContract
      );

    _deployedContracts.tokenImplementation = token;

    _deployedContracts.bridgeSetup = (await bridgeSetupFactory.deploy()) as BridgeSetup;
    _deployedContracts.bridgeImplementation = (await bridgeImplFactory.deploy()) as BridgeImplementation;

    const bridgeSetupTx = await _deployedContracts.bridgeSetup.setup(
      _deployedContracts.bridgeImplementation.address,
      props.bridgeChainId,
      _deployedContracts.wormholeContract.address,
      props.bridgeGovernanceChainId, // testGovernanceChainId,
      props.bridgeGovernanceContract, // testGovernanceContract,
      _deployedContracts.tokenImplementation.address,
      _deployedContracts.WETH.address
    );

    const susyTokenBridgeFactory = (await ethers.getContractFactory("SuSyTokenBridge")) as SuSyTokenBridge__factory;

    _deployedContracts.susyTokenBridge = await susyTokenBridgeFactory.deploy(
      _deployedContracts.bridgeSetup.address,
      bridgeSetupTx.data
    );

    _deployedContracts.bridgeImplementation = await bridgeImplFactory.attach(_deployedContracts.susyTokenBridge!.address);

    expect(_deployedContracts.wormholeContract.address).to.equal(await _deployedContracts.bridgeImplementation.wormhole()); // Recommended
  }

  // суета происходит перед каждым тестом
  // this fn runs before each test (pure state persist)
  beforeEach("deploy test contracts", async () => {
    await runInitialMigration(deployedContracts);
  });

  async function runRegisterChainTest(_deployedContracts: DeployedContracts, overrideProps: Partial<TestProps> = {}) {
    const bridgeImplDeployed = _deployedContracts.bridgeImplementation!;

    const moduleName = "SuSyBridge";
    const moduleNameHex = buildOfLen(Buffer.from(moduleName, "utf8"), 32).toString("hex");

    const props = {
      ...testProps,
      ...overrideProps,
    };

    // console.log(
    //   "000000000000000000000000000000000000000000546f6b656e427269646765",
    //   web3.utils.hexToString("0x000000000000000000000000000000000000000000546f6b656e427269646765")
    // );

    const dataRaw = [
      "0x",
      moduleNameHex,
      "01", // chain action
      props.emitterChainId, // "0000", // chain id
      web3.eth.abi.encodeParameter("uint16", props.bridgeChainId).substring(2 + (64 - 4)),
      web3.eth.abi.encodeParameter("bytes32", props.bridgeGovernanceContract).substring(2),
    ];

    const data = dataRaw.join("");

    const vm = await signAndEncodeVM(
      1,
      1,
      props.governanceChainId, // testGovernanceChainId,
      props.governanceContract, // testGovernanceContract,
      0,
      data,
      guardianSet.privateKeysList,
      0,
      0
    );

    let before = await bridgeImplDeployed.bridgeContracts(props.bridgeChainId);

    expect(before).to.equal("0x0000000000000000000000000000000000000000000000000000000000000000");

    await bridgeImplDeployed.registerChain("0x" + vm, {
      from: guardianSet.addressList[0],
      gasLimit: 2000000,
    });

    let after = await bridgeImplDeployed.bridgeContracts(props.bridgeChainId);

    expect(after).to.equal(props.bridgeGovernanceContract);
  }

  const randomN = (n: number) => Math.ceil(Math.random() * n);
  const randomTwo = (n: number): [number, number] => {
    const a = randomN(n);
    const b = randomN(n);
    return a !== b ? [a, b] : randomTwo(n);
  };

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
    // console.log({ mockData });
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

  async function runTokenAttestationTest(
    _deployedContracts: DeployedContracts,
    result: DeployedTokenResponse,
    props: Partial<TestProps>
  ) {
    const { mockData, wrappedAsset, nameWithPostfix } = result;

    // console.log({ result });

    const bridgeImpl = _deployedContracts.bridgeImplementation!;

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

    // console.log("attest token", {
    //   dataRaw,
    //   dataDict: {
    //     act: "0x0" + BridgeStructs_PayloadEnum.AssetMeta,
    //     // tokenAddress
    //     tokenAddr: web3_wrappedAssetAddress,
    //     // token chain
    //     tokenChain: web3_wrappedAssetChainId,
    //     // decimals
    //     decimals: String(mockData.decimals),
    //     // symbol
    //     symbol: buildOfLen(Buffer.from(mockData.symbol, "utf8"), 32).toString("hex"),
    //     // name
    //     name: buildOfLen(Buffer.from(mockData.name, "utf8"), 32).toString("hex"),
    //   },
    //   web3_wrappedAssetChainId,
    //   web3_wrappedAssetAddress,
    // });
    const data = dataRaw.join("");

    const vm = await signAndEncodeVM(
      0,
      0,
      props.bridgeChainId!,
      props.bridgeGovernanceContract!,
      0,
      data,
      guardianSet.privateKeysList,
      0,
      0
    );

    // const bridge = de

    // console.log("attest", { result, web3_wrappedAssetChainId }, { web3_wrappedAssetAddress });

    const tx = await bridgeImpl.createWrapped("0x" + vm);

    const wrappedAssetRetrieved = await bridgeImpl.wrappedAsset("0x" + web3_wrappedAssetChainId, "0x" + web3_wrappedAssetAddress);

    // console.log("run attestation", {
    //   wrappedAssetRetrieved,
    //   args: ["0x" + web3_wrappedAssetChainId, "0x" + web3_wrappedAssetAddress],
    // });
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

  async function runDepositFundsTest(result: DeployedTokenResponse) {
    let { wrappedAsset, tokenOwner } = result;

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
      // console.log("Transfer", { to, amount, from });
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
        // console.log({ lastLogMessageEvent });
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

    await new Promise((resolve) => setTimeout(resolve, 200));

    const accountBalanceAfter = await wrappedAsset.balanceOf(tokenOwner.address);

    const bridgeBalanceAfter = await wrappedAsset.balanceOf(tokenBridge.address);

    assert.equal(accountBalanceAfter.toString(), "0");

    assert.equal(bridgeBalanceAfter.toString(), props.amount.toString());

    // console.log({ lastLogMessageEvent });
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
  }

  it("should register a foreign bridge implementation correctly", async () => {
    await runRegisterChainTest(deployedContracts);
  });

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

  it("should correctly deploy a wrapped asset for a token attestation", async function name() {
    await runRegisterChainTest(deployedContracts);

    const result = await deployWrappedToken();

    await runTokenAttestationTest(deployedContracts, result, testProps);
  });

  it("should deposit and log transfers correctly", async function () {
    await runRegisterChainTest(deployedContracts);

    const result = await deployWrappedToken();

    await runTokenAttestationTest(deployedContracts, result, testProps);
    await runDepositFundsTest(result);
  });

  async function runMintCompleteTransferTest(
    _deployedContracts: DeployedContracts,
    result: DeployedTokenResponse,
    props: Partial<TestProps>,
    caseProps: {
      amount: Big;
      receiverAddress: string;
      fee?: string;
    }
  ) {
    let { wrappedAsset, tokenOwner } = result;

    wrappedAsset = wrappedAsset.connect(tokenOwner);

    const tokenBridge = _deployedContracts.bridgeImplementation!;

    await wrappedAsset.mint(tokenBridge.address, caseProps.amount.toFixed());

    const accountBalanceBefore = await wrappedAsset.balanceOf(tokenOwner.address);
    const bridgeBalanceBefore = await wrappedAsset.balanceOf(tokenBridge.address);

    assert.equal(accountBalanceBefore.toString(), "0");
    assert.equal(bridgeBalanceBefore.toString(), caseProps.amount.toFixed());

    // const dataRaw = [
    //   "0x0" + BridgeStructs_PayloadEnum.Transfer,
    //   // amount
    //   web3.eth.abi.encodeParameter("uint256", caseProps.amount.div(1e10).toString()).substring(2),
    //   // tokenaddress
    //   web3.eth.abi.encodeParameter("address", wrappedAsset.address).substr(2),
    //   // tokenchain
    //   web3.eth.abi.encodeParameter("uint16", await wrappedAsset.chainId()).substring(2 + (64 - 4)),
    //   // receiver
    //   web3.eth.abi.encodeParameter("address", caseProps.receiverAddress).substr(2),
    //   // receiving chain
    //   web3.eth.abi.encodeParameter("uint16", await tokenBridge.chainId()).substring(2 + (64 - 4)),
    //   // fee
    //   caseProps.fee ?? "0000000000000000000000000000000000000000000000000000000000000000",
    // ];

    // console.log({ dataRaw });
    // const data = dataRaw.join("");
    const data = VMArgsEncoder.instance
      .completeTransferMethodEncode(
        caseProps.amount,
        wrappedAsset.address,
        await wrappedAsset.chainId(),
        caseProps.receiverAddress,
        await tokenBridge.chainId(),
        "0x0"
      )
      .join("");

    const vm = await signAndEncodeVM(
      0,
      0,
      props.bridgeChainId!,
      props.bridgeGovernanceContract!,
      0,
      data,
      guardianSet.privateKeysList,
      0,
      0
    );

    await tokenBridge.completeTransfer("0x" + vm);

    const accountBalanceAfter = await wrappedAsset.balanceOf(tokenOwner.address);

    const bridgeBalanceAfter = await wrappedAsset.balanceOf(tokenBridge.address);

    assert.equal(accountBalanceAfter.toString(), "0");

    assert.equal(bridgeBalanceAfter.toString(), caseProps.amount.toFixed());
  }

  // this test emulates destination operation
  // like, some funds were locked on other chain
  // and this test validates correct behaviour on
  // destination chain -> mint operation
  it("test behaviour as destination chain (mint): (should transfer out locked assets for a valid transfer vm)", async () => {
    const receiver = Wallet.createRandom();

    await runRegisterChainTest(deployedContracts);

    const caseProps = { amount: new Big(10).mul(1e18), receiverChain: 137, receiverAddress: receiver.address };

    const result = await deployWrappedToken({ chainId: caseProps.receiverChain });

    await runTokenAttestationTest(deployedContracts, result, testProps);

    await runMintCompleteTransferTest(deployedContracts, result, testProps, caseProps);
  });

  it("should revert on transfer out of a total of > max(uint64) tokens", async () => {
    await runRegisterChainTest(deployedContracts);

    const plain = {
      supply: "184467440737095516160000000000",
      firstTransfer: "1000000000000",
    };
    const caseProps = {
      supply: new Big(plain.supply),
      firstTransfer: new Big(plain.firstTransfer),
      receiverChain: 135,
    };

    const result = await deployWrappedToken({ chainId: caseProps.receiverChain });

    await runTokenAttestationTest(deployedContracts, result, testProps);

    let { wrappedAsset, tokenOwner } = result;

    wrappedAsset = wrappedAsset.connect(tokenOwner);

    const bridgeImpl = deployedContracts.bridgeImplementation!.connect(tokenOwner);

    await wrappedAsset.mint(tokenOwner.address, caseProps.supply.toFixed());
    await wrappedAsset.approve(bridgeImpl.address, caseProps.supply.toFixed());

    await bridgeImpl.transferTokens(
      wrappedAsset.address,
      caseProps.firstTransfer.toFixed(),
      "10",
      await wrappedAsset.nativeContract(),
      "0",
      "0"
    );

    let failed = false;
    try {
      await bridgeImpl.transferTokens(
        wrappedAsset.address,
        caseProps.supply.sub(caseProps.firstTransfer).toFixed(),
        "10",
        await wrappedAsset.nativeContract(),
        "0",
        "0"
      );
    } catch (err) {
      assert.equal(
        // @ts-ignore
        err.message,
        "VM Exception while processing transaction: reverted with reason string 'transfer exceeds max outstanding bridged token amount'"
      );
      failed = true;
    }

    assert.isTrue(failed);
  });

  it("rotate validators: add new, drop existing", async () => {
    const newGuardianSet = new GuardianSet(5);

    await runRegisterChainTest(deployedContracts);

    const result = await deployWrappedToken();

    await runTokenAttestationTest(deployedContracts, result, testProps);
  });

  it("mesh network transfer test (imitate chains)", async () => {
    type ChainCfg = {
      name: string;
      chainId: number;
      implSlot: string;
      deployment: DeployedContracts;
      tokenDeployment?: DeployedTokenResponse;
      wrapAssetFromChainIds: number[];
    };

    const tokenDeployFn = async (
      cfg: ChainCfg,
      overrideToken: Partial<TestTokenProps>,
      props: Partial<TestProps>
    ): Promise<DeployedTokenResponse> => {
      const result = await deployWrappedToken({
        ...overrideToken,
        nativeContract: cfg.implSlot,
      });

      await runTokenAttestationTest(cfg.deployment, result, props);
      return result;
    };

    const origin = 333;
    const tokenConfig: Partial<TestTokenProps> = {
      name: "Super-Summetry Token",
      symbol: "SUSY",
      decimals: 18,
      sequence: 0,
      chainId: 333,
    };
    const configs: {
      [x: string]: ChainCfg;
    } = {
      Polygon: {
        name: "polygon",
        chainId: 2,
        implSlot: "0x" + buildOfLen("0x2", 32).toString("hex"),
        deployment: {},
        wrapAssetFromChainIds: [3, 4],
      },
      Harmony: {
        name: "harmony",
        chainId: 3,
        implSlot: "0x" + buildOfLen("0x3", 32).toString("hex"),
        deployment: {},
        wrapAssetFromChainIds: [3, 2],
      },
      Fantom: {
        name: "fantom",
        chainId: 4,
        implSlot: "0x" + buildOfLen("0x4", 32).toString("hex"),
        deployment: {},
        wrapAssetFromChainIds: [2, 4],
      },
    };

    /**
     *
     * SUSY token
     *
     *
     * Token flow:
     *                 –---------> (Polygon)
     *              /               /   \
     *            /               /      \
     * (Solana) –----------> (Harmony)    \
     *           \               \       /
     *            \               \    /
     *              ----------> (Fantom)
     *
     *
     * Legend:
     *  ----- - lock/unlock token flow
     *  – – – - mint/lock token flow
     */

    const overrideProps = (cfg: ChainCfg): Partial<TestProps> => ({
      chainId: web3.utils.numberToHex(cfg.chainId),
      governanceChainId: web3.utils.numberToHex(cfg.chainId),
      governanceContract: cfg.implSlot,
      bridgeChainId: web3.utils.numberToHex(cfg.chainId),
      bridgeGovernanceChainId: web3.utils.numberToHex(cfg.chainId),
      bridgeGovernanceContract: cfg.implSlot,
      emitterChainId: buildOfLen(web3.utils.numberToHex(cfg.chainId), 2).toString("hex"),
    });

    console.log("\n=================================== Mesh testing ===================================\n");

    await runInitialMigration(configs.Polygon.deployment, overrideProps(configs.Polygon));
    await runInitialMigration(configs.Harmony.deployment, overrideProps(configs.Harmony));
    await runInitialMigration(configs.Fantom.deployment, overrideProps(configs.Fantom));

    await runRegisterChainTest(configs.Polygon.deployment, overrideProps(configs.Polygon));
    await runRegisterChainTest(configs.Harmony.deployment, overrideProps(configs.Harmony));
    await runRegisterChainTest(configs.Fantom.deployment, overrideProps(configs.Fantom));

    configs.Polygon.tokenDeployment = await tokenDeployFn(configs.Polygon, tokenConfig, overrideProps(configs.Polygon));
    configs.Harmony.tokenDeployment = await tokenDeployFn(configs.Harmony, tokenConfig, overrideProps(configs.Harmony));
    configs.Fantom.tokenDeployment = await tokenDeployFn(configs.Fantom, tokenConfig, overrideProps(configs.Fantom));

    const toDecimals = (amt: Big) => amt.mul(new Big(10).pow(tokenConfig.decimals!));
    const globalAssetInfo = {
      initSupply: toDecimals(new Big(100_000)),
      polygonReceiver: guardianSet.guardians[1],
      harmonyReceiver: guardianSet.guardians[2],
      fantomReceiver: guardianSet.guardians[3],
    };

    /**
     * Test Mesh flow:
     *  1. Locked 100,000 SUSY on Solana => Minted 100,000 SUSY on Polygon
     *  2. Transfer 50,000 SUSY from Polygon to Harmony (.1 - lock, .2 - mint)
     *  3. Transfer 25,000 SUSY from Harmony to Fantom (.1 - lock, .2 - mint)
     */

    // 1. initial mint occured on Polygon for instance
    await runMintCompleteTransferTest(
      configs.Polygon.deployment,
      configs.Polygon.tokenDeployment!,
      overrideProps(configs.Polygon),
      {
        amount: globalAssetInfo.initSupply,
        receiverAddress: globalAssetInfo.polygonReceiver.address,
      }
    );

    const lockFunds = async (cfg: ChainCfg, amount: Big, recipientChain: string, receiverAddress: string, from: Wallet) => {
      const asset = cfg.tokenDeployment!.wrappedAsset!.connect(from);
      const bridge = cfg.deployment.bridgeImplementation!.connect(from);

      await asset.approve(bridge.address, amount.toFixed());

      // 2.1
      await bridge.transferTokens(
        asset.address,
        amount.toFixed(),
        recipientChain.toString(),
        web3.eth.abi.encodeParameter("address", receiverAddress),
        "0",
        "235"
      );
    };

    const redeemFunds = async (cfg: ChainCfg, amount: Big, from: Wallet, receiverAddr: string, props: Partial<TestProps>) => {
      // const asset = cfg.tokenDeployment!.wrappedAsset!.connect(from);
      const bridge = cfg.deployment.bridgeImplementation!.connect(from);

      // await asset.approve(bridge.address, amount.toFixed());

      const data = VMArgsEncoder.instance
        .completeTransferMethodEncode(
          amount,
          cfg.tokenDeployment!.wrappedAsset.address,
          await cfg.tokenDeployment!.wrappedAsset.chainId(),
          receiverAddr,
          await cfg.deployment.bridgeImplementation!.chainId(),
          "0x0"
        )
        .join("");

      const vm = await signAndEncodeVM(
        0,
        0,
        // web3.utils.numberToHex(await cfg.deployment.bridgeImplementation!.chainId()),
        // web3.utils.numberToHex(await cfg.deployment.bridgeImplementation!.governanceContract()),
        props.bridgeChainId!,
        props.bridgeGovernanceContract!,
        // web3.utils.numberToHex(await cfg.deployment.bridgeImplementation!.chainId()),
        // testProps.bridgeGovernanceContract
        0,
        data,
        guardianSet.privateKeysList,
        0,
        0
      );

      // 2.2
      await bridge.completeTransfer("0x" + vm);
    };

    await configs.Polygon.tokenDeployment!.wrappedAsset!.mint(
      globalAssetInfo.polygonReceiver.address,
      toDecimals(new Big(50_000)).toFixed()
    );

    //2.1
    await lockFunds(
      configs.Polygon,
      toDecimals(new Big(50_000)),
      String(configs.Harmony.chainId),
      globalAssetInfo.harmonyReceiver.address,
      globalAssetInfo.polygonReceiver
    );

    await redeemFunds(
      configs.Harmony,
      toDecimals(new Big(50_000)),
      globalAssetInfo.harmonyReceiver,
      globalAssetInfo.harmonyReceiver.address,
      overrideProps(configs.Harmony)
    );
    
  });
});
