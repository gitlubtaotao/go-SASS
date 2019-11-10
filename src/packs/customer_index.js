var app = new Vue({
    el: '#page_content',
    delimiters: ['{', '}'],
    data: {
        pageCount: 1, colNames: [], objects: [], actions: [],
        customer: {
            Name: '', Telephone: '', Email: '',
            AccountPeriod: '', IsVip: '',
            Status: '', BusinessTypeName: '',
            AuditUser: '', CreateUser: '',
            SaleUser: '', Company: ''
        },
        select2Data: [],
        companyOptions: [],
        userOptions: [],
    },
    mounted: function () {
        this.getList();
    },
    methods: {
        //获取select2数据
        select2Method: function (actionType) {
            this.select2Data = window.selectApi("/customer/status", {actionType: actionType}, 1);
        },
        clickCompany: function (search) {
            let str = "";
            if (search) {
                str += ("Name:" + search)
            }
            this.companyOptions = this.$select2Company({query: str})
        },
        clickUser: function (search) {
            let str = "";
            let company = this.customer.Company;
            if (company !== '' && company !== null) {
                str = "Company:" + company;
            }
            if (search) {
                str += (",Name:" + search)
            }
            this.userOptions = this.$select2User({query: str})
        },
        //获取对应的员工数据
        clickCallback: function (pageNum) {
            this.getList(pageNum);
        },
        //获取所有的数据
        getList: function (page = 1) {
            let _this = this;
            let url = "";
            if (location.pathname === "/supplier/index") {
                url = "/supplier";
            } else {
                url = "/customer";
            }
            console.log(this.getFilerResult());
            this.$indexData(url, {query: this.getFilerResult()}).then(res => {
                    _this.actions = res.actions;
                    _this.colNames = res.colNames;
                    _this.pageCount = res.countPage;
                    if (res.data !== null && typeof (res.data) !== 'undefined') {
                        _this.objects = res.data;
                    } else {
                        _this.objects = [];
                    }

                },
                error => {
                    console.log(error);
                });
        },
        //过滤部分数据
        filterResult: function () {
            this.getList();
            toastr.success("刷新数据成功");
        },
        //对form 表单对数据进行过滤
        getFilerResult: function () {
            let str = [];
            $.each(this.customer, function (k, v) {
                if (v) {
                    str.push(k + ":" + v);
                }
            });
            return str.join(',');
        },
        //清空数据
        refreshResult: function () {
            let _this = this;
            $.each(this.customer, function (k, v) {
                _this.$data.customer[k] = "";
            });
            this.getList();
            toastr.success("刷新数据成功");
        },
    }
});

