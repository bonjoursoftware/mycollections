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
            Tag.list = result.map(o => o.name).sort()
        }).catch(function (err) {
            console.log('User not logged in')
        })
    },

    current: undefined,
    item: [],
    loadItem: function (tagName) {
        Tag.current = tagName
        Tag.item = []
        return m.request({
            method: 'GET',
            url: '/api/v1/tag/' + encodeURIComponent(tagName),
            headers: { 'Authorization': 'Basic ' + localStorage.getItem('basicauth') },
            withCredentials: true
        }).then(function (result) {
            Tag.item = result.sort(function (item1, item2) {
                return item1.name.localeCompare(item2.name)
            })
        })
    },

    newTag: undefined,
    addTag: function () {
        Tag.list.push(Tag.newTag)
        Tag.list.sort()
        Tag.newTag = undefined
    }
}

module.exports = Tag