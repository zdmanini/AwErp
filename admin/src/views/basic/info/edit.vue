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
                <el-form-item label="字典名称" prop="name">
                    <el-input v-model="formData.name" placeholder="请输入字典名称" clearable />
                </el-form-item>
                <el-form-item label="字典类型" prop="type">
                  <el-tree-select
                      class="flex-1"
                      v-model="formData.type"
                      :data="optionsData.typeLists"
                      clearable
                      node-key="name"
                      :props="{
                            value: 'name',
                            label: 'name'
                        }"
                      check-strictly
                      :default-expand-all="true"
                      placeholder="请选择分类">
                  </el-tree-select>


                </el-form-item>
                <el-form-item label="字典状态" required prop="status">
                    <el-radio-group v-model="formData.status">
                        <el-radio :label="1">正常</el-radio>
                        <el-radio :label="0">停用</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input
                        v-model="formData.remark"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                        clearable
                        maxlength="200"
                        show-word-limit
                    />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import Popup from '@/components/popup/index.vue'
import { infoAdd, infoEdit, typeLists } from '@/api/basic/info'
import feedback from '@/utils/feedback'
import { useDictOptions } from '@/hooks/useDictOptions'
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑字典项' : '新增字典项'
})
const formData = reactive({
    id: '',
    name: '',
    type: '',
    status: 1,
    remark: ''
})

const rules = {
    name: [
        {
            required: true,
            message: '请输入字典名称',
            trigger: ['blur']
        }
    ],
    type: [
        {
            required: true,
            message: '请输入字典项',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    await formRef.value?.validate()
    mode.value == 'edit' ? await infoEdit(formData) : await infoAdd(formData)
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

const { optionsData } = useDictOptions<{
    typeLists: any[]
}>({
    typeLists: {
        api: typeLists
    }
})

defineExpose({
    open,
    setFormData
})
</script>
