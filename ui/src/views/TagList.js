var m = require('mithril')
var Tag = require('../models/Tag')

module.exports = {
    oninit: Tag.loadList,
    view: function() {
        return m('.tag-list', Tag.list.map(function (tag) {
            return m(m.route.Link, {
                class: 'tag-list-item',
                href: '/tag/' + tag.name
            }, tag.name)
        }))
    }
}