<template>
	<sc-page-header :title="id?'编辑订单':'新增订单'" icon="el-icon-burger" :descripttion="id?form.code:'新增订单'">
		<template #main>
			<el-button type="primary" @click="$router.go(-1)">返回</el-button>
		</template>
	</sc-page-header>

	<el-main>
		<el-card shadow="never">
			<el-steps :active="active" align-center style="margin-bottom: 20px;">
				<el-step title="基础"></el-step>
				<el-step title="款式"></el-step>
				<el-step title="工序"></el-step>
				<el-step title="物料"></el-step>
				<el-step title="完成"></el-step>
			</el-steps>
			<el-row>
				<el-col :lg="{span: 20, offset: 3}">
					<el-form v-if="active==0" ref="stepForm_0" :model="form" :rules="rules" label-position="top">
						<el-form-item label="合同名称" prop="name">
							<el-input v-model="form.name" clearable style="width: 360px"></el-input>
						</el-form-item>
						<el-form-item label="合同编号" prop="code">
							<el-input v-model="form.code" clearable style="width: 360px">
								<template #append>
									<el-button type="primary" icon="el-icon-refresh" @click="generateCode" circle>自动生成</el-button>
								</template>
							</el-input>
						</el-form-item>
						<el-form-item label="客户名称" prop="customer">
							<el-select v-model="form.customer" placeholder="请输入客户名称" clearable filterable allow-create style="width: 360px">
								<el-option v-for="item in customers" :key="item.id" :label="item.name" :value="item.name"></el-option>
							</el-select>
						</el-form-item>
						<el-form-item label="交货日期" prop="delivery_date">
							<el-date-picker v-model="form.delivery_date" type="date" placeholder="选择日期" style="width: 360px" value-format="YYYY-MM-DD"> </el-date-picker>
						</el-form-item>
						<el-form-item label="订单类型" prop="order_type">
							<el-autocomplete v-model="form.order_type" :fetch-suggestions="querySearchOrderType" placeholder="请输入订单类型" clearable style="width: 360px"></el-autocomplete>
						</el-form-item>
						<el-form-item label="业务员" prop="salesman">
							<el-autocomplete v-model="form.salesman" :fetch-suggestions="querySearchSalesman" placeholder="请输入业务员" clearable style="width: 360px"></el-autocomplete>
						</el-form-item>
						<el-form-item label="备注" prop="remark">
							<el-input v-model="form.remark" clearable type="textarea" rows="3" style="width: 360px"></el-input>
						</el-form-item>
					</el-form>
					<el-form v-if="active==1" ref="stepForm_1" :model="form" :rules="rules" label-position="top">
						<el-form-item label="款号" prop="cloth.code">
							<el-select v-model="form.cloth.code" placeholder="请输入款号"
							:remote-method="querySearchCloth"
							 clearable 
							 filterable 
							 @change="handleClothCodeChange"
							 style="width: 360px"
							 allow-create>
								<el-option v-for="item in cloths" :key="item.id" :label="item.code" :value="item.code"></el-option>
							</el-select>
						</el-form-item>
						<el-form-item label="款名" prop="cloth.name">
							<el-select v-model="form.cloth.name" placeholder="请输入款名"
							:remote-method="querySearchCloth"
							 clearable 
							 filterable 
							 @change="handleClothNameChange"
							 allow-create
							 style="width: 360px"
							 >
								<el-option v-for="item in cloths" :key="item.id" :label="item.name" :value="item.name"></el-option>
							</el-select>
						</el-form-item>
						<el-form-item label="图片" prop="cloth.picture">
							<sc-upload v-model="form.cloth.picture" title="上传图片"></sc-upload>
						</el-form-item>
						<el-form-item label="颜色" prop="cloth.colors">
							<el-select v-model="form.cloth.colors" multiple filterable allow-create clearable style="width: 360px">
								<el-option v-for="item in colors" :key="item.value" :label="item.label" :value="item.value"></el-option>
							</el-select>
						</el-form-item>
						<el-form-item label="尺码" prop="cloth.sizes">
							<el-select v-model="form.cloth.sizes" multiple filterable allow-create clearable style="width: 360px">
								<el-option v-for="item in sizes" :key="item.value" :label="item.label" :value="item.value"></el-option>
							</el-select>
						</el-form-item>
						<el-form-item label="年份" prop="cloth.year">
							<el-input v-model="form.cloth.year" clearable style="width: 360px"></el-input>
						</el-form-item>
						<el-form-item label="季节" prop="cloth.season">
							<el-input v-model="form.cloth.season" clearable style="width: 360px"></el-input>
						</el-form-item>
						<el-form-item label="单价" prop="cloth.unit_price">
							<el-input-number v-model="form.cloth.unit_price" clearable controls-position="right" :precision="2" :min="0"></el-input-number>
							<span style="margin-left: 10px;">
								元
							</span>
						</el-form-item>
						<el-form-item prop="contains" label="数量" v-if="form.cloth.sizes.length > 0 && form.cloth.colors.length > 0">
							<table class="contains">
						      <thead>
						        <tr>
						          <th>颜色</th>
						          <th v-for="size in form.cloth.sizes" :key="size">{{ size }}</th>
						          <th>合计</th>
						        </tr>
						      </thead>
						      <tbody>
						        <tr v-for="color in form.cloth.colors" :key="color">
								  <td class="color">{{ color }}</td>
								  <td v-for="size in form.cloth.sizes" :key="size">
									<el-input-number v-model="form.contains.find(item=>item.color==color&&item.size==size).nums" :min="0" :precision="0"></el-input-number>
								  </td>
								  <td>{{ form.contains.filter(item=>item.color==color).reduce((total,item)=>total+item.nums,0) }}</td>
								  		</tr>
						      </tbody>
							  <tfoot>
								<tr>
								  <td>合计</td>
								  <td v-for="size in form.cloth.sizes" :key="size">{{ form.contains.filter(item=>item.size==size).reduce((total,item)=>total+item.nums,0) }}</td>
								  <td>{{ form.contains.reduce((total,item)=>total+item.nums,0) }}</td>
								</tr>
							  </tfoot>
						    </table>		
						</el-form-item>
						<el-form-item label="金额" prop="total_price">
							<el-input-number v-model="form.total_price" clearable controls-position="right" :precision="2" :min="0.01"></el-input-number>
							<span style="margin-left: 10px;">
								元
							</span>
						</el-form-item>
					</el-form>
					
					<el-form v-if="active==2" ref="stepForm_2" :model="form" :rules="rules" label-position="top">
						<el-form-item label="工序" prop="procedure">
							<sc-form-table ref="procedure" v-model="form.procedure" :addTemplate="addProcedureTemplate" placeholder="暂无数据">
								<el-table-column prop="name" label="工序" width="180">
										<template #default="scope">
											<span style="display: none;" >{{ scope.row.sort = scope.$index+1 }}</span>
											<el-select v-model="scope.row.name" placeholder="请选择" filterable allow-create>
												<el-option v-for="item in procedures" :key="item.value" :label="item.label" :value="item.value"></el-option>
											</el-select>
										</template>
								</el-table-column>
								<el-table-column prop="price" label="单价" width="180">
									<template #default="scope">
										<el-input-number v-model="scope.row.price" placeholder="请输入单价" clearable :precision="2"></el-input-number>
									</template>
								</el-table-column>
								<el-table-column prop="indiscriminately" label="不分扎上报" >
									<template #default="scope">
										<el-switch v-model="scope.row.indiscriminately"></el-switch>
									</template>
								</el-table-column>
								<el-table-column prop="is_completed" label="是否为完成工序" >
									<template #default="scope">
										<el-switch v-model="scope.row.is_completed"></el-switch>
									</template>
								</el-table-column>
							</sc-form-table>
						</el-form-item>
					</el-form>
					<el-form v-if="active==3" ref="stepForm_3" :model="form" :rules="rules" label-position="top">
						<el-form-item label="物料" prop="consumption">
							<table>
								<thead>
  <th>部位</th>
  <th>物料名称</th>
  <th>单价/单位</th>
  <th>单件用量</th>
  <th>全部颜色</th>
  <th v-for="size in form.cloth.sizes" :key="size">{{ size }}</th>
  <th>全部尺码</th>
  <th v-for="color in form.cloth.colors" :key="color">{{ color }}</th>
  <th>操作</th>
