var m = require('mithril')
var Tag = require('../models/Tag')

module.exports = {
    oninit: function (vnode) {
        Tag.loadItem(vnode.attrs.name)
    },
    view: function () {
        return m('div', Tag.item.map(function (item) {
            return m('div', item.name)
        }))
    }
}