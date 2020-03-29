var m = require('mithril')
var Tab = require('../models/Tab')
var Collector = require('../models/Collector')

module.exports = {
    view: function () {
        return m('div', {class: 'tabs'}, [
            m('ul', Collector.authenticated ? [
                m('li', {
                    class: Tab.active.name === 'tag' ? 'is-active' : '', onclick: function () {
                        Tab.select('tag')
                    }
                }, [
                    m('a', [
                        m('span', {class: 'icon is-small'}, [m('i', {class: 'fas fa-file-alt'})]),
                        m('span', 'Tags')
                    ])
                ]),
                m('li', {
                    class: Tab.active.name === 'item' ? 'is-active' : '', onclick: function () {
                        Tab.select('item')
                    }
                }, [
                    m('a', [
                        m('span', {class: 'icon is-small'}, [m('i', {class: 'fas fa-edit'})]),
                        m('span', 'Add item')
                    ])
                ])
            ] : [])
        ])
    }
}