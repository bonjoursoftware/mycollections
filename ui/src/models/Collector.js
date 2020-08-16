var m = require('mithril')
var Tag = require('./Tag')

var Collector = {
    current: {},
    loginRequestRegistered: false,
    authenticated: false,

    reset: function () {
        localStorage.clear()
        Collector.current = {}
        Collector.authenticated = false
        Collector.loginRequestRegistered = false
        m.route.set('/')
    },

    requestLogin: function () {
    Collector.loginRequestRegistered = true
        return m.request({
            method: 'GET',
            url: '/api/v1/login/' + encodeURIComponent(Collector.current.username),
            withCredentials: false
        }).catch(function (err) {
            Collector.authenticated = false
            m.route.set('/')
        })
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