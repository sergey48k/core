const io = require('socket.io-client')
const createClient = require('./createClient')

module.exports = async () => {
  let connected = false
  const socket = io(process.env.BITCOIN_MAINNET)

  socket.on('connect', () => {
    socket.emit('subscribe', 'inv')
    connected = true
  })
  socket.on('disconnect', () => (connected = false))
  socket.on('error', error => {
    connected = false
    throw error
  })

  return createClient({
    type: 'BITCOIN',
    network: 'MAINNET',
    isConnected: () => connected,
    onTransaction: callback => socket.on('block', block => {
      callback({}, block)
    })
  })
}
