var m = require('mithril')
var Tag = require('../models/Tag')

module.exports = {
    oninit: Tag.loadList,
    view: function () {
        return m('aside', {class: 'menu'}, [
            m('ul', {class: 'menu-list'}, Tag.list.map(function (tag) {
                return m('li', [
                    m('a', {href: '/#!/tag/' + tag.name}, tag.name)
                ])
            }))
        ])
    }
}