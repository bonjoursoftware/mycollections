var m = require('mithril')

var TagList = require('./views/TagList')

var mycollections = document.getElementById('mycollections')
m.mount(mycollections, TagList)