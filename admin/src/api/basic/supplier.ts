import request from '@/utils/request'

// 供应商列表
export function supplierLists(params?: any) {
    return request.get({ url: '/basic/supplier/list', params })
}

// 添加供应商
export function supplierAdd(params: any) {
    return request.post({ url: '/basic/supplier/add', params })
}

// 编辑供应商
export function supplierEdit(params: any) {
    return request.post({ url: '/basic/supplier/edit', params })
}

// 删除供应商
export function supplierDelete(params: any) {
    return request.post({ url: '/basic/supplier/del', params })
}

// 供应商详情
export function supplierDetail(params: any) {
    return request.get({ url: '/basic/supplier/detail', params })
}

// 供应商状态
export function supplierChange(params: any) {
    return request.post({ url: '/basic/supplier/change', params })
}