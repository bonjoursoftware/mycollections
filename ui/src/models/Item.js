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

    load: function (id) {
        return m.request({
            method: 'GET',
            url: '/api/v1/item/' + id,
            headers: {'Authorization': 'Basic ' + localStorage.getItem('basicauth')},
            withCredentials: true
        }).then(function (result) {
            Item.current = result
        })
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

    delete: function () {
        if (Item.isEmpty()) {
            return
        }
        return m.request({
            method: 'DELETE',
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
    },

    exists: function () {
        return Item.current.id !== undefined
    }
}

module.exports = Item