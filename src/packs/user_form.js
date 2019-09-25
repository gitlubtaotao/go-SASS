console.log("测试");
var app = new Vue({
    el: '#create_form',
    delimiters: ['{', '}'],
    data: {
        action: '/user',
        user: {
            Id: 0,
            Name: '',//姓名
            Email: '',//邮箱
            Phone: '',//电话
            Password: '',//密码
            confirmPassword: '',//确认密码
            Gender: '', // 性别
            Company: '',// 所属公司
            EntryTime: '', //入职时间
        },
        errors: [], //错误
    },
    methods: {
        //对表单的数据进行验证
        checkForm: function () {
            console.log('sdsdsd');
            if (!this.user.Name) {
                this.errors.push("姓名必填项");
                return false;
            }
            this.errors = [];
            axios({
                method: 'post',
                url: this.action + "?companyId=" + this.user.Company,
                data: this.user
            }).then(function (response) {
                console.log(response);
                if (response.status === 200) {
                    app.errors.push(response.data);
                } else {
                    toastr.success("新增员工成功");
                }
            }).catch(function (error) {
                console.log(error);
            });
        }
    }

});