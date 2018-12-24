var HelloWorld = artifacts.require("HelloWorld");
module.exports = function(deployer) {
  deployer.deploy(HelloWorld, "hello");
  // Additional contracts can be deployed here
};