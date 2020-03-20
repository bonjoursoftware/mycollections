var m = require('mithril')

var Item = {
    current: {},

    reset: function () {
        Item.current = {}
    },

    save: function () {
        Item.current.name = Item.current.name ? Item.current.name : 'unnamed'
        Item.current.tags = Item.current.tags ? Item.current.tags.split(',') : 'untagged'
        return m.request({
            method: 'POST',
            url: '/api/v1/item',
            body: Item.current,
            withCredentials: true
        }).then(function () {
            Item.reset()
        })
    }
}

module.exports = Item