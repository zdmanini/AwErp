<template>
    <div class="style-edit">
        <el-card class="!border-none" shadow="never">
            <el-page-header content="款式编辑" @back="$router.back()" />
        </el-card>
        <el-card class="mt-4 !border-none" shadow="never">
            <el-form
                ref="formRef"
                class="ls-form"
                :model="formData"
                label-width="85px"
                :rules="rules"
            >
                <el-form-item label="款名" prop="name">
                    <el-input v-model="formData.name" placeholder="请输入款名" clearable></el-input>
                </el-form-item>
                <el-form-item label="编号" prop="code">
                    <el-input v-model="formData.code" placeholder="请输入编号" clearable></el-input>
                </el-form-item>
                <el-form-item label="款式图片" prop="picture">
                    <material-picker v-model="formData.picture" :limit="1" />
                </el-form-item>
                <el-form-item label="颜色" prop="colors">
                    <el-select
                        v-model="formData.colors"
                        multiple
                        filterable
                        allow-create
                        style="width: 100%"
                    >
                        <el-option
                            v-for="item in optionsData.dict['颜色']"
                            :key="item.id"
                            :label="item.name"
                            :value="item.name"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="尺码" prop="sizes">
                    <el-select
                        v-model="formData.sizes"
                        multiple
                        filterable
                        allow-create
                        style="width: 100%"
                    >
                        <el-option
                            v-for="item in optionsData.dict['尺码']"
                            :key="item.id"
                            :label="item.name"
                            :value="item.name"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="年份" prop="years">
                    <el-input v-model="formData.year" placeholder="请输入年份" clearable></el-input>
                </el-form-item>
                <el-form-item label="季节" prop="seasons">
                    <el-input
                        v-model="formData.season"
                        placeholder="请输入季节"
                        clearable
                    ></el-input>
                </el-form-item>
                <el-form-item label="单价" prop="unit_price">
                    <el-input-number
                        v-model="formData.unit_price"
                        placeholder="请输入单价"
                        clearable
                        :precision="2"
                    ></el-input-number>
                </el-form-item>
                <el-form-item label="工序" prop="procedures">
                    <zdm-form-table
                        ref="procedures"
                        v-model="formData.procedures"
                        :addTemplate="data.addTemplate"
                        drag-sort
                        placeholder="暂无数据"
                    >
                        <el-table-column prop="name" label="工序" width="180">
                            <template #default="scope">
                                <span style="display: none">{{
                                    (scope.row.sort = scope.$index + 1)
                                }}</span>
                                <el-select
                                    v-model="scope.row.name"
                                    placeholder="请选择"
                                    filtename
                                    allow-create
                                    filterable
                                >
                                    <el-option
                                        v-for="item in optionsData.dict['工序']"
                                        :key="item.id"
                                        :label="item.name"
                                        :value="item.name"
                                    ></el-option>
                                </el-select>
                            </template>
                        </el-table-column>
                        <el-table-column prop="price" label="单价" width="180">
                            <template #default="scope">
                                <el-input-number
                                    v-model="scope.row.price"
                                    placeholder="请输入单价"
                                    clearable
                                    :precision="2"
                                ></el-input-number>
                            </template>
                        </el-table-column>
                        <el-table-column prop="indiscriminately" label="不分扎上报">
                            <template #default="scope">
                                <el-switch v-model="scope.row.indiscriminately"></el-switch>
                            </template>
                        </el-table-column>
                        <el-table-column prop="is_completed" label="是否为完成工序">
                            <template #default="scope">
                                <el-switch v-model="scope.row.is_completed"></el-switch>
                            </template>
                        </el-table-column>
                    </zdm-form-table>
                </el-form-item>
                <el-form-item label="工艺" prop="crafts">
                    <table class="crafts">
                        <thead>
                            <tr>
                                <th style="width: 50px;">
                                  <el-button
                                      type="primary"
                                      @click="methods.addCrafts"
                                      size="small"
                                      circle
                                  >
                                    <icon name="el-icon-plus" />
                                  </el-button>
                                </th>
                                <th>名称</th>
                                <th>内容</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(item, index) in formData.crafts" :key="index">
                              <td></td>
                                <td>
                                    <el-select
                                        v-model="formData.crafts[index].name"
                                        placeholder="请选择"
                                        filterable
                                        allow-create
                                        style="width: 160px"
                                    >
                                        <el-option
                                            v-for="craft in optionsData.dict['工艺']"
                                            :key="craft.id"
                                            :label="craft.name"
                                            :value="craft.name"
                                        ></el-option>
                                    </el-select>
                                </td>
                                <td>
                                    <div v-if="!item.children || item.children.length == 0">
                                        <el-input
                                            v-model="formData.crafts[index].description"
                                            placeholder="请输入内容"
                                            clearable
                                            border-bottom
                                        ></el-input>
                                    </div>
                                    <div style="display: flex; flex-direction: column" v-else>
                                        <div
                                            style="display: flex"
                                            v-for="(chil, i) in formData.crafts[index].children"
                                            :key="i"
                                        >
                                            <div
                                                style="
                                                    display: block;
                                                    width: 130px;
                                                    float: left;
                                                    margin-right: 10px;
                                                "
                                            >
                                                <el-select
                                                    v-model="chil.name"
                                                    placeholder="请选择"
                                                    filterable
                                                    allow-create
                                                    style="width: 120px"
                                                >
                                                    <el-option
                                                        v-for="craft in optionsData.dict['工艺']"
                                                        :key="craft.id"
                                                        :label="craft.name"
                                                        :value="craft.name"
                                                    ></el-option>
                                                </el-select>
                                            </div>
                                            <div
                                                style="
                                                    display: block;
                                                    width: calc(100% - 140px);
                                                    float: left;
                                                "
                                            >
                                                <el-input
                                                    v-model="chil.description"
                                                    placeholder="请输入内容"
                                                    clearable
                                                    border-bottom
                                                ></el-input>
                                            </div>
                                        </div>
                                        <div
                                            style="
                                                display: flex;
                                                justify-content: flex-end;
                                                margin-top: 10px;
                                            "
                                        >
                                            <el-button
                                                type="danger"
                                                @click="methods.deleteCraftsChildren(index, i)"
                                                circle
                                            >
                                              <icon v-if="i == 0" name="el-icon-delete" />
                                              <icon v-else name="el-icon-minus" />
                                            </el-button>
                                        </div>
                                    </div>
                                    <!-- <button @click="addCraftsChildren(index)" v-if="item.children.length<4">添加</button> -->
                                    <el-button
                                        type="primary"
                                        @click="methods.addCraftsChildren(index)"
                                        circle
                                    >
                                      <icon name="el-icon-plus" />
                                    </el-button>
                                    <el-button
                                        type="danger"
                                        @click="methods.deleteCrafts(index)"
                                        circle
                                    >
                                      <icon v-if="index == 0" name="el-icon-delete" />
                                      <icon v-else name="el-icon-minus" />
                                    </el-button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input
                        type="textarea"
                        v-model="formData.remark"
                        placeholder="请输入备注"
                        clearable
                    ></el-input>
                </el-form-item>
            </el-form>
        </el-card>
        <footer-btns>
            <el-button type="primary" @click="handleSave">保存</el-button>
        </footer-btns>
    </div>
