var m = require('mithril')
var Collector = require('../models/Collector')

module.exports = {
    oninit: Collector.reset,
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
        }, Collector.authenticated ? [
                m('div', {class: 'field is-grouped'}, [
                    m('p', {class: 'control is-expanded'}, [
                        m('label', {class: 'label'}, 'Hi, ' + localStorage.getItem('friendlyname') + '!')
                    ]),
                    m('p', {class: 'control'}, [
                        m('button', {type: 'reset', class: 'button is-danger'}, 'Logout')
                    ]),
                ])
            ]
            :
            [
                m('div', {class: 'field is-grouped'}, [
                    m('p', {class: 'control has-icons-left is-expanded'}, [
                        m('input.input[type=email][placeholder=Email]', {
                            oninput: function (e) {
                                Collector.current.username = e.target.value
                            }, value: Collector.current.username
                        }),
                        m('span', {class: 'icon is-small is-left'}, [
                            m('i', {class: 'fas fa-envelope'})
                        ])
                    ]),
                    m('p', {class: 'control has-icons-left is-expanded'}, [
                        m('input.input[type=password][placeholder=Password]', {
                            oninput: function (e) {
                                Collector.current.secret = e.target.value
                            }, value: Collector.current.secret
                        }),
                        m('span', {class: 'icon is-small is-left'}, [
                            m('i', {class: 'fas fa-lock'})
                        ])
                    ]),
                    m('p', {class: 'control'}, [
                        m('button', {type: 'submit', class: 'button is-success'}, 'Login')
                    ])
                ])
            ])
    }
}