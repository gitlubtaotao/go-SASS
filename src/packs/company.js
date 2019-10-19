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
            let hashParams = {};
            let currentFilter = $('.filter-form');
            let url='';
            if (currentFilter.length !== 0) {
                $.each(currentFilter.serializeArray(), function (key, value) {
                    hashParams[value['name']] = value['value'];
                });
                url =currentFilter.attr('action');
            }else{
                url = $(this.$el).attr('data-url');
            }
            hashParams['page'] = page;
            axios.get(url, {
                headers: {'X-Requested-With': 'XMLHttpRequest'},
                params: hashParams,
                dataType: 'json',
            }).then(function (response) {
                console.log(response);
                if (response['data'] !== null) {
                    app.objects = response['data']['data'];
                    app.pageCount = response['data']['countPage'];
                    app.colNames = response['data']['colNames'];
                    if (response['data']['actions'] !== "") {
                        app.actions = response['data']["actions"];
                    }
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
            console.log(this.editUrl + Id);
            location.href = this.editUrl + Id;
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

