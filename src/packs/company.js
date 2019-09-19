var app = new Vue({
    el: '#page_content',
    delimiters: ['{', '}'],
    data: {
        companyList: [],
        pageCount: 1
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
            let hashParams = {};
            let currentFilter = $('.filter-form');
            $.each(currentFilter.serializeArray(), function (key, value) {
                hashParams[value['name']] = value['value'];
            });
            hashParams['page'] = page;
            axios.get(currentFilter.attr('action'), {
                headers: {'X-Requested-With': 'XMLHttpRequest'},
                params: hashParams,
                dataType: 'json',
            }).then(function (response) {
                console.log(response);
                if (response['data'] !== null) {
                    app.companyList = response['data']['data'];
                    app.pageCount = response['data']['countPage']
                }
            }).catch(function (error) {
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
        editCompany: function (Id, index) {
            console.log(Id);

        },
        deleteCompany: function (Id, index) {
            console.log(index);
            if (confirm("确定删除该记录？")) {
                axios.delete("/company/" + Id, {
                    headers: {'X-Requested-With': 'XMLHttpRequest'},
                    dataType: 'json',
                }).then(function (response) {
                    app.companyList.splice(index, 1)
                }).catch(function (error) {
                    console.log(error);
                });
            }
        },
    }
});