</template>

<script lang="ts" setup name="styleListsEdit">
import { getCurrentInstance } from 'vue'
import type { FormInstance } from 'element-plus'
import feedback from '@/utils/feedback'
import { useDictOptions } from '@/hooks/useDictOptions'
import { styleDetail, styleEdit, styleAdd } from '@/api/cloth/style'
import useMultipleTabs from '@/hooks/useMultipleTabs'
import { infoQuery } from '@/api/basic/info'
import zdmFormTable from "@/components/zdmFormTable/index.vue";
const instance = getCurrentInstance()
const route = useRoute()
const router = useRouter()
const formData = reactive({
    id: '',
    name: '', // 款名
    code: '', // 款号
    picture: '', // 款式图片
    colors: [], // 颜色
    sizes: [], // 尺码
    year: '', // 年份
    season: '', // 季节
    status: 1, // 状态
    unit_price: 0.0, // 单价
    procedures: [], // 工序
    crafts: [], // 工艺
    remark: '' // 备注
})
const procedures = ref([])

const { removeTab } = useMultipleTabs()
const formRef = shallowRef<FormInstance>()
const rules = reactive({
    title: [{ required: true, message: '请输入款式标题', trigger: 'blur' }],
    cid: [{ required: true, message: '请选择款式栏目', trigger: 'blur' }]
})

const data = reactive({
    addTemplate: {
        name: '',
        sort: 0,
        price: 0,
        indiscriminately: false,
        is_completed: false
    }
})

const methods = {
    deleteCrafts: (i) => {
        formData.crafts.splice(i, 1)
    },
    addCrafts:() => {
        formData.crafts.push({
            name: '',
            sort: 0,
            description: '',
            children: []
        })
    },
    addCraftsChildren:(i) => {
        formData.crafts[i].children.push({
            name: '',
            content: '',
            sort: 0
        })
    },
    deleteCraftsChildren:(i, j) => {
        formData.crafts[i].children.splice(j, 1)
    }
}

const getDetails = async () => {
    const data = await styleDetail({
        id: route.query.id
    })
    Object.keys(formData).forEach((key) => {
        //@ts-ignore
        formData[key] = data[key]
    })
}

const { optionsData } = useDictOptions<{
    dict: any[]
}>({
    dict: {
        api: infoQuery,
        params: {
            type: '颜色,尺码,工序,工艺'
        }
    }
})

const handleSave = async () => {
    await formRef.value?.validate()
    if (route.query.id) {
        await styleEdit(formData)
    } else {
        await styleAdd(formData)
    }
    feedback.msgSuccess('操作成功')
    removeTab()
    router.back()
}

route.query.id && getDetails()
</script>

<style scoped>
.el-steps:deep(.is-finish) .el-step__line {background: var(--el-color-primary);}

.contains {
  width: auto;
  border-collapse: collapse;
}
.contains th {
  border-bottom: 1px solid #ddd;
  padding: 5px 10px;
  background-color: #f5f5f5;
}
.contains td {
  min-width: 100px;
  border-bottom: 1px solid #ddd;
  padding: 5px 10px;
  text-align: center;
}
.contains tbody tr:nth-child(2n) {
  background-color: #f5f5f5;
}
.contains .color {
  text-align: left;
}

table.crafts{
  width: 100%;
  border-collapse: collapse;
}
table.crafts th {
  border-bottom: 1px solid #ddd;
  padding: 5px 10px;
  background-color: #f5f5f5;
}
table.crafts td {
  min-width: 100px;
  border-bottom: 1px solid #ddd;
  padding: 5px 10px;
  text-align: center;
}
table.crafts tbody tr:nth-child(2n) {
  background-color: #f5f5f5;
}
</style>