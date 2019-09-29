require('../widget/common');
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
            Pwd: '',//密码
            confirmPassword: '',//确认密码
            Gender: '', // 性别
            Company: '',// 所属公司
            EntryTime: '', //入职时间
        },
        errors: [], //错误
        companyList: [],
    },
    mounted: function () {
        this.selectData();
        this.initData();

    },
    methods: {
        //对表单的数据进行验证
        checkForm: function () {
            this.errors = [];
            // if (!this.validateSubmit()) {
            //     return ;
            // }
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
        },
        //提交表单前进行验证
        validateSubmit: function () {
            let regPhone = /^1[3456789]\d{9}$/;
            if (!this.user.Name) {
                this.errors.push("姓名必填项");
                return false;
            }
            if (!this.user.Email) {
                this.errors.push("邮箱必填项");
                return false;
            }
            if (!this.user.Phone) {
                this.errors.push("电话必填项");
                return false;
            }
            console.log(this.user.Id);
            if (this.user.Id === 0) {
                if(!this.validatePassword()){
                    return false;
                }
            }else if(this.user.Id !== 0 && this.user.Pwd !==''){
                if(!this.validatePassword()){
                    return false;
                }
            }
            if (!regPhone.test(this.user.Phone)) {
                this.errors.push('请输入有效的手机号码');
            }

            return true;
        },
        validatePassword: function(){
            if (!this.user.Pwd) {
                this.errors.push('密码必填项');
                return false;
            }
            if (!this.user.confirmPassword) {
                this.errors.push("确认密码必填项");
                return false;
            }
            if (!(this.user.Pwd === this.user.confirmPassword)) {
                this.errors.push('两次输入的密码不一致');
                return false;
            }
            return true;
        },
        //所属公司下拉数据
        selectData: function () {
            let result = selectApi('/company', ['Name', 'Id'], 1);
            if (result.status) {
                if (Array.isArray(result.data.data)) {
                    console.log(result.data);
                    this.companyList = result.data['data'];
                }
            } else {
                console.log(result.data)
            }
        },
        //修改时初始化数据
        initData: function () {
            let id = $('#create_form').attr('data-id');
            if (id !== "" && id !== null) {
                console.log(id);
                this.user.Id = parseInt(id);
                axios({
                    method: 'get',
                    url: "/user/" + id,
                    data: this.user
                }).then(function (response) {
                    console.log(response);
                    let data = response.data;
                    app.user.Name = data['Name'];
                    app.user.Email = data['Email'];
                    app.user.Phone = data['Phone'];
                    app.user.Id = data['Id'];
                    app.user.Gender = data['Gender'];
                    app.user.EntryTime = data['EntryTime'];
                    app.user.Company = data['Company']['Id'];
                }).catch(function (error) {
                    console.log(error);
                });
            }
        }

    }
});

