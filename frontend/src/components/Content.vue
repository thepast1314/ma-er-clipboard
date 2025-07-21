<template>
  <div style="background-color: aqua" >
    <Search v-model:model-value="inputValue" v-model:change-func="copy"/>
    <div class="option">
      <div v-for="(item, index) in valueTypeEnum" @click="getContentByType(item.value)">
        <span>{{item.label}}</span>
      </div>
    </div>
    <div :style="{height: (height - 50 - 38 - 10) + 'px' }" class="slide">
      <div v-for="(item, index) in contentStore.profile.data">
        <div class="itme" @dblclick="copy(item.value)">
          {{index + 1}}  {{ item.type }} - {{ item.value }}
        </div>
        <div class="delete" @click="deleteItem(item.value)">
          DEL
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref} from 'vue'
import { useContentStore } from '../stores/ContentStore'
import { QueryAllValueTypeEnum } from "../../wailsjs/go/controller/EnumController";
import {
  QueryAllContent,
  QueryContentByCondition,
  DeleteContentByPrimary
} from "../../wailsjs/go/controller/ContentController";
import Search from "./Search.vue";
import {ValueType} from "../enums/ValueType.js";

const contentStore = useContentStore();

let valueTypeEnum = ref([]);
let valueType = ref(-1)
const inputValue = ref('')

defineOptions({
  name: "Content"
})

const props = defineProps({
  height: {
    type: Number,
    default: 0
  }
})

onMounted(() => {
  QueryAllValueTypeEnum().then(result => {
    result.unshift({value: -1, label: '全部'})
    valueTypeEnum = result
  })
})


function copy() {
  console.log(inputValue)


  getContentByCondition(valueType.value, inputValue.value)
}

function getContentByType(type) {
  valueType.value = type
  console.log(inputValue)

  getContentByCondition(type, inputValue.value)
}

function getContentByCondition(type, value) {
  if (value != null && value !== '' && type === ValueType.PICTURE) {
    alert("图片不支持搜索!")
    return
  }

  console.log(type)
  if (type === -1 && (value == null || value === '')) {
    QueryAllContent().then(result => contentStore.setData(result))
  } else {
    QueryContentByCondition({type: type, value: value}).then(result => {
      contentStore.setData(result)
    })
  }
}

function deleteItem(value) {
  DeleteContentByPrimary(value)
  copy()
}

</script>

<style scoped>
.itme {
  height: 80px;
  width: 88%;
  text-align: left;
  padding: 10px;
}
.itme:hover {
  background-color: #add8e6; /* 淡蓝色 */
}
.slide {
  overflow: auto;       /* 启用滚动 */
  scrollbar-width: none; /* Firefox 隐藏滚动条 */
}
.option {
  height: 50px;
  display: flex;
  justify-content: space-around;
  align-items: center;
}
.option > div {
  width: 120px;
  height: 100%;
  align-content: center;
}

.option > div:hover {
  background-color: #add8e6; /* 淡蓝色 */
}

.slide > div {
  display: flex;
  justify-content: space-between;
}

.slide .delete {
  width: 12%;
  /*height: 100%;*/
  align-content: center;
}
.slide .delete:hover {
  background-color: red;
  color: white;
}
</style>