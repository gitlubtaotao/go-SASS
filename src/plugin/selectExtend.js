var selectExtend = {};
selectExtend.install = function (Vue, options) {
    //选择所属公司
    //options: 其他条件
    //page: 当前的页
    Vue.prototype.$select2Company = (options = {}, page = 1) => {
        let result = window.selectApi('/company', $.extend({fields: 'Name,Id'}, options), page);
        if (result.status) {
            if (Array.isArray(result.data.data)) {
                return result.data.data;
            }
        } else {
            return [];
        }
    };

    //选择员工信息
    Vue.prototype.$select2User = function (options = {}, page = 1) {
        let result = selectApi('/user', $.extend({fields: 'Name,Id'}, options), page);
        if (result.status) {
            let options = [];
            if (Array.isArray(result.data.data)) {
                $.each(result.data.data,function (key,value) {
                   options.push({Id: value.Id,Name: value.Name});
                });
                return options;
            } else {
                return []
            }
        }
    };

    //选择对应的部门信息
    Vue.prototype.$select2Department = function (options = {}, page = 1) {
        let result = selectApi('/department', $.extend({fields: "Name,Id"}, options), page);
        if (result.status) {
            if (Array.isArray(result.data.data)) {
                return result.data.data;
            } else {
                return [];
            }
        }
    }
};
module.exports = selectExtend;
