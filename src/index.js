'use strict';
import $ from 'jquery';
window.$ = $;
require('bootstrap-datepicker/dist/js/bootstrap-datepicker');
require("./plugin/datapicker");
$(document).ready(function () {
    //时间格式化
    let formDatetime = $(".datepicker");
    if (formDatetime.length > 0) {
        $('.datepicker').datepicker({
            format: 'yyyy-mm-dd ',
            language: 'zh-CN',
        }).next().on(ace.click_event, function () {
            $(this).prev().focus();
        });
    }
});




