import { ethers, waffle } from "hardhat"
import { BigNumber } from "ethers"

import { TestERC20 } from "../../typechain/TestERC20"
import { WrappedNative } from "../../typechain/WrappedNative"
import { QuickPair } from "../../typechain/QuickPair"
import { QuickFactory } from "../../typechain/QuickFactory"
import { QuickRouter01 } from "../../typechain/QuickRouter01"
import { Calibrator } from "../../typechain/Calibrator"

import { MdexPair } from "../../typechain/MdexPair"
import { MdexFactory } from "../../typechain/MdexFactory"
import { MdexRouter } from "../../typechain/MdexRouter"

import {
  expandTo18Decimals,
} from "./utilities"

import { Fixture } from "ethereum-waffle"

interface TokensFixture {
  token0: TestERC20
  token1: TestERC20
  token2: TestERC20
  weth: WrappedNative
}

export const tokensFixture: Fixture<TokensFixture> =
  async function (
    [wallet, other],
    provider
  ): Promise<TokensFixture> {
  const tokenFactory = await ethers.getContractFactory("TestERC20")
  const tokenA = (await tokenFactory.deploy(
    BigNumber.from(2).pow(255)
  )) as TestERC20
  const tokenB = (await tokenFactory.deploy(
    BigNumber.from(2).pow(255)
  )) as TestERC20
  const tokenC = (await tokenFactory.deploy(
    BigNumber.from(2).pow(255)
  )) as TestERC20

  const [token0, token1, token2] = [tokenA, tokenB, tokenC].sort(
    (tokenA, tokenB) =>
      tokenA.address.toLowerCase() < tokenB.address.toLowerCase() ? -1 : 1
  )

  const wethFactory = await ethers.getContractFactory("WrappedNative")
  let weth = await wethFactory.deploy() as WrappedNative
  while (weth.address.toLowerCase() > token0.address.toLowerCase()) {
    weth = await wethFactory.deploy() as WrappedNative
  }

  return { token0, token1, token2, weth }
}

interface UniswapFixture extends TokensFixture {
  uniswapV2Factory: QuickFactory
  uniswapV2Router01: QuickRouter01
  uniswapV2PairGTON_WETH: QuickPair
  uniswapV2PairGTON_USDC: QuickPair
  mdexFactory: MdexFactory
  mdexRouter: MdexRouter
  mdexPairGTON_WETH: MdexPair
}

