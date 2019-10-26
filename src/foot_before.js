'use strict';
window.Vue = require('vue');
import Paginate from 'vuejs-paginate';
import vueSelect from 'vue-select'
import indexTable from './widget/vue/index_table';
window.axios = require('axios');
Vue.component('paginate', Paginate);
Vue.component('vue-select',vueSelect);
Vue.component('index-table', indexTable);

import toastr from 'toastr';
window.toastr = toastr;

import {selectApi} from './widget/data_api';
window.selectApi = selectApi;