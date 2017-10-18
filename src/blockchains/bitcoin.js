const io = require('socket.io-client')
const createClient = require('./createClient')

const request = endpoint => fetch([process.env.BITCOIN_MAINNET, endpoint].join(''))
  .then(x => x.json())

const getBlock = hash => request(`/insight-api/block/${hash}`)
const getTransaction = hash => request(`/insight-api/tx/${hash}`)

const onTransaction = socket => callback => socket.on('block', async blockHash => {
  const blockDetails = await getBlock(blockHash)
  console.log('start loop')
  blockDetails.tx.forEach(async transactionId => {
    const transaction = await getTransaction(transactionId)
    callback(transaction, blockDetails)
  })
  console.log('end loop')
})

module.exports = () => {
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
    onTransaction: onTransaction(socket)
  })
}
