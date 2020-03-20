var m = require('mithril')
var Tag = require('../models/Tag')

module.exports = {
    oninit: function (vnode) {
        Tag.loadItem(vnode.attrs.name)
    },
    view: function () {
        return m('.item-list', Tag.item.map(function (item) {
            return m('.item-list-item', item.name)
        }))
    }
}