'use strict';

window.Vue = require('vue');
window.axios = require('axios');

import Paginate from 'vuejs-paginate';
import vueSelect from 'vue-select';
import indexTable from './components/vue/index_table';
import selectExtend from './plugin/selectExtend';
import indexExtend from './plugin/indexExtend';
import dateExtend from './plugin/dateExtend';
import { Datetime } from 'vue-datetime';
import { Settings } from 'luxon'
import downloadExtend from "./components/vue/download_extend";
Settings.defaultLocale = 'zh-CN';
Vue.component('datetime', Datetime);
Vue.component('paginate', Paginate);
Vue.component('vue-select',vueSelect);
Vue.component('index-table', indexTable);
Vue.component('download-extend',downloadExtend);

Vue.use(selectExtend);
Vue.use(indexExtend);
Vue.use(dateExtend);
import toastr from 'toastr';
window.toastr = toastr;

import {selectApi} from './plugin/data_api';
window.selectApi = selectApi;




