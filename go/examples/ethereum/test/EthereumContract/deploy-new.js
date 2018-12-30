module.exports = async function(done) {
  try {
    const HelloWorld = artifacts.require('./contracts/HelloWorld.sol');
    const instance = await HelloWorld.new();
    const output = {
      "ContractAddress": instance.address,
      "ContractJsonAbi": JSON.stringify(instance.abi)
    };
    console.log(JSON.stringify(output, null, 2));
  } catch(e) {
    console.error(e);
  }
  done();
};