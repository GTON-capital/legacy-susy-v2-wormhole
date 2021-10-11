require('dotenv').config({ path: "../.env" });

const Setup = artifacts.require("Setup");
const Implementation = artifacts.require("Implementation");
const Wormhole = artifacts.require("Wormhole");

const buildDeployerProps = require('../scripts/override')

// CONFIG
const initialSigners = JSON.parse(process.env.INIT_SIGNERS);
const chainId = process.env.INIT_CHAIN_ID;
const governanceChainId = process.env.INIT_GOV_CHAIN_ID;
const governanceContract = process.env.INIT_GOV_CONTRACT; // bytes32

module.exports = async function (deployer) {
    // deploy setup
    await deployer.deploy(Setup, buildDeployerProps(deployer));

    // deploy implementation
    await deployer.deploy(Implementation, buildDeployerProps(deployer));

    // encode initialisation data
    const setup = new web3.eth.Contract(Setup.abi, Setup.address);
    const initData = setup.methods.setup(
        Implementation.address,
        initialSigners,
        chainId,
        governanceChainId,
        governanceContract
    ).encodeABI();

    // deploy proxy
    await deployer.deploy(Wormhole, Setup.address, initData, buildDeployerProps(deployer));
};
