var m = require('mithril')
var Share = require('../models/Share')

module.exports = {
    oninit: function (vnode) {
        Share.reset()
        if (localStorage.getItem('readonly') === 'true') {
            m.route.set('/')
        }
    },
    view: function () {
        return m('section', {class: 'section'}, [
            m('form', {
                onsubmit: function (e) {
                    e.preventDefault()
                    Share.requestShare()
                }
            }, [
                m('div', {class: 'field is-grouped'}, [
                    m('p', {class: 'control is-expanded'}, [
                        m('input.input[type=email][placeholder=Email address]', {
                            id: 'shareInput',
                            oninput: function (e) {
                                Share.current.email = e.target.value
                            }, value: Share.current.email
                        })
                    ]),
                    m('p', {class: 'control'}, [
                        m('button', {
                            type: 'submit',
                            class: 'button is-success',
                            disabled: Share.isEmpty()
                        }, 'Share')
                    ]),
                    m('p', {class: 'control', type: 'reset'}, [
                        m('button', {
                            class: 'button is-danger',
                            disabled: Share.isEmpty(),
                            onclick: function () {
                                Share.reset()
                                document.getElementById('shareInput').focus()
                            }
                        }, 'Reset')
                    ]),
                ]),
            ])
        ])
    }
}
