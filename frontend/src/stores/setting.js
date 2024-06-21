import {computed, ref} from "vue";
import {defineStore} from "pinia";

export const useSettingStore = defineStore("setting",()=>{
    const setting  = ref({
        source:""
    })
    const getSetting = computed(()=>{
        return setting.value
    })
    function setSetting(data){
        for (let i in data){
            if (setting.value[i] === data[i]){
                continue
            }
            setting.value[i] = data[i]
        }
    }
    return {
        setting,
        getSetting,
        setSetting
    }
})