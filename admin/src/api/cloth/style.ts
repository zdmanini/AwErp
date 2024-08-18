import request from '@/utils/request'

// 款式列表
export function styleLists(params?: any) {
    return request.get({ url: '/cloth/style/list', params })
}

// 添加款式
export function styleAdd(params: any) {
    return request.post({ url: '/cloth/style/add', params })
}

// 编辑款式
export function styleEdit(params: any) {
    return request.post({ url: '/cloth/style/edit', params })
}

// 删除款式
export function styleDelete(params: any) {
    return request.post({ url: '/cloth/style/del', params })
}

// 款式详情
export function styleDetail(params: any) {
    return request.get({ url: '/cloth/style/detail', params })
}

// 款式状态
export function styleChange(params: any) {
    return request.post({ url: '/cloth/style/change', params })
}