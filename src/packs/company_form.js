var app = new Vue({
    el: '#create_form',
    delimiters: ['{', '}'],
    data: {
        create_action: '/company',
        update_action: '',
        company: {
            Name: '',
            Telephone: '',
            Email: '',
            Address: '',
            Website: '',
            Remarks: '',
        },
        errors: []
    },
    methods: {
        //提交表单
        submitForm: function () {
            if(this.validateForm()){
                axios({
                    method: 'post',
                    url: this.create_action,
                    data: this.company,
                }).then(function (response) {
                    console.log(response);
                    if (response.status === 200) {
                        $.each(response.data,function (key,value) {
                            app.errors.push(key+':'+value);
                        });
                    } else {
                        toastr.success("新增公司成功");
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            }
        },
        //验证form
        validateForm: function () {
            if(!this.company.Name){
                this.errors.push("公司名称为必填项");
                return false;
            }
            if(!this.company.Telephone){
                this.errors.push("公司电话为必填项");
                return false;
            }
            if(!this.company.Email){
                this.errors.push("公司邮箱为必填项");
                return false;
            }
            return true
        }

    }

});