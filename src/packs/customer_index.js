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
    },
    mounted: function () {
        this.getList();
    },
    methods: {
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
                    _this.actions = res.actions;
                    _this.colNames = res.colNames;
                    _this.pageCount = res.countPage;
                    _this.objects = res.data;
                },
                error => {
                    console.log(error);
                });
        },
        //过滤部分数据
        filterResult: function () {
            app.getList();
        },
        //清空数据
        refreshResult: function () {
            $('.filter-form')[0].reset();
            app.getList();
        },
    }
});

