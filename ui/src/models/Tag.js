var m = require('mithril')

var Tag = {
    list: [],
    loadList: function () {
        return m.request({
            method: 'GET',
            url: '/api/v1/tag',
            withCredentials: true
        }).then(function (result) {
            Tag.list = result
        })
    }
}

module.exports = Tag