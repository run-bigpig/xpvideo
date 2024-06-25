<script setup>
import Hls from 'hls.js'
import Flv from 'flv.js'  //用于播放m3u8格式
import DPlayer from 'dplayer'
import {onMounted, ref} from "vue";
import {useRoute} from "vue-router";
import {Detail} from "@/../wailsjs/go/controller/AppController.js";

const playinfo = ref({})
const route = useRoute()
const currentPlayUrl = ref("")
var dp = null
const loadVideo = () => {
  dp = new DPlayer({
    container: document.getElementById('player'), //播放器容器元素
    autoplay: true, //是否自动播放
    live: false, //是否直播
    theme: '#b7daff', //主题色 底部进度条相关颜色
    loop: false, //是否循环播放
    lang: 'zh-cn', //语言
    screenshot: true, //开启截图，如果开启，视频和视频封面需要允许跨域
    hotkey: true, //开启热键，支持快进、快退、音量控制、播放暂停
    preload: 'auto', //视频预加载，可选值: 'none', 'metadata', 'auto'
    volume: 20, //初始化音量
    playbackSpeed: [0.5, 1, 2, 4, 8], //播放速度
    mutex: false, //互斥，阻止多个播放器同时播放，当前播放器播放时暂停其他播放器
    preventClickToggle: false, //阻止点击播放器时候自动切换播放/暂停
    logo: '', //在左上角展示一个 logo，你可以通过 CSS 调整它的大小和位置
    video: {
      pic: '', // 视频封面
      url: '', //视频链接
      thumbnails: '',//视频缩略图，可以使用 DPlayer-thumbnails生成
      type: 'customHls', //可选值: 'auto', 'hls', 'flv', 'dash', 'webtorrent', 'normal' 或其他自定义类型,
      customType: {
        //自定义播放类型文件《type需要设置为'customHls'》
        customHls: function (video, player) {
          const hls = new Hls()
          hls.loadSource(video.src)
          hls.attachMedia(video)
        },
        //自定义播放类型文件《type需要设置为'customFlv'》
        customFlv: function (video, player) {
          const flvPlayer = Flv.createPlayer({
            type: 'flv',
            url: video.src,
          })
          flvPlayer.attachMediaElement(video)
          flvPlayer.load()
        },
      },
    },
  })
  //考虑隐私问题，初始化时把音量设为0,才可自动播放
  dp.volume(10, true, true)
  dp.on('ended', ()=>{
    nextPlay()
  })
  //视频组件加载完成后自动播放
  dp.on('loadedmetadata', function () {
    dp.play()
  })
}
const playVideo = (playurl) => {
  if (playurl.indexOf('m3u8') > 0) {
    dp.switchVideo({
      url: playurl,
      type:"customHls"
    })
  } else {
    dp.switchVideo({
      url: playurl,
      type:"customFlv"
    })
  }
  dp.play()
  currentPlayUrl.value = playurl
}
const getDetail = () => {
  Detail( {
    id: parseInt(route.params.id)
  }).then(res => {
    playinfo.value = res.data
    playVideo(playinfo.value.play[0].url)
  })
}
//获取下一集
const nextPlay = () => {
  if (playinfo.value.play.length > 0) {
    let index = playinfo.value.play.findIndex(item => item.url === currentPlayUrl.value)
    if (index < playinfo.value.play.length - 1) {
      playVideo(playinfo.value.play[index + 1].url)
      return
    }
    playVideo(currentPlayUrl.value)
  }
}
onMounted(() => {
  loadVideo()
  getDetail()
})
</script>

<template>
  <a-row>
    <a-col :span="20">
      <div id="player" class="video-box"></div>
      <a-card :title="playinfo.name" :bordered="false">
        <a-descriptions layout="inline-horizontal" :column="1">
          <a-descriptions-item label="演员:">{{playinfo.actor}}</a-descriptions-item>
          <a-descriptions-item label="描述:">{{playinfo.desc}}</a-descriptions-item>
        </a-descriptions>
      </a-card>
    </a-col>
    <a-col :span="4">
      <a-card title="播放列表" :bordered="false">
        <a-list :grid-props="{ gutter: [10, 10], xs: 24, sm: 24, md: 12, lg: 12, xl:12, xxl:12 }" :bordered="false"
                :data="playinfo.play" :max-height="500">
          <template #item="{item}">
            <a-list-item style="padding: 0;border-radius: 0" :action-layout="'vertical'">
              <div  @click="playVideo(item.url)" :class="['button',item.url===currentPlayUrl?'active':'']">{{ item.name }}</div>
            </a-list-item>
          </template>
        </a-list>
      </a-card>
    </a-col>
  </a-row>
</template>

<style scoped>
.video-box {
  width: 100%;
  height: 65vh;
}

.button {
  width: 100%;
  height: 100%;
  color: #4E5969;
  box-sizing: border-box;
  border: 1px solid #81ecec;
  text-align: center;
  white-space: nowrap; /* 禁止文本换行 */
  overflow: hidden; /* 隐藏溢出内容 */
  text-overflow: ellipsis; /* 用省略号表示被截断的文本 */
  cursor: pointer;
}

.button:hover {
  background-color: #81ecec;
  color:white;
  opacity: 1;
}

.active {
  background-color: #81ecec;
  color:white;
  opacity: 1;
}

</style>