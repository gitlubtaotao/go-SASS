// 获取公司的数据的Api
//selectApi 获取下拉数据
//resourceName: 对应资源url,options: 其他options,page: 对应的分页数量

function selectApi(url, options = {}, page = 1) {
    let returnValue = {};
    if (url === "") {
        return "url不能为空"
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
                returnValue = data.obj;
            }
        },
        error: function (data) {
            console.log(data);
        },
    });
    return returnValue
}

function getCookie(cname) {
    let name = cname + "=";
    let ca = document.cookie.split(';');
    for(var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) === ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name)  === 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

export {selectApi,getCookie};

