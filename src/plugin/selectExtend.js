var selectExtend = {};
selectExtend.install = function (Vue, options) {
    //选择所属公司
    //options: 其他条件
    //page: 当前的页
    Vue.prototype.$select2Company = (options = {}, page = 1) => {
        let result = window.selectApi('/company', $.extend({fields: 'Name,Id'}, options), page);
        if (Array.isArray(result.data)) {
            return result.data;
        } else {
            return []
        }
    };

    //选择员工信息
    Vue.prototype.$select2User = function (options = {}, page = 1) {
        let result = selectApi('/user', $.extend({fields: 'Name,Id'}, options), page);
        let return_result = [];
        if (Array.isArray(result.data)) {
            $.each(result.data, function (key, value) {
                return_result.push({Id: value.Id, Name: value.Name});
            });
            return return_result;
        } else {
            return []
        }
    };

    //选择对应的部门信息
    Vue.prototype.$select2Department = function (options = {}, page = 1) {
        let result = selectApi('/department', $.extend({fields: "Name,Id"}, options), page);
        if (Array.isArray(result.data)) {
            return result.data;
        } else {
            return [];
        }
    };

};
module.exports = selectExtend;
