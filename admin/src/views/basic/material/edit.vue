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
            <el-form ref="formRef" :model="formData" label-width="84px" :rules="formRules">
                <el-form-item label="名称" prop="name">
                    <el-input
                        v-model="formData.name"
                        placeholder="请输入物料名称"
                        clearable
                        :maxlength="100"
                    />
                </el-form-item>
                <el-form-item label="编码" prop="code">
                    <el-input v-model="formData.code" placeholder="请输入物料编码" clearable />
                </el-form-item>
                <el-form-item label="类型" prop="type">
                    <el-select v-model="formData.type" placeholder="请选择物料类型">
                        <el-option></el-option>
                      </el-select>
                  </el-form-item>
                <el-form-item label="供应商" prop="supplier_id">
                    <el-select v-model="formData.supplier_id" placeholder="请选择供应商">
                        <el-option v-for="item in optionsData.supplier.lists" :key="item.id" :label="item.name" :value="item.id" />
                    </el-select>
                  </el-form-item>
                <el-form-item label="单价" prop="unit_price">
                    <el-input-number v-model="formData.unit_price" :precision="2" />
                  </el-form-item>
                <el-form-item label="单位" prop="unit">
                    <el-input v-model="formData.unit" placeholder="请输入单位" clearable />
                  </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input
                        v-model="formData.remark"
                        placeholder="请输入备注"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                        maxlength="200"
                        show-word-limit
                    />
                </el-form-item>
                <el-form-item label="状态" prop="status">
                    <el-switch v-model="formData.status" :active-value="1" :inactive-value="0" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import { materialEdit, materialAdd, materialDetail } from '@/api/basic/material'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import {useDictOptions} from "@/hooks/useDictOptions";
import {supplierLists} from "@/api/basic/supplier";
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑物料' : '新增物料'
})
const formData = reactive({
    id: '',
    name: '',
    code: '',
    type: '',
    supplier_id: 0,
    unit_price: 0.0,
    unit: '',
    remark: '',
    status: 1
})

const formRules = {
    code: [],
    name: [
        {
            required: true,
            message: '请输入物料名称',
            trigger: ['blur']
        }
    ],
}

const { optionsData } = useDictOptions<{
  supplier: any[]
}>({
  supplier: {
    api: supplierLists,
    params: {
      status: 1,
      pageNo: 1,
      pageSize: 10000
    }
  }
})

const handleSubmit = async () => {
    await formRef.value?.validate()
    mode.value == 'edit' ? await materialEdit(formData) : await materialAdd(formData)
    feedback.msgSuccess('操作成功')
    popupRef.value?.close()
    emit('success')
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
        //TODO：因为后端返回字段为is_stop
        else {
            //@ts-ignore
            formData[key] = data['is_stop']
        }
    }
}

const getDetail = async (row: Record<string, any>) => {
    const data = await materialDetail({
        id: row.id
    })
    setFormData(data)
}

const handleClose = () => {
    emit('close')
}

defineExpose({
    open,
    setFormData,
    getDetail
})
</script>
