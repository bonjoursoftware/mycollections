var m = require('mithril')

var Tab = {
    active: {
        name: 'tag'
    },

    select: function (tab) {
        Tab.active.name = tab
        m.route.set('/' + Tab.active.name)
    }
}

module.exports = Tab