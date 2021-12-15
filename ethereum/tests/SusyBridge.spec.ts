import { ethers, waffle } from "hardhat";
import BigNumber from "big.js";

import { Wormhole } from "../typechain/Wormhole";
import { TokenBridge } from "../typechain/SuSyBridge";
import { BridgeImplementation } from "../typechain/BridgeImplementation";
import { TokenImplementation } from "../typechain/TokenImplementation";

describe("Tests: SuSy Bridge", () => {
  const [wallet, other] = waffle.provider.getWallets();
});
