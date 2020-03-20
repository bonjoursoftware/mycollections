var m = require('mithril')

var Tag = {
    list: [],
    loadList: function () {
        Tag.list = []
        return m.request({
            method: 'GET',
            url: '/api/v1/tag',
            withCredentials: true
        }).then(function (result) {
            Tag.list = result
        }).catch(function (err) {
            console.log('User not logged in')
        })
    },

    item: [],
    loadItem: function (tagName) {
        Tag.item = []
        return m.request({
            method: 'GET',
            url: '/api/v1/item/' + tagName,
            withCredentials: true
        }).then(function (result) {
            Tag.item = result
        })
    }
}

module.exports = Tag