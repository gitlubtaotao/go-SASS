var app = new Vue({
    el: '#page_content',
    delimiters: ['{', '}'],
    data: {
        pageCount: 1,
        //col names
        colNames: [],
        //objects
        objects: [],
        //对应的actions
        actions: [],
        customer: {
            Name: '',
            Telephone: '',
            Email: '',
            AccountPeriod: '',
            Aging: '',
            Amount: '',
            IsVip: '',
            Status: '',
            BusinessTypeName: '',
            AuditUser: '',
            CreateUser: '',
            SaleUser: '',
            Company: ''

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
            let currentFilter = $('.filter-form');
            let url = '';
            if (currentFilter.length !== 0) {
                $.each(currentFilter.serializeArray(), function (key, value) {
                    hashParams[value['name']] = value['value'];
                });
                url = currentFilter.attr('action');
            } else {
                url = $(this.$el).attr('data-url');
            }
            hashParams['page'] = page;
            window.indexData(url, hashParams).then(res => {
                console.log(res);
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
        },
        //清空数据
        refreshResult: function () {
            $('.filter-form')[0].reset();
            this.getList();
        },
    }
});

