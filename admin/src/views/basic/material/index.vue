<template>
    <div class="material-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="编码">
                    <el-input
                        class="w-[280px]"
                        v-model="queryParams.code"
                        clearable
                        @keyup.enter="resetPage"
                    />
                </el-form-item>
                <el-form-item label="名称">
                    <el-input
                        class="w-[280px]"
                        v-model="queryParams.name"
                        clearable
                        @keyup.enter="resetPage"
                    />
                </el-form-item>
                <el-form-item label="状态">
                    <el-select class="w-[280px]" v-model="queryParams.status">
                        <el-option label="全部" :value="-1" />
                        <el-option label="正常" :value="1" />
                        <el-option label="停用" :value="0" />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <el-button v-perms="['basic:material:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <el-table class="mt-4" size="large" v-loading="pager.loading" :data="pager.lists">
                <el-table-column label="编码" prop="code" min-width="100" />
                <el-table-column label="名称" prop="name" min-width="100" />
                <el-table-column label="联系电话" prop="phone" min-width="160" />
                <el-table-column label="公司名称" prop="company" min-width="160" show-overflow-tooltip/>
                <el-table-column label="联系地址" prop="address" min-width="160" show-overflow-tooltip/>
                <el-table-column
                    label="备注"
                    prop="remarks"
                    min-width="100"
                    show-overflow-tooltip
                />
                <el-table-column label="添加时间" prop="createTime" min-width="180" />
                <el-table-column label="状态" prop="status" min-width="100">
                    <template #default="{ row }">
                        <el-switch
                            v-model="row.status"
                            :active-value="1"
                            :inactive-value="0"
                            @change="handleSwitch(row)"
                        />
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['basic:material:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['basic:material:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.id)"
                        >
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
    </div>
</template>
<script lang="ts" setup name="material">
import { materialDelete, materialLists, materialChange } from '@/api/basic/material'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive({
    code: '',
    name: '',
    status: -1
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: materialLists,
    params: queryParams
})

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await materialDelete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

const handleSwitch = async (data: any) => {
    await materialChange({
        id: data.id,
    })
    feedback.msgSuccess('操作成功')
    getLists()
}

getLists()
</script>
