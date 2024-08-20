<template>
    <div class="order-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="款名">
                    <el-input
                        class="w-[280px]"
                        v-model="queryParams.name"
                        clearable
                        @keyup.enter="resetPage"
                    />
                </el-form-item>
                <el-form-item label="款号">
                    <el-input
                        class="w-[280px]"
                        v-model="queryParams.code"
                        clearable
                        @keyup.enter="resetPage"
                    />
                </el-form-item>
                <el-form-item label="年份">
                    <el-input
                        class="w-[180px]"
                        v-model="queryParams.year"
                        clearable
                        @keyup.enter="resetPage"
                    />
                </el-form-item>
                <el-form-item label="款号">
                    <el-input
                        class="w-[180px]"
                        v-model="queryParams.season"
                        clearable
                        @keyup.enter="resetPage"
                    />
                </el-form-item>
                <el-form-item label="款式状态">
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
                <router-link
                    v-perms="['cloth:order:add', 'cloth:order:add/edit']"
                    :to="{
                        path: getRoutePath('cloth:order:add/edit')
                    }"
                >
                    <el-button type="primary" class="mb-4">
                        <template #icon>
                            <icon name="el-icon-Plus" />
                        </template>
                        发布款式
                    </el-button>
                </router-link>
            </div>
            <el-table size="large" v-loading="pager.loading" :data="pager.lists">
                <el-table-column label="ID" prop="id" min-width="80" show-overflow-tooltip/>
                <el-table-column label="版样" min-width="100">
                    <template #default="{ row }">
                        <image-contain
                            v-if="row.picture"
                            :src="row.picture"
                            :width="80"
                            :height="80"
                            :preview-src-list="[row.picture]"
                            preview-teleported
                            fit="contain"
                        />
                    </template>
                </el-table-column>
                <el-table-column
                    label="款名"
                    prop="name"
                    min-width="160"
                    show-tooltip-when-overflow
                />
                <el-table-column label="款号" prop="code" min-width="100" />
                <el-table-column label="单价(元)" prop="unit_price" min-width="100" />
                <el-table-column label="年份" prop="year" min-width="80" />
                <el-table-column label="季节" prop="season" min-width="100" />
                <el-table-column label="颜色" prop="colors" min-width="80" />
                <el-table-column label="尺寸" prop="sizes" min-width="100" />
                <el-table-column label="状态" min-width="100">
                    <template #default="{ row }">
                        <el-switch
                            v-perms="['cloth:order:change']"
                            v-model="row.status"
                            :active-value="1"
                            :inactive-value="0"
                            @change="changeStatus(row.id)"
                        />
                    </template>
                </el-table-column>
                <el-table-column label="备注" prop="remark" min-width="120" show-overflow-tooltip/>
                <el-table-column label="发布时间" prop="create_time" min-width="120" show-overflow-tooltip/>
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['cloth:order:edit', 'cloth:order:publish']"
                            type="primary"
                            link
                        >
                            <router-link
                                :to="{
                                    path: getRoutePath('cloth:order:publish'),
                                    query: {
                                        id: row.id
                                    }
                                }"
                            >
                                编辑
                            </router-link>
                        </el-button>
                        <el-button
                            v-perms="['cloth:order:del']"
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
    </div>
</template>
<script lang="ts" setup name="orderLists">
import { orderLists, orderDelete, orderChange } from '@/api/cloth/order'
import { usePaging } from '@/hooks/usePaging'
import { getRoutePath } from '@/router'
import feedback from '@/utils/feedback'
const queryParams = reactive({
    name: '', // 款名
    code: '', // 款号
    year: '', // 年份
    season: '', // 季节
    status: -1 // 状态
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: orderLists,
    params: queryParams
})


const changeStatus = async (id: number) => {
    try {
        await orderChange({ id })
        feedback.msgSuccess('修改成功')
        getLists()
    } catch (error) {
        getLists()
    }
}

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await orderDelete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

onActivated(() => {
    getLists()
})

getLists()
</script>
