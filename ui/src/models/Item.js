var m = require('mithril')

var Item = {
    current: {
        tags: []
    },

    reset: function () {
        Item.current = {
            tags: []
        }
    },

    save: function () {
        if (Item.isEmpty()) {
            return
        }
        return m.request({
            method: 'POST',
            url: '/api/v1/item',
            headers: {'Authorization': 'Basic ' + localStorage.getItem('basicauth')},
            body: Item.current,
            withCredentials: true
        }).then(function () {
            Item.reset()
        })
    },

    addTag: function (tag) {
        Item.current.tags.push(tag)
    },

    removeTag: function (targetTag) {
        Item.current.tags = Item.current.tags.filter(tag => tag !== targetTag)
    },

    isEmpty: function () {
        return !Item.current.name || Item.current.tags.length < 1
    }
}

module.exports = Item