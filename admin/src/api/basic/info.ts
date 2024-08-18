import request from '@/utils/request'

// 字典项列表
export function itemLists(params?: any) {
    return request.get({ url: '/basic/info/item', params })
}

// 字典类型列表
export function typeLists(params?: any) {
    return request.get({ url: '/basic/info/type', params })
}

// 字典查询
export function infoQuery(params?: any) {
    return request.get({ url: '/basic/info/query', params })
}

// 添加字典项
export function infoAdd(params: any) {
    return request.post({ url: '/basic/info/add', params })
}

// 编辑字典项
export function infoEdit(params: any) {
    return request.post({ url: '/basic/info/edit', params })
}

// 删除字典项
export function infoDelete(params: any) {
    return request.post({ url: '/basic/info/del', params })
}
// 字典类型添加
export function typeAdd(params: any) {
    return request.post({ url: '/basic/info/addType', params })
}

// 编辑字典类型
export function typeEdit(params: any) {
    return request.post({ url: '/basic/info/editType', params })
}

// 删除字典类型
export function typeDelete(params: any) {
    return request.post({ url: '/basic/info/delType', params })
}

// 字典项状态
export function infoChange(params: any) {
    return request.post({ url: '/basic/info/change', params })
}
