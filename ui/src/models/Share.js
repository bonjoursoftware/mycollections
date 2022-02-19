var m = require('mithril')

var Share = {
    current: {},

    reset: function () {
        Share.current = {}
    },

    requestShare: function () {
        return m.request({
            method: 'POST',
            url: '/api/v1/login/' + encodeURIComponent(Share.current.email),
            headers: { 'Authorization': 'Basic ' + localStorage.getItem('basicauth') },
            withCredentials: true
        }).then(function (result) {
            Share.reset()
        }).catch(function (err) {
            console.log('unable to share collection')
        })
    },

    isEmpty: function () {
        return !Share.current.email
    }
}

module.exports = Share
