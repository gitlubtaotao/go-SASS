var app = new Vue({
    el: '#page_content',
    delimiters: ['{', '}'],
    data: {
        pageCount: 1, colNames: [], objects: [], actions: [],
        index: {
            name__icontains: "", Department: '', email__contains: '', telephone__icontains: '',
            Company: '', Gender: '', phone__contains: '', entry_time__lte: '',entry_time__gte: '',
            created_at__lte: "",created_at__gte: '',
        },
        url: '', DepartmentArray: [], CompanyArray: [],
    },
    mounted: function () {
        this.setUrl();
        this.getList();

    },
    methods: {
        clickDepartment: function (search) {
            let str = "company.id:" + this.index.Company;
            if (search) {
                str += ("name__icontains:" + search)
            }
            this.DepartmentArray = this.$select2Department({"query": str})
        },
        clickCompany: function (search) {
            let str = "";
            if (search) {
                str += ("name__icontains:" + search)
            }
            this.CompanyArray = this.$select2Company({"query": str});
        },
        setUrl: function () {
            let url = location.pathname;
            this.url = url.slice(0, url.length - 6);
        },
        clickCallback: function (pageNum) {
            this.getList(pageNum);
        },
        //获取所有的数据
        getList: function (page = 1) {
            let _this = this;
            this.$indexData(this.url, {"query": this.getQueryStr(),page: page}).then(res => {
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
        getQueryStr: function () {
            let str = [];
            $.each(this.index, function (k, v) {
                if (v !== "") {
                    str.push(k + ":" + v)
                }
            });
            return str.join(',');
        },
        //清空数据
        refreshResult: function () {
            let _this = this;
            $.each(this.index, function (k, v) {
                _this.$data.index[k] = '';
            });
            this.getList();
            toastr.success("刷新数据成功");

        },
    }
});

