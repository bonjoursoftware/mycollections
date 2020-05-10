var m = require('mithril')
var Tag = require('./Tag')

var Collector = {
    current: {},
    authenticated: false,

    reset: function () {
        localStorage.clear()
        Collector.current = {}
        Collector.authenticated = false
        m.route.set('/')
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
            localStorage.setItem('friendlyname', result.friendlyname)
            Collector.authenticated = true
            Tag.reset()
        }).catch(function (err) {
            Collector.authenticated = false
            m.route.set('/')
        })
    }
}

module.exports = Collector