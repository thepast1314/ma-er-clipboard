import {defineStore} from 'pinia'

export const useStyleStore = defineStore('style-store', {
    state: () => ({
        profile: {
            config: {
                window: {
                    width: 0,
                    height: 0
                }
            }
        }
    }),
    actions: {
        async setConfig(obj) {
            this.profile.config.window = obj
        }
    }
})