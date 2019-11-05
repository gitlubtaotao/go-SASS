var indexExtend = {};
indexExtend.install = function (Vue, options) {
    //获取index data
    Vue.prototype.$indexData = function (url, options) {
        return new Promise((resolve, reject) => {
            axios.get(url, {
                headers: {'X-Requested-With': 'XMLHttpRequest'},
                params: options,
                dataType: 'json',
            }).then(function (response) {
                console.log(response);
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
                    toastr.success("保存成功");
                    resolve([])
                } else {
                    let data = response.data.msg;
                    if (!Array.isArray(response.data)) {
                        data = [data];
                    }
                    resolve(data)
                }
            }).catch(function (error) {
                reject(error)
            });
        });
    }
};
module.exports = indexExtend;