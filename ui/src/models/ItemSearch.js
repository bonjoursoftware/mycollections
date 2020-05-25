var m = require('mithril')

var ItemSearch = {
    search: {
        hasRun: false,
        itemName: null,
        result: []
    },

    reset: function () {
        ItemSearch.search = {
            hasRun: false,
            itemName: null,
            result: []
        }
    },

    findByName: function (name) {
        return m.request({
            method: 'GET',
            url: '/api/v1/item/name/' + encodeURIComponent(name),
            headers: {'Authorization': 'Basic ' + localStorage.getItem('basicauth')},
            withCredentials: true
        }).then(function (result) {
            ItemSearch.search.hasRun = true,
            ItemSearch.search.result = result.sort(function (item1, item2) {
                return item1.name.localeCompare(item2.name)
            })
        })
    },

    nameIsEmpty: function () {
        return !ItemSearch.search.itemName
    },
}

module.exports = ItemSearch