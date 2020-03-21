var m = require('mithril')
var Collector = require('../models/Collector')

module.exports = {
    view: function () {
        return m('form', {
            onsubmit: function (e) {
                e.preventDefault()
                Collector.store()
            },
            onreset: function (e) {
                e.preventDefault()
                Collector.reset()
            }
        }, [
            m('input.input[type=text][placeholder=Username]', {
                oninput: function (e) {
                    Collector.current.username = e.target.value
                }, value: Collector.current.username
            }),
            m('input.input[type=password][placeholder=Secret]', {
                oninput: function (e) {
                    Collector.current.secret = e.target.value
                }, value: Collector.current.secret
            }),
            m('button.button[type=submit]', 'Login'),
            m('button.button[type=reset]', 'Logout'),
        ])
    }
}