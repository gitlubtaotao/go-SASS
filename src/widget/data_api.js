// 获取公司的数据的Api
//selectApi 获取下拉数据
//resourceName: 对应资源url,options: 其他options,page: 对应的分页数量

function selectApi(url, options = {}, page = 1) {
    let returnValue = {};
    if (url === "") {
        return {
            status: false,
            data: 'url不能为空'
        }
    }
    $.ajax({
        url: url,
        data: $.extend(options, {page: page}),
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

//method: 对应的method
//url: url
function submitForm(method, url, data) {
    return new Promise((resolve, reject) => {
        axios({
            method: method,
            url: url,
            dataType: 'json',
            data: data
        }).then(function (response) {
            if(response.data === 'OK'){
                toastr.success("保存成功");
                resolve([])
            }else{
                if(Array.isArray(response.data)){
                    resolve(response.data);
                }else{
                    resolve([response.data]);
                }

            }
        }).catch(function (error) {
            reject(error)
        });
    });
}

function initData(url) {
    return new Promise((resolve, reject) => {
        axios({
            method: 'get',
            url: url
        }).then(function (response) {
            resolve(response.data);
        }).catch(function (error) {
            reject(error);
        });
    });
}

//初始化form对应的数据


export {selectApi, submitForm, initData};

