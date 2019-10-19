'use strict';
window.$ = $;
import Paginate from 'vuejs-paginate';
Vue.component('paginate', Paginate);
// index table
import indexTable from './widget/vue/index_table';
Vue.component('index-table', indexTable);

window.axios = require('axios');

import toastr from 'toastr';

window.toastr = toastr;


require('./widget/common');
$(document).ready(function () {
    //时间格式化
    let formDatetime = $(".form_datetime");
    if (formDatetime.length > 0) {
        formDatetime.datetimepicker({
            format: 'YYYY-MM-DD',
            icons: {
                time: 'fa fa-clock-o',
                date: 'fa fa-calendar',
                up: 'fa fa-chevron-up',
                down: 'fa fa-chevron-down',
                previous: 'fa fa-chevron-left',
                next: 'fa fa-chevron-right',
                today: 'fa fa-arrows ',
                clear: 'fa fa-trash',
                close: 'fa fa-times'
            }
        }).next().on(ace.click_event, function () {
            $(this).prev().focus();
        });
    }
    let select = $('select');
    // if (select.length > 0) {
    //     //     select.select2({
    //     //         placeholder: '请选择'
    //     //     });
    //     // }
});




