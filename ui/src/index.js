var m = require('mithril')

var EmptyView = require('./views/EmptyView')
var TabView = require('./views/TabView')
var TagList = require('./views/TagList')
var TagItem = require('./views/TagItem')
var ItemForm = require('./views/ItemForm')
var CollectorForm = require('./views/CollectorForm')
var ItemSearchForm = require('./views/ItemSearchForm')
var ShareForm = require('./views/ShareForm')
var LoginView = require('./views/LoginView')

var collector = document.getElementById('collector')
m.mount(collector, CollectorForm)

var tab = document.getElementById('tab')
m.mount(tab, TabView)

var mycollections = document.getElementById('mycollections')
m.route(mycollections, '/', {
    '/': EmptyView,
    '/tag': TagList,
    '/tag/:name': TagItem,
    '/item': ItemForm,
    '/item/:id': ItemForm,
    '/search': ItemSearchForm,
    '/share': ShareForm,
    '/login/:username/:token': LoginView,
})
