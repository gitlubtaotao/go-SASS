'use strict';
import $ from 'jquery';
import {Settings} from "luxon";
window.$ = $;
require("./plugin/datapicker");

$(document).ready(function () {
    //选择语言
    $(document).on('click', '.lang-changed', function () {
        let $e = $(this);
        let lang = $e.data('lang');
        let d = new Date();
        d.setTime(d.getTime() + (24 * 60 * 60 * 1000));
        Settings.defaultLocale = lang;
        let expires = "expires="+d.toUTCString();
        document.cookie = "lang=" + lang + ";" + expires + ";path=/";
        window.location.reload();
    });
});






