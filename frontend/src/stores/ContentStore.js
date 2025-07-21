import { defineStore } from 'pinia'
import { ValueType } from '../enums/ValueType'

export const useContentStore = defineStore('content-store', {
    state: () => ({
        profile: {
            data: [{
                value: '',
                type: ValueType.NULL
            }]
        }
    }),

    getters:{
        getData() {
            return this.profile.data
        }
    },
    actions: {
        setData(obj) {
            this.profile.data = []
            if (obj === null) {
                return
            }
            for (const objKey of obj) {
                let item = {
                    value: objKey.value,
                    type: objKey.type
                }
                this.profile.data.push(item)
            }
        }
    }
})