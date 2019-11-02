var app = new Vue({
    el: '#page_content',
    delimiters: ['{', '}'],
    data: {
        pageCount: 1,colNames: [],objects: [],actions: [],
        customer: {
            Name: '',Telephone: '',Email: '',
            AccountPeriod: '', IsVip: '',
            Status: '',BusinessTypeName: '',
            AuditUser: '',CreateUser: '',
            SaleUser: '',Company: ''
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
            let data = window.selectApi("/customer/get_status", {actionType: actionType}, 1);
            if (data.status) {
                this.select2Data = data.data
            }
        },
        clickCompany: function () {
            this.companyOptions = this.$select2Company()
        },
        clickUser: function () {
            let str = "";
            let company = this.customer.Company;
            if (company !== '' && company !== null) {
                str = "Company:" + company.Id;
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
            let hashParams = {};
            let url = "/customer";
            hashParams["query"] = this.getFilerResult();
            hashParams['page'] = page;
            window.indexData(url, hashParams).then(res => {
                    _this.actions = res.actions;
                    _this.colNames = res.colNames;
                    _this.pageCount = res.countPage;
                    if (res.data !== null && typeof (res.data) !== 'undefined') {
                        _this.objects = res.data;
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
                    if (['IsVip'].indexOf(k) > -1) {
                        let value = 0;
                        if (v.code) {
                            value = 1;
                        }
                        str.push(k + ":" + value);
                    } else if (['AccountPeriod', 'Status'].indexOf(k) > -1) {
                        str.push(k + ":" + v['code']);
                    } else if (['AuditUser', 'CreateUser',
                        'SaleUser', 'Company'].indexOf(k) > -1) {
                        str.push(k+":"+v.Id);
                    } else {
                        str.push(k + ":" + v);
                    }
                }
            });
            console.log(str);
            return str.join(',');

        },
        //清空数据
        refreshResult: function () {
            $('.filter-form')[0].reset();
            this.getList();
            toastr.success("刷新数据成功");
        },
    }
});

