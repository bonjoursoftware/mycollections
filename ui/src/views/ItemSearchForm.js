var m = require('mithril')
var ItemSearch = require('../models/ItemSearch')

module.exports = {
    oninit: function(vnode) {
        ItemSearch.find()
    },
    view: function () {
        return m('section', {class: 'section'}, [
            m('form', {
                onsubmit: function (e) {
                    e.preventDefault()
                    ItemSearch.find()
                }
            }, [
                m('div', {class: 'field is-grouped'}, [
                    m('p', {class: 'control is-expanded'}, [
                        m('input.input[type=text][placeholder=Name]', {
                            id: 'searchInput',
                            oninput: function (e) {
                                ItemSearch.search.itemName = e.target.value
                            }, value: ItemSearch.search.itemName
                        })
                    ]),
                    m('p', {class: 'control'}, [
                        m('button', {
                            type: 'submit',
                            class: 'button is-success',
                            disabled: ItemSearch.nameIsEmpty()
                        }, 'Search')
                    ]),
                    m('p', {class: 'control', type: 'reset'}, [
                        m('button', {
                            class: 'button is-danger',
                            disabled: ItemSearch.nameIsEmpty(),
                            onclick: function () {
                                ItemSearch.reset()
                                document.getElementById('searchInput').focus()
                            }
                        }, 'Reset')
                    ]),
                ]),
            ]),
            m('br'),
            m('nav', {class: 'breadcrumb is-medium', 'aria-label': 'breadcrumbs'}, [
                m('ul', [
                    m('li', {class: 'is-active'}, [m('a', ItemSearch.search.hasRun ? `Found ${ItemSearch.search.result.length} item(s)` : '')])
                ])
            ]),
            m('aside', {class: 'menu'}, [
                m('ul', {class: 'menu-list'}, ItemSearch.search.result.map(function (item) {
                    return m('li', [
                        m('a', {href: '/#!/item/' + encodeURIComponent(item.id)}, `${item.name} (${item.tags.join(', ')})`)
                    ])
                }))
            ])
        ])
    }
}