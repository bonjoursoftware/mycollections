var Collector = {
    current: {},

    reset: function () {
        Collector.current = {}
        localStorage.clear()
    },

    store: function () {
        localStorage.setItem('username', Collector.current.username)
        localStorage.setItem('secret', Collector.current.secret)
    }
}

module.exports = Collector