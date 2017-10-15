module.exports = trigger => {
  const { chain, address } = trigger.connector.bitcoinTransaction

  return {
    match: ({ type, network, transaction, block }) => {
      if (type !== 'BITCOIN') { return false }
      if (network !== chain) { return false }
      return address.toLowerCase() === (transaction.from || '').toLowerCase() ||
             address.toLowerCase() === (transaction.to || '').toLowerCase()
    },
    normalizeEvent: ({ transaction, block }) => ({
      blockId: '...',
      fees: '...',
      from: '...',
      payload: {},
      to: '...',
      transactionId: '...',
      value: '...'
    })
  }
}
