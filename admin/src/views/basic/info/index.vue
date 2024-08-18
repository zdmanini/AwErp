<template>
    <div class="dict-type">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" inline>
                <el-form-item label="字典名称">
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
            <el-row :gutter="10">
                <el-col :span="6">
                    <el-button type="primary" @click="handleAddType">新增分类</el-button>
                    <el-scrollbar height="600px">
                        <el-tree
                            :data="optionsData.typeLists"
                            node-key="id"
                            class="mt-4"
                            default-expand-all
                            :expand-on-click-node="false"
                            :props="{ label: 'name', children: 'children' }"
                        >
                            <template #default="{ node, data }">
                                <div class="custom-tree-node">
                                    <div class="name" @click="handleNodeClick(data)">
                                        {{ node.label }}
                                    </div>
                                    <div class="operator">
                                        <a
                                            href="javascript:void(0);"
                                            class="primary"
                                            v-perms="['basic:info:add']"
                                            @click="append(data)"
                                            >新增</a
                                        >
                                        <a
                                            href="javascript:void(0);"
                                            class="success"
                                            v-perms="['basic:info:edit']"
                                            @click="edit(data)"
                                            >编辑</a
                                        >
                                        <a
                                            href="javascript:void(0);"
                                            class="danger"
                                            v-perms="['basic:info:del']"
                                            @click="remove(node, data)"
                                            >删除</a
                                        >
                                    </div>
                                </div>
                            </template>
                        </el-tree>
                    </el-scrollbar>
                </el-col>
                <el-col :span="18">
                    <div>
                        <el-button v-perms="['basic:info:add']" type="primary" @click="handleAdd">
                            <template #icon>
                                <icon name="el-icon-Plus" />
                            </template>
                            新增
                        </el-button>
                        <el-button
                            v-perms="['basic:info:del']"
                            :disabled="!selectData.length"
                            type="danger"
                            @click="handleDelete(selectData)"
                        >
                            <template #icon>
                                <icon name="el-icon-Delete" />
                            </template>
                            删除
                        </el-button>
                    </div>
                    <div class="mt-4" v-loading="pager.loading">
                        <div>
                            <el-table
                                :data="pager.lists"
                                size="large"
                                @selection-change="handleSelectionChange"
                            >
                                <el-table-column type="selection" width="55" />
                                <!--                                <el-table-column label="ID" prop="id" />-->
                                <el-table-column label="字典名称" prop="name" min-width="120" />
                                <el-table-column label="字典类型" prop="type" min-width="120" />
                                <el-table-column label="状态">
                                    <template v-slot="{ row }">
                                        <el-switch
                                            v-model="row.status"
                                            :active-value="1"
                                            :inactive-value="0"
                                            @change="handleSwitch(row)"
                                        />
                                    </template>
                                </el-table-column>
                                <el-table-column
                                    label="备注"
                                    prop="dictRemark"
                                    show-tooltip-when-overflow
                                />
                                <el-table-column
                                    label="创建时间"
                                    prop="create_time"
                                    min-width="180"
                                />
                                <el-table-column label="操作" width="190" fixed="right">
                                    <template #default="{ row }">
                                        <el-button
                                            v-perms="['basic:info::edit']"
                                            link
                                            type="primary"
                                            @click="handleEdit(row)"
                                        >
                                            编辑
                                        </el-button>
                                        <!--                                        <el-button-->
                                        <!--                                            v-perms="['setting:dict:data:list']"-->
                                        <!--                                            type="primary"-->
                                        <!--                                            link-->
                                        <!--                                        >-->
                                        <!--                                            <router-link-->
                                        <!--                                                :to="{-->
                                        <!--                                                    path: getRoutePath('setting:dict:data:list'),-->
                                        <!--                                                    query: {-->
                                        <!--                                                        type: row.type-->
                                        <!--                                                    }-->
                                        <!--                                                }"-->
                                        <!--                                            >-->
                                        <!--                                                数据管理-->
                                        <!--                                            </router-link>-->
                                        <!--                                        </el-button>-->
                                        <el-button
                                            v-perms="['basic:info::del']"
                                            link
                                            type="danger"
                                            @click="handleDelete(row.id)"
                                        >
                                            删除
                                        </el-button>
                                    </template>
                                </el-table-column>
                            </el-table>
                        </div>
                        <div class="flex justify-end mt-4">
                            <pagination v-model="pager" @change="getLists" />
                        </div>
                    </div>
                </el-col>
            </el-row>
        </el-card>
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
        <type-popup v-if="showType" ref="typeRef" @success="getInfo" @close="showType = false" />
    </div>
</template>

<script lang="ts" setup name="type">
import { typeDelete, typeLists, itemLists, infoDelete, infoChange } from '@/api/basic/info'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
import TypePopup from './type.vue'
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const typeRef = shallowRef<InstanceType<typeof TypePopup>>()
const showType = ref(false)
const showEdit = ref(false)
const queryParams = reactive({
    name: '',
    type: '',
    status: -1
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: itemLists,
    params: queryParams
})

const handleNodeClick = (data: any) => {
    queryParams.type = data.name
    resetPage()
}

const optionsData = reactive({
    typeLists: []
})

const getInfo = async () => {
    const res = await typeLists()
    console.log(res)
    optionsData.typeLists = res
}

const selectData = ref<any[]>([])

const handleSelectionChange = (val: any[]) => {
    selectData.value = val.map(({ id }) => id)
}

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
    if (queryParams.type) {
        const params = {
            type: queryParams.type
        }
        editRef.value?.setFormData(params)
    }
}

const handleAddType = async () => {
    showType.value = true
    await nextTick()
    typeRef.value?.open('add')
}

const append = async (data: any) => {
    showType.value = true
    await nextTick()
    typeRef.value?.open('add')
    const params = {
        parent: data.name
    }
    typeRef.value?.setFormData(params)
}

const edit = async (data: any) => {
    showType.value = true
    await nextTick()
    typeRef.value?.open('edit')
    typeRef.value?.setFormData(data)
}

const remove = async (node: any, data: any) => {
    await feedback.confirm('确定要删除？')
    await typeDelete({ id: data.id })
    feedback.msgSuccess('删除成功')
    getInfo()
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.setFormData(data)
}

// 删除角色
const handleDelete = async (id: string) => {
    await feedback.confirm('确定要删除？')
    await infoDelete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

const handleSwitch = async (data: any) => {
    await infoChange({
        id: data.id
    })
    feedback.msgSuccess('操作成功')
    getLists()
}
getInfo()
getLists()
</script>

<style>
.custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
}
.custom-tree-node .name {
    cursor: pointer;
    width: 70%;
}

.custom-tree-node .operator {
    flex-shrink: 0;
    margin-left: 10px;
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.custom-tree-node .operator a {
    padding: 0 5px;
}
.custom-tree-node .operator a.primary {
    color: var(--el-color-primary);
}
.custom-tree-node .operator a.success {
    color: var(--el-color-success);
}
.custom-tree-node .operator a.danger {
    color: var(--el-color-danger);
}
.custom-tree-node .operator a.primary:hover {
    color: var(--el-color-primary-light-5);
}
.custom-tree-node .operator a.success:hover {
    color: var(--el-color-success-light-5);
}
.custom-tree-node .operator a.danger:hover {
    color: var(--el-color-danger-light-5);
}
</style>
