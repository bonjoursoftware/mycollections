var m = require('mithril')
var Login = require('../models/Login')

module.exports = {
    oninit: function (vnode) {
        Login.reset()
        Login.store(vnode.attrs.username, vnode.attrs.token)
    },
    view: function () {
        return m('div')
    }
}