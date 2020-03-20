var m = require('mithril')
var Item = require('../models/Item')

module.exports = {
    oninit: Item.reset,
    view: function () {
        return m('form', {
            onsubmit: function (e) {
                e.preventDefault()
                Item.save()
            }
        }, [
            m('label.label', 'Item'),
            m('input.input[type=text][placeholder=Name]', {
                oninput: function (e) {
                    Item.current.name = e.target.value
                }, value: Item.current.name
            }),
            m('label.label', 'Tags'),
            m('input.input[type=text][placeholder=Tags]', {
                oninput: function (e) {
                    Item.current.tags = e.target.value
                }, value: Item.current.tags
            }),
            m('button.button[type=submit]', 'Add')
        ])
    }
}