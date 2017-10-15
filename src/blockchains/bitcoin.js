const io = require('socket.io-client')
const createClient = require('./createClient')

module.exports = async (testnet = false) => {

  const url = testnet ? 'https://test-insight.bitpay.com/' : 'https://insight.bitpay.com/'
  let connected = false
  const socket = io(url)

  socket.on('connect', () => {
    socket.emit('subscribe', 'inv')
    connected = true
  })
  socket.on('disconnect', () => (connected = false))
  socket.on('error', error => {
    debugger
    connected = false
    throw error
  })

  return createClient({
    type: 'BITCOIN',
    network: testnet ? 'TESTNET' : 'MAINNET',
    isConnected: () => connected,
    onTransaction: callback => socket.on('tx', transaction => {
      callback(transaction)
    })
  })
}
