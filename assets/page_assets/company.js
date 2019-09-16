var app = new Vue({
    el: '#page_content',
    delimiters: ['{', '}'],
    data: {
        companyList: []
    },
    mounted: function () {
        this.getList()
    },
    methods: {
        //获取所有的数据
        getList: function () {
            var hashParams = {};
            $.each($('#company_filter').serializeArray(), function (key, value) {
                hashParams[value['name']] = value['value'];
            });
            axios.get('/company', {
                headers: {'X-Requested-With': 'XMLHttpRequest'},
                params: hashParams,
                dataType: 'json',
            }).then(function (response) {
                console.log(response);
                app.companyList = response['data']['data'];
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
            $('#company_filter')[0].reset();
            app.getList();
        },
        editCompany: function(Id,index){
            console.log(Id);

        },
        deleteCompany: function (Id,index) {
            console.log(index);
            if(confirm("确定删除该记录？")){
                axios.delete("/company/"+Id, {
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
