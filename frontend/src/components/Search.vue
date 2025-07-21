<template>
  <div style="background-color: darkkhaki; height: 50px">
<!--  <div style="height: 50px">-->
<!--    <input type="text" v-model="value">-->
    <input type="text" placeholder="请输入内容" class="my-input" :value="modelValue" @input="change" />
<!--    <button @click="queryData">查询</button>-->
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useContentStore } from '../stores/ContentStore'
import { QueryContentByValue, QueryAllContent } from "../../wailsjs/go/controller/ContentController";
import { ValueType } from "../enums/ValueType.js";

defineOptions({
  name: "Search"
})

const props = defineProps(['modelValue', 'changeFunc']);
const emit = defineEmits(['update:modelValue']);

const contentStore = useContentStore();

const value = ref('')

function change(e) {
  const v = e.target.value;
  emit('update:modelValue', v)

  props.changeFunc()


  // if (props.modelValue === ValueType.PICTURE) {
  //   alert('图片不支持搜索!')
  //   return
  // }
  //
  // let queryContent;
  // console.log(v)
  // console.log(typeof v)
  // if (v === null || v === '') {
  //   queryContent = QueryAllContent();
  // } else {
  //   queryContent = QueryContentByValue(v);
  // }
  // queryContent.then(result => {
  //   contentStore.setData(result)
  // })
}

function queryData() {
  // let data = contentStore.profile.data;
  // console.log(value)
  // console.log(value.value)
  // contentStore.setData(data.filter(item => item.value.includes(value.value)))
  const queryContent = QueryContent(value.value);
  // console.log(queryContent)
  // let newContent = [];
  // contentStore.profile.data = []
  queryContent.then(result => {
    // console.log(typeof result)
    // console.log(result)
    contentStore.setData(result)
  })
  // console.log(newContent)
  // contentStore.setData(newContent)
}



</script>

<style scoped>
.my-input {
  width: 97%;
  height: 38px;
  padding: 6px 10px;
  border: 0px solid #ccc;
  /*border-radius: 6px;*/
  font-size: 14px;
  outline: none;
}

.my-input:focus {
  border-color: #42b983;
  box-shadow: 0 0 3px rgba(66, 185, 131, 0.6);
}
</style>