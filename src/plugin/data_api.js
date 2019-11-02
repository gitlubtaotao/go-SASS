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
            if (response.data === 'OK') {
                toastr.success("保存成功");
                resolve([])
            } else {
                if (Array.isArray(response.data)) {
                    resolve(response.data[0]);
                } else {
                    resolve(response.data);
                }
            }
        }).catch(function (error) {
            reject(error)
        });
    });
}

//initData: form表单初始化数据
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

//初始化index对应的数据
function indexData(url, hashParams) {
    return new Promise((resolve, reject) => {
        axios.get(url, {
            headers: {'X-Requested-With': 'XMLHttpRequest'},
            params: hashParams,
            dataType: 'json',
        }).then(function (response) {
            if (response['data'] !== null) {
                resolve(response['data'])
            }
        }).catch(function (error) {
            reject(error);
        });
    });
}

export {selectApi, submitForm, initData,indexData};

