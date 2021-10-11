var Migrations = artifacts.require("Migrations");

const buildDeployerProps = require('../scripts/override')

module.exports = function(deployer) {
    // Deploy the Migrations contract as our only task

    deployer.deploy(Migrations, buildDeployerProps(deployer));
};