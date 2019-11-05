// 获取公司的数据的Api
//selectApi 获取下拉数据
//resourceName: 对应资源url,options: 其他options,page: 对应的分页数量

function selectApi(url, options = {}, page = 1) {
    let returnValue = {};
    if (url === "") {
    }
    $.ajax({
        url: url,
        data: $.extend(options, {page: page}),
        dataType: 'json',
        async: false,
        type: 'get',
        headers: {'X-Requested-With': 'XMLHttpRequest'},
        success: function (data) {
            if (data.code === 200) {
                returnValue = data.obj
            }
        },
        error: function (data) {
            console.log(data);
        },
    });
    return returnValue
}


export {selectApi};

