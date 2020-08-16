var m = require('mithril')
var Collector = require('./Collector')

var Login = {

    reset: function () {
        localStorage.clear()
    },

    store: function (username, token) {
        localStorage.setItem('username', username)
        localStorage.setItem('basicauth', btoa(username + ':' + token))
        Collector.isAuthenticated()
    }
}

module.exports = Login