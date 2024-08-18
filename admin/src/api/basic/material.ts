import request from '@/utils/request'

// 物料列表
export function materialLists(params?: any) {
    return request.get({ url: '/basic/material/list', params })
}

// 添加物料
export function materialAdd(params: any) {
    return request.post({ url: '/basic/material/add', params })
}

// 编辑物料
export function materialEdit(params: any) {
    return request.post({ url: '/basic/material/edit', params })
}

// 删除物料
export function materialDelete(params: any) {
    return request.post({ url: '/basic/material/del', params })
}

// 物料详情
export function materialDetail(params: any) {
    return request.get({ url: '/basic/material/detail', params })
}

// 物料状态
export function materialChange(params: any) {
    return request.post({ url: '/basic/material/change', params })
}