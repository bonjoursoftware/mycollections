var m = require('mithril')
var Tag = require('../models/Tag')

module.exports = {
    oninit: Tag.loadList,
    view: function() {
        return m('.tag-list', Tag.list.map(function (tag) {
            return m('.tag-list-item', tag.name)
        }))
    }
}