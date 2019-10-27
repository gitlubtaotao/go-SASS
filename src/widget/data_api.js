// 获取公司的数据的Api
//selectApi 获取下拉数据
//resourceName: 对应资源url,options: 其他options,page: 对应的分页数量

function selectApi(url, options={}, page = 1) {
    let returnValue ={};
    if (url === "") {
        return {
            status: false,
            data: 'url不能为空'
        }
    }
    $.ajax({
        url: url,
        data: $.extend(options,{page: page}),
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

