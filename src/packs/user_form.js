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
        this.companySelect();
    },
    methods: {
        //对表单的数据进行验证
        checkForm: function () {
            if(!this.validateSubmit()){
                return;
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
        },
        //提交表单前进行验证
        validateSubmit: function(){
            let regPhone=/^1[3456789]\d{9}$/;
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
            if(!this.user.Pwd){
                this.errors.push('密码必填项');
                return  false;
            }
            if(!regPhone.test(this.user.Phone)){
                this.errors.push('请输入有效的手机号码');
            }
            if(!(this.user.Pwd === this.user.confirmPassword)){
                this.errors.push('两次输入的密码不一致');
                return false;
            }
            return true;
        },
        //所属公司下拉数据
        companySelect: function () {
            let result = selectApi('/company', ['Name', 'Id'], 1);
            if (result.status) {
                if (Array.isArray(result.data.data)) {
                    console.log(result.data);
                    this.companyList = result.data['data'];
                }
            } else {
                console.log(result.data)
            }
        }
    }
});

