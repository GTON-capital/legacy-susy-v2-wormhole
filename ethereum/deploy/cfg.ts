export class Config {
  initialSigners: string[]; // array of guardians
  privateKey: string;

  /** governance */
  chainId: string;
  governanceChainId: string;
  governanceContract: string; // impl

  /** app specific */
  bridgeChainId: string;
  bridgeGovernanceChainId: string;
  bridgeGovernanceContract: string; // impl
  WETH: string;

  constructor() {
    this.initialSigners = JSON.parse(process.env.INIT_SIGNERS || "[]") as string[];

    this.privateKey = process.env.PRIVATE_KEY!;
    this.chainId = process.env.INIT_CHAIN_ID!;
    this.governanceChainId = process.env.INIT_GOV_CHAIN_ID!;
    this.governanceContract = process.env.INIT_GOV_CONTRACT!; // bytes32;

    this.bridgeChainId = process.env.BRIDGE_INIT_CHAIN_ID!;
    this.bridgeGovernanceChainId = process.env.BRIDGE_INIT_GOV_CHAIN_ID!;
    this.bridgeGovernanceContract = process.env.BRIDGE_INIT_GOV_CONTRACT!; // bytes32

    this.WETH = process.env.BRIDGE_INIT_WETH!;
  }
}
