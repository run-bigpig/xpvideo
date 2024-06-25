<template>
  <a-layout-sider breakpoint="md">
    <a-row>
      <a-col :span="24" :style="{textAlign: 'center'}">
        <a-avatar :size="64" shape="square" @click="$router.push('/')"  :image-url="logoSrc" :style="{backgroundColor: 'rba(var(--color-primary))',cursor: 'pointer'}">Arco
        </a-avatar>
      </a-col>
    </a-row>
    <a-row>
      <a-col :span="24">
      <a-menu
            :style="{ width: '100%' }"
            :selected-keys=selectedKeys
            :default-open-keys="defaultOpenKeys"
        >
          <template v-for="item in menuList">
            <a-menu-item v-if="!item.children || item.children.length===0" :key="'/list/'+item.id"
                         @click="$router.push('/list/'+item.id)">
              <template #icon>
                <IconVideoCamera/>
              </template>
              {{ item.name }}
            </a-menu-item>
            <a-sub-menu v-else :key="'pid_'+item.id">
              <template #title>
                <span><IconCalendar/>{{ item.name }}</span>
              </template>
              <a-menu-item v-for="item1 in item.children" :key="'/list/'+item1.id"
                           @click="$router.push('/list/'+item1.id)">
                {{ item1.name }}
              </a-menu-item>
            </a-sub-menu>
          </template>
          <a-menu-item @click="$router.push('/setting')" :key="'/setting'">
            <template #icon>
              <IconSettings/>
            </template>
            设置数据源
          </a-menu-item>
        </a-menu>
      </a-col>
    </a-row>
  </a-layout-sider>
</template>

<script setup>
import {computed, onMounted, ref, watch} from 'vue';
import logoSrc from "@/assets/logo.png"
import {IconCalendar, IconVideoCamera, IconAlignLeft, IconAlignRight, IconSettings} from '@arco-design/web-vue/es/icon';
import {useRoute} from "vue-router";
import {useSettingStore} from "@/stores/setting.js";
import {Class} from "@/../wailsjs/go/controller/AppController.js";

const store = useSettingStore()
const emit = defineEmits(['show-search'])
const route = useRoute()
const menuList = ref([])
const defaultOpenKeys = ref([])

const selectedKeys = computed(() => {
  if (route.name === "play") {
    return [localStorage.getItem("lastPath")]
  }
  return [route.path]
})

const getMenuList = () => {
  Class().then(res => {
    menuList.value = []
    let list = res.data
    for (let i = 0; i < list.length; i++) {
      if (list[i].pid === 0) {
        list[i].children = []
        for (let j = 0; j < list.length; j++) {
          if (list[j].pid === list[i].id) {
            if (list[j].id === parseInt(route.params.id)) {
              defaultOpenKeys.value.push("pid_" + list[i].id)
            }
            list[i].children.push(list[j])
          }
        }
        menuList.value.push(list[i])
      }
    }
  })
}

watch(() => route.path, () => {
  if (route.name === "setting") {
    emit('show-search', false)
    return
  }
  emit('show-search', true)
})

onMounted(() => {
  getMenuList()
})

// 监听设置
watch(() => store.getSetting.sources, () => {
  getMenuList()
}, {deep: true})
</script>

<style scoped>
.logo img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

</style>