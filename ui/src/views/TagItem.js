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
                    m('li', [m('a', {href: '/#!/tag'}, 'Tags')]),
                    m('li', {class: 'is-active'}, [m('a', {href: '#', 'aria-current': 'page'}, Tag.current)])
                ])
            ]),
            Tag.item.map(function (item) {
                return m('div', item.name)
            })
        ])
    }
}