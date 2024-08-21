<template>
    <div class="style-edit">
        <el-card class="!border-none" shadow="never">
            <el-page-header content="款式编辑" @back="$router.back()" />
        </el-card>
        <el-card class="mt-4 !border-none" shadow="never">
            <el-steps :active="state.active" align-center style="margin-bottom: 20px">
                <el-step title="基础"></el-step>
                <el-step title="款式"></el-step>
                <el-step title="工序"></el-step>
                <el-step title="物料"></el-step>
                <el-step title="完成"></el-step>
            </el-steps>
            <el-row>
                <el-col :lg="{ span: 20, offset: 3 }">
                    <el-form
                        v-if="state.active == 0"
                        ref="stepForm_0"
                        :model="formData"
                        :rules="rules"
                        label-position="top"
                    >
                        <el-form-item label="合同名称" prop="name">
                            <el-input
                                v-model="formData.name"
                                clearable
                                style="width: 360px"
                            ></el-input>
                        </el-form-item>
                        <el-form-item label="合同编号" prop="code">
                            <el-input v-model="formData.code" clearable style="width: 360px">
                                <template #append>
                                    <el-button
                                        type="primary"
                                        @click="methods.generateCode"
                                        circle
                                        >
                                      <icon name="el-icon-refresh" />
                                      自动生成</el-button
                                    >
                                </template>
                            </el-input>
                        </el-form-item>
                        <el-form-item label="客户名称" prop="customer">
                            <el-select
                                v-model="formData.customer"
                                placeholder="请输入客户名称"
                                clearable
                                filterable
                                allow-create
                                style="width: 360px"
                            >
                                <el-option
                                    v-for="item in state.customers"
                                    :key="item.id"
                                    :label="item.name"
                                    :value="item.name"
                                ></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="交货日期" prop="delivery_date">
                            <el-date-picker
                                v-model="formData.delivery_date"
                                type="date"
                                placeholder="选择日期"
                                style="width: 360px"
                                value-format="YYYY-MM-DD"
                            >
                            </el-date-picker>
                        </el-form-item>
                        <el-form-item label="订单类型" prop="order_type">
                            <el-autocomplete
                                v-model="formData.order_type"
                                :fetch-suggestions="methods.querySearchOrderType"
                                placeholder="请输入订单类型"
                                clearable
                                style="width: 360px"
                            ></el-autocomplete>
                        </el-form-item>
                        <el-form-item label="业务员" prop="salesman">
                            <el-autocomplete
                                v-model="formData.salesman"
                                :fetch-suggestions="methods.querySearchSalesman"
                                placeholder="请输入业务员"
                                clearable
                                style="width: 360px"
                            ></el-autocomplete>
                        </el-form-item>
                        <el-form-item label="备注" prop="remark">
                            <el-input
                                v-model="formData.remark"
                                clearable
                                type="textarea"
                                rows="3"
                                style="width: 360px"
                            ></el-input>
                        </el-form-item>
                    </el-form>
                    <el-form
                        v-if="state.active == 1"
                        ref="stepForm_1"
                        :model="formData"
                        :rules="rules"
                        label-position="top"
                    >
                        <el-form-item label="款号" prop="cloth.code">
                            <el-select
                                v-model="formData.cloth.code"
                                placeholder="请输入款号"
                                :remote-method="methods.querySearchCloth"
                                clearable
                                filterable
                                @change="methods.handleClothCodeChange"
                                style="width: 360px"
                                allow-create
                            >
                                <el-option
                                    v-for="item in state.cloths"
                                    :key="item.id"
                                    :label="item.code"
                                    :value="item.code"
                                ></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="款名" prop="cloth.name">
                            <el-select
                                v-model="formData.cloth.name"
                                placeholder="请输入款名"
                                :remote-method="methods.querySearchCloth"
                                clearable
                                filterable
                                @change="methods.handleClothNameChange"
                                allow-create
                                style="width: 360px"
                            >
                                <el-option
                                    v-for="item in cloths"
                                    :key="item.id"
                                    :label="item.name"
                                    :value="item.name"
                                ></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="图片" prop="cloth.picture">
                            <material-picker v-model="formData.cloth.picture" :limit="1" />
                        </el-form-item>
                        <el-form-item label="颜色" prop="cloth.colors">
                            <el-select
                                v-model="formData.cloth.colors"
                                multiple
                                filterable
                                allow-create
                                clearable
                                style="width: 360px"
                            >
                                <el-option
                                    v-for="item in state.colors"
                                    :key="item.id"
                                    :label="item.name"
                                    :value="item.name"
                                ></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="尺码" prop="cloth.sizes">
                            <el-select
                                v-model="formData.cloth.sizes"
                                multiple
                                filterable
                                allow-create
                                clearable
                                style="width: 360px"
                            >
                                <el-option
                                    v-for="item in state.sizes"
                                    :key="item.id"
                                    :label="item.name"
                                    :value="item.name"
                                ></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="年份" prop="cloth.year">
                            <el-input
                                v-model="formData.cloth.year"
                                clearable
                                style="width: 360px"
                            ></el-input>
                        </el-form-item>
                        <el-form-item label="季节" prop="cloth.season">
                            <el-input
                                v-model="formData.cloth.season"
                                clearable
                                style="width: 360px"
                            ></el-input>
                        </el-form-item>
                        <el-form-item label="单价" prop="cloth.unit_price">
                            <el-input-number
                                v-model="formData.cloth.unit_price"
                                clearable
                                controls-position="right"
                                :precision="2"
                                :min="0"
                            ></el-input-number>
                            <span style="margin-left: 10px"> 元 </span>
                        </el-form-item>
                        <el-form-item
                            prop="contains"
                            label="数量"
                            v-if="
                                formData.cloth.sizes.length > 0 && formData.cloth.colors.length > 0
                            "
                        >
                            <table class="contains">
                                <thead>
                                    <tr>
                                        <th>颜色</th>
                                        <th v-for="size in formData.cloth.sizes" :key="size">
                                            {{ size }}
                                        </th>
                                        <th>合计</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr v-for="color in formData.cloth.colors" :key="color">
                                        <td class="color">{{ color }}</td>
                                        <td v-for="size in formData.cloth.sizes" :key="size">
                                            <el-input-number
                                                v-model="
                                                    formData.contains.find(
                                                        (item) =>
                                                            item.color == color && item.size == size
                                                    ).nums
                                                "
                                                :min="0"
                                                :precision="0"
                                            ></el-input-number>
                                        </td>
                                        <td>
                                            {{
                                                formData.contains
                                                    .filter((item) => item.color == color)
                                                    .reduce((total, item) => total + item.nums, 0)
                                            }}
                                        </td>
                                    </tr>
                                </tbody>
                                <tfoot>
                                    <tr>
                                        <td>合计</td>
                                        <td v-for="size in formData.cloth.sizes" :key="size">
                                            {{
                                                formData.contains
                                                    .filter((item) => item.size == size)
                                                    .reduce((total, item) => total + item.nums, 0)
                                            }}
                                        </td>
                                        <td>
                                            {{
                                                formData.contains.reduce(
                                                    (total, item) => total + item.nums,
                                                    0
                                                )
                                            }}
                                        </td>
                                    </tr>
                                </tfoot>
                            </table>
                        </el-form-item>
                        <el-form-item label="金额" prop="total_price">
                            <el-input-number
                                v-model="formData.total_price"
                                clearable
                                controls-position="right"
                                :precision="2"
                                :min="0.01"
                            ></el-input-number>
                            <span style="margin-left: 10px"> 元 </span>
                        </el-form-item>
                    </el-form>

                    <el-form
                        v-if="state.active == 2"
                        ref="stepForm_2"
                        :model="formData"
                        :rules="rules"
                        label-position="top"
                    >
                        <el-form-item label="工序" prop="procedure">
                            <zdm-form-table
                                ref="procedure"
                                v-model="formData.procedure"
                                :addTemplate="state.addProcedureTemplate"
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
                                            filterable
                                            allow-create
                                        >
                                            <el-option
                                                v-for="item in procedures"
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
                    </el-form>
                    <el-form
                        v-if="state.active == 3"
                        ref="stepForm_3"
                        :model="formData"
                        :rules="rules"
                        label-position="top"
                    >
                        <el-form-item label="物料" prop="consumption">
                            <table>
                                <thead>
                                    <th>部位</th>
                                    <th>物料名称</th>
                                    <th>单价/单位</th>
                                    <th>单件用量</th>
                                    <th>全部颜色</th>
                                    <th v-for="size in formData.cloth.sizes" :key="size">
                                        {{ size }}
                                    </th>
                                    <th>全部尺码</th>
                                    <th v-for="color in formData.cloth.colors" :key="color">
                                        {{ color }}
                                    </th>
                                    <th>操作</th>
                                </thead>
                                <tbody>
                                    <tr v-for="(item, index) in formData.consumption" :key="index">
                                        <td>
                                            <el-input
                                                v-model="item.part"
                                                placeholder="请输入部位"
                                                clearable
                                            ></el-input>
                                        </td>
                                        <td>
                                            <el-input
                                                v-model="item.name"
                                                placeholder="请输入物料名称"
                                                clearable
                                            ></el-input>
                                        </td>
                                        <td>
                                            <el-input-number
                                                v-model="item.unit_price"
                                                placeholder="请输入单价"
                                                clearable
                                                :precision="2"
                                            ></el-input-number>
                                        </td>
                                        <td>
                                            <el-input-number
                                                v-model="item.nums"
                                                placeholder="请输入单件用量"
                                                clearable
                                                :precision="2"
                                            ></el-input-number>
                                        </td>
                                        <td>
                                            <el-switch
                                                @change="allColor(item, index)"
                                                v-model="allColors[index]"
                                            ></el-switch>
                                        </td>
                                        <td v-for="color in formData.cloth.colors" :key="color">
                                            <el-switch v-model="item[color]"></el-switch>
                                        </td>
                                        <td>
                                            <el-switch
                                                @change="allSize(item, index)"
                                                v-model="allSizes[index]"
                                            ></el-switch>
                                        </td>

                                        <td v-for="size in formData.cloth.sizes" :key="size">
                                            <el-switch v-model="item[size]"></el-switch>
                                        </td>
                                        <td>
                                            <el-button
                                                type="danger"
                                                plain
                                                icon="el-icon-delete"
                                                @click="formData.consumption.splice(index, 1)"
                                            ></el-button>
                                        </td>
                                    </tr>
                                </tbody>
                                <tfoot>
                                    <tr>
                                        <td colspan="11">
                                            <el-button
                                                type="primary"
                                                plain
                                                icon="el-icon-plus"
                                                @click="
                                                    formData.consumption.push(
                                                        JSON.parse(JSON.stringify(addConsumption))
                                                    )
                                                "
                                                >添加</el-button
                                            >
                                        </td>
                                    </tr>
                                </tfoot>
                            </table>
                        </el-form-item>
                    </el-form>
                    <div v-if="state.active == 4">
                        <el-result icon="success" title="操作成功" sub-title="您的操作已完成！">
                            <template #extra>
                                <el-button type="primary" @click="methods.again">再来一单</el-button>
                                <el-button @click="$router.go(-1)">返回列表</el-button>
                            </template>
                        </el-result>
                    </div>
                </el-col>
            </el-row>
        </el-card>
        <footer-btns>
            <el-button
                v-if="state.active > 0 && state.active < 4"
                @click="methods.pre"
                :disabled="state.submitLoading"
                >上一步</el-button
            >
            <el-button v-if="state.active < 3" type="primary" @click="methods.next"
                >下一步</el-button
            >
            <el-button v-if="state.active == 3" type="primary" @click="handleSave">保存</el-button>
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
import zdmFormTable from '@/components/zdmFormTable/index.vue'
const route = useRoute()
const router = useRouter()
const instance = getCurrentInstance()
const formData = reactive({
    id: '',
    // 基础
    name: '', // 合同名称
    code: '', // 合同编号
    customer: '', // 客户名称
    delivery_date: '', // 交货日期
    order_type: '', // 订单类型
    salesman: '', // 业务员
    remark: '', // 备注
    // 款式
    cloth: {
        name: '', // 款名
        code: '', // 款号
        picture: '', // 图片
        colors: [], // 颜色
        sizes: [], // 尺码
        year: '', // 年份
        season: '', // 季节
        unit_price: 0 // 单价
    },
    total: 0, // 总数量
    total_price: 0, // 总金额
    contains: [],
    consumption: [], // 消耗
    // 工序
    procedure: []
})
const procedures = ref([])
const state = reactive({
    active: 0,
    submitLoading: false,
    sizes: [],
    colors: [],
    crafts: [],
    procedures: [],
    order_types: [],
    salesmans: [],
    customers: [],
    cloths: [],
    addContainsTemplate: {
        size: '',
        color: '',
        nums: 0
    },
    addProcedureTemplate: {
        name: '',
        sort: 0,
        price: 0,
        indiscriminately: false,
        is_completed: false
    },
    addConsumption: {
        name: '',
        part: '',
        unit: '',
        nums: 0,
        unit_price: 0
    },
    contains: [],
    consumption: [],
    allColors: [],
    allSizes: []
})
const { removeTab } = useMultipleTabs()
const formRef = shallowRef<FormInstance>()
const rules = reactive({
  name: [{ required: true, message: '请输入合同名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入合同编号', trigger: 'blur' }],
  customer: [{ required: true, message: '请输入客户名称', trigger: 'blur' }],
  delivery_date: [{ required: true, message: '请选择交货日期', trigger: 'blur' }],
  order_type: [{ required: true, message: '请选择订单类型', trigger: 'blur' }],
  salesman: [{ required: true, message: '请选择业务员', trigger: 'blur' }],
  remark: [],
  // 款式
  'cloth.name': [{ required: true, message: '请输入款名', trigger: 'blur' }],
  'cloth.code': [{ required: true, message: '请输入款号', trigger: 'blur' }],
  'cloth.picture': [],
  'cloth.colors': [{ required: true, message: '请选择颜色', trigger: 'blur' }],
  'cloth.sizes': [{ required: true, message: '请选择尺码', trigger: 'blur' }],
  'cloth.year': [{ required: true, message: '请输入年份', trigger: 'blur' }],
  'cloth.season': [{ required: true, message: '请输入季节', trigger: 'blur' }],
  'cloth.unit_price': [{ required: true, message: '请输入单价', trigger: 'blur' }],
  // 工序
  procedure: [{ required: true, message: '请添加工序', trigger: 'blur' }]
})

