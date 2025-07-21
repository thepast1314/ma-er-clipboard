<script setup>
import Content from "./components/Content.vue";
import Tool from "./components/Tool.vue";
import { useContentStore } from './stores/ContentStore'
import {onMounted, ref} from 'vue'
import { QueryAllConfig } from '../wailsjs/go/controller/ConfigController'
import { QueryAllContent } from "../wailsjs/go/controller/ContentController";
import { useStyleStore } from "./stores/StyleStore.js";

let contentStore = useContentStore();
let styleStore = useStyleStore();

let height = ref(500)

onMounted(() => {
  QueryAllContent().then(result => contentStore.setData(result))
  QueryAllConfig().then(result => {
    styleStore.setConfig(result)
    console.log(result)
    height.value = styleStore.profile.config.window.height - 50
    console.log(height)
  })
  console.log(styleStore.profile.config.window)

})

</script>

<template>
  <div class="periphery">
    <div class="container">
      <Content style="width: 90%" :style="{ height: height + 'px' }" :height="height"/>
      <Tool style="width: 10%" :style="{ height: height + 'px' }"/>
    </div>
  </div>
</template>

<style>
.container {
  display: flex;         /* 设置为水平布局 */
  gap: 10px;             /* 可选：两个 box 之间的间距 */
}

.periphery {
  padding: 0px 10px 10px 10px;
  background-color: white;
}

#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