const uniswapFixture: Fixture<UniswapFixture> =
  async function (
    [wallet, other],
    provider
  ): Promise<UniswapFixture> {
  const { token0: gton, token1: usdt, token2: usdc, weth: weth } = await tokensFixture([wallet, other], provider)

  const uniswapV2FactoryFactory = await ethers.getContractFactory(
    "QuickFactory"
  )
  const uniswapV2Factory = await uniswapV2FactoryFactory.deploy(
    wallet.address
  ) as QuickFactory

  await uniswapV2Factory.setFeeTo(other.address)

  const uniswapV2Router01Factory = await ethers.getContractFactory(
    "QuickRouter01"
  )
  const uniswapV2Router01 = await uniswapV2Router01Factory.deploy(
    uniswapV2Factory.address,
    weth.address
  ) as QuickRouter01

  await uniswapV2Factory.createPair(weth.address, gton.address)

  const uniswapV2PairFactory = await ethers.getContractFactory(
    "QuickPair"
  )
  // log pairV2 bytecode for init code hash in the router
  // let bytecode = uniswapV2PairFactory.bytecode
  // console.log(ethers.utils.solidityKeccak256(["bytes"],[bytecode]))
  let pairAddressGTON_WETH = await uniswapV2Factory.getPair(weth.address, gton.address)
  const uniswapV2PairGTON_WETH = uniswapV2PairFactory.attach(pairAddressGTON_WETH) as QuickPair

  let liquidityGTON
  let block
  let timestamp

  liquidityGTON = expandTo18Decimals(10)
  let liquidityWETH = expandTo18Decimals(20)
  await gton.approve(uniswapV2Router01.address, liquidityGTON)
  block = await wallet.provider.getBlock("latest")
  timestamp = block.timestamp
  // let k = await uniswapV2PairGTON_WETH.kLast()
  // console.log(k.toString())
  await uniswapV2Router01.addLiquidityETH(
    gton.address,
    liquidityGTON,
    liquidityGTON,
    liquidityWETH,
    wallet.address,
    timestamp + 3600,
    {value: liquidityWETH}
  )
  // k = await uniswapV2PairGTON_WETH.kLast()
  // console.log(k.toString())

  await uniswapV2Factory.createPair(usdc.address, gton.address)

  // log pairV2 bytecode for init code hash in the router
  // let bytecode = uniswapV2PairFactory.bytecode
  // console.log(ethers.utils.solidityKeccak256(["bytes"],[bytecode]))
  let pairAddressGTON_USDC = await uniswapV2Factory.getPair(usdc.address, gton.address)
  const uniswapV2PairGTON_USDC = uniswapV2PairFactory.attach(pairAddressGTON_USDC) as QuickPair

  liquidityGTON = expandTo18Decimals(10)
  let liquidityUSDC = expandTo18Decimals(50)
  await gton.approve(uniswapV2Router01.address, liquidityGTON)
  await usdc.approve(uniswapV2Router01.address, liquidityUSDC)
  block = await wallet.provider.getBlock("latest")
  timestamp = block.timestamp
  let k = await uniswapV2PairGTON_USDC.kLast()
  // console.log(k.toString())
  await uniswapV2Router01.addLiquidity(
    gton.address,
    usdc.address,
    liquidityGTON,
    liquidityUSDC,
    liquidityGTON,
    liquidityUSDC,
    wallet.address,
    timestamp + 3600
  )

  k = await uniswapV2PairGTON_USDC.kLast()
  // console.log(k.toString())
  liquidityGTON = BigNumber.from(1)
  liquidityUSDC = BigNumber.from(5)
  await gton.transfer(other.address, liquidityGTON)
  await usdc.transfer(other.address, liquidityUSDC)
  await gton.connect(other).approve(uniswapV2Router01.address, liquidityGTON)
  await usdc.connect(other).approve(uniswapV2Router01.address, liquidityUSDC)
  block = await wallet.provider.getBlock("latest")
  timestamp = block.timestamp
  await uniswapV2Router01.connect(other).addLiquidity(
    gton.address,
    usdc.address,
    liquidityGTON,
    liquidityUSDC,
    liquidityGTON,
    liquidityUSDC,
    wallet.address,
    timestamp + 3600
  )

  const mdexFactoryFactory = await ethers.getContractFactory(
    "MdexFactory"
  )
  const mdexFactory = await mdexFactoryFactory.deploy(
    wallet.address
  ) as MdexFactory

  // console.log(await mdexFactory.getInitCodeHash())
  await mdexFactory.setInitCodeHash("0x0fd8d405689cfdaa9e8621294709df601788f38378a22ebcaff404049bf31af9")

  await mdexFactory.setFeeTo(other.address)
  await mdexFactory.setFeeToRate(2)

  const mdexRouterFactory = await ethers.getContractFactory(
    "MdexRouter"
  )

  const mdexRouter = await mdexRouterFactory.deploy(
    mdexFactory.address,
    weth.address
  ) as MdexRouter

  await mdexFactory.createPair(weth.address, gton.address)

  const mdexPairFactory = await ethers.getContractFactory(
    "MdexPair"
  )
  // log pairV2 bytecode for init code hash in the router
  // let bytecode = uniswapV2PairFactory.bytecode
  // console.log(ethers.utils.solidityKeccak256(["bytes"],[bytecode]))
  pairAddressGTON_WETH = await mdexFactory.getPair(weth.address, gton.address)
  const mdexPairGTON_WETH = mdexPairFactory.attach(pairAddressGTON_WETH) as MdexPair

  liquidityGTON = BigNumber.from("733147580690169526")
  liquidityWETH = BigNumber.from("307447487")
  await gton.approve(mdexRouter.address, liquidityGTON)
  block = await wallet.provider.getBlock("latest")
  timestamp = block.timestamp
  await mdexRouter.addLiquidityETH(
    gton.address,
    liquidityGTON,
    liquidityGTON,
    liquidityWETH,
    wallet.address,
    timestamp + 3600,
    {value: liquidityWETH}
  )
  // console.log(
  //   "lp balance",
  //   (await mdexPairGTON_WETH.balanceOf(wallet.address)).toString()
  // )

  liquidityGTON = BigNumber.from("7331475806")
  liquidityWETH = BigNumber.from("3")
  await gton.transfer(other.address, liquidityGTON)
  await wallet.sendTransaction({
    to: other.address,
    value: liquidityWETH
  })
  await gton.connect(other).approve(mdexRouter.address, liquidityGTON)
  block = await wallet.provider.getBlock("latest")
  timestamp = block.timestamp
  await mdexRouter.connect(other).addLiquidityETH(
    gton.address,
    liquidityGTON,
    liquidityGTON,
    liquidityWETH,
    wallet.address,
    timestamp + 3600,
    {value: liquidityWETH}
  )

  return {
    token0: gton,
    token1: usdt,
    token2: usdc,
    weth,
    uniswapV2Factory,
    uniswapV2Router01,
    uniswapV2PairGTON_WETH,
    uniswapV2PairGTON_USDC,
    mdexFactory,
    mdexRouter,
    mdexPairGTON_WETH
  }
}

interface CalibratorFixture extends UniswapFixture {
  calibrator: Calibrator
  calibratorMdex: Calibrator
}

export const calibratorFixture: Fixture<CalibratorFixture> =
  async function ([wallet, other, nebula], provider): Promise<CalibratorFixture> {
    const {
      token0: gton,
      token1: usdt,
      token2: usdc,
      weth,
      uniswapV2Factory,
      uniswapV2Router01,
      uniswapV2PairGTON_WETH,
      uniswapV2PairGTON_USDC,
      mdexFactory,
      mdexRouter,
      mdexPairGTON_WETH
    } = await uniswapFixture([wallet, other], provider)

    const calibratorFactory = await ethers.getContractFactory(
      "Calibrator"
    )
    const calibrator = (await calibratorFactory.deploy(
      gton.address,
      uniswapV2Router01.address,
      "QUICK"
    )) as Calibrator

    const calibratorMdex = (await calibratorFactory.deploy(
      gton.address,
      mdexRouter.address,
      "MDEX"
    )) as Calibrator

    return {
      token0: gton,
      token1: usdt,
      token2: usdc,
      weth,
      uniswapV2Factory,
      uniswapV2Router01,
      uniswapV2PairGTON_WETH,
      uniswapV2PairGTON_USDC,
      mdexFactory,
      mdexRouter,
      mdexPairGTON_WETH,
      calibrator,
      calibratorMdex
    }
  }
