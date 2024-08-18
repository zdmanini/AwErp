import request from '@/utils/request'

// 客户列表
export function customerLists(params?: any) {
    return request.get({ url: '/basic/customer/list', params })
}

// 添加客户
export function customerAdd(params: any) {
    return request.post({ url: '/basic/customer/add', params })
}

// 编辑客户
export function customerEdit(params: any) {
    return request.post({ url: '/basic/customer/edit', params })
}

// 删除客户
export function customerDelete(params: any) {
    return request.post({ url: '/basic/customer/del', params })
}

// 客户详情
export function customerDetail(params: any) {
    return request.get({ url: '/basic/customer/detail', params })
}

// 客户状态
export function customerChange(params: any) {
    return request.post({ url: '/basic/customer/change', params })
}