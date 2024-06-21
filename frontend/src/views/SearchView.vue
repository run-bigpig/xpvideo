<script setup>
import {onMounted, ref, watch} from "vue";
import {useRoute} from "vue-router";
import {Search} from "@/../wailsjs/go/controller/AppController.js";

const route = useRoute()
const dataSource = ref([])
const total = ref(0)
const loading = ref(false)
const current = ref(1)
const getList = () => {
  loading.value = true
  Search({
    keyword: route.params.keyword
    , page: current.value
    , pageSize: 24
  }).then(res => {
    dataSource.value = res.data.list
    total.value = res.data.total
    loading.value = false
  })
}

const initData = (page)=>{
  current.value = page
  getList()
}

onMounted(() => {
  initData(1)
})

watch(()=>route.path,()=>{
  initData(1)
})

const handlePageChange = (page, pageSize) => {
  current.value = page
  getList()
}
</script>

<template>
  <a-list
      :grid-props="{ gutter: [10, 10], xs: 24, sm: 12, md: 6, lg: 6, xl:4, xxl: 4 }"
      :data="dataSource"
      :bordered="false"
      :max-height="1080"
      :scrollbar="false"
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
  <a-pagination
      :current=current
      :total="total"
      :page-size="24"
      simple
      @change=handlePageChange
      v-show="total>0"
  />
</template>

<style scoped>

</style>