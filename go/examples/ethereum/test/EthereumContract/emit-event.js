module.exports = async function(done) {
  try {
    const HelloWorld = artifacts.require('./contracts/HelloWorld.sol');
    const instance = await HelloWorld.new();
    const result = await instance.emitHello();
    const output = {
      "TxHash": result.tx,
      "ContractAddress": instance.address,
      "ContractJsonAbi": JSON.stringify(instance.abi)
    };
    console.log(JSON.stringify(output, null, 2));
  } catch (e) {
    console.error(e);
  }
  done();
};