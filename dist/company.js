/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, { enumerable: true, get: getter });
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 			Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 		}
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// create a fake namespace object
/******/ 	// mode & 1: value is a module id, require it
/******/ 	// mode & 2: merge all properties of value into the ns
/******/ 	// mode & 4: return value when already ns object
/******/ 	// mode & 8|1: behave like require
/******/ 	__webpack_require__.t = function(value, mode) {
/******/ 		if(mode & 1) value = __webpack_require__(value);
/******/ 		if(mode & 8) return value;
/******/ 		if((mode & 4) && typeof value === 'object' && value && value.__esModule) return value;
/******/ 		var ns = Object.create(null);
/******/ 		__webpack_require__.r(ns);
/******/ 		Object.defineProperty(ns, 'default', { enumerable: true, value: value });
/******/ 		if(mode & 2 && typeof value != 'string') for(var key in value) __webpack_require__.d(ns, key, function(key) { return value[key]; }.bind(null, key));
/******/ 		return ns;
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = "./src/packs/company.js");
/******/ })
/************************************************************************/
/******/ ({

/***/ "./src/packs/company.js":
/*!******************************!*\
  !*** ./src/packs/company.js ***!
  \******************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
eval("\n\nvar app = new Vue({\n  el: '#page_content',\n  delimiters: ['{', '}'],\n  data: {\n    companyList: [],\n    pageCount: 1\n  },\n  mounted: function mounted() {\n    this.getList();\n  },\n  methods: {\n    clickCallback: function clickCallback(pageNum) {\n      this.getList(pageNum);\n    },\n    //获取所有的数据\n    getList: function getList() {\n      var page = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 1;\n      var hashParams = {};\n      $.each($('#company_filter').serializeArray(), function (key, value) {\n        hashParams[value['name']] = value['value'];\n      });\n      hashParams['page'] = page;\n      axios.get('/company', {\n        headers: {\n          'X-Requested-With': 'XMLHttpRequest'\n        },\n        params: hashParams,\n        dataType: 'json'\n      }).then(function (response) {\n        app.companyList = response['data']['data'];\n        app.pageCount = response['data']['pageResult']['CountPage'];\n      })[\"catch\"](function (error) {\n        console.log(error);\n      });\n    },\n    //过滤部分数据\n    filterResult: function filterResult() {\n      app.getList();\n    },\n    //清空数据\n    refreshResult: function refreshResult() {\n      $('#company_filter')[0].reset();\n      app.getList();\n    },\n    editCompany: function editCompany(Id, index) {\n      console.log(Id);\n    },\n    deleteCompany: function deleteCompany(Id, index) {\n      console.log(index);\n\n      if (confirm(\"确定删除该记录？\")) {\n        axios[\"delete\"](\"/company/\" + Id, {\n          headers: {\n            'X-Requested-With': 'XMLHttpRequest'\n          },\n          dataType: 'json'\n        }).then(function (response) {\n          app.companyList.splice(index, 1);\n        })[\"catch\"](function (error) {\n          console.log(error);\n        });\n      }\n    }\n  }\n});\n\n//# sourceURL=webpack:///./src/packs/company.js?");

/***/ })

/******/ });