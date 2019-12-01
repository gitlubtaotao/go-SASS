var indexExtend = {};
indexExtend.install = function (Vue, options) {
    //获取index data
    Vue.prototype.$indexData = function (url, options) {
        return new Promise((resolve, reject) => {
            axios.get(url, {
                headers: {'X-Requested-With': 'XMLHttpRequest'},
                params: $.extend({page: 1}, options),
                dataType: 'json',
            }).then(function (response) {
                if (response.data.code === 200) {
                    resolve(response.data.obj);
                } else {
                    resolve(response.data.msg);
                }
            }).catch(function (error) {
                reject(error);
            });
        });
    };
    Vue.prototype.$showInitData = function (url) {
        return new Promise((resolve, reject) => {
            axios({
                method: 'get',
                url: url
            }).then(function (response) {
                if (response.data.code === 200) {
                    resolve(response.data.obj);
                } else {
                    resolve(response.data.msg);
                }
            }).catch(function (error) {
                reject(error);
            });
        });
    };
    Vue.prototype.$submitFormData = function (method, url, data) {
        return new Promise((resolve, reject) => {
            axios({
                method: method,
                url: url,
                dataType: 'json',
                data: data
            }).then(function (response) {
                if (response.data.code === 200) {
                    toastr.success(response.data.msg);
                    resolve([])
                } else {
                    let str = '';
                    let data = response.data.msg;
                    if (!Array.isArray(response.data)) {
                        data = [data];
                    }
                    for (let i = 0; i <= data.length; i++) {
                        $.each(data[i], function (k, v) {
                            str += (k + ":" + v);
                        });
                    }
                    toastr.error(str);
                }
            }).catch(function (error) {
                console.log(error);
                toastr.error(error);
            });
        });
    }
};
module.exports = indexExtend;