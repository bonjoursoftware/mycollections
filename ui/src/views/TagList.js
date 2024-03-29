var m = require('mithril')
var Tag = require('../models/Tag')

module.exports = {
    oninit: Tag.load,
    view: function () {
        return m('section', {class: 'section'}, [
            m('nav', {class: 'breadcrumb is-medium', 'aria-label': 'breadcrumbs'}, [
                m('ul', [
                    m('li', {class: 'is-active'}, [m('a', {
                        onclick: function () {
                            Tag.reset()
                        }
                    }, 'Tags')])
                ])
            ]),
            m('aside', {class: 'menu'}, [
                m('ul', {class: 'menu-list'}, Tag.list.map(function (tag) {
                    return m('li', [
                        m('a', {href: '/#!/tag/' + encodeURIComponent(tag)}, tag)
                    ])
                }))
            ]), localStorage.getItem('readonly') === 'true' ?
            m('p')
            :
            m('form', [
                m('div', {class: 'field is-grouped'}, [
                    m('p', {class: 'control is-expanded'}, [
                        m('input.input[type=text][placeholder=New tag]', {
                            oninput: function (e) {
                                Tag.newTag = e.target.value
                            }, value: Tag.newTag
                        })
                    ]),
                    m('p', {class: 'control'}, [
                        m('button', {
                            class: 'button is-success', onclick: function (e) {
                                e.preventDefault()
                                Tag.addTag()
                            }
                        }, 'Add')
                    ])
                ])
            ])
        ])
    }
}