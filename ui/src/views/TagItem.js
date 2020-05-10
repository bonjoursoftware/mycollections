var m = require('mithril')
var Tag = require('../models/Tag')

module.exports = {
    oninit: function (vnode) {
        Tag.loadItem(vnode.attrs.name)
    },
    view: function () {
        return m('section', {class: 'section'}, [
            m('nav', {class: 'breadcrumb is-medium', 'aria-label': 'breadcrumbs'}, [
                m('ul', [
                    m('li', [m('a', {
                        onclick: function () {
                            Tag.reset()
                        }
                    }, 'Tags')]),
                    m('li', {class: 'is-active'}, [m('a', {href: '#', 'aria-current': 'page'}, Tag.current)])
                ])
            ]),
            m('aside', {class: 'menu'}, [
                m('ul', {class: 'menu-list'}, Tag.item.map(function (item) {
                    return m('li', [
                        m('a', {href: '/#!/item/' + encodeURIComponent(item.id)}, item.name)
                    ])
                }))
            ])
        ])
    }
}