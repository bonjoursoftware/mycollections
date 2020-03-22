var m = require('mithril')

var Collector = {
    current: {},
    authenticated: false,

    reset: function () {
        localStorage.clear()
        Collector.current = {}
        Collector.authenticated = false
    },

    store: function () {
        localStorage.setItem('username', Collector.current.username)
        localStorage.setItem('basicauth', btoa(Collector.current.username + ':' + Collector.current.secret))
        Collector.isAuthenticated()
    },

    isAuthenticated: function () {
        return m.request({
            method: 'GET',
            url: '/api/v1/collector',
            headers: { 'Authorization': 'Basic ' + localStorage.getItem('basicauth') },
            withCredentials: true
        }).then(function (result) {
            Collector.authenticated = true
        }).catch(function (err) {
            Collector.authenticated = false
        })
    }
}

module.exports = Collector