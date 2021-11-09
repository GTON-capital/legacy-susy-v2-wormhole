// workaround for
// only replay-protected (EIP-155) transactions allowed over RPC.
// error
module.exports = function(deployer) {
  const networkCfg = deployer.networks[deployer.network];
  return {
    gasLimit: networkCfg.gas,
    gasPrice: networkCfg.gasPrice,
    chainId: networkCfg.network_id,
  };
};
