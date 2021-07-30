var m = require('mithril')

var Tab = {
    select: function (tab) {
        m.route.set('/' + encodeURIComponent(tab))
    },

    isTagActive: function () {
        return m.route.get().startsWith('/tag')
    },

    isItemActive: function () {
        return m.route.get().startsWith('/item')
    },

    isSearchActive: function () {
        return m.route.get().startsWith('/search')
    },

    isShareActive: function () {
        return m.route.get().startsWith('/share')
    }
}

module.exports = Tab
