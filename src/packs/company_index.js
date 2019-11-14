var app = new Vue({
    el: '#page_content',
    delimiters: ['{', '}'],
    data: {
        company: {name__icontains: '', telephone__icontains: '', email__contains: '', address__contains: ''},
        pageCount: 1, colNames: [], objects: [], actions: [],
    },
    mounted: function () {
        this.getList();
    },
    methods: {
        clickCallback: function (pageNum) {
            this.getList(pageNum);
        },
        //获取所有的数据
        getList: function () {
            let url = "/company";
            let _this = this;
            this.$indexData(url, {query: this.getFilerResult()}).then(res => {
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
            toastr.success("刷新数据成功");
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
            toastr.success("刷新数据成功");
        },
    }
});

