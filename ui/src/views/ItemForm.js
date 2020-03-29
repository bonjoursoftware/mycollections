var m = require('mithril')
var Item = require('../models/Item')
var Tag = require('../models/Tag')

module.exports = {
    oninit: Item.reset,
    view: function () {
        return m('section', {class: 'section'}, [
            m('form', {
                onsubmit: function (e) {
                    e.preventDefault()
                    Item.save()
                    Array.from(document.getElementsByClassName('is-info')).map(function (el) {
                        el.classList.replace('is-info', 'is-light')
                    })
                }
            }, [
                m('div', {class: 'field is-grouped'}, [
                    m('p', {class: 'control is-expanded'}, [
                        m('input.input[type=text][placeholder=Name]', {
                            oninput: function (e) {
                                Item.current.name = e.target.value
                            }, value: Item.current.name
                        })
                    ]),
                    m('p', {class: 'control'}, [
                        m('button', {
                            type: 'submit',
                            class: 'button is-success',
                            disabled: Item.isEmpty()
                        }, 'Add')
                    ])
                ]),
                m('div', {class: 'field'}, [
                    m('div', {class: 'tags are-medium'}, [
                        Tag.list.map(function (tag) {
                            return m('a', {
                                class: 'tag is-light', onclick: function (e) {
                                    var tag = e.target
                                    if (tag.classList.contains('is-light')) {
                                        tag.classList.replace('is-light', 'is-info')
                                        Item.addTag(tag.textContent)
                                    } else {
                                        tag.classList.replace('is-info', 'is-light')
                                        Item.removeTag(tag.textContent)
                                    }
                                }
                            }, tag)
                        })
                    ])
                ])
            ])
        ])
    }
}