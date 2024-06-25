import {computed, ref} from "vue";
import {defineStore} from "pinia";

export const useSettingStore = defineStore("setting",()=>{
    const setting  = ref({
        sources:[],
        refresh:0
    })
    const getSetting = computed(()=>{
        return setting.value
    })
    function setSetting(data){
        for (let i in data){
            if (setting.value[i] === data[i]){
                continue
            }
            setting.value[i] = {...data[i]}
        }
        setting.value.refresh++
    }
    return {
        setting,
        getSetting,
        setSetting
    }
})