</thead>
<tbody>
  <tr v-for="(item, index) in form.consumption" :key="index">
    <td>
      <el-input v-model="item.part" placeholder="请输入部位" clearable></el-input>
    </td>
    <td>
      <el-input v-model="item.name" placeholder="请输入物料名称" clearable></el-input>
    </td>
    <td>
      <el-input-number v-model="item.unit_price" placeholder="请输入单价" clearable :precision="2"></el-input-number>
    </td>
    <td>
      <el-input-number v-model="item.nums" placeholder="请输入单件用量" clearable :precision="2"></el-input-number>
    </td>
    <td>
	  <el-switch @change="allColor(item,index)" v-model="allColors[index]"></el-switch>

		</td>
		<td v-for="color in form.cloth.colors" :key="color">
      <el-switch v-model="item[color]"></el-switch>
    </td>
    <td>
      <el-switch @change="allSize(item,index)" v-model="allSizes[index]"></el-switch>
    </td>
    
	<td v-for="size in form.cloth.sizes" :key="size">
      <el-switch v-model="item[size]"></el-switch>
    </td>
    <td>
      <el-button type="danger" plain icon="el-icon-delete" @click="form.consumption.splice(index, 1)"></el-button>
    </td>
  </tr>
</tbody>
<tfoot>
  <tr>
    <td colspan="11">
      <el-button type="primary" plain icon="el-icon-plus" @click="form.consumption.push(JSON.parse(JSON.stringify(addConsumption)))">添加</el-button>
    </td>
  </tr>
