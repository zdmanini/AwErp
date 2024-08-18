<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form
                class="ls-form"
                ref="formRef"
                :rules="rules"
                :model="formData"
                label-width="84px"
            >
              <el-form-item label="上级栏目" prop="parent">
                <el-tree-select
                    class="flex-1"
                    v-model="formData.parent"
                    :data="options.typeList"
                    clearable
                    node-key="name"
                    :props="{
                            value: 'name',
                            label: 'name'
                        }"
                    check-strictly
                    :default-expand-all="true"
                    placeholder="请选择上级分类">
                </el-tree-select>
              </el-form-item>
                <el-form-item label="分类名称" prop="name">
                    <el-input v-model="formData.name" placeholder="请输入分类名称" clearable />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import {typeLists,typeAdd,typeEdit} from "@/api/basic/info";
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑字典类型' : '新增字典类型'
})
const formData = reactive({
    id: '',
    name: '',
    parent: '',
})

const rules = {

}

const options = reactive({
   typeList: []
})

const handleSubmit = async () => {
    await formRef.value?.validate()
    mode.value == 'edit' ? await typeEdit(formData) : await typeAdd(formData)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const handleClose = () => {
    emit('close')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = (data: Record<any, any>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            //@ts-ignore
            formData[key] = data[key]
        }
    }
}

const getParents = async () => {
  options.typeList = [
      {name: '顶级分类'},
  ]
  const data = await typeLists()
  data.forEach(item => {
      options.typeList.push(item)
  })
}

onMounted(() => {
    getParents()
})

defineExpose({
    open,
    setFormData
})
</script>
