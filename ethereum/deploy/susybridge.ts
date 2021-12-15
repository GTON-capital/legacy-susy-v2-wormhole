import { BridgeImplementation__factory } from "./../typechain/factories/BridgeImplementation__factory";
import "@typechain/hardhat";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";
import "hardhat-abi-exporter";
import "@nomiclabs/hardhat-web3";

import { ethers, web3 } from "hardhat";
import { Signer, Wallet } from "ethers";

import { Implementation__factory } from "./../typechain/factories/Implementation__factory";
import { Setup__factory } from "./../typechain/factories/Setup__factory";
import { Config } from "./cfg";
import { Setup } from "../typechain/Setup";
import { Implementation } from "../typechain/Implementation";
import { Wormhole } from "../typechain/Wormhole";
import { Wormhole__factory } from "../typechain/factories/Wormhole__factory";

import { SuSyBridge } from "../typechain/SuSyBridge";
import { SuSyBridge__factory } from "./../typechain/factories/SuSyBridge__factory";
import { SuSyTokenBridge } from "../typechain/SuSyTokenBridge";
import { SuSyTokenBridge__factory } from "./../typechain/factories/SuSyTokenBridge__factory";

import { BridgeSetup } from "../typechain/BridgeSetup";
import { BridgeImplementation } from "../typechain/BridgeImplementation";
import { TokenImplementation } from "../typechain/TokenImplementation";
import { TokenImplementation__factory } from "../typechain/factories/TokenImplementation__factory";

import { TestERC20 } from "../typechain/TestERC20";
import { TestERC20__factory } from "./../typechain/factories/TestERC20__factory";

type DeploymentContracts = {
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

async function deployGovernance(
  deployedContracts: DeploymentContracts,
  deployer: Signer,
  guardians: string[],
  chainId: string,
  governanceContract: string
) {
  const setupFactory = (await ethers.getContractFactory("Setup")) as Setup__factory;
  const implFactory = (await ethers.getContractFactory("Implementation")) as Implementation__factory;
  const wormholeFactory = (await ethers.getContractFactory("Wormhole")) as Wormhole__factory;

  deployedContracts.setupContract = await setupFactory.connect(deployer).deploy();
  await deployedContracts.setupContract.deployed();

  deployedContracts.implementationContract = await implFactory.connect(deployer).deploy();
  await deployedContracts.implementationContract.deployed();

  const setupTx = await deployedContracts.setupContract.setup(
    deployedContracts.implementationContract.address,
    guardians,
    chainId,
    chainId,
    governanceContract
  );

  deployedContracts.wormholeContract = await wormholeFactory
    .connect(deployer)
    .deploy(deployedContracts.setupContract.address, setupTx.data);
  await deployedContracts.wormholeContract.deployed();
}

// async function deployBridge(deployedContracts: DeploymentContracts, deployer: Signer) {
//   /**
//    * Migration N2
//    */
//   const bridgeSetupFactory = (await ethers.getContractFactory("BridgeSetup")) as BridgeImplementation__factory;
//   const bridgeImplFactory = (await ethers.getContractFactory("BridgeImplementation")) as BridgeImplementation__factory;
//   const tokenImplFactory = (await ethers.getContractFactory("TokenImplementation")) as TokenImplementation__factory;

//   /** Token Migration */
//   const legitOwner = guardianSet.guardians[0];
//   const token = await tokenImplFactory.connect(legitOwner).deploy();
//   await token
//     .connect(legitOwner)
//     .initialize(
//       testProps.testToken!.name,
//       testProps.testToken!.symbol,
//       testProps.testToken!.decimals,
//       testProps.testToken!.sequence,
//       legitOwner.address,
//       testProps.testToken!.chainId,
//       testProps.testToken!.nativeContract
//     );

//   deployedContracts.tokenImplementation = token;

//   deployedContracts.bridgeSetup = (await bridgeSetupFactory.deploy()) as BridgeSetup;
//   deployedContracts.bridgeImplementation = (await bridgeImplFactory.deploy()) as BridgeImplementation;

//   const bridgeSetupTx = await deployedContracts.bridgeSetup.setup(
//     deployedContracts.bridgeImplementation.address,
//     testProps.bridgeChainId,
//     deployedContracts.wormholeContract.address,
//     testProps.bridgeGovernanceChainId, // testGovernanceChainId,
//     testProps.bridgeGovernanceContract, // testGovernanceContract,
//     deployedContracts.tokenImplementation.address,
//     deployedContracts.WETH.address
//   );

//   const susyTokenBridgeFactory = (await ethers.getContractFactory("SuSyTokenBridge")) as SuSyTokenBridge__factory;

//   deployedContracts.susyTokenBridge = await susyTokenBridgeFactory.deploy(
//     deployedContracts.bridgeSetup.address,
//     bridgeSetupTx.data
//   );

//   deployedContracts.bridgeImplementation = await bridgeImplFactory.attach(deployedContracts.susyTokenBridge.address);
// }

const privateKeys = [
  "0xd77f0f006ddba135ff78bd7f9fb92a828b2b45a6eeb66fba85d7ab6393badd8c",
  "0x807bb23c8cbdaa9a2c6bc14928bbd33beab3b97331f98b76ec536a71816c18f8",
  "0x36fbabdecac294469b8b83e7c814e628c7c9a5995f9808af5b1ed205bbc26840",
  "0x66efdac9af9469833757e07ec56e01e06361ba4c9c3847e1d9fba1b35a20a4a8",
  "0xb29adf24c32b5aa2de6f072ad2586264cd75c1a873162602483f13bda2954b1f",
];

async function deploy() {
  // const n = 5;
  const walletsList = privateKeys.map((x) => new Wallet(x));
  const guardiansList = walletsList.map((x) => x.address);

  const cfg = new Config();

  // const [deployer] = waffle.provider.getWallets();
  const [deployer] = await ethers.getSigners();

  const harmonyChainId = 166;

  const governanceDeployResult: DeploymentContracts = {};

  await deployGovernance(
    governanceDeployResult,
    deployer,
    guardiansList,
    web3.utils.numberToHex(harmonyChainId),
    // "0x63564c40", // 1666600000,
    "0x0000000000000000000000000000000000000000000000000000000000000004"
  );

  console.log("deployment finished");
  console.log({
    setup: governanceDeployResult.setupContract!.address,
    implementation: governanceDeployResult.implementationContract!.address,
    wormhole: governanceDeployResult.wormholeContract!.address,
  });
}

deploy()
  .then(() => process.exit(0))
  .catch((err: Error) => {
    console.log(err.message);
    process.exit(1);
  });
