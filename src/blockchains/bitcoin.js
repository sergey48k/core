const createClient = require('./createClient')

module.exports = async => {
  return createClient({
    type: 'BITCOIN',
    network: 'MAINNET',
    isConnected: () => false,
    onTransaction: callback => {}
  })
}
