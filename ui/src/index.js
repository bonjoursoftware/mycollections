var m = require('mithril')

var TagList = require('./views/TagList')
var TagItem = require('./views/TagItem')
var ItemForm = require('./views/ItemForm')

var mycollections = document.getElementById('mycollections')
m.route(mycollections, '/tag', {
    '/tag': TagList,
    '/tag/:name': TagItem,
    '/item': ItemForm
})