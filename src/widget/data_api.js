// 获取公司的数据的Api
//selectApi 获取下拉数据
//resourceName: 对应资源url,fields: 需要显示的字段,page: 对应的分页数量

function selectApi(url, fields, page = 1) {
    let returnValue;
    if (Array.isArray(fields)) {
        fields = fields.join(',')
    } else {
        return {
            status: false,
            data: 'fields必须为数组形式'
        };
    }
    if (url === "") {
        return {
            status: false,
            data: 'url不能为空'
        }
    }
    $.ajax({
        url: url,
        data: {fields: fields, page: page},
        dataType: 'json',
        async: false,
        type: 'get',
        headers: {'X-Requested-With': 'XMLHttpRequest'},
        success: function (data) {
            returnValue = {
                status: true,
                data: data
            }
        },
        error: function (data) {
            returnValue = {
                status: false,
                data: data
            }
        },
    });
    return returnValue;
}

export {selectApi};

