
var dateExtend = {};
dateExtend.install = function (Vue, options) {
    //获取index data
    Vue.prototype.$dateFormat = function (time,format = 'YYYY-MM-DD') {
       return moment(new Date(time)).format(format);
    };
};
module.exports = dateExtend;