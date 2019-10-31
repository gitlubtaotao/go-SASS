'use strict';
window.Vue = require('vue');
import Paginate from 'vuejs-paginate';
import vueSelect from 'vue-select'
import indexTable from './components/vue/index_table';
import selectExtend from './plugin/selectExtend';
window.axios = require('axios');
Vue.component('paginate', Paginate);
Vue.component('vue-select',vueSelect);
Vue.component('index-table', indexTable);
Vue.use(selectExtend);

import toastr from 'toastr';
window.toastr = toastr;

import {selectApi,submitForm,initData,indexData} from './plugin/data_api';
window.selectApi = selectApi;
window.submitForm = submitForm;
window.initData = initData;
window.indexData = indexData;