</tfoot>
							</table>
						</el-form-item>
					</el-form>
					<div v-if="active==4">
						<el-result icon="success" title="操作成功" sub-title="您的操作已完成！">
							<template #extra>
								<el-button type="primary" @click="again">再来一单</el-button>
								<el-button @click="$router.go(-1)">返回列表</el-button>
							</template>
						</el-result>
					</div>
					<el-button v-if="active>0 && active<4" @click="pre" :disabled="submitLoading">上一步</el-button>
					<el-button v-if="active<3" type="primary" @click="next">下一步</el-button>
					<el-button v-if="active==3" type="primary" @click="submit" :loading="submitLoading">提交</el-button>
				</el-col>
			</el-row>
		</el-card>
	</el-main>
</template>

<script>


	export default {
		name: 'listCrud-detail',
		data() {
			return {
				id: this.$route.query.id,
				active: 0,
				form:{
					// 基础
					name: '', // 合同名称
					code: '', // 合同编号
					customer: '', // 客户名称
					delivery_date: '', // 交货日期
					order_type: '', // 订单类型
					salesman: '', // 业务员
					remark: '', // 备注
					// 款式
					cloth:{
						name: '', // 款名
						code: '', // 款号
						picture: '', // 图片
						colors: [], // 颜色
						sizes: [], // 尺码
						year: '', // 年份
						season: '', // 季节
						unit_price:0, // 单价
					},
					total: 0, // 总数量
					total_price: 0, // 总金额
					contains:[],
					consumption:[],// 消耗
					// 工序
					procedure: []
				},
				rules:{
					name: [
						{ required: true, message: '请输入合同名称', trigger: 'blur' },
					],
					code: [
						{ required: true, message: '请输入合同编号', trigger: 'blur' },
					],
					customer: [
						{ required: true, message: '请输入客户名称', trigger: 'blur' },
					],
					delivery_date: [
						{ required: true, message: '请选择交货日期', trigger: 'blur' },
					],
					order_type: [
						{ required: true, message: '请选择订单类型', trigger: 'blur' },
					],
					salesman: [
						{ required: true, message: '请选择业务员', trigger: 'blur' },
					],
					remark: [],
					// 款式
					'cloth.name': [
						{ required: true, message: '请输入款名', trigger: 'blur' },
					],
					'cloth.code': [
						{ required: true, message: '请输入款号', trigger: 'blur' },
					],
					'cloth.picture': [],
					'cloth.colors': [
						{ required: true, message: '请选择颜色', trigger: 'blur' },
					],
					'cloth.sizes': [
						{ required: true, message: '请选择尺码', trigger: 'blur' },
					],
					'cloth.year': [
						{ required: true, message: '请输入年份', trigger: 'blur' },
					],
					'cloth.season': [
						{ required: true, message: '请输入季节', trigger: 'blur' },
					],
					'cloth.unit_price': [
						{ required: true, message: '请输入单价', trigger: 'blur' },
					],
					// 工序
					procedure: [
						{ required: true, message: '请添加工序', trigger: 'blur' },
					],
				},
				sizes: [],
				colors: [],
				procedures: [],
				order_types: [],
				salesmans: [],
				customers: [],
				cloths: [],
				addContainsTemplate:{
					size : "",
					color : "",
					nums : 0,
				},
				addProcedureTemplate:{
					name : "",
					sort : 0,
					price : 0,
					indiscriminately : false,
					is_completed : false,
				},
				addConsumption:{
					name:"",
					part:"",
					unit:"",
					nums:0,
					unit_price:0,

				},
				contains: [

				],
				consumption: [

				],
				allColors: [],
				allSizes: [],
			}
		},
		created() {

		},
		watch: {
			// 监听 form.contains 的变化,计算总数量和总金额
			'form.contains': {
				handler: function (val, oldVal) {
					let total = 0
					let total_price = 0
					val.forEach(item=>{
						total += item.nums
						total_price += item.nums * this.form.cloth.unit_price
					})
					this.form.total = total
					this.form.total_price = total_price
				},
				deep: true
			},
			// 监听 form.cloth.unit_price 的变化,计算总金额
			'form.cloth.unit_price': {
				handler: function (val, oldVal) {
					let total_price = 0
					this.form.contains.forEach(item=>{
						total_price += item.nums * val
					})
					this.form.total_price = total_price
				},
				deep: true
			},
			// 监听 form.total 的变化,计算总金额
			'form.total': {
				handler: function (val, oldVal) {
					let total_price = 0
					this.form.contains.forEach(item=>{
						total_price += item.nums * this.form.cloth.unit_price
					})
					this.form.total_price = total_price
				},
				deep: true
			},
			// 监听 form.cloth.colors 的变化,计算 form.contains
			'form.cloth.colors': {
				handler: function (val, oldVal) {
					let oldContains = this.form.contains
					let contains = []
					val.forEach(color=>{
						this.form.cloth.sizes.forEach(size=>{
							contains.push({
								color: color,
								size: size,
								nums: 0
							})
						})
					})
					contains.forEach(item=>{
						let oldItem = oldContains.find(oldItem=>oldItem.color==item.color&&oldItem.size==item.size)
						if(oldItem){
							item.nums = oldItem.nums
						}
					})
					this.form.contains = contains
				},
				deep: true
			},
			// 监听 form.cloth.sizes 的变化,计算 form.contains
			'form.cloth.sizes': {
				handler: function (val, oldVal) {
					let oldContains = this.form.contains
					let contains = []
					this.form.cloth.colors.forEach(color=>{
						val.forEach(size=>{
							contains.push({
								color: color,
								size: size,
								nums: 0
							})
						})
					})
					contains.forEach(item=>{
						let oldItem = oldContains.find(oldItem=>oldItem.color==item.color&&oldItem.size==item.size)
						if(oldItem){
							item.nums = oldItem.nums
						}
					})
					this.form.contains = contains
				},
				deep: true
			},
		},
		mounted() {
			//修改tab名称
			this.$store.commit("updateViewTagsTitle", this.id?`CURD编辑ID:${this.id}`:"CURD新增")
			this.getInfo()
			if(this.id){
				this.getDetail()
			}
		},
		methods: {
			allColor(item,i){
				if(this.allColor[i]){
					this.form.cloth.colors.forEach(color=>{
						this.form.consumption[i][color] = true
					})
					this.allSize[i] = true
				}else{
					this.form.cloth.colors.forEach(color=>{
						this.form.consumption[i][color] = false
					})
					this.allSize[i] = false
				}
				return this.allColor[i]
			},
			allSize(item,i){
				if(this.allSize[i]){
					this.form.cloth.sizes.forEach(size=>{
						item[size] = true
					})
					this.allColor[i] = true
				}else{
					this.form.cloth.sizes.forEach(size=>{
						item[size] = false
					})
					this.allColor[i] = false
				}
				return this.allSize[i]
			},
			addConsumptionBtn(){
				let consumption = JSON.parse(JSON.stringify(this.addConsumption))
				this.form.consumption.push(consumption)
			},
			async getDetail(){
				let res = await this.$API.order.order.show.get({id:this.id})
				if(res.code == 200){
					this.form = res.data.item
				}
			},
			handleContainsChange(val){
				// form.contains 中的 color-size 组合不能重复
				let contains = this.form.contains
				let containsMap = {}
				contains.forEach(item=>{
					if(containsMap[item.color+item.size]){
						this.$notify({
							title: '警告',
							message: '颜色尺码组合不能重复',
							type: 'warning',
							duration: 3500
						});
						item.color = ""
					}else{
						containsMap[item.color+item.size] = true
					}
				})
			},
			handleClothCodeChange(val){
				let cloth = this.cloths.find(item=>item.code == val)
				if(cloth){
					this.form.cloth = cloth
					this.form.procedure = cloth.procedures
				}
			},
			handleClothNameChange(val){
				let cloth = this.cloths.find(item=>item.name == val)
				if(cloth){
					this.form.cloth = cloth
					this.form.procedure = cloth.procedures
				}
			},
			async querySearchCloth(query){
				let params = {
					page:1,
					pageSize:10,
					keywords:query
				}
				let res = await this.$API.order.style.list.get(params)
				if(res.code == 200){
					this.cloths = res.data.items
				}
			},
			generateCode(){
				let date = `${new Date().getFullYear()}${new Date().getMonth()+1}${new Date().getDate()}`.slice(-8)
				this.form.code = `AWS${date}${Math.random().toString().slice(-6)}`
			},
			generateClothCode(){
				let date = `${new Date().getFullYear()}${new Date().getMonth()+1}${new Date().getDate()}`.slice(-8)
				this.form.cloth.code = `AWS${date}${Math.random().toString().slice(-6)}`
			},
			querySearchCustomer(queryString, cb) {
				let results = queryString ? this.customers.filter(this.createFilterCustomer(queryString)) : this.customers;
				cb(results);
			},
			createFilterCustomer(queryString) {
				return (item) => {
					return (item.name.toLowerCase().indexOf(queryString.toLowerCase()) === 0);
				};
			},
			querySearchSalesman(queryString, cb) {
				let results = queryString ? this.salesmans.filter(this.createFilter(queryString)) : this.salesmans;
				cb(results);
			},
			querySearchOrderType(queryString, cb) {
				let results = queryString ? this.order_types.filter(this.createFilter(queryString)) : this.order_types;
				cb(results);
			},
			createFilter(queryString) {
				return (item) => {
					return (item.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0);
				};
			},
			async getInfo(){
				let params = {
					key : "业务员,订单类型,颜色,尺码,工序"
				}
				let res = await this.$API.basic.dic.query.get(params)
				if(res.code == 200){
					this.salesmans = res.data.items['业务员']
					this.order_types = res.data.items['订单类型']
					this.colors = res.data.items['颜色']
					this.sizes = res.data.items['尺码']
					this.procedures = res.data.items['工序']
				}
				let customer  = await this.$API.basic.customer.list.get({page:1,pageSize:1000})
				if(customer.code == 200){
					this.customers = customer.data.items
					console.log(this.customers);
				}
			},
			//下一步
			next(){
				const formName = `stepForm_${this.active}`
				this.$refs[formName].validate((valid) => {
					if (valid) {
						this.active += 1
					}else{
						return false
					}
				})
			},
			//上一步
			pre(){
				this.active -= 1
			},
			//提交
			submit(){
				const formName = `stepForm_${this.active}`
				this.$refs[formName].validate(async (valid) => {
					if (valid) {
						this.submitLoading = true
						var res = null
						if (this.id){
							res = await this.$API.order.order.edit.put(this.form)
						}else{
							res = await this.$API.order.order.add.post(this.form)
						}
						console.log(res);
						this.submitLoading = false
						if (res.code == 200){
							this.$notify({
								title: '成功',
								message: '操作成功',
								type: 'success',
								duration: 3500
							});
							this.active += 1
						}else{
								this.$notify({
									title: '失败',
									message: res.message,
									type: 'error',
									duration: 3500
								});
							}
						}
				})
			},
			//再来一次
			again(){
				this.active = 0
				this.form.code = ""
			}
		}
	}
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
</style>

