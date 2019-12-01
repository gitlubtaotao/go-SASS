'use strict';

window.Vue = require('vue');
window.axios = require('axios');
import {selectApi,getCookie} from './plugin/data_api';

import Paginate from 'vuejs-paginate';
import vueSelect from 'vue-select';
import indexTable from './components/vue/index_table';
import selectExtend from './plugin/selectExtend';
import indexExtend from './plugin/indexExtend';
import dateExtend from './plugin/dateExtend';
import {Datetime} from 'vue-datetime';
import {Settings} from 'luxon'
import downloadExtend from "./components/vue/download_extend";
import VueI18n from 'vue-i18n'
import toastr from 'toastr';
Settings.defaultLocale = getCookie('lang');

Vue.component('datetime', Datetime);
Vue.component('paginate', Paginate);
Vue.component('vue-select', vueSelect);
Vue.component('index-table', indexTable);
Vue.component('download-extend', downloadExtend);

Vue.use(selectExtend);
Vue.use(indexExtend);
Vue.use(dateExtend);
Vue.use(VueI18n);

//设置全局的实例
window.toastr = toastr;
window.selectApi = selectApi;
window.getCookie = getCookie;
window.VueI18n = VueI18n;




