var m = require('mithril')

var Tag = {
    list: [],
    loadList: function () {
        Tag.list = []
        return m.request({
            method: 'GET',
            url: '/api/v1/tag',
            headers: { 'Authorization': 'Basic ' + localStorage.getItem('basicauth') },
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
            headers: { 'Authorization': 'Basic ' + localStorage.getItem('basicauth') },
            withCredentials: true
        }).then(function (result) {
            Tag.item = result
        })
    }
}

module.exports = Tag