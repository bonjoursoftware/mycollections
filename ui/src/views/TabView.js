var m = require('mithril')
var Tab = require('../models/Tab')
var Collector = require('../models/Collector')

module.exports = {
    view: function () {
        return m('div', {class: 'tabs'}, [
            m('ul', Collector.authenticated ? [
                m('li', {
                    class: Tab.isTagActive() ? 'is-active' : '', onclick: function () {
                        Tab.select('tag')
                    }
                }, [
                    m('a', [
                        m('span', {class: 'icon is-small'}, [m('i', {class: 'fas fa-file-alt'})]),
                        m('span', 'Tags')
                    ])
                ]),
                m('li', {
                    class: Tab.isItemActive() ? 'is-active' : '', onclick: function () {
                        Tab.select('item')
                    }
                }, [
                    m('a', [
                        m('span', {class: 'icon is-small'}, [m('i', {class: 'fas fa-edit'})]),
                        m('span', 'Item')
                    ])
                ]),
                m('li', {
                    class: Tab.isSearchActive() ? 'is-active' : '', onclick: function () {
                        Tab.select('search')
                    }
                }, [
                    m('a', [
                        m('span', {class: 'icon is-small'}, [m('i', {class: 'fas fa-search'})]),
                        m('span', 'Search')
                    ])
                ]),
                m('li', {
                    class: localStorage.getItem('readonly') === 'true' ? 'is-hidden' : Tab.isShareActive() ? 'is-active' : '', onclick: function () {
                        if (localStorage.getItem('readonly') === 'true') {
                            return
                        }
                        Tab.select('share')
                    }
                }, [
                    m('a', [
                        m('span', {class: 'icon is-small'}, [m('i', {class: 'fas fa-share-alt'})]),
                        m('span', 'Share')
                    ])
                ])
            ] : [])
        ])
    }
}
