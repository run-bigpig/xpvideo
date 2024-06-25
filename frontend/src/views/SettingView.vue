<script setup>
import {onMounted, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import {useSettingStore} from "@/stores/setting.js";
import {GetSourceList, Setting} from "@/../wailsjs/go/controller/AppController.js";
import {IconPlus, IconMinus} from '@arco-design/web-vue/es/icon';

const store = useSettingStore()
const form = ref({
  sources: [],
})
const AddSource = (index) => {
  if (form.value.sources[index].url === '' || form.value.sources[index].name === '') {
    Message.error({
      content: '请输入完整信息',
      showIcon: true,
      duration: 1000
    })
    return
  }
  form.value.sources.push({url: '', name: ''})
}

const CheckDefaultSource = (index) => {
  let flag = false
  for (let i = 0; i < form.value.sources.length; i++) {
    if (i===index && form.value.sources[index].default===true){
      flag = true
      continue
    }
    form.value.sources[i].default = false
  }
  if (!flag) {
    form.value.sources[0].default = true
  }
}

const RemoveSource = (index) => {
  form.value.sources.splice(index, 1)
}

const getSource = () => {
  GetSourceList().then(res => {
    let data = res.data
    form.value.sources = data.list
    form.value.defaultSource = data.default
  })
}

onMounted(() => {
  getSource()
})

const onSubmit = () => {
  Setting(form.value).then(res => {
    store.setSetting(form.value)
    Message.info({
      content: res.msg,
      showIcon: true,
      duration: 1000
    })
  })
}

</script>

<template>
  <a-card :style="{padding:'20px',textAlign:'center'}" :bordered="false">
    <a-form :model="form" layout="horizontal">
      <a-row>
        <a-checkbox-group :max="1">
          <a-form-item v-for="(source,index) of form.sources" :field="`sources[${index}].value`"
                       :label="`数据源-${index+1}`" :key="index">
            <a-space>
              <a-input v-model="source.name" placeholder="数据源名称" :style="{width:'100px'}"/>
              <a-input v-model="source.url" placeholder="数据源地址" :style="{width:'300px'}"/>
              <a-switch type="round" v-model="source.default" :style="{width:'32px'}" @click="CheckDefaultSource(index)">
                <template #checked>
                  主
                </template>
                <template #unchecked>
                  备
                </template>
              </a-switch>
              <a-button type="outline" status="danger" @click="RemoveSource(index)">
                <template #icon>
                  <IconMinus/>
                </template>
              </a-button>
              <a-button type="outline" @click="AddSource(index)" v-show="form.sources.length===index+1">
                <template #icon>
                  <IconPlus/>
                </template>
              </a-button>
            </a-space>
          </a-form-item>
        </a-checkbox-group>
      </a-row>
      <a-row>
        <a-form-item>
          <a-button-group>
            <a-button type="primary" @click="onSubmit">保存</a-button>
          </a-button-group>
        </a-form-item>
      </a-row>
    </a-form>
  </a-card>
</template>

<style scoped>

</style>