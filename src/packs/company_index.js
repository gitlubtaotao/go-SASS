var app = new Vue({
    el: '#page_content',
    delimiters: ['{', '}'],
    data: {
        company: {name__icontains: '', telephone__icontains: '', email__contains: '',
            address__contains: '',created_at__gte:'',created_at__lte: ''},
        pageCount: 1, colNames: [], objects: [], actions: [],page: 1,
    },
    i18n: new VueI18n({
        locale: getCookie('lang'),
        messages: Messages,
    }),
    mounted: function () {
        this.getList();
    },
    methods: {
        clickCallback: function (pageNum) {
            this.page = pageNum;
            this.getList(pageNum);
        },
        //获取所有的数据
        getList: function () {
            let url = "/company";
            let _this = this;
            this.$indexData(url, {query: this.getFilerResult(),page: this.page}).then(res => {
                    _this.actions = res.actions;
                    _this.colNames = res.colNames;
                    _this.pageCount = res.countPage;
                    if (res.data !== null) {
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
        getFilerResult: function () {
            let str = [];
            $.each(this.company, function (k, v) {
                if (v) {
                    str.push(k + ":" + v);
                }
            });
            return str.join(',');
        },
        //清空数据
        refreshResult: function () {
            let _this = this;
            $.each(this.company,function (k,v) {
                _this.$data.company[k] = ""
            });
            this.getList();
            toastr.success(this.$i18n.t("refresh"));
        },
    }
});

