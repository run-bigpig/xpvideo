<script setup>
import {onMounted, ref, watch} from "vue";
import {useRoute} from "vue-router";
import {List} from "@/../wailsjs/go/controller/AppController.js";

const route = useRoute()
const dataSource = ref([])
const loading = ref(false)
const getList = () => {
  loading.value = true
  let params = {
    class: 0
    , page: 1
    , pageSize: 24
  }
  List(params).then(res => {
    dataSource.value = res.data.list
    loading.value = false
  })
}
onMounted(() => {
  getList()
})
</script>

<template>
  <a-list
      :grid-props="{ gutter: [10, 10], xs: 24, sm: 12, md: 6, lg: 6, xl:4, xxl: 4 }"
      :data="dataSource"
      :max-height="1080"
      :scrollbar="false"
      :bordered="false"
      :loading=loading
  >
    <template #item="{ item }">
      <a-list-item @click="$router.push('/play/'+item.id)">
        <a-image
            :src=item.img
            :style="{ width: '100%', height: '280px'}"
            :width="'100%'"
            :height="'100%'"
            :fit="'cover'"
            footer-position="outer"
            show-loader
        >
          <template #extra>
            <h4 style="overflow: hidden;text-overflow: ellipsis;white-space: nowrap;">{{ item.name }}</h4>
          </template>
        </a-image>
      </a-list-item>
    </template>
  </a-list>
</template>

<style scoped>

</style>