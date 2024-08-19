import request from '@/utils/request'

// 訂單列表
export function orderLists(params?: any) {
    return request.get({ url: '/cloth/order/list', params })
}

// 添加訂單
export function orderAdd(params: any) {
    return request.post({ url: '/cloth/order/add', params })
}

// 编辑訂單
export function orderEdit(params: any) {
    return request.post({ url: '/cloth/order/edit', params })
}

// 删除訂單
export function orderDelete(params: any) {
    return request.post({ url: '/cloth/order/del', params })
}

// 訂單详情
export function orderDetail(params: any) {
    return request.get({ url: '/cloth/order/detail', params })
}

// 訂單状态
export function orderChange(params: any) {
    return request.post({ url: '/cloth/order/change', params })
}