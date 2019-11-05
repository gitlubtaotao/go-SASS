console.log('sdsdsds');
var app = new Vue({
    el: '#login_in',
    data: {
        userName: '',
        Password: '',
        Remember: false,
        action: '/login'
    },
    methods: {
        submitForm: function () {
            if (!this.userName) {
                toastr.error('账号不能为空');
                return false;
            }
            if (!this.Password) {
                toastr.error('密码不能为空');
                return false;
            }
            axios({
                method: 'post',
                url: this.action,
                data: {
                    UserName: this.userName,
                    Password: this.Password,
                    Remember: this.Remember
                },
            }).then(function (response) {
                if (response.data.code === 200) {
                    toastr.success('登录成功');
                    location.href = response.data.obj;
                } else {
                    toastr.error(response.data.msg);
                }
            }).catch(function (error) {
                console.log(error);
            });
        }
    }
});