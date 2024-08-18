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
                        placeholder="请输入客户名称"
                        clearable
                        :maxlength="100"
                    />
                </el-form-item>
                <el-form-item label="编码" prop="code">
                    <el-input v-model="formData.code" placeholder="请输入客户编码" clearable />
                </el-form-item>
                <el-form-item label="手机号码" prop="phone">
                    <el-input v-model="formData.phone" placeholder="请输入手机号码" clearable />
                </el-form-item>
                <el-form-item label="公司" prop="company">
                    <el-input v-model="formData.company" placeholder="请输入公司名称" clearable />
                </el-form-item>
                <el-form-item label="地址" prop="address">
                    <el-input v-model="formData.address" placeholder="请输入地址" clearable />
                </el-form-item>
                <el-form-item label="性别" prop="gender">
                    <el-radio-group v-model="formData.gender">
                        <el-radio label="女">女</el-radio>
                        <el-radio label="男">男</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="地址" prop="address">
                    <el-input v-model="formData.address" placeholder="请输入地址" clearable />
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
import { customerEdit, customerAdd, customerDetail } from '@/api/basic/customer'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑客户' : '新增客户'
})
const formData = reactive({
    id: '',
    name: '',
    code: '',
    gender: '男',
    phone: '',
    company: '',
    address: '',
    remark: '',
    status: 1
})

const formRules = {
    code: [],
    name: [
        {
            required: true,
            message: '请输入客户名称',
            trigger: ['blur']
        }
    ],
    phone: [
        {
            required: true,
            message: '请输入手机号码',
            trigger: ['blur']
        },
        {
            pattern: /^1[3456789]\d{9}$/,
            message: '手机号格式错误',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    await formRef.value?.validate()
    mode.value == 'edit' ? await customerEdit(formData) : await customerAdd(formData)
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
    const data = await customerDetail({
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
