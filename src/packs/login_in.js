console.log('sdsdsds');
var app = new Vue({
    el: '#login_in',
    data:{
        userName: '',
        Password: '',
        Remember: false,
        action: '/login'
    },
    methods:{
        submitForm: function () {
            if(!this.userName){
                toastr.error('账号不能为空');
                return false;
            }
            if(!this.Password){
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
                console.log(response);
                if (response.data.status) {
                    toastr.success('登录成功');
                    location.href = response.data['url'];
                } else {
                    toastr.error(response.data['message']);
                }
            }).catch(function (error) {
                console.log(error);
            });
        }
    }
});