module.exports = async () => [
  await require('./ethereum')('MAINNET'),
  await require('./ethereum')('KOVAN'),
  await require('./bitcoin')(),
  await require('./bitcoin')(true)
]
