'use strict';
Vue.component('paginate', VuejsPaginate);
window.$ = $;

$(document).ready(function () {

    $(".form_datetime").datetimepicker({
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
    }).next().on(ace.click_event, function(){
        $(this).prev().focus();
    });
});

// import ss from "./packs/home";
// console.log(ss);



