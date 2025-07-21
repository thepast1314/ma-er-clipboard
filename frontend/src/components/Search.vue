<template>
  <div style="background-color: darkkhaki; height: 50px">
    <input type="text" placeholder="请输入内容" class="my-input" :value="modelValue" @input="change" />
  </div>
</template>

<script setup>
defineOptions({
  name: "Search"
})

const props = defineProps(['modelValue', 'changeFunc']);
const emit = defineEmits(['update:modelValue']);

function change(e) {
  const v = e.target.value;
  emit('update:modelValue', v)

  props.changeFunc()
}
</script>

<style scoped>
.my-input {
  width: 97%;
  height: 38px;
  padding: 6px 10px;
  border: 0px solid #ccc;
  font-size: 14px;
  outline: none;
}

.my-input:focus {
  border-color: #42b983;
  box-shadow: 0 0 3px rgba(66, 185, 131, 0.6);
}
</style>