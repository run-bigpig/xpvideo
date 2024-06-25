<script setup>
import Sider  from "@/components/Sider.vue";
import {ref} from "vue";
import {IconSearch} from '@arco-design/web-vue/es/icon';
import router from "@/router/index.js";
const search = ref(true)
const keyword = ref('')
const showSearch = (val) => {
  search.value = val
}
const searchKeyword = () => {
  if (keyword.value===""){
    return
  }
  router.push('/search/'+keyword.value)
}
</script>

<template>
  <a-layout class="layout-main">
    <Sider @show-search="showSearch" />
    <a-layout>
      <a-layout style="padding: 0;">
        <a-layout-header :style="{textAlign:'center'}" v-show=search>
          <a-input-search v-model="keyword" :style="{width:'500px',background:'rgba(0,0,0,0.05)'}" placeholder="请输入您要搜索的内容" @search="searchKeyword" @keyup.enter="searchKeyword">
            <template #button-icon>
              <IconSearch style="font-size: 20px"/>
            </template>
          </a-input-search>
        </a-layout-header>
        <a-layout-content><RouterView /></a-layout-content>
      </a-layout>
    </a-layout>
  </a-layout>
</template>

<style scoped>
.layout-main {
  height: 100vh;
  background: var(--color-fill-2);
}
.layout-main :deep(.arco-layout-header)  {
  height: 64px;
  line-height: 64px;
  background: var(--color-bg-3);
}

.layout-main :deep(.arco-layout-content) {
  color: var(--color-text-2);
  padding: 24px;
  background: var(--color-bg-3);
}

.layout-main :deep(.arco-list){
  border-radius: 0;
}
</style>