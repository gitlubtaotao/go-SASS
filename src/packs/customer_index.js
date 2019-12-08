var app = new Vue({
    el: '#page_content',
    delimiters: ['{', '}'],
    data: {
        pageCount: 1, colNames: [], objects: [], actions: [],
        page: 1,
        customer: {
            name__icontains: '', telephone__icontains: '', email__contains: '',
            AccountPeriod: '', IsVip: '',
            Status: '', business_type_name__contains: '',
            AuditUser: '', CreateUser: '',
            SaleUser: '', Company: '',
            created_at__gte: '',created_at__lte: '',
        },
        select2Data: [],
        companyOptions: [],
        userOptions: [],
    },
    i18n: new VueI18n({
        locale: getCookie('lang'),
        messages: Messages,
    }),
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
                str += ("name__icontains:" + search)
            }
            this.companyOptions = this.$select2Company({query: str})
        },
        clickUser: function (search) {
            let str = "";
            let company = this.customer.Company;
            if (company !== '' && company !== null) {
                str = "company.id:" + company;
            }
            if (search) {
                str += (",name__icontains:" + search)
            }
            this.userOptions = this.$select2User({query: str})
        },
        //获取对应的员工数据
        clickCallback: function (pageNum) {
            this.page = pageNum;
            this.getList(pageNum);
        },
        //获取所有的数据
        getList: function () {
            let _this = this;
            let url = "";
            if (location.pathname === "/supplier/index") {
                url = "/supplier";
            } else {
                url = "/customer";
            }
            this.$indexData(url, {query: this.getFilerResult(),page: this.page}).then(res => {
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
            toastr.success(this.$i18n.t("refresh"));
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
            toastr.success(this.$i18n.t("refresh"));
        },
    }
});

