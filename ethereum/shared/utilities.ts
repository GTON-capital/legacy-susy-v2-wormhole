import { ethers } from "hardhat"
import { BigNumber, Bytes } from "ethers"

export const MAX_UINT = "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
export const ZERO_ADDR = "0x0000000000000000000000000000000000000000"

export function expandTo18Decimals(n: number): BigNumber {
  return BigNumber.from(n).mul(BigNumber.from(10).pow(18))
}
