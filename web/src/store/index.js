import { createStore } from 'vuex'

const store = createStore({
    state () {
        return {
            token: '',

            // 动态标签
            editableTabsValue: '/product/list',
            editableTabs: [{
                title: '商品列表',
                name: '/product/list',
            }],

            // 动态步骤条激活状态
            stepsActive: 0,

            // 结果页标题状态
            pageTitle: ''
        }
    },
    mutations: {
        SET_TOKEN: (state, token) => {
            state.token = token
            localStorage.setItem("token", token)
        },
        resetState: (state) => {
            state.token = ''
        },
        addTab(state, tab) {
            let index = state.editableTabs.findIndex(s => s.name === tab.name)
            if (index === -1){
                state.editableTabs.push({
                    title: tab.title,
                    name: tab.name,
                });
            }
            state.editableTabsValue = tab.name;
        },
        setActive(state, active) {
            state.stepsActive = active
        },
        setPageTitle(state, title) {
            state.pageTitle = title
        }
    },
    actions: {},
    modules: {}
})

export default store;