const methods = {
    allColor: (item, i) => {
        if (state.allColor[i]) {
            formData.cloth.colors.forEach((color) => {
                formData.consumption[i][color] = true
            })
            state.allSize[i] = true
        } else {
            formData.cloth.colors.forEach((color) => {
                formData.consumption[i][color] = false
            })
            state.allSize[i] = false
        }
        return state.allColor[i]
    },
    allSize: (item, i) => {
        if (state.allSize[i]) {
            formData.cloth.sizes.forEach((size) => {
                formData.consumption[i][size] = true
            })
            state.allColor[i] = true
        } else {
            formData.cloth.sizes.forEach((size) => {
                formData.consumption[i][size] = false
            })
            state.allColor[i] = false
        }
        return state.allSize[i]
    },
    addConsumptionBtn: () => {
        const consumption = JSON.parse(JSON.stringify(state.addConsumption))
        formData.consumption.push(consumption)
    },
    handleClothCodeChange: (val) => {
        const cloth = state.cloths.find((item) => item.code == val)
        if (cloth) {
            formData.cloth = cloth
            formDataprocedure = cloth.procedures
        }
    },
    handleClothNameChange: (val) => {
        const cloth = state.cloths.find((item) => item.name == val)
        if (cloth) {
            formData.cloth = clothformData.procedure = cloth.procedures
        }
    },
    generateCode: () => {
        const date = `${new Date().getFullYear()}${
            new Date().getMonth() + 1
        }${new Date().getDate()}`.slice(-8)
        formData.code = `AWS${date}${Math.random().toString().slice(-6)}`
    },
    createFilterCustomer: (queryString) => {
        return (item) => {
            return item.name.toLowerCase().indexOf(queryString.toLowerCase()) === 0
        }
    },
    querySearchSalesman: (queryString, cb) => {
        const results = queryString
            ? state.salesmans.filter(methods.createFilter(queryString))
            : state.salesmans
        cb(results)
    },
    querySearchOrderType: (queryString, cb) => {
        const results = queryString
            ? state.order_types.filter(methods.createFilter(queryString))
            : state.order_types
        cb(results)
    },
    createFilter: (queryString) => {
        return (item) => {
            return item.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
        }
    },
    next: () => {
        const formName = `stepForm_${state.active}`
        instance.refs[formName].validate((valid) => {
            if (valid) {
                state.active++
            } else {
                return false
            }
        })
    },
    pre: () => {
        state.active--
    },
    again() {
        state.active = 0
        formData.code = ''
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

const getInfo = async () => {
    const data = await infoQuery({
        type: '颜色,尺码,工序,工艺,订单类型'
    })
    state.sizes = data['尺码']
    state.colors = data['颜色']
    state.procedures = data['工序']
    state.crafts = data['工艺']
    state.order_types = data['订单类型']
}

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
watch(
    () => formData.contains,
    (val, oldVal) => {
        let total = 0
        let total_price = 0
        val.forEach((item) => {
            total += item.nums
            total_price += item.nums * formData.cloth.unit_price
        })
        formData.total = total
        formData.total_price = total_price
    }
)

watch(
    () => formData.cloth.unit_price,
    (val, oldVal) => {
        let total_price = 0
        formData.contains.forEach((item) => {
            total_price += item.nums * val
        })
        formData.total_price = total_price
    }
)

watch(
    () => formData.total,
    (val, oldVal) => {
        let total_price = 0
        formData.contains.forEach((item) => {
            total_price += item.nums * formData.cloth.unit_price
        })
        formData.total_price = total_price
    }
)

watch(
    () => formData.cloth.colors,
    (val, oldVal) => {
        const oldContains = formData.contains
        const contains = []
        val.forEach((color) => {
            formData.cloth.sizes.forEach((size) => {
                contains.push({
                    color: color,
                    size: size,
                    nums: 0
                })
            })
        })
        contains.forEach((item) => {
            const oldItem = oldContains.find(
                (oldItem) => oldItem.color == item.color && oldItem.size == item.size
            )
            if (oldItem) {
                item.nums = oldItem.nums
            }
        })
        formData.contains = contains
    }
)

watch(
    () => formData.cloth.sizes,
    (val, oldVal) => {
        const oldContains = formData.contains
        const contains = []
        formData.cloth.colors.forEach((color) => {
            val.forEach((size) => {
                contains.push({
                    color: color,
                    size: size,
                    nums: 0
                })
            })
        })
        contains.forEach((item) => {
            const oldItem = oldContains.find(
                (oldItem) => oldItem.color == item.color && oldItem.size == item.size
            )
            if (oldItem) {
                item.nums = oldItem.nums
            }
        })
        formData.contains = contains
    }
)

route.query.id && getDetails()
getInfo()
</script>

<style scoped>
.el-steps:deep(.is-finish) .el-step__line {
    background: var(--el-color-primary);
}

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

table.crafts {
